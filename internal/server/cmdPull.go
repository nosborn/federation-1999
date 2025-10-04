package server

import (
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

// Player attempts to pull a lever.
func (p *Player) CmdPullLever() {
	if p.IsInSolSystem() {
		if p.CurLoc().Number() == sol.Chamber {
			// increaseIntelligence(*this) -- FIXME
			return
		}
		if p.CurLoc().Number() == 711 { // FIXME: where is this?
			p.Outputm(text.MN565)
			return
		}
	}
	p.Outputm(text.MN55)
}
