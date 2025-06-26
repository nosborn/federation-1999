package server

import (
	"github.com/dustin/go-humanize"
	"github.com/nosborn/federation-1999/internal/model"
)

// Reports a players current score.
func (p *Player) CmdScore() {
	var sex string
	switch p.Sex() {
	case model.SexFemale:
		sex = "Female"
	case model.SexMale:
		sex = "Male"
	default:
		sex = "Neuter"
	}
	p.outputf("Name: %s   Sex: %s   Rank: %s\n", p.Name(), sex, p.rankName())

	if p.HasWallet() {
		p.outputf("Bank balance: %s IG\n", humanize.Comma(int64(p.Balance())))
	}

	if p.IsInsured() {
		p.Output("You are insured\n")
	} else {
		p.Output("You are -* NOT *- insured!\n")
	}

	p.outputf("Reward value: %s IG   Games played: %s\n",
		humanize.Comma(int64(p.Reward())),
		humanize.Comma(int64(p.Games)))

	if p.Rank() == model.RankCommander && p.Loan() > 0 {
		p.outputf("Outstanding loan: %s IG\n", humanize.Comma(int64(p.Loan())))
	}

	if p.Rank() >= model.RankCommander && p.Rank() <= model.RankAdventurer {
		p.outputf("Trader rating: %s\n", humanize.Comma(int64(p.TradeCredits)))
	}

	if p.Rank() == model.RankGM {
		p.outputf("Tonnage shipped: %s\n", humanize.Comma(int64(p.Shipped)))
	}

	p.outputf("Strength     max: %3d current: %3d\n", p.Str.Max, p.Str.Cur)
	p.outputf("Stamina      max: %3d current: %3d\n", p.Sta.Max, p.Sta.Cur)
	p.outputf("Intelligence max: %3d current: %3d\n", p.Int.Max, p.Int.Cur)
	p.outputf("Dexterity    max: %3d current: %3d\n", p.Dex.Max, p.Dex.Cur)

	if p.HasSpaceship() { //nolint:staticcheck
		shipClass := GetShipClass(p.ShipKit.Tonnage)
		article := "a"
		// TODO
		p.outputf("You own %s %s class spaceship\n", article, shipClass)
	}

	if p.company != nil {
		p.outputf("CEO of %s\n", p.company.Name)
	}

	if p.Rank() >= model.RankSquire && p.Rank() <= model.RankDuke && p.OwnSystem() != nil {
		p.outputf("Overlord of %s\n", p.OwnSystem().Name())
	}

	if p.Rank() >= model.RankDuke {
		if !p.CurSys().IsHidden() && p.Rank() == model.RankDeity {
			var locNo uint32
			if p.IsFlyingSpaceship() {
				locNo = p.ShipLoc
			} else {
				locNo = p.LocNo
			}
			p.outputf("%s system, teleport address %d\n", p.CurSys().Name(), locNo)
		}
	}

	if p.Rank() == model.RankBaron || p.Rank() == model.RankDuke { // nolint:staticcheck
		// TODO
		//  static const char* facilities[7] = {
		//   "Upside",
		//   "Downside",
		//   "Accelerator",
		//   "Processor",
		//   "C/C/C",
		//   "Teleporter",
		//   "Time machine"
		//  };
		//
		//  bool first = true;
		//
		//  for (size_t i = 0; i < DIM_OF(m_facilities); ++i) {
		//   if (m_facilities[i] > 0) {
		//    if (first) {
		//     output("Facilities:\n");
		//     first = false;
		//    }
		//    output(" %-12s %3d%%\n", facilities[i], m_facilities[i]);
		//   }
		//  }
	}
}
