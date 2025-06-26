package core

import (
	"github.com/nosborn/federation-1999/internal/text"
)

type Event struct { // Structure defines the action for an event
	Type uint16 // Type of event
	// bool           (*handler)( Player& );
	MessageNo text.MsgNum // Text output to user
	Field1    int16       // hull/str
	Field2    int16       // shield/sta
	Field3    int16       // engine/int
	Field4    int16       // tractor/dex
	Field7    int16       // change amount for type 3 - 6 events
	Field8    int16       // object number, +ve carried/-ve not carried
	NewLoc    uint16      // New location to move player/ship to
}
