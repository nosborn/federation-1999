package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdTune(channel int32) {
	if !p.HasCommUnit() {
		p.Outputm(text.NoCommUnit)
		return
	}

	if channel < 1 || channel > model.MAX_XT_CHANNEL {
		p.Outputm(text.MN186, model.MAX_XT_CHANNEL-1)
		return
	}

	if channel == model.MAX_XT_CHANNEL && p.Rank() < model.RankHostess {
		p.Outputm(text.TUNE_RESERVED)
		return
	}

	p.SetChannel(channel)
	if channel == 1 {
		p.Outputm(text.TuneOK_1)
	} else {
		p.Outputm(text.TuneOK)
	}
}

func (p *Player) CmdTuneOff() {
	if !p.HasCommUnit() {
		p.Outputm(text.NoCommUnit)
		return
	}

	if p.Channel() == 0 {
		p.Outputm(text.XT_ALREADY_OFF)
		if !p.IsCommsOff() {
			p.Outputm(text.TRY_COMMS_OFF)
		}
		return
	}

	p.SetChannel(0)
	p.Outputm(text.XT_OFF)

	if !p.IsCommsOff() {
		p.Outputm(text.COMMS_STILL_ON)
	}
}
