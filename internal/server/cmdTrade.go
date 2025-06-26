package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdTrade() {
	noTrade := true

	if p.CurLoc().IsShipyard() {
		p.Outputm(text.TradeInShipyard)
		noTrade = false
	}
	if p.CurLoc().IsGeneralStore() {
		p.Outputm(text.TradeInGeneralStore)
		noTrade = false
	}
	if p.CurLoc().IsWeaponsShop() {
		p.Outputm(text.TradeInWeaponsShop)
		noTrade = false
	}
	if p.CurLoc().IsCafe() {
		p.Outputm(text.TradeInCafe)
		noTrade = false
	}
	if p.CurLoc().IsRepairShop() {
		p.Outputm(text.TradeInRepairShop)
		noTrade = false
	}
	if p.CurLoc().IsElectronicsStore() {
		p.Outputm(text.TradeInElectronicsStore)
		noTrade = false
	}
	if p.CurLoc().IsClothingStore() {
		p.Outputm(text.TradeInClothingStore)
		noTrade = false
	}

	switch {
	case p.IsInSolLocation(sol.PamperUHealthFarm):
		p.Outputm(text.TradeBuyDexterity)
		noTrade = false
	case p.IsInSolLocation(sol.InstaLernFactORamaUniversity):
		p.Outputm(text.TradeBuyIntelligence)
		noTrade = false
	case p.IsInSolLocation(sol.BattleCreekSanitarium):
		p.Outputm(text.TradeBuyStamina)
		noTrade = false
	case p.IsInSolLocation(sol.FeeldaBurnesAerobicsAndWorkoutClasses):
		p.Outputm(text.TradeBuyStrength)
		noTrade = false
	}

	if p.CurLoc().IsExchange() {
		if p.Rank() < model.RankMerchant {
			p.Outputm(text.MN477)
		} else {
			p.Outputm(text.TradeInExchange)
		}
		noTrade = false
	}

	if noTrade {
		p.Outputm(text.TradeNothingAvailable)
	}
}

// Allows Merchants to set the commodity group they wish to watch.
func (p *Player) CmdTradeGroup(group model.CommodityGroup) {
	if !p.CurLoc().IsExchange() {
		p.Outputm(text.MN157)
		return
	}
	if p.Rank() < model.RankMerchant {
		p.Outputm(text.MN477)
		return
	}
	p.SetTradeGroup(group)
	p.Outputm(text.MN482)
}
