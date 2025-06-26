package server

import (
	"math/rand/v2"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

// CmdCheat is for those who like to cheat at games! It just produces a text
// answer and a fine!
func (p *Player) CmdCheat() {
	if p.Rank() == model.RankEmperor {
		p.Outputm(text.DontBeSilly)
		return
	}
	var title string
	switch p.Sex() {
	case model.SexFemale:
		title = "Empress"
	case model.SexMale:
		title = "Emperor"
	default:
		title = "Emperoid"
	}
	if p.HasWallet() {
		p.ChangeBalance(-(rand.Int32N(100) + 1)) //nolint:gosec // "It's Just A Game"
	}
	p.Outputm(text.Cheat, title, p.Name())
	p.Save(database.SaveNow)
}
