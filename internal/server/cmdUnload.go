package server

import (
	"github.com/nosborn/federation-1999/internal/text"
)

// Unload cargo from a player's ship and put it in the appropriate place.
func (p *Player) CmdUnload() {
	if !p.HasSpaceship() {
		p.Outputm(text.NoSpaceship)
		return
	}

	p.Output("Not implemented. Check back in 2 weeks.\n")
}
