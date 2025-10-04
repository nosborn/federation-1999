package server

import (
	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

var routes = [7][6][]uint32{
	{ // Titan...
		{sol.TitanOrbit, 13, sol.CallistoOrbit},                    // ...to Callisto
		{sol.TitanOrbit, 20, 25, sol.MarsOrbit},                    // ...to Mars
		{sol.TitanOrbit, 20, 25, 26, 32, sol.EarthOrbit},           // ...to Earth
		{sol.TitanOrbit, 20, 25, 26, 32, 33, 40, sol.LunarOrbit},   // ...to Moon
		{sol.TitanOrbit, 19, 23, 28, 35, 42, sol.VenusOrbit},       // ...to Venus
		{sol.TitanOrbit, 20, 25, 30, 36, 44, 52, sol.MercuryOrbit}, // ...to.Nercury
	},
	{ // Callisto...
		{sol.CallistoOrbit, 13, sol.TitanOrbit},                   // ...to Titan
		{sol.CallistoOrbit, 25, sol.MarsOrbit},                    // ...to Mars
		{sol.CallistoOrbit, 26, 32, sol.EarthOrbit},               // ...to Earth
		{sol.CallistoOrbit, 26, 32, 33, 40, sol.LunarOrbit},       // ...to Moon
		{sol.CallistoOrbit, 25, 30, 36, 43, sol.VenusOrbit},       // ...to Venus
		{sol.CallistoOrbit, 25, 30, 36, 44, 52, sol.MercuryOrbit}, // ...to Mercury
	},
	{ // Mars...
		{sol.MarsOrbit, 25, 20, sol.TitanOrbit},       // ...to Titan
		{sol.MarsOrbit, 25, sol.CallistoOrbit},        // ...to Callisto
		{sol.MarsOrbit, 32, sol.EarthOrbit},           // ...to Earth
		{sol.MarsOrbit, 32, 33, 40, sol.LunarOrbit},   // ...to Moon
		{sol.MarsOrbit, 36, 43, sol.VenusOrbit},       // ...to Venus
		{sol.MarsOrbit, 38, 46, 53, sol.MercuryOrbit}, // ...to Mercury
	},
	{ // Earth...
		{sol.EarthOrbit, 32, 26, 25, 20, sol.TitanOrbit}, // ...to Titan
		{sol.EarthOrbit, 32, 26, sol.CallistoOrbit},      // ...to Callisto
		{sol.EarthOrbit, 32, sol.MarsOrbit},              // ...to Mars
		{sol.EarthOrbit, sol.LunarOrbit},                 // ...to Moon
		{sol.EarthOrbit, 45, 51, sol.VenusOrbit},         // ...to Venus
		{sol.EarthOrbit, 46, 53, sol.MercuryOrbit},       // ...to Mercury
	},
	{ // Moon...
		{sol.LunarOrbit, 40, 33, 32, 26, 25, 20, sol.TitanOrbit}, // ...to Titan
		{sol.LunarOrbit, 40, 33, 32, 26, sol.CallistoOrbit},      // ...to Callisto
		{sol.LunarOrbit, 40, 33, 32, sol.MarsOrbit},              // ...to Mars
		{sol.LunarOrbit, sol.EarthOrbit},                         // ...to Earth
		{sol.LunarOrbit, 53, 52, 51, sol.VenusOrbit},             // ...to Venus
		{sol.LunarOrbit, 53, sol.MercuryOrbit},                   // ...to Mercury
	},
	{ // Venus...
		{sol.VenusOrbit, 42, 35, 28, 23, 19, sol.TitanOrbit}, // ...to Titan
		{sol.VenusOrbit, 43, 36, 30, 25, sol.CallistoOrbit},  // ...to Callisto
		{sol.VenusOrbit, 43, 36, sol.MarsOrbit},              // ...to Mars
		{sol.VenusOrbit, 51, 45, sol.EarthOrbit},             // ...to Earth
		{sol.VenusOrbit, 51, 52, 53, sol.VenusOrbit},         // ...to Venus -- FIXME: ...to Moon
		{sol.VenusOrbit, 51, 52, sol.MercuryOrbit},           // ...to Mercury
	},
	{ // Mercury...
		{sol.MercuryOrbit, 52, 44, 36, 30, 25, 20, sol.TitanOrbit}, // ...to Titan
		{sol.MercuryOrbit, 52, 44, 36, 30, 25, sol.CallistoOrbit},  // ...to Callisto
		{sol.MercuryOrbit, 53, 46, 38, sol.MarsOrbit},              // ...to Mars
		{sol.MercuryOrbit, 53, 46, sol.EarthOrbit},                 // ...to Earth
		{sol.MercuryOrbit, 53, sol.LunarOrbit},                     // ...to Moon
		{sol.MercuryOrbit, 52, 51, sol.VenusOrbit},                 // ...to Venus
	},
}

// Allows a player to go directly to the planet if the computer knows the way!
func (p *Player) CmdGoto(where string) {
	// Are we flying the ship?
	if !p.IsFlyingSpaceship() {
		p.Outputm(text.MN549)
		return
	}
	// Are we in the Solar system?
	if !p.IsInSolSystem() {
		p.Outputm(text.GotoOutsideSol)
		return
	}
	destination, ok := FindPlanet(where)
	if !ok {
		p.Outputm(text.GotoBadDestination)
		return
	}
	if destination.System() == p.CurSys() {
		if destination.orbitLocNo == p.ShipLoc {
			p.Outputm(text.GotoHere)
			return
		}
	}
	// Does the computer know how to get there?
	if (p.Flags1 & destination.RouteFlag()) == 0 {
		p.Outputm(text.GotoNotProgrammed)
		return
	}
	// We're not going anywhere without fuel!
	if p.ShipKit.CurFuel == 0 {
		p.Outputm(text.MN551)
		return
	}
	//
	route, ok := findRoute(p.ShipLoc, destination.orbitLocNo)
	if !ok {
		p.Outputm(text.MN549)
		return
	}
	// FIXME: Put back random failure.
	msg := text.Msg(text.ShipLeaves, p.Name, GetShipClass(p.ShipKit.Tonnage))
	p.curLoc.Talk(msg, p)

	// Let's go!
	for i := range route {
		p.Count[model.PL_G_JOB]++
		if p.Count[model.PL_G_MAINT]++; p.Count[model.PL_G_MAINT]%350 == 349 {
			ServiceCall(p)
		}
		p.ShipLoc = route[i]
		p.ShipKit.CurFuel -= (1 + p.ShipKit.MaxEngine/60)
		p.CheckShields()
		p.setLocation(p.ShipLoc)
		p.curLoc.Describe(p, BriefDescription)
		if p.ShipKit.CurFuel < 0 { // FIXME: should be <= ?
			p.Outputm(text.MN551)
			p.ShipKit.CurFuel = 0
			break
		}
	}
	p.Save(database.SaveNow)
	// Tell any spectators.
	msg = text.Msg(text.ShipArrives, p.Name, GetShipClass(p.ShipKit.Tonnage))
	p.curLoc.Talk(msg, p)
}

func findRoute(from, to uint32) ([]uint32, bool) {
	for i := range routes {
		if routes[i][0][0] != from {
			continue
		}
		for j := range len(routes[i]) {
			debug.Check(routes[i][0][0] == from)
			if routes[i][j][len(routes[i][j])-1] != to {
				continue
			}
			return routes[i][j], true
		}
	}
	return []uint32{}, false
}
