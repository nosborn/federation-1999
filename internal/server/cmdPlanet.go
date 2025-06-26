package server

import "github.com/nosborn/federation-1999/internal/model"

func (p *Player) CmdPlanet() {
	if !p.CurLoc().IsSpace() {
		if p.Rank() == model.RankDeity {
			planet := p.GuessCurrentPlanet()
			if planet == nil {
				p.Output("You check your bearings and decide that you're lost!\n")
			} else {
				p.outputf("You check your bearings and decide that you're on %s.\n", planet.Name())
			}
			return
		}
	}
	if p.CurLoc().IsOrbit() {
		p.Output("You're already close enough to the nearest planet!\n")
		return
	}
	p.CmdGo(model.DirectionPlanet)
}
