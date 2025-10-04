package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

// Places the player in ship command mode, and moves the ship into planetary
// orbit.
//
// NOTE: There'll be a bit of a mess if someone takes off outside Sol with the
// n-space converter fitted!
func (p *Player) CmdOrbit() {
	// Is the player already flying their spaceship?
	if p.IsFlyingSpaceship() {
		p.Outputm(text.MN562)
		return
	}

	// Is player in the Command Centre?
	if !p.IsInsideSpaceship() || p.CurLocNo() != 1 /* shipCommandCentre */ {
		p.Outputm(text.MN560)
		return
	}

	// Find the corresponding orbit location.
	orbit := p.CurSys().OrbitLocNo(p.ShipLocNo())
	if orbit == 0 {
		p.Outputm(text.MN562)
		return
	}

	// Do they have any fuel?
	fuelRequired := 1 + (p.ShipKit.MaxEngine / 30)
	if p.ShipKit.CurFuel < fuelRequired {
		p.Outputm(text.INSUFFICIENT_FUEL)
		return
	}

	p.CurSys().FindLocation(p.ShipLocNo()).Talk(text.Msg(text.MN561, p.Name()))
	p.Outputm(text.LIFT_OFF, p.Name())

	p.Flags0 |= model.PL0_FLYING
	p.ShipKit.CurFuel -= fuelRequired
	p.Count[model.PL_G_JOB] += 1
	p.Count[model.PL_G_MAINT] += 1

	// debug.Check(ship_kit.cur_fuel >= 0)

	if p.Count[model.PL_G_MAINT] >= 349 { // ship needs a service!
		p.Count[model.PL_G_MAINT] = 0
		ServiceCall(p)
	}

	// Take off with the n-space converter installed?
	if (p.Flags1 & model.PL1_HILBERT) != 0 {
		p.Outputm(text.ORBIT_TO_HILBERT)
		multiTalk(text.Msg(text.SANITY_CHECK), p)

		p.SetCurSta(max(0, p.CurSta()-70))
		if p.CurSta() < 1 {
			// FIXME: needs message here.
			p.Die()
			return
		}

		// FIXME: this all needs attention...
		p.curSys, _ = FindSystem("Snark") // FIXME
		// debug.Check(m_currentSystem != NULL);
		// p.ShipLoc = Random(SnarkSystem::HilbertSpace1, SnarkSystem::HilbertSpace12)
		p.setLocation(p.ShipLoc)

		// p.checkSpyers()

		p.CurLoc().Describe(p, FullDescription)
	} else {
		// Normal take off.
		p.Outputm(text.ORBIT)
		p.ShipLoc = orbit
		p.setLocation(p.ShipLoc)
		p.CurLoc().Describe(p, DefaultDescription)
	}

	msg := text.Msg(text.ShipArrives, p.Name(), GetShipClass(p.ShipKit.Tonnage))
	p.CurLoc().Talk(msg, p)

	p.quickStatus()
	p.Save(database.SaveNow)
}
