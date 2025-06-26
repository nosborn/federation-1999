package server

import (
	"log"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdSuicide() {
	// Fetch out the Swiss Army penknife...
	log.Printf("%s has committed suicide", p.Name())
	p.Outputm(text.Suicide)

	// ...tell any spectators...
	if !p.IsInsideSpaceship() && !p.IsInSolLocation(sol.MeetingPoint) {
		var msgNo text.MsgNum
		switch p.Sex() {
		case model.SexFemale:
			msgNo = text.Suicide_F
		case model.SexMale:
			msgNo = text.Suicide_M
		default:
			msgNo = text.Suicide_N
		}
		msg := text.Msg(msgNo, p.Name)
		p.curLoc.Talk(msg, p)
	}

	// ...and slash!
	p.Die()
}
