package server

import (
	"fmt"
	"log"
	"math"
	"math/rand/v2"

	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/snark"
	"github.com/nosborn/federation-1999/internal/server/sol"
)

var (
	LaunchCode        string
	LaunchCoordinates string
)

type SnarkSystem struct {
	CoreSystem
}

func NewSnarkSystem(duchy *Duchy) *SnarkSystem {
	s := &SnarkSystem{
		CoreSystem: CoreSystem{
			System: System{
				events:    make([]*Event, len(snark.Events)),
				locations: make([]*Location, len(snark.Locations)),
				objects:   make([]*Object, len(snark.Objects)),
				planets:   make([]*Planet, 1),
				balance:   math.MaxInt32 - 1,
				duchy:     duchy,
				loadState: SystemOnline,
				name:      "Snark",
				taxRate:   10,
			},
		},
	}

	allSystems.Insert(s)
	systemIndex.Insert(s.Name(), s)
	duchy.AddMember(s)

	for i := range snark.Events {
		s.events[i] = NewCoreEvent(&snark.Events[i])
	}
	for i := range snark.Locations {
		debug.Check(snark.Locations[i].Number == uint32(i+9))
		s.locations[i] = NewCoreLocation(&snark.Locations[i], s)
	}
	for i := range snark.Objects {
		s.objects[i] = NewCoreObject(&snark.Objects[i], s)
	}
	s.planets[0] = NewCorePlanet(s, &snark.Planet)

	log.Printf("%s system initialized", s.Name())
	return s
}

func (s *SnarkSystem) DrinkEvent(player *Player, object *Object) bool {
	debug.Precondition(player != nil)
	debug.Precondition(object != nil)

	if object.Number() == snark.ObPotion {
		player.DrinkPotion(object)
		return true
	}
	return s.CoreSystem.DrinkEvent(player, object)
}

func (s *SnarkSystem) IsHidden() bool {
	return true
}

func (s *SnarkSystem) IsSnark() bool {
	return true
}

func (s *SnarkSystem) RecycleHook(object *Object) model.HookResult {
	debug.Precondition(object != nil)
	debug.Precondition(object.HomeSystem() == s)

	if object.number == snark.ObConverter {
		tuningFork, _ := s.FindObject(snark.ObTuningFork)
		if !tuningFork.IsRecycling() {
			tuningFork.Recycle()
		}
		return model.HookStop
	}
	return model.HookContinue
}

func (s *SnarkSystem) UseHook(player *Player, object *Object) model.HookResult {
	debug.Precondition(player != nil)
	debug.Precondition(object != nil)

	if object.number == sol.ObSoap {
		_, ok := player.curLoc.FindObject(snark.ObLever)
		if ok {
			s.useSoap(player, object)
			return model.HookStop
		}
	}

	return model.HookContinue
}

func (s *SnarkSystem) useSoap(player *Player, soap *Object) {
	// TODO
}

func GetLaunchCode() string {
	return LaunchCode
}

func GetLaunchCoordinates() string {
	return LaunchCoordinates
}

func setLaunchCode() {
	code := (rand.IntN(9)+1)*10000 + //nolint:gosec // "It's Just A Game"
		(rand.IntN(9)+1)*1000 + //nolint:gosec // "It's Just A Game"
		(rand.IntN(9)+1)*100 + //nolint:gosec // "It's Just A Game"
		(rand.IntN(9)+1)*10 + //nolint:gosec // "It's Just A Game"
		(rand.IntN(9) + 1) //nolint:gosec // "It's Just A Game"
	LaunchCode = fmt.Sprintf("%d", code)
	// debug.Check(dbgIsValidString(launchCode))
	log.Printf("Snark launch code set to %s", LaunchCode)
}

func setLaunchCoordinates() {
	LaunchCoordinates = fmt.Sprintf("%03d:%03d", rand.IntN(360), rand.IntN(360)) //nolint:gosec // "It's Just A Game"
	// debug.Check(dbgIsValidString(launchCoordinates));
	log.Printf("Snark launch co-ordinates set to %s", LaunchCoordinates)
}

// Takes 10 stam off when player moves through strange light zone on Snark.
// func snarkSapStamina(p *Player) bool {
// 	p.Sta.Cur = max(0, p.Sta.Cur-10)
// 	p.Outputm(text.MN240)
// 	if p.Sta.Cur < 1 {
// 		p.Die()
// 		return false
// 	}
// 	return true
// }

// Deals with possible damage to a player's health when a missile is launched
// in his/her vicinity.
// func snarkSiloDamage(p *Player) bool {
// 	if rand.IntN(100) < 75 { //nolint:gosec // "It's Just A Game"
// 		p.Sta.Cur = max(0, p.Sta.Cur-50)
// 		p.Outputm(text.MN241)
//
// 		if p.Sta.Cur < 1 {
// 			p.Die()
// 			return false
// 		}
// 	}
// 	return true
// }
