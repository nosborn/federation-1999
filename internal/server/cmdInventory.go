package server

import "github.com/nosborn/federation-1999/internal/text"

// Gets a complete list of a player's personal kit and of all the objects s/he
// is carrying.
func (p *Player) CmdInventory() {
	var items []string

	if p.HasCommUnit() {
		items = append(items, "a comm unit")
	}
	if p.HasLamp() {
		items = append(items, "a lamp")
	}
	if p.HasShipPermit() {
		items = append(items, "a ship-owner's permit")
	}
	if p.HasTradingPermit() {
		items = append(items, "a trading permit")
	}
	if p.HasPlanetPermit() {
		items = append(items, "a planet-owner's permit")
	}
	if p.HasSpyBeam() {
		items = append(items, "a spybeam receiver")
	}
	if p.HasSpyScreen() {
		items = append(items, "a spybeam screen")
	}
	if p.HasIDCard() {
		items = append(items, "an ID card")
	}
	if p.Games >= 1000 {
		items = append(items, "a long service medal")
	}
	items = append(items, "a vac suit")

	p.Output("Your personal kit includes ")
	p.Output(text.ListOfObjects(items))

	if len(p.inventory) > 0 {
		var items []string
		for _, o := range p.inventory {
			items = append(items, o.DisplayName(false))
		}
		p.Output(" Apart from that you are carrying ")
		p.Output(text.ListOfObjects(items))
	}

	p.Output("\n")
}
