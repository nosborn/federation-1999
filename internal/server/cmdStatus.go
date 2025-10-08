package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/text"
)

// Displays the status of the player's ship.
func (p *Player) CmdStatus() {
	if !p.HasSpaceship() {
		p.Outputm(text.NoSpaceship)
		return
	}

	p.Output("Status report for your spaceship:\n")
	p.outputf("  Hull strength: %4d/%-4d\n", p.ShipKit.CurHull, p.ShipKit.MaxHull)

	p.outputf("  Shields:       %4d/%-4d", p.ShipKit.CurShield, p.ShipKit.MaxShield)
	if p.ShipKit.CurShield > 0 {
		if (p.Flags1 & model.PL1_SHIELDS) == 0 {
			p.Output(" - off")
		} else {
			p.Output(" - on")
		}
	}
	p.Output("\n")

	p.outputf("  Engines:       %4d/%-4d", p.ShipKit.CurEngine, p.ShipKit.MaxEngine)
	if (p.Flags1 & model.PL1_HILBERT) != 0 {
		p.Output(" - n-space capable")
	}
	p.Output("\n")

	p.outputf("  Computer:      %4d/%-4d", p.ShipKit.CurComputer, p.ShipKit.MaxComputer)
	if p.ShipKit.CurComputer > 0 {
		if (p.Flags1 & model.PL1_AUTO) == 0 {
			p.Output(" - manual")
		} else {
			p.Output(" - automatic")
		}
	}
	p.Output("\n")

	p.outputf("  Cargo space:   %4d/%-4d\n", p.ShipKit.CurHold, p.ShipKit.MaxHold)
	p.outputf("  Fuel:          %4d/%-4d\n", p.ShipKit.CurFuel, p.ShipKit.MaxFuel)
	p.Output("  Weaponry carried:\n")

	noWeaponry := true

	// for (size_t slot = 0; slot < MAX_GUNS; ++slot) {
	// 	const s_guns_t& gun = guns[slot];
	// 	const arms_t* weapon = getWeapon(gun.type);
	//
	// 	if (weapon != 0) {
	// 		output("   %-14s   %2d/%-2d\n",
	// 		message(weapon->messageNo),
	// 		gun.damage,
	// 		weapon->gun.damage);
	// 		noWeaponry = false;
	// 	}
	// }

	if p.Missiles > 0 {
		p.outputf("   %d missiles available\n", p.Missiles)
		noWeaponry = false
	}

	if p.Ammo > 0 {
		p.outputf("   %d mag gun shots available\n", p.Ammo)
		noWeaponry = false
	}

	if noWeaponry {
		p.Outputm(text.MN633)
	}

	if p.Target != 0 { //nolint:staticcheck // SA9003: empty branch
		// const Player* targetPlayer = Player::find(m_target);
		//
		// if (targetPlayer != 0 && targetPlayer->IsPlaying()) {
		// 	output("Computer auto-target is: %s\n", targetPlayer->name());
		// } else {
		// 	m_target = 0;
		// }
	}

	// if (pl_job.status != JOB_NONE) {
	// 	displayJob();
	// }
}
