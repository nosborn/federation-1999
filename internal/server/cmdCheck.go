package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/text"
)

// CmdCheckCargo gives an inventory of the cargo carried by a player's ship.
func (p *Player) CmdCheckCargo() {
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

func (p *Player) CmdCheckPrice(commodity model.Commodity, duchyName string) {
	// opts := &checkPriceOptions{
	// 	duchyName: "", // FIXME: default to player current duchy
	// }
	//
	// for _, opt := range options {
	// 	opt(opts)
	// }
	//
	// if opts.duchyName != "" {
	// 	// Logic for checking price in a specific duchy...
	// 	//fmt.Printf("Checking price of %s in duchy %s\n", commodity, opts.duchyName)
	// } else {
	// 	// Logic for checking price in the local duchy...
	// 	//fmt.Printf("Checking price of %s in local duchy\n", commodity)
	// }

	p.Output("Not implemented. Check back in 2 weeks.\n")
}

// CmdCheckWarehouse inventories the contents of the warehouse on the current
// planet.
func (p *Player) CmdCheckWarehouse(planetName string) {
	var planet *Planet
	switch {
	case planetName == "":
		if p.Rank() < model.RankMerchant {
			p.Output("You need to be a Merchant before you can do that!\n")
			return
		}
		planet, _ = FindPlanet(planetName)
	case p.Rank() >= model.RankSquire && p.Rank() <= model.RankDuke:
		if p.OwnSystem() != nil {
			planet = p.OwnSystem().Planets()[0]
		}
	default:
		planet = p.GuessCurrentPlanet()
	}
	if planet == nil {
		p.Outputm(text.MN408)
		return
	}
	if p.IsInHorsellSystem() {
		// if !p.isInTime() { -- FIXME
		// 	p.Outputm(text.TemporalAnomoly)
		// 	return
		// }
		p.Outputm(text.TemporalRequest)
	}
	// if warehouse := p.findWarehouse(planet.name); warehouse == nil {
	// 	p.Output("mnNoWarehouseOnThatPlanet")
	// 	return
	// }
	// if planet.IsClosed() {
	// 	p.Outputm(text.ClosedForBusiness, planet.name)
	// 	return
	// }
	p.Output("Not implemented. Check back in 2 weeks.\n")
}
