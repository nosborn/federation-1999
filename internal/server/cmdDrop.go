package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/text"
)

// CmdDrop allows a player to drop an object being carried. Checks for any
// applicable events.
func (p *Player) CmdDrop(objectName model.Name) {
	// See if player is carrying it!
	o, ok := p.FindInventoryName(objectName)
	if !ok {
		p.Outputm(text.MN514)
		return
	}
	// Deal with any drop events. If the object doesn't belong in the current
	// star system, events are skipped; this mostly avoids dealing with a Baron
	// who drops the TDX in Horsell.
	if p.CurSys() == o.HomeSystem() {
		if o.DropEvent() != 0 {
			if !XEventHandler(p, o.DropEvent()) {
				return
			}
		}
		if !EventHandler(o.Events()[2], o, p) {
			return
		}
	}
	// Drop the object.
	p.RemoveFromInventory(o)
	if p.IsInsideSpaceship() || p.curLoc.IsLocked() || p.CurSys() != o.HomeSystem() {
		o.Recycle()
		p.Outputm(text.MN516)
		return
	}
	o.SetCurLocNo(p.LocNo) // FIXME
	// Tell the player.
	p.Outputm(text.DropOK, o.DisplayName(false))
	// Tell the others in the location.
	if !p.IsInsideSpaceship() {
		msg := text.Msg(text.DropAudienceTell, p.Name, o.DisplayName(false))
		p.curLoc.Talk(msg, p)
	}
}
