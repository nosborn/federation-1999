package workbench

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/text"
)

var (
	eventData    []model.DBEvent
	locationData []model.DBLocation
	objectData   []model.DBObject
)

var (
	numEvents    int
	numLocations int
	numObjects   int
)

// Checks to see that the various parameters of players' planets are within
// bounds. Returns true if no problems detected, false if a problem is
// detected.
func CheckPlanet() bool {
	var checkedOK bool

	if loadData() {
		if checkLocations() && checkEvents() && checkObjects() {
			checkedOK = true
			NeedsCheck = false
		}
	}

	numEvents = 0
	numLocations = 0
	numObjects = 0

	if !checkedOK {
		fmt.Print(text.Msg(text.Workbench_PlanetCheckFailed))
	}

	return checkedOK
}

// Checks the event file for dubious activities. Returns true if no errors
// found, false if errors found.
func checkEvents() bool {
	fmt.Print(text.Msg(text.Workbench_CheckingEvents))
	checkedOK := true

	for i := range numEvents {
		event := &eventData[i]
		eventNo := i + 1

		fmt.Print(text.Msg(text.Workbench_CheckingEvent, eventNo))

		if event.Type != 1 && event.Type != 3 && event.Type != 9 {
			fmt.Println("Invalid event type")
			checkedOK = false
			continue
		}

		if isBlank(string(event.Desc[:])) {
			fmt.Println("Event text is missing.")
			checkedOK = false
			continue
		}
		if strings.Contains(string(event.Desc[:]), "\n") {
			fmt.Println("Event text may not contain blank lines.")
			checkedOK = false
			continue
		}

		if event.NewLoc != 0 {
			if int(event.NewLoc) > numLocations {
				fmt.Println("Moves to non-existent location!")
				checkedOK = false
				continue
			}

			if event.NewLoc < 9 {
				fmt.Println("Moves to spaceship location!")
				checkedOK = false
				continue
			}

			newLocation := &locationData[event.NewLoc-1]

			if (newLocation.MapFlag & model.LfSpace) != 0 {
				fmt.Println("Moves to SPACE location!")
				checkedOK = false
				continue
			}
		}

		if event.Type == 9 {
			objectNumber := int32(math.Abs(float64(event.Field8)))
			objectFound := false

			for j := range numObjects {
				if objectData[j].Number == objectNumber {
					objectFound = true
					break
				}
			}

			if !objectFound {
				fmt.Println("Requires non-existent object!")
				checkedOK = false
				continue
			}
		}

		fmt.Println("OK")
	}

	if checkedOK {
		fmt.Print(text.Msg(text.Workbench_CheckEventsOK))
	} else {
		fmt.Print(text.Msg(text.Workbench_CheckEventsError))
	}

	return checkedOK
}

// Checks to see that the movement tables don't call out of range locations.
// Returns true if all is OK, false if out of range locs.
func checkLocations() bool {
	fmt.Print(text.Msg(text.Workbench_CheckingLocations))

	checkedOK := true
	var hospitalLocationNo, landingLocationNo, linkLocationNo, orbitLocationNo, tradeLocationNo int

	for locationIndex := 8; locationIndex < numLocations; locationIndex++ {
		location := &locationData[locationIndex]
		locationNo := locationIndex + 1

		fmt.Print(text.Msg(text.Workbench_CheckingLocation, locationNo))

		deathFlag := (location.MapFlag & model.LfDeath)
		peaceFlag := (location.MapFlag & model.LfPeace)
		spaceFlag := (location.MapFlag & model.LfSpace)

		// Check desc.

		count := 0
		allSpaces := true

		for i := 0; i < 79 && i < len(location.Desc); i++ {
			if location.Desc[i] == '\n' {
				break
			}
			// If there's only one line in the description, there
			// won't be a new-line. I'm not sure if this is the way
			// it should be!
			if location.Desc[i] == 0 {
				count = 0
				break
			}
			if location.Desc[i] > 32 && location.Desc[i] != 0xA0 { // FIXME: check for printable
				allSpaces = false
			}
			count++
		}

		if count == 0 || count == 79 || allSpaces {
			fmt.Println("No short description!")
			checkedOK = false
			continue
		}

		// Check events.

		if location.Events[0] != 0 {
			eventNo := int(location.Events[0])
			if eventNo > numEvents {
				fmt.Println("Non-existent IN event!")
				checkedOK = false
				continue
			}
			if spaceFlag == model.LfSpace {
				fmt.Println("Useless IN event in SPACE!")
				checkedOK = false
				continue
			}
		}

		if location.Events[1] != 0 {
			eventNo := int(location.Events[1])
			if eventNo > numEvents {
				fmt.Println("Non-existent OUT event!")
				checkedOK = false
				continue
			}
			if spaceFlag == model.LfSpace {
				fmt.Println("Useless OUT event in SPACE!")
				checkedOK = false
				continue
			}
		}

		// Check map_flag.

		if (location.MapFlag & model.LfHospital) != 0 {
			if hospitalLocationNo != 0 {
				fmt.Println("Can't have multiple Hospitals!")
				checkedOK = false
				continue
			}
			hospitalLocationNo = locationNo
		}

		if (location.MapFlag & model.LfLanding) != 0 {
			if landingLocationNo != 0 {
				fmt.Println("Can't have multiple Landing Pads!")
				checkedOK = false
				continue
			}
			landingLocationNo = locationNo
		}

		if (location.MapFlag & model.LfLink) != 0 {
			if linkLocationNo != 0 {
				fmt.Println("Can't have multiple Interstellar Links!")
				checkedOK = false
				continue
			}
			linkLocationNo = locationNo
		}

		if (location.MapFlag & model.LfOrbit) != 0 {
			if orbitLocationNo != 0 {
				fmt.Println("Can't have multiple Planetary Orbits!")
				checkedOK = false
				continue
			}
			orbitLocationNo = locationNo
		}

		if (location.MapFlag & model.LfTrade) != 0 {
			if tradeLocationNo != 0 {
				fmt.Println("Can't have multiple Trading Exchanges!")
				checkedOK = false
				continue
			}
			tradeLocationNo = locationNo
		}

		// Facilities can't be death locations
		facilityFlags := model.LfCafe | model.LfClth | model.LfCom | model.LfGen |
			model.LfHospital | model.LfIns | model.LfLanding | model.LfLink |
			model.LfOrbit | model.LfRep | model.LfTrade | model.LfWeap | model.LfYard
		if (location.MapFlag & facilityFlags) != 0 {
			if deathFlag == model.LfDeath {
				fmt.Println("Can't be a death location!")
				checkedOK = false
				continue
			}
		}

		if (location.MapFlag&model.LfLink) != 0 || (location.MapFlag&model.LfOrbit) != 0 {
			if peaceFlag == 0 {
				fmt.Println("Must be a peaceful location!")
				checkedOK = false
				continue
			}
			if spaceFlag == 0 {
				fmt.Println("Must be a space location!")
				checkedOK = false
				continue
			}
		}

		groundFacilities := model.LfCafe | model.LfClth | model.LfCom | model.LfGen |
			model.LfHospital | model.LfIns | model.LfLanding | model.LfRep |
			model.LfTrade | model.LfWeap | model.LfYard
		if (location.MapFlag & groundFacilities) != 0 {
			if spaceFlag == model.LfSpace {
				fmt.Println("Can't be a space location!")
				checkedOK = false
				continue
			}
		}

		if spaceFlag == model.LfSpace {
			allowedFlags := model.LfDeath | model.LfLink | model.LfOrbit | model.LfPeace | model.LfSpace | model.LfVacuum
			if (location.MapFlag &^ allowedFlags) != 0 {
				fmt.Println("One or more silly flags set with SPACE!")
				checkedOK = false
				continue
			}
		} else {
			disallowedFlags := model.LfLink | model.LfOrbit | model.LfSpace
			if (location.MapFlag & disallowedFlags) != 0 {
				fmt.Println("One or more silly flags set without SPACE!")
				checkedOK = false
				continue
			}
		}

		// Check mov_tab.

		movementOK := true

		for moveIndex := range 13 {
			if location.MovTab[moveIndex] == 0 {
				continue
			}

			if deathFlag == model.LfDeath {
				fmt.Println("Useless exit from DEATH location!")
				movementOK = false
				break
			}

			toLocationNo := int(location.MovTab[moveIndex])
			if toLocationNo > numLocations {
				fmt.Println("Movement to non-existent location!")
				movementOK = false
				break
			}

			if toLocationNo < 9 {
				fmt.Println("Movement to spaceship location!")
				movementOK = false
				break
			}

			toLocation := &locationData[toLocationNo-1]

			if (toLocation.MapFlag & model.LfSpace) != spaceFlag {
				if spaceFlag != 0 {
					fmt.Println("Movement from space to ground!")
				} else {
					fmt.Println("Movement from ground to space!")
				}
				movementOK = false
				break
			}

			if locationNo == hospitalLocationNo {
				if (toLocation.MapFlag & model.LfDeath) != 0 {
					fmt.Println("Movement from HOSPITAL to DEATH location!")
					movementOK = false
					break
				}
			}
		}

		if !movementOK {
			checkedOK = false
			continue
		}

		// Check IN/OUT vs PLANET movement restrictions
		if spaceFlag == model.LfSpace {
			if location.MovTab[10] != 0 || location.MovTab[11] != 0 { // mvIn=10, mvOut=11
				fmt.Println("IN/OUT movement is not valid in SPACE locations!")
				checkedOK = false
				continue
			}
		} else {
			if location.MovTab[12] != 0 { // mvPlanet=12
				fmt.Println("PLANET movement is only valid in SPACE locations!")
				checkedOK = false
				continue
			}
		}

		// Check sys_loc.

		if !checkSys(locationIndex) {
			fmt.Println("Out of range system message!")
			// checkedOK = false -- FIXME: uncomment this
			// continue
		}

		fmt.Println("OK")
	}

	if checkedOK {
		checkedOK = locInfo()
	}

	if checkedOK {
		if linkLocationNo != 0 && orbitLocationNo != 0 {
			fmt.Print("Checking route from Link to Orbit... ")

			if !checkRoute(linkLocationNo, orbitLocationNo) {
				checkedOK = false
			}

			fmt.Print("Checking route from Orbit to Link... ")

			if !checkRoute(orbitLocationNo, linkLocationNo) {
				checkedOK = false
			}
		}

		if landingLocationNo != 0 && tradeLocationNo != 0 && !NoExchange {
			fmt.Print("Checking route from Landing Pad to Exchange... ")

			if !checkRoute(landingLocationNo, tradeLocationNo) {
				checkedOK = false
			}

			fmt.Print("Checking route from Exchange to Landing Pad... ")

			if !checkRoute(tradeLocationNo, landingLocationNo) {
				checkedOK = false
			}
		}

		if hospitalLocationNo != 0 && landingLocationNo != 0 {
			fmt.Print("Checking route from Hospital to Landing Pad... ")

			if !checkRoute(hospitalLocationNo, landingLocationNo) {
				checkedOK = false
			}
		}
	}

	if checkedOK {
		fmt.Print(text.Msg(text.Workbench_CheckLocationsOK))
	} else {
		fmt.Print(text.Msg(text.Workbench_CheckLocationsError))
	}

	return checkedOK
}

// Checks the objects file to make sure that there is nothing untoward. Returns
// true if everything is OK, false otherwise.
func checkObjects() bool {
	fmt.Print(text.Msg(text.Workbench_CheckingObjects))

	flag := true

	for objectIndex := range numObjects {
		obj := &objectData[objectIndex]

		// common stuff
		objName := text.CStringToString(obj.Name[:])
		fmt.Printf("%-18s", objName)

		if objName == "" || !text.IsAlpha(objName[0]) {
			fmt.Println("The name must begin with a letter.")
			flag = false
			continue
		}

		length := len(objName)

		if length < 3 {
			fmt.Println("The name must be at least 3 characters long.")
			// We used to allow 2 character names even though we
			// shouldn't have. This is a bodge to avoid breaking
			// planets that have them.
			if length < 2 {
				flag = false
				continue
			}
		}

		// Check name characters
		nameOK := true
		for i := 1; i < length; i++ {
			c := objName[i]
			if !(text.IsAlnum(c) || c == '-') {
				fmt.Println("Names can only use letters, numbers and hyphens.")
				fmt.Printf("objName=%v length=%d\n", objName, length)
				flag = false
				nameOK = false
				break
			}
		}
		if !nameOK {
			continue
		}

		if isBlank(string(obj.Desc[:])) {
			fmt.Println("Description text is missing.")
			flag = false
			continue
		}
		if strings.Contains(string(obj.Desc[:]), "\n") {
			fmt.Println("Description text may not contain blank lines.")
			flag = false
			continue
		}

		if isBlank(string(obj.Scan[:])) {
			fmt.Println("Scan text is missing.")
			flag = false
			continue
		}
		if strings.Contains(string(obj.Scan[:]), "\n") {
			fmt.Println("Scan text may not contain blank lines.")
			flag = false
			continue
		}

		if int(obj.StartLoc) < 9 || int(obj.StartLoc) > numLocations {
			fmt.Printf("\n    *** Start location (%d) out of range (1-%d)! ***\n", obj.StartLoc, numLocations)
			flag = false
			continue
		}
		fmt.Print(".")

		if int(obj.GetEvent) > numEvents {
			fmt.Println("\n   *** Get event is out of range ***")
			flag = false
			continue
		}
		fmt.Print(".")

		if int(obj.GiveEvent) > numEvents {
			fmt.Println("\n   *** Give event is out of range ***")
			flag = false
			continue
		}
		fmt.Print(".")

		if int(obj.ConsumeEvent) > numEvents {
			fmt.Println("\n   *** Consume event is out of range ***")
			flag = false
			continue
		}
		fmt.Print(".")

		if int(obj.DropEvent) > numEvents {
			fmt.Println("\n   *** Drop event is out of range ***")
			flag = false
			continue
		}
		fmt.Print(".")

		if int(obj.StartLoc+obj.Offset) > numLocations {
			fmt.Println("\n   *** Recycle offset will take object out of location range! ***")
			flag = false
			continue
		}
		fmt.Print(".")

		if (obj.Flags & model.OfAnimate) == 0 {
			if obj.Value < 0 {
				fmt.Println("\n   *** Value of an object can't be negative ***")
				flag = false
				continue
			}
			fmt.Print("\n")
			continue
		}

		// mobile stuff
		if obj.MaxLoc < 9 || int(obj.MaxLoc) > numLocations {
			fmt.Println("\n   *** Maximum location for movement is out of range! ***")
			flag = false
			continue
		}
		fmt.Print(".")

		if obj.MinLoc < 9 || int(obj.MinLoc) > numLocations {
			fmt.Println("\n   *** Minimum location for movement is out of range! ***")
			flag = false
			continue
		}
		fmt.Print(".")

		if obj.MinLoc > obj.MaxLoc {
			fmt.Println("\n   *** The lowest location for movement is higher than the highest! ***")
			flag = false
			continue
		}
		fmt.Print(".")

		var spaceFlag uint32
		if (obj.Flags & model.OfShip) != 0 {
			spaceFlag = model.LfSpace
		}
		for locNo := obj.MinLoc; locNo <= obj.MaxLoc; locNo++ {
			if (locationData[locNo-1].MapFlag & model.LfSpace) == spaceFlag {
				continue
			}
			if spaceFlag == model.LfSpace {
				fmt.Println("\n   *** Moves through non-SPACE location! ***")
			} else {
				fmt.Println("\n   *** Moves through SPACE location! ***")
			}
			flag = false
			break
		}
		if !flag {
			continue
		}

		if obj.PrefObject > 0 {
			objFlag := false
			for index := range numObjects {
				if int32(obj.PrefObject) == objectData[index].Number {
					objFlag = true
					break
				}
			}
			if !objFlag {
				fmt.Printf("\n   *** Preferred object (#%d) doesn't exist! ***\n", obj.PrefObject)
				flag = false
			}
		} else {
			fmt.Print(".")
		}

		fmt.Print("\n")
	}

	if flag {
		fmt.Print(text.Msg(text.Workbench_CheckObjectsOK))
	} else {
		fmt.Print(text.Msg(text.Workbench_CheckObjectsError))
	}

	return flag
}

func checkRoute(fromLocNo, toLocNo int) bool {
	// This will mung the location and event data, so save a copy first.
	locationDataBackup := make([]model.DBLocation, len(locationData))
	copy(locationDataBackup, locationData)
	eventDataBackup := make([]model.DBEvent, len(eventData))
	copy(eventDataBackup, eventData)

	routeOK := findRoute(fromLocNo, toLocNo, false)
	if routeOK {
		fmt.Println("OK")
	} else {
		fmt.Println("Failed!")
	}

	// Restore the original data.
	copy(locationData, locationDataBackup)
	copy(eventData, eventDataBackup)

	return routeOK
}

// Check that the system messages are the one that are allowed. Returns true if
// in range, false if out of range.
func checkSys(index int) bool {
	switch index {
	case 0, 1, 2:
		return true
	case 7:
		return true
	case 12:
		return true
	case 23:
		return true
	case 26:
		return true
	case 30:
		return true
	case 32, 33:
		return true
	case 351, 352:
		return true
	default:
		log.Printf("checkSys: index=%#v", index)
		return false
	}
}

// FIXME: Well actually, not a FIXME, but review this function when making
// changes to the main event handling code. In particular, this assumes that IN
// events and DEATH flags are ignored when moved to a location by an event.
func findRoute(fromLocNo, toLocNo int, checkEntry bool) bool {
	location := &locationData[fromLocNo-1]

	if checkEntry && location.Events[0] != 0 {
		event := &eventData[location.Events[0]-1]

		if event.Field2 < 0 {
			return false // IN event removes stamina.
		}

		if event.NewLoc > 0 {
			if event.NewLoc == math.MaxUint16 {
				return false // Give up, we're looping
			}

			nextLocNo := event.NewLoc
			event.NewLoc = math.MaxUint16

			if findRoute(int(nextLocNo), toLocNo, false) {
				return true // We got there alive!
			}
		}
	}

	if checkEntry && (location.MapFlag&model.LfDeath) != 0 {
		return false // Plain DEATH location.
	}

	if fromLocNo == toLocNo {
		return true // We got there alive!
	}

	if location.Events[1] != 0 {
		event := &eventData[location.Events[1]-1]

		if event.Field2 < 0 {
			return false // OUT event removes stamina.
		}

		if event.NewLoc > 0 {
			mustFire := true

			for i := range len(location.MovTab) {
				if location.MovTab[i] != 0 {
					mustFire = false
					break
				}
			}

			if mustFire {
				if event.NewLoc == math.MaxUint16 {
					return false // Give up, we're looping
				}

				nextLocNo := event.NewLoc
				event.NewLoc = math.MaxUint16

				if findRoute(int(nextLocNo), toLocNo, false) {
					return true // We got there alive!
				}
			}
		}
	}

	for i := 0; i < len(location.MovTab); i++ {
		if location.MovTab[i] > 0 {
			if location.MovTab[i] != math.MaxUint16 {
				nextLocNo := location.MovTab[i]
				location.MovTab[i] = math.MaxUint16

				if findRoute(int(nextLocNo), toLocNo, true) {
					return true // We got there alive!
				}
			}
		}
	}

	return false // We ran out of routes to try!
}

func isBlank(line string) bool {
	return strings.TrimSpace(line) == ""
}

// Loads in the location, object and event files and sets up the max record
// numbers. Returns true if there is no problem, false if unable to open files
// or out of memory.
func loadData() bool {
	// load locations

	locFile, err := OpenLocFile()
	if err != nil {
		return false
	}
	defer locFile.Close()

	fileInfo, err := locFile.Stat()
	if err != nil {
		return false
	}
	fileSize := fileInfo.Size()

	if fileSize > 0 {
		numLocations = int(min(fileSize/model.DBLocationSize, 8+LOCATION_LIMIT))
		locationData = make([]model.DBLocation, numLocations)

		for i := range numLocations {
			err := binary.Read(locFile, binary.LittleEndian, &locationData[i])
			if err != nil {
				return false
			}
		}
	}

	_ = locFile.Close()
	fmt.Print(text.Msg(text.Workbench_CheckLocationCount, numLocations))

	// load events

	evFile, err := OpenEvFile()
	if err != nil {
		return false
	}
	defer evFile.Close()

	fileInfo, err = evFile.Stat()
	if err != nil {
		return false
	}
	fileSize = fileInfo.Size()

	if fileSize > 0 {
		numEvents = int(min(fileSize/model.DBEventSize, EVENT_LIMIT))
		eventData = make([]model.DBEvent, numEvents)

		for i := range numEvents {
			err := binary.Read(evFile, binary.LittleEndian, &eventData[i])
			if err != nil {
				return false
			}
		}
	}

	_ = evFile.Close()
	fmt.Print(text.Msg(text.Workbench_CheckEventCount, numEvents))

	// load objects

	objFile, err := OpenObjFile()
	if err != nil {
		return false
	}
	defer objFile.Close()

	fileInfo, err = objFile.Stat()
	if err != nil {
		return false
	}
	fileSize = fileInfo.Size()

	if fileSize > 0 {
		numObjects = int(min(fileSize/model.DBObjectSize, OBJECT_LIMIT))
		objectData = make([]model.DBObject, numObjects)

		for i := range numObjects {
			err := binary.Read(objFile, binary.LittleEndian, &objectData[i])
			if err != nil {
				return false
			}
		}
	}

	objFile.Close()
	fmt.Print(text.Msg(text.Workbench_CheckObjectCount, numObjects))

	return true
}

// Gives the explorer a list of the locs that have different flags set and
// check for unallowed multiples. Returns true if no unallowed multiples, false
// if unallowed multipl
func locInfo() bool {
	returnFlag := true

	// TODO: fill this out

	return returnFlag
}

// // Prints out the brief desc of the specified location.
// func printBriefDesc(locNo int) {
// 	// TODO: fill this out
// }
