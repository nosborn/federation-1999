package server

import (
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

const (
	bangCost = 15 // 15 IG
)

// CmdBang lets the player maka a noise by banging on something.
func (p *Player) CmdBang() {
	if !p.IsInSolLocation(sol.Cell) {
		p.Outputm(text.BangWrongLocation)
		return
	}
	if p.Balance() > 0 {
		p.ChangeBalance(-bangCost)
		if p.Balance() < 0 {
			p.SetBalance(0)
		}
	}
	XEventHandler(p, 9)
	p.Save(database.SaveNow)
}
