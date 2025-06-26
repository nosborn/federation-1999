package server

import (
	"log"
	"strings"
	"time"

	"github.com/nosborn/federation-1999/internal/collections"
	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/horsell"
	"github.com/nosborn/federation-1999/internal/text"
)

var (
	allSystems  = collections.NewOrderedCollection[Systemer]()
	systemIndex = collections.NewNameIndex[Systemer]()
)

type Systemer interface {
	// Core methods
	Name() string
	Balance() int32
	Duchy() *Duchy
	LinkLocNo() int32
	Owner() *Player

	// State methods
	Flags() uint32
	LoadState() SystemLoadState
	LastOnline() int32
	TaxRate() int32
	TouristTime() int32

	// Collection accessors -- FIXME: bad!
	Events() []*Event
	Objects() []*Object
	Planets() []*Planet

	// Location methods
	FindLocation(locNo uint32) *Location
	LandingLocNo(orbitLocNo uint32) uint32
	OrbitLocNo(landingLocNo uint32) uint32

	// Object methods
	FindObject(number uint32) (*Object, bool)
	FindObjectName(name model.Name) (*Object, bool)

	// Planet methods
	FindPlanet(name string) (*Planet, bool)
	GuessPlanet(locNo uint32) *Planet

	// Status methods
	IsArena() bool
	IsCapital() bool
	IsClosed() bool
	IsHidden() bool
	IsHorsell() bool
	IsLoading() bool
	IsOffline() bool
	IsOwnerPlaying() bool
	IsSnark() bool
	IsSol() bool
	IsUnloading() bool

	// Financial methods
	Income(amount int32, taxable bool)
	Expenditure(amount int32)
	ExchangeTicks() int

	// Lifecycle methods
	Start()
	Stop()

	// Player interaction
	Talk(message string, omit ...*Player)
	UpdateTime(caller *Player)
	Overlord() string

	// Configuration
	SetFlags(v uint32)
	SetLoadState(state SystemLoadState)
	SetTaxRate(newRate int32)
	SetDuchy(d *Duchy)

	// Persistence
	Save(when database.SaveWhen)
	Serialize(dbp *model.DBPlanet)

	// Hooks
	BlastHook(player *Player, name model.Name, object *Object) model.HookResult
	CleanupHook(player *Player)
	CmdGiveHook(player *Player, amount int32, mobile *Object) bool
	CmdSlideHook(player *Player) bool
	DisplayRoutesHook(caller *Player)
	DrinkEvent(player *Player, object *Object) bool
	GetHook(player *Player, object *Object) bool
	GiveHook(player *Player, object *Object, mobile *Object) model.HookResult
	KillMobileHook(mobile *Object, player *Player)
	NoExitHook(player *Player, direction model.Direction) model.HookResult
	PlaceHook(object *Object) model.HookResult
	RecycleHook(object *Object) model.HookResult
	UseHook(player *Player, object *Object) model.HookResult
	WaitHook(player *Player)
	WantedHook(caller *Player)
}

// System is the concrete implementation of Systemer
type System struct {
	name string

	events    []*Event
	locations []*Location
	objects   []*Object
	planets   []*Planet

	balance int32 // Treasury balance
	duchy   *Duchy
	// Events
	flags      uint32
	loadState  SystemLoadState
	linkLocNo  int32 // Interstellar Link location number
	owner      *Player
	lastOnline int32
	// done        chan struct{}
	populated   int32
	taxRate     int32
	touristTime int32
}

func AllSystems() []Systemer {
	return allSystems.All()
}

func FindSystem(name string) (Systemer, bool) {
	return systemIndex.Find(name)
}

// // NewPlayerSystem creates a minimal player system from DBPersona data.
// // This creates the system structure without the planet implementation.
// // In Federation terminology, players own planets, but we often refer to "player systems"
// // when talking about the system containing their planet.
// func NewPlayerSystem(owner *Player, systemName string, dbPlanet model.DBPlanet) *System {
// 	s := &System{
// 		flags:       dbPlanet.Flags &^ model.PLT_CLOSED,
// 		owner:       owner,
// 		balance:     dbPlanet.Balance,
// 		lastOnline:  dbPlanet.LastOnline,
// 		loadState:   SystemLoading,
// 		name:        systemName,
// 		taxRate:     dbPlanet.Tax,
// 		touristTime: dbPlanet.Time,
// 	}
//
// 	var ok bool
// 	s.duchy, ok = FindDuchy(text.CStringToString(dbPlanet.Duchy[:]))
// 	if !ok {
// 		log.Printf("Moving %s system to Sol duchy", s.name)
// 		s.duchy = SolDuchy
// 	}
// 	s.duchy.AddMember(s)
//
// 	if err := allSystems.Insert(s); err != nil {
// 		log.Panic("PANIC: Duplicate system added: ", err)
// 	}
// 	systemIndex.Insert(s.name, s)
//
// 	planet := NewPlayerPlanet(s, dbPlanet)
// 	s.planets = append(s.planets, planet)
//
// 	if (s.flags & model.PLT_CLOSED) == 0 && !p.IsLockedOut()
// 		Enqueue(s)
// 	}
//
// 	return s
// }

func (s *System) Balance() int32 {
	return s.balance
}

func (s *System) BlastHook(_ *Player, _ model.Name, _ *Object) model.HookResult {
	return model.HookContinue
}

func (s *System) CleanupHook(_ *Player) {} // Do nothing.

func (s *System) CmdGiveHook(_ *Player, _ int32, _ *Object) bool {
	return false
}

func (s *System) CmdSlideHook(_ *Player) bool {
	return false
}

func (s *System) destroyEvents() {
	for i, e := range s.events {
		e.Destroy()
		s.events[i] = nil
	}
	s.events = s.events[:0]
}

func (s *System) destroyLocations() {
	for i, l := range s.locations {
		l.Destroy()
		s.locations[i] = nil
	}
	s.locations = s.locations[:0]
}

func (s *System) destroyObjects() {
	for i, o := range s.objects {
		o.Destroy()
		s.objects[i] = nil
	}
	s.objects = s.objects[:0]
}

func (s *System) DisplayRoutesHook(caller *Player) {
	// TODO
}

func (s *System) DrinkEvent(_ *Player, _ *Object) bool {
	return false
}

// Duchy returns the duchy this system belongs to.
func (s *System) Duchy() *Duchy {
	return s.duchy
}

func (s *System) Events() []*Event { // FIXME
	return s.events
}

func (s *System) ExchangeTicks() int {
	return s.duchy.ExchangeTicks()
}

func (s *System) Expenditure(amount int32) {
	if s.IsClosed() {
		log.Panicf("expenditure: %s system is closed! (%d IG)", s.Name(), amount)
	}
	changeBalance(&s.balance, -amount)
}

// FIXME: The spaceship locations should be coming from an independent object
// rather than having knowledge of SolSystem here.
func (s *System) FindLocation(locNo uint32) *Location {
	debug.Check(s != nil)

	// for i := range s.locations {
	// 	log.Printf("%#v", s.locations[i])
	// }

	if s.name == "Sol" {
		if locNo >= 1 && locNo <= uint32(len(s.locations)) {
			return s.locations[locNo-1]
		}
		log.Printf("Can't find %s/%d", s.name, locNo)
		return nil
	}
	if locNo >= 1 {
		if locNo <= SPACESHIP_SIZE {
			return SolSystem.FindLocation(locNo)
		}
		if locNo <= uint32(len(s.locations)+SPACESHIP_SIZE) {
			return s.locations[locNo-(SPACESHIP_SIZE+1)]
		}
	}
	log.Printf("Can't find %s/%d", s.name, locNo)
	return nil
}

func (s *System) FindObject(number uint32) (*Object, bool) {
	for _, o := range s.objects {
		if o.number == number {
			return o, true
		}
	}
	return nil, false
}

func (s *System) FindObjectName(name model.Name) (*Object, bool) {
	for _, o := range s.objects {
		if name.The && o.noThe() {
			continue
		}
		if strings.EqualFold(o.Name(), name.Text) {
			return o, true
		}
		if name.Words != 1 {
			continue
		}
		for j := range o.Synonyms {
			if strings.EqualFold(o.Synonyms[j], name.Text) {
				return o, true
			}
		}
	}
	return nil, false
}

func (s *System) FindPlanet(name string) (*Planet, bool) {
	for _, p := range s.planets {
		if strings.EqualFold(p.name, name) || strings.EqualFold(p.synonym, name) {
			return p, true
		}
	}
	return nil, false
}

func (s *System) Flags() uint32 {
	return s.flags
}

func (s *System) GetHook(_ *Player, _ *Object) bool {
	return false
}

func (s *System) GiveHook(_ *Player, _ *Object, _ *Object) model.HookResult {
	return model.HookContinue
}

func (s *System) GuessPlanet(locNo uint32) *Planet {
	loc := s.FindLocation(locNo)
	if loc == nil || loc.IsSpace() {
		return nil
	}

	if len(s.planets) == 1 {
		return s.planets[0]
	}

	for _, p := range s.planets {
		if locNo == p.exchangeLocNo || locNo == p.hospitalLocNo || locNo == p.landingLocNo {
			return p
		}
	}

	if s.name != "Sol" {
		return nil
	}

	if locNo < 70 { // IsSpace() should have handled this already
		return nil
	}
	if locNo < 92 { // Callisto
		return s.planets[1]
	}
	if locNo < 131 { // Titan
		return s.planets[0]
	}
	if locNo < 225 { // Moon
		return s.planets[4]
	}
	if locNo < 368 { // Mars
		return s.planets[2]
	}
	if locNo <= 515 /* LadiesLoo */ { // Earth
		return s.planets[3]
	}
	if locNo < 596 { // Venus
		return s.planets[5]
	}
	if locNo <= 671 /* Chamber */ { // Mercury
		return s.planets[6]
	}
	if locNo <= 696 /* MeetingPoint */ { // Earth
		return s.planets[3]
	}
	if locNo <= 697 /* DieselsBoudoir */ { // Mars
		return s.planets[2]
	}
	return nil
}

func (s *System) Income(amount int32, taxable bool) {
	if s.IsClosed() {
		log.Panicf("income: %s system is closed! (%d IG)", s.Name(), amount)
	}

	// Take the duchy tax.
	if amount > 0 { //nolint:staticcheck // SA9003: empty branch
		// Slice off the Duke's cut of the money.
		// FIXME: use Duchy.taxRate().

		// TODO
	}

	// The planet gets what's left after Duchy taxes.
	changeBalance(&s.balance, amount)
}

func (s *System) IsArena() bool {
	return s.name == "Arena"
}

func (s *System) IsCapital() bool {
	for _, p := range s.planets {
		if p.level == model.LevelCapital {
			return true
		}
	}
	return false
}

func (s *System) IsClosed() bool {
	return s.loadState != SystemOnline
}

func (s *System) IsHidden() bool {
	return s.IsHorsell() || s.IsSnark()
}

func (s *System) IsHorsell() bool {
	return horsell.NamePattern.MatchString(s.name)
}

func (s *System) IsLoading() bool {
	return s.loadState == SystemLoading
}

func (s *System) IsOffline() bool {
	return s.loadState == SystemOffline
}

func (s *System) IsOwnerPlaying() bool {
	if s.owner == nil || !s.owner.OwnsPlanet() {
		return false
	}
	return s.owner.IsPlaying() && !s.owner.IsOnDutyNavigator()
}

func (s *System) IsSnark() bool {
	return s.name == "Snark"
}

func (s *System) IsSol() bool {
	return s.name == "Sol"
}

func (s *System) IsUnloading() bool {
	return s.loadState == SystemUnloading
}

func (s *System) KillMobileHook(_ *Object, _ *Player) {} // Do nothing.

func (s *System) LandingLocNo(orbitLocNo uint32) uint32 {
	for _, p := range s.planets {
		if p.orbitLocNo == orbitLocNo {
			return p.landingLocNo
		}
	}
	return 0
}

func (s *System) LastOnline() int32 {
	return s.lastOnline
}

// LoadState returns the system's current load state
func (s *System) LoadState() SystemLoadState {
	return s.loadState
}

// Name returns the system's name.
func (s *System) Name() string {
	return s.name
}

// LinkLocNo returns the interstellar link location number.
func (s *System) LinkLocNo() int32 {
	return s.linkLocNo
}

func (s *System) NoExitHook(_ *Player, _ model.Direction) model.HookResult {
	return model.HookContinue
}

func (s *System) Objects() []*Object { // FIXME
	return s.objects
}

func (s *System) OrbitLocNo(landingLocNo uint32) uint32 {
	for _, p := range s.planets {
		if p.landingLocNo == landingLocNo {
			return p.orbitLocNo
		}
	}
	return 0
}

func (s *System) Overlord() string {
	if s.owner == nil {
		return text.Msg(text.OverlordMing)
	}
	return s.owner.Name()
}

func (s *System) Owner() *Player {
	return s.owner
}

func (s *System) PlaceHook(_ *Object) model.HookResult {
	return model.HookContinue
}

func (s *System) Planets() []*Planet { // FIXME
	return s.planets
}

func (s *System) RecycleHook(_ *Object) model.HookResult {
	return model.HookContinue
}

func (s *System) Save(when database.SaveWhen) {
	if s.owner == nil { // don't try to save core systems
		return
	}
	debug.Trace("System.Save(%s,%d)", s.name, when)
	s.owner.Save(when)
}

func (s *System) Serialize(dbp *model.DBPlanet) {
	debug.Precondition(dbp != nil)

	debug.Check(len(s.planets) == 1 && s.planets[0] != nil)
	s.planets[0].Serialize(dbp)

	copy(dbp.Duchy[:], s.duchy.Name())

	dbp.Tax = s.taxRate
	dbp.Time = s.touristTime
	dbp.Balance = s.balance
	dbp.Flags = s.flags
	dbp.LastOnline = s.lastOnline

	if s.loadState == SystemOffline || s.loadState == SystemUnloading {
		dbp.Flags |= model.PLT_CLOSED
	}
}

func (s *System) SetDuchy(d *Duchy) {
	if s.duchy != nil {
		s.duchy.RemoveMember(s)
	}
	s.duchy = d
	s.duchy.AddMember(s)
}

func (s *System) SetFlags(v uint32) {
	s.flags = v
}

func (s *System) SetLoadState(state SystemLoadState) {
	s.loadState = state
}

func (s *System) SetTaxRate(newRate int32) {
	s.taxRate = max(0, min(newRate, 30))
}

func (s *System) Start() {
	if s.loadState != SystemOnline {
		return
	}
	log.Printf("Starting %s system", s.name)

	for _, planet := range s.planets {
		planet.StartExchange()
	}
}

func (s *System) Stop() {
	for _, planet := range s.planets {
		planet.StopExchange()
	}

	// if s.name == "Sol" {
	// 	if PublicAddressTimer != nil {
	// 		PublicAddressTimer.Stop()
	// 		PublicAddressTimer = nil
	// 	}
	// 	if ShuttleTimer != nil {
	// 		ShuttleTimer.Stop()
	// 		ShuttleTimer = nil
	// 	}
	// }
}

func (s *System) Talk(message string, omit ...*Player) {
players:
	for _, player := range Players {
		for i := range omit {
			if player == omit[i] {
				continue players
			}
		}
		player.Output(message)
		if !player.isActiveThread() {
			player.FlushOutput()
		}
	}
}

func (s *System) TaxRate() int32 {
	return s.taxRate
}

func (s *System) TouristTime() int32 {
	return s.touristTime
}

func (s *System) UpdateTime(caller *Player) {
	if s != caller.OwnSystem() && !s.IsClosed() {
		s.touristTime++
	}
	s.populated = int32(time.Now().Unix()) /*Transaction::time()*/
}

func (s *System) UseHook(_ *Player, _ *Object) model.HookResult {
	return model.HookContinue
}

func (s *System) WaitHook(_ *Player) {} // Do nothing

// NOTE: Includes RECYCLING objects.
func (s *System) WantedHook(caller *Player) {
	debug.Precondition(caller != nil)

	if s.IsClosed() || s.balance <= 0 {
		return
	}

	// TODO
}
