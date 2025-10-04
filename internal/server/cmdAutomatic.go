package server

import (
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

// CmdAutomatic sets the battle computer to fight on automatic.
func (p *Player) CmdAutomatic() {
	if !p.HasSpaceship() {
		p.Outputm(text.NoSpaceship)
		return
	}
	if !p.IsAuto() {
		p.SetAuto(true)
		p.Save(database.SaveNow)
	}
	p.Outputm(text.AutomaticOK)
	p.quickStatus()
}
