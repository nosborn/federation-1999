package server

import "github.com/nosborn/federation-1999/internal/model"

// Allows the player to put an object into the Snark Workshop jig. Whatever
// goes into it will be destroyed, along with a possible dex hit if it's not
// the tuning fork. If it is the tuning fork, the player comes away with the
// n-space converter.
func (p *Player) CmdPutIntoJig(name model.Name) {
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

func (p *Player) CmdPutIntoSlot(objectName model.Name) {
	p.Output("Not implemented. Check back in 2 weeks.\n")
}
