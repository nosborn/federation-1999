package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/text"
)

// CmdAcceptFactory allows a planet owner to agree to the establishment of a
// factory on his/her planet. This is all a bit dodgy because both factory
// owner and planet owner share the zone_req pointer.
func (p *Player) CmdAcceptFactory() {
	if p.zoneReq == nil {
		p.Outputm(text.MN1021)
		return
	}

	// TODO:
	p.Output("Not implemented. Check back in 2 weeks.\n")
}

func (p *Player) CmdAcceptJob(jobNo int32) {
	if jobNo < 1 || jobNo > TRANS_MAX_JOBS {
		p.Outputm(text.MN991)
		return
	}
	if !p.HasSpaceship() {
		if p.Rank() == model.RankGroundHog {
			p.Outputm(text.AcceptNoSpaceship)
		} else {
			p.Outputm(text.NoSpaceship)
		}
		return
	}
	if p.Job.Status != JOB_NONE {
		if p.Job.Status == JOB_OFFERED {
			p.Outputm(text.AcceptOutstandingOffer)
		} else {
			p.Outputm(text.AcceptSecondJob)
		}
		return
	}
	if p.Rank() >= model.RankTrader && p.Rank() <= model.RankExplorer {
		p.ChangeBalance(int32(-150 * int(p.Rank())))
		p.Outputm(text.AcceptBidDeclined)
		return
	}
	if p.IsOnDutyNavigator() || p.IsPromoCharacter() {
		p.ChangeBalance(int32(-150 * int(p.Rank())))
		p.Outputm(text.AcceptBidDeclined)
		return
	}
	p.Outputm(text.AcceptApplying)
	if !AcceptJob(p, jobNo) {
		return
	}
	p.Save(database.SaveNow)
}
