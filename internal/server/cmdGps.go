package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdGPS() {
	if !p.Session().IsRoboBod() && p.Rank() != model.RankDeity {
		p.UnknownCommand()
		return
	}
	locNo := p.CurLocNo()
	if p.IsFlyingSpaceship() {
		locNo = p.ShipLocNo()
	}
	p.Nsoutputm(text.GPS, p.CurSysName(), locNo, p.ShipLocNo())
}
