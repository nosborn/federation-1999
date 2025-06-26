package server

import (
	"log"
	"math"

	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/server/arena"
)

type ArenaSystem struct {
	CoreSystem
}

func NewArenaSystem(duchy *Duchy) *ArenaSystem {
	s := &ArenaSystem{
		CoreSystem: CoreSystem{
			System: System{
				locations: make([]*Location, len(arena.Locations)),
				objects:   make([]*Object, len(arena.Objects)),
				planets:   make([]*Planet, 1),
				balance:   math.MaxInt32 - 1,
				duchy:     duchy,
				linkLocNo: arena.InterstellarLink,
				loadState: SystemOnline,
				name:      "Arena",
				taxRate:   30,
			},
		},
	}

	allSystems.Insert(s)
	systemIndex.Insert(s.Name(), s)
	duchy.AddMember(s)

	for i := range arena.Locations {
		debug.Check(arena.Locations[i].Number == uint32(i+9))
		s.locations[i] = NewCoreLocation(&arena.Locations[i], s)
	}
	for i := range arena.Objects {
		s.objects[i] = NewCoreObject(&arena.Objects[i], s)
	}
	s.planets[0] = NewCorePlanet(s, &arena.Planet)

	log.Printf("%s system initialized", s.Name())
	return s
}

func (s *ArenaSystem) IsArena() bool {
	return true
}

func (s *ArenaSystem) OrbitLocNo(landingLocNo uint32) uint32 {
	starbase1, ok := s.FindObject(arena.ObStarBase1)
	if !ok {
		return 0
	}
	return starbase1.curLocNo
}
