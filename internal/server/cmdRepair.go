package server

import (
	"github.com/dustin/go-humanize"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

// Effect repairs to a damaged computer on a player's ship.
func (p *Player) CmdRepairComputer() {
	if !p.repairPreamble() {
		return
	}
	if p.ShipKit.CurComputer >= p.ShipKit.MaxComputer {
		p.repairNotNeeded(text.RepairComputerNoDamage)
		return
	}
	cost := (p.ShipKit.MaxComputer - p.ShipKit.CurComputer) * 50000
	if cost > p.balance && p.HasWallet() {
		p.Outputm(text.MN848, humanize.Comma(int64(cost)))
		return
	}
	p.ShipKit.CurComputer = p.ShipKit.MaxComputer
	p.Outputm(text.MN878, humanize.Comma(int64(cost)))
	p.repairPostamble(cost)
}

// Effect repairs to a damaged engines on a player's ship.
func (p *Player) CmdRepairEngines() {
	if !p.repairPreamble() {
		return
	}
	if p.ShipKit.CurEngine >= p.ShipKit.MaxEngine {
		p.repairNotNeeded(text.RepairEnginesNoDamage)
		return
	}
	cost := (p.ShipKit.MaxEngine - p.ShipKit.CurEngine) * 10000
	if cost > p.balance && p.HasWallet() {
		p.Outputm(text.MN848, humanize.Comma(int64(cost)))
		return
	}
	p.ShipKit.CurEngine = p.ShipKit.MaxEngine
	p.Outputm(text.MN880, humanize.Comma(int64(cost)))
	p.repairPostamble(cost)
}

// Repair armament on an existing spaceship.
func (p *Player) CmdRepairGun(gunType int) {
	if !p.repairPreamble() {
		return
	}
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

// Effect repairs to a damaged hull on a player's ship.
func (p *Player) CmdRepairHull() {
	if !p.repairPreamble() {
		return
	}
	if p.ShipKit.CurHull >= p.ShipKit.MaxHull {
		p.repairNotNeeded(text.RepairHullNoDamage)
		return
	}
	cost := (p.ShipKit.MaxHull - p.ShipKit.CurHull) * 1000
	if cost > p.balance && p.HasWallet() {
		p.Outputm(text.MN848, humanize.Comma(int64(cost)))
		return
	}
	p.ShipKit.CurHull = p.ShipKit.MaxHull
	p.Outputm(text.MN882, humanize.Comma(int64(cost)))
	p.repairPostamble(cost)
}

// Effect repairs to damaged shields on a player's ship.
func (p *Player) CmdRepairShields() {
	if !p.repairPreamble() {
		return
	}
	if p.ShipKit.MaxShield == 0 {
		p.repairNotNeeded(text.RepairShieldsAbsent)
		return
	}
	if p.ShipKit.CurShield >= p.ShipKit.MaxShield {
		p.repairNotNeeded(text.RepairShieldsNoDamage)
		return
	}
	cost := (p.ShipKit.MaxShield - p.ShipKit.CurShield) * 1500
	if cost > p.balance && p.HasWallet() {
		p.Outputm(text.MN848, humanize.Comma(int64(cost)))
		return
	}
	p.ShipKit.CurShield = p.ShipKit.MaxShield
	p.Outputm(text.MN888, humanize.Comma(int64(cost)))
	p.repairPostamble(cost)
}

func (p *Player) repairPostamble(cost int32) {
	if p.HasWallet() {
		p.ChangeBalance(-cost)
		p.CurSys().Income(cost, true)
		p.Save(database.SaveNow)
	}
	p.Count[model.PL_G_MAINT] = 0
	p.quickStatus()
}

func (p *Player) repairNotNeeded(messageNo text.MsgNum) {
	if p.balance > 0 && p.HasWallet() {
		p.balance -= min((500 * int32(p.rank)), p.balance)
		p.Save(database.SaveNow)
	}
	p.Outputm(messageNo)
}

func (p *Player) repairPreamble() bool {
	if !p.curLoc.IsRepairShop() && !p.curLoc.IsShipyard() {
		if !p.HasSpaceship() {
			p.Outputm(text.NoSpaceship)
			return false
		}
		p.Outputm(text.MN1346)
		return false
	}
	if p.CurSys().IsClosed() {
		p.Outputm(text.ClosedForBusiness, p.CurSys().Name())
		return false
	}
	if !p.HasSpaceship() {
		p.repairNotNeeded(text.RepairNoSpaceship)
		return false
	}
	return true
}
