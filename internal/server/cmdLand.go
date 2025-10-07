package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdLand() {
	if !p.IsFlyingSpaceship() {
		p.Outputm(text.MN556)
		return
	}
	// debug.Check(IsInsideSpaceship());

	// Check that we're in orbit. This may appear to be redundant, but it
	// lets us play tricks in Arena to prevent landing again after takeoff
	// from StarBase1.
	if !p.curLoc.IsOrbit() {
		p.Outputm(text.MN558)
		return
	}

	// Is there a landing pad corresponding to the current location?
	landing := p.CurSys().LandingLocNo(p.ShipLoc)
	if landing == 0 {
		p.Outputm(text.MN558)
		return
	}

	msg := text.Msg(text.ShipLeaves, p.Name, GetShipClass(p.ShipKit.Tonnage))
	p.curLoc.Talk(msg, p)

	p.Flags0 &^= model.PL0_FLYING
	p.setLocation(1 /* shipCommandCentre */)
	p.ShipLoc = landing
	p.LocNo = 1 /* shipCommandCentre */

	landingPad := p.CurSys().FindLocation(p.ShipLoc)
	// debug.Check(landingPad != NULL);

	fuelUsed := 1 + (p.ShipKit.MaxEngine / 30)

	if p.ShipKit.CurFuel >= fuelUsed {
		p.ShipKit.CurFuel -= fuelUsed
		landingPad.Talk(text.Msg(text.ShipLands, p.Name))
	} else {
		p.ShipKit.CurFuel = 0
		if p.ShipKit.CurHull > 1 {
			p.ShipKit.CurHull--
		}
		landingPad.Talk(text.Msg(text.ShipCrashes, p.Name))
	}

	p.Count[model.PL_G_JOB]++ // Increment jobs counter
	p.Count[model.PL_G_MAINT]++
	if (p.Count[model.PL_G_MAINT] % 500) == 499 { // Ship needs a service!
		ServiceCall(p)
		p.Count[model.PL_G_MAINT] = 0
	}

	p.Flags1 &^= model.PL1_SHIELDS

	planet := p.GuessCurrentPlanet()
	if planet != nil {
		p.Flags1 |= planet.RouteFlag()
	}

	p.Outputm(text.MN49)
	p.curLoc.Describe(p, BriefDescription)
	p.Save(database.SaveNow)
}
