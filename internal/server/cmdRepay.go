package server

import (
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

// CmdRepay enables player to transfer money out of his/her account to repay
// ship purchase loan - checks for promotion to Captain.
func (p *Player) CmdRepay(amount int32) {
	if p.Loan() <= 0 {
		p.Outputm(text.MN793)
		return
	}

	if amount <= 0 /*|| amount >= math.MaxInt32*/ {
		p.Outputm(text.MN792)
		return
	}

	if amount > p.Balance() {
		p.Outputm(text.InsufficientFunds)
		return
	}

	// FIXME: Make smartass comments if the loan is overpayed.
	p.RepayLoan(amount)
	p.ChangeBalance(-amount)
	p.Outputm(text.MN794)
	p.CheckForPromotion()
	p.Save(database.SaveNow)
}
