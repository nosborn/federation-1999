package server

import "github.com/nosborn/federation-1999/internal/model"

// CmdDeallocate deallocates the discretionary production from a specified
// commodity.
func (p *Player) CmdDeallocate(commodity model.Commodity) {
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

// CmdDeallocateSocialSecurity allows a planet owner to deallocate points in
// the social security categories.
func (p *Player) CmdDeallocateSocialSecurity(points int32) {
	p.Output("Not implemented. Check back in 2 weeks.\n")
}
