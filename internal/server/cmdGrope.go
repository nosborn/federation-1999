package server

import (
	"math/rand/v2"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

var gropeMsgNos = [7][3]text.MsgNum{
	{text.Grope_Friendly, text.GropeRecipient_Friendly, text.GropeBystander_Friendly},
	{text.Grope_Hot, text.GropeRecipient_Hot, text.GropeBystander_Hot},
	{text.Grope_Nice, text.GropeRecipient_Nice, text.GropeBystander_Nice},
	{text.Grope_Passionate, text.GropeRecipient_Passionate, text.GropeBystander_Passionate},
	{text.Grope_Sloppy, text.GropeRecipient_Sloppy, text.GropeBystander_Sloppy},
	{text.Grope_Tender, text.GropeRecipient_Tender, text.GropeBystander_Tender},
	{text.Grope_Warm, text.GropeRecipient_Warm, text.GropeBystander_Warm},
}

func (p *Player) CmdGrope(targetName string) {
	target, ok := p.CurLoc().FindPlayer(targetName)
	if !ok {
		mobileName := model.Name{The: false, Words: 1, Text: targetName}
		mobile, ok := p.CurLoc().FindObjectName(mobileName)
		if ok && mobile.Number() == sol.ObDiesel {
			p.Outputm(text.SLAP_RECIPIENT, mobile.DisplayName(true))
			locMsg := text.Msg(text.SLAP_AUDIENCE, mobile.DisplayName(true), p.Name())
			p.CurLoc().Talk(locMsg, p)
			if p.CurSta() > 2 {
				p.SetCurSta(p.CurSta() - 1)
				p.Save(database.SaveNow)
			}
			return
		}
	}
	if !ok || !target.IsPlaying() {
		p.Outputm(text.MN20)
		return
	}
	if target == p {
		p.Outputm(text.GropeSelf)
		return
	}
	if p.IsInsideSpaceship() {
		p.Outputm(text.MN20)
		return
	}

	idx := rand.IntN(len(gropeMsgNos))
	// Tell the person doing the groping.
	p.Outputm(gropeMsgNos[idx][0], target.Name())
	// Tell the recipient.
	target.Outputm(gropeMsgNos[idx][1], target.Name())
	target.FlushOutput()
	// Tell any other players in the room.
	locMsg := text.Msg(gropeMsgNos[idx][2], p.Name(), target.Name())
	p.CurLoc().Talk(locMsg, p, target)
}
