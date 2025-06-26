package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/text"
)

// CmdBump can be used by on-duty Navigators so all output needs to be
// spyproof.
func (p *Player) CmdBump(name string) {
	effectiveRank := p.rank
	if p.IsOnDutyNavigator() {
		effectiveRank = model.RankHostess
	}
	if effectiveRank < model.RankHostess || effectiveRank > model.RankManager {
		p.UnknownCommand()
		return
	}

	// if p.MsgOut.spyDepth == spyPublic {
	// 	p.MsgOut.spyDepth = spyPrivate; // Suppress command echo
	// }

	offender, ok := FindPlayer(name)
	if !ok {
		p.Nsoutputm(text.PlayerNotFound)
		return
	}
	if offender.IsPlaying() {
		p.Nsoutputm(text.PlayerNotPresent, offender.Name())
		return
	}
	if offender == p {
		p.Nsoutput("Why don't you use QUIT like everyone else?\n")
		return
	}
	if offender.rank == effectiveRank {
		p.Nsoutputm(text.DontBeSilly)
		return
	}
	if offender.rank > effectiveRank {
		p.Nsoutputm(text.MustBeJoking)
		return
	}

	// TODO:
	p.Output("Not implemented. Check back in 2 weeks.\n")
}
