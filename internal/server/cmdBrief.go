package server

import (
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

// CmdBrief sets flag to indicate that the player only requires brief
// descriptions.
func (p *Player) CmdBrief() {
	if !p.WantsBrief() {
		p.SetBrief(true)
		p.Save(database.SaveNow)
	}
	p.Outputm(text.BriefOK)
}
