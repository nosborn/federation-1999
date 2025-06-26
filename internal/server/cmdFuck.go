package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

// Yes, well... players will be players. :)
func (p *Player) CmdFuck(name model.Name) {
	if p.IsInSolLocation(sol.ChezDiesel) {
		o, ok := p.CurLoc().FindObjectName(name)
		if ok && o.Number() == sol.ObDiesel {
			if o.Value() < 0 {
				p.ChangeReward(-o.Value())
			}
			p.Outputm(text.DieselSlugs, o.DisplayName(true))
			locMsg := text.Msg(text.DieselSlugs_Audience, p.Name())
			p.CurLoc().Talk(locMsg, p)
			p.Die()
			return
		}
	}
	p.Output("Behave!\n")
}
