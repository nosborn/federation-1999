package server

import (
	"github.com/nosborn/federation-1999/internal/text"
)

// CmdAct reports player action with player defined modifier.
func (p *Player) CmdAct(action string) {
	if p.IsInsideSpaceship() {
		p.Outputm(text.ActInsideSpaceship)
		return
	}
	// The actor gets the message as well.
	msg := text.Msg(text.ActAudienceTell, p.Name, action)
	p.CurLoc().Talk(msg)
}
