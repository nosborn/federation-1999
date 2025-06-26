package server

import (
	"github.com/nosborn/federation-1999/internal/text"
)

// Gives the player a one line status display.
func (p *Player) CmdQuickStatus() {
	if !p.HasSpaceship() {
		p.Outputm(text.NoSpaceship)
		return
	}
	p.quickStatus()
}
