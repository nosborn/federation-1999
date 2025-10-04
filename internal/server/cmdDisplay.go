package server

import (
	"log"

	"github.com/dustin/go-humanize"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/goods"
	"github.com/nosborn/federation-1999/internal/text"
)

// CmdDisplayCompany displays details of a player's company on-screen.
func (p *Player) CmdDisplayCompany(name string) {
	if name == "" {
		if p.company == nil {
			p.Outputm(text.MN1050)
			return
		}
		p.company.Display(p)
		return
	}
	if p.Rank() != model.RankDeity {
		p.UnknownCommand()
		return
	}
	company, ok := FindCompany(name)
	if !ok {
		p.Outputm(text.NoSuchCompany)
		return
	}
	company.Display(p)
}

func (p *Player) CmdDisplayDuchies() {
	for d := range allDuchies.Values() {
		switch d.Name() {
		case "Sol":
			p.Nsoutputm(text.MN1291, d.Name(), d.CustomsRate())
		case "Horsell":
			if p.Rank() != model.RankDeity {
				continue
			}
			p.Nsoutputm(text.MN1291a, d.Name(), d.CustomsRate())
		default:
			p.Nsoutputm(text.MN1292, d.Name(), d.Owner().Name(), d.CustomsRate())
			if d.Embargo() != nil {
				p.Nsoutputm(text.MN1293, d.Embargo().Name())
			}
			if d.Favoured() != nil {
				p.Nsoutputm(text.MN1294, d.Favoured().Name(), d.FavouredRate())
			}
		}
		if d.systems.Len() == 0 {
			p.Nsoutput("  No member planets.\n")
			continue
		}
		for member := range d.systems.Values() {
			if member.IsHidden() && p.Rank() != model.RankDeity {
				continue
			}
			// FIX ME! Does this need to exclude Explorer's planets
			// when they're loading for the first time?
			p.Nsoutputf("%-15s ", member.Name())
		}
		p.Nsoutput("\n")
	}
}

// CmdDisplayDuchy allows a Duke to display details of what on-line planets are
// in his/her duchy.
//
// FIXME: this is probably a bad idea for Sol because `duchy.Planet`
func (p *Player) CmdDisplayDuchy(duchyName string) {
	var duchy *Duchy
	if duchyName == "" {
		if p.Rank() != model.RankDuke {
			p.Outputm(text.MUST_BE_DUKE)
			return
		}
		if p.ownDuchy == nil {
			p.Outputm(text.NO_DUCHY_FOR_DUKE)
			return
		}
		duchy = p.ownDuchy
	} else {
		if p.Rank() != model.RankDeity {
			p.UnknownCommand()
			return
		}
		var ok bool
		duchy, ok = FindDuchy(duchyName)
		if !ok {
			p.Outputm(text.NoSuchDuchy)
			return
		}
	}
	p.Outputm(text.DisplayDuchyHeader, duchy.Name(), duchy.systems.Len())
	for member := range duchy.systems.Values() {
		msgNo := text.DisplayDuchyMember
		if member.IsClosed() {
			msgNo = text.DisplayDuchyMember_Closed
		}
		p.Outputm(msgNo, member.Name(), member.Planets()[0].LevelDescription())
	}
	p.Outputm(text.DisplayDuchyTax, duchy.TaxRate(), duchy.CustomsRate())
	if duchy.Favoured() != nil {
		p.Outputm(text.DisplayDuchyFavoured, duchy.Favoured().Name(), duchy.FavouredRate())
	}
	if duchy.Embargo() != nil {
		p.Outputm(text.DisplayDuchyEmbargo, duchy.Embargo().Name())
	}
}

func (p *Player) CmdDisplayFactories(planetName string) {
	if planetName == "" {
		if p.Rank() < model.RankSquire || p.Rank() > model.RankDuke || p.OwnSystem() == nil {
			p.Outputm(text.NOT_A_PLANET_OWNER)
			return
		}
		if p.OwnSystem().IsClosed() {
			p.Outputm(text.ClosedForBusiness, p.OwnSystem().Name())
			return
		}
		p.OwnSystem().Planets()[0].DisplayFactories(p)
		return
	}
	if p.Rank() != model.RankDeity {
		p.UnknownCommand()
		return
	}
	planet, ok := FindPlanet(planetName)
	if !ok {
		p.Outputm(text.NoSuchPlanet)
		return
	}
	planet.DisplayFactories(p)
}

// CmdDisplayFactory displays the operating details of a factory for player.
func (p *Player) CmdDisplayFactory(factoryNo int32) {
	factory, ok := p.findFactory(factoryNo)
	if !ok {
		return
	}
	if factory.IsClosed() {
		p.Outputm(text.ClosedForBusiness, factory.Planet().Name())
		return
	}
	factory.Display(p)
}

func (p *Player) CmdDisplayInformation(project model.Project) {
	template, ok := getBuildTemplate(project)
	if !ok || p.Rank() < template.LowRank {
		p.Outputm(text.DI_INFO_UNKNOWN_PROJECT)
		return
	}

	p.outputf("Commodities needed to build %s:\n", template.Name)

	for i := range template.Components {
		if template.Components[i].Quantity == 0 {
			break
		}
		// debug.Check(theTemplate->abldc[i].index >= 0);
		p.outputf("  %-15s  %6s tons\n",
			goods.GoodsArray[template.Components[i].Index].Name,
			humanize.Comma(int64(template.Components[i].Quantity)))
	}

	if template.Labour > 0 {
		p.outputf("Labor needed: %s workers\n", humanize.Comma(int64(template.Labour)))
	}

	if template.Cash > 0 {
		p.outputf("Advance costs: %s IG\n", humanize.Comma(int64(template.Cash)))
	}

	if template.Duration > 0 {
		minutes := template.Duration / 60
		p.outputf("Time to build: %s\n", text.HumanizeMinutes(int64(minutes)))
	}
}

// CmdDisplayPlanet gives a planet's parameters and currently producing
// factories.
func (p *Player) CmdDisplayPlanet(name string) {
	var planet *Planet

	if name == "" {
		if p.Rank() >= model.RankSquire && p.Rank() <= model.RankDuke {
			planet = p.OwnSystem().Planets()[0]
		}
	}
	if name == "" {
		p.Output("Not yet...\n")
		return
	}

	if planet == nil {
		planet, _ = FindPlanet(name)
	}
	if planet == nil {
		p.Outputm(text.NoSuchPlanet)
		return
	}
	if p.Rank() != model.RankDeity {
		if planet.System().IsHidden() {
			p.Outputm(text.NoSuchPlanet)
			return
		}
	}

	ownerDisplay := false
	dukeDisplay := false
	if p.Rank() >= model.RankSquire && p.Rank() <= model.RankDuke && planet.System() == p.OwnSystem() {
		ownerDisplay = true
	} else if p.Rank() == model.RankDuke && planet.System().Duchy() == p.ownDuchy {
		dukeDisplay = true
	}

	planet.Display(p, ownerDisplay, dukeDisplay)
}

// CmdDisplayProduction displays the current use of production points by a
// planet owner.
func (p *Player) CmdDisplayProduction() {
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

func (p *Player) CmdDisplayProject() {
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

// CmdDisplayRoutes lists the links to other star systems from the current
// system.
func (p *Player) CmdDisplayRoutes() {
	log.Print("cmdDisplayRoutes")

	if p.CurSys().IsHidden() {
		p.Outputm(text.MustBeJoking)
		return
	}

	p.Output("Star systems available from this link:\n")
	noDestination := true
	for s := range p.CurSys().Duchy().systems.Values() {
		if s == p.CurSys() || s.IsClosed() || s.IsHidden() {
			continue
		}
		p.outputf("%-15s ", s.Name())
		noDestination = false
	}
	if noDestination {
		p.Output("  None.")
	}
	p.Output("\n")

	if !p.CurSys().IsCapital() {
		log.Print("cmdDisplayRoutes: not capital system")
		return
	}

	p.Output("\nDuchies available from this link:\n")
	noDestination = true
	for d := range allDuchies.Values() {
		if d == p.CurSys().Duchy() {
			continue
		}
		if d.IsClosed() || d.IsHidden() {
			continue
		}
		if p.CurSys().Duchy().Embargo() == d || d.Embargo() == p.CurSys().Duchy() {
			continue
		}
		p.outputf("%-15s ", d.Name())
		noDestination = false
	}
	if noDestination {
		p.Output("  None.")
	}
	p.Output("\n")
}

// CmdDisplayWarehouse provides the player with a summary of warehouses
// currently in use.
func (p *Player) CmdDisplayWarehouses() {
	p.Output("Not implemented. Check back in 2 weeks.\n")
}
