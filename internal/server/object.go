package server

import (
	"log"
	"math/rand/v2"
	"time"

	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/monitoring"
	"github.com/nosborn/federation-1999/internal/server/core"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/global"
	"github.com/nosborn/federation-1999/internal/server/snark"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

const (
	MIN_RECYCLE_SECONDS = 10 * 60
	MAX_RECYCLE_SECONDS = 40 * 60
)

const (
	// fightTimerPeriod = 1 * time.Second
	moveTimerPeriod = 7 * time.Second
)

type ObjectEvents [4]int

type ObjectShipGun struct {
	Damage int
	Name   string
	Power  int
	Type   int
}

type ObjectShipKit struct {
	Computer int
	Engine   int
	Fuel     int
	Hold     int
	Hull     int
	Shield   int
	Tonnage  int
}

type MobileGun struct {
	gunType  int
	damage   int
	opponent int64
}

type Object struct {
	Flags     uint32
	GiveEvent int
	Sex       model.Sex
	ShipGuns  [4]MobileGun
	ShipKit   ObjectShipKit
	Synonyms  []string

	attackPercent uint32
	consumeEvent  int
	desc          string
	descMsgNo     text.MsgNum
	dropEvent     int
	events        ObjectEvents
	getEvent      int
	maxCounter    int
	maxLocNo      uint32
	minLocNo      uint32
	name          string
	number        uint32
	prefObject    int
	scan          string
	scanMsgNo     text.MsgNum
	value         int32
	weight        uint32

	cleanTimer *time.Timer
	curLocNo   uint32
	// descText     string
	fightTimer   *time.Timer
	homeSystem   Systemer
	moveTimer    *time.Timer
	recycleTimer *time.Timer
}

func NewCoreObject(data *core.Object, system Systemer) *Object {
	o := &Object{
		events:        [4]int{0, 0, 0, 0}, // Events need to be mapped separately
		Flags:         data.Flags,
		GiveEvent:     int(data.GiveEvent),
		Sex:           data.Sex,
		Synonyms:      data.Synonyms,
		attackPercent: uint32(data.AttackPercent),
		desc:          text.Msg(data.DescMessageNo),
		dropEvent:     int(data.DropEvent),
		getEvent:      int(data.GetEvent),
		homeSystem:    system,
		maxCounter:    int(data.MaxCounter),
		maxLocNo:      uint32(data.MaxLoc),
		minLocNo:      uint32(data.MinLoc),
		name:          data.Name,
		number:        uint32(data.Number),
		prefObject:    int(data.PrefObject),
		scan:          text.Msg(data.ScanMessageNo),
		value:         data.Value,
		weight:        uint32(data.Weight),

		ShipKit: ObjectShipKit{
			Computer: int(data.ShipKit.Computer),
			Engine:   int(data.ShipKit.Engine),
			Fuel:     int(data.ShipKit.Fuel),
			Hold:     int(data.ShipKit.Hold),
			Hull:     int(data.ShipKit.Hull),
			Shield:   int(data.ShipKit.Shield),
			Tonnage:  int(data.ShipKit.Tonnage),
		},
	}
	for i := range data.ShipGuns {
		o.ShipGuns[i] = MobileGun{
			gunType: int(data.ShipGuns[i].Type),
			damage:  int(data.ShipGuns[i].Damage),
		}
	}
	if o.Flags&model.OfHidden == 0 {
		o.Place()
	}
	return o
}

// NewObjectFromDB converts a DBObject from disk to an Object
func NewObjectFromDB(dbObject model.DBObject, system Systemer) *Object {
	o := &Object{
		events:        [4]int{0, 0, 0, 0}, // Events need to be mapped separately
		Flags:         dbObject.Flags,
		GiveEvent:     int(dbObject.GiveEvent),
		Sex:           model.Sex(dbObject.Sex),
		attackPercent: uint32(dbObject.AttackPercent),
		desc:          text.CStringToString(dbObject.Desc[:]),
		dropEvent:     int(dbObject.DropEvent),
		getEvent:      int(dbObject.GetEvent),
		maxCounter:    int(dbObject.MaxCounter),
		maxLocNo:      uint32(dbObject.MaxLoc),
		minLocNo:      uint32(dbObject.MinLoc),
		name:          text.CStringToString(dbObject.Name[:]),
		number:        uint32(dbObject.Number),
		prefObject:    int(dbObject.PrefObject),
		scan:          text.CStringToString(dbObject.Scan[:]),

		ShipKit: ObjectShipKit{
			Computer: int(dbObject.Computer),
			Engine:   int(dbObject.Engine),
			Fuel:     int(dbObject.Fuel),
			Hold:     int(dbObject.Hold),
			Hull:     int(dbObject.Hull),
			Shield:   int(dbObject.Shield),
			Tonnage:  int(dbObject.Tonnage),
		},

		value:      dbObject.Value,
		weight:     uint32(dbObject.Weight),
		curLocNo:   uint32(dbObject.CurLoc), // TODO: check this
		homeSystem: system,
	}

	// Convert ship guns
	for i := range dbObject.ShipGuns {
		o.ShipGuns[i] = MobileGun{
			gunType: int(dbObject.ShipGuns[i].GunType),
			damage:  int(dbObject.ShipGuns[i].Damage),
		}
	}

	return o
}

func (o *Object) ConsumeEvent() int {
	return o.consumeEvent
}

func (o *Object) CurLocNo() uint32 {
	return o.curLocNo
}

func (o *Object) Desc() string {
	if o.descMsgNo > 0 {
		return text.Msg(o.descMsgNo)
	}
	debug.Check(o.desc != "")
	return o.desc
}

func (o *Object) Destroy() {
	// TODO
}

func (o *Object) DisplayName(capitalize bool) string {
	if o.noThe() {
		return o.name
	}
	if capitalize {
		return text.Msg(text.ObjectName_Cap, o.name)
	} else {
		return text.Msg(text.ObjectName_NoCap, o.name)
	}
}

func (o *Object) DropEvent() int {
	return o.dropEvent
}

func (o *Object) Events() [4]int {
	return o.events
}

func (o *Object) GetEvent() int {
	return o.getEvent
}

func (o *Object) Name() string {
	return o.name
}

func (o *Object) Number() uint32 {
	return o.number
}

func (o *Object) Scan() string {
	if o.scanMsgNo > 0 {
		switch o.Number() {
		case sol.ObShare:
			return text.Msg(o.scanMsgNo, GetLaunchCode())
		case snark.ObPaper:
			return text.Msg(o.scanMsgNo, GetLaunchCoordinates())
		default:
			return text.Msg(o.scanMsgNo)
		}
	}
	debug.Check(o.scan != "")
	return o.scan
}

func (o *Object) cleanTimerHandler() {
	global.Lock()
	defer global.Unlock()

	monitoring.CleanTimerTickTotal.WithLabelValues(o.homeSystem.Name()).Inc()

	// defer database.CommitDatabase() -- doesn't do anything to warrant this
	o.cleanTimerProc()
}

func (o *Object) cleanTimerProc() {
	// log.Printf("Clean timer proc for %s in %s", o.name, o.homeSystem.Name())

	o.cleanTimer = nil

	if o.IsHidden() || o.IsRecycling() {
		return
	}

	o.homeSystem.FindLocation(o.curLocNo).Clean(o)
}

func (o *Object) Hide() {
	o.Flags |= model.OfHidden
	o.curLocNo = 0

	// Clear any fights.
	if o.IsSpaceship() {
		for i := range len(o.ShipGuns) {
			o.ShipGuns[i].opponent = 0
		}
	}
}

func (o *Object) HomeSystem() Systemer {
	return o.homeSystem
}

func (o *Object) immobilize() {
	o.Hide()
	o.maxCounter = 0
}

func (o *Object) isCleaner() bool {
	return (o.Flags & model.OfCleaner) != 0
}

func (o *Object) isDisplaced(l *Location) bool {
	if o.IsHidden() || o.IsRecycling() { // caller should have done this already?
		return false
	}
	if o.curLocNo < o.minLocNo || o.curLocNo > o.maxLocNo {
		return true
	}
	if (o.Flags&model.OfIndoors) != 0 && !l.IsIndoors() {
		return true
	}
	if (o.Flags&model.OfOutdoors) != 0 && !l.IsOutdoors() {
		return true
	}
	return false
}

func (o *Object) IsEdible() bool {
	return (o.Flags & model.OfEdible) != 0
}

func (o *Object) IsHidden() bool {
	return (o.Flags & model.OfHidden) != 0
}

func (o *Object) IsLiquid() bool {
	return (o.Flags & model.OfLiquid) != 0
}

func (o *Object) IsMobile() bool {
	return (o.Flags & model.OfAnimate) != 0
}

func (o *Object) IsRecycling() bool {
	return o.recycleTimer != nil
}

func (o *Object) IsSpaceship() bool {
	return (o.Flags & model.OfShip) != 0
}

func (o *Object) isValidPlacement(locNo uint32) bool {
	l := o.homeSystem.FindLocation(locNo)
	if l == nil {
		return false
	}

	// Avoid HIDDEN locations.
	if l.IsHidden() {
		return false
	}

	if o.IsSpaceship() {
		// Keep SHIP mobiles in SPACE.
		if !l.IsSpace() {
			log.Print("SPACE mobile passes through non-SPACE location!")
			return false
		}
	} else {
		// Keep ground-bound mobiles out of SPACE.
		if l.IsSpace() {
			return false
		}

		// Keep INDOOR mobiles indoors and OUTDOOR mobiles outdoors.
		if (o.Flags&model.OfIndoors) != 0 && !l.IsIndoors() {
			return false
		}

		// Keep OUTDOOR mobiles outdoors.
		if (o.Flags&model.OfOutdoors) != 0 && !l.IsOutdoors() {
			return false
		}
	}

	return true
}

func (o *Object) moveTimerHandler() {
	global.Lock()
	defer global.Unlock()

	monitoring.MoveTimerTickTotal.WithLabelValues(o.homeSystem.Name()).Inc()

	defer database.CommitDatabase()
	o.moveTimerProc()
}

func (o *Object) moveTimerProc() {
	// log.Printf("moveTimerProc for %s in %s", o.name, o.homeSystem.Name())

	o.moveTimer = nil

	// Don't move hidden or recycling objects.
	if o.IsHidden() || o.IsRecycling() {
		return
	}

	// Don't move if it's not a mobile or has been immobilised.
	if (o.Flags & model.OfAnimate) == 0 {
		return
	}
	if o.minLocNo == o.maxLocNo || o.maxCounter <= 0 {
		return
	}

	// Leave the current location.
	fromLoc := o.homeSystem.FindLocation(o.curLocNo)
	if fromLoc == nil {
		log.Printf("object.MoveTimerProc: FindLocation() failed")
		o.immobilize()
		return
	}
	if (o.Flags & model.OfShip) == 0 {
		message := text.Msg(text.MOBILE_DEPARTURE, o.DisplayName(true))
		fromLoc.Talk(message)
	} else {
		message := text.Msg(text.SHIP_MOBILE_DEPARTURE, o.DisplayName(true))
		fromLoc.Talk(message)
	}

	// Enter the next location.
	toLoc := o.nextLocation()
	if toLoc == nil {
		log.Printf("object.MoveTimerProc: nextLocation() failed")
		o.immobilize()
		return
	}

	o.curLocNo = toLoc.Number()
	if o.IsSpaceship() {
		message := text.Msg(text.SHIP_MOBILE_ARRIVAL, o.name)
		toLoc.Talk(message)

		if toLoc.IsPeaceful() {
			// o.ShipKit.CurHull = o.ShipKit.MaxHull
		} else if o.attackPercent > 0 && o.fightTimer == nil { //nolint:staticcheck // SA9003: empty branch
			// FIXME:
		}
	} else {
		message := text.Msg(text.MOBILE_ARRIVAL, o.DisplayName(true))
		toLoc.Talk(message)
	}

	//
	if o.isCleaner() {
		duration := time.Duration(1+rand.IntN((o.maxCounter*7)-1)) * time.Second //nolint:gosec // "It's Just A Game"
		o.cleanTimer = time.AfterFunc(duration, o.cleanTimerHandler)
	}

	duration := moveTimerPeriod * time.Duration(o.maxCounter)
	o.moveTimer = time.AfterFunc(duration, o.moveTimerHandler)
}

// Move on to the next (possible) location, wrapping around if needed.
func (o *Object) nextLocation() *Location {
	nextLocNo := o.curLocNo + 1

	for {
		if nextLocNo > o.maxLocNo {
			nextLocNo = o.minLocNo
		}
		if nextLocNo == o.curLocNo {
			log.Printf("nextLocation: %s mobile is stuck (min=%d, max=%d, cur=%d)",
				o.name,
				o.minLocNo,
				o.maxLocNo,
				o.curLocNo)
			return nil
		}
		// See if the candidate location is valid.
		if o.isValidPlacement(nextLocNo) {
			// debug.Trace("%s -> %s/%d", m_name.c_str(), pstarHome->name(), cur_loc);
			return o.homeSystem.FindLocation(nextLocNo)
		}
		nextLocNo++ // CRITICAL FIX: advance to next location
	}
}

func (o *Object) noThe() bool {
	return (o.Flags & model.OfNoThe) != 0
}

func (o *Object) Place() {
	if o.minLocNo == 0 || o.maxLocNo == 0 {
		return
	}

	if o.minLocNo == o.maxLocNo {
		o.curLocNo = o.minLocNo
	} else {
		o.curLocNo = o.minLocNo + uint32(rand.Int32N(int32(o.maxLocNo-o.minLocNo))) //nolint:gosec // "It's Just A Game"
	}
	// debug.Check(cur_loc >= min_loc && cur_loc <= max_loc)

	// if !o.isValidPlacement() && o.nextLocation() == nil {
	// 	log.Printf("Unable to place %s", o.Name(false))
	// 	immobilize()
	// 	return
	// }

	if o.IsMobile() {
		if o.minLocNo != o.maxLocNo && o.maxCounter > 0 {
			duration := moveTimerPeriod * time.Duration(o.maxCounter)
			o.moveTimer = time.AfterFunc(duration, o.moveTimerHandler)
		}
	}
}

func (o *Object) Recycle() {
	// log.Printf("object.recycle(%s)", o.name)

	// Remove the object from its current location.
	o.curLocNo = 0

	// Clear any fights.
	if o.IsSpaceship() {
		for i := range o.ShipGuns {
			o.ShipGuns[i].opponent = 0
		}
	}

	// if (pstarHome->recycleHook(this) != hContinue) {
	// 	return;
	// }

	// Set a time until the object's reappearance.
	seconds := MIN_RECYCLE_SECONDS + rand.IntN(MAX_RECYCLE_SECONDS-MIN_RECYCLE_SECONDS) //nolint:gosec // "It's Just A Game"
	o.recycleTimer = time.AfterFunc(time.Duration(seconds)*time.Second, o.recycleTimerHandler)
}

func (o *Object) recycleTimerHandler() {
	global.Lock()
	defer global.Unlock()

	monitoring.RecycleTimerTickTotal.WithLabelValues(o.homeSystem.Name()).Inc()

	o.recycleTimerProc()
}

func (o *Object) recycleTimerProc() {
	// log.Printf("object.recycleTimerProc(%s)", o.name)

	o.recycleTimer = nil
	o.Place()
}

func (o *Object) SetCurLocNo(v uint32) {
	o.curLocNo = v
}

func (o *Object) SetDesc(v string) {
	o.desc = v
}

func (o *Object) SetScan(v string) {
	o.scan = v
}

func (o *Object) SetSystem(v Systemer) { // FIXME: stop doing this
	o.homeSystem = v
}

func (o *Object) Value() int32 {
	return o.value
}

func (o *Object) Unhide() {
	o.Flags &^= model.OfHidden
}

func (o *Object) Weight() uint32 {
	return o.weight
}
