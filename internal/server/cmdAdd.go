package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/parser"
	"github.com/nosborn/federation-1999/internal/text"
)

// CmdAdd combines the contents of two bays in the specified warehouse.
func (p *Player) CmdAdd(bayList *parser.BayList, planetName string) {
	// FIXME: there must be a better solution!
	if bayList.Size < 2 {
		p.UnknownCommand()
		return
	}

	//
	if p.rank < model.RankMerchant {
		p.Outputm(text.MN446)
		return
	}

	//
	var planet *Planet
	if planetName == "" {
		if p.rank >= model.RankSquire && p.rank <= model.RankDuke {
			if p.OwnSystem() != nil {
				planet = p.OwnSystem().Planets()[0]
			}
		} else {
			planet = p.GuessCurrentPlanet()
		}
	} else {
		planet, _ = FindPlanet(planetName)
	}
	if planet == nil {
		p.Outputm(text.MN447)
		return
	}

	// Find the warehouse on that planet and ensure it's open.
	warehouse := p.FindWarehouse(planet.Name())
	if warehouse == nil {
		p.Outputm(text.NoWarehouseOnThatPlanet)
		return
	}
	if planet.IsClosed() {
		p.Outputm(text.ClosedForBusiness, planet.Name())
		return
	}

	//
	if bayList.Size > len(bayList.Bay) {
		p.outputf("The stevedores can't cope with more than %d bays at once!\n", len(bayList.Bay))
		return
	}

	// Check the destination bay number for sanity.
	toBayNo := bayList.Bay[0]
	if toBayNo < 1 || toBayNo > 20 {
		p.Outputm(text.BadWarehouseBayNo)
		return
	}

	//
	var maxQuantity int
	if p.rank >= model.RankSquire && p.rank <= model.RankDuke {
		maxQuantity = 1000000
	} else {
		maxQuantity = 300
	}
	_ = maxQuantity // FIXME

	// TODO:
	p.Output("Not implemented. Check back in 2 weeks.\n")
}
