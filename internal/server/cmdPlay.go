package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/horsell"
	"github.com/nosborn/federation-1999/internal/text"
)

// Allows a player to play a musical instrument which is either in the room, or
// the player's inventory.
func (p *Player) CmdPlay(music string) {
	if p.CurSys().Name() == "Horsell" && p.LocNo == horsell.OrganLoft {
		p.playOrgan(music)
		return
	}

	var instrument *Object
	for _, o := range p.inventory {
		if (o.Flags & model.OfMusic) != 0 {
			instrument = o
			break
		}
	}
	if instrument == nil {
		for _, o := range p.CurSys().Objects() {
			if o.IsHidden() || o.IsRecycling() {
				continue
			}
			if (o.Flags & model.OfMusic) != 0 {
				instrument = o
				break
			}
		}
	}
	if instrument == nil {
		p.Outputm(text.PlayNoInstrument)
		return
	}

	p.Outputm(text.MN786, music, instrument.Name)
	if !p.IsInsideSpaceship() {
		msg := text.Msg(text.MN785, p.Name, music, instrument.Name)
		p.curLoc.Talk(msg, p)
	}
}

func (p *Player) playOrgan(music string) {
	p.Output("Not implemented. Check back in 2 weeks.\n")
}
