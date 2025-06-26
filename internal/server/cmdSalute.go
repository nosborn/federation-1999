package server

import (
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdSalute() {
	// FIXME: Should be checking for SALUTE KATOV.

	if !p.IsInSolLocation(sol.ControlRoom3) || p.HasIDCard() {
		p.Outputm(text.Salute)
		return
	}
	p.SetMI6Offered(false)
	p.SetIDCard(true)
	p.Outputm(text.SaluteKatov)
	p.Save(database.SaveNow)
}
