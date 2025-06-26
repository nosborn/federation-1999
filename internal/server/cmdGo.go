package server

import (
	"log"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

// CmdGo moves a player or ship to a new location.
func (p *Player) CmdGo(direction model.Direction) {
	// Fix up the movement table if we're in the ship's airlock.
	if p.LocNo == 3 {
		p.curLoc.MovTab[model.MvOut] = p.ShipLoc
	}

	// Does that direction lead anywhere?
	if p.curLoc.MovTab[direction] == 0 {
		if p.curSys.NoExitHook(p, direction) == model.HookStop {
			return
		}

		eventNo := p.curLoc.Events[1]
		if eventNo != 0 { // TODO
			if stop := XEventHandler(p, int(eventNo)); stop {
				return
			}
		}

		p.Output(p.curLoc.noExitMessage())
		return
	}

	// Move the ship/player.
	if p.IsFlyingSpaceship() { //nolint:staticcheck
		p.moveSpaceship(direction)
	} else if !p.movePlayer(direction) {
		return
	}

	// Is she still with us?
	if p.curLoc.IsDeath() {
		p.Die()
		return
	}

	p.Save(database.SaveNow)
}

func (p *Player) movePlayer(direction model.Direction) bool {
	msgNo := [12]text.MsgNum{
		text.PlayerHasLeft_N,
		text.PlayerHasLeft_NE,
		text.PlayerHasLeft_E,
		text.PlayerHasLeft_SE,
		text.PlayerHasLeft_S,
		text.PlayerHasLeft_SW,
		text.PlayerHasLeft_W,
		text.PlayerHasLeft_NW,
		text.PlayerHasLeft,
		text.PlayerHasLeft,
		text.PlayerHasLeft,
		text.PlayerHasLeft_O,
	}

	if p.Rank() == model.RankGroundHog && p.LocNo == sol.MeetingPoint {
		if p.channel == 0 {
			p.channel = 1
		}
		p.flags2 &^= PL2_COMMS_OFF
		// debug.Check(!IsCommsOff())
	}

	var oldLocation *Location
	var suppressInEvent bool

	if p.IsInsideSpaceship() {
		oldLocation = nil
		suppressInEvent = true
	} else {
		oldLocation = p.curLoc
		suppressInEvent = false
	}

	p.CurSys().CleanupHook(p)
	p.LocNo = p.curLoc.MovTab[direction]
	p.setLocation(p.LocNo)

	// IN events on the landing pad don't fire when leaving the ship's
	// airlock. This could perhaps check that the player is moving to a
	// landing pad, but if they're not then they must be moving around
	// inside the ship and there shouldn't be an IN event anyway.

	if !suppressInEvent {
		eventNo := p.curLoc.Events[0]
		if eventNo != 0 {
			if !XEventHandler(p, int(eventNo)) {
				return false
			}
		}
	}

	if oldLocation != nil {
		msg := text.Msg(msgNo[direction], p.MoodAndName())
		oldLocation.Talk(msg, p)
	}

	if p.Rank() == model.RankAdventurer && p.GMLocation() > 0 {
		if p.CurSysName() == "Sol" && p.CurLocNo() == p.GMLocation() {
			p.SetRank(model.RankTrader)
			p.SetDeaths(0)
			p.ClearGMLocation()

			// p.checkSpyers() -- TODO
			p.Outputm(text.GMEvent)

			log.Printf("Trader promotion for %s [%d]", p.Name(), p.UID())
		}
	}

	p.curLoc.Describe(p, DefaultDescription)

	if !p.IsInsideSpaceship() {
		msg := text.Msg(text.PlayerHasArrived, p.MoodAndName())
		p.curLoc.Talk(msg, p)
	}

	if p.Rank() == model.RankCommander {
		if (p.Flags0 & model.PL0_OFFER_TOUR) != 0 {
			if p.IsInSolLocation(1 /* shipCommandCentre */) {
				p.Outputm(text.TOUR_OFFER)
				p.Flags0 &^= model.PL0_OFFER_TOUR
			}
		}
	}

	// if IsOnSnark(p) && !CheckSnarkTraps(p) {
	// 	return false
	// }

	return true
}

func (p *Player) moveSpaceship(direction model.Direction) {
	if p.ShipKit.CurFuel <= 0 {
		p.Outputm(text.MN551)
		return
	}

	p.curLoc.Talk(text.Msg(text.ShipLeaves, p.name, GetShipClass(p.ShipKit.Tonnage)), p)

	p.Count[model.PL_G_JOB]++

	if p.Count[model.PL_G_MAINT]++; (p.Count[model.PL_G_MAINT] % 350) == 349 {
		ServiceCall(p)
	}

	p.ShipLoc = p.curLoc.MovTab[direction]
	p.ShipKit.CurFuel -= (1 + p.ShipKit.MaxEngine/60)
	p.CheckShields()
	p.setLocation(p.ShipLoc)

	eventNo := p.curLoc.Events[0]
	if eventNo != 0 {
		if !XEventHandler(p, int(eventNo)) {
			return
		}
	}

	p.curLoc.Describe(p, DefaultDescription)
	p.curLoc.Talk(text.Msg(text.ShipArrives, p.name, GetShipClass(p.ShipKit.Tonnage)), p)

	if p.Target != 0 {
		target, ok := FindPlayerByID(p.Target)
		if !ok || target.Session() == nil {
			// Target has quit. This shouldn't happen!
			p.Target = 0
		} else if p.ShipLoc == target.ShipLoc {
			p.CmdFire(0, &model.Name{Text: target.Name(), The: false, Words: 1})
		}
	}

	if p.ShipKit.CurFuel < 10 {
		p.Outputm(text.WarningBuzzer)
	}
}
