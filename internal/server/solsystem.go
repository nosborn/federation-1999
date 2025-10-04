package server

import (
	"log"
	"math"
	"math/rand/v2"
	"time"

	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/monitoring"
	"github.com/nosborn/federation-1999/internal/server/global"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

const (
	publicAddressTimerPeriod = 29 * time.Second
	shuttleTimerPeriod       = 31 * time.Second
)

type stationInfo struct {
	locNo uint32
	msgNo text.MsgNum
}

type solSystem struct {
	CoreSystem

	publicAddressQuickStartIdx int
	publicAddressTimer         *time.Timer
	shuttleStationIdx          int
	shuttleTimer               *time.Timer
}

func NewSolSystem(duchy *Duchy) *solSystem {
	s := &solSystem{
		CoreSystem: CoreSystem{
			System: System{
				events:    make([]*Event, len(sol.Events)),
				locations: make([]*Location, len(sol.Locations)),
				objects:   make([]*Object, len(sol.Objects)),
				planets:   make([]*Planet, len(sol.Planets)),
				balance:   math.MaxInt32 - 1,
				duchy:     duchy,
				linkLocNo: sol.SolarSystemInterstellarLink,
				loadState: SystemOnline,
				name:      "Sol",
				taxRate:   10,
			},
		},
	}

	allSystems.Insert(s)
	systemIndex.Insert(s.name, s)
	duchy.AddMember(s)

	for i := range sol.Events {
		s.events[i] = NewCoreEvent(&sol.Events[i])
	}
	for i := range sol.Locations {
		debug.Check(sol.Locations[i].Number == uint32(i+1))
		s.locations[i] = NewCoreLocation(&sol.Locations[i], s)
	}
	for i := range sol.Objects {
		s.objects[i] = NewCoreObject(&sol.Objects[i], s)
	}
	for i := range sol.Planets {
		s.planets[i] = NewCorePlanet(s, &sol.Planets[i])
	}

	// Create timer handlers.
	s.publicAddressTimer = time.AfterFunc(publicAddressTimerPeriod, s.publicAddressTimerHandler)
	s.shuttleTimer = time.AfterFunc(shuttleTimerPeriod, s.shuttleTimerHandler)

	log.Printf("%s system initialized", s.name)
	return s
}

func (s *solSystem) CleanupHook(player *Player) {
	debug.Precondition(player != nil)

	// Fake it for any event that would normally have been triggered. I
	// suppose this could just check for an OUT event and trigger it
	// regardless. Maybe sometime I'll get around to checking whether
	// that's the case or not...

	switch player.LocNo {
	case sol.SecretRoom:
		player.StopSpying()
	case sol.Terminal:
		player.ClearDNIPassword()
	}
}

func (s *solSystem) CmdSlideHook(player *Player) bool {
	debug.Precondition(player != nil)

	switch player.CurLocNo() {
	case sol.MainStaircase, sol.MainStairway:
		player.Outputm(text.MN207)
		return true
	}
	return false
}

func (s *solSystem) NoExitHook(player *Player, direction model.Direction) model.HookResult {
	switch player.LocNo {
	case sol.GuardRoom1:
		if direction == model.DirectionEast {
			player.EnterNavalBase()
			return model.HookStop
		}
	case sol.ShuttleStation1, sol.ShuttleStation2, sol.ShuttleStation3:
		if direction == model.DirectionIn {
			player.Output("The shuttle isn't here!\n") // FIXME
			return model.HookStop
		}
	case sol.SlartisConstructionAndDesignWorkshop:
		if direction == model.DirectionWest {
			player.EnterWorkbench()
			return model.HookStop
		}
	case sol.TreasureRoom2:
		if _, ok := player.FindInventoryID(sol.ObOpal); !ok {
			player.Outputm(text.MN214)
			player.LocNo = uint32(Random(sol.MazeOfAlleys1, sol.MazeOfAlleys12))
			player.setLocation(player.LocNo)
			return model.HookStop
		}
	}
	return model.HookContinue
}

func (s *solSystem) publicAddressTimerHandler() {
	global.Lock()
	defer global.Unlock()

	monitoring.PublicAddressTimerTickTotal.WithLabelValues(s.name).Inc()

	s.publicAddressTimerProc()
}

func (s *solSystem) publicAddressTimerProc() {
	quickstart := []text.MsgNum{
		text.QuickStart_1,
		text.QuickStart_2,
		text.QuickStart_3,
		text.QuickStart_4,
		text.QuickStart_5,
		text.QuickStart_6,
		text.QuickStart_7,
	}
	general := []text.MsgNum{
		text.EarthPA_1,
		text.EarthPA_2,
		text.EarthPA_3,
		text.EarthPA_4,
		text.EarthPA_5,
		text.EarthPA_6,
		text.EarthPA_7,
	}

	for i := range Players {
		if !Players[i].IsPlaying() {
			continue
		}
		if !Players[i].IsInSolSystem() || Players[i].LocNo != sol.MeetingPoint {
			continue
		}
		if Players[i].IsOnDutyNavigator() {
			continue
		}
		switch Players[i].Rank() { //nolint:exhaustive
		case model.RankGroundHog, model.RankHostess, model.RankManager:
			Players[i].Nsoutputm(quickstart[s.publicAddressQuickStartIdx])
			Players[i].FlushOutput()
		}
	}
	s.publicAddressQuickStartIdx = (s.publicAddressQuickStartIdx + 1) % len(quickstart)

	msgID := general[rand.IntN(len(general)-1)] //nolint:gosec // "It's Just A Game"
	for i := range Players {
		if !Players[i].IsPlaying() {
			continue
		}
		if !Players[i].IsInSolSystem() {
			continue
		}
		switch Players[i].LocNo {
		case sol.StarshipCantina, sol.Terminus1, sol.Terminus2, sol.TerminusExit, sol.TerminusEntrance:
			Players[i].Nsoutputm(msgID)
			Players[i].FlushOutput()
		}
	}

	s.publicAddressTimer = time.AfterFunc(publicAddressTimerPeriod, s.publicAddressTimerHandler)
}

func (s *solSystem) shuttleTimerHandler() {
	global.Lock()
	defer global.Unlock()

	monitoring.ShuttleTimerTickTotal.WithLabelValues(s.name).Inc()

	s.ShuttleTimerProc()
}

func (s *solSystem) ShuttleTimerProc() { // FIXME: should be private
	stations := [4]stationInfo{
		{sol.ShuttleStation1, text.MN129_548}, // Cargon City
		{sol.ShuttleStation2, text.MN129_560}, // West Mine
		{sol.ShuttleStation3, text.MN129_561}, // Nisrik Mining Corporation
		{sol.ShuttleStation4, text.MN129_590}, // East Mine
	}

	if s.shuttleTimer != nil {
		s.shuttleTimer.Stop()
		s.shuttleTimer = nil
	}

	// Move out of the old location.
	oldStation := s.FindLocation(stations[s.shuttleStationIdx].locNo)
	debug.Check(oldStation != nil)
	oldStation.MovTab[model.DirectionIn] = 0
	oldStation.Talk(text.Msg(text.MN781))

	// Move into the new location.
	s.shuttleStationIdx = (s.shuttleStationIdx + 1) % len(stations)

	shuttle := s.FindLocation(sol.Shuttle)
	debug.Check(shuttle != nil)
	shuttle.MovTab[model.DirectionOut] = stations[s.shuttleStationIdx].locNo
	shuttle.Talk(text.Msg(stations[s.shuttleStationIdx].msgNo))

	newStation := s.FindLocation(stations[s.shuttleStationIdx].locNo)
	debug.Check(newStation != nil)
	newStation.MovTab[model.DirectionIn] = sol.Shuttle
	newStation.Talk(text.Msg(text.MN780))

	s.shuttleTimer = time.AfterFunc(shuttleTimerPeriod, s.shuttleTimerHandler)
}
