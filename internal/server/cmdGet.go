package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

// Allows a player to pick up an object in the game. Checks to see if the
// object is in the same room, and calls the event handler if there is an event
// associated with getting the object.
func (p *Player) CmdGet(objectName model.Name) {
	if p.IsInsideSpaceship() {
		p.Outputm(text.GetNotFound)
		return
	}
	o, ok := p.CurLoc().FindObjectName(objectName)
	if !ok {
		p.Outputm(text.GetNotFound)
		return
	}
	// Deal with any get event.
	// if p.CurSys().GetHook(p, o) {
	// 	return
	// }
	if o.HomeSystem() == p.CurSys() {
		if o.GetEvent() != 0 { //nolint:staticcheck // SA9003: empty branch
			// if (!x_event_handler(*this, theObject->get_event)) {
			// 	return
			// }
		}
	}
	if o.IsMobile() {
		p.Outputm(text.DontBeSilly)
		return
	}
	if !p.canCarry(o.Weight()) {
		p.Outputm(text.GetTooHeavy)
		return
	}
	o.SetCurLocNo(0) // FIXME: remove from location's object list too
	p.inventory = append(p.inventory, o)
	p.Outputm(text.GetOK)
	if o.Number() == sol.ObBlackBox {
		p.checkForBlackBox()
	}
	msg := text.Msg(text.GetTellAudience, p.Name, o.DisplayName(false))
	p.curLoc.Talk(msg, p)
}
