package server

import (
	"math/rand/v2"

	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

// Duplicates another player's input and output to the SPYing player's screen.
func (p *Player) CmdSpy(playerName string) {
	if !p.HasSpyBeam() && !p.IsInSolLocation(sol.SecretRoom) {
		p.Outputm(text.SPY_NO_SPYBEAM)
		return
	}

	if playerName == "" {
		if p.spied == nil {
			p.Outputm(text.NotSpying)
			return
		}
		p.Outputm(text.Spy, p.spied.Name)
		return
	}

	subject, ok := p.CurLoc().FindPlayer(playerName)
	if !ok {
		p.Outputm(text.SPY_CANT_SPY)
		return
	}

	p.StopSpying()

	// Trying to spy themself?
	if subject == p {
		p.Outputm(text.DontBeSilly)
		return
	}

	// Is the subject spyable?
	if !p.canSpy(subject) {
		p.Outputm(text.SPY_CANT_SPY)
		return
	}

	// If spyer is at the public beam give them an access code for Nisrik.
	if p.IsInSolLocation(sol.SecretRoom) {
		p.Crypto = rand.IntN(9000) + 1000 //nolint:gosec // "It's Just A Game"
		p.Outputm(text.MN616, p.Crypto)
	}

	// Locked on!
	subject.AddSpyer(p)
	p.spied = subject
	p.Outputm(text.SPY_OK, p.spied.Name)
}

// Turns off the player's spybeam equipment.
func (p *Player) CmdSpyOff() {
	if !p.HasSpyBeam() && !p.IsInSolLocation(sol.SecretRoom) {
		p.Outputm(text.SPY_NO_SPYBEAM)
		return
	}

	if p.spied == nil {
		p.Outputm(text.NotSpying)
		return
	}

	// p.StopSpying()
	// p.Outputm(text.SPYOFF_OK)

	p.Output("Not implemented. Check back in 2 weeks.\n")
}
