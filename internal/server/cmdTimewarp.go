package server

import (
	"log"
	"math/rand/v2"

	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/global"
	"github.com/nosborn/federation-1999/internal/server/horsell"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
	"github.com/nosborn/federation-1999/pkg/ibgames"
)

type TimewarpFuel struct {
	facility  uint32
	commodity model.Commodity
	quantity  int32
}

var timewarpFuel = []TimewarpFuel{
	{model.DK_DOWNSIDE, model.CommodityAntiMatter, 3000},
	{model.DK_ACCEL, model.CommodityRads, 4000},
	{model.DK_MATPROC, model.CommoditySoya, 8000},
	{model.DK_MATPROC, model.CommodityNitros, 5000},
	{model.DK_MATTRANS, model.CommodityMonopoles, 2500},
	{model.DK_MATTRANS, model.CommodityPetros, 3500},
	{model.DK_MATTRANS, model.CommodityXmetals, 1700},
	{model.DK_MATTRANS, model.CommodityCrystals, 2500},
}

// TODO: Fix this up to call doSolCleanups somewhere
func (p *Player) CmdTimewarp() {
	if p.Rank() != model.RankDeity {
		if p.Rank() < model.RankSquire || p.Rank() > model.RankDuke || p.OwnSystem() == nil {
			p.Outputm(text.NOT_A_PLANET_OWNER)
			return
		}
		if p.OwnSystem().IsClosed() {
			p.Outputm(text.MN444)
			return
		}
	}
	if !p.IsInHorsellSystem() && !p.IsInSolSystem() {
		p.Outputm(text.TimewarpOutsideSol)
		return
	}
	if p.IsInsideSpaceship() {
		p.Outputm(text.TimewarpInShip)
		return
	}

	// Figure out the chance of success. This should probably be based on
	// the weakest component instead of the average (nastier that way).
	var chance int
	switch {
	case p.Rank() == model.RankDeity:
		chance = 100
	case p.CurSys().IsHorsell():
		if p.Warper == p.UID() {
			chance = 100
		} else {
			chance = 0
		}
	default:
		for i := range p.Facilities {
			if p.Facilities[i] <= 0 {
				p.Outputm(text.TimewarpFacilitiesIncomplete)
				return
			}
			chance += int(p.Facilities[i])
		}

		chance /= len(p.Facilities)
		debug.Trace("Facilities at %d%%", chance)
		debug.Check(chance >= 0 && chance <= 100)

		if p.storage == nil || p.storage.Warehouse[0] == nil { // FIXME
			p.Outputm(text.NO_MEGA_WAREHOUSE)
			return
		}

		warehouse := p.storage.Warehouse[0]
		debug.Check(warehouse.Planet == p.OwnSystem().Name())

		ownPlanet := p.OwnSystem().Planets()[0]
		debug.Check(ownPlanet != nil)

		// Check for the necessary inputs.
		for i := range timewarpFuel {
			if p.Facilities[timewarpFuel[i].facility] <= 0 {
				continue
			}

			required := timewarpFuel[i].quantity * 2

			// FIXME FIXME FIXME!!
			if global.TestFeaturesEnabled {
				required /= 100
			}

			for nBay := 0; nBay < 20 && required > 0; nBay++ {
				if warehouse.Bay[nBay].Type == timewarpFuel[i].commodity && warehouse.Bay[nBay].Quantity > 0 {
					required -= warehouse.Bay[nBay].Quantity
				}
			}

			// FIX ME - look in the player's exchange as well.
			//
			// This is (supposed to be) a temporary measure until
			// there's a way for a PO to easily accumulate goods in
			// their mega-warehouse.
			if required > 0 && ownPlanet.ActualStock(timewarpFuel[i].commodity) > 0 {
				required -= ownPlanet.ActualStock(timewarpFuel[i].commodity)
			}

			// Did we find enough?
			if required > 0 {
				p.Outputm(text.InsufficientFuelForTimewarp)
				return
			}
		}

		// The inputs are present in the warehouse, so go through again
		// and remove them.
		//
		// FIXME: Only remove half the inputs if the warp fails.
		for i := range timewarpFuel {
			if p.Facilities[timewarpFuel[i].facility] <= 0 {
				continue
			}

			required := timewarpFuel[i].quantity * 2
			if global.TestFeaturesEnabled {
				required /= 100
			}

			for nBay := 0; nBay < 20 && required > 0; nBay++ {
				if warehouse.Bay[nBay].Type == timewarpFuel[i].commodity && warehouse.Bay[nBay].Quantity > 0 {
					if warehouse.Bay[nBay].Quantity >= required {
						warehouse.Bay[nBay].Quantity -= required
						required = 0
					} else {
						required -= warehouse.Bay[nBay].Quantity
						warehouse.Bay[nBay].Quantity = 0
					}
				}
			}

			// FIXME: look in the player's exchange as well.
			if required > 0 && ownPlanet.ActualStock(timewarpFuel[i].commodity) > 0 {
				debug.Check(ownPlanet.ActualStock(timewarpFuel[i].commodity) >= required)
				ownPlanet.RemoveGoods(timewarpFuel[i].commodity, required)
			}
		}
	}

	//
	if chance < rand.IntN(100)+1 {
		if p.IsInHorsellSystem() {
			p.CurSys().Talk(text.Msg(text.TimewarpFailed))
		} else {
			p.curLoc.Talk(text.Msg(text.TimewarpFailed))
		}
		p.Save(database.SaveNow)
		return
	}
	if global.TestFeaturesEnabled {
		switch p.UID() {
		case 100000: // Moi!
		case 102400: // Shaunamari (live persona)
			break
		default:
			p.Outputm(text.TimewarpFailed)
			return
		}
	}

	if !p.IsInHorsellSystem() {
		debug.Trace("cmdTimewarp: Moving to Horsell")
		if !timewarpToHorsell(p) {
			p.Outputm(text.TimewarpFailed)
		}
	} else {
		debug.Trace("cmdTimewarp: Moving from Horsell")
		timewarpFromHorsell(p)
	}

	//
	damaged := false
	if p.Rank() <= model.RankDuke {
		maxDamage := 12
		if chance == 100 {
			maxDamage = 7
		}
		for i := model.DK_UPSIDE; i <= model.DK_TIMEMACH; i++ {
			if p.Facilities[i] > 0 {
				if damage := rand.IntN(maxDamage + 1); damage > 0 {
					p.Facilities[i] -= int32(damage)
					if p.Facilities[i] < 1 {
						p.Facilities[i] = 1
					}
					damaged = true
				}
			}
		}
	}
	if damaged {
		p.Outputm(text.TimewarpDamage)
	} else {
		p.Outputm(text.TimewarpNoDamage)
	}
}

func beginTimewarp(player *Player, toSystem Systemer, toLocNo, shipLocNo uint32) {
	debug.Precondition(toSystem != nil)

	// FIXME: TEMP: turn on logging.
	if toSystem.IsHorsell() {
		player.Session().Logging(true)
	}

	//
	if (player.flags2 & PL2_CORPSE) == 0 {
		player.Outputm(text.TimewarpOK)
		player.FlushOutput()
	}

	player.SetCurSys(toSystem)
	player.LocNo = toLocNo
	player.setLocation(toLocNo)

	if player.HasSpaceship() {
		player.ShipLoc = shipLocNo
	}

	player.flags2 |= PL2_TIMEWARPED

	player.CheckSpyers()
}

func EndTimewarp(player *Player, warper ibgames.AccountID) {
	// dbgPrecondition(player == 0 || player >= ibgames.MinAccountID);

	if (player.flags2 & PL2_CORPSE) == 0 {
		player.curLoc.Describe(player, LongDescription)
		player.FlushOutput()
	}

	player.flags2 &^= PL2_TIMEWARPED
	player.Warper = warper

	if (player.flags2 & PL2_CORPSE) == 0 {
		player.Save(database.SaveNow)
	}

	// // Remove any Horsell objects.
	// player.inventory.erase(remove_if(m_inventory.begin(),
	//                             m_inventory.end(),
	//                             isHorsellObject()),
	//                   m_inventory.end());

	// FIXME: TEMP: turn off logging.
	if (player.Rank() < model.RankSenator || player.Rank() == model.RankDeity) && !player.IsPromoCharacter() {
		if player.curSys.IsSol() {
			player.Session().Logging(global.TestFeaturesEnabled)
		}
	}
}

func timewarpFromHorsell(caller *Player) {
	horsellSystem := caller.CurSys()
	horsellDuchy := horsellSystem.Duchy()

	for _, warpee := range Players {
		if !warpee.IsPlaying() {
			continue
		}
		if warpee.curSys != horsellSystem {
			continue
		}
		beginTimewarp(warpee, SolSystem, sol.EarthLandingArea, sol.EarthLandingArea)
	}

	caller.curLoc.EndTimewarp(0)

	// Destroy this Horsell.
	// horsellSystem.(*HorsellSystem).Destroy() -- FIXME: call "destructor"
	horsellDuchy.Destroy()
}

func timewarpToHorsell(caller *Player) bool {
	debug.Check(caller.curSys.IsSol())

	// Create a new Horsell.
	horsellSystem := createHorsell(caller)
	if horsellSystem == nil {
		caller.Output("createHorsell() failed!\n") // FIXME
		return false
	}

	// o/~ Let's do the timewarp again... o/~ :)
	fromLoc := caller.curLoc
	for _, warpee := range Players {
		if !warpee.IsPlaying() {
			continue
		}
		if warpee.curLoc != fromLoc {
			continue
		}
		if warpee.Rank() < model.RankAdventurer {
			continue
		}
		if warpee.IsOnDutyNavigator() {
			continue
		}
		beginTimewarp(warpee, horsellSystem, horsell.Downs1, horsell.GoodiesStore)
	}

	// Tell anyone who was left behind.
	fromLoc.Talk(text.Msg(text.TimewarpFailed))

	// Finish up the trip for those who went.
	caller.curLoc.EndTimewarp(caller.UID())

	// Start the puzzle.
	debug.Check(caller.curSys == horsellSystem)
	if caller.curSys != horsellSystem {
		log.Panic("cmdTimewarp: caller not moved to Horsell")
	}
	horsellSystem.StartPuzzle()

	return true
}

func createHorsell(caller *Player) *HorsellSystem {
	duchy := NewHorsellDuchy("Horsell0") // FIXME
	return NewHorsellSystem(duchy, caller)
}
