package server

import (
	"log"

	"github.com/nosborn/federation-1999/internal/collections"
	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/global"
	"github.com/nosborn/federation-1999/internal/server/horsell"
	"github.com/nosborn/federation-1999/internal/text"
)

// const (
// 	maxPlayerDuchyMembers = 20
// )

var (
	allDuchies = collections.NewOrderedCollection[*Duchy]()
	duchyIndex = collections.NewNameIndex[*Duchy]()
)

type Duchy struct {
	customsRate    int32  // import duty from non-favoured duchy
	embargo        *Duchy // string // embargoed duchy (may be empty)
	favoured       *Duchy // string // favoured duchy (may be empty)
	favouredRate   int32  // import duty from favoured duchy
	name           string
	owner          *Player
	systems        *collections.OrderedCollection[Systemer] // planets currently in the duchy
	taxRate        int32                                    // tax rate on member planets' income
	transportation *Transportation
}

func AllDuchies() []*Duchy {
	return allDuchies.All()
}

func FindDuchy(name string) (*Duchy, bool) {
	return duchyIndex.Find(name)
}

func (d *Duchy) AddMember(s Systemer) {
	if err := d.systems.Insert(s); err != nil {
		log.Panic("PANIC: Duplicate system added into duchy: ", err)
	}
}

func (d *Duchy) AllMembers() []Systemer {
	return d.systems.All()
}

func (d *Duchy) Broadcast(text string, omit *Player) {
	for _, p := range Players {
		if p == omit || p.CurSys().Duchy() != d || !p.HasCommUnit() {
			continue
		}
		p.Output(text)
		p.FlushOutput()
	}
}

func (d *Duchy) CapitalSystem() Systemer {
	if d.owner != nil {
		return d.owner.ownSystem
	}
	s, ok := FindSystem(d.name)
	if !ok {
		log.Panic("PANIC: can't find capital system for ", d.name)
	}
	return s
}

func (d *Duchy) ClearEmbargo(caller *Player) bool {
	if d.embargo == nil {
		if caller != nil {
			caller.Outputm(text.ClearEmbargoNotSet)
			return false
		}
	}
	d.embargo = nil
	if caller != nil {
		caller.Outputm(text.ClearEmbargoOK)
	}
	return true
}

func (d *Duchy) ClearFavoured(caller *Player) bool {
	if d.favoured == nil {
		if caller != nil {
			caller.Outputm(text.ClearFavouredNotSet)
			return false
		}
	}
	d.favoured = nil
	d.favouredRate = 0
	if caller != nil {
		caller.Outputm(text.ClearFavouredOK)
	}
	return true
}

func (d *Duchy) CustomsRate() int32 {
	return d.customsRate
}

func (d *Duchy) Embargo() *Duchy {
	return d.embargo
}

func (d *Duchy) Delete() { // destructor
	if d.transportation != nil {
		d.transportation.Destroy()
		d.transportation = nil
	}
	duchyIndex.Remove(d.name)
	allDuchies.Remove(d)
	debug.Trace("%s duchy deleted", d.name)
}

func (d *Duchy) Destroy() {
	if d.IsSol() || d.IsHorsell() { // sanity check
		log.Panicf("Duchy.Destroy called for %s", d.name)
	}

	log.Printf("Destroying %s duchy [%d]", d.name, d.owner.UID())

	for member := range d.systems.Values() { // iterate over a copy of d.systems
		member.SetDuchy(SolDuchy)
	}

	// Clean up favoured status and embargo of other duchies.
	for other := range allDuchies.Values() {
		if other.Embargo() == d {
			other.ClearEmbargo(nil)
		}
		if other.Favoured() == d {
			other.ClearFavoured(nil)
		}
	}

	// Remove the owner. This might be a bad idea...
	d.owner = nil
}

func (d *Duchy) ExchangeTicks() int {
	return global.ExchangeTicks
}

func (d *Duchy) Favoured() *Duchy {
	return d.favoured
}

func (d *Duchy) FavouredRate() int32 {
	return d.favouredRate
}

func (d *Duchy) ImportDuty(c model.Cargo) int32 {
	taxRate := d.customsRate
	if d.favoured != nil {
		origin, ok := FindPlanet(c.Origin)
		if ok && origin.System().Duchy() == d.favoured {
			taxRate = d.favouredRate
		}
	}
	return (c.Quantity * c.Cost * taxRate) / 100
}

func (d *Duchy) IsClosed() bool {
	switch {
	case d.IsSol():
		return false
	case d.IsHorsell():
		return true
	case d.owner != nil:
		return d.owner.ownSystem.IsClosed()
	default: // probably shouldn't happen
		return true
	}
}

func (d *Duchy) IsHidden() bool {
	return d.IsHorsell()
}

func (d *Duchy) IsHorsell() bool {
	return horsell.NamePattern.MatchString(d.name)
}

func (d *Duchy) IsSol() bool {
	return d.name == "Sol"
}

func (d *Duchy) ListWork(caller *Player) {
	d.transportation.ListWork(caller)
}

func (d *Duchy) Members() int {
	return d.systems.Len()
}

func (d *Duchy) Name() string {
	return d.name
}

func (d *Duchy) Owner() *Player {
	return d.owner
}

func (d *Duchy) RemoveMember(s Systemer) {
	if err := d.systems.Remove(s); err != nil {
		log.Fatal("not found: ", err)
	}
}

func (d *Duchy) Serialize(dbd *model.DBDuchy) {
	if d.IsSol() || d.IsHorsell() { // sanity check
		log.Panicf("Duchy.Serialize called for %s", d.name)
	}

	if d.favoured != nil {
		copy(dbd.Favoured[:], d.favoured.Name())
	}
	if d.embargo != nil {
		copy(dbd.Embargo[:], d.embargo.Name())
	}
	dbd.CustomsRate = d.customsRate
	dbd.FavouredRate = d.favouredRate
}

func (d *Duchy) SetCustomsRate(newRate int32, caller *Player) bool {
	debug.Check(newRate >= 0 && newRate <= 50)

	d.customsRate = newRate
	if caller != nil {
		caller.Outputm(text.MN1310)
	}
	if d.customsRate < d.favouredRate {
		d.SetFavouredRate(d.customsRate, caller)
	}
	return true
}

func (d *Duchy) SetEmbargo(duchyName string, caller *Player) bool {
	embargoedDuchy, ok := FindDuchy(duchyName)
	if !ok || embargoedDuchy.IsHidden() {
		if caller != nil {
			caller.Outputm(text.NoSuchDuchy)
		}
		return false
	}
	if embargoedDuchy == d || embargoedDuchy.IsSol() {
		if caller != nil {
			caller.Outputm(text.DontBeSilly)
		}
		return false
	}
	if embargoedDuchy == d.favoured {
		d.ClearFavoured(caller)
	}
	d.embargo = embargoedDuchy
	if caller != nil {
		caller.Outputm(text.SetEmbargoOK)
	}
	return true
}

func (d *Duchy) SetFavoured(duchyName string, caller *Player) bool {
	debug.Precondition(duchyName != "")

	favouredDuchy, ok := FindDuchy(duchyName)
	if !ok || favouredDuchy.IsHidden() {
		if caller != nil {
			caller.Outputm(text.NoSuchDuchy)
		}
		return false
	}
	if favouredDuchy == d {
		if caller != nil {
			caller.Outputm(text.DontBeSilly)
		}
		return false
	}
	if favouredDuchy == d.embargo {
		d.ClearEmbargo(caller)
	}
	d.favoured = favouredDuchy
	if caller != nil {
		caller.Outputm(text.SetFavouredOK)
	}
	return true
}

func (d *Duchy) SetFavouredRate(newRate int32, caller *Player) bool {
	debug.Check(newRate >= 0 && newRate <= 50)

	d.favouredRate = newRate
	if caller != nil {
		caller.Outputm(text.SetFavouredTaxOK, d.favouredRate)
	}
	return true
}

func (d *Duchy) SetTaxRate(rate int32) {
	d.taxRate = rate
}

func (d *Duchy) StartTransportation() {
	if d.transportation != nil {
		d.transportation.Start()
	}
}

func (d *Duchy) Stop() {
	//
}

func (d *Duchy) TaxRate() int32 {
	return d.taxRate
}
