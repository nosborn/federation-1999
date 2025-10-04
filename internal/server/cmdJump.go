package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

// Allows player to make hyperspace jumps between star systems.
func (p *Player) CmdJump(destName string) {
	if p.curLoc.IsLink() {
		p.Outputm(text.JumpNotAtLink)
		return
	}

	if p.Rank() < model.RankCaptain {
		p.Outputm(text.JumpNotAllowed)
		return
	}

	dest, ok := FindSystem(destName)
	if !ok || dest.IsHidden() {
		p.Outputm(text.NoSuchSystem)
		return
	}
	if dest == p.CurSys() {
		p.outputf("You're already in %s!\n", dest.Name())
		return
	}

	var toDuchy *Duchy
	if dest.Duchy() != p.CurSys().Duchy() {
		if !p.CurSys().IsCapital() || !dest.IsCapital() {
			p.Output("I don't know how to get there from here!\n")
			return
		}
		if p.CurSys().Duchy().Embargo() == dest.Duchy() || dest.Duchy().Embargo() == p.CurSys().Duchy() {
			p.Outputm(text.LINK_CLOSED, dest.Name())
			return
		}
		toDuchy = dest.Duchy()
	}

	if dest.IsClosed() {
		p.Outputm(text.LINK_CLOSED, dest.Name())
		return
	}

	// Check for embargoed goods and calculate any import duties.
	importDuty := int32(0)
	if toDuchy != nil {
		// FIXME: check for embargoed job
		for i := range p.Load {
			if p.Load[i].Quantity == 0 {
				continue
			}
			origin, ok := FindSystem(p.Load[i].Origin)
			if ok && origin.Duchy() == toDuchy.Embargo() { //nolint:staticcheck // SA9003: empty branch
				// FIXME:
				// output(mn?);
				// return;
			}
			importDuty += toDuchy.ImportDuty(p.Load[i])
		}
		// debug.Check(importDuty >= 0);

		// Can the player afford it?
		if importDuty > 0 && p.Balance() < importDuty {
			p.Outputm(text.MN1126, importDuty)
			return
		}
	}

	// Deal with on-duty Navigators leaving Sol.
	if p.IsOnDutyNavigator() && !dest.IsSol() {
		p.StopNavigating()
		p.FlushOutput()
	}

	// We don't want objects moving between star systems.
	if len(p.inventory) > 0 {
		p.Outputm(text.MN103)
		p.clearInventory()
	}

	// Take the money for the import duty.
	if importDuty > 0 {
		p.ChangeBalance(-importDuty)
		dest.Income(importDuty, true)
		p.Outputm(text.MN1125, importDuty)
	}

	// Sanity check!
	// debug.Check(IsFlyingSpaceship());
	// debug.Check(IsInsideSpaceship());

	// Notify any other players at the departure point.
	msg := text.Msg(text.ShipJumpsFrom, p.Name, GetShipClass(p.ShipKit.Tonnage))
	p.curLoc.Talk(msg, p)

	// Move the player.
	p.SetCurSys(dest)
	p.ShipLoc = uint32(dest.LinkLocNo())
	p.setLocation(p.ShipLoc)
	p.Outputm(text.MN48)
	p.curLoc.Describe(p, DefaultDescription)
	p.Save(database.SaveNow)

	// Notify any other players at the arrival point.
	msg = text.Msg(text.ShipJumpsTo, p.Name, GetShipClass(p.ShipKit.Tonnage))
	p.curLoc.Talk(msg, p)

	// TODO:
	// // Mark the star system as populated.
	// currentSystem()->m_populated = Transaction::time();

	// Restart the player's tourism timer.
	p.RestartTourismTimer()
}
