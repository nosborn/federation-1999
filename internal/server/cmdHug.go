package server

import (
	"math/rand/v2"

	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

var hugMsgNos = [7][3]text.MsgNum{
	{text.Hug_Friendly, text.HugRecipient_Friendly, text.HugBystander_Friendly},
	{text.Hug_Hot, text.HugRecipient_Hot, text.HugBystander_Hot},
	{text.Hug_Nice, text.HugRecipient_Nice, text.HugBystander_Nice},
	{text.Hug_Passionate, text.HugRecipient_Passionate, text.HugBystander_Passionate},
	{text.Hug_Sloppy, text.HugRecipient_Sloppy, text.HugBystander_Sloppy},
	{text.Hug_Tender, text.HugRecipient_Tender, text.HugBystander_Tender},
	{text.Hug_Warm, text.HugRecipient_Warm, text.HugBystander_Warm},
}

func (p *Player) CmdHug(targetName string) {
	target, ok := p.CurLoc().FindPlayer(targetName)
	if !ok || !target.IsPlaying() {
		p.Outputm(text.MN20)
		return
	}
	if target == p {
		p.Outputm(text.DontBeSilly)
		return
	}
	if p.IsInsideSpaceship() {
		p.Outputm(text.MN20)
		return
	}

	idx := rand.IntN(len(hugMsgNos))
	// Tell the person doing the hugging.
	p.Outputm(hugMsgNos[idx][0], target.Name())
	// Tell the recipient.
	target.Outputm(hugMsgNos[idx][1], p.Name())
	target.FlushOutput()
	// Tell any other players in the room.
	locMsg := text.Msg(hugMsgNos[idx][2], p.Name(), target.Name())
	p.CurLoc().Talk(locMsg, p, target)

	// Stop sulking.
	if p.IsSulking() {
		p.SetSulking(false)
		p.Save(database.SaveNow)
	}
}
