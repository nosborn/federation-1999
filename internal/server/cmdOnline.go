package server

import (
	"log"
	"strings"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/build"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdOnline(name string) {
	if name == "" {
		// Make an allowance here for confused and/or impatient Explorers.
		if p.Rank() < model.RankSquire || p.Rank() > model.RankDuke || p.OwnSystem() == nil {
			if p.Rank() == model.RankExplorer {
				if p.OwnSystem() != nil && p.OwnSystem().IsLoading() {
					p.Outputm(text.OnlineAlreadyLoading)
					return
				}
				// if p.workbenchAccess(p.uid) != WB_ACCESS_OK { -- TODO
				// 	p.Outputm(text.OnlineSupplyName);
				// 	return;
				// }
			}
			p.Outputm(text.NOT_A_PLANET_OWNER)
			return
		}

		// Make sure the player's system is currently offline.
		if !p.OwnSystem().IsOffline() {
			switch {
			case p.OwnSystem().IsLoading():
				p.Outputm(text.OnlineAlreadyLoading)
			case p.OwnSystem().IsUnloading():
				p.Outputm(text.OnlineStillUnloading)
			default:
				p.Outputm(text.OnlineAlreadyOnline)
			}
			return
		}

		// Queue up for loading.
		log.Printf("%s is coming on-line at the owner's request", p.OwnSystem().Name())
		p.OwnSystem().(*PlayerSystem).BeginLoad()
		p.Outputm(text.OnlineLoadStarted, LoaderQueueLength())
		p.Save(database.SaveNow)
	} else {
		if p.Rank() != model.RankExplorer {
			if p.OwnsPlanet() && strings.EqualFold(name, p.OwnSystem().Name()) {
				p.CmdOnline("")
				return
			}
			p.UnknownCommand()
			return
		}

		// Do they already have a planet load in progress?
		if p.OwnSystem() != nil {
			p.Outputm(text.OnlineAlreadyLoading)
			return
		}

		// Have they built an Interstellar Link?
		if p.BuildProject != build.LINK || !p.IsConstructionComplete() {
			p.Outputm(text.OnlineNoLink)
			return
		}

		p.Output("Not implemented. Check back in 2 weeks.\n")
	}
}
