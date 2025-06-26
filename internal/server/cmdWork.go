package server

import (
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdWork(duchyName string) {
	var d *Duchy
	if duchyName == "" {
		d = p.CurSys().Duchy()
	} else {
		var ok bool
		d, ok = FindDuchy(duchyName)
		if !ok || d.IsHidden() {
			p.Outputm(text.NoSuchDuchy)
			return
		}
	}
	d.ListWork(p)
}
