package server

import (
	"os"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/exchange"
	"github.com/nosborn/federation-1999/internal/server/global"
	"github.com/nosborn/federation-1999/internal/server/goods"
	"github.com/nosborn/federation-1999/internal/text"
	"github.com/nosborn/federation-1999/internal/version"
)

func (p *Player) CmdShowConfiguration() {
	if !p.canShow(model.RankHostess) {
		p.UnknownCommand()
		return
	}

	p.showHeader()

	if global.TestFeaturesEnabled {
		p.Nsoutputm(text.ShowConfig_TestEnabled)
	} else {
		p.Nsoutputm(text.ShowConfig_TestDisabled)
	}

	p.Nsoutput("\n")
	p.Nsoutputf("CFG_ADVENTURER_TCR: %s\n", humanize.Comma(int64(getConfig(CFG_ADVENTURER_TCR))))
	p.Nsoutputf("    CFG_TRADER_TCR: %s\n", humanize.Comma(int64(getConfig(CFG_TRADER_TCR))))
	p.Nsoutputf("CFG_MERCHANT_STATS: %s\n", humanize.Comma(int64(getConfig(CFG_MERCHANT_STATS))))
	p.Nsoutputf("     CFG_JP_PROFIT: %s\n", humanize.Comma(int64(getConfig(CFG_JP_PROFIT))))
	p.Nsoutputf("     CFG_GM_PROFIT: %s\n", humanize.Comma(int64(getConfig(CFG_GM_PROFIT))))
	p.Nsoutputf(" CFG_EXPLORER_STAT: %s\n", humanize.Comma(int64(getConfig(CFG_EXPLORER_STAT))))
	p.Nsoutputf(" CFG_EXPLORER_TONS: %s\n", humanize.Comma(int64(getConfig(CFG_EXPLORER_TONS))))
	p.Nsoutputf(" CFG_RP_JOB_CREDIT: %s\n", humanize.Comma(int64(getConfig(CFG_RP_JOB_CREDITS))))
	p.Nsoutputf("  CFG_PLANET_PROMO: %s\n", humanize.Comma(int64(getConfig(CFG_PLANET_PROMO))))
}

func (p *Player) CmdShowProduction() {
	if p.Rank() != model.RankDeity {
		p.UnknownCommand()
		return
	}
	for i := range goods.GoodsArray {
		p.Nsoutputf("%-15s %13s %13s\n",
			goods.GoodsArray[i].Name,
			humanize.Comma(int64(exchange.Production[i])),
			humanize.Comma(int64(FactoryProduction[i])))
	}
}

func (p *Player) CmdShowStatus() {
	if !p.canShow(model.RankHostess) {
		p.UnknownCommand()
		return
	}

	p.showHeader()

	ftime := global.GameStartTime.Format("Mon Jan _2 15:04:05 2006 MDT")
	p.Nsoutputm(text.ShowStatus_Started, ftime)

	uptime := time.Since(global.GameStartTime).Minutes()
	p.Nsoutputm(text.ShowStatus_UpTime, text.HumanizeMinutes(int64(uptime)))

	p.Nsoutputm(text.ShowStatus_Players, humanize.Comma(int64(global.PeakPlayers)))

	loadQueueLength := LoaderQueueLength()
	if loadQueueLength == 0 {
		p.Nsoutputm(text.ShowStatus_LoadQueue_Empty)
	} else {
		plural := "s"
		if loadQueueLength == 1 {
			plural = ""
		}
		if loader.IsFrozen() {
			p.Nsoutputm(text.ShowStatus_LoadQueue_Frozen, humanize.Comma(int64(loadQueueLength)), plural)
		} else {
			p.Nsoutputm(text.ShowStatus_LoadQueue, humanize.Comma(int64(loadQueueLength)), plural)
		}
	}

	// TODO:
	// if GetCraterLife() > 0 {
	// 	p.Nsoutputm(text.ShowStatus_Crater, GetCraterLife())
	// }

	// TODO:
	// if GetLookingGlassLife() > 0 {
	// 	p.Nsoutputm(text.ShowStatus_MirrorRoom, GetLookingGlassLife())
	// }
}

func (p *Player) canShow(minRank model.Rank) bool {
	if p.Rank() >= minRank {
		return true
	}
	if global.TestFeaturesEnabled {
		p.Nsoutputm(text.TestFeature)
		p.FlushOutput()
		return true
	}
	return false
}

func (p *Player) showHeader() {
	hostname, _ := os.Hostname()
	ftime := time.Now().Local().Format("Mon Jan _2 15:04:05 2006 MDT")
	p.Nsoutputm(text.Show_Header, ftime, version.String(), hostname)
}
