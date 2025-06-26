package server

import (
	"github.com/nosborn/federation-1999/internal/text"
)

// CmdAlarm allows a player to set up an alarm to go off at a later time (but
// within the next 24 hours).
func (p *Player) CmdAlarm(minutes int32) {
	if minutes == 0 {
		p.CmdAlarmOff()
		return
	}
	if minutes < 0 || minutes > 24*60 {
		p.Outputm(text.DontBeSilly)
		return
	}
	p.SetAlarm(minutes)
	p.Outputm(text.AlarmOK, text.HumanizeMinutes(int64(minutes)))
}

// CmdAlarmOff cancels any pending alarm call.
func (p *Player) CmdAlarmOff() {
	p.CancelAlarm()
	p.Outputm(text.AlarmCancelled)
}
