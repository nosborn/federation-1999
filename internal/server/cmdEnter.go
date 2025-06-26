package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/text"
)

// CmdEnter is used to get a player into his/her spaceship.
func (p *Player) CmdEnter() {
	if !p.HasSpaceship() {
		p.Outputm(text.NoSpaceship)
		return
	}
	if p.IsInsideSpaceship() {
		p.Outputm(text.MN519)
		return
	}
	if p.CurLocNo() != p.ShipLocNo() {
		p.Outputm(text.MN520)
		return
	}
	// Stash the landing pad location, we'll need it later.
	landingPad := p.CurLoc()
	// Move into the airlock.
	p.setLocation(4) // shipAirlock
	p.CurLoc().Describe(p, DefaultDescription)
	// Notify any players on the landing pad.
	var msgNo text.MsgNum
	switch p.Sex() {
	case model.SexFemale:
		msgNo = text.EnteredShip_F
	case model.SexMale:
		msgNo = text.EnteredShip_M
	case model.SexNeuter:
		msgNo = text.EnteredShip_N
	}
	msg := text.Msg(msgNo, p.Name)
	landingPad.Talk(msg, p)
}
