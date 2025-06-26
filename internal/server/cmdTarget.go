package server

// Allocates the specified player as the preferred target for ship weapons.
func (p *Player) CmdTarget(name string) {
	if name == "" {
		p.Output("Not implemented. Check back in 2 weeks.\n")
		return
	}

	p.Output("Not implemented. Check back in 2 weeks.\n")
}

func (p *Player) CmdTargetOff() {
	p.Output("Not implemented. Check back in 2 weeks.\n")
}
