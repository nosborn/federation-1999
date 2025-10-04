package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/global"
)

func (p *Player) CmdTrace(trace model.Trace) {
	if !global.TestFeaturesEnabled && p.Rank() != model.RankDeity {
		p.UnknownCommand()
		return
	}

	if p.MsgOutSpyDepth == spyPublic {
		p.MsgOutSpyDepth = spyPrivate
	}

	var buf []byte
	if trace == model.TracePerivale {
		buf = []byte{model.DLE, model.LeTrace, 'P'}
	} else {
		buf = []byte{model.DLE, model.LeTrace, '-'}
	}
	p.sendOutput(string(buf), spyPrivate)

	p.Nsoutput("OK\n")
}
