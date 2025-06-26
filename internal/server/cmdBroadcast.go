package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/text"
)

// CmdBroadcast passes a message to all players.
func (p *Player) CmdBroadcast(message string) {
	if p.Rank() != model.RankManager && p.Rank() != model.RankDeity {
		p.UnknownCommand()
		return
	}
	if !p.HasCommUnit() {
		p.Outputm(text.NoCommUnit)
		return
	}
	broadcast := text.Msg(text.BROADCAST, p.Name, message)
	for _, recipient := range Players {
		if !recipient.IsPlaying() {
			continue
		}
		if recipient.IsCommsOff() {
			continue
		}
		if recipient == p {
			continue
		}
		recipient.Nsoutput(broadcast)
		recipient.FlushOutput()
	}
	p.Outputm(text.BROADCAST_SENT)
}
