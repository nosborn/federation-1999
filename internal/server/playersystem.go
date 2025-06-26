package server

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/global"
	"github.com/nosborn/federation-1999/internal/text"
	"github.com/nosborn/federation-1999/internal/workbench"
)

const (
	unloadTimerPeriod = 60 * time.Second
)

type PlayerSystem struct {
	System

	unloadTicks int
	unloadTimer *time.Timer
}

// func NewPlayerSystem(owner *Player, systemName string, noProduction bool) {
// 	s := &System{
// 		duchy:     SolDuchy,
// 		owner:     owner,
// 		loadState: SystemLoading,
// 		name:      systemName,
// 		taxRate:   10,
// 	}
// 	s.duchy.AddMember(s)
// 	// : System(systemName, 0, solDuchy, 0, owner, 10, 0, 0),
// 	//   m_destroy(false),
// 	//   m_planet(0),
// 	//   m_unloadTicks(0),
// 	//   m_unloadTimer(0)
// 	s.Planets[0] = NewPlanet(s, noProduction)
// 	s.beginLoad()
// }

func NewPlayerSystemFromDB(owner *Player, systemName string, dbPlanet model.DBPlanet) *PlayerSystem {
	s := &PlayerSystem{
		System: System{
			flags:       dbPlanet.Flags &^ model.PLT_CLOSED,
			owner:       owner,
			balance:     dbPlanet.Balance,
			lastOnline:  dbPlanet.LastOnline,
			loadState:   SystemLoading,
			name:        systemName,
			taxRate:     dbPlanet.Tax,
			touristTime: dbPlanet.Time,
		},
	}

	var ok bool
	s.duchy, ok = FindDuchy(text.CStringToString(dbPlanet.Duchy[:]))
	if !ok {
		log.Printf("Moving %s system to Sol duchy", s.name)
		s.duchy = SolDuchy
	}
	s.duchy.AddMember(s)

	if err := allSystems.Insert(s); err != nil {
		log.Panic("PANIC: Duplicate system added: ", err)
	}
	systemIndex.Insert(s.name, s)

	planet := NewPlayerPlanet(s, dbPlanet)
	s.planets = append(s.planets, planet)

	if (s.flags&model.PLT_CLOSED) == 0 && !owner.IsLockedOut() {
		Enqueue(s)
	}

	return s
}

func (s *PlayerSystem) BeginLoad() {
	s.loadState = SystemLoading
	Enqueue(s)
}

func (s *PlayerSystem) BeginUnload() {
	if s.loadState != SystemOnline {
		log.Panic("system.BeginUnload: loadState != SystemOnline")
	}

	s.loadState = SystemUnloading

	// Clean up factories.
	// FIXME: Should be a destroy somewhere if the PO is deceased.

	if len(s.planets) > 0 {
		s.planets[0].StopExchange()
		s.planets[0].StopFactories()
	}

	s.unloadTimer = time.AfterFunc(unloadTimerPeriod, s.unloadTimerHandler)
}

func (s *PlayerSystem) Load() bool {
	if s.owner == nil {
		return false
	}

	// Clear out the key location numbers. loadLocations() will reset them.
	s.planets[0].exchangeLocNo = 0
	s.planets[0].hospitalLocNo = 0
	s.planets[0].landingLocNo = 0
	s.planets[0].orbitLocNo = 0

	if !s.loadEvents() {
		log.Printf("%s: Can't load event data", s.name)
		return false
	}
	if !s.loadLocations() {
		log.Printf("%s: Can't load location data", s.name)
		s.destroyEvents()
		return false
	}
	if !s.loadObjects() {
		log.Printf("%s: Can't load object data", s.name)
		s.destroyLocations()
		s.destroyEvents()
		return false
	}
	log.Printf("%s [%d] loaded E=%d H=%d L=%d O=%d", s.name, s.owner.UID(), s.planets[0].exchangeLocNo,
		s.planets[0].hospitalLocNo, s.planets[0].landingLocNo, s.planets[0].orbitLocNo)

	// Fix up anything that might have missed a Galactic Midnight.
	today := int32(time.Now().Unix() / 86400) // Transaction::dayNumber())
	if s.lastOnline != today {
		s.flags &^= model.PLT_T4_HOLDER
		s.touristTime = 0
	}
	s.lastOnline = today

	s.loadState = SystemOnline
	s.owner.Save(database.SaveNow)

	s.planets[0].StartExchange()
	s.owner.resumeConstruction()

	return true
}

func (s *PlayerSystem) loadEvents() bool {
	pathname := fmt.Sprintf("data/workbench%d/%d.e", s.owner.UID()%10, s.owner.UID())
	f, err := os.Open(pathname)
	if err != nil {
		log.Printf("loadEvents: %#v", err)
		return false
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		log.Printf("loadEvents: %#v", err)
		return false
	}
	fileSize := fileInfo.Size()
	const eventSize = 176
	if fileSize%eventSize != 0 {
		log.Printf("loadEvents: %s has dubious size %d (not multiple of %d)", pathname, fileSize, eventSize)
		return false
	}

	if fileSize == 0 {
		return true
	}

	numEvents := fileSize / eventSize

	for i := range numEvents {
		var dbEvent model.DBEvent
		err := binary.Read(f, binary.LittleEndian, &dbEvent)
		if err != nil {
			log.Printf("loadEvents: failed to read event %d: %v", i, err)
			return false
		}

		event := NewEventFromDB(dbEvent)
		s.events = append(s.events, event)
	}
	return true
}

func (s *PlayerSystem) loadLocations() bool {
	pathname := fmt.Sprintf("data/workbench%d/%d.l", s.owner.UID()%10, s.owner.UID())
	f, err := os.Open(pathname)
	if err != nil {
		log.Printf("loadLocations: %#v", err)
		return false
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		log.Printf("loadLocations: %#v", err)
		return false
	}
	fileSize := fileInfo.Size()
	const locationSize = 1060
	if fileSize%locationSize != 0 {
		log.Printf("loadLocations: %s has dubious size %d", pathname, fileSize)
		return false
	}

	numLocations := fileSize / locationSize

	if numLocations <= SPACESHIP_SIZE {
		if numLocations == 0 {
			log.Printf("loadLocations: %s is empty", pathname)
		} else {
			log.Printf("loadLocations: %s has dubious size", pathname)
		}
		return false
	}

	if numLocations > SPACESHIP_SIZE+workbench.LOCATION_LIMIT {
		log.Printf("Dropping excess locations from %s", s.name)
		numLocations = SPACESHIP_SIZE + workbench.LOCATION_LIMIT
	}

	for i := range numLocations {
		var dbLocation model.DBLocation
		err := binary.Read(f, binary.LittleEndian, &dbLocation)
		if err != nil {
			log.Printf("loadLocations: failed to read location %d: %v", i, err)
			return false
		}

		if i < SPACESHIP_SIZE {
			continue
		}

		location := NewLocationFromDB(uint32(i+1), dbLocation, s)
		s.locations = append(s.locations, location)

		if location.IsLink() {
			s.linkLocNo = int32(i + 1)
		}

		if location.IsHospital() {
			s.planets[0].hospitalLocNo = uint32(i + 1)
		}
		if location.IsLandingPad() {
			s.planets[0].landingLocNo = uint32(i + 1)
		}
		if location.IsOrbit() {
			s.planets[0].orbitLocNo = uint32(i + 1)
		}

		if location.IsExchange() {
			switch s.planets[0].level {
			case model.LevelNoProduction, model.LevelCapital:
				// location.Flags &^= model.LfTrade -- FIXME
			default:
				s.planets[0].exchangeLocNo = uint32(i + 1)
			}
		}
	}
	return true
}

func (s *System) loadObjects() bool {
	pathname := fmt.Sprintf("data/workbench%d/%d.o", s.owner.UID()%10, s.owner.UID())
	f, err := os.Open(pathname)
	if err != nil {
		log.Printf("loadObjects: %#v", err)
		return false
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		log.Printf("loadObjects: %#v", err)
		return false
	}
	fileSize := fileInfo.Size()
	const objectSize = 464
	if fileSize%objectSize != 0 {
		log.Printf("loadObjects: %s has dubious size %d (not multiple of %d)", pathname, fileSize, objectSize)
		return false
	}

	if fileSize == 0 {
		return true
	}

	numObjects := fileSize / objectSize

	for i := range numObjects {
		var dbObject model.DBObject
		err := binary.Read(f, binary.LittleEndian, &dbObject)
		if err != nil {
			log.Printf("loadObjects: failed to read object %d: %v", i, err)
			return false
		}

		object := NewObjectFromDB(dbObject, s)
		s.objects = append(s.objects, object)
	}
	return true
}

func (s *PlayerSystem) unloadTimerHandler() {
	global.Lock()
	defer global.Unlock()

	defer database.CommitDatabase()

	s.unloadTimerProc()

	if s.IsOffline() && s.owner == nil { //nolint:staticcheck // SA9003: empty branch
		// delete thisSystem -- TODO
	}
}

func (s *PlayerSystem) unloadTimerProc() {
	log.Printf("Unload timer proc for %s", s.name)

	s.unloadTicks++
	s.unloadTimer = nil

	vacant := true
	for _, player := range Players {
		if player.curSys == s {
			if s.unloadTicks >= 10 && player != s.owner {
				player.Deport()
				continue
			}
			if player.HasCommUnit() {
				player.Outputm(text.MN1543, s.name)
				player.FlushOutput()
			}
			vacant = false
		}
	}
	if !vacant {
		s.unloadTimer = time.AfterFunc(unloadTimerPeriod, s.unloadTimerHandler)
		return
	}

	s.loadState = SystemOffline

	s.destroyObjects()
	s.destroyLocations()
	s.destroyEvents()

	if s.owner != nil {
		s.owner.Outputm(text.OfflineNotify, s.name)
		s.owner.FlushOutput()
	}
}
