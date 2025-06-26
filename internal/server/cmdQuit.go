package server

import (
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdQuit() {
	p.Outputm(text.Quit)
	p.Session().Quit()
}
