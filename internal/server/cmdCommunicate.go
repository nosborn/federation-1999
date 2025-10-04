package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdCommunicate(message string) {
	if !p.HasCommUnit() {
		p.Outputm(text.NoCommUnit)
		return
	}
	if p.IsCommsOff() {
		if p.Rank() != model.RankGroundHog || p.LocNo != sol.MeetingPoint {
			p.Outputm(text.COM_WHILE_COMMSOFF)
			return
		}
		p.flags2 &^= PL2_COMMS_OFF
		// debug.Check(!IsCommsOff())
	}
	if p.IsInSolSystem() {
		if p.Rank() < model.RankHostess && !p.IsOnDutyNavigator() {
			p.Outputm(text.MN440)
			return
		}
	}
	// Tell the player what's happened.
	// punctuation string
	// if (isShouting(theMessage)) {
	// 	punctuation = "!"
	// 	output("mn10")
	// } else {
	// 	punctuation = "."
	// 	output("mn9")
	// }
	// if (ispunct(theMessage[strlen(theMessage) - 1])) {
	// 	punctuation = ""
	// }
	// Broadcast the message.
	// FIXME: Sort out punctuation
	incomingComm := text.Msg(text.MN897, p.Name, message)
	for _, recipient := range Players {
		if recipient.curSys.Duchy() != p.CurSys().Duchy() {
			continue
		}
		if !recipient.IsPlaying() {
			continue
		}
		if recipient.IsCommsOff() {
			continue
		}
		if recipient == p {
			continue
		}
		recipient.Nsoutput(incomingComm)
		recipient.FlushOutput()
	}
}
