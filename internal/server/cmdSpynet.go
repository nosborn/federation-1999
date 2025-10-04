package server

import (
	"github.com/dustin/go-humanize"
	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdSpynetNotice() {
	p.Flags0 ^= model.PL0_INFO
	if (p.Flags0 & model.PL0_INFO) == 0 {
		p.Outputm(text.SpynetNoticeOff)
	} else {
		p.Outputm(text.SpynetNoticeOn)
	}
	p.Save(database.SaveNow)
}

func (p *Player) CmdSpynetReport(name string) {
	if p.Rank() == model.RankGroundHog {
		p.Outputm(text.MN77)
		return
	}

	if p.IsInHorsellSystem() && p.Rank() < model.RankHostess {
		p.Outputm(text.SpynetReportInHorsell)
		return
	}

	debug.Trace("Reporting on subject '%s'", name)

	// Try for a player name first.
	player, ok := FindPlayer(name)
	if ok {
		debug.Trace("Reporting on player")
		if p.HasWallet() {
			cost := int32(p.Rank()+1) * 200
			if p.Balance() < cost {
				p.Outputm(text.SpynetInsufficientFunds)
				return
			}
			p.ChangeBalance(-cost)
			p.Save(database.SaveNow)
		}
		reportStatus(p, player)
		return
	}

	// It wasn't a player name, let's try a company name instead.
	company, ok := FindCompany(name)
	if ok {
		debug.Trace("Reporting on company")
		if p.HasWallet() {
			if p.Balance() < 1200 {
				p.Outputm(text.SpynetInsufficientFunds)
				return
			}
			p.ChangeBalance(-1200)
			p.Save(database.SaveNow)
		}
		p.Outputm(text.SpynetReportCompany, company.Name, company.CEO.Name, humanize.Comma(int64(company.Shares)))
		return
	}

	// No information available.
	p.Outputm(text.MN78)
}

// Produces SPYNET status report for player.
func reportStatus(p *Player, subject *Player) {
	// No SpyNet reports on deities by mere mortals.
	if p.Rank() < model.RankManager && subject.Rank() >= model.RankManager {
		p.Outputm(text.MN74)
		return
	}

	p.Outputm(text.MN603, subject.Name, subject.rankName(), humanize.Comma(int64(subject.Reward())))

	if subject.HasSpaceship() {
		p.Outputm(text.MN607, GetShipClass(subject.ShipKit.Tonnage))
	}

	if subject.company != nil {
		p.Outputm(text.MN608, subject.company.Name)
	}

	if subject.Rank() >= model.RankSquire && subject.Rank() <= model.RankDuke {
		debug.Check(subject.ownSystem != nil)
		p.Outputm(text.MN609, subject.ownSystem.Name())
	}

	if subject.IsPlaying() {
		if subject.curSys.IsHidden() && p.Rank() < model.RankHostess {
			p.Output("   Current whereabouts unknown\n")
		} else {
			if p.Rank() >= model.RankDuke {
				p.outputf("   Currently in %s system location %d\n", subject.curSys.Name(), subject.LocNo)
			} else {
				p.outputf("   Currently in %s system\n", subject.curSys.Name())
			}
		}
	}

	// Everything from here is extra info about mortals for deities.
	if p.Rank() < model.RankManager || subject.Rank() >= model.RankDeity {
		return
	}

	p.outputf("   Played %d games\n", subject.Games)
	p.outputf("   Personal balance: %s IG\n", humanize.Comma(int64(subject.Balance())))

	if subject.company != nil {
		p.outputf("   Company balance: %s IG\n", humanize.Comma(int64(subject.company.Balance)))
	}

	if subject.ownSystem != nil {
		p.outputf("   Treasury balance: %s IG\n", humanize.Comma(int64(subject.ownSystem.Balance())))
	}

	// Puzzle status.
	if subject.Rank() == model.RankAdventurer && subject.GMLocation() != 0 {
		p.outputf("   Grand Master is at location %d.\n", subject.GMLocation())
	}

	if subject.HasIDCard() {
		p.Output("   Holds a commission in Naval Intelligence.\n")
	} else if (subject.Flags1 & model.PL1_MI6_OFFERED) != 0 {
		p.Output("   Commission in Naval Intelligence offered.\n")
	}

	if subject.hasSnarkAssignment() {
		p.Output("   Snark puzzle has been assigned.\n")
	}

	if subject.hasCompletedSnarkPuzzle() {
		p.Output("   Has completed Snark puzzle.\n")
	}

	if subject.hasHorsellAssignment() {
		p.Output("   Horsell puzzle has been assigned.\n")
	}

	if (subject.Flags0 & model.PL0_KNOWS_ORSONITE) != 0 {
		p.Output("   Knows about Orsonite.\n")
	}

	// DataSpace Navigator status.
	if subject.IsNavigator() {
		if subject.IsOnDutyNavigator() {
			p.Output("   DataSpace Navigator - on duty\n")
		} else {
			p.Output("   DataSpace Navigator - off duty\n")
		}
	}

	// Lockout status.
	if subject.IsLockedOut() {
		p.Output("   Player is locked out.\n")
	}

	// Gag status.
	if !subject.HasCommUnit() {
		p.Output("   Player is gagged.\n")
	}

	// Promotional character status.
	if subject.IsPromoCharacter() {
		p.Output("   Promotional character.\n")
	}
}
