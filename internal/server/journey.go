package server

import (
	"github.com/nosborn/federation-1999/internal/text"
)

// Event 20 - checks to see whether player is allowed to enter
// Naval  Intelligence HQ.
func EnterMI6(p *Player) bool {
	p.Outputm(text.MN1005)
	if p.HasIDCard() {
		p.Outputm(text.MN1006)
		return false
	}
	p.Outputm(text.MN194)
	p.LocNo = 674
	p.setLocation(p.LocNo)
	return true
}
