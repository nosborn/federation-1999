package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/text"
)

// Fire a mag gun or laser at an opponent...
func (p *Player) CmdFire(weapon int, name *model.Name) {
	if !p.IsFlyingSpaceship() {
		p.Outputm(text.MN248)
		return
	}

	// TODO:
	p.Output("Not implemented. Check back in 2 weeks.\n")
}
