package server

import (
	"log"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdBlessHostess(name string) {
	if p.rank != model.RankDeity {
		p.UnknownCommand()
		return
	}

	subject, ok := FindPlayer(name)
	if !ok {
		p.Outputm(text.MN20)
		return
	}
	if subject.rank == model.RankGroundHog || subject.rank > model.RankExplorer {
		p.Outputm(text.DontBeSilly)
		return
	}
	if subject.IsNavigator() || subject.IsPromoCharacter() {
		p.Outputm(text.DontBeSilly)
		return
	}
	if subject.company != nil || subject.storage != nil {
		p.Outputm(text.BlessHasAssets, subject.Name())
		return
	}

	subject.rank = model.RankHostess
	subject.Str = PlayerStat{Max: 128, Cur: 120}
	subject.Sta = PlayerStat{Max: 128, Cur: 120}
	subject.Int = PlayerStat{Max: 128, Cur: 120}
	subject.Dex = PlayerStat{Max: 128, Cur: 120}
	subject.Shipped = 0
	subject.Flags0 &^= model.PL0_JOB | model.PL0_OFFER_TOUR | model.PL0_KNOWS_ORSONITE | model.PL0_SNARK_ASSIGNED | model.PL0_HORSELL_ASSIGNED
	subject.Flags0 |= model.PL0_COMM_UNIT | model.PL0_LIT | model.PL0_INSURED | model.PL0_SPYBEAM | model.PL0_SPYSCREEN
	subject.Flags1 &^= model.PL1_MI6_OFFERED | model.PL1_HILBERT | model.PL1_PO_PERMIT | model.PL1_NAVIGATOR | model.PL1_PROMO_CHAR
	subject.Flags1 |= model.PL1_SHIP_PERMIT | model.PL1_MI6 | model.PL1_DONE_STA | model.PL1_DONE_STR | model.PL1_DONE_INT | model.PL1_DONE_SNARK
	subject.flags2 &^= PL2_ON_DUTY_NAV

	subject.balance = 100000000 // 100,000,000 IG (!)
	subject.loan = 0
	subject.reward = 0
	subject.deaths = 0
	subject.TradeCredits = 0
	subject.gmLocation = 0

	// TODO:
	// memset(&theSubject->pl_job, '\0', sizeof(theSubject->pl_job));
	subject.Job.Status = JOB_NONE

	subject.Save(database.SaveNow)

	p.outputf("OK - %s is now a DataSpace Hostess!\n", subject.Name())
	log.Printf("{H} %s has been blessed by %s", subject.Name(), p.Name())

	// Start session logging for the subject.
	if subject.IsPlaying() {
		subject.session.Logging(true)
	}
}

func (p *Player) CmdBlessSenator(name string) {
	if p.rank != model.RankDeity {
		p.UnknownCommand()
		return
	}

	subject, ok := FindPlayer(name)
	if !ok {
		p.Outputm(text.MN20)
		return
	}
	if subject.rank == model.RankGroundHog {
		p.Outputm(text.DontBeSilly)
		return
	}
	if subject.IsNavigator() || subject.IsPromoCharacter() {
		p.Outputm(text.DontBeSilly)
		return
	}
	if subject.company != nil || subject.storage != nil {
		p.Outputm(text.BlessHasAssets, subject.Name())
		return
	}

	subject.rank = model.RankSenator
	subject.Str = PlayerStat{Max: 128, Cur: 120}
	subject.Sta = PlayerStat{Max: 128, Cur: 120}
	subject.Int = PlayerStat{Max: 128, Cur: 120}
	subject.Dex = PlayerStat{Max: 128, Cur: 120}
	subject.Shipped = 0
	subject.Flags0 &^= model.PL0_JOB | model.PL0_OFFER_TOUR | model.PL0_SNARK_ASSIGNED | model.PL0_HORSELL_ASSIGNED
	subject.Flags0 |= model.PL0_COMM_UNIT | model.PL0_LIT | model.PL0_INSURED | model.PL0_SPYBEAM | model.PL0_SPYSCREEN
	subject.Flags1 &^= model.PL1_PO_PERMIT | model.PL1_MI6_OFFERED | model.PL1_MI6 | model.PL1_HILBERT | model.PL1_PO_PERMIT | model.PL1_NAVIGATOR | model.PL1_PROMO_CHAR
	subject.Flags1 |= model.PL1_SHIP_PERMIT | model.PL1_DONE_STA | model.PL1_DONE_STR | model.PL1_DONE_INT
	subject.flags2 &^= PL2_ON_DUTY_NAV

	subject.balance = 0
	subject.loan = 0
	subject.reward = 0
	subject.deaths = 0
	subject.TradeCredits = 0
	subject.gmLocation = 0

	// FIXME:
	// memset(&theSubject->pl_job, '\0', sizeof(theSubject->pl_job));
	subject.Job.Status = JOB_NONE

	subject.Save(database.SaveNow)

	p.outputf("OK - %s is now a Senator!\n", subject.Name())
	log.Printf("{H} %s has been blessed by %s", subject.Name(), p.Name())

	// Start session logging for the subject.
	if subject.IsPlaying() {
		subject.session.Logging(true)
	}
}
