package server

import (
	"log"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/global"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdOrder(name string) {
	if p.IsInSolLocation(sol.SlartisConstructionAndDesignWorkshop) {
		p.Outputm(text.MN1525)
		return
	}
	p.CmdOrderSpaceship()
}

func (p *Player) CmdOrderPlanet(name string) {
	// We should be in the planet shop.
	if !p.IsInSolLocation(sol.SlartisConstructionAndDesignWorkshop) {
		p.Outputm(text.ORDER_1)
		return
	}

	// Does the player already have a planet?
	if p.OwnsPlanet() {
		p.Outputm(text.ORDER_2)
		return
	}

	// Do we have the appropriate paperwork?
	if !p.HasPlanetPermit() {
		p.Outputm(text.ORDER_3)
		return
	}

	// // Make sure we don't get a second order from an Explorer.
	// if p.workbenchAccess(uid) != WB_NO_FILES {
	// 	p.Outputm(text.ORDER_4)
	// 	return
	// }

	p.Output("Not implemented. Check back in 2 weeks.\n")
}

func (p *Player) CmdOrderSpaceship() {
	if !p.curLoc.IsShipyard() {
		p.Outputm(text.MN824)
		return
	}

	if !p.HasShipPermit() {
		p.Outputm(text.MN140)
		return
	}

	if p.HasSpaceship() {
		p.Outputm(text.MN141)
		return
	}

	// model.Ranks above GroundHog get to specify the spaceship.
	if p.Rank() > model.RankGroundHog {
		log.Printf("%s is buying a ship", p.name)

		funds := p.Balance()
		if funds < 0 {
			funds = 0
		} else if funds > 100000000 {
			funds = 100000000
		}
		_ = funds // FIXME

		// TODO

		p.Output("Not implemented. Check back in 2 weeks.\n")
		return
	}

	// GroundHogs get a standard spaceship, and a promotion to Commander.

	log.Printf("Setting up starter ship for %s", p.name)
	p.ShipKit = Equipment{
		MaxHull:     10,
		CurHull:     10,
		MaxShield:   0,
		CurShield:   0,
		MaxEngine:   40,
		CurEngine:   40,
		MaxComputer: 1,
		CurComputer: 1,
		MaxFuel:     80,
		CurFuel:     80,
		MaxHold:     75,
		CurHold:     75,
		Tonnage:     200,
	}
	p.loan = 200000
	p.Missiles = 0
	p.Ammo = 0
	p.Count[model.PL_G_MAINT] = 0 // FIXME
	p.shipDesc = ""
	p.ShipLoc = sol.EarthLandingArea
	p.rank = model.RankCommander
	p.Flags0 |= model.PL0_INSURED    // Hand out free insurance
	p.Flags0 |= model.PL0_JOB        // Turn on Transportation jobs
	p.Flags0 |= model.PL0_OFFER_TOUR // Selena awaits
	p.Registry = "Panama"

	p.Outputm(text.MN143, p.name)
	p.Save(database.SaveNow)

	log.Printf("%s is now insured", p.name)

	global.ExchangeTicks++
	global.Haulers++
}
