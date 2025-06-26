package server

import (
	"log"
	"math/rand/v2"
	"time"

	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/monitoring"
	"github.com/nosborn/federation-1999/internal/server/global"
	"github.com/nosborn/federation-1999/internal/server/goods"
	"github.com/nosborn/federation-1999/internal/server/jobs"
	"github.com/nosborn/federation-1999/internal/text"
)

const (
	transportationTimerPeriod = 30 * time.Second
)

type Transportation struct {
	duchy     *Duchy
	freeSlots []int // Stack of free slot numbers (1-99)
	jobArray  [TRANS_MAX_JOBS]JobSlot
	name      string
	timer     *time.Timer
}

func NewTransportation(duchy *Duchy, name string) *Transportation {
	t := &Transportation{
		duchy: duchy,
		name:  name,
	}
	for slotNo := TRANS_MAX_JOBS - 1; slotNo >= 1; slotNo-- {
		t.freeSlots = append(t.freeSlots, slotNo)
	}
	for i := range t.jobArray {
		t.jobArray[i].job.JobType = jobs.JOB_INVALID
	}
	return t
}

func (t *Transportation) CancelJobs(planet *Planet) {
	// TODO: Cancel all jobs from/to this planet
}

func (t *Transportation) Destroy() {
	t.Stop()
}

func (t *Transportation) ListJobs(caller *Player) {
	// TODO: List available jobs for acceptance
}

func (t *Transportation) ListWork(caller *Player) {
	// TODO: List current transportation workboard
	transportationReport(caller, t)
}

func (t *Transportation) Start() {
	if t.timer == nil {
		t.timer = time.AfterFunc(transportationTimerPeriod, t.timerHandler)
		log.Printf("%s transportation started", t.duchy.Name())
	}
}

func (t *Transportation) Stop() {
	if t.timer != nil {
		t.timer.Stop()
		t.timer = nil
		log.Printf("%s transportation stopped", t.duchy.Name())
	}
}

func (t *Transportation) advertiseJob(job model.Work, jobNo int) {
	for _, p := range Players { // FIXME: in current duchy only
		if p.Session() == nil {
			continue
		}
		if p.CurSys().IsHidden() {
			continue
		}
		if (p.Flags0 & model.PL0_JOB) == 0 {
			continue
		}
		if p.Rank() < model.RankCaptain {
			if job.Pallet.Quantity > 75 {
				continue
			}
			if job.JobType == jobs.JOB_FACTORY || job.JobType == jobs.JOB_PLANET {
				continue
			}
		}
		if (p.flags2 & PL2_IN_JOB_ADVERT) == 0 {
			p.Nsoutputm(text.JobAdvertHeader)
			p.flags2 |= PL2_IN_JOB_ADVERT
		}
		p.Nsoutput(JobReport(job, jobNo))
	}
}

func (t *Transportation) allocateJobSlot(job *model.Work) bool {
	// Check if we have any free slots
	if len(t.freeSlots) == 0 {
		return false // No slots available
	}

	// Pop a slot from the free stack
	slotIndex := t.freeSlots[len(t.freeSlots)-1]
	t.freeSlots = t.freeSlots[:len(t.freeSlots)-1]

	// Assign the job to this slot
	t.jobArray[slotIndex].job = *job
	t.jobArray[slotIndex].lastused = time.Now()

	debug.Trace("Transportation: allocated job slot %d for %s->%s", slotIndex+1, job.From, job.To)

	// Advertise the new job to eligible players
	t.advertiseJob(*job, slotIndex+1)

	return true
}

func (t *Transportation) ageJobs() {
	for i := range t.jobArray {
		job := &t.jobArray[i].job
		if job.JobType == jobs.JOB_INVALID {
			continue
		}

		// Age jobs at different rates
		switch job.JobType {
		case jobs.JOB_PLANET, jobs.JOB_FACTORY:
			job.Age += 3
		case jobs.JOB_GENERAL:
			job.Age += GENERAL_JOB_AGING
		default:
			job.Age++
		}

		// Handle expired jobs
		if job.Age > JOB_DELIVERY_AGE {
			t.handleExpiredJob(i)
		}
	}
}

func (t *Transportation) generateSolJobs() {
	// Sol job generation based on original C++ algorithm
	if global.Haulers <= 0 {
		return
	}

	// Calculate number of jobs: (haulers / Random(2,6)) + 1
	//nolint:gosec // "It's Just A Game"
	divisor := 2 + rand.IntN(5) // Random(2,6) means 2-6 inclusive
	numJobs := (global.Haulers / divisor) + 1
	if numJobs > global.Haulers {
		numJobs = global.Haulers
	}

	debug.Trace("Transportation: creating %d jobs for %d haulers", numJobs, global.Haulers)

	// Sol planet names matching C++ exactly
	places := []string{"Castillo", "Titan", "Mars", "Earth", "Moon", "Venus", "Mercury"}

	for i := range numJobs {
		// Pick random start and finish planets (different)
		//nolint:gosec // "It's Just A Game"
		start := rand.IntN(7) // 0-6 inclusive
		var finish int
		for {
			finish = rand.IntN(7) //nolint:gosec // "It's Just A Game"
			if finish != start {
				break
			}
		}

		// Create the job
		var job model.Work

		// Create cargo pallet
		//nolint:gosec // "It's Just A Game"
		job.Pallet.Type = model.Commodity(rand.IntN(52)) // 0-51 commodities
		job.Pallet.Quantity = 75
		job.Pallet.Origin = places[start]
		job.Pallet.Cost = int32(goods.GoodsArray[job.Pallet.Type].BasePrice)

		// Job details
		job.JobType = jobs.JOB_MILKRUN
		job.From = places[start]
		job.To = places[finish]
		job.Status = JOB_NONE
		//nolint:gosec // "It's Just A Game"
		job.Value = (150 + rand.Int32N(50)) / job.Pallet.Quantity // (150-199) / 75
		//nolint:gosec // "It's Just A Game"
		job.Gtu = 9 + rand.Int32N(4) // 9-12 inclusive
		job.Credits = 2

		// Apply value multipliers from C++
		//nolint:gosec // "It's Just A Game"
		job.Value *= (7 + rand.Int32N(4)) // Random(7,10) = 7-10 inclusive

		// 3% chance for 5x value bonus
		if rand.IntN(100) < 3 { //nolint:gosec // "It's Just A Game"
			job.Value *= 5
		}

		// Try to allocate the job to a slot
		if !t.allocateJobSlot(&job) {
			debug.Trace("Transportation: no slots available for job %d", i+1)
			break // No more slots available
		}
	}
}

func (t *Transportation) handleExpiredJob(slotIndex int) {
	job := &t.jobArray[slotIndex].job

	switch job.JobType {
	case jobs.JOB_PLANET:
		// Return goods to planet
		t.returnGoods(job)
		debug.Trace("Transportation: returning planet goods for job %d", slotIndex+1)
	default:
		if (10 + rand.Int32N(10)) < job.Value { //nolint:gosec // "It's Just A Game"
			// Auto-deliver if job value is high enough
			payment := job.Value * job.Pallet.Quantity
			// TODO: _DeliverGoods(job, payment)
			debug.Trace("Transportation: auto-delivering job %d for %d IG", slotIndex+1, payment)
		} else {
			// Return goods to origin
			t.returnGoods(job)
			debug.Trace("Transportation: returning goods for job %d", slotIndex+1)
		}
	}

	// Mark job as invalid and add slot back to free list
	job.JobType = jobs.JOB_INVALID
	t.freeSlots = append(t.freeSlots, slotIndex)
}

func (t *Transportation) returnGoods(job *model.Work) {
	switch job.JobType {
	case jobs.JOB_FACTORY:
		fromFactory := GetFactory(job.FactryWk.PickUp)
		if fromFactory != nil && !fromFactory.IsClosed() {
			fromFactory.ReturnGoods(job)
		}
	case jobs.JOB_PLANET:
		fromPlanet, ok := FindPlanet(job.From)
		if ok && !fromPlanet.IsClosed() {
			fromPlanet.ReturnGoods(job)
		}
	}

	job.JobType = jobs.JOB_INVALID
}

func (t *Transportation) timerHandler() {
	global.Lock()
	defer global.Unlock()

	monitoring.TransportationTimerTickTotal.WithLabelValues(t.duchy.Name()).Inc()

	t.timerProc()
}

func (t *Transportation) timerProc() {
	// log.Printf("Transportation timer proc for %s", t.duchy.Name())

	// Age all jobs and handle expired ones
	t.ageJobs()

	// Generate new jobs if applicable (Sol system gets special treatment)
	if t.duchy.Name() == "Sol" {
		t.generateSolJobs()
	}

	// Generate planet jobs
	// TODO: gen_planet_jobs()

	// Clean up after the adverts
	EndJobAdvert()

	t.timer = time.AfterFunc(transportationTimerPeriod, t.timerHandler)
}

func transportationReport(caller *Player, t *Transportation) {
	caller.Outputm(text.TransportationReportHeader)

	for i := range t.jobArray {
		if t.jobArray[i].job.JobType == jobs.JOB_INVALID {
			continue
		}
		caller.Output(JobReport(t.jobArray[i].job, i+1))
	}

	// if index == 0 {
	// 	caller.Outputm(text.MN633)
	// } else {
	caller.Outputm(text.TransportationReportTrailer)
	// }
}
