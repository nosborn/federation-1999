package server

import (
	"log"

	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/core"
	"github.com/nosborn/federation-1999/internal/text"
)

type Event struct {
	Type      uint32
	Handler   func(p *Player) bool
	message   string
	messageNo text.MsgNum
	Field1    int32
	Field2    int32
	Field3    int32
	Field4    int32
	Field7    int32
	Field8    int32
	NewLoc    uint32
}

func NewCoreEvent(data *core.Event) *Event {
	e := &Event{
		Type:    uint32(data.Type),
		Handler: nil, // Handler functions need to be mapped separately
		Field1:  int32(data.Field1),
		Field2:  int32(data.Field2),
		Field3:  int32(data.Field3),
		Field4:  int32(data.Field4),
		Field7:  int32(data.Field7),
		Field8:  int32(data.Field8),
		NewLoc:  uint32(data.NewLoc),
	}
	if data.MessageNo > 0 {
		e.message = text.Msg(data.MessageNo)
	}
	return e
}

// NewEventFromDB converts a DBEvent from disk to an Event
func NewEventFromDB(dbEvent model.DBEvent) *Event {
	return &Event{
		Type:    uint32(dbEvent.Type),
		Handler: nil, // Handler functions need to be mapped separately
		message: text.CStringToString(dbEvent.Desc[:]),
		Field1:  int32(dbEvent.Field1),
		Field2:  int32(dbEvent.Field2),
		Field3:  int32(dbEvent.Field3),
		Field4:  int32(dbEvent.Field4),
		Field7:  int32(dbEvent.Field7),
		Field8:  int32(dbEvent.Field8),
		NewLoc:  uint32(dbEvent.NewLoc),
	}
}

func (e *Event) Destroy() {
	// nothing to do
}

func (e *Event) Message() string {
	if e.messageNo > 0 {
		return text.Msg(e.messageNo)
	}
	debug.Check(e.message != "")
	return e.message
}

func EventHandler(eventNo int, object *Object, player *Player) bool {
	// log.Printf("EventHandler: eventNo=%#v object=%#v", eventNo, object)

	if eventNo == 0 {
		return true
	}

	event := player.curSys.Events()[eventNo-1]

	switch event.Type {
	case 1:
		return xType1(player, event)
	case 3:
		return xType3(player, event)
	case 8: // Sol events only.
		if !object.HomeSystem().IsSol() {
			log.Fatal("eventHandler: type 8 outside Sol!")
		}
		switch eventNo {
		case 51: // Place event for paper.
			setLaunchCoordinates()
			return true
		case 52: // Place event for sargeur.
			object.Flags &^= model.OfDuke
			object.desc = text.Msg(text.Sargeur_Desc)
			object.scan = text.Msg(text.Sargeur_Scan)
			return true
		case 53: // Place event for share.
			setLaunchCode()
			return true
		case 54: // Place or drop event for black box.
			object.desc = text.Msg(text.BlackBox_Desc)
			object.scan = text.Msg(text.BlackBox_Scan)
			return true
		}
		return true
	case 9:
		return xType9(player, event)
	default:
		return true
	}
}

func XEventHandler(player *Player, eventNo int) bool {
	// log.Printf("XEventHandler: eventNo=%#v", eventNo)

	// if (eventNumber < 0) {
	//	log.Printf("EventHandler: got a negative event number!")
	//	return true
	// }

	event := player.curSys.Events()[eventNo-1]
	switch event.Type {
	case 1:
		return xType1(player, event)
	case 3:
		return xType3(player, event)
	case 8:
		if !player.IsInSolSystem() && !player.IsInHorsellSystem() {
			log.Fatal("EventHandler: type 8 on player planet!")
		}
		// return event.handler(player)
		return true // FIXME
	case 9:
		return xType9(player, event)
	default:
		log.Printf("XEventHandler: unexpected type %v", event.Type)
		return true
	}
}

// Handler for type 1 x_events. Change personal attributes.
func xType1(player *Player, event *Event) bool {
	// log.Printf("xType1: %#v", event)

	if player.IsInsideSpaceship() { // no events of this type in ships
		return true
	}

	player.Output(event.Message())
	player.Output("\n\n")

	player.Str.Cur = max(player.Str.Cur+event.Field1, 1)
	player.Str.Cur = min(player.Str.Cur, player.Str.Max)

	if player.Sta.Cur += event.Field2; player.Sta.Cur < 1 {
		player.Die()
		return false
	}
	player.Sta.Cur = min(player.Sta.Cur, player.Sta.Max)

	player.Int.Cur = max(player.Int.Cur+event.Field3, 1)
	player.Int.Cur = min(player.Int.Cur, player.Int.Max)

	player.Dex.Cur = max(player.Dex.Cur+event.Field4, 1)
	player.Dex.Cur = min(player.Dex.Cur, player.Dex.Max)

	if event.NewLoc != 0 {
		// Move to the new location.
		player.LocNo = event.NewLoc
		player.setLocation(player.LocNo)
		player.curLoc.Describe(player, LongDescription)

		// See if we moved them to somewhere non-spyable.
		// player.checkSpyers() -- FIXME
	}

	return false
}

// Handler for type 3 x_events. Persona attribute test for minimum value.
func xType3(player *Player, event *Event) bool {
	// log.Printf("xType3: %#v", event)

	var curField, maxField *int32
	eventField := int32(1)

	if player.IsInsideSpaceship() { // No events of this type in ships
		return true
	}

	switch {
	case event.Field1 > 0:
		curField = &player.Str.Cur
		maxField = &player.Str.Max
		eventField = event.Field1
	case event.Field2 > 0:
		curField = &player.Sta.Cur
		maxField = &player.Sta.Max
		eventField = event.Field2
	case event.Field3 > 0:
		curField = &player.Int.Cur
		maxField = &player.Int.Max
		eventField = event.Field3
	case event.Field4 > 0:
		curField = &player.Dex.Cur
		maxField = &player.Dex.Max
		eventField = event.Field4
	}

	if curField == nil { // Duff event - don't process it
		return true
	}

	if *curField < eventField {
		*curField += event.Field7

		player.Output(event.Message())
		player.Output("\n\n")

		if *curField > *maxField {
			*curField = *maxField
		}

		if player.Sta.Cur < 1 { // Check for dying...
			player.Die()
			return false
		}

		if *curField < 1 {
			*curField = 1
		}

		if event.NewLoc != 0 {
			player.LocNo = event.NewLoc
			player.setLocation(player.LocNo)
			player.curLoc.Describe(player, LongDescription)
		}

		return false
	}

	return true // passes test if it gets to this point!
}

// func xType8(player *Player, eventNo int) bool {
// 	log.Printf("xType8: eventNo=%v", eventNo)
//
// 	if !player.IsInSolSystem() && !player.IsInHorsellSystem() {
// 		log.Fatal("x_event_handler: Type 8 on player planet!")
// 		// TODO
// 	}
//
// 	switch eventNo {
// 	case 51: // place event for paper
// 		// TODO
// 		return true
// 	case 52: // place event for sargeur
// 		// TODO
// 		return true
// 	case 53: // place event for share
// 		// TODO
// 		return true
// 	case 54: // place or drop event for black box
// 		// TODO
// 		return true
// 	}
//
// 	return true // passes test if it gets to this point!
// }

// Handler for type 9 x_events. Test whether player is carrying/not carrying an
// object.
func xType9(player *Player, event *Event) bool {
	log.Printf("xType9: %v", event)

	objNo := uint32(max(event.Field8, -event.Field8))
	var obj *Object
	for _, o := range player.inventory {
		if o.Number() == objNo {
			obj = o
			break
		}
	}
	if event.Field8 > 0 && obj == nil {
		return true
	}
	if event.Field8 < 0 && obj != nil {
		return true
	}
	return xType1(player, event)
}

// See's if the player is strong enough to lift the barbells.
// Player can peg out if not!
// func increaseStrength(p *Player) bool {
// 	if p.Rank() < model.RankAdventurer {
// 		p.Sta.Cur = min(p.Sta.Max, 6)
// 		p.Outputm(text.MN120)
// 		return false
// 	}
//
// 	if p.Str.Max >= 120 {
// 		p.Outputm(text.DoneStrengthPuzzle)
// 		return false
// 	}
//
// 	if p.Sta.Cur -= 50; p.Sta.Cur < 1 {
// 		p.Outputm(text.MN122)
// 		p.Die()
// 		return false
// 	}
//
// 	if p.Balance() < STRENGTH_COST {
// 		p.Outputm(text.MN124, STRENGTH_COST)
// 		return false
// 	}
//
// 	p.Str.Max = min(p.Str.Max+4, 120)
// 	p.Str.Cur = min(p.Str.Cur+4, p.Str.Max)
// 	p.Flags1 |= model.PL1_DONE_STR
//
// 	p.ChangeBalance(-STRENGTH_COST)
//
// 	p.Outputm(text.MN125, STRENGTH_COST)
// 	p.Save(database.SaveNow)
//
// 	// Yes, people can accumulate enough stat points to promote without
// 	// having done -all- of the puzzles!
// 	p.CheckForPromotion()
//
// 	return false
// }

// func restoreDogMovementTable(p *Player) bool {
// 	SolSystem.FindLocation(sol.Backyard).MovTab[model.DirectionDown] = 0
// 	return true
// }

// func restoreNisrikMovementTable(p *Player) bool {
// 	SolSystem.FindLocation(sol.SecurityArea).MovTab[model.DirectionWest] = 0
// 	return true
// }

// func silentSapStamina(p *Player) bool {
// 	if p.Sta.Cur -= 10; p.Sta.Cur < 1 {
// 		p.Die()
// 		return false
// 	}
// 	return true
// }

// Checks to ee whether the player has custom clothes before
// allowing access to Chez Diesel.
// func _DressCode(p *Player) bool {
// 	if p.IsDressed() {
// 		return true
// 	}
// 	p.Outputm(text.DRESS_CODE)
// 	p.LocNo = 246 /* SolSystem::TuxDeluxe */
// 	p.setLocation(246 /* SolSystem::TuxDeluxe */)
// 	p.curLoc.Describe(p, LongDescription)
// 	return false
// }
