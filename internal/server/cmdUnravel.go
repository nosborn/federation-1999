package server

func (p *Player) CmdUnravelNavigator(name string) {
	if !p.CanKnit() {
		p.UnknownCommand()
		return
	}
	p.Output("Not implemented. Check back in 2 weeks.\n")
}
