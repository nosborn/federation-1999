package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

const (
	buyAleCost        = 5       //         5 IG
	AMMUNITION_COST   = 1000    //     1,000 IG
	buyClothesCost    = 750     //       750 IG
	DEXTERITY_COST    = 1500000 // 1,500,000 IG
	FOOD_COST         = 10      //        10 IG
	FUEL_COST         = 10      //        10 IG
	INTELLIGENCE_COST = 3000000 // 3,000,000 IG
	LAMP_COST         = 10      //        10 IG
	MISSILE_COST      = 25000   //    25,000 IG
	STAMINA_COST      = 1500000 // 1,500,000 IG
	STRENGTH_COST     = 1500000 // 1,500,000 IG
)

const (
	SPYBEAM_WEIGHT = 50 //         50 tons
)

// CmdBuyAle allows player to buy a drink in a cafe location and consume it
// then and there.
func (p *Player) CmdBuyAle() {
	if !p.curLoc.IsCafe() {
		p.Outputm(text.BuyAleWrongLocation)
		return
	}
	if p.CurSys().IsClosed() {
		p.Outputm(text.ClosedForBusiness, p.CurSys().Name())
		return
	}
	if p.HasWallet() {
		if p.Balance() < buyAleCost {
			p.Outputm(text.MN757)
			return
		}
		p.ChangeBalance(-buyAleCost)
	}
	p.Sta.Cur = min(p.Sta.Cur+2, p.Sta.Max)
	p.Outputm(text.BuyAleOK)
	p.Save(database.SaveNow)
}

// CmdBuyAmmunition allows a player to buy mag gun ammo from a weapon shop.
func (p *Player) CmdBuyAmmunition(number int32) {
	if !p.curLoc.IsWeaponsShop() {
		p.Outputm(text.MN676)
		return
	}
	if !p.HasSpaceship() {
		p.Outputm(text.BUY_AMMO_NO_SPACESHIP)
		return
	}

	if number < 1 {
		p.Outputm(text.MN243)
		return
	}
	if number > 100 {
		number = 100
	}
	_ = number // FIXME

	// TODO:
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

// CmdBuyClothes allows users to write their own description.
func (p *Player) CmdBuyClothes(clothes string) {
	if !p.curLoc.IsClothingStore() {
		p.Outputm(text.BuyClothesWrongLocation)
		return
	}
	if p.CurSys().IsClosed() {
		p.Outputm(text.ClosedForBusiness, p.CurSys().Name())
		return
	}
	if p.HasWallet() && p.Balance() < buyClothesCost {
		p.Outputm(text.InsufficientFunds)
		return
	}
	if len(clothes) >= model.MAX_PER_DESC {
		p.Outputm(text.DescriptionTooLong, model.MAX_PER_DESC-1)
		return
	}
	if clothes[0] == '/' || clothes[0] == '>' {
		p.Outputm(text.DescriptionBadLeader)
		return
	}
	p.Desc = clothes
	if p.HasWallet() {
		p.ChangeBalance(-buyClothesCost)
		p.CurSys().Income(buyClothesCost, true)
	}
	p.Outputm(text.BuyClothesOK, p.Desc)
	p.Save(database.SaveNow)
}

// CmdBuyCompany purchases a company for the player.
func (p *Player) CmdBuyCompany(name string) {
	if p.rank < model.RankMerchant || p.rank > model.RankExplorer {
		p.Outputm(text.MN1028)
		return
	}

	// TODO:
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

func (p *Player) CmdBuyDexterity(points int32) {
	if !p.IsInSolLocation(sol.PamperUHealthFarm) {
		p.Outputm(text.BuyDexterityWrongLocation)
		return
	}
	if (p.Flags1 & model.PL1_DONE_DEX) == 0 {
		p.Outputm(text.DO_DEX_PUZZLE)
		return
	}
	points = min(max(1, points), 120) // 1 <= points <= 120
	cost := DEXTERITY_COST * points
	if p.Balance() < cost {
		p.Outputm(text.InsufficientFunds)
		return
	}
	p.Dex.Max = min(p.Dex.Max+points, 120)
	p.Dex.Cur = min(p.Dex.Cur+points, p.Dex.Max)
	p.ChangeBalance(-cost)
	p.Outputm(text.BuyDexterityOK)
	p.Save(database.SaveNow)
	p.CheckForPromotion()
}

// CmdBuyFactory allows a player to buy a factory on the specified planet, if
// the planet owner agrees...
func (p *Player) CmdBuyFactory(where string, commodity model.Commodity) {
	// Player must own a company in order to buy a factory.
	if p.company == nil {
		p.Outputm(text.MN1035)
		return
	}

	// TODO:
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

// CmdBuyFood allows a player in a cafe location to buy & eat a meal.
func (p *Player) CmdBuyFood() {
	if !p.curLoc.IsCafe() {
		p.Outputm(text.BuyFoodWrongLocation)
		return
	}
	if p.CurSys().IsClosed() {
		p.Outputm(text.ClosedForBusiness, p.CurSys().Name())
		return
	}
	if p.HasWallet() {
		if p.Balance() < FOOD_COST {
			p.Outputm(text.MN757)
			return
		}
		p.ChangeBalance(-FOOD_COST)
	}
	p.Sta.Cur = min(p.Sta.Cur+5, p.Sta.Max)
	p.Outputm(text.BuyFoodOK)
	p.Save(database.SaveNow)
}

func (p *Player) CmdBuyFuel(quantity int32) {
	if !p.HasSpaceship() {
		p.Outputm(text.NoSpaceship)
		return
	}
	if quantity == -1 { // Top up the tank!
		quantity = p.ShipKit.MaxFuel - p.ShipKit.CurFuel
	} else if quantity > p.ShipKit.MaxFuel {
		quantity = p.ShipKit.MaxFuel
	}
	if p.HasWallet() {
		cost := FUEL_COST * quantity
		if p.curLoc.IsSpace() {
			cost *= 5
		}
		if p.Balance() < cost {
			p.Outputm(text.InsufficientFunds)
			return
		}
		p.ChangeBalance(-cost)
	}
	p.ShipKit.CurFuel = min(p.ShipKit.CurFuel+quantity, p.ShipKit.MaxFuel)
	if p.curLoc.IsSpace() {
		p.Outputm(text.MN34)
	}
	p.Outputm(text.MN492, quantity)
	p.quickStatus()
	p.Save(database.SaveNow)
}

func (p *Player) CmdBuyGoods(commodity model.Commodity, planetName string, quantity int32) {
	if p.rank < model.RankTrader {
		p.Outputm(text.MN453)
		return
	}
	if !p.curLoc.IsExchange() {
		p.Outputm(text.MN157)
		return
	}
	if !p.HasSpaceship() {
		p.Outputm(text.NoSpaceship)
		return
	}

	// TODO:
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

func (p *Player) CmdBuyIntelligence(points int32) {
	if !p.IsInSolLocation(sol.InstaLernFactORamaUniversity) {
		p.Outputm(text.BuyIntelligenceWrongLocation)
		return
	}
	if (p.Flags1 & model.PL1_DONE_INT) == 0 {
		p.Outputm(text.DO_INT_PUZZLE)
		return
	}
	points = min(max(1, points), 120) // 1 <= points <= 120
	cost := INTELLIGENCE_COST * points
	if p.Balance() < cost {
		p.Outputm(text.InsufficientFunds)
		return
	}
	p.Int.Max = min(p.Int.Max+points, 120)
	p.Int.Cur = min(p.Int.Cur+points, p.Int.Max)
	p.ChangeBalance(-cost)
	p.Outputm(text.BuyIntelligenceOK)
	p.Save(database.SaveNow)
	p.CheckForPromotion()
}

func (p *Player) CmdBuyLamp() {
	if !p.curLoc.IsGeneralStore() {
		p.Outputm(text.MN493)
		return
	}
	if p.CurSys().IsClosed() {
		p.Outputm(text.ClosedForBusiness, p.CurSys().Name())
		return
	}
	if p.HasWallet() {
		if p.Balance() < LAMP_COST {
			p.Outputm(text.InsufficientFunds)
			return
		}
		p.ChangeBalance(-LAMP_COST)
	}
	p.Flags0 |= model.PL0_LIT
	p.Outputm(text.MN494)
	p.Save(database.SaveNow)
}

// CmdBuyMissiles allows a player to buy missiles to arm her/his ship with.
func (p *Player) CmdBuyMissiles(number int32) {
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

// CmdBuyPizza buys a round of pizza for other players in the cafe/bar.
func (p *Player) CmdBuyPizza() {
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

// CmdBuyRound buys a round of drinks for other players in the same room.
func (p *Player) CmdBuyRound(description string) {
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

// CmdBuySpyBeam allows player to buy equipment suitable for spying on other
// players.
func (p *Player) CmdBuySpyBeam() {
	if !p.curLoc.IsElectronicsStore() {
		p.Outputm(text.MN582)
		return
	}
	if p.HasSpyBeam() {
		p.Output("You already have a SpyBeam!\n")
		return
	}
	if p.Balance() < model.SPYBEAM_COST || p.Rank() == model.RankSenator {
		p.Outputm(text.InsufficientFunds)
		return
	}
	if !p.HasSpaceship() {
		p.Outputm(text.NoSpaceship)
		return
	}
	if p.ShipKit.MaxHold != p.ShipKit.CurHold {
		p.Outputm(text.MN69)
		return
	}
	if p.ShipKit.MaxHold <= SPYBEAM_WEIGHT {
		p.Outputm(text.MN584)
		return
	}
	if p.Rank() > model.RankDuke {
		p.Output("Now why would you want to do that?\n")
		return
	}
	p.ChangeBalance(-model.SPYBEAM_COST)
	p.Flags0 |= model.PL0_SPYBEAM
	p.ShipKit.MaxHold -= SPYBEAM_WEIGHT
	p.ShipKit.CurHold -= SPYBEAM_WEIGHT
	p.Outputm(text.MN70)
	p.Save(database.SaveNow)
}

// CmdBuySpyScreen allows the player - for a -*vast*- amount of groats - to buy
// a screen which protects from own level spying.
func (p *Player) CmdBuySpyScreen() {
	if !p.curLoc.IsElectronicsStore() {
		p.Outputm(text.MN580)
		return
	}
	if p.HasSpyScreen() {
		p.Output("You already have a SpyBeam Screen!\n")
		return
	}
	if p.Balance() < model.SPYSCREEN_COST {
		p.Outputm(text.MN581)
		p.Outputm(text.InsufficientFunds)
		return
	}
	if !p.HasSpaceship() {
		p.Outputm(text.NoSpaceship)
		return
	}
	if p.Rank() > model.RankDuke {
		p.Output("Now why would you want to do that?")
		return
	}
	p.ChangeBalance(-model.SPYSCREEN_COST)
	p.CurSys().Income(model.SPYSCREEN_COST, true)
	p.Flags0 |= model.PL0_SPYSCREEN
	p.Outputm(text.MN68)
	p.Save(database.SaveNow)
	p.CheckSpyers()
}

func (p *Player) CmdBuyStamina(points int32) {
	if !p.IsInSolLocation(sol.BattleCreekSanitarium) {
		p.Outputm(text.BuyStaminaWrongLocation)
		return
	}
	if (p.Flags1 & model.PL1_DONE_STA) == 0 {
		p.Outputm(text.DO_STA_PUZZLE)
		return
	}
	points = min(max(1, points), 120) // 1 <= points <= 120
	cost := STAMINA_COST * points
	if p.Balance() < cost {
		p.Outputm(text.InsufficientFunds)
		return
	}
	p.Sta.Max = min(p.Sta.Max+points, 120)
	p.Sta.Cur = min(p.Sta.Cur+points, p.Sta.Max)
	p.ChangeBalance(-cost)
	p.Outputm(text.BuyStaminaOK)
	p.Save(database.SaveNow)
	p.CheckForPromotion()
}

func (p *Player) CmdBuyStrength(points int32) {
	if !p.IsInSolLocation(sol.FeeldaBurnesAerobicsAndWorkoutClasses) {
		p.Outputm(text.BuyStrengthWrongLocation)
		return
	}
	if (p.Flags1 & model.PL1_DONE_STR) == 0 {
		p.Outputm(text.DO_STR_PUZZLE)
		return
	}
	points = min(max(1, points), 120) // 1 <= points <= 120
	cost := STRENGTH_COST * points
	if p.Balance() < cost {
		p.Outputm(text.InsufficientFunds)
		return
	}
	p.Str.Max = min(p.Str.Max+points, 120)
	p.Str.Cur = min(p.Str.Cur+points, p.Str.Max)
	p.ChangeBalance(-cost)
	p.Outputm(text.BuyStrengthOK)
	p.Save(database.SaveNow)
	p.CheckForPromotion()
}

// CmdBuyWarehouse gives the player a warehouse linked to the exchange in which
// the warehouse is bought.
func (p *Player) CmdBuyWarehouse() {
	// Are we in a Trading exchange?
	if !p.curLoc.IsExchange() {
		p.Outputm(text.MN734)
		return
	}

	// Planet owners can't buy warehouses.
	if p.rank == model.RankExplorer && p.OwnSystem() != nil {
		p.Outputm(text.NiceTry)
		return
	}
	if p.rank > model.RankExplorer {
		p.Outputm(text.PO_WAREHOUSE_PURCHASE)
		return
	}

	p.Output("Not implemented. Check back in 2 weeks.\n")
}
