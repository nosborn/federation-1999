package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

// Allow the Emperor to reduce another player's stamina right down to 2 + 1.
func (p *Player) CmdFlog(name string) {
	if p.Rank() != model.RankEmperor {
		p.UnknownCommand()
		return
	}
	floggee, ok := p.CurLoc().FindPlayer(name)
	if !ok {
		p.Outputm(text.MN20)
		return
	}
	if floggee == p {
		p.Outputm(text.DontBeSilly)
		return
	}
	if p.IsInsideSpaceship() {
		p.Outputm(text.MN20)
		return
	}
	// Tell the flogger...
	var msgNo text.MsgNum
	switch floggee.Sex() {
	case model.SexFemale:
		msgNo = text.FlogEmperor_F
	case model.SexMale:
		msgNo = text.FlogEmperor_M
	default:
		msgNo = text.FlogEmperor_N
	}
	p.Outputm(msgNo, floggee.Name)
	// ...the flogee...
	floggee.Outputm(text.FlogFloggee)
	floggee.FlushOutput()
	if floggee.CurSta() > 2 {
		floggee.SetCurSta(2)
		floggee.SetTimerCount(99)
	}
	floggee.Save(database.SaveNow)
	// ...and any onlookers.
	switch floggee.Sex() {
	case model.SexFemale:
		msgNo = text.FlogSpectator_F
	case model.SexMale:
		msgNo = text.FlogSpectator_M
	default:
		msgNo = text.FlogSpectator_N
	}
	locMsg := text.Msg(msgNo, floggee.Name())
	p.CurLoc().Talk(locMsg, p, floggee)
}
