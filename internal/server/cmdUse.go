package server

import "github.com/nosborn/federation-1999/internal/model"

// Catchall verb for manipulation objects...
func (p *Player) CmdUse(objectName model.Name) {
	p.Output("Not implemented. Check back in 2 weeks.\n")
}
