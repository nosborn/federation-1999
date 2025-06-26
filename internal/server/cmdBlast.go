package server

import (
	"strings"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/horsell"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdBlast(what, with model.Name) {
	o, ok := p.FindInventoryName(with)
	if !ok || o.Number() != sol.ObTDX {
		p.Outputm(text.BlastNoExplosive)
		return
	}
	if p.IsInSolSystem() {
		if p.LocNo != sol.TreasureRoom1 {
			p.UnknownCommand()
			return
		}
		if !strings.EqualFold(what.Text, "wall") {
			p.UnknownCommand()
			return
		}
		if p.Dex.Cur < 75 {
			p.Outputm(text.MN337)
			p.Die()
			return
		}
		p.Outputm(text.MN85)
		p.LocNo = sol.CaveIn
		p.setLocation(p.LocNo)
		p.Save(database.SaveNow)
	} else if p.IsInHorsellSystem() {
		if p.CurLocNo() != horsell.Cylinder2 {
			p.UnknownCommand()
			return
		}
		if !strings.EqualFold(what.Text, "nest") {
			p.UnknownCommand()
			return
		}
		// TODO
		// useTDX(thePlayer, theTDX)
		return
	}
	p.RemoveFromInventory(o)
	o.Recycle()
	p.Save(database.SaveNow)
}
