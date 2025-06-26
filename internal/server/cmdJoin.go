package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

const (
	joinFee = 20000000 // 20,000,000 IG
)

// Planet owner joins a dutchy.
func (p *Player) CmdJoin(duchyName string) {
	if p.Rank() < model.RankSquire || p.Rank() > model.RankDuke || p.OwnSystem() == nil {
		p.Outputm(text.NOT_A_PLANET_OWNER)
		return
	}

	if p.Rank() == model.RankDuke || p.IsPromoCharacter() {
		p.Outputm(text.DontBeSilly)
		return
	}

	// Find this person's planet.
	if p.OwnSystem().IsClosed() {
		p.Outputm(text.ClosedForBusiness, p.OwnSystem().Name())
		return
	}

	// FIXME:
	// // Is the planet already in a duchy?
	// if !p.OwnSystem().duchy().Name != "Sol" {
	// 	p.Outputm(text.JoinAlreadyJoined)
	// 	return
	// }

	// Try to find the duchy.
	d, ok := FindDuchy(duchyName)
	if !ok || d.IsHidden() {
		p.Outputm(text.NoSuchDuchy)
		return
	}

	// Nobody can join Sol, that's what SECEDE is for.
	if d.Name() == "Sol" {
		p.Outputm(text.DontBeSilly)
		return
	}

	// Find the capital planet and make sure it's open for business. This
	// is the wrong message, but it will do!
	s, ok := FindSystem(d.Name())
	if !ok || s.IsClosed() {
		p.Outputm(text.JoinFullDuchy, d.Name())
		return
	}

	// Is the duchy full?
	// if d.members() >= 20 /* Duchy::MAX_MEMBERS */) {
	// 	p.outputf(messages["JoinFullDuchy"], d.Name)
	// 	return
	// }

	// Enough cash in the planet treasury?
	if p.OwnSystem().Balance() < joinFee {
		p.Outputm(text.NO_TREASURY_FUNDS)
		return
	}

	// All OK, sign the planet up and take the money...
	p.OwnSystem().SetDuchy(d)
	p.OwnSystem().Expenditure(joinFee)

	// ...pay the joining fee over to the duchy...
	s.Income(joinFee, true)
	// s.save(database.SaveNow) -- FIXME

	// ...and tell the player.
	p.Outputm(text.JoinReport, p.OwnSystem().Name(), d.Name())
	p.Save(database.SaveNow)

	// Tell the Duke.
	// duke := findPlayer(d.Owner())
	if d.Owner().IsPlaying() {
		d.Owner().Outputm(text.JoinReport, p.OwnSystem().Name(), d.Name())
		d.Owner().FlushOutput()
	}
}
