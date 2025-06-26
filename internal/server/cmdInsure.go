package server

import (
	"log"

	"github.com/dustin/go-humanize"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

// Allows player to insure against being killed in the game. The function also
// tells how much insurance costs.
func (p *Player) CmdInsure(action model.InsureAction) {
	// Check player is in insurance brokers!
	if !p.CurLoc().IsInsuranceBroker() {
		if !p.IsInsured() {
			locNo := p.CurLocNo()
			if p.IsFlyingSpaceship() {
				locNo = p.ShipLocNo()
			}
			log.Printf("%s is looking for insurance (%s/%d)", p.Name(), p.CurSysName(), locNo)
		}
		p.Outputm(text.InsureWrongLocation)
		return
	}

	// Are they already insured?
	if p.IsInsured() {
		log.Printf("%s is already insured (%s/%d)", p.Name(), p.CurSysName(), p.CurLocNo())
		p.Outputm(text.InsureAlreadyInsured)
		return
	}

	// Calculate cost of insurance.
	var cost int32
	if p.Deaths() < 10 {
		cost = 1000 + (3000 * p.Deaths())
	} else {
		cost = 10000 + (50000 * p.Deaths())
	}

	// Does the player just want to know how much it will cost?
	if action == model.InsureGetQuote {
		if !p.IsInsured() {
			log.Printf("%s got insurance quotation (%s/%d)", p.Name(), p.CurSysName(), p.CurLocNo())
		}
		p.Outputm(text.InsureQuotation, humanize.Comma(int64(cost)))
		return
	}

	// debug.Check(action == insureBuyPolicy);

	// Are they already insured?
	if p.IsInsured() {
		p.Outputm(text.InsureAlreadyInsured)
		return
	}

	// Can they afford to get insured?
	if cost > p.Balance() {
		log.Printf("%s can't afford insurance (%s/%d)", p.Name(), p.CurSysName(), p.CurLocNo())
		p.Outputm(text.InsureInsufficientFunds, humanize.Comma(int64(cost)))
		return
	}

	// All OK - issue them with a policy...
	// p.Flags0 |= model.PL0_INSURED
	p.SetInsured(true)
	p.ChangeBalance(-cost)
	p.CurSys().Income(cost, true)
	p.Outputm(text.InsureOK)
	p.Save(database.SaveNow)

	// ...and log the fact.
	log.Printf("%s has purchased insurance (%s/%d)", p.Name(), p.CurSysName(), p.CurLocNo())
}
