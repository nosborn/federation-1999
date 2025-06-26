package database

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"slices"
	"sync"
	"time"
	"unsafe"

	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/monitoring"
	"github.com/nosborn/federation-1999/internal/server/build"
	"github.com/nosborn/federation-1999/internal/server/global"
	"github.com/nosborn/federation-1999/internal/server/jobs"
	"github.com/nosborn/federation-1999/internal/text"
	"github.com/nosborn/federation-1999/pkg/ibgames"
	"golang.org/x/sys/unix"
)

const (
	dbPersonaSize = 8144 // historically calculated as sizeof(dbpersona_t)

	// Round the page size up to a multiple of 4K.
	// This doesn't matter if not mmap()'ing the persona file.
	dbPageSize = (dbPersonaSize + 0x0FFF) &^ 0x0FFF
)

type SaveWhen int

const (
	SaveNow SaveWhen = iota
	SaveLater
)

type Offset int64

type dbPage [dbPageSize]byte

type Database struct {
	elist      *os.File
	file       *os.File
	nextOffset Offset // Next offset where a new page will be written
	idTable    map[uint32]DatabasePlayer

	queueMu       sync.Mutex
	fastSaveQueue []DatabasePlayer
	slowSaveQueue []DatabasePlayer
	idleTimer     *time.Timer
}

type DatabasePlayer interface {
	IsDeceased() bool
	Name() string
	Offset() Offset
	Serialize() *model.DBPersona
	UID() ibgames.AccountID
}

var database *Database

func init() {
	if (dbPageSize % 4096) != 0 {
		log.Panicf("dbPageSize is not a multiple of 4K: %d", dbPageSize)
	}
	database = &Database{}
}

// updateQueueMetrics updates the Prometheus metrics for queue sizes
func (db *Database) updateQueueMetrics() {
	monitoring.DatabaseFastSaveQueueSize.Set(float64(len(db.fastSaveQueue)))
	monitoring.DatabaseSlowSaveQueueSize.Set(float64(len(db.slowSaveQueue)))
}

// OpenDatabase opens the global persona database
func OpenDatabase(loadPlayer func(*model.DBPersona, Offset) DatabasePlayer) error {
	return database.Open(loadPlayer)
}

func Modify(player DatabasePlayer, when SaveWhen) {
	database.Modify(player, when)
}

func CommitDatabase() {
	database.Commit()
}

func CloseDatabase() {
	database.Close()
}

func NextPersonaOffset() Offset {
	return database.nextPersonaOffset()
}

// PackDBPersona packs a DBPersona back into raw page format with union handling
//
//nolint:gosec // Audited: entire function uses unsafe for 1999 C binary format compatibility
func packDBPersona(persona *model.DBPersona) *dbPage {
	page := &dbPage{}
	data := page[:dbPersonaSize]

	// Personal info.
	copy(data[0:16], persona.Name[:])
	*(*uint32)(unsafe.Pointer(&data[16])) = persona.ID
	data[20] = persona.Sex
	*(*uint32)(unsafe.Pointer(&data[24])) = persona.Rank
	copy(data[28:180], persona.Desc[:])
	copy(data[180:216], persona.Mood[:])

	// Statistics.
	*(*uint32)(unsafe.Pointer(&data[216])) = persona.MaxStr
	*(*uint32)(unsafe.Pointer(&data[220])) = persona.CurStr
	*(*uint32)(unsafe.Pointer(&data[224])) = persona.MaxSta
	*(*uint32)(unsafe.Pointer(&data[228])) = persona.CurSta
	*(*uint32)(unsafe.Pointer(&data[232])) = persona.MaxInt
	*(*uint32)(unsafe.Pointer(&data[236])) = persona.CurInt
	*(*uint32)(unsafe.Pointer(&data[240])) = persona.MaxDex
	*(*uint32)(unsafe.Pointer(&data[244])) = persona.CurDex
	*(*uint32)(unsafe.Pointer(&data[248])) = persona.Shipped
	*(*uint32)(unsafe.Pointer(&data[252])) = persona.Games
	*(*uint32)(unsafe.Pointer(&data[256])) = persona.Flags[0]
	*(*uint32)(unsafe.Pointer(&data[260])) = persona.Flags[1]

	// Money, etc.
	*(*int32)(unsafe.Pointer(&data[264])) = persona.Balance
	*(*int32)(unsafe.Pointer(&data[268])) = persona.Loan
	*(*int32)(unsafe.Pointer(&data[272])) = persona.Reward
	*(*[4]int32)(unsafe.Pointer(&data[276])) = persona.Frame

	// Locations.
	*(*uint32)(unsafe.Pointer(&data[292])) = persona.LocNo
	copy(data[296:312], persona.StarSystem[:])

	// Work.
	*(*model.DBCargo)(unsafe.Pointer(&data[312])) = persona.Job.Pallet
	*(*int32)(unsafe.Pointer(&data[340])) = persona.Job.JobType
	copy(data[344:360], persona.Job.From[:])
	copy(data[360:376], persona.Job.To[:])
	*(*int32)(unsafe.Pointer(&data[376])) = persona.Job.Status
	*(*int32)(unsafe.Pointer(&data[380])) = persona.Job.Value
	*(*int32)(unsafe.Pointer(&data[384])) = persona.Job.GTU
	*(*int32)(unsafe.Pointer(&data[388])) = persona.Job.Credits
	switch persona.Job.JobType {
	case jobs.JOB_FACTORY:
		*(*model.DBFactoryJob)(unsafe.Pointer(&data[392])) = persona.Job.Type.FactryWk
	case jobs.JOB_GENERAL:
		*(*model.DBGeneralJob)(unsafe.Pointer(&data[396])) = persona.Job.Type.GenWk
	}
	*(*int32)(unsafe.Pointer(&data[464])) = persona.Job.Age

	// Trading.
	*(*int32)(unsafe.Pointer(&data[468])) = persona.LastTrade

	// Spaceship.
	*(*uint32)(unsafe.Pointer(&data[472])) = persona.ShipLoc
	copy(data[476:636], persona.ShipDesc[:])
	copy(data[636:652], persona.Registry[:])
	*(*model.DBEquipment)(unsafe.Pointer(&data[652])) = persona.ShipKit
	*(*uint32)(unsafe.Pointer(&data[704])) = persona.Missiles
	*(*uint32)(unsafe.Pointer(&data[708])) = persona.Ammo
	*(*[model.MAX_GUNS]model.SGuns)(unsafe.Pointer(&data[712])) = persona.Guns
	*(*[model.MAX_LOAD]model.DBCargo)(unsafe.Pointer(&data[776])) = persona.Load

	// Misc game stuff.
	*(*int32)(unsafe.Pointer(&data[1196])) = persona.LastOn
	*(*[4]int32)(unsafe.Pointer(&data[1200])) = persona.Count
	*(*model.DBBuild)(unsafe.Pointer(&data[1216])) = persona.Build

	// Other structures.
	if persona.Rank < uint32(model.RankSquire) {
		*(*model.DBPPData)(unsafe.Pointer(&data[1228])) = *persona.PP
	} else if persona.Rank <= uint32(model.RankDuke) {
		*(*model.DBRPData)(unsafe.Pointer(&data[1228])) = *persona.RP
	}

	return page
}

// Open opens the persona database and performs cleanup/migration
func (db *Database) Open(loadPlayer func(*model.DBPersona, Offset) DatabasePlayer) error {
	realPath := "data/person.d" // FIXME
	newPath := realPath + "-NEW"
	oldPath := realPath + "-OLD"

	// Open the current persona file.
	inFile, err := os.Open(realPath)
	if err != nil {
		return fmt.Errorf("failed to open %s: %v", realPath, err)
	}
	defer inFile.Close()

	// Check the current file size.
	fileInfo, err := inFile.Stat()
	if err != nil {
		return fmt.Errorf("failed to stat data/person.d: %v", err)
	}
	fileSize := fileInfo.Size()
	if fileSize%int64(dbPageSize) != 0 {
		log.Panic("Size of the persona file is suspect")
	}

	// Create new persona file.
	outFile, err := os.Create(newPath)
	if err != nil {
		return fmt.Errorf("failed to create data/person.d-NEW: %v", err)
	}
	defer outFile.Close()

	// Write the dummy first record.
	zeroPage := make([]byte, dbPageSize)
	n, err := outFile.Write(zeroPage)
	if err != nil || n != dbPageSize {
		return fmt.Errorf("failed to write zero page: %v", err)
	}
	db.nextOffset += dbPageSize

	// Open the external player list.
	db.elist, _ = os.Create("log/persona.list")
	unix.CloseOnExec(int(db.elist.Fd()))

	var numPersonas, numLoaded int
	poLoadList := []int64{}    // Offsets for Planet Owner personas
	otherLoadList := []int64{} // Offsets for other personas

	// Pre-scan the persona file. Personas for cancelled accounts are
	// dropped here, Dukes are loaded, the remainder are tagged into 2
	// lists (planet-owners and others).
	for offset := int64(0); offset < fileSize; offset += dbPageSize {
		var page dbPage
		n, err := inFile.Read(page[:])
		if err != nil || n != dbPageSize {
			log.Panicf("Read failed: %v", err)
		}

		// Skip the dummy first record
		if offset == 0 {
			continue
		}

		dbPersona := unpackDBPersona(&page)

		// Skip anything without a UID or name (not that there should
		// be any).
		if dbPersona.ID == 0 || dbPersona.Name[0] == 0 {
			continue
		}

		// Nuke cancelled accounts, but ensure the persona has been
		// absent 7 days before doing so.
		// TODO

		// Load or tag the persona.
		numPersonas++
		switch {
		case model.Rank(dbPersona.Rank) == model.RankDuke:
			load(dbPersona, outFile, loadPlayer)
			numLoaded++
		case model.Rank(dbPersona.Rank) >= model.RankSquire && model.Rank(dbPersona.Rank) < model.RankDuke:
			poLoadList = append(poLoadList, offset)
		default:
			otherLoadList = append(otherLoadList, offset)
		}
	}

	log.Printf("There are %d personas in the file:", numPersonas)
	log.Printf("%6d Dukes", numLoaded)
	log.Printf("%6d planet-owners", len(poLoadList))
	log.Printf("%6d others", len(otherLoadList))

	// Fix up the duchies.
	log.Print("Fixing up duchies")
	// for (DuchyFixups::const_iterator iter = duchyFixups.begin();
	// iter != duchyFixups.end();
	// iter++)
	// {
	// 	PlayerDuchy* thisDuchy = iter->duchy;
	// 	if (strlen(iter->favoured) > 0) {
	// 		thisDuchy->setFavoured(iter->favoured, NULL);
	// 	}
	// 	if (strlen(iter->embargo) > 0) {
	// 		thisDuchy->setEmbargo(iter->embargo, NULL);
	// 	}
	// }
	// duchyFixups.clear();

	// Load non-Duke personas.
	for _, offset := range poLoadList {
		var page dbPage
		n, err := inFile.ReadAt(page[:], offset)
		if err != nil || n != dbPageSize {
			log.Panicf("ReadAt failed: %v", err)
		}
		dbPersona := unpackDBPersona(&page)
		load(dbPersona, outFile, loadPlayer)
		numLoaded++
	}
	for _, offset := range otherLoadList {
		var page dbPage
		n, err := inFile.ReadAt(page[:], offset)
		if err != nil || n != dbPageSize {
			log.Panicf("ReadAt failed: %v", err)
		}
		dbPersona := unpackDBPersona(&page)
		load(dbPersona, outFile, loadPlayer)
		numLoaded++
	}

	// Close the input file.
	if err := inFile.Close(); err != nil {
		log.Panicf("Close: %v", err)
	}

	// Make sure we processed the expected number of personas.
	if numLoaded != numPersonas {
		log.Panicf("Open: Loaded %d, expected %d?", numLoaded, numPersonas)
	}

	// Close the output file.
	if err := outFile.Close(); err != nil {
		log.Panicf("Close: %v", err)
	}

	// Rename person.d to person.d-OLD, and person.d-NEW to person.d.
	err = os.Rename("data/person.d", "data/person.d-OLD")
	if err != nil {
		// OLD file might already exist, remove it first
		err = os.Rename("data/person.d", "data/person.d-OLD")
		if err != nil {
			return fmt.Errorf("failed to rename person.d to person.d-OLD: %v", err)
		}
	}
	log.Printf("%s -> %s", realPath, oldPath)
	err = os.Rename("data/person.d-NEW", "data/person.d")
	if err != nil {
		return fmt.Errorf("failed to rename person.d-NEW to person.d: %v", err)
	}
	log.Printf("%s -> %s", newPath, realPath)

	// Open the (new) persona file for real.
	db.file, err = os.OpenFile(realPath, os.O_RDWR, 0o600)
	if err != nil {
		return fmt.Errorf("failed to open final database: %v", err)
	}
	finalFileInfo, err := db.file.Stat()
	if err != nil {
		return fmt.Errorf("failed to stat final database: %v", err)
	}
	if finalFileInfo.Size() != int64(db.nextOffset) {
		return fmt.Errorf("database size mismatch: nextOffset=%d, actual file size=%d",
			db.nextOffset, finalFileInfo.Size())
	}

	unix.CloseOnExec(int(db.file.Fd()))

	// All done.
	return nil
}

func (db *Database) Close() {
	db.queueMu.Lock()
	defer db.queueMu.Unlock()

	if db.file != nil {
		// Save any dirty records that we've been holding on to.
		for _, player := range db.slowSaveQueue {
			// debug.Trace("Saving %d from slow queue", player.UID())
			db.write(player)
		}
		// debug.Trace("Delayed database updates flushed")

		db.idleTimer = time.AfterFunc(time.Duration(rand.IntN(int(time.Second))), db.idleHandler)

		// Close the persona file.
		if err := db.file.Close(); err != nil {
			log.Printf("Database.Close: Close() failed: %v", err)
		}
		log.Print("Persona file closed")
		db.file = nil
	}

	db.idTable = nil
	db.fastSaveQueue = nil
	db.slowSaveQueue = nil

	// Update metrics to reflect empty queues
	db.updateQueueMetrics()

	log.Print("Database closed")
}

func (db *Database) Commit() {
	db.queueMu.Lock()
	defer db.queueMu.Unlock()

	// debug.Trace("Database.Commit()")
	monitoring.DatabaseCommitTotal.Inc()

	for _, player := range db.fastSaveQueue {
		// debug.Trace("Saving %d from fast queue", player.UID())
		db.write(player)

		for j := range db.slowSaveQueue {
			if db.slowSaveQueue[j] == player {
				db.slowSaveQueue = append(db.slowSaveQueue[:j], db.slowSaveQueue[j+1:]...)
				break
			}
		}

		if player.IsDeceased() {
			debug.Trace("Deleting deceased player: %s", player.Name())
			// player.Delete() -- TODO
		}
	}
	db.fastSaveQueue = db.fastSaveQueue[:0]

	// Update metrics after queue changes
	db.updateQueueMetrics()

	// debug.Trace("Immediate database updates flushed")
}

// load processes a player and writes their page to the new database file
func load(persona *model.DBPersona, outFile *os.File, loadPlayer func(*model.DBPersona, Offset) DatabasePlayer) {
	if persona == nil {
		log.Panic("persona is nil")
	}
	if outFile == nil {
		log.Panic("outFile is nil")
	}

	// Clean up any dross in the persona record.
	cleanPersonaRecord(persona)

	// Roll out planets with absentee owners
	if persona.Rank >= uint32(model.RankSquire) && persona.Rank <= uint32(model.RankDuke) {
		lastOn := time.Unix(int64(persona.LastOn), 0)
		daysOff := int(time.Since(lastOn).Hours() / 24) // Transaction::dayNumber() - lastOn

		if daysOff >= 10 && (persona.RP.Planet.Flags&model.PLT_CLOSED) == 0 {
			log.Printf("Closing %s: %s last on %d days ago",
				text.CStringToString(persona.RP.Fief[:]),
				text.CStringToString(persona.Name[:]),
				daysOff)
			persona.RP.Planet.Flags |= model.PLT_CLOSED
		}
	}

	// Load into in-memory database.
	player := loadPlayer(persona, database.nextOffset)
	if player == nil {
		log.Panic("player is nil")
	}

	// Pack the dbPersona back into page format and write to output file
	page := packDBPersona(persona)
	n, err := outFile.Write(page[:])
	if err != nil || n != dbPageSize {
		log.Panicf("load: outFile.Write() failed: %v", err)
	}
	database.nextOffset += dbPageSize

	_, _ = fmt.Fprintf(database.elist, "%s,%d\n", player.Name(), player.UID())
}

func cleanPersonaRecord(persona *model.DBPersona) {
	// // Clear bogus last_on fields -- FIXME
	// if persona.LastOn == 999999999 {
	// 	persona.LastOn = 0
	// }
	// if persona.LastOn == 0 && persona.Rank == uint32(model.RankDuke) {
	// 	persona.LastOn = int32(time.Now().Unix())
	// 	persona.RP.Planet.Flags = 0
	// }
	if persona.Rank == uint32(model.RankDuke) { // FIXME: testing
		persona.LastOn = int32(time.Now().Unix())
		persona.RP.Planet.Flags = 0
	}

	// Clear out any redundant flags.
	persona.Flags[0] &= model.PL0_FLYING +
		model.PL0_BRIEF +
		model.PL0_COMM_UNIT +
		model.PL0_LIT +
		model.PL0_INSURED +
		model.PL0_JOB +
		model.PL0_INFO +
		model.PL0_SULKING +
		model.PL0_SPYBEAM +
		model.PL0_OFFER_TOUR +
		model.PL0_SPYSCREEN +
		model.PL0_KNOWS_ORSONITE +
		model.PL0_SNARK_ASSIGNED +
		model.PL0_HORSELL_ASSIGNED
	persona.Flags[1] &= model.PL1_SHIP_PERMIT +
		model.PL1_SHIELDS +
		model.PL1_AUTO +
		model.PL1_MI6_OFFERED +
		model.PL1_MI6 +
		model.PL1_TITAN +
		model.PL1_CALLISTO +
		model.PL1_MARS +
		model.PL1_EARTH +
		model.PL1_MOON +
		model.PL1_VENUS +
		model.PL1_MERCURY +
		model.PL1_HILBERT +
		model.PL1_DONE_STA +
		model.PL1_DONE_STR +
		model.PL1_DONE_DEX +
		model.PL1_DONE_INT +
		model.PL1_PO_PERMIT +
		model.PL1_DONE_SNARK +
		model.PL1_NAVIGATOR +
		model.PL1_PROMO_CHAR
	if persona.Rank != uint32(model.RankGroundHog) {
		persona.Flags[1] |= model.PL1_SHIP_PERMIT
	}
	if persona.Rank != uint32(model.RankExplorer) {
		if persona.Rank >= uint32(model.RankSquire) && persona.Rank <= uint32(model.RankDuke) {
			persona.Flags[1] |= model.PL1_PO_PERMIT
		} else {
			persona.Flags[1] &^= model.PL1_PO_PERMIT
		}
	}
	if (persona.Flags[1] & model.PL1_MI6) != 0 {
		persona.Flags[1] &^= model.PL1_MI6_OFFERED
	}

	// Clear spurious loans.
	if persona.Loan < 0 || persona.Rank != uint32(model.RankCommander) {
		persona.Loan = 0
	}

	// Clear negative rewards.
	if persona.Reward < 0 {
		persona.Reward = 0
	}

	// Clear special-purpose indicator for all player ranks.
	if persona.Rank < uint32(model.RankSenator) {
		persona.Frame[1] = 0
	}

	// Clear descriptions that start with a reserved output prefix.
	if persona.Desc[0] == '/' || persona.Desc[0] == '>' {
		clear(persona.Desc[:])
	}
	if persona.Mood[0] == '/' || persona.Mood[0] == '>' {
		clear(persona.Mood[:])
	}
	if persona.ShipDesc[0] == '/' || persona.ShipDesc[0] == '>' {
		clear(persona.ShipDesc[:])
	}

	// Clear the active build from players who can't possibly be building.
	if persona.Rank < uint32(model.RankExplorer) || persona.Rank > uint32(model.RankDuke) {
		persona.Build = model.DBBuild{}
	}

	if persona.Rank < uint32(model.RankSquire) {
		// Fix warehouse bays with cost of less than 10 IG per ton.
		for i := range persona.PP.Storage {
			warehouse := &persona.PP.Storage[i]
			for j := range warehouse.Bay {
				if warehouse.Bay[j].Quantity > 0 {
					if warehouse.Bay[j].Cost < 10 {
						warehouse.Bay[j].Cost = 10
					}
				}
			}
		}
		for i := range persona.PP.Factory {
			factory := &persona.PP.Factory[i]
			if factory.Wages < 0 {
				factory.Wages = 0
			} else if factory.Wages > model.MAX_WAGE {
				factory.Wages = model.MAX_WAGE
			}
			if factory.Layoff < 0 {
				factory.Layoff = 0
			} else if factory.Layoff > factory.Wages {
				factory.Layoff = factory.Wages
			}
		}
	} else if persona.Rank <= uint32(model.RankDuke) {
		planet := &persona.RP.Planet
		if planet.Population < 2 {
			planet.Population = 2
		}
		if planet.Tax < 0 {
			planet.Tax = 0
		} else if planet.Tax > 30 {
			planet.Tax = 30
		}

		// Fix warehouse bays with cost of less than 10 IG per ton.
		warehouse := &persona.RP.Storage
		for i := range warehouse.Bay {
			if warehouse.Bay[i].Quantity > 0 {
				if warehouse.Bay[i].Cost < 10 {
					warehouse.Bay[i].Cost = 10
				}
			}
		}
	}

	// Clean up promo characters.
	if (persona.Flags[1] & model.PL1_PROMO_CHAR) != 0 {
		persona.Flags[0] |= model.PL0_COMM_UNIT
		persona.Flags[0] |= model.PL0_LIT
		persona.Flags[0] |= model.PL0_INSURED
		persona.Flags[0] |= model.PL0_SPYSCREEN

		persona.Flags[0] &^= model.PL0_SPYBEAM
		persona.Flags[0] &^= model.PL0_OFFER_TOUR
		persona.Flags[0] &^= model.PL0_KNOWS_ORSONITE
		persona.Flags[0] &^= model.PL0_SNARK_ASSIGNED
		persona.Flags[0] &^= model.PL0_HORSELL_ASSIGNED

		persona.Flags[1] |= model.PL1_DONE_STA
		persona.Flags[1] |= model.PL1_DONE_STR
		persona.Flags[1] |= model.PL1_DONE_DEX
		persona.Flags[1] |= model.PL1_DONE_INT
		persona.Flags[1] |= model.PL1_DONE_SNARK

		persona.Flags[1] &^= model.PL1_MI6_OFFERED
		persona.Flags[1] &^= model.PL1_MI6
		persona.Flags[1] &^= model.PL1_HILBERT
		persona.Flags[1] &^= model.PL1_NAVIGATOR

		if persona.Rank >= uint32(model.RankSquire) {
			persona.Build.IDProject = build.NOTHING

			persona.RP.Storage = model.DBWarehouse{}

			for i := range persona.RP.Facilities {
				persona.RP.Facilities[i] = 0
			}

			if persona.Rank == uint32(model.RankBaron) {
				persona.RP.Facilities[model.DK_UPSIDE] = 100
				persona.RP.Facilities[model.DK_DOWNSIDE] = 100
				persona.RP.Facilities[model.DK_MATTRANS] = 100
			}

			persona.RP.Planet.Duchy = [model.NAME_SIZE]byte{}
			copy(persona.RP.Planet.Duchy[:], "Sol")
		}
	}

	// Fix oversized ships (again).
	if persona.ShipLoc > 0 && persona.ShipKit.MaxHull > 1000 {
		log.Printf("BUG EXPLOIT: Fixing %s", persona.Name)
		if persona.Balance > 0 {
			persona.Balance /= 2
			log.Print("BUG EXPLOIT: Fine levied against personal balance")
		}
		if persona.Rank < uint32(model.RankSquire) {
			if persona.PP.Company.Balance > 0 {
				persona.PP.Company.Balance /= 2
				log.Print("BUG EXPLOIT: Fine levied against company balance")
			}
		} else if persona.Rank < uint32(model.RankSenator) {
			if persona.RP.Planet.Balance > 0 {
				persona.RP.Planet.Balance /= 2
				log.Print("BUG EXPLOIT: Fine levied against planet treasury")
			}
		}
		persona.Flags[0] &^= model.PL0_FLYING
		persona.Flags[0] &^= model.PL0_SPYBEAM
		persona.Flags[1] &^= model.PL1_SHIELDS
		persona.Flags[1] &^= model.PL1_TITAN
		persona.Flags[1] &^= model.PL1_CALLISTO
		persona.Flags[1] &^= model.PL1_MARS
		persona.Flags[1] &^= model.PL1_EARTH
		persona.Flags[1] &^= model.PL1_MOON
		persona.Flags[1] &^= model.PL1_VENUS
		persona.Flags[1] &^= model.PL1_MERCURY
		persona.Flags[1] &^= model.PL1_HILBERT
		persona.LocNo = 426
		persona.StarSystem = [model.NAME_SIZE]byte{}
		copy(persona.StarSystem[:], "Sol")
		persona.Job = model.DBWork{}
		persona.ShipLoc = 0
		clear(persona.ShipDesc[:])
		clear(persona.Registry[:])
		persona.ShipKit = model.DBEquipment{}
		persona.Missiles = 0
		persona.Ammo = 0
		for i := range persona.Guns {
			persona.Guns[i] = model.SGuns{}
		}
		for i := range persona.Load {
			persona.Load[i] = model.DBCargo{}
		}
	}
}

func (db *Database) Modify(player DatabasePlayer, when SaveWhen) {
	if player == nil {
		log.Panicf("Database.Modify: player is nil")
	}
	if when != SaveNow && when != SaveLater {
		log.Panicf("Database.Modify: when is not SaveNow or SaveLater")
	}
	// debug.Trace("Database.Modify(%s,%d)", player.Name(), when)

	monitoring.DatabaseModifyTotal.Inc()

	database.queueMu.Lock()
	defer database.queueMu.Unlock()
	defer db.updateQueueMetrics() // Update metrics when queues change

	if when == SaveNow {
		if slices.Contains(db.fastSaveQueue, player) {
			return
		}
		// debug.Trace("Adding %d to fast save queue", player.UID())
		db.fastSaveQueue = append(db.fastSaveQueue, player)
	} else {
		if slices.Contains(db.slowSaveQueue, player) {
			return
		}
		// debug.Trace("Adding %d to slow save queue", player.UID())
		db.slowSaveQueue = append(db.slowSaveQueue, player)
		if db.idleTimer != nil {
			return
		}
		// debug.Trace("Database.Modify: Requesting idle call")
		db.idleTimer = time.AfterFunc(time.Duration(rand.IntN(int(time.Second))), db.idleHandler)
	}
}

func (db *Database) idleHandler() {
	global.Lock()
	defer global.Unlock()

	db.IdleProc()
}

func (db *Database) IdleProc() {
	db.queueMu.Lock()
	defer db.queueMu.Unlock()

	db.idleTimer = nil

	if len(db.slowSaveQueue) == 0 {
		return
	}

	player := db.slowSaveQueue[0]
	debug.Trace("Saving %d from slow queue (idle)", player.UID())
	db.write(player)

	db.slowSaveQueue = db.slowSaveQueue[1:]

	// Update metrics after queue changes
	db.updateQueueMetrics()

	if len(db.slowSaveQueue) == 0 {
		debug.Trace("Delayed database updates flushed")
		return
	}

	db.idleTimer = time.AfterFunc(time.Duration(rand.IntN(int(time.Second))), db.idleHandler)
}

func (db *Database) write(player DatabasePlayer) {
	if player == nil {
		log.Panicf("Database.write: player is nil")
	}

	// Paranoia check on the offset!
	if player.Offset() <= 0 || (player.Offset()%dbPageSize) != 0 {
		log.Panicf("Database.write: Bad offset (%d)", player.Offset())
	}

	// If the player's dead then write an empty page to destroy the
	// existing record.
	if player.IsDeceased() {
		log.Printf("Deleting the late %s [%d]", player.Name(), player.UID())
		var page dbPage
		n, err := db.file.WriteAt(page[:], int64(player.Offset()))
		if err != nil || n != dbPageSize {
			log.Panicf("Database.write: fatal I/O error writing player %d to database: %v", player.UID(), err)
		}
		monitoring.DatabaseWriteTotal.WithLabelValues("delete").Inc()
	} else {
		persona := player.Serialize()
		page := packDBPersona(persona)
		n, err := db.file.WriteAt(page[:], int64(player.Offset()))
		if err != nil || n != dbPageSize {
			log.Panicf("Database.write: fatal I/O error writing player %d to database: %v", player.UID(), err)
		}
		monitoring.DatabaseWriteTotal.WithLabelValues("write").Inc()
	}
}

func (db *Database) nextPersonaOffset() Offset {
	offset := db.nextOffset
	if offset == 0 {
		log.Panic("offset is zero")
	}
	if (offset % dbPageSize) != 0 {
		log.Panic("offset is not a multiple of page size")
	}

	db.nextOffset += dbPageSize
	if db.nextOffset == 0 {
		log.Panic("db.nextOffset is zero")
	}
	if (db.nextOffset % dbPageSize) != 0 {
		log.Panic("db.nextOffset is not a multiple of page size")
	}

	return offset
}

// UnpackDBPersona unpacks raw page data into a DBPersona with union handling
//
//nolint:gosec // Audited: entire function uses unsafe for 1999 C binary format compatibility
func unpackDBPersona(page *dbPage) *model.DBPersona {
	var persona model.DBPersona
	data := page[:dbPersonaSize]

	// Personal info.
	copy(persona.Name[:], data[0:16])
	persona.ID = *(*uint32)(unsafe.Pointer(&data[16]))
	persona.Sex = data[20]
	persona.Rank = *(*uint32)(unsafe.Pointer(&data[24]))
	copy(persona.Desc[:], data[28:180])
	copy(persona.Mood[:], data[180:216])

	// Statistics.
	persona.MaxStr = *(*uint32)(unsafe.Pointer(&data[216]))
	persona.CurStr = *(*uint32)(unsafe.Pointer(&data[220]))
	persona.MaxSta = *(*uint32)(unsafe.Pointer(&data[224]))
	persona.CurSta = *(*uint32)(unsafe.Pointer(&data[228]))
	persona.MaxInt = *(*uint32)(unsafe.Pointer(&data[232]))
	persona.CurInt = *(*uint32)(unsafe.Pointer(&data[236]))
	persona.MaxDex = *(*uint32)(unsafe.Pointer(&data[240]))
	persona.CurDex = *(*uint32)(unsafe.Pointer(&data[244]))
	persona.Shipped = *(*uint32)(unsafe.Pointer(&data[248]))
	persona.Games = *(*uint32)(unsafe.Pointer(&data[252]))
	persona.Flags[0] = *(*uint32)(unsafe.Pointer(&data[256]))
	persona.Flags[1] = *(*uint32)(unsafe.Pointer(&data[260]))

	// Money, etc.
	persona.Balance = *(*int32)(unsafe.Pointer(&data[264]))
	persona.Loan = *(*int32)(unsafe.Pointer(&data[268]))
	persona.Reward = *(*int32)(unsafe.Pointer(&data[272]))
	persona.Frame = *(*[4]int32)(unsafe.Pointer(&data[276]))

	// Locations.
	persona.LocNo = *(*uint32)(unsafe.Pointer(&data[292]))
	copy(persona.StarSystem[:], data[296:312])

	// Work.
	persona.Job.Pallet = *(*model.DBCargo)(unsafe.Pointer(&data[312]))
	persona.Job.JobType = *(*int32)(unsafe.Pointer(&data[340]))
	copy(persona.Job.From[:], data[344:360])
	copy(persona.Job.To[:], data[360:376])
	persona.Job.Status = *(*int32)(unsafe.Pointer(&data[376]))
	persona.Job.Value = *(*int32)(unsafe.Pointer(&data[380]))
	persona.Job.GTU = *(*int32)(unsafe.Pointer(&data[384]))
	persona.Job.Credits = *(*int32)(unsafe.Pointer(&data[388]))
	switch persona.Job.JobType {
	case jobs.JOB_FACTORY:
		persona.Job.Type.FactryWk = *(*model.DBFactoryJob)(unsafe.Pointer(&data[392]))
	case jobs.JOB_GENERAL:
		persona.Job.Type.GenWk = *(*model.DBGeneralJob)(unsafe.Pointer(&data[396]))
	}
	persona.Job.Age = *(*int32)(unsafe.Pointer(&data[464]))

	// Trading.
	persona.LastTrade = *(*int32)(unsafe.Pointer(&data[468]))

	// Spaceship.
	persona.ShipLoc = *(*uint32)(unsafe.Pointer(&data[472]))
	copy(persona.ShipDesc[:], data[476:636])
	copy(persona.Registry[:], data[636:652])
	persona.ShipKit = *(*model.DBEquipment)(unsafe.Pointer(&data[652]))
	persona.Missiles = *(*uint32)(unsafe.Pointer(&data[704]))
	persona.Ammo = *(*uint32)(unsafe.Pointer(&data[708]))
	persona.Guns = *(*[model.MAX_GUNS]model.SGuns)(unsafe.Pointer(&data[712]))
	persona.Load = *(*[model.MAX_LOAD]model.DBCargo)(unsafe.Pointer(&data[776]))

	// Misc game stuff.
	persona.LastOn = *(*int32)(unsafe.Pointer(&data[1196]))
	persona.Count = *(*[4]int32)(unsafe.Pointer(&data[1200]))
	persona.Build = *(*model.DBBuild)(unsafe.Pointer(&data[1216]))

	// Other structures.
	if persona.Rank < uint32(model.RankSquire) {
		persona.PP = (*model.DBPPData)(unsafe.Pointer(&data[1228]))
	} else if persona.Rank <= uint32(model.RankDuke) {
		persona.RP = (*model.DBRPData)(unsafe.Pointer(&data[1228]))
	}

	return &persona
}

func (db *Database) AddPlayerID(player DatabasePlayer) {
	// TODO
}

func (db *Database) DeletePlayerID(uid uint32) {
	// TODO
}

func (db *Database) FindPlayerID(uid uint32) {
	// TODO
}
