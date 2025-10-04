package server

import "github.com/nosborn/federation-1999/internal/model"

func (p *Player) CmdWhereIs(name model.Name) {
	if p.Rank() != model.RankDeity {
		p.UnknownCommand()
		return
	}

	o, ok := p.CurSys().FindObjectName(name)
	if !ok {
		p.Output("I haven't the faintest idea!\n")
		return
	}

	if o.IsHidden() {
		p.Output("Currently out of the game.\n")
		return
	}
	if o.IsRecycling() {
		p.Output("Currently recycling.\n")
		return
	}
	if o.CurLocNo() == 0 {
		p.Output("Currently being carried.\n")
		return
	}
	p.outputf("Currently in location %d.\n", o.CurLocNo())
}
