package server

import (
	"log"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdSellCargo(bayNo int32) {
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

// CmdSellFactory allows a player to sell an unwanted factory.
func (p *Player) CmdSellFactory(factoryNo int32) { // FIXME: optional factoryNo
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

// CmdSellSpaceship sells a player's ship in the auction...
func (p *Player) CmdSellSpaceship() {
	if !p.HasSpaceship() {
		p.Outputm(text.NoSpaceship)
		return
	}

	if p.loan > 0 {
		p.Outputm(text.MN41)
		return
	}

	if p.IsInsideSpaceship() {
		p.Outputm(text.LEAVE_SPACESHIP)
		return
	}

	if p.curSys.IsHorsell() { // FIXME: Check for Snark too.
		p.Outputm(text.SellSpaceshipInHorsell)
		return
	}

	value := shipValue(p) / 2
	p.Outputm(text.MN137)
	p.Outputm(text.MN821, value/5, value/4, value/3, value/2, value/2, (3*value)/4, value)
	for range 3 {
		// if (random() % 100) > 60 { -- FIXME
		// 	break
		// }
		value += value / 10
		p.Outputm(text.MN822, value)
	}
	p.Outputm(text.MN823, value)

	if p.HasWallet() {
		p.ChangeBalance(value)
	}

	p.ShipLoc = 0
	p.shipDesc = ""
	p.Registry = ""
	// memset(&ship_kit, '\0', sizeof(ship_kit)); -- FIXME
	p.Missiles = 0
	p.Ammo = 0
	// memset(guns, '\0', sizeof(guns)); -- FIXME
	p.Target = 0

	if p.HasSpyBeam() {
		p.destroySpyBeam()

		if p.HasWallet() {
			p.ChangeBalance(model.SPYBEAM_RESALE)
			p.Outputm(text.SpybeamSold, model.SPYBEAM_RESALE)
		}
	}

	// Just in case a wannabe-GM tries to be clever!
	p.Flags1 &^= model.PL1_HILBERT

	// Reset the navigation computer.
	p.Flags1 &^= model.PL1_TITAN
	p.Flags1 &^= model.PL1_CALLISTO
	p.Flags1 &^= model.PL1_MARS
	p.Flags1 &^= model.PL1_EARTH
	p.Flags1 &^= model.PL1_MOON
	p.Flags1 &^= model.PL1_VENUS
	p.Flags1 &^= model.PL1_MERCURY

	p.Save(database.SaveNow)
	log.Printf("%s has sold their ship", p.Name())
}

// CmdSellWarehouse allows a player to sell a warehouse.
func (p *Player) CmdSellWarehouse(planetName string) {
	p.Output("Not implemented. Check back in 2 weeks.\n")
}
