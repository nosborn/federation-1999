package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/goods"
	"github.com/nosborn/federation-1999/internal/text"
)

// CmdAllocate allows a planet owner to allocate production points to a
// commodity, or development points.
func (p *Player) CmdAllocate(points int32, commodity model.Commodity) {
	if points == 0 {
		p.CmdDeallocate(commodity)
		return
	}

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

	available := 15 + 5*int32(p.rank-model.RankExplorer)
	if available-ownPlanet.AllocatedProductionPoints() < points {
		p.Outputm(text.MN622, available-ownPlanet.AllocatedProductionPoints())
		return
	}

	if !ownPlanet.Allocate(points, commodity) {
		p.Outputm(text.MN623)
		return
	}

	p.Outputm(text.MN624, goods.GoodsArray[commodity].Name, points)
	p.Save(database.SaveNow)
}

// CmdAllocateSocialSecurity allows a planet owner to allocate development
// points to the different categories.
func (p *Player) CmdAllocateSocialSecurity(points int32) {
	if p.rank < model.RankSquire || p.rank > model.RankDuke || p.OwnSystem() == nil {
		p.Outputm(text.NOT_A_PLANET_OWNER)
		return
	}
	if p.OwnSystem().IsClosed() {
		p.Outputm(text.ClosedForBusiness, p.OwnSystem().Name())
		return
	}
	ownPlanet := p.OwnSystem().Planets()[0]

	//
	if points <= 0 {
		p.Outputm(text.NiceTry)
		return
	}

	// Enforce a maximum of 10 points here, or nasty things will end up
	// happening to population, treasury, or disaffection!
	if points > 10 || ownPlanet.socialSec+points > 10 {
		p.Outputm(text.MN1524)
		return
	}

	ownPlanet.socialSec += points
	ownPlanet.disaffection -= min(points, ownPlanet.disaffection)

	p.Outputm(text.MN265)
	p.Save(database.SaveNow)
}
