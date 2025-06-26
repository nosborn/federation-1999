package server

import (
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

const (
	cureCost = 25000 // 25,000 IG
)

func (p *Player) CmdCure() {
	if !p.curLoc.IsHospital() {
		p.Outputm(text.CureNotHere)
		return
	}
	if p.CurSys().IsClosed() {
		p.Outputm(text.ClosedForBusiness, p.CurSys().Name())
		return
	}
	if p.Sta.Cur >= p.Sta.Max && p.Str.Cur >= p.Str.Max && p.Dex.Cur >= p.Dex.Max && p.Int.Cur >= p.Int.Max {
		p.Outputm(text.CureNotNeeded)
		return
	}
	if p.HasWallet() && p.Balance() < cureCost {
		p.Outputm(text.TooExpensive, cureCost)
		return
	}
	// Hmmm... this will nuke the 5 point dex boost that Diesel can give out.
	// Do we actually care?
	p.Sta.Cur = p.Sta.Max
	p.Str.Cur = p.Str.Max
	p.Dex.Cur = p.Dex.Max
	p.Int.Cur = p.Int.Max
	if p.HasWallet() {
		p.ChangeBalance(-cureCost)
		p.CurSys().Income(cureCost, true)
	}
	p.Outputm(text.CureOK)
	p.Save(database.SaveNow)
}
