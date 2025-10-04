package snark

import (
	"math"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/core"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

var Events = [8]core.Event{
	// 1: ?
	{
		Type: 8,
		// Handler: snarkSapStamina,
	},
	// 2: ?
	{
		Type: 8,
		// Handler: silentSapStamina,
	},
	// 3: ?
	{
		Type: 8,
		// Handler: snarkSiloDamage,
	},
	// 4: ?
	{
		Type:      1,
		MessageNo: text.SolEvent33,
		Field2:    -65,
		NewLoc:    sol.LinesRoom,
	},
	// 5: ?
	{
		Type:      9,
		MessageNo: text.SolEvent42,
		Field2:    -999,
		Field8:    -sol.ObBadge,
	},
	// 6: ?
	{
		Type:      1,
		MessageNo: text.SolEvent46,
		Field2:    -199,
	},
	// 7: Obsolete, was consume event for potion
	{
		Type: 8,
	},
	// 8: Place event for paper.
	{
		Type: 8,
	},
}

var Locations = [86]core.Location{
	{
		Number:     Corridor1,
		BriefMsgNo: text.SnarkLocation9_Brief,
		FullMsgNo:  text.SnarkLocation9_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			Corridor2,  // SE
			0,          // S
			0,          // SW
			Laboratory, // W
			0,          // NW
			0,          // Up
			0,          // Down
			Laboratory, // In
			0,          // Out
		},
	},
	{
		Number:     Corridor2,
		BriefMsgNo: text.SnarkLocation10_Brief,
		FullMsgNo:  text.SnarkLocation10_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			Octagon,   // SE
			0,         // S
			0,         // SW
			0,         // W
			Corridor1, // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     Corridor3,
		BriefMsgNo: text.SnarkLocation11_Brief,
		FullMsgNo:  text.SnarkLocation11_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,       // N
			Bar,     // NE
			0,       // E
			0,       // SE
			0,       // S
			Octagon, // SW
			0,       // W
			0,       // NW
			0,       // Up
			0,       // Down
			Bar,     // In
			0,       // Out
		},
	},
	{
		Number:     Octagon,
		BriefMsgNo: text.SnarkLocation12_Brief,
		FullMsgNo:  text.SnarkLocation12_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,         // N
			Corridor3, // NE
			0,         // E
			Corridor4, // SE
			Lift1,     // S
			0,         // SW
			0,         // W
			Corridor2, // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     Corridor4,
		BriefMsgNo: text.SnarkLocation13_Brief,
		FullMsgNo:  text.SnarkLocation13_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			Corridor5, // SE
			0,         // S
			0,         // SW
			0,         // W
			Octagon,   // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     Corridor5,
		BriefMsgNo: text.SnarkLocation14_Brief,
		FullMsgNo:  text.SnarkLocation14_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,               // N
			0,               // NE
			TransferControl, // E
			0,               // SE
			MissileRoom,     // S
			0,               // SW
			0,               // W
			Corridor4,       // NW
			0,               // Up
			0,               // Down
			0,               // In
			0,               // Out
		},
	},
	{
		Number:     Corridor6,
		BriefMsgNo: text.SnarkLocation15_Brief,
		Events:     [2]uint16{1, 2},
		FullMsgNo:  text.SnarkLocation15_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,          // N
			Laboratory, // NE
			0,          // E
			0,          // SE
			Corridor7,  // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			Laboratory, // Out
		},
	},
	{
		Number:     Corridor7,
		BriefMsgNo: text.SnarkLocation16_Brief,
		Events:     [2]uint16{1, 2},
		FullMsgNo:  text.SnarkLocation16_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			Corridor6,    // N
			0,            // NE
			0,            // E
			0,            // SE
			WideCorridor, // S
			0,            // SW
			0,            // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     WideCorridor,
		BriefMsgNo: text.SnarkLocation17_Brief,
		Events:     [2]uint16{1, 2},
		FullMsgNo:  text.SnarkLocation17_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			Corridor6,      // N
			0,              // NE
			ReactorControl, // E
			0,              // SE
			0,              // S
			Corridor8,      // SW
			0,              // W
			0,              // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     Corridor8,
		BriefMsgNo: text.SnarkLocation18_Brief,
		FullMsgNo:  text.SnarkLocation18_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			SecurityRoom, // N
			WideCorridor, // NE
			0,            // E
			0,            // SE
			0,            // S
			0,            // SW
			Cave1,        // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     Cave1,
		BriefMsgNo: text.SnarkLocation19_Brief,
		FullMsgNo:  text.SnarkLocation19_Full,
		Flags:      model.LfDark | model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			Corridor8, // E
			0,         // SE
			0,         // S
			0,         // SW
			0,         // W
			Tunnel,    // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     Tunnel,
		BriefMsgNo: text.SnarkLocation20_Brief,
		FullMsgNo:  text.SnarkLocation20_Full,
		Flags:      model.LfDark | model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,     // N
			0,     // NE
			0,     // E
			Cave1, // SE
			0,     // S
			0,     // SW
			Shaft, // W
			0,     // NW
			0,     // Up
			Shaft, // Down
			0,     // In
			Cave1, // Out
		},
	},
	{
		Number:     Shaft,
		BriefMsgNo: text.SnarkLocation21_Brief,
		FullMsgNo:  text.SnarkLocation21_Full,
		Flags:      model.LfDark | model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			0,        // E
			0,        // SE
			MineFace, // S
			0,        // SW
			0,        // W
			0,        // NW
			Tunnel,   // Up
			0,        // Down
			MineFace, // In
			Tunnel,   // Out
		},
	},
	{
		Number:     MineFace,
		BriefMsgNo: text.SnarkLocation22_Brief,
		FullMsgNo:  text.SnarkLocation22_Full,
		Flags:      model.LfDark | model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			Shaft, // N
			0,     // NE
			0,     // E
			0,     // SE
			0,     // S
			0,     // SW
			0,     // W
			0,     // NW
			0,     // Up
			0,     // Down
			0,     // In
			Shaft, // Out
		},
	},
	{
		Number:     PackingRoom,
		BriefMsgNo: text.SnarkLocation23_Brief,
		FullMsgNo:  text.SnarkLocation23_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			Corridor1, // E
			0,         // SE
			0,         // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			Corridor1, // Out
		},
	},
	{
		Number:     Laboratory,
		BriefMsgNo: text.SnarkLocation24_Brief,
		Events:     [2]uint16{5, 0},
		FullMsgNo:  text.SnarkLocation24_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			Corridor1, // E
			0,         // SE
			0,         // S
			Corridor6, // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			Corridor1, // Out
		},
	},
	{
		Number:     Bar,
		BriefMsgNo: text.SnarkLocation25_Brief,
		FullMsgNo:  text.SnarkLocation25_Full,
		Flags:      model.LfCafe | model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			0,         // SE
			0,         // S
			Corridor3, // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     Office,
		BriefMsgNo: text.SnarkLocation26_Brief,
		FullMsgNo:  text.SnarkLocation26_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,   // N
			0,   // NE
			0,   // E
			0,   // SE
			0,   // S
			0,   // SW
			0,   // W
			Bar, // NW
			0,   // Up
			0,   // Down
			0,   // In
			Bar, // Out
		},
	},
	{
		Number:     Lift1,
		BriefMsgNo: text.SnarkLocation27_Brief,
		FullMsgNo:  text.SnarkLocation27_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			Octagon, // N
			0,       // NE
			0,       // E
			0,       // SE
			0,       // S
			0,       // SW
			0,       // W
			0,       // NW
			Lift2,   // Up
			0,       // Down
			0,       // In
			Octagon, // Out
		},
	},
	{
		Number:     SeptilateralRoom,
		BriefMsgNo: text.SnarkLocation28_Brief,
		Events:     [2]uint16{6, 0},
		FullMsgNo:  text.SnarkLocation28_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			0,            // E
			0,            // SE
			SecurityRoom, // S
			0,            // SW
			0,            // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			SecurityRoom, // Out
		},
	},
	{
		Number:     CannonRoom,
		BriefMsgNo: text.SnarkLocation29_Brief,
		FullMsgNo:  text.SnarkLocation29_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,               // N
			0,               // NE
			TransferControl, // E
			0,               // SE
			0,               // S
			0,               // SW
			Corridor5,       // W
			0,               // NW
			0,               // Up
			0,               // Down
			0,               // In
			0,               // Out
		},
	},
	{
		Number:     TransferControl,
		BriefMsgNo: text.SnarkLocation30_Brief,
		FullMsgNo:  text.SnarkLocation30_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			Corridor5,  // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			CannonRoom, // Out
		},
	},
	{
		Number:     MissileRoom,
		BriefMsgNo: text.SnarkLocation31_Brief,
		FullMsgNo:  text.SnarkLocation31_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			Corridor5, // N
			0,         // NE
			0,         // E
			0,         // SE
			0,         // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			Corridor5, // Out
		},
	},
	{
		Number:     SecurityRoom,
		BriefMsgNo: text.SnarkLocation32_Brief,
		FullMsgNo:  text.SnarkLocation32_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			SeptilateralRoom, // N
			0,                // NE
			0,                // E
			0,                // SE
			Corridor8,        // S
			0,                // SW
			0,                // W
			0,                // NW
			0,                // Up
			0,                // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     ReactorControl,
		BriefMsgNo: text.SnarkLocation33_Brief,
		FullMsgNo:  text.SnarkLocation33_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			0,            // E
			0,            // SE
			ReactorRoom,  // S
			0,            // SW
			WideCorridor, // W
			0,            // NW
			0,            // Up
			ReactorRoom,  // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     ReactorRoom,
		BriefMsgNo: text.SnarkLocation34_Brief,
		Events:     [2]uint16{2, 2},
		FullMsgNo:  text.SnarkLocation34_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			ReactorControl, // N
			0,              // NE
			0,              // E
			0,              // SE
			0,              // S
			0,              // SW
			0,              // W
			0,              // NW
			ReactorControl, // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     LivingQuarters,
		BriefMsgNo: text.SnarkLocation35_Brief,
		FullMsgNo:  text.SnarkLocation35_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			OpenSpace, // SE
			0,         // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			OpenSpace, // Out
		},
	},
	{
		Number:     TransmitterRoom,
		BriefMsgNo: text.SnarkLocation36_Brief,
		FullMsgNo:  text.SnarkLocation36_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			0,         // SE
			OpenSpace, // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			OpenSpace, // Out
		},
	},
	{
		Number:     ComputerRoom,
		BriefMsgNo: text.SnarkLocation37_Brief,
		FullMsgNo:  text.SnarkLocation37_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			0,         // SE
			0,         // S
			OpenSpace, // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			OpenSpace, // Out
		},
	},
	{
		Number:     Storeroom,
		BriefMsgNo: text.SnarkLocation38_Brief,
		FullMsgNo:  text.SnarkLocation38_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			0,         // SE
			0,         // S
			Corridor9, // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			Corridor9, // Out
		},
	},
	{
		Number:     Galley,
		BriefMsgNo: text.SnarkLocation39_Brief,
		FullMsgNo:  text.SnarkLocation39_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			OpenSpace, // E
			0,         // SE
			0,         // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			OpenSpace, // Out
		},
	},
	{
		Number:     OpenSpace,
		BriefMsgNo: text.SnarkLocation40_Brief,
		FullMsgNo:  text.SnarkLocation40_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			TransmitterRoom, // N
			ComputerRoom,    // NE
			Corridor9,       // E
			NarrowPassage,   // SE
			Workshop,        // S
			Passage,         // SW
			Galley,          // W
			LivingQuarters,  // NW
			0,               // Up
			0,               // Down
			0,               // In
			0,               // Out
		},
	},
	{
		Number:     Corridor9,
		BriefMsgNo: text.SnarkLocation41_Brief,
		FullMsgNo:  text.SnarkLocation41_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,         // N
			Storeroom, // NE
			0,         // E
			HoloRoom,  // SE
			0,         // S
			0,         // SW
			OpenSpace, // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     Passage,
		BriefMsgNo: text.SnarkLocation42_Brief,
		FullMsgNo:  text.SnarkLocation42_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,         // N
			OpenSpace, // NE
			0,         // E
			0,         // SE
			0,         // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     Workshop,
		BriefMsgNo: text.SnarkLocation43_Brief,
		FullMsgNo:  text.SnarkLocation43_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			OpenSpace, // N
			0,         // NE
			0,         // E
			0,         // SE
			0,         // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			OpenSpace, // Out
		},
	},
	{
		Number:     NarrowPassage,
		BriefMsgNo: text.SnarkLocation44_Brief,
		FullMsgNo:  text.SnarkLocation44_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			0,         // SE
			Cave2,     // S
			0,         // SW
			0,         // W
			OpenSpace, // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     HoloRoom,
		BriefMsgNo: text.SnarkLocation45_Brief,
		FullMsgNo:  text.SnarkLocation45_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			0,         // SE
			0,         // S
			0,         // SW
			0,         // W
			Corridor9, // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     MirrorRoom,
		BriefMsgNo: text.SnarkLocation46_Brief,
		FullMsgNo:  text.SnarkLocation46_Full,
		Flags:      model.LfDeath | model.LfLock | model.LfShield,
	},
	{
		Number:     Cave2,
		BriefMsgNo: text.SnarkLocation47_Brief,
		FullMsgNo:  text.SnarkLocation47_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			0,             // E
			0,             // SE
			EnergyCurtain, // S
			0,             // SW
			0,             // W
			0,             // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     EnergyCurtain,
		BriefMsgNo: text.SnarkLocation48_Brief,
		Events:     [2]uint16{2, 2},
		FullMsgNo:  text.SnarkLocation48_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			Cave2, // N
			0,     // NE
			0,     // E
			0,     // SE
			Cave3, // S
			0,     // SW
			0,     // W
			0,     // NW
			0,     // Up
			0,     // Down
			Cave2, // In
			Cave3, // Out
		},
	},
	{
		Number:     CannonControl,
		BriefMsgNo: text.SnarkLocation49_Brief,
		FullMsgNo:  text.SnarkLocation49_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,              // N
			0,              // NE
			MissileControl, // E
			0,              // SE
			BunkerEntrance, // S
			0,              // SW
			0,              // W
			0,              // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     MissileControl,
		BriefMsgNo: text.SnarkLocation50_Brief,
		FullMsgNo:  text.SnarkLocation50_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			0,             // E
			0,             // SE
			0,             // S
			0,             // SW
			CannonControl, // W
			0,             // NW
			0,             // Up
			0,             // Down
			0,             // In
			CannonControl, // Out
		},
	},
	{
		Number:     BunkerEntrance,
		BriefMsgNo: text.SnarkLocation51_Brief,
		FullMsgNo:  text.SnarkLocation51_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			CannonControl, // N
			0,             // NE
			0,             // E
			0,             // SE
			Blockhouse,    // S
			0,             // SW
			0,             // W
			0,             // NW
			0,             // Up
			0,             // Down
			CannonControl, // In
			Blockhouse,    // Out
		},
	},
	{
		Number:     Cave3,
		BriefMsgNo: text.SnarkLocation52_Brief,
		FullMsgNo:  text.SnarkLocation52_Full,
		Flags:      model.LfIndoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			EnergyCurtain, // N
			0,             // NE
			0,             // E
			0,             // SE
			Ravine1,       // S
			0,             // SW
			0,             // W
			0,             // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     Ravine1,
		BriefMsgNo: text.SnarkLocation53_Brief,
		FullMsgNo:  text.SnarkLocation53_Full,
		Flags:      model.LfDark | model.LfOutdoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			Cave3,   // N
			0,       // NE
			0,       // E
			0,       // SE
			0,       // S
			Ravine4, // SW
			0,       // W
			0,       // NW
			0,       // Up
			0,       // Down
			Cave3,   // In
			0,       // Out
		},
	},
	{
		Number:     Ravine2,
		BriefMsgNo: text.SnarkLocation54_Brief,
		FullMsgNo:  text.SnarkLocation54_Full,
		Flags:      model.LfOutdoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			Ravine3,  // E
			0,        // SE
			Surface1, // S
			0,        // SW
			0,        // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     Ravine3,
		BriefMsgNo: text.SnarkLocation55_Brief,
		FullMsgNo:  text.SnarkLocation55_Full,
		Flags:      model.LfOutdoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			Ravine4, // E
			0,       // SE
			0,       // S
			0,       // SW
			Ravine2, // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Ravine4,
		BriefMsgNo: text.SnarkLocation56_Brief,
		FullMsgNo:  text.SnarkLocation56_Full,
		Flags:      model.LfOutdoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			0,       // N
			Ravine1, // NE
			0,       // E
			0,       // SE
			0,       // S
			0,       // SW
			Ravine3, // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Ravine5,
		BriefMsgNo: text.SnarkLocation57_Brief,
		FullMsgNo:  text.SnarkLocation57_Full,
		Flags:      model.LfOutdoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			0,          // N
			Ravine2,    // NE
			Surface1,   // E
			Surface4,   // SE
			SpoilHeap1, // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Surface1,
		BriefMsgNo: text.SnarkLocation58_Brief,
		FullMsgNo:  text.SnarkLocation58_Full,
		Flags:      model.LfOutdoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			Ravine2,  // N
			0,        // NE
			Surface2, // E
			0,        // SE
			Surface4, // S
			0,        // SW
			0,        // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     Surface2,
		BriefMsgNo: text.SnarkLocation59_Brief,
		FullMsgNo:  text.SnarkLocation59_Full,
		Flags:      model.LfOutdoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			Surface3, // E
			Surface5, // SE
			0,        // S
			Surface4, // SW
			Surface1, // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     Surface3,
		BriefMsgNo: text.SnarkLocation60_Brief,
		FullMsgNo:  text.SnarkLocation60_Full,
		Flags:      model.LfOutdoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			SpoilHeap2, // E
			Path1,      // SE
			Surface5,   // S
			0,          // SW
			Surface2,   // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Surface4,
		BriefMsgNo: text.SnarkLocation61_Brief,
		FullMsgNo:  text.SnarkLocation61_Full,
		Flags:      model.LfOutdoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			Surface1,     // N
			Surface2,     // NE
			0,            // E
			Surface7,     // SE
			Surface6,     // S
			MissileSilos, // SW
			0,            // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     Surface5,
		BriefMsgNo: text.SnarkLocation62_Brief,
		FullMsgNo:  text.SnarkLocation62_Full,
		Flags:      model.LfOutdoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			Surface3,   // N
			SpoilHeap2, // NE
			Path1,      // E
			0,          // SE
			Surface8,   // S
			Surface7,   // SW
			0,          // W
			Surface2,   // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Surface6,
		BriefMsgNo: text.SnarkLocation63_Brief,
		FullMsgNo:  text.SnarkLocation63_Full,
		Flags:      model.LfOutdoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			Surface4,     // N
			0,            // NE
			Surface7,     // E
			0,            // SE
			0,            // S
			0,            // SW
			MissileSilos, // W
			0,            // NW
			Surface4,     // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     Surface7,
		BriefMsgNo: text.SnarkLocation64_Brief,
		FullMsgNo:  text.SnarkLocation64_Full,
		Flags:      model.LfOutdoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			Airlock,  // N
			Surface5, // NE
			Surface8, // E
			Silos,    // SE
			0,        // S
			0,        // SW
			Surface6, // W
			Surface4, // NW
			0,        // Up
			0,        // Down
			Airlock,  // In
			0,        // Out
		},
	},
	{
		Number:     Surface8,
		BriefMsgNo: text.SnarkLocation65_Brief,
		FullMsgNo:  text.SnarkLocation65_Full,
		Flags:      model.LfOutdoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			Surface5, // N
			Path1,    // NE
			Pit,      // E
			0,        // SE
			Silos,    // S
			0,        // SW
			Surface7, // W
			0,        // NW
			0,        // Up
			Pit,      // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     SpoilHeap1,
		BriefMsgNo: text.SnarkLocation66_Brief,
		FullMsgNo:  text.SnarkLocation66_Full,
		Flags:      model.LfOutdoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			Ravine5,  // N
			Surface1, // NE
			Surface4, // E
			Surface6, // SE
			0,        // S
			0,        // SW
			0,        // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     SpoilHeap2,
		BriefMsgNo: text.SnarkLocation67_Brief,
		FullMsgNo:  text.SnarkLocation67_Full,
		Flags:      model.LfOutdoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			0,        // E
			Path2,    // SE
			Path1,    // S
			Surface5, // SW
			Surface3, // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     SmallSpoilHeap,
		BriefMsgNo: text.SnarkLocation68_Brief,
		FullMsgNo:  text.SnarkLocation68_Full,
		Flags:      model.LfOutdoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			Surface6, // N
			0,        // NE
			0,        // E
			0,        // SE
			0,        // S
			0,        // SW
			0,        // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     MissileSilos,
		BriefMsgNo: text.SnarkLocation69_Brief,
		FullMsgNo:  text.SnarkLocation69_Full,
		Flags:      model.LfOutdoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			0,        // N
			Surface4, // NE
			Surface6, // E
			0,        // SE
			0,        // S
			0,        // SW
			0,        // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     Airlock,
		BriefMsgNo: text.SnarkLocation70_Brief,
		FullMsgNo:  text.SnarkLocation70_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			Lift2,    // N
			0,        // NE
			0,        // E
			0,        // SE
			Surface7, // S
			0,        // SW
			0,        // W
			0,        // NW
			0,        // Up
			0,        // Down
			Lift2,    // In
			Surface7, // Out
		},
	},
	{
		Number:     Path1,
		BriefMsgNo: text.SnarkLocation71_Brief,
		FullMsgNo:  text.SnarkLocation71_Full,
		Flags:      model.LfOutdoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			SpoilHeap2, // N
			0,          // NE
			Path2,      // E
			0,          // SE
			0,          // S
			Surface8,   // SW
			Surface5,   // W
			Surface3,   // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Path2,
		BriefMsgNo: text.SnarkLocation72_Brief,
		FullMsgNo:  text.SnarkLocation72_Full,
		Flags:      model.LfOutdoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Path3,      // E
			0,          // SE
			0,          // S
			0,          // SW
			Path1,      // W
			SpoilHeap2, // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Path3,
		BriefMsgNo: text.SnarkLocation73_Brief,
		FullMsgNo:  text.SnarkLocation73_Full,
		Flags:      model.LfOutdoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			Blockhouse,  // E
			0,           // SE
			SensorArray, // S
			0,           // SW
			Path2,       // W
			0,           // NW
			0,           // Up
			0,           // Down
			Blockhouse,  // In
			0,           // Out
		},
	},
	{
		Number:     Pit,
		BriefMsgNo: text.SnarkLocation74_Brief,
		Events:     [2]uint16{2, 2},
		FullMsgNo:  text.SnarkLocation74_Full,
		Flags:      model.LfOutdoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			0,        // E
			0,        // SE
			0,        // S
			0,        // SW
			Surface8, // W
			0,        // NW
			Surface8, // Up
			0,        // Down
			0,        // In
			Surface8, // Out
		},
	},
	{
		Number:     Blockhouse,
		BriefMsgNo: text.SnarkLocation75_Brief,
		FullMsgNo:  text.SnarkLocation75_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,     // N
			0,     // NE
			0,     // E
			0,     // SE
			0,     // S
			0,     // SW
			Path3, // W
			0,     // NW
			0,     // Up
			0,     // Down
			0,     // In
			Path3, // Out
		},
	},
	{
		Number:     SensorArray,
		BriefMsgNo: text.SnarkLocation76_Brief,
		FullMsgNo:  text.SnarkLocation76_Full,
		Flags:      model.LfOutdoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			Path3,          // N
			0,              // NE
			ParticleCannon, // E
			0,              // SE
			LandingArea,    // S
			0,              // SW
			0,              // W
			0,              // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     ParticleCannon,
		BriefMsgNo: text.SnarkLocation77_Brief,
		FullMsgNo:  text.SnarkLocation77_Full,
		Flags:      model.LfOutdoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			0,           // E
			0,           // SE
			0,           // S
			0,           // SW
			SensorArray, // W
			Path3,       // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     Lift2,
		BriefMsgNo: text.SnarkLocation78_Brief,
		FullMsgNo:  text.SnarkLocation78_Full,
		Flags:      model.LfIndoors | model.LfShield,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			0,       // E
			0,       // SE
			Airlock, // S
			0,       // SW
			0,       // W
			0,       // NW
			0,       // Up
			Lift1,   // Down
			0,       // In
			Airlock, // Out
		},
	},
	{
		Number:     LandingArea,
		BriefMsgNo: text.SnarkLocation79_Brief,
		FullMsgNo:  text.SnarkLocation79_Full,
		Flags:      model.LfLanding | model.LfOutdoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			SensorArray, // N
			0,           // NE
			0,           // E
			0,           // SE
			0,           // S
			0,           // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			SensorArray, // Out
		},
	},
	{
		Number:     Silos,
		BriefMsgNo: text.SnarkLocation80_Brief,
		Events:     [2]uint16{3, 3},
		FullMsgNo:  text.SnarkLocation80_Full,
		Flags:      model.LfOutdoors | model.LfShield | model.LfVacuum,

		MovTab: [13]uint16{
			Surface8, // N
			0,        // NE
			0,        // E
			0,        // SE
			0,        // S
			0,        // SW
			0,        // W
			Surface7, // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     SnarkOrbit,
		BriefMsgNo: text.SnarkSnarkOrbit_Brief,
		FullMsgNo:  text.SnarkSnarkOrbit_Full,
		Flags:      model.LfOrbit | model.LfSpace,

		MovTab: [13]uint16{
			HilbertSpace6,  // N
			HilbertSpace7,  // NE
			HilbertSpace10, // E
			HilbertSpace4,  // SE
			HilbertSpace3,  // S
			HilbertSpace5,  // SW
			HilbertSpace2,  // W
			HilbertSpace12, // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     HilbertSpace1,
		BriefMsgNo: text.SnarkLocation82_Brief,
		FullMsgNo:  text.SnarkLocation82_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			HilbertSpace11, // N
			HilbertSpace12, // NE
			HilbertSpace2,  // E
			HilbertSpace3,  // SE
			HilbertSpace5,  // S
			HilbertSpace4,  // SW
			HilbertSpace10, // W
			HilbertSpace7,  // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     HilbertSpace2,
		BriefMsgNo: text.SnarkLocation83_Brief,
		FullMsgNo:  text.SnarkLocation83_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			HilbertSpace12, // N
			HilbertSpace10, // NE
			SnarkOrbit,     // E
			HilbertSpace4,  // SE
			HilbertSpace3,  // S
			HilbertSpace5,  // SW
			HilbertSpace1,  // W
			HilbertSpace11, // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     HilbertSpace3,
		BriefMsgNo: text.SnarkLocation84_Brief,
		FullMsgNo:  text.SnarkLocation84_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			SnarkOrbit,     // N
			HilbertSpace10, // NE
			HilbertSpace4,  // E
			HilbertSpace7,  // SE
			HilbertSpace6,  // S
			HilbertSpace9,  // SW
			HilbertSpace5,  // W
			HilbertSpace2,  // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     HilbertSpace4,
		BriefMsgNo: text.SnarkLocation85_Brief,
		FullMsgNo:  text.SnarkLocation85_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			HilbertSpace10, // N
			HilbertSpace1,  // NE
			HilbertSpace5,  // E
			HilbertSpace8,  // SE
			HilbertSpace7,  // S
			HilbertSpace6,  // SW
			HilbertSpace3,  // W
			SnarkOrbit,     // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     HilbertSpace5,
		BriefMsgNo: text.SnarkLocation86_Brief,
		FullMsgNo:  text.SnarkLocation86_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			HilbertSpace1,  // N
			HilbertSpace2,  // NE
			HilbertSpace3,  // E
			HilbertSpace9,  // SE
			HilbertSpace8,  // S
			HilbertSpace7,  // SW
			HilbertSpace4,  // W
			HilbertSpace10, // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     HilbertSpace6,
		BriefMsgNo: text.SnarkLocation87_Brief,
		FullMsgNo:  text.SnarkLocation87_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			HilbertSpace3,  // N
			HilbertSpace4,  // NE
			HilbertSpace7,  // E
			HilbertSpace10, // SE
			SnarkOrbit,     // S
			HilbertSpace1,  // SW
			HilbertSpace9,  // W
			HilbertSpace5,  // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     HilbertSpace7,
		BriefMsgNo: text.SnarkLocation88_Brief,
		FullMsgNo:  text.SnarkLocation88_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			HilbertSpace4,  // N
			HilbertSpace5,  // NE
			HilbertSpace8,  // E
			HilbertSpace11, // SE
			HilbertSpace10, // S
			SnarkOrbit,     // SW
			HilbertSpace6,  // W
			HilbertSpace3,  // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     HilbertSpace8,
		BriefMsgNo: text.SnarkLocation89_Brief,
		FullMsgNo:  text.SnarkLocation89_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			HilbertSpace5,  // N
			HilbertSpace3,  // NE
			HilbertSpace9,  // E
			HilbertSpace12, // SE
			HilbertSpace11, // S
			HilbertSpace10, // SW
			HilbertSpace7,  // W
			HilbertSpace4,  // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     HilbertSpace9,
		BriefMsgNo: text.SnarkLocation90_Brief,
		FullMsgNo:  text.SnarkLocation90_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			HilbertSpace3,  // N
			HilbertSpace4,  // NE
			HilbertSpace6,  // E
			HilbertSpace10, // SE
			HilbertSpace12, // S
			HilbertSpace11, // SW
			HilbertSpace8,  // W
			HilbertSpace5,  // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     HilbertSpace10,
		BriefMsgNo: text.SnarkLocation91_Brief,
		FullMsgNo:  text.SnarkLocation91_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			HilbertSpace7,  // N
			HilbertSpace8,  // NE
			HilbertSpace11, // E
			HilbertSpace1,  // SE
			HilbertSpace4,  // S
			HilbertSpace2,  // SW
			SnarkOrbit,     // W
			HilbertSpace6,  // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     HilbertSpace11,
		BriefMsgNo: text.SnarkLocation92_Brief,
		FullMsgNo:  text.SnarkLocation92_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			HilbertSpace8,  // N
			HilbertSpace9,  // NE
			HilbertSpace12, // E
			HilbertSpace2,  // SE
			HilbertSpace1,  // S
			HilbertSpace4,  // SW
			HilbertSpace10, // W
			HilbertSpace7,  // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     HilbertSpace12,
		BriefMsgNo: text.SnarkLocation93_Brief,
		FullMsgNo:  text.SnarkLocation93_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			HilbertSpace9,  // N
			HilbertSpace6,  // NE
			HilbertSpace10, // E
			SnarkOrbit,     // SE
			HilbertSpace2,  // S
			HilbertSpace1,  // SW
			HilbertSpace11, // W
			HilbertSpace8,  // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     OutThroughTheLookingGlass,
		BriefMsgNo: text.SnarkLocation94_Brief,
		Events:     [2]uint16{4, 0},
		FullMsgNo:  text.SnarkLocation94_Full,
		Flags:      model.LfHidden | model.LfLock | model.LfShield,
	},
}

var Objects = [10]core.Object{
	{
		AttackPercent: 100,
		DescMessageNo: text.Care_Desc,
		Flags:         model.OfAnimate | model.OfShip,
		MaxCounter:    6,
		MaxLoc:        HilbertSpace6,
		MinLoc:        HilbertSpace1,
		Name:          "Care satellite",
		Number:        ObCare,
		ScanMessageNo: text.Care_Scan,
		Sex:           'n',
		Synonyms:      []string{"Care"},
		Value:         61000,

		ShipGuns: [4]model.OldSGuns{
			core.SHIP_GUN_QUAD_LASER,
			core.SHIP_GUN_QUAD_LASER,
			core.SHIP_GUN_TWIN_LASER,
			core.SHIP_GUN_LASER,
		},

		ShipKit: core.Equipment{
			Computer: 5,
			Engine:   40,
			Fuel:     500,
			Hold:     45,
			Hull:     60,
			Shield:   40,
			Tonnage:  600,
		},
	},
	{
		AttackPercent: 50,
		DescMessageNo: text.Caretaker_Desc,
		Flags:         model.OfAnimate | model.OfCleaner | model.OfIndoors,
		MaxCounter:    4,
		MaxLoc:        Silos,
		MinLoc:        Corridor1,
		Name:          "caretaker",
		Number:        ObCaretaker,
		ScanMessageNo: text.Caretaker_Scan,
		Sex:           'm',
	},
	{
		DescMessageNo: text.Converter_Desc,
		Flags:         model.OfHidden,
		Name:          "n-space converter",
		Number:        ObConverter,
		ScanMessageNo: text.Converter_Scan,
		Sex:           'n',
		Synonyms:      []string{"converter"},
		Value:         1000,
		Weight:        4,
	},
	{
		AttackPercent: 50,
		DescMessageNo: text.HilbertWind_Desc,
		Flags:         model.OfAnimate | model.OfOutdoors,
		MaxCounter:    4,
		MaxLoc:        Silos,
		MinLoc:        Corridor1,
		Name:          "Hilbert wind",
		Number:        ObHilbertWind,
		ScanMessageNo: text.HilbertWind_Scan,
		Sex:           'n',
		Synonyms:      []string{"wind"},
	},
	{
		AttackPercent: 100,
		DescMessageNo: text.Hope_Desc,
		Flags:         model.OfAnimate | model.OfShip,
		MaxCounter:    6,
		MaxLoc:        HilbertSpace12,
		MinLoc:        HilbertSpace7,
		Name:          "Hope satellite",
		Number:        ObHope,
		ScanMessageNo: text.Hope_Scan,
		Sex:           'n',
		Synonyms:      []string{"Hope"},
		Value:         61000,

		ShipGuns: [4]model.OldSGuns{
			core.SHIP_GUN_QUAD_LASER,
			core.SHIP_GUN_QUAD_LASER,
			core.SHIP_GUN_QUAD_LASER,
			core.SHIP_GUN_QUAD_LASER,
		},

		ShipKit: core.Equipment{
			Computer: 6,
			Engine:   30,
			Fuel:     1000,
			Hold:     70,
			Hull:     50,
			Shield:   30,
			Tonnage:  500,
		},
	},
	{
		// TODO: needs a GET event!
		DescMessageNo: text.Lever_Desc,
		MaxLoc:        MineFace,
		MinLoc:        MineFace,
		Name:          "lever",
		Number:        ObLever,
		ScanMessageNo: text.Lever_Scan,
		Sex:           'n',
		Weight:        math.MaxUint16,
	},
	{
		DescMessageNo: text.Paper_Desc,
		Events:        [4]uint16{0, 8, 0, 0},
		MaxLoc:        HoloRoom,
		MinLoc:        LivingQuarters,
		Name:          "paper",
		Number:        ObPaper,
		ScanMessageNo: text.Paper_Scan,
		Sex:           'n',
		Value:         78,
		Weight:        1,
	},
	{
		AttackPercent: 75,
		DescMessageNo: text.Patrol_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    3,
		MaxLoc:        Silos,
		MinLoc:        Corridor1,
		Name:          "patrol",
		Number:        ObPatrol,
		ScanMessageNo: text.Patrol_Scan,
		Sex:           'n',
	},
	{
		DescMessageNo: text.Potion_Desc,
		Flags:         model.OfLiquid,
		MaxLoc:        Laboratory,
		MinLoc:        Laboratory,
		Name:          "potion",
		Number:        ObPotion,
		ScanMessageNo: text.Potion_Scan,
		Sex:           'n',
		Synonyms:      []string{"flask"},
		Value:         250,
		Weight:        2,
	},
	{
		DescMessageNo: text.TuningFork_Desc,
		MaxLoc:        ReactorRoom,
		MinLoc:        ReactorRoom,
		Name:          "tuning fork",
		Number:        ObTuningFork,
		ScanMessageNo: text.TuningFork_Scan,
		Sex:           'n',
		Synonyms:      []string{"fork"},
		Value:         145,
		Weight:        1,
	},
}

var Planet = core.Planet{
	Name:    "Snark",
	Level:   uint16(model.LevelNoProduction),
	Orbit:   SnarkOrbit,
	Landing: LandingArea,
}
