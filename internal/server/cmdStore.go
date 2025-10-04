package server

import (
	"github.com/nosborn/federation-1999/internal/server/parser"
	"github.com/nosborn/federation-1999/internal/text"
)

// Stores goods from a ship to a warehouse.
func (p *Player) CmdStore(bayList *parser.BayList) {
	// It helps to have a spaceship for this!
	if !p.HasSpaceship() {
		p.Outputm(text.NoSpaceship)
		return
	}

	// Find the current planet...
	planet := p.GuessCurrentPlanet()
	if planet == nil {
		p.Outputm(text.MN39)
		return
	}

	// ...and the warehouse there.
	warehouse := p.FindWarehouse(planet.Name())
	if warehouse == nil {
		p.Outputm(text.NoWarehouseOnThisPlanet)
		return
	}

	//
	if planet.IsClosed() {
		p.Outputm(text.ClosedForBusiness, planet.Name())
		return
	}

	//
	if bayList.Size > len(bayList.Bay) {
		p.outputf("The stevedores can't cope with more than %d bays at once!\n", len(bayList.Bay))
		return
	}

	// TODO:
	p.Output("Not implemented. Check back in 2 weeks.\n")
}
