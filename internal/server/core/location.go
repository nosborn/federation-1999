package core

import (
	"github.com/nosborn/federation-1999/internal/text"
)

type Location struct {
	Number                                   uint32 // Location number
	FullMsgNo/*fullMessageNo*/ text.MsgNum   // Full description
	BriefMsgNo/*briefMessageNo*/ text.MsgNum // Brief description
	Events                                   [2]uint16   // [0] = enter event, [1] = move in non-exit direction
	Flags                                    uint32      // Location flags - 32 bits
	MovTab                                   [13]uint16  // Movement table
	SysLoc                                   text.MsgNum // Message for non-movement
}
