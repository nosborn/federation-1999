package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdSystems() {
	p.Outputm(text.SystemsHeader)

	for s := range allSystems.Values() {
		if s.LoadState() != SystemOnline {
			continue
		}
		msgID := text.SystemsNormalEntry
		if s.IsHidden() {
			if p.Rank() != model.RankDeity {
				continue
			}
			msgID = text.SystemsHiddenEntry
		}
		switch s.Name() {
		case "Arena", "Horsell", "Snark", "Sol":
			for i, planet := range s.Planets() {
				name := s.Name()
				if i > 0 {
					name = ""
				}
				p.Outputm(msgID, name, planet.name, planet.LevelDescription())
			}
		default:
			p.Outputm(msgID, s.Name(), s.Planets()[0].name, s.Planets()[0].LevelDescription())
		}
	}
}
