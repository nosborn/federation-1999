package server

func (p *Player) CmdGlance() {
	p.CurLoc().Describe(p, BriefDescription)
}
