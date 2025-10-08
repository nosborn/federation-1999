package server

import (
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdPost(message string) {
	if p.IsInHorsellSystem() {
		p.Outputm(text.FakeNoBarboardHere)
		return
	}
	if !p.CurLoc().IsCafe() {
		p.Outputm(text.NoBarboardHere)
		return
	}
	p.Output("Not implemented. Check back in 2 weeks.\n")
}
