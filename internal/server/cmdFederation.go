package server

import (
	"github.com/nosborn/federation-1999/internal/text"
	"github.com/nosborn/federation-1999/internal/version"
)

func (p *Player) CmdFederation() {
	p.Outputm(text.Federation, version.String())
}
