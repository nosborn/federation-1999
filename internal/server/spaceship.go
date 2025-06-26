package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/text"
)

const (
	SPACESHIP_SIZE = 8
)

func GetShipClass(tonnage int32) string {
	switch {
	case tonnage < 200:
		return "Fleet"
	case tonnage < 400:
		return "Harrier"
	case tonnage < 600:
		return "Mesa"
	case tonnage < 800:
		return "Dragon"
	case tonnage < 1000:
		return "Guardian"
	case tonnage < 1400:
		return "Mammoth"
	default:
		return "Imperial"
	}
}

// Reduces a spaceships stats so that it needs servicing, and warns the player.
func ServiceCall(p *Player) {
	if p.Rank() < model.RankTrader {
		return
	}

	// Reduce effectiveness of kit.
	if p.ShipKit.CurShield -= 1; p.ShipKit.CurShield < 0 {
		p.ShipKit.CurShield = 0
	}
	if p.ShipKit.CurEngine -= 10; p.ShipKit.CurEngine < 10 {
		p.ShipKit.CurEngine = 10
	}
	p.Outputm(text.MN893)
}
