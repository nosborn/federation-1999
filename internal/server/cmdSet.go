package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdSetDuchyTax(newRate int32) {
	if p.Rank() != model.RankDuke {
		p.Outputm(text.MUST_BE_DUKE)
		return
	}
	if p.ownDuchy == nil {
		p.Outputm(text.NO_DUCHY_FOR_DUKE)
		return
	}
	if newRate < 0 || newRate > 50 {
		p.Outputm(text.NiceTry)
		return
	}
	if p.ownDuchy.SetCustomsRate(newRate, p) {
		p.Save(database.SaveNow)
	}
}

// Allows a Duke to specify an embargoed planet.
func (p *Player) CmdSetEmbargo(duchyName string) {
	if p.Rank() != model.RankDuke {
		p.Outputm(text.MUST_BE_DUKE)
		return
	}
	if p.ownDuchy == nil {
		p.Output("NO_DUCHY_FOR_DUKE")
		return
	}
	if p.ownDuchy.SetEmbargo(duchyName, p) {
		p.Save(database.SaveNow)
	}
}

// Allows a Duke to specify a favoured planet.
func (p *Player) CmdSetFavoured(duchyName string) {
	if p.Rank() != model.RankDuke {
		p.Outputm(text.MUST_BE_DUKE)
		return
	}
	if p.ownDuchy == nil {
		p.Output("NO_DUCHY_FOR_DUKE")
		return
	}
	if p.ownDuchy.SetFavoured(duchyName, p) {
		p.Save(database.SaveNow)
	}
}

// Allows a Duke to specify a favoured planet.
func (p *Player) CmdSetFavouredTax(newRate int32) {
	if p.Rank() != model.RankDuke {
		p.Outputm(text.MUST_BE_DUKE)
		return
	}
	if p.ownDuchy == nil {
		p.Outputm(text.NO_DUCHY_FOR_DUKE)
		return
	}
	if newRate < 0 || newRate > 50 {
		p.Outputm(text.NiceTry)
		return
	}
	if p.ownDuchy.SetFavouredRate(newRate, p) {
		p.Save(database.SaveNow)
	}
}

func (p *Player) CmdSetLayoff(number int32) {
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

func (p *Player) CmdSetMarkup(level int32) {
	if p.Rank() < model.RankSquire || p.Rank() > model.RankDuke || p.OwnSystem() == nil {
		p.Outputm(text.NOT_A_PLANET_OWNER)
		return
	}
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

func (p *Player) CmdSetMarkupCommodity(commodity model.Commodity, level int32) {
	if p.Rank() < model.RankSquire || p.Rank() > model.RankDuke || p.OwnSystem() == nil {
		p.Outputm(text.NOT_A_PLANET_OWNER)
		return
	}
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

// Allows planet owners to set up not more than four 'milkrun' jobs for lower
// players.
func (p *Player) CmdSetMilkrun(slotNo int32, commodity model.Commodity, planetName string, carriage int32) {
	if p.Rank() < model.RankSquire || p.Rank() > model.RankDuke || p.OwnSystem() == nil {
		p.Outputm(text.NOT_A_PLANET_OWNER)
		return
	}
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

func (p *Player) CmdSetOutput(delivery model.Delivery) {
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

func (p *Player) CmdSetPlanetTax(newRate int32) {
	if p.Rank() < model.RankSquire || p.Rank() > model.RankDuke || p.ownSystem == nil {
		p.Outputm(text.NOT_A_PLANET_OWNER)
		return
	}
	if p.ownSystem.IsClosed() {
		p.Outputm(text.ClosedForBusiness, p.ownSystem.Name())
		return
	}
	if newRate < 0 {
		p.Outputm(text.NiceTry)
		return
	}
	p.ownSystem.SetTaxRate(newRate)
	p.Outputm(text.MN653, p.ownSystem.Name(), p.ownSystem.TaxRate())
	p.Save(database.SaveNow)
}

func (p *Player) CmdSetStockpile(commodity model.Commodity, level int32) {
	if p.Rank() < model.RankSquire || p.Rank() > model.RankDuke || p.ownSystem == nil {
		p.Outputm(text.NOT_A_PLANET_OWNER)
		return
	}
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

func (p *Player) CmdSetTax(rate int32) {
	if p.Rank() == model.RankDuke {
		p.CmdSetDuchyTax(rate)
	} else {
		p.CmdSetPlanetTax(rate)
	}
}

func (p *Player) CmdSetWages(number int32) {
	p.Output("Not implemented. Check back in 2 weeks.\n")
}
