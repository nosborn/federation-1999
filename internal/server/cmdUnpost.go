package server

import "github.com/nosborn/federation-1999/internal/model"

func (p *Player) CmdUnpost(name string) {
	if p.Rank() != model.RankHostess && p.Rank() != model.RankManager {
		if !p.IsOnDutyNavigator() {
			p.UnknownCommand()
			return
		}
	}

	p.Output("Not implemented. Check back in 2 weeks.\n")
}
