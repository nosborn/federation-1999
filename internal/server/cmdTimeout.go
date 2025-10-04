package server

import (
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdTimeout(minutes ...int32) {
	// options := &TimeoutOptions{}
	// for _, opt := range opts {
	// 	opt(options)
	// }

	if len(minutes) == 0 {
		timeout := p.Session().GetTimeout()
		if timeout > 0 {
			p.Outputm(text.TimeoutIsMinutes, text.HumanizeMinutes(int64(timeout)))
		} else {
			p.Outputm(text.TimeoutIsOff)
		}
		return
	}

	if minutes[0] == 0 {
		p.CmdTimeoutOff()
		return
	}
	if minutes[0] < 0 || minutes[0] > 24*60 {
		p.Outputm(text.DontBeSilly)
		return
	}
	p.Session().SetTimeout(minutes[0])
	p.Outputm(text.TimeoutSet, text.HumanizeMinutes(int64(minutes[0])))
}

func (p *Player) CmdTimeoutOff() {
	p.Session().SetTimeout(0)
	p.Outputm(text.TimeoutTurnedOff)
}
