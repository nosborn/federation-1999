package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

const (
	bashStaminaCost = 6000000 // 6,000,000 IG
)

func (p *Player) CmdBash() {
	// The player must be carrying the spanner...
	spanner, ok := p.FindInventoryID(sol.ObSpanner)
	if !ok {
		p.Outputm(text.BashNoTool)
		return
	}
	// ...standing in the same location as the slot machine...
	if !p.IsInSolLocation(sol.WaitingRoom) {
		p.Outputm(text.BashWrongLocation)
		return
	}
	if (p.Flags1&model.PL1_DONE_STA) != 0 || p.Sta.Max >= 120 {
		p.Outputm(text.DoneStaminaPuzzle)
		return
	}
	// ...and have 6 MegaGroats in the bank to pay for it.
	if p.Balance() < bashStaminaCost {
		p.Outputm(text.BashInsufficientFunds)
		return
	}
	// Success! Tell them the result of the bashing.
	p.Outputm(text.BashSuccessful)

	// Take away the spanner. If there's a problem with placing it on the
	// recycle list we'll have to nuke their entire inventory and bale out.
	p.RemoveFromInventory(spanner)
	spanner.Recycle()

	// Add 4 stamina points and flag the player as having done the puzzle.
	p.Sta.Max = min(p.Sta.Max+4, 120)
	p.Flags1 |= model.PL1_DONE_STA

	// Finally, take their money.
	p.ChangeBalance(-bashStaminaCost)
	// debug.Check(balance >= 0)

	p.Save(database.SaveNow)
}
