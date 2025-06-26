package server

import "github.com/nosborn/federation-1999/internal/model"

// Allows low level players to gamble their job ratings at roulette.
func (p *Player) CmdGamble(colorWager model.GambleColour, stake int32) {
	p.Output("Not implemented. Check back in 2 weeks.\n")
}
