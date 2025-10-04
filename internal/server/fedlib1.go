package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/sol"
)

func FindSolLandingPad(locNo uint32) uint32 {
	switch {
	case locNo < 70:
		return sol.EarthLandingArea
	case locNo < 92:
		return sol.CallistoLandingPad
	case locNo < 131:
		return sol.TitanLandingArea
	case locNo < 225:
		return sol.SelenaLandingPad
	case locNo < 368:
		return sol.MarsLandingArea
	case locNo < 516:
		return sol.EarthLandingArea
	case locNo < 596:
		return sol.LandingBay3
	case locNo < 672:
		return sol.LandingBay5
	case locNo < 697:
		return sol.EarthLandingArea
	case locNo < 699:
		return sol.MarsLandingArea
	default:
		return sol.EarthLandingArea
	}
}

func shipValue(player *Player) int32 {
	var cost int32

	switch player.ShipKit.Tonnage { // hull
	case 100:
		cost = 50000
	case 200:
		cost = 60000
	case 400:
		cost = 80000
	case 600:
		cost = 150000
	case 800:
		cost = 200000
	case 1000:
		cost = 360000
	case 1400:
		cost = 700000
	}

	cost += cost / 4 // drive

	if player.ShipKit.CurHull > 10 { // plating
		cost += (player.ShipKit.CurHull - 10) * 5000
	}

	cost += player.ShipKit.CurShield * 2500 // shielding

	for i := range model.MAX_GUNS { // weapons
		if player.Guns[i].Type == 0 {
			break
		}
		switch player.Guns[i].Type {
		case 1:
			cost += 2000
		case 2:
			cost += 8000
		case 3:
			cost += 40000
		case 4:
			cost += 75000
		}
	}

	switch player.ShipKit.CurComputer { // computer
	case 1:
		cost += 20000
	case 2:
		cost += 55000
	case 3:
		cost += 80000
	case 4:
		cost += 200000
	case 5:
		cost += 500000
	case 6:
		cost += 1200000
	case 7:
		cost += 2500000
	}

	cost += 10000 + player.ShipKit.CurEngine*10
	return cost
}
