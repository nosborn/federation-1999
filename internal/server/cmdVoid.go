package server

// Allows a player to get out of a job contract, at a price.
//
// Future enhancement: VOIDed jobs should go back onto the job board if at all
// possible.
func (p *Player) CmdVoid() {
	p.Output("Not implemented. Check back in 2 weeks.\n")
}
