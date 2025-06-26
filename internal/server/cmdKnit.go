package server

func (p *Player) CmdKnitBaron(name string) {
	if !p.CanKnit() {
		p.UnknownCommand()
		return
	}
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

func (p *Player) CmdKnitExplorer(name string, promo bool) {
	if !p.CanKnit() {
		p.UnknownCommand()
		return
	}
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

func (p *Player) CmdKnitNavigator(name string) {
	if !p.CanKnit() {
		p.UnknownCommand()
		return
	}
	p.Output("Not implemented. Check back in 2 weeks.\n")
}
