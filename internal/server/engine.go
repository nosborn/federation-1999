package server

import (
	"log"
	"strings"
	"time"

	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/global"
	"github.com/nosborn/federation-1999/internal/server/parser"
	"github.com/nosborn/federation-1999/internal/text"
)

var (
	CurrentPlayers int
	PeakPlayers    int
)

var (
	SolDuchy    *Duchy
	SolSystem   *solSystem
	arenaSystem *ArenaSystem
	snarkSystem *SnarkSystem
)

var Players map[string]*Player

func InitializeGame(startListener func()) bool {
	global.Lock()
	defer global.Unlock()

	log.Print(text.Msg(text.Dashs))

	Players = make(map[string]*Player) // out of place?

	//
	defer database.CommitDatabase()

	// Initialize the core star systems.
	SolDuchy = NewSolDuchy()
	SolSystem = NewSolSystem(SolDuchy)
	arenaSystem = NewArenaSystem(SolDuchy)
	_ = arenaSystem
	snarkSystem = NewSnarkSystem(SolDuchy)

	NewLoader(startListener)

	if err := database.OpenDatabase(func(dbp *model.DBPersona, off database.Offset) database.DatabasePlayer {
		return LoadPlayer(dbp, off)
	}); err != nil {
		log.Panicf("Failed to open persona database: %v", err)
	}

	//
	// initializeJobs()

	global.GameStartTime = time.Now().Local() // Use the real time, not the transaction time.

	//
	// const int seconds = gameStartTime % SECS_IN_A_DAY;
	// starday = 200501 + (gameStartTime / SECS_IN_A_DAY);
	// startime = (seconds * 1000) / SECS_IN_A_DAY;
	// const int remaining = ((startime + 1) * SECS_IN_A_DAY) - (seconds * 1000);
	// dbgTrace("Initial startime: %d %d %d", starday, startime, remaining);
	// Tcl_CreateTimerHandler(remaining, startimeTimerProc, NULL);

	//
	log.Print("Initialisation complete")
	log.Print(text.Msg(text.Dashs))

	return true

	// // solDuchy := NewCoreDuchy("Sol", 25)
	// // if err := loadSolSystem(SolDuchy); err != nil {
	// // 	log.Panicf("%v", err)
	// // }
	// // if err := loadArenaSystem(SolDuchy); err != nil {
	// // 	log.Panicf("%v", err)
	// // }
	// // if err := loadSnarkSystem(SolDuchy); err != nil {
	// // 	log.Panicf("%v", err)
	// // }
	// SolDuchy.Start()
	//
	// SolSystem.Start()
	// ArenaSystem.Start()
	// SnarkSystem.Start()
	//
	// NewLoader(startListener)
	//
	// if err := OpenDatabase(LoadPlayer); err != nil {
	// 	log.Panicf("Failed to open persona database: %v", err)
	// }
	//
	// StartLoader()
	//
	// for duchy := range allDuchies.Values() {
	// 	duchy.StartTransportation()
	// }
	//
	// // d := NewCoreDuchy("Sol", 10)
	// // Duchies["Sol"] = d
	// //
	// // s, err := LoadSystem("Sol", d.Name)
	// // if err != nil {
	// // 	log.Panicf("Sol: %v", err)
	// // }
	// // Systems["Sol"] = s
	// //
	// // s, err = LoadSystem("Arena", d.Name)
	// // if err != nil {
	// // 	log.Panicf("Arena: %v", err)
	// // }
	// // Systems["Arena"] = s
	//
	// // log.Printf("Sol.Events[0]: %#v\n", &Systems["Sol"].Events[0])
	// // log.Printf("Sol.Locations[0]: %#v\n", &Systems["Sol"].Locations[0])
	// // log.Printf("Sol.Objects[0]: %#v\n", &Systems["Sol"].Objects[0])
}

func changeBalance(balance *int32, amount int32) {
	result := int64(*balance) + int64(amount)
	switch {
	case result > int64(model.MAX_BALANCE):
		*balance = model.MAX_BALANCE
	case result < int64(model.MIN_BALANCE):
		*balance = model.MIN_BALANCE
	default:
		*balance = int32(result)
	}
}

// func loadArenaSystem(duchy *Duchy) error {
// 	data := CoreSystem{
// 		Locations: ArenaLocations,
// 		Objects:   ArenaObjects,
// 		Planets:   ArenaPlanets,
// 	}
// 	NewCoreSystem("Arena", 30, data, duchy)
// 	// s.Start()
// 	return nil
// }

// func loadSnarkSystem(duchy *Duchy) error {
// 	data := CoreSystem{
// 		Locations: SnarkLocations,
// 		Objects:   SnarkObjects,
// 		Planets:   SnarkPlanets,
// 	}
// 	NewCoreSystem("Snark", 10, data, duchy)
// 	// s.Start()
// 	return nil
// }

// func loadSolSystem(duchy *Duchy) error {
// 	data := CoreSystem{
// 		Events:    SolEvents,
// 		Locations: SolLocations,
// 		Objects:   SolObjects,
// 		Planets:   SolPlanets,
// 	}
// 	NewCoreSystem("Sol", 10, data, duchy)
// 	// s.Start()
// 	return nil
// }

// Checks name chosen by player against persona records and dictionary. For
// sanity's sake, we'll only accept names containing 7-bit ASCII.
func IsNameAvailable(name string) bool {
	debug.Precondition(name != "")

	// Names must be at least 3 characters long.
	if len(name) < 3 {
		log.Printf("Rejecting '%s' name: too short", name)
		return false
	}

	// Name must start with a letter. Names consisting of a single letter
	// repeated are deemed unacceptable.
	if !text.IsAlpha(name[0]) {
		log.Printf("Rejecting '%s' name: non-alpha first character", name)
		return false
	}

	stupidName := true
	for i, ch := range []byte(name) {
		if !text.IsPrint(ch) {
			log.Printf("Rejecting '%s' name: not visible ASCII", name)
			return false
		}
		if i > 0 && ch != name[0] {
			stupidName = false
			break
		}
	}
	if stupidName {
		log.Printf("Rejecting '%s' name: stupid name", name)
		return false
	}

	// XX...XX names indicate troublemakers on AOL, so we'll continue the
	// tradition of blocking them here.
	if len(name) >= 5 {
		if strings.HasPrefix(strings.ToLower(name), "xx") && strings.HasSuffix(strings.ToLower(name), "xx") {
			log.Printf("Rejecting '%s' name: troublemaker", name)
			return false
		}
	}

	// Check for a name in the database.
	_, ok := FindCompany(name)
	if ok {
		log.Printf("Rejecting '%s' name: matches company name", name)
		return false
	}
	_, ok = FindDuchy(name)
	if ok {
		log.Printf("Rejecting '%s' name: matches duchy name", name)
		return false
	}
	_, ok = FindPlanet(name)
	if ok {
		log.Printf("Rejecting '%s' name: matches planet name", name)
		return false
	}
	_, ok = FindPlayer(name)
	if ok {
		log.Printf("Rejecting '%s' name: matches player name", name)
		return false
	}
	_, ok = FindSystem(name)
	if ok {
		log.Printf("Rejecting '%s' name: matches star system name", name)
		return false
	}

	// Not sure if we still need this.
	// #if 0
	//    for (SystemList::const_iterator systemIter = newSystemList.begin();
	//         systemIter != newSystemList.end();
	//         systemIter++)
	//    {
	//       const System* newSystem = *systemIter;
	//
	//       if (strcasecmp(name, newSystem->name()) == 0) {
	//          log.Printf("Rejecting '%s' name: matches new system name", theName);
	//          return false;
	//       }
	//    }
	// #endif

	// Check for an object/mobile in a core system.
	fancyName := model.Name{The: false, Words: 1, Text: name}

	coreSystem, _ := FindSystem("Sol")
	if _, ok := coreSystem.FindObjectName(fancyName); ok {
		log.Printf("Rejecting '%s' name: matches Sol object", name)
		return false
	}

	coreSystem, _ = FindSystem("Arena")
	if _, ok := coreSystem.FindObjectName(fancyName); ok {
		log.Printf("Rejecting '%s' name: matches Arena object", name)
		return false
	}

	coreSystem, _ = FindSystem("Snark")
	if _, ok := coreSystem.FindObjectName(fancyName); ok {
		log.Printf("Rejecting '%s' name: matches Snark object", name)
		return false
	}

	// FIX ME! Check other core systems too.
	// Needs some kind of a kludge to cope with Horsell names.

	// Names of bodies in the Solar System.
	solNames := []string{
		"mercury",
		"venus",
		"earth",
		"moon",
		"mars",
		"phobos",
		"deimos",
		"jupiter",
		"callisto",
		"castillo", // Our favourite misspelling
		"ganymede",
		"europa",
		"io", // (Used in production game)
		"amalthea",
		"saturn",
		"titan",
		"mimas",
		"enceladus",
		"tethys",
		"dione",
		"rhea",
		"iapetus",
		"phoebe",
		"hyperion",
		"uranus",
		"titania",
		"oberon",
		"ariel",
		"umbrial",
		"miranda",
		"neptune",
		"triton",
		"nereid",
		"naiad",
		"thalassa",
		"despina",
		"galatea",
		"larissa",
		"proteus",
		"pluto",
		"charon",
	}
	for _, solName := range solNames {
		if strings.EqualFold(name, solName) {
			log.Printf("Rejecting '%s' name: matches Solar System name", name)
			return false
		}
	}

	// Names that we specifically block. Mostly we block on prefix below.
	badNames := []string{
		"alenton",
		"dataspace",
		"deity",
		"fcraig",
		"ficraig",
		"fionacraig",
		"god",
		"horsell",
		"limbo",
		"nick",
		"nickosborn",
		"nosborn",
		"osborn",
		"tellurian", // I'm so bad! :)
	}
	for _, badName := range badNames {
		if strings.EqualFold(name, badName) {
			log.Printf("Rejecting '%s' name: matches blocked name", name)
			if !global.TestFeaturesEnabled {
				return false
			}
		}
	}

	// Prefixes that we specifically block.
	badPrefixes := []string{
		"alan",
		"assist",
		"bella",
		"bug",
		"crypt",
		"cunt",
		"dbg",
		"dbug",
		"debug",
		"emperor",
		"greet",
		"fed",
		"fuck",
		"hazed",
		"help",
		"hilbert",
		"horsell",
		"host",
		"ib",
		"lenton",
		"limbo",
		"merciless",
		"ming",
		"navigator",
		"perivale",
		"piss",
		"shit",
		"snark",
		"starbase",
		"tech",
		"test",
		"theemperor",
		"thevile",
		"welcome",
	}
	for _, badPrefix := range badPrefixes {
		if strings.HasPrefix(strings.ToLower(name), badPrefix) {
			log.Printf("Rejecting '%s' name: matches blocked prefix", name)
			if !global.TestFeaturesEnabled {
				return false
			}
		}
	}

	// Segments that we specifically block.
	badSegments := []string{
		"bella",
		"bugger",
		"crypt",
		"emperor",
		"fuck",
		"lenton",
		"merciless",
		"ming",
	}
	for _, badSegment := range badSegments {
		if strings.Contains(strings.ToLower(name), badSegment) {
			log.Printf("Rejecting '%s' name: matches blocked segment", name)
			if !global.TestFeaturesEnabled {
				return false
			}
		}
	}

	// Check for anything that the lexer understands.
	if parser.IsReservedWord(name) {
		log.Printf("Rejecting '%s' name: Recognized by lexer", name)
		return false
	}

	// It's available for use.
	return true
}

func NormalizeName(name string) (string, bool) {
	debug.Precondition(name != "")

	// Is the name within allowable length limits?
	if len(name) < 3 || len(name) >= model.NAME_SIZE {
		return "", false
	}

	nameBytes := []byte(name)

	// Convert the name to lowercase, checking for alphabetic ASCII
	// characters only.
	for pos := range nameBytes {
		if !text.IsASCII(nameBytes[pos]) || !text.IsAlpha(nameBytes[pos]) {
			return "", false
		}
		nameBytes[pos] = text.ToLower(nameBytes[pos])
	}

	// Finally, capitalize the name.
	nameBytes[0] = text.ToUpper(nameBytes[0])

	// It's OK.
	return string(nameBytes), true
}
