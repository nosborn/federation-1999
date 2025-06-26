package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/text"
)

// Quick who - formatted listing of the names (only) of players on the game.
func (p *Player) CmdQuickWho() {
	playing := 0
	for i := range Players {
		if Players[i].IsPlaying() {
			p.outputf("%-15s ", Players[i].Name())
			playing++
		}
	}
	if playing == 1 {
		p.Outputm(text.WHO_1_PLAYER)
	} else {
		p.Outputm(text.WHO_ALL_PLAYERS, playing)
	}
}

// Quick who - formatted listing of the names (only) of players on the game.
func (p *Player) CmdQuickWhoChannel(channel int32) {
	if channel < 0 || channel > model.MAX_XT_CHANNEL {
		p.Outputm(text.WHO_BAD_CHANNEL)
		return
	}

	p.Output("Not implemented. Check back in 2 weeks.\n")
}

// Quick who - formatted listing of the names (only) of players on the game.
func (p *Player) CmdQuickWhoSystem(sysName string) {
	s, ok := FindSystem(sysName)
	if !ok || (s.IsHidden() && p.Rank() < model.RankHostess) {
		p.Outputm(text.WHO_BAD_SYSTEM)
		return
	}

	p.Output("Not implemented. Check back in 2 weeks.\n")
}
