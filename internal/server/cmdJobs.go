package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdJobs() {
	if (p.Flags0 & model.PL0_JOB) == 0 {
		p.Outputm(text.JobsAreOff)
	} else {
		p.Outputm(text.JobsAreOn)
	}
}

// Clears the player's PL0_JOB flag so that s/he is no longer notified when
// jobs become available.
func (p *Player) CmdJobsOff() {
	if (p.Flags0 & model.PL0_JOB) != 0 {
		p.Flags0 &^= model.PL0_JOB
		p.Save(database.SaveNow)
	}
	p.Outputm(text.JobsNowOff)
}

// Sets the player's PL0_JOB flag so that s/he is notified when jobs become
// available.
func (p *Player) CmdJobsOn() {
	if (p.Flags0 & model.PL0_JOB) == 0 {
		p.Flags0 |= model.PL0_JOB
		p.Save(database.SaveNow)
	}
	p.Outputm(text.JobsNowOn)
}
