package server

import (
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

// CmdDividend allows player to pay out a company dividend.
func (p *Player) CmdDividend(number int32) {
	if p.company == nil {
		p.Outputm(text.MN1066)
		return
	}
	if number <= 0 {
		p.Outputm(text.MN1067)
		return
	}
	if number > p.company.Balance {
		p.Outputm(text.MN1068, p.company.Name)
		return
	}
	p.company.ChangeBalance(-number)
	p.ChangeBalance(number)
	p.Outputm(text.MN1069)
	p.Save(database.SaveNow)
}
