package server

import (
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdHamsters() {
	p.Outputm(text.Hamsters)
}
