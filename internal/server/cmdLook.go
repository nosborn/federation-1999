package server

func (p *Player) CmdLook() {
	p.CurLoc().Describe(p, FullDescription)
}
