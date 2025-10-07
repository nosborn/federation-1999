package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

const (
	changeCost = 500 // 500 IG
)

// CmdChange changes a player's sex.
func (p *Player) CmdChange() {
	if !p.curLoc.IsHospital() {
		p.Outputm(text.ChangeNotHere)
		return
	}
	if p.CurSys().IsClosed() {
		p.Outputm(text.ClosedForBusiness, p.CurSys().Name())
		return
	}
	// Ming can't change sex.
	// Selena can't change sex.
	if p.Rank() == model.RankEmperor || p.CustomRank == CustomRankOfTheSpaceways {
		p.Outputm(text.DontBeSilly)
		return
	}
	if p.Sex() == model.SexNeuter {
		p.Outputm(text.ChangeCantChange)
		return
	}
	if p.HasWallet() {
		if p.Balance() < changeCost {
			p.Outputm(text.TooExpensive, changeCost)
			return
		}
		p.ChangeBalance(-changeCost)
		p.CurSys().Income(changeCost, true)
	}
	switch p.Sex() { //nolint:exhaustive // model.SexNeuter excluded above
	case model.SexFemale:
		p.sex = model.SexMale
	case model.SexMale:
		p.sex = model.SexFemale
	}
	p.Outputm(text.ChangeOK)
	p.Save(database.SaveNow)
}
