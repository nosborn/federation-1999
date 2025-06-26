package server

import (
	"log"
	"math"
	"strings"
	"time"

	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/global"
	"github.com/nosborn/federation-1999/internal/server/horsell"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
	"github.com/nosborn/federation-1999/pkg/ibgames"
)

type HorsellSystem struct {
	CoreSystem

	artilleryBarrage   bool
	artillerymanPlaced bool
	organActive        bool
	shieldRaised       bool
	timeUp             bool
	timer              *time.Timer
	warper             ibgames.AccountID
}

func NewHorsellSystem(duchy *Duchy, warper *Player) *HorsellSystem {
	s := &HorsellSystem{
		CoreSystem: CoreSystem{
			System: System{
				locations: make([]*Location, len(horsell.Locations)),
				objects:   make([]*Object, len(horsell.Objects)),
				planets:   make([]*Planet, 1),
				balance:   math.MaxInt32 - 1,
				duchy:     duchy,
				loadState: SystemOnline,
				name:      duchy.Name(),
				taxRate:   30,
			},
		},
		warper: warper.UID(),
	}

	allSystems.Insert(s)
	systemIndex.Insert(s.name, s)
	duchy.AddMember(s)

	log.Printf("%s system created for %s", s.name, warper.Name())

	for i := range horsell.Locations {
		debug.Check(horsell.Locations[i].Number == uint32(i+9))
		s.locations[i] = NewCoreLocation(&horsell.Locations[i], s)
	}
	for i := range horsell.Objects {
		s.objects[i] = NewCoreObject(&horsell.Objects[i], s)
	}
	s.planets[0] = NewCorePlanet(s, &horsell.Planet)

	return s
}

func (s *HorsellSystem) artillerymansTale(player *Player, dogtag, artilleryman *Object) {
	// TODO
}

func (s *HorsellSystem) astronomersTale(player *Player, marsmetal, astronomer *Object) {
	// TODO
}

func (s *HorsellSystem) BlastHook(player *Player, what model.Name, tdx *Object) model.HookResult {
	if player.LocNo != horsell.Cylinder2 {
		player.UnknownCommand()
		return model.HookStop
	}
	if !strings.EqualFold(what.Text, "nest") {
		player.UnknownCommand()
		return model.HookStop
	}
	s.useTDX(player, tdx)
	return model.HookStop
}

func (s *HorsellSystem) Destroy(reason string) {
	// TODO

	s.duchy.RemoveMember(s)
	s.duchy = nil

	allSystems.Remove(s)
	systemIndex.Remove(s.name)
}

func (s *HorsellSystem) GiveHook(player *Player, object, mobile *Object) model.HookResult {
	debug.Precondition(player != nil)
	debug.Precondition(object != nil && !object.IsMobile())
	debug.Precondition(mobile != nil && mobile.IsMobile())

	debug.Trace("HorsellSystem::giveHook %s", s.name)

	if object.Number() == horsell.ObDogTag && mobile.Number() == horsell.ObArtilleryman {
		s.artillerymansTale(player, object, mobile)
		return model.HookStop
	}
	if object.Number() == horsell.ObMarsmetal && mobile.Number() == horsell.ObAstronomer {
		s.astronomersTale(player, object, mobile)
		return model.HookStop
	}
	return model.HookContinue
}

func (s *HorsellSystem) IsHidden() bool {
	return true
}

func (s *HorsellSystem) IsInCrater(locationNo uint32) bool {
	return locationNo >= horsell.Crater1 && locationNo <= horsell.Crater3
}

func (s *HorsellSystem) martianVictory() {
	if s.shieldRaised { //nolint:staticcheck // SA9003: empty branch
		// TODO
	}
	s.Talk(text.Msg(text.MartianVictory))
	s.StopPuzzle()
}

func (s *HorsellSystem) NoExitHook(player *Player, direction model.Direction) model.HookResult {
	if player.LocNo == horsell.Crater3 && s.shieldRaised {
		if direction == model.DirectionNorth || direction == model.DirectionNW {
			opal, ok := player.FindInventoryID(sol.ObOpal)
			if !ok {
				player.Outputm(text.NoExitEnergyCurtain)
			} else {
				// useOpal(player, opal) -- FIXME
				_ = opal // FIXME
			}
		} else {
			player.Outputm(text.NoExitHeatRay)
		}
		return model.HookStop
	}
	return model.HookContinue
}

// func (s *HorsellSystem) playOrgan(player *Player, music string) {
// 	// TODO
// }

func (s *HorsellSystem) RecycleHook(object *Object) model.HookResult {
	// Nothing recycles in Horsell, it all goes into hiding.
	object.Hide()
	return model.HookStop
}

func (s *HorsellSystem) StartPuzzle() {
	dogtag, _ := s.FindObject(horsell.ObDogTag)
	debug.Check(dogtag != nil && dogtag.IsHidden())

	// Place the dog tag.
	debug.Trace("%s: Placing the dog tag", s.Name())
	dogtag.Unhide()
	dogtag.Place()

	// Do the green flash thing.
	s.Talk(text.Msg(text.MN359))

	// The next timer event is the artillery barrage.
	s.timer = time.AfterFunc(8*60*time.Second, s.timerHandler)
}

// If a Horsell system has to be created for a player at logon, this function
// is called to put the puzzle into its final segment. Use martianVictory() to
// stop the puzzle on other occasions.
func (s *HorsellSystem) StopPuzzle() {
	// Ensure the artilleryman has been placed.
	if !s.artillerymanPlaced {
		artilleryman, ok := s.FindObject(horsell.ObArtilleryman)
		debug.Check(artilleryman != nil)
		if ok && artilleryman.IsHidden() {
			debug.Trace("%s: Placing the artilleryman", s.name)
			artilleryman.Unhide()
			artilleryman.Place()
			debug.Check(artilleryman.curLocNo == horsell.Cellar)
		}
		s.artillerymanPlaced = true
	}

	// Fiddle with the crater locations.
	location := s.FindLocation(horsell.Cylinder1)
	location.SetBriefMsgNo(text.HorsellLocation82_P3_Brief)
	location.SetFullMsgNo(text.HorsellLocation82_P3_Full)

	location = s.FindLocation(horsell.Cylinder2)
	location.SetBriefMsgNo(text.HorsellLocation83_P3_Brief)
	location.SetFullMsgNo(text.HorsellLocation83_P3_Full)

	location = s.FindLocation(horsell.Crater3)
	location.SetBriefMsgNo(text.HorsellLocation84_P3_Brief)
	location.SetFullMsgNo(text.HorsellLocation84_P3_Full)
	location.MovTab[model.MvNorth] = horsell.Cylinder2
	location.MovTab[model.MvNW] = horsell.Cylinder1
	location.MovTab[model.MvDown] = horsell.Cylinder2
	location.MovTab[model.MvIn] = horsell.Cylinder1

	// Hide the marsmetal.
	marsmetal, ok := s.FindObject(horsell.ObMarsmetal)
	debug.Check(marsmetal != nil)

	if ok && marsmetal.curLocNo == horsell.Cylinder1 {
		debug.Trace("%s: Hiding the marsmetal", s.name)
		marsmetal.Hide()
		debug.Check(marsmetal.curLocNo == 0)
	}

	//
	s.artilleryBarrage = true
	s.organActive = false
	s.shieldRaised = false
	s.timeUp = true

	// Kill any existing timer and set up for destroying out the system.
	if s.timer != nil {
		s.timer.Stop()
		s.timer = nil
	}
	s.timer = time.AfterFunc(5*60*time.Second, s.timerHandler)
}

func (s *HorsellSystem) timerHandler() {
	global.Lock()
	defer global.Unlock()

	s.timerProc()
}

func (s *HorsellSystem) timerProc() {
	s.timer = nil

	// Check for destroying the system first.
	if s.timeUp {
		s.Destroy(text.Msg(text.TimewarpTimeUp))
		return
	}

	// Do the artillery barrage and set up for placing the artilleryman.
	if !s.artilleryBarrage {
		s.Talk(text.Msg(text.HorsellPuzzleLastSegment)) // Do something about name
		s.artilleryBarrage = true
		s.timer = time.AfterFunc(2*60*time.Second, s.timerHandler)
		return
	}

	// Place the artilleryman and set up for running out of time.
	if !s.artillerymanPlaced {
		artilleryman, ok := s.FindObject(horsell.ObArtilleryman)
		debug.Check(artilleryman != nil)

		if ok && artilleryman.IsHidden() {
			debug.Trace("%s: Placing the artilleryman", s.Name())
			artilleryman.Unhide()
			artilleryman.Place()
			debug.Check(artilleryman.curLocNo == horsell.Cellar)

			cellar := s.FindLocation(artilleryman.curLocNo)
			debug.Check(cellar != nil)

			if cellar != nil {
				msg := text.Msg(text.MOBILE_ARRIVAL, artilleryman.DisplayName(true))
				cellar.Talk(msg)
			}
		}

		s.artillerymanPlaced = true
		s.timer = time.AfterFunc(15*60*time.Second, s.timerHandler)
		return
	}

	// Stop the puzzle. This sets up for destroying out the system as well.
	s.martianVictory()
}

func (s *HorsellSystem) UseHook(player *Player, object *Object) model.HookResult {
	debug.Precondition(player != nil)
	debug.Precondition(object != nil)

	switch player.CurLocNo() {
	case horsell.Cylinder2:
		if object.Number() == sol.ObTDX {
			s.useTDX(player, object)
			return model.HookStop
		}
	case horsell.Crater3:
		if object.Number() == sol.ObOpal && s.shieldRaised {
			s.useOpal(player, object)
			return model.HookStop
		}
	}
	return model.HookContinue
}

func (s *HorsellSystem) useOpal(player *Player, opal *Object) {
	debug.Precondition(player != nil)
	debug.Precondition(opal != nil)

	player.Outputm(text.UseOpal)
	player.LocNo = horsell.Cylinder1
	player.setLocation(player.LocNo)
	player.curLoc.Describe(player, DefaultDescription)
	player.Save(database.SaveNow)
}

func (s *HorsellSystem) useTDX(player *Player, tdx *Object) {
	// TODO
}
