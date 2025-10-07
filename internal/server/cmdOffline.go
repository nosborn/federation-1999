package server

import (
	"log"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdOffline() {
	// FIXME: This should allow for over-enthusiastic Explorers.

	if p.Rank() < model.RankSquire || p.Rank() > model.RankDuke || p.OwnSystem() == nil {
		p.Outputm(text.NOT_A_PLANET_OWNER)
		return
	}

	if p.OwnSystem().IsClosed() {
		switch {
		case p.OwnSystem().IsLoading():
			p.Outputm(text.OfflineStillLoading)
		case p.OwnSystem().IsUnloading():
			p.Outputm(text.OfflineAlreadyUnloading)
		default:
			p.Outputm(text.OfflineAlreadyOffline)
		}
		return
	}

	if p.IsInHorsellSystem() {
		p.Outputm(text.OfflineInHorsell) // FIXME: Temporary message text.
		return
	}

	log.Printf("%s is going off-line at the owner's request", p.OwnSystem().Name())
	p.OwnSystem().(*PlayerSystem).BeginUnload()
	p.Save(database.SaveNow)

	// Stop any build that's in progress.
	if p.buildTimer != nil {
		p.buildTimer.Stop()
		p.buildTimer = nil
	}

	// Tell the player what's happening.
	msgNo := text.OfflineSuccess
	for _, player := range Players {
		if player.curSys == p.OwnSystem() {
			msgNo = text.OfflineDelay
			break
		}
	}
	p.Outputm(msgNo, p.OwnSystem().Name())
}
