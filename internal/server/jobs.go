package server

import (
	"time"

	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/goods"
	"github.com/nosborn/federation-1999/internal/server/jobs"
	"github.com/nosborn/federation-1999/internal/text"
)

const ( // Where to put the goods
	JOB_EXCHANGE = iota
	JOB_WAREHOUSE
)

const ( // Job status values
	JOB_OFFERED   = iota - 1 // Player offered job not yet signed for
	JOB_NONE                 // No current job
	JOB_ACCEPTED             // Job accepted - but not yet collected
	JOB_COLLECTED            // Cargo collected - but not yet delivered
)

const ( // Transportation requests and reports
	TRANS_NO_SLOT     = iota - 1 // No more slots available - try later
	TRANS_REQ_SLOT               // Request slot for job
	TRANS_REPORT                 // Request report of jobs available
	TRANS_JOB_REFUSED            // Generic job not given
	TRANS_JOB_OK                 // Request completed OK
	TRANS_AGE_JOBS               // Age all the jobs on hand by 1
)

const (
	TRANS_MAX_JOBS    = 100 // Max number of job slots in transportation
	GENERAL_JOB_AGING = 5   // Age increment for general jobs
	JOB_DELIVERY_AGE  = 20  // Age at which jobs expire
)

type JobSlot struct {
	lastused time.Time
	job      model.Work
}

// type PlanetJob struct { // used to store owner generated milkruns
// 	name      string          // planet to deliver to
// 	commodity model.Commodity // type of goods to deliver
// 	carriage  int32           // IG/ton for the hauler
// }

var jobArray [TRANS_MAX_JOBS]JobSlot

// last_job_no int

func AcceptJob(p *Player, jobNo int32) bool {
	job := &jobArray[jobNo].job
	if job.JobType == jobs.JOB_INVALID {
		p.Outputm(text.AcceptAlreadyTaken)
		return false
	}

	p.Count[model.PL_G_JOB] = 0 // zero the jobs counter
	// memcpy(&thisPlayer.pl_job, pJob, sizeof(work_t))

	quantity := job.Pallet.Quantity

	switch job.JobType {
	case jobs.JOB_FACTORY:
		if p.ShipKit.CurHold < quantity {
			p.Outputm(text.AcceptJobTooBig)
			p.ChangeBalance(-150 * int32(p.Rank()))
			// debug.Check(p.pl_job.status == JOB_NONE)
			return false
		}
		job.JobType = jobs.JOB_INVALID

	case jobs.JOB_GENERAL, jobs.JOB_MILKRUN:
		if p.ShipKit.CurHold < quantity {
			p.Outputm(text.AcceptJobTooBig)
			p.ChangeBalance(-150 * int32(p.Rank()))
			// debug.Check(p.pl_job.status == JOB_NONE)
			return false
		}
		job.JobType = jobs.JOB_INVALID

	case jobs.JOB_PLANET:
		if p.ShipKit.CurHold < 1 {
			p.Outputm(text.AcceptFullCargoHold)
			p.ChangeBalance(-150 * int32(p.Rank()))
			// FIXME: This assertion is incorrect...
			// debug.Check(p.Job.status == JOB_NONE)
			return false
		}
		if quantity > p.ShipKit.CurHold {
			quantity = p.ShipKit.CurHold
		}
		p.Job.Pallet.Quantity = quantity
		job.Pallet.Quantity -= quantity
		if job.Pallet.Quantity <= 0 {
			job.JobType = jobs.JOB_INVALID
		}
	}

	debug.Trace("Job accepted for %s", p.Name())
	p.Job.Status = JOB_ACCEPTED
	p.Outputm(text.AcceptBidAccepted)
	return true
}

func BeginJobAdvert() {
	for _, p := range Players {
		if p.Session() == nil {
			continue
		}
		p.flags2 |= PL2_IN_JOB_ADVERT
	}
}

func EndJobAdvert() {
	for _, p := range Players {
		if p.Session() == nil {
			continue
		}
		if (p.flags2 & PL2_IN_JOB_ADVERT) == 0 {
			continue
		}
		p.FlushOutput()
		p.flags2 &^= PL2_IN_JOB_ADVERT
	}
}

func JobReport(job model.Work, jobNo int) string {
	jobType := ' '
	if job.JobType == jobs.JOB_PLANET {
		jobType = '*'
	}
	return text.Msg(text.JobReport,
		jobNo,
		job.Pallet.Quantity,
		goods.GoodsArray[job.Pallet.Type].Name,
		job.From,
		job.To,
		job.Gtu,
		job.Value,
		job.Credits,
		jobType)
}
