package server

import (
	"github.com/nosborn/federation-1999/internal/text"
)

// Allows a player to read the messages on the bar's noticeboard.
func (p *Player) CmdRead() {
	if p.IsInHorsellSystem() {
		p.Outputm(text.FakeNoBarboardHere)
		return
	}
	if !p.curLoc.IsCafe() {
		p.Outputm(text.NoBarboardHere)
		return
	}
	p.Output(ReadBarBoard(p, false)) // FIXME
}
