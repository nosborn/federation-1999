package server

import (
	"github.com/dustin/go-humanize"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/text"
)

// Gives the player a table of the game's ranks.
func (p *Player) CmdRanks() {
	p.Outputm(text.RanksHeader)

	for r := range model.RankDuke {
		if r > model.RankGroundHog && ((r-model.RankGroundHog)%3) == 0 {
			p.Output("\n")
		}

		switch p.Sex() {
		case model.SexFemale:
			p.Outputm(text.RanksEntry, r+1, model.FemaleRanks[r])
		case model.SexMale:
			p.Outputm(text.RanksEntry, r+1, model.MaleRanks[r])
		default:
			p.Outputm(text.RanksEntry, r+1, model.NeuterRanks[r])
		}
	}
	p.Output("\n")

	switch p.Rank() {
	case model.RankGroundHog:
		p.Outputm(text.RanksGroundHog)
	case model.RankCommander:
		p.Outputm(text.RanksCommander)
	case model.RankCaptain:
		tcr := int64(getConfig(CFG_ADVENTURER_TCR))
		p.Outputm(text.RanksCaptain, humanize.Comma(tcr))
	case model.RankAdventurer:
		if p.GMLocation() == 0 {
			tcr := int64(getConfig(CFG_TRADER_TCR))
			p.Outputm(text.RanksAdventurer_1, humanize.Comma(tcr))
		} else {
			p.Outputm(text.RanksAdventurer_2)
		}
	case model.RankTrader:
		stats := getConfig(CFG_MERCHANT_STATS)
		p.Outputm(text.RanksTrader, stats)
	case model.RankMerchant:
		profit := getConfig(CFG_JP_PROFIT)
		p.Outputm(text.RanksMerchant, humanize.Comma(int64(profit)))
	case model.RankJP:
		profit := getConfig(GM_PROFIT)
		p.Outputm(text.RanksJP, humanize.Comma(int64(profit)))
	case model.RankGM:
		stat := getConfig(EXPLORER_STAT)
		tons := int64(getConfig(EXPLORER_TONS))
		p.Outputm(text.RanksGM, stat, humanize.Comma(tons))
	case model.RankExplorer:
		p.Outputm(text.RanksExplorer)
	case model.RankSquire:
		if p.IsPromoCharacter() {
			break
		}
		level := text.Msg(text.Level_Mining)
		p.Outputm(text.RanksPO, level)
	case model.RankThane:
		level := text.Msg(text.Level_Industrial)
		p.Outputm(text.RanksPO, level)
	case model.RankIndustrialist:
		level := text.Msg(text.Level_Technological)
		p.Outputm(text.RanksPO, level)
	case model.RankTechnocrat:
		level := text.Msg(text.Level_Leisure)
		p.Outputm(text.RanksPO, level)
	case model.RankBaron:
		if p.IsPromoCharacter() {
			break
		}
		p.Outputm(text.RanksBaron)
	case model.RankDuke:
		// No promotion prospects!
	case model.RankSenator, model.RankMascot, model.RankHostess, model.RankManager, model.RankDeity, model.RankEmperor:
		// No promotion prospects!
	}
}
