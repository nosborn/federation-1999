package server

import (
	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdTour() {
	if !p.IsInSolLocation(1 /* shipCommandCentre */) {
		if !p.HasSpaceship() {
			p.Outputm(text.TourNoSpaceship)
		} else {
			p.Outputm(text.TourWrongPlace)
		}
		return
	}
	if p.ShipLocNo() != sol.EarthLandingArea {
		p.Outputm(text.TourWrongPlace)
		return
	}
	debug.Check(p.Session() != nil)
	p.Session().SwitchToTour()
}
