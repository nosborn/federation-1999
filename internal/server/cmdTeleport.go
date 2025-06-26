package server

import (
	"slices"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/snark"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

// Teleports a player to the specified location in the game.
func (p *Player) CmdTeleport(locNo int32, sysName string) {
	// Do we have a functioning teleporter?
	if p.Rank() < model.RankDuke {
		for i := range []int{model.DK_UPSIDE, model.DK_DOWNSIDE, model.DK_MATTRANS} {
			if p.Facilities[i] < 100 {
				p.Outputm(text.TeleportNoFacilities)
				return
			}
		}
	}

	// Are we in Horsell?
	if p.CurSys().IsHorsell() && p.Rank() < model.RankManager {
		p.Outputm(text.TeleportInTime)
		return
	}

	// Figure out which star system we're going to.
	s := p.CurSys()
	if sysName != "" {
		var ok bool
		s, ok = FindSystem(sysName)
		if !ok {
			p.Outputm(text.TeleportInvalidAddress)
			return
		}
		// Are we trying to get to a closed system?
		if s.IsClosed() {
			if p.Rank() < model.RankHostess {
				p.Outputm(text.LINK_CLOSED, s.Name)
				return
			}
			if !s.IsHorsell() {
				p.Outputm(text.LINK_CLOSED, s.Name)
				return
			}
		}
	}

	// Redirect anything in the (ruined) Martian ruins to the crater.
	// if system.IsSol() && solSystem->areMartianRuinsBlown()) {
	// 	if (locNo >= SolSystem::BeforeTheRuins && locNo <= SolSystem::MazeOfAlleys12) {
	// 		locNo = SolSystem::Crater
	// 	}
	// }

	// Can we find the destination location?
	l := s.FindLocation(uint32(locNo))
	if l == nil {
		p.Outputm(text.TeleportInvalidAddress)
		return
	}

	// Are we inside our spaceship?
	if p.IsInsideSpaceship() {
		p.Outputm(text.TeleportInShip)
		return
	}

	// Are we holding any objects?
	if len(p.inventory) > 0 && p.Rank() != model.RankDeity {
		p.Outputm(text.TeleportHasObjects)
		return
	}

	if s.IsSol() && isTeleportShielded(uint32(locNo)) && p.Rank() < model.RankDeity {
		p.Outputm(text.TeleportIsShielded)
		return
	}
	if s.IsSnark() && locNo == snark.Storeroom && p.Rank() < model.RankDeity {
		p.Outputm(text.TeleportIsShielded)
		return
	}

	// Is the ship carrying goods?
	if p.HasSpaceship() {
		if p.Job.Status == JOB_COLLECTED {
			p.Outputm(text.TeleportHasCargo)
			return
		}
		for i := range p.Load {
			if p.Load[i].Quantity > 0 {
				p.Outputm(text.TeleportHasCargo)
				return
			}
		}
	}

	// No teleporting to hidden locations.
	if l.IsHidden() && p.Rank() < model.RankDeity {
		p.Outputm(text.TeleportInvalidAddress)
		return
	}

	// No teleporting to shielded locations. Staff can override shielding
	// so long as they're not trying to teleport into their ship.
	if l.IsShielded() {
		if p.Rank() < model.RankHostess || locNo <= SPACESHIP_SIZE {
			p.Outputm(text.TeleportIsShielded)
			return
		}
	}

	// No teleporting to locations with IN events.
	if l.Events[0] != 0 {
		p.Outputm(text.MN440)
		return
	}

	// Do any cleaning up prior to leaving the location.
	p.CurSys().CleanupHook(p)

	// Deal with on-duty Navigators leaving Sol.
	if p.IsOnDutyNavigator() && !s.IsSol() {
		p.StopNavigating()
		p.FlushOutput()
	}

	// OK - lets move the player...
	msg := text.Msg(text.TeleportDepartureTell, p.MoodAndName())
	p.curLoc.Talk(msg, p)

	//
	p.SetCurSys(s)
	p.LocNo = uint32(locNo)
	p.setLocation(p.LocNo)

	// ...and move the spaceship.
	if p.HasSpaceship() {
		if s.IsSol() {
			p.ShipLoc = FindSolLandingPad(uint32(locNo))
		} else {
			p.ShipLoc = s.Planets()[0].landingLocNo
		}
	}

	p.curLoc.Describe(p, DefaultDescription)

	//
	if p.curLoc.IsDeath() {
		p.Die()
		return
	}
	if p.curLoc.IsSpace() {
		p.Outputm(text.TeleportToSpaceLocation)
		p.Die()
		return
	}
	p.Save(database.SaveNow)

	msg = text.Msg(text.TeleportArrivalTell, p.MoodAndName())
	p.curLoc.Talk(msg, p)

	//
	// currentSystem()->m_populated = Transaction::time(); -- TODO
}

// Checks to see if the player is trying to teleport into a barred location.
func isTeleportShielded(locNo uint32) bool {
	barred := []uint32{
		sol.Office2,           // 217 - Moon    - Office
		sol.Courtyard2,        // 299 - Mars    - Courtyard
		sol.TreasureRoom1,     // 300 - Mars    - Treasure room
		sol.TreasureRoom2,     // 301 - Mars    - Treasure room
		sol.ControlRoom3,      // 421 - Earth   - Control room
		sol.SecretRoom,        // 467 - Earth   - Secret room     - Public spybeam
		sol.Terminal,          // 470 - Earth   - Terminal
		sol.UnusedLocation509, // 509 - Earth   - Deserted gates  - Spare location
		sol.Corridor43,        // 564 - Mercury - Corridor
		sol.DrugProcessing,    // 587 - Mercury - Drug processing
		sol.StoreRoom2,        // 588 - Mercury - Store room
		sol.LivingQuarters3,   // 589 - Mercury - Living quarters
		sol.PoliceCell,        // 596 - Mercury - Police cell
		sol.DieselsBoudoir,    // 697 - Mars    - Diesel's boudoir
	}

	// Spaceship, Snark or Horsell?
	if locNo <= SPACESHIP_SIZE || locNo >= sol.ThroughTheLookingGlass /*698*/ {
		return true
	}

	// Specifically barred locations?
	if slices.Contains(barred, locNo) {
		return true
	}

	// Martian ruins when they're blown?
	// if solSystem.areMartianRuinsBlown() {
	// 	if locNo >= sol.Archway1 /*267*/ && locNo <= sol.MazeOfAlleys12 /*367*/ {
	// 		return true
	// 	}
	// }

	// It's OK.
	return false
}
