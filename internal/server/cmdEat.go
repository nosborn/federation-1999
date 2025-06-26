package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

// CmeEat allows player to eat an edible object carried by the player.
func (p *Player) CmdEat(name model.Name) {
	o, ok := p.FindInventoryName(name)
	if !ok {
		p.Outputm(text.EatNotCarried)
		return
	}
	// TODO:
	// if o.homeSystem == p.CurSys() && o.consumeEvent != 0 {
	// 	// if !XEventHandler(p, o.consumeEvent) {
	// 	// 	return
	// 	// }
	// }
	if !o.IsEdible() {
		p.Outputm(text.EatNotEdible)
		return
	}
	p.RemoveFromInventory(o)
	o.Recycle()
	p.SetCurSta(min(p.CurSta()+5, p.MaxSta()))
	p.Outputm(text.EatOK)
	p.Save(database.SaveNow)
}
