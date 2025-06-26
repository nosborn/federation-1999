package server

import (
	"math/rand/v2"

	"github.com/dustin/go-humanize"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

const (
	bribeTechnicianCost = 100000000 // 100,000,000 IG
)

// CmdBribe bribes a galactic official for a ship or planet permit.
func (p *Player) CmdBribe(amount int32) {
	if p.CurSysName() == "Sol" && p.CurLoc().Number() == sol.PublicCounter2 {
		if p.Rank() == model.RankGroundHog {
			p.Outputm(text.BribeWrongLocation_Hint)
		} else {
			p.Outputm(text.BribeWrongLocation)
		}
		return
	}
	if amount <= 0 {
		p.Outputm(text.NiceTry)
		return
	}
	if amount > p.Balance() {
		p.Outputm(text.InsufficientFunds)
		return
	}
	p.ChangeBalance(-amount)
	if amount < 80+rand.Int32N(100) { //nolint:gosec // "It's Just A Game"
		p.Outputm(text.BribeNotEnough)
		return
	}
	switch {
	case p.Rank() == model.RankGroundHog && !p.HasShipPermit():
		// p.Flags1 |= model.PL1_SHIP_PERMIT
		p.SetShipPermit(true)
		p.Outputm(text.BribeOK_ShipPermit)
	case p.Rank() == model.RankExplorer && !p.HasPlanetPermit():
		// p.Flags1 |= model.PL1_PO_PERMIT
		p.SetPlanetPermit(true)
		p.Outputm(text.BribeOK_PlanetPermit)
	default:
		p.Outputm(text.BribeOK)
	}
	p.Save(database.SaveNow)
}

// CmdBribeTechnician allows a player to bribe the technician to zero their
// insurance.
func (p *Player) CmdBribeTechnician() {
	if p.CurSysName() != "Sol" {
		p.Outputm(text.BribeWrongLocation)
		return
	}
	if _, ok := p.CurLoc().FindObject(sol.ObTechnician); !ok {
		p.Outputm(text.BribeTechnicianMobileMissing)
		return
	}
	if p.Balance() < bribeTechnicianCost {
		p.Outputm(text.BribeTechnicianInsufficientFunds, humanize.Comma(bribeTechnicianCost))
		return
	}
	p.ChangeBalance(-bribeTechnicianCost)
	p.SetDeaths(0)
	p.Outputm(text.BribeTechnicianOK, humanize.Comma(bribeTechnicianCost))
	p.Save(database.SaveNow)
}
