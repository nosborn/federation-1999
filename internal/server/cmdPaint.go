package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

const (
	paintSpaceshipCost = 750 // 750 IG
)

// Enters new text into a player's ship_desc field.
func (p *Player) CmdPaintSpaceship(description string) {
	if !p.HasSpaceship() {
		p.Outputm(text.NoSpaceship)
		return
	}
	if p.HasWallet() && p.Balance() < paintSpaceshipCost {
		p.Outputm(text.MN783, paintSpaceshipCost)
		p.Outputm(text.InsufficientFunds)
		return
	}
	if len(description) >= model.SHIP_DESC_SIZE {
		p.Outputm(text.DescriptionTooLong, model.SHIP_DESC_SIZE-1)
		return
	}
	if description[0] == '/' || description[0] == '>' {
		p.Outputm(text.DescriptionBadLeader)
		return
	}
	p.SetShipDesc(description)
	if p.HasWallet() {
		p.ChangeBalance(-paintSpaceshipCost)
	}
	p.Outputm(text.PaintSpaceshipOK, p.ShipDesc())
	p.Save(database.SaveNow)
}
