package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/text"
)

// Create a job for a high level player to offer to a low level player.
func (p *Player) CmdOffer(haulerName, fromPlanetName string, bayNo int32, toPlanetName string, time, carriage int32, toWarehouse bool) {
	if p.rank < model.RankJP {
		p.Outputm(text.MN939)
		return
	}

	// Find the hauler.
	hauler, ok := FindPlayer(haulerName)
	if !ok || !hauler.IsPlaying() {
		p.Outputm(text.MN36)
		return
	}
	if hauler == p {
		p.Outputm(text.MustBeJoking)
		return
	}

	// Check that the from planet is accessible.
	fromPlanet, ok := FindPlanet(fromPlanetName)
	if !ok {
		p.Outputm(text.MN944)
		return
	}
	if fromPlanet.IsClosed() {
		p.Outputm(text.ClosedForBusiness, fromPlanet.Name())
		return
	}

	// Find the warehouse on the from planet.
	warehouse := p.FindWarehouse(fromPlanet.Name())
	if warehouse == nil {
		p.Outputm(text.NoWarehouseOnThatPlanet)
		return
	}

	// Check for a valid bay number.
	if bayNo < 1 || bayNo > 20 {
		p.Outputm(text.BadWarehouseBayNo)
		return
	}

	// See if there's anything in that warehouse bay.
	quantity := warehouse.Bay[bayNo-1].Quantity
	if quantity == 0 {
		p.Outputm(text.MN473)
		return
	}

	// Check that the to planet is accessible.
	toPlanet, ok := FindPlanet(toPlanetName)
	if !ok {
		p.Outputm(text.MN944)
		return
	}
	if toPlanet.IsClosed() {
		p.Outputm(text.ClosedForBusiness, toPlanet.Name())
		return
	}
	if fromPlanet.Name() == toPlanet.Name() {
		p.Outputm(text.MN527)
		return
	}

	// Check the GTUs for sanity.
	if time < 1 {
		p.Outputm(text.MN947)
		return
	}

	// Check the carriage for sanity.
	if carriage < 1 {
		p.Outputm(text.MN946)
		return
	}
	if carriage > 30 {
		p.Outputm(text.OfferCarriageCap)
		return
	}

	// Check that the player can afford the payment.
	payment := carriage * quantity
	if payment > p.balance {
		p.Outputm(text.MN25)
		return
	}

	//
	if !toPlanet.HasExchange() && !toWarehouse {
		p.Outputm(text.NO_EXCHANGE_ON_PLANET, toPlanet.Name())
		return
	}

	// Now let's find out if the hauler's in a position to take the job.
	if !hauler.HasSpaceship() {
		p.Outputm(text.MN941, hauler.Name())
		return
	}
	if hauler.Job.Status != JOB_NONE {
		p.Outputm(text.MN942, hauler.Name())
		return
	}
	if hauler.ShipKit.CurHold < quantity {
		p.Outputm(text.MN948)
		return
	}

	// TODO:
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

// Allows a planet owner whose planet is on line to write a planet job directly
// into transportation.
func (p *Player) CmdOfferJob(commodity model.Commodity, toPlanetName string, payment int32) {
	if p.rank < model.RankSquire || p.rank > model.RankDuke || p.OwnSystem() == nil {
		p.Outputm(text.NOT_A_PLANET_OWNER)
		return
	}
	if p.OwnSystem().IsClosed() {
		p.Outputm(text.ClosedForBusiness, p.OwnSystem().Name())
		return
	}

	ownPlanet := p.OwnSystem().Planets()[0]
	if !ownPlanet.HasExchange() {
		p.Outputm(text.NO_EXCHANGE_ON_PLANET, ownPlanet.Name())
		return
	}
	if ownPlanet.ActualStock(commodity) < 250 {
		p.Outputm(text.MN1130)
		return
	}

	toPlanet, ok := FindPlanet(toPlanetName)
	if !ok || toPlanet.IsHidden() {
		p.Outputm(text.NoSuchPlanet)
		return
	}
	if !toPlanet.HasExchange() {
		p.Outputm(text.NO_EXCHANGE_ON_PLANET, toPlanet.Name())
		return
	}
	if payment < 10 || payment > 30 {
		p.Outputm(text.MN1132)
		return
	}

	// FIXME:
	// if isEmbargoed(toPlanet, ownPlanet) { // Is an embargo in effect?
	// 	p.Outputm(text.MN1198, toPlanet.Name())
	// 	return
	// }

	// TODO:
	p.Output("Not implemented. Check back in 2 weeks.\n")
}
