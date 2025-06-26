package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

// Player signs contract to carry goods for another player.
func (p *Player) CmdSign() {
	if p.deal != nil {
		p.signContract()
		return
	}

	if p.Job.Status != JOB_OFFERED {
		p.Outputm(text.MN982)
		return
	}

	p.Job.Status = JOB_ACCEPTED
	p.Count[model.PL_G_JOB] = 0
	p.Outputm(text.MN983, p.Job.From)

	p.Save(database.SaveNow)

	// Notify the owner of the cargo.
	owner, ok := FindPlayer(p.Job.GenWk.Owner)
	if ok && owner.IsPlaying() {
		owner.Outputm(text.MN984, p.name)
		owner.FlushOutput()
	}
}

// Completes the transfer of funds and goods in player/player trading.
func (p *Player) signContract() {
	if p.deal == nil {
		p.Outputm(text.MN611)
		return
	}
	if p.balance < p.deal.Value {
		p.Outputm(text.InsufficientFunds)
		return
	}

	// Sort out buyer end.
	p.ChangeBalance(-p.deal.Value)

	warehouse := p.storage.FindWarehouse(p.deal.Planet)
	if warehouse == nil {
		p.Outputm(text.MN75)
	} else {
		switch {
		case (p.rank < model.RankSquire || p.rank > model.RankDuke) && p.deal.Pallet.Quantity > 300:
			p.Outputm(text.MN76)
		case !StorePallet(warehouse, p.deal.Pallet):
			p.Outputm(text.MN76)
		default:
			p.Outputm(text.MN612)
		}
	}

	// Sort out vendor end.
	vendor, ok := FindPlayerByID(p.deal.Owner)
	if ok && vendor.IsPlaying() {
		// FIX ME -- guard against overflow.
		vendor.ChangeBalance(p.deal.Value)
		vendor.ChangeBalance(-(p.deal.Value / 20))
		vendor.Outputm(text.MN336, p.name)
		vendor.FlushOutput()
		vendor.Save(database.SaveNow)
	}

	p.deal = nil // FIXME: delete?
}
