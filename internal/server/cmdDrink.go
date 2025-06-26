package server

import (
	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

// Player drinks something s/he is carried.
func (p *Player) CmdDrink(name model.Name) {
	object, ok := p.FindInventoryName(name)
	if !ok {
		p.Outputm(text.DrinkNotCarried)
		return
	}

	if p.curSys.DrinkEvent(p, object) {
		return
	}

	if p.CurSys() == object.HomeSystem() && object.ConsumeEvent() != 0 { //nolint:staticcheck // SA9003: empty branch
		// TODO:
		// if !XEventHandler(p, o.consumeEvent) {
		// 	return
		// }
	}

	if !object.IsLiquid() {
		p.Outputm(text.DrinkNotLiquid, object.DisplayName(false))
		return
	}

	p.RemoveFromInventory(object)
	object.Recycle()

	p.Sta.Cur = min(p.Sta.Cur+2, p.Sta.Max)

	p.Outputm(text.DrinkOK)
	p.Save(database.SaveNow)
}

func (p *Player) DrinkPotion(o *Object) {
	debug.Precondition(o != nil)

	if !p.CurSys().IsSnark() {
		p.Outputm(text.DrinkPotionNoEffect)
		return
	}

	p.RemoveFromInventory(o)
	o.Recycle()

	p.Sta.Max = min(p.Sta.Max+5, 120)
	p.Sta.Cur = min(p.Sta.Cur+5, p.Sta.Max)
	p.Str.Max = min(p.Str.Max+5, 120)
	p.Str.Cur = min(p.Str.Cur+5, p.Str.Max)
	p.Dex.Max = max(p.Dex.Max-5, 1)
	p.Dex.Cur = min(p.Dex.Cur, p.Dex.Max)
	p.Int.Max = max(p.Int.Max-5, 1)
	p.Int.Cur = min(p.Int.Cur, p.Int.Max)

	p.Outputm(text.DrinkPotionOK)
	p.Save(database.SaveNow)
}

func (p *Player) DrinkWHOOSH(o *Object) {
	// Part 1: Drink the WHOOSH.
	if o != nil {
		debug.Precondition(o.Number() == sol.ObWHOOSH)

		if p.IsInsideSpaceship() {
			p.Outputm(text.DrinkWHOOSHNoEffect)
			return
		}

		p.RemoveFromInventory(o)
		o.Recycle()

		if !p.IsInSolSystem() {
			p.flags2 |= PL2_WHOOSH
			p.Outputm(text.DrinkWHOOSHDelayed)
			return
		}

		p.Output("DrinkWHOOSH")
	}

	// Part 2: Run to the loo.
	debug.Check(p.IsInSolSystem())

	if !p.IsInSolSystem() { // Shouldn't happen
		return
	}

	if o == nil {
		debug.Check((p.flags2 & PL2_WHOOSH) != 0)
		p.Outputm(text.DrinkWHOOSHDelayedReaction)
		p.flags2 &^= PL2_WHOOSH
	}

	if len(p.inventory) > 0 {
		for i := range p.inventory {
			p.inventory[i].curLocNo = p.LocNo
		}
		// p.curLoc.Talk(text.Msg("DropOK", text.ListOfObjects(p.inventory)), p)
		p.inventory = nil
	}

	p.CurSys().CleanupHook(p)
	p.curLoc.Talk(text.Msg(text.PlayerHasLeft, p.Name), p)

	var locNo uint32
	switch p.Sex() {
	case model.SexFemale:
		locNo = sol.LadiesLoo
	case model.SexMale:
		locNo = sol.GentsLoo
	default:
		locNo = sol.Loo
	}
	p.setLocation(locNo)

	if p.HasSpaceship() {
		p.ShipLoc = sol.EarthLandingArea
	}

	msg := text.Msg(text.PlayerHasArrived, p.name)
	p.curLoc.Talk(msg, p)

	p.Save(database.SaveNow)
}
