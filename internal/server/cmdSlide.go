package server

import (
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

// CmdSlide allows a player to slide an object aside - like shelving for
// instance!
func (p *Player) CmdSlide() {
	if p.CurSys().CmdSlideHook(p) {
		return
	}
	p.Outputm(text.MN1011)
}

// CmdSlideShelving allows a player to slide an object aside - like shelving
// for instance!
func (p *Player) CmdSlideShelving() {
	if !p.IsInSolLocation(sol.Storeroom3) {
		p.Outputm(text.MN1011)
		return
	}
	if p.CurDex() < 75 {
		p.Outputm(text.MN208)
		return
	}
	p.LocNo = sol.BareRoom3
	p.setLocation(p.LocNo)
	p.Save(database.SaveNow)
	p.Outputm(text.MN209)
}
