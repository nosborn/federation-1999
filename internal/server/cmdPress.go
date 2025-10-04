package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdPress() {
	p.Outputm(text.PressWhat)
}

// Allows a player to push a button.
func (p *Player) CmdPressButton0() {
	if !p.IsInSolSystem() {
		p.Outputm(text.PressWhat)
		return
	}
	if p.CurLocNo() == sol.BareRoom3 || p.CurLocNo() == sol.BareRoom4 {
		p.Outputm(text.PressWhichButton)
		return
	}
	p.Outputm(text.MN58)
}

// Allows a player to push a button.
func (p *Player) CmdPressButton1(colour model.ButtonColour) {
	if !p.IsInSolSystem() {
		p.Outputm(text.PressWhat)
		return
	}
	p.Output("Not implemented. Check back in 2 weeks.\n")
}
