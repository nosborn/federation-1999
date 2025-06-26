package server

import (
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

// Clears flag indicating that the player only requires brief descriptions.
func (p *Player) CmdFull() {
	if p.WantsBrief() {
		p.SetBrief(false)
		p.Save(database.SaveNow)
	}
	p.Outputm(text.FullOK)
}
