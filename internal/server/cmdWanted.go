package server

import (
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdWanted() {
	p.Nsoutput(stardate())
	p.Nsoutputm(text.WantedHeader)

	// Mobiles first...

	// p.currentSystem.wantedHook(actor) -- TODO

	// ...then players.

	// TODO

	p.Output("Not implemented. Check back in 2 weeks.\n")
}
