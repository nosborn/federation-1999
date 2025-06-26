package server

import (
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

// Sets the battle computer to fight on manual.
func (p *Player) CmdManual() {
	if !p.HasSpaceship() {
		p.Outputm(text.NoSpaceship)
		return
	}
	if p.IsAuto() {
		p.SetAuto(false)
		p.Save(database.SaveNow)
	}
	p.Outputm(text.ManualOK)
}
