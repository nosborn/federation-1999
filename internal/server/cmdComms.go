package server

import (
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdComms() {
	if !p.HasCommUnit() {
		p.Outputm(text.NoCommUnit)
		return
	}
	if p.IsCommsOff() {
		p.Outputm(text.CommsAreOff)
	} else {
		p.Outputm(text.CommsAreOn)
	}
}

func (p *Player) CmdCommsOff() {
	if !p.HasCommUnit() {
		p.Outputm(text.NoCommUnit)
		return
	}
	if p.IsCommsOff() {
		p.Outputm(text.CommsAlreadyOff)
		if p.Channel() != 0 {
			p.Outputm(text.TRY_XT_OFF)
		}
		return
	}
	p.SetCommsOff(true)
	p.Outputm(text.CommsNowOff)
	if p.Channel() != 0 {
		p.Outputm(text.XT_STILL_ON)
	}
}

func (p *Player) CmdCommsOn() {
	if !p.HasCommUnit() {
		p.Outputm(text.NoCommUnit)
		return
	}
	p.SetCommsOff(false)
	p.Outputm(text.CommsNowOn)
}
