package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

// Passes a message to all players who have their comms on.
func (p *Player) CmdTransmit(message string) {
	if !p.HasCommUnit() {
		p.Outputm(text.NoCommUnit)
		return
	}

	if p.channel == 0 {
		if p.Rank() == model.RankGroundHog && p.LocNo == sol.MeetingPoint {
			p.channel = 1
		} else {
			p.Outputm(text.MN915)
			return
		}
	}

	var punctuation string
	if text.IsShouting(message) {
		// punctuation = "!"
		p.Outputm(text.MN10)
	} else {
		// punctuation = "."
		p.Outputm(text.MN9)
	}
	// if ispunct(message[1 : len(message)-1]) {
	// 	punctuation = ""
	// }
	incoming := text.Msg(text.IncomingXT, p.Name, message, punctuation)

	for _, op := range Players {
		if !op.IsPlaying() {
			continue
		}
		if !op.HasCommUnit() {
			continue
		}
		if op.Channel() != p.channel {
			continue
		}
		if p.IsInHorsellSystem() || op.IsInHorsellSystem() {
			if p.CurSys() != op.CurSys() {
				continue
			}
		}
		if op != p {
			op.Nsoutput(incoming)
			op.FlushOutput()
		}
	}
}
