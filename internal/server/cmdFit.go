package server

import "github.com/nosborn/federation-1999/internal/model"

// Catch-all verb for manipulation objects...
func (p *Player) CmdFit(name model.Name) {
	p.Output("Not implemented. Check back in 2 weeks.\n")
}
