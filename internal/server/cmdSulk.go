package server

import (
	"github.com/nosborn/federation-1999/internal/text"
)

// Sets the sulk flag, and tells other players in the room that the player
// concerned is sulking.
func (p *Player) CmdSulk() {
	if p.IsSulking() {
		p.Outputm(text.SulkAlreadySulking)
		return
	}

	p.SetSulking(true)
	p.Outputm(text.SulkOK)

	// Tell other folks in the room that they've gone off in a huff.
	if !p.IsInsideSpaceship() {
		locMsg := text.Msg(text.SulkAudienceTell, p.Name)
		p.CurLoc().Talk(locMsg, p)
	}
}
