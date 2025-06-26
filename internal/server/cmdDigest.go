package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdDigest1(group model.CommodityGroup) {
	if p.Rank() < model.RankSquire || p.Rank() > model.RankDuke || p.OwnSystem() == nil {
		p.Outputm(text.NOT_A_PLANET_OWNER)
		return
	}
	if p.OwnSystem().IsClosed() {
		p.Outputm(text.ClosedForBusiness, p.OwnSystem().Name())
		return
	}
	p.OwnSystem().Planets()[0].Digest(group, p)
}

func (p *Player) CmdDigest2(group model.CommodityGroup, planetName string) {
	if p.Rank() != model.RankDeity {
		p.UnknownCommand()
		return
	}
	planet, ok := FindPlanet(planetName)
	if !ok {
		p.Outputm(text.NoSuchPlanet)
		return
	}
	if planet.IsClosed() {
		p.Outputm(text.ClosedForBusiness, planet.name)
		return
	}
	planet.Digest(group, p)
}
