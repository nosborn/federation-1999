package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

// CmdClearEmbargo enables a Duke to clear the embargoed status on planets in
// his/her duchy.
func (p *Player) CmdClearEmbargo() {
	if p.Rank() != model.RankDuke {
		p.Outputm(text.MUST_BE_DUKE)
		return
	}
	if p.ownDuchy == nil {
		p.Outputm(text.NO_DUCHY_FOR_DUKE)
		return
	}
	if p.ownDuchy.ClearEmbargo(p) {
		p.Save(database.SaveNow)
	}
}

func (p *Player) CmdClearFactory(factoryNo int32) {
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

// CmdClearFavoured enables a Duke to clear the favoured status on planets in
// his/her duchy.
func (p *Player) CmdClearFavoured() {
	if p.Rank() != model.RankDuke {
		p.Outputm(text.MUST_BE_DUKE)
		return
	}
	if p.ownDuchy == nil {
		p.Output("NO_DUCHY_FOR_DUKE")
		return
	}
	if p.ownDuchy.ClearFavoured(p) {
		p.Save(database.SaveNow)
	}
}

func (p *Player) CmdClearMilkrun(slotNo int32) {
	p.Output("Not implemented. Check back in 2 weeks.\n")
}
