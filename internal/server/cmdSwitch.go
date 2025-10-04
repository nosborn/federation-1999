package server

import (
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdSwitchShields() {
	if !p.HasSpaceship() {
		p.Outputm(text.NoSpaceship)
		return
	}
	if p.MaxShield() == 0 {
		p.Outputm(text.MN715)
		return
	}
	p.ToggleShields()
	p.Outputm(text.MN716)
	p.quickStatus()
	p.Save(database.SaveNow)
}
