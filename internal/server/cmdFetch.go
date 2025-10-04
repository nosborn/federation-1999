package server

import (
	"github.com/nosborn/federation-1999/internal/server/parser"
	"github.com/nosborn/federation-1999/internal/text"
)

// Fetches goods from a player's warehouse.
func (p *Player) CmdFetch(bayList *parser.BayList) {
	if !p.HasSpaceship() {
		p.Outputm(text.NoSpaceship)
		return
	}

	planet := p.GuessCurrentPlanet()
	if planet == nil { // No planet found!
		p.Outputm(text.MN39)
		return
	}

	// Get the warehouse.
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
