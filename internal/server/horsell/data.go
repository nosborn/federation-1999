package horsell

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/core"
	"github.com/nosborn/federation-1999/internal/text"
)

var Locations = [80]core.Location{
	{
		Number:     Downs1,
		BriefMsgNo: text.HorsellLocation9_Brief,
		FullMsgNo:  text.HorsellLocation9_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			0,     // N
			0,     // NE
			0,     // E
			0,     // SE
			0,     // S
			0,     // SW
			Road1, // W
			0,     // NW
			0,     // Up
			0,     // Down
			0,     // In
			0,     // Out
		},
	},
	{
		Number:     Road1,
		BriefMsgNo: text.HorsellLocation10_Brief,
		FullMsgNo:  text.HorsellLocation10_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			Road2,  // N
			0,      // NE
			Downs1, // E
			0,      // SE
			0,      // S
			0,      // SW
			0,      // W
			0,      // NW
			0,      // Up
			0,      // Down
			0,      // In
			0,      // Out
		},
	},
	{
		Number:     Road2,
		BriefMsgNo: text.HorsellLocation11_Brief,
		FullMsgNo:  text.HorsellLocation11_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			Road3, // N
			0,     // NE
			0,     // E
			0,     // SE
			Road1, // S
			0,     // SW
			0,     // W
			0,     // NW
			0,     // Up
			0,     // Down
			0,     // In
			0,     // Out
		},
	},
	{
		Number:     Road3,
		BriefMsgNo: text.HorsellLocation12_Brief,
		FullMsgNo:  text.HorsellLocation12_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			Road4, // N
			0,     // NE
			0,     // E
			0,     // SE
			Road2, // S
			0,     // SW
			0,     // W
			Eaves, // NW
			0,     // Up
			0,     // Down
			0,     // In
			0,     // Out
		},
	},
	{
		Number:     Road4,
		BriefMsgNo: text.HorsellLocation13_Brief,
		FullMsgNo:  text.HorsellLocation13_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			BendInTheRoad1, // N
			0,              // NE
			0,              // E
			0,              // SE
			Road1,          // S
			0,              // SW
			Eaves,          // W
			0,              // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     BendInTheRoad1,
		BriefMsgNo: text.HorsellLocation14_Brief,
		FullMsgNo:  text.HorsellLocation14_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			0,           // E
			0,           // SE
			Road4,       // S
			Eaves,       // SW
			WindingRoad, // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     WindingRoad,
		BriefMsgNo: text.HorsellLocation15_Brief,
		FullMsgNo:  text.HorsellLocation15_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			0,              // N
			0,              // NE
			BendInTheRoad1, // E
			0,              // SE
			Eaves,          // S
			Copse3,         // SW
			BirchTrees2,    // W
			Road5,          // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     Road5,
		BriefMsgNo: text.HorsellLocation16_Brief,
		FullMsgNo:  text.HorsellLocation16_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			RoadJunction1, // N
			0,             // NE
			0,             // E
			WindingRoad,   // SE
			BirchTrees2,   // S
			BirchTrees1,   // SW
			Copse1,        // W
			0,             // NW
			RoadJunction1, // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     Road6,
		BriefMsgNo: text.HorsellLocation17_Brief,
		FullMsgNo:  text.HorsellLocation17_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			RoadJunction1, // E
			0,             // SE
			Copse1,        // S
			Downs2,        // SW
			Road7,         // W
			0,             // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     Road7,
		BriefMsgNo: text.HorsellLocation18_Brief,
		FullMsgNo:  text.HorsellLocation18_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			0,      // N
			0,      // NE
			Road6,  // E
			Copse1, // SE
			Downs2, // S
			Downs3, // SW
			Road8,  // W
			0,      // NW
			0,      // Up
			0,      // Down
			0,      // In
			0,      // Out
		},
	},
	{
		Number:     Road8,
		BriefMsgNo: text.HorsellLocation19_Brief,
		FullMsgNo:  text.HorsellLocation19_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			Road7,   // E
			Downs2,  // SE
			Downs3,  // S
			0,       // SW
			Roadway, // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Roadway,
		BriefMsgNo: text.HorsellLocation20_Brief,
		FullMsgNo:  text.HorsellLocation20_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			0,                    // N
			0,                    // NE
			Road8,                // E
			Downs3,               // SE
			0,                    // S
			0,                    // SW
			DamagedSectionOfRoad, // W
			0,                    // NW
			0,                    // Up
			0,                    // Down
			0,                    // In
			0,                    // Out
		},
	},
	{
		Number:     DamagedSectionOfRoad,
		BriefMsgNo: text.HorsellLocation21_Brief,
		FullMsgNo:  text.HorsellLocation21_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			Roadway,     // E
			0,           // SE
			0,           // S
			0,           // SW
			DamagedRoad, // W
			Crater3,     // NW
			0,           // Up
			Crater3,     // Down
			Crater3,     // In
			0,           // Out
		},
	},
	{
		Number:     DamagedRoad,
		BriefMsgNo: text.HorsellLocation22_Brief,
		FullMsgNo:  text.HorsellLocation22_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			Crater3,              // N
			0,                    // NE
			DamagedSectionOfRoad, // E
			0,                    // SE
			0,                    // S
			0,                    // SW
			Road9,                // W
			Downs5,               // NW
			0,                    // Up
			Crater3,              // Down
			Crater3,              // In
			0,                    // Out
		},
	},
	{
		Number:     Road9,
		BriefMsgNo: text.HorsellLocation23_Brief,
		FullMsgNo:  text.HorsellLocation23_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			Downs5,      // N
			0,           // NE
			DamagedRoad, // E
			0,           // SE
			0,           // S
			0,           // SW
			Road10,      // W
			Downs4,      // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     Road10,
		BriefMsgNo: text.HorsellLocation24_Brief,
		FullMsgNo:  text.HorsellLocation24_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			Downs4,         // N
			Downs5,         // NE
			Road9,          // E
			0,              // SE
			0,              // S
			0,              // SW
			BendInTheRoad2, // W
			0,              // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     BendInTheRoad2,
		BriefMsgNo: text.HorsellLocation25_Brief,
		FullMsgNo:  text.HorsellLocation25_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			0,      // N
			Downs4, // NE
			Road10, // E
			0,      // SE
			0,      // S
			0,      // SW
			0,      // W
			Road11, // NW
			0,      // Up
			0,      // Down
			0,      // In
			0,      // Out
		},
	},
	{
		Number:     Road11,
		BriefMsgNo: text.HorsellLocation26_Brief,
		FullMsgNo:  text.HorsellLocation26_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			0,              // N
			0,              // NE
			0,              // E
			BendInTheRoad2, // SE
			0,              // S
			0,              // SW
			EdgeOfGrove,    // W
			Road12,         // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     Road12,
		BriefMsgNo: text.HorsellLocation27_Brief,
		FullMsgNo:  text.HorsellLocation27_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			Grove1,      // N
			0,           // NE
			0,           // E
			Road11,      // SE
			EdgeOfGrove, // S
			Thicket2,    // SW
			Grove2,      // W
			Road13,      // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     Road13,
		BriefMsgNo: text.HorsellLocation28_Brief,
		FullMsgNo:  text.HorsellLocation28_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			Grove1,   // E
			Road12,   // SE
			Grove2,   // S
			Thicket1, // SW
			Trees,    // W
			Road14,   // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     Road14,
		BriefMsgNo: text.HorsellLocation29_Brief,
		FullMsgNo:  text.HorsellLocation29_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			0,      // N
			0,      // NE
			0,      // E
			Road13, // SE
			Trees,  // S
			0,      // SW
			0,      // W
			Road19, // NW
			0,      // Up
			0,      // Down
			0,      // In
			0,      // Out
		},
	},
	{
		Number:     Trees,
		BriefMsgNo: text.HorsellLocation30_Brief,
		FullMsgNo:  text.HorsellLocation30_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_351,

		MovTab: [13]uint16{
			Road14,   // N
			0,        // NE
			Road13,   // E
			Grove2,   // SE
			Thicket1, // S
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
		Number:     Grove1,
		BriefMsgNo: text.HorsellLocation31_Brief,
		FullMsgNo:  text.HorsellLocation31_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_351,

		MovTab: [13]uint16{
			0,      // N
			0,      // NE
			0,      // E
			0,      // SE
			Road12, // S
			0,      // SW
			Road13, // W
			0,      // NW
			0,      // Up
			0,      // Down
			0,      // In
			0,      // Out
		},
	},
	{
		Number:     Thicket1,
		BriefMsgNo: text.HorsellLocation32_Brief,
		FullMsgNo:  text.HorsellLocation32_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_351,

		MovTab: [13]uint16{
			Trees,    // N
			Road13,   // NE
			Grove2,   // E
			Thicket2, // SE
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
		Number:     Grove2,
		BriefMsgNo: text.HorsellLocation33_Brief,
		FullMsgNo:  text.HorsellLocation33_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_351,

		MovTab: [13]uint16{
			Road13,      // N
			0,           // NE
			Road12,      // E
			EdgeOfGrove, // SE
			Thicket2,    // S
			0,           // SW
			Thicket1,    // W
			Trees,       // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     Thicket2,
		BriefMsgNo: text.HorsellLocation34_Brief,
		FullMsgNo:  text.HorsellLocation34_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_351,

		MovTab: [13]uint16{
			Grove2,      // N
			Road12,      // NE
			EdgeOfGrove, // E
			0,           // SE
			0,           // S
			0,           // SW
			0,           // W
			Thicket1,    // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     EdgeOfGrove,
		BriefMsgNo: text.HorsellLocation35_Brief,
		FullMsgNo:  text.HorsellLocation35_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_351,

		MovTab: [13]uint16{
			Road12,   // N
			0,        // NE
			Road11,   // E
			0,        // SE
			0,        // S
			0,        // SW
			Thicket2, // W
			Grove2,   // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     RoadJunction1,
		BriefMsgNo: text.HorsellLocation36_Brief,
		FullMsgNo:  text.HorsellLocation36_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			Road15, // N
			0,      // NE
			0,      // E
			0,      // SE
			Road5,  // S
			Copse1, // SW
			Road6,  // W
			0,      // NW
			0,      // Up
			0,      // Down
			0,      // In
			0,      // Out
		},
	},
	{
		Number:     Road15,
		BriefMsgNo: text.HorsellLocation37_Brief,
		FullMsgNo:  text.HorsellLocation37_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			RoadJunction2, // N
			0,             // NE
			0,             // E
			0,             // SE
			RoadJunction1, // S
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
		Number:     RoadJunction2,
		BriefMsgNo: text.HorsellLocation38_Brief,
		FullMsgNo:  text.HorsellLocation38_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			0,         // N
			Path1,     // NE
			0,         // E
			0,         // SE
			Road15,    // S
			0,         // SW
			0,         // W
			Outskirts, // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     Outskirts,
		BriefMsgNo: text.HorsellLocation39_Brief,
		FullMsgNo:  text.HorsellLocation39_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_351,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			0,             // E
			RoadJunction2, // SE
			0,             // S
			0,             // SW
			0,             // W
			Road16,        // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     Downs2,
		BriefMsgNo: text.HorsellLocation40_Brief,
		FullMsgNo:  text.HorsellLocation40_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			Road7,       // N
			Road6,       // NE
			Copse1,      // E
			BirchTrees1, // SE
			Copse2,      // S
			0,           // SW
			Downs3,      // W
			Road8,       // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     Downs3,
		BriefMsgNo: text.HorsellLocation41_Brief,
		FullMsgNo:  text.HorsellLocation41_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			Road8,   // N
			Road7,   // NE
			Downs2,  // E
			Copse2,  // SE
			0,       // S
			0,       // SW
			0,       // W
			Roadway, // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Road16,
		BriefMsgNo: text.HorsellLocation42_Brief,
		FullMsgNo:  text.HorsellLocation42_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			Outskirts,  // SE
			Garden,     // S
			0,          // SW
			MainStreet, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     MainStreet,
		BriefMsgNo: text.HorsellLocation43_Brief,
		FullMsgNo:  text.HorsellLocation43_Full,
		Flags:      model.LfOutdoors,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			Road16,       // E
			0,            // SE
			GeneralStore, // S
			0,            // SW
			Road17,       // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     Copse1,
		BriefMsgNo: text.HorsellLocation44_Brief,
		FullMsgNo:  text.HorsellLocation44_Full,
		Flags:      model.LfOutdoors,

		MovTab: [13]uint16{
			Road6,         // N
			RoadJunction1, // NE
			Road5,         // E
			BirchTrees2,   // SE
			BirchTrees1,   // S
			Copse2,        // SW
			Downs2,        // W
			Road7,         // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     Copse2,
		BriefMsgNo: text.HorsellLocation45_Brief,
		FullMsgNo:  text.HorsellLocation45_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			Downs2,             // N
			Copse1,             // NE
			BirchTrees1,        // E
			TangledUndergrowth, // SE
			0,                  // S
			0,                  // SW
			0,                  // W
			Downs3,             // NW
			0,                  // Up
			0,                  // Down
			0,                  // In
			0,                  // Out
		},
	},
	{
		Number:     BirchTrees1,
		BriefMsgNo: text.HorsellLocation46_Brief,
		FullMsgNo:  text.HorsellLocation46_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_351,

		MovTab: [13]uint16{
			Copse1,             // N
			Road5,              // NE
			BirchTrees2,        // E
			Copse3,             // SE
			TangledUndergrowth, // S
			0,                  // SW
			Copse2,             // W
			Downs2,             // NW
			0,                  // Up
			0,                  // Down
			0,                  // In
			0,                  // Out
		},
	},
	{
		Number:     BirchTrees2,
		BriefMsgNo: text.HorsellLocation47_Brief,
		FullMsgNo:  text.HorsellLocation47_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_351,

		MovTab: [13]uint16{
			Road5,              // N
			0,                  // NE
			WindingRoad,        // E
			Eaves,              // SE
			Copse3,             // S
			TangledUndergrowth, // SW
			BirchTrees1,        // W
			Copse1,             // NW
			0,                  // Up
			0,                  // Down
			0,                  // In
			0,                  // Out
		},
	},
	{
		Number:     TangledUndergrowth,
		BriefMsgNo: text.HorsellLocation48_Brief,
		FullMsgNo:  text.HorsellLocation48_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_351,

		MovTab: [13]uint16{
			BirchTrees1, // N
			BirchTrees2, // NE
			Copse3,      // E
			EdgeOfTrees, // SE
			0,           // S
			0,           // SW
			0,           // W
			Copse2,      // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     Copse3,
		BriefMsgNo: text.HorsellLocation49_Brief,
		FullMsgNo:  text.HorsellLocation49_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_351,

		MovTab: [13]uint16{
			BirchTrees2,        // N
			WindingRoad,        // NE
			Eaves,              // E
			0,                  // SE
			EdgeOfTrees,        // S
			0,                  // SW
			TangledUndergrowth, // W
			BirchTrees1,        // NW
			0,                  // Up
			0,                  // Down
			0,                  // In
			0,                  // Out
		},
	},
	{
		Number:     Eaves,
		BriefMsgNo: text.HorsellLocation50_Brief,
		FullMsgNo:  text.HorsellLocation50_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			WindingRoad,    // N
			BendInTheRoad1, // NE
			Road4,          // E
			Road3,          // SE
			0,              // S
			EdgeOfTrees,    // SW
			Copse3,         // W
			BirchTrees2,    // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     EdgeOfTrees,
		BriefMsgNo: text.HorsellLocation51_Brief,
		FullMsgNo:  text.HorsellLocation51_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			Copse3,             // N
			Eaves,              // NE
			0,                  // E
			0,                  // SE
			0,                  // S
			0,                  // SW
			0,                  // W
			TangledUndergrowth, // NW
			0,                  // Up
			0,                  // Down
			0,                  // In
			0,                  // Out
		},
	},
	{
		Number:     Road17,
		BriefMsgNo: text.HorsellLocation52_Brief,
		FullMsgNo:  text.HorsellLocation52_Full,
		Flags:      model.LfOutdoors,

		MovTab: [13]uint16{
			SideRoad,   // N
			0,          // NE
			MainStreet, // E
			0,          // SE
			Hallway,    // S
			0,          // SW
			Road18,     // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     SideRoad,
		BriefMsgNo: text.HorsellLocation53_Brief,
		FullMsgNo:  text.HorsellLocation53_Full,
		Flags:      model.LfOutdoors,

		MovTab: [13]uint16{
			DeadEnd,        // N
			0,              // NE
			TheExpressPub,  // E
			0,              // SE
			Road17,         // S
			0,              // SW
			HorsellStation, // W
			Observatory,    // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     DeadEnd,
		BriefMsgNo: text.HorsellLocation54_Brief,
		FullMsgNo:  text.HorsellLocation54_Full,
		Flags:      model.LfOutdoors,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			0,        // E
			0,        // SE
			SideRoad, // S
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
		Number:     Path1,
		BriefMsgNo: text.HorsellLocation55_Brief,
		FullMsgNo:  text.HorsellLocation55_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			Path3,         // N
			0,             // NE
			Path2,         // E
			0,             // SE
			0,             // S
			RoadJunction2, // SW
			0,             // W
			0,             // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     Path2,
		BriefMsgNo: text.HorsellLocation56_Brief,
		FullMsgNo:  text.HorsellLocation56_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Church,     // E
			Graveyard1, // SE
			0,          // S
			0,          // SW
			Path1,      // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			Path1,      // Out
		},
	},
	{
		Number:     Path3,
		BriefMsgNo: text.HorsellLocation57_Brief,
		FullMsgNo:  text.HorsellLocation57_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			VicarageGarden, // N
			0,              // NE
			0,              // E
			0,              // SE
			Path1,          // S
			0,              // SW
			0,              // W
			0,              // NW
			0,              // Up
			0,              // Down
			0,              // In
			Path1,          // Out
		},
	},
	{
		Number:     Garden,
		BriefMsgNo: text.HorsellLocation58_Brief,
		FullMsgNo:  text.HorsellLocation58_Full,
		Flags:      model.LfOutdoors,

		MovTab: [13]uint16{
			Road16, // N
			0,      // NE
			0,      // E
			0,      // SE
			0,      // S
			0,      // SW
			0,      // W
			0,      // NW
			0,      // Up
			0,      // Down
			0,      // In
			Road16, // Out
		},
	},
	{
		Number:     GeneralStore,
		BriefMsgNo: text.HorsellLocation59_Brief,
		FullMsgNo:  text.HorsellLocation59_Full,
		Flags:      model.LfIndoors,

		MovTab: [13]uint16{
			MainStreet, // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			MainStreet, // Out
		},
	},
	{
		Number:     Hallway,
		BriefMsgNo: text.HorsellLocation60_Brief,
		FullMsgNo:  text.HorsellLocation60_Full,
		Flags:      model.LfIndoors,

		MovTab: [13]uint16{
			Road17, // N
			0,      // NE
			0,      // E
			0,      // SE
			0,      // S
			0,      // SW
			0,      // W
			0,      // NW
			0,      // Up
			0,      // Down
			0,      // In
			Road17, // Out
		},
	},
	{
		Number:     Stable,
		BriefMsgNo: text.HorsellLocation61_Brief,
		FullMsgNo:  text.HorsellLocation61_Full,
		Flags:      model.LfIndoors,

		MovTab: [13]uint16{
			Road18, // N
			0,      // NE
			0,      // E
			0,      // SE
			0,      // S
			0,      // SW
			0,      // W
			0,      // NW
			0,      // Up
			0,      // Down
			0,      // In
			0,      // Out
		},
	},
	{
		Number:     Undertakers,
		BriefMsgNo: text.HorsellLocation62_Brief,
		FullMsgNo:  text.HorsellLocation62_Full,
		Flags:      model.LfOutdoors,

		MovTab: [13]uint16{
			0,      // N
			0,      // NE
			Road18, // E
			0,      // SE
			0,      // S
			0,      // SW
			0,      // W
			0,      // NW
			0,      // Up
			0,      // Down
			0,      // In
			0,      // Out
		},
	},
	{
		Number:     HorsellStation,
		BriefMsgNo: text.HorsellLocation63_Brief,
		FullMsgNo:  text.HorsellLocation63_Full,
		Flags:      model.LfOutdoors,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			SideRoad, // E
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
		Number:     Observatory,
		BriefMsgNo: text.HorsellLocation64_Brief,
		FullMsgNo:  text.HorsellLocation64_Full,
		Flags:      model.LfIndoors,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			0,        // E
			SideRoad, // SE
			0,        // S
			0,        // SW
			0,        // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			SideRoad, // Out
		},
	},
	{
		Number:     TheExpressPub,
		BriefMsgNo: text.HorsellLocation65_Brief,
		FullMsgNo:  text.HorsellLocation65_Full,
		Flags:      model.LfCafe | model.LfIndoors,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			0,        // E
			0,        // SE
			0,        // S
			0,        // SW
			SideRoad, // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			SideRoad, // Out
		},
	},
	{
		Number:     VicarageGarden,
		BriefMsgNo: text.HorsellLocation66_Brief,
		FullMsgNo:  text.HorsellLocation66_Full,
		Flags:      model.LfOutdoors,

		MovTab: [13]uint16{
			0,     // N
			Hall,  // NE
			0,     // E
			0,     // SE
			Path3, // S
			0,     // SW
			0,     // W
			0,     // NW
			0,     // Up
			0,     // Down
			0,     // In
			Path3, // Out
		},
	},
	{
		Number:     Graveyard1,
		BriefMsgNo: text.HorsellLocation67_Brief,
		FullMsgNo:  text.HorsellLocation67_Full,
		Flags:      model.LfOutdoors,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Graveyard2, // E
			0,          // SE
			0,          // S
			0,          // SW
			0,          // W
			Path2,      // NW
			0,          // Up
			0,          // Down
			0,          // In
			Path2,      // Out
		},
	},
	{
		Number:     Graveyard2,
		BriefMsgNo: text.HorsellLocation68_Brief,
		FullMsgNo:  text.HorsellLocation68_Full,
		Flags:      model.LfOutdoors,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			Graveyard1, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Church,
		BriefMsgNo: text.HorsellLocation69_Brief,
		FullMsgNo:  text.HorsellLocation69_Full,
		Flags:      model.LfIndoors,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			ChurchNave,   // E
			0,            // SE
			0,            // S
			0,            // SW
			Path2,        // W
			0,            // NW
			ChurchBelfry, // Up
			0,            // Down
			0,            // In
			Path2,        // Out
		},
	},
	{
		Number:     ChurchNave,
		BriefMsgNo: text.HorsellLocation70_Brief,
		FullMsgNo:  text.HorsellLocation70_Full,
		Flags:      model.LfIndoors,

		MovTab: [13]uint16{
			OrganLoft, // N
			0,         // NE
			0,         // E
			0,         // SE
			0,         // S
			0,         // SW
			Church,    // W
			0,         // NW
			OrganLoft, // Up
			0,         // Down
			0,         // In
			Church,    // Out
		},
	},
	{
		Number:     ChurchBelfry,
		BriefMsgNo: text.HorsellLocation71_Brief,
		FullMsgNo:  text.HorsellLocation71_Full,
		Flags:      model.LfIndoors,

		MovTab: [13]uint16{
			0,      // N
			0,      // NE
			0,      // E
			0,      // SE
			0,      // S
			0,      // SW
			0,      // W
			0,      // NW
			0,      // Up
			Church, // Down
			0,      // In
			Church, // Out
		},
	},
	{
		Number:     OrganLoft,
		BriefMsgNo: text.HorsellLocation72_Brief,
		FullMsgNo:  text.HorsellLocation72_Full,
		Flags:      model.LfIndoors,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			ChurchNave, // SE
			0,          // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			ChurchNave, // Down
			0,          // In
			ChurchNave, // Out
		},
	},
	{
		Number:     Hall,
		BriefMsgNo: text.HorsellLocation73_Brief,
		FullMsgNo:  text.HorsellLocation73_Full,
		Flags:      model.LfIndoors,

		MovTab: [13]uint16{
			Study,          // N
			Kitchen,        // NE
			Bedroom,        // E
			0,              // SE
			0,              // S
			VicarageGarden, // SW
			0,              // W
			0,              // NW
			0,              // Up
			0,              // Down
			0,              // In
			VicarageGarden, // Out
		},
	},
	{
		Number:     Study,
		BriefMsgNo: text.HorsellLocation74_Brief,
		FullMsgNo:  text.HorsellLocation74_Full,
		Flags:      model.LfIndoors,

		MovTab: [13]uint16{
			0,    // N
			0,    // NE
			0,    // E
			0,    // SE
			Hall, // S
			0,    // SW
			0,    // W
			0,    // NW
			0,    // Up
			0,    // Down
			0,    // In
			Hall, // Out
		},
	},
	{
		Number:     Kitchen,
		BriefMsgNo: text.HorsellLocation75_Brief,
		FullMsgNo:  text.HorsellLocation75_Full,
		Flags:      model.LfIndoors,

		MovTab: [13]uint16{
			0,      // N
			0,      // NE
			0,      // E
			0,      // SE
			0,      // S
			Hall,   // SW
			0,      // W
			0,      // NW
			0,      // Up
			Cellar, // Down
			0,      // In
			Hall,   // Out
		},
	},
	{
		Number:     Bedroom,
		BriefMsgNo: text.HorsellLocation76_Brief,
		FullMsgNo:  text.HorsellLocation76_Full,
		Flags:      model.LfIndoors,

		MovTab: [13]uint16{
			0,    // N
			0,    // NE
			0,    // E
			0,    // SE
			0,    // S
			0,    // SW
			Hall, // W
			0,    // NW
			0,    // Up
			0,    // Down
			0,    // In
			Hall, // Out
		},
	},
	{
		Number:     Cellar,
		BriefMsgNo: text.HorsellLocation77_Brief,
		FullMsgNo:  text.HorsellLocation77_Full,
		Flags:      model.LfIndoors,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			0,       // E
			Tunnel,  // SE
			0,       // S
			0,       // SW
			0,       // W
			0,       // NW
			Kitchen, // Up
			0,       // Down
			0,       // In
			Kitchen, // Out
		},
	},
	{
		Number:     Tunnel,
		BriefMsgNo: text.HorsellLocation78_Brief,
		FullMsgNo:  text.HorsellLocation78_Full,
		Flags:      model.LfDark | model.LfIndoors,

		MovTab: [13]uint16{
			0,      // N
			0,      // NE
			0,      // E
			0,      // SE
			0,      // S
			0,      // SW
			0,      // W
			Cellar, // NW
			0,      // Up
			0,      // Down
			0,      // In
			Cellar, // Out
		},
	},
	{
		Number:     Road18,
		BriefMsgNo: text.HorsellLocation79_Brief,
		FullMsgNo:  text.HorsellLocation79_Full,
		Flags:      model.LfOutdoors,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			Road17,      // E
			0,           // SE
			Stable,      // S
			0,           // SW
			Undertakers, // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     Crater1,
		BriefMsgNo: text.HorsellLocation80_Brief,
		FullMsgNo:  text.HorsellLocation80_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			Cylinder2, // SE
			Cylinder1, // S
			Crater2,   // SW
			0,         // W
			0,         // NW
			0,         // Up
			Cylinder1, // Down
			Cylinder1, // In
			0,         // Out
		},
	},
	{
		Number:     Crater2,
		BriefMsgNo: text.HorsellLocation81_Brief,
		FullMsgNo:  text.HorsellLocation81_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			0,         // N
			Crater1,   // NE
			Cylinder1, // E
			0,         // SE
			0,         // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			Cylinder1, // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     Cylinder1,
		BriefMsgNo: text.HorsellLocation82_Brief,
		FullMsgNo:  text.HorsellLocation82_Full,
		Flags:      model.LfOutdoors,

		MovTab: [13]uint16{
			Crater1,   // N
			0,         // NE
			Cylinder2, // E
			Crater3,   // SE
			0,         // S
			0,         // SW
			Crater2,   // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			Crater3,   // Out
		},
	},
	{
		Number:     Cylinder2,
		BriefMsgNo: text.HorsellLocation83_Brief,
		FullMsgNo:  text.HorsellLocation83_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			0,         // SE
			Crater3,   // S
			0,         // SW
			Cylinder1, // W
			Crater1,   // NW
			Crater3,   // Up
			0,         // Down
			0,         // In
			Crater3,   // Out
		},
	},
	{
		Number:     Crater3,
		BriefMsgNo: text.HorsellLocation84_Brief,
		FullMsgNo:  text.HorsellLocation84_Full,
		Flags:      model.LfOutdoors,

		MovTab: [13]uint16{
			Cylinder2,            // N
			0,                    // NE
			0,                    // E
			DamagedSectionOfRoad, // SE
			DamagedRoad,          // S
			0,                    // SW
			0,                    // W
			Cylinder1,            // NW
			0,                    // Up
			Cylinder2,            // Down
			Cylinder1,            // In
			DamagedRoad,          // Out
		},
	},
	{
		Number:     Downs4,
		BriefMsgNo: text.HorsellLocation85_Brief,
		FullMsgNo:  text.HorsellLocation85_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			0,                    // N
			0,                    // NE
			Downs5,               // E
			DamagedSectionOfRoad, // SE
			Road10,               // S
			BendInTheRoad2,       // SW
			0,                    // W
			0,                    // NW
			0,                    // Up
			0,                    // Down
			0,                    // In
			0,                    // Out
		},
	},
	{
		Number:     Downs5,
		BriefMsgNo: text.HorsellLocation86_Brief,
		FullMsgNo:  text.HorsellLocation86_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			0,           // E
			DamagedRoad, // SE
			Road9,       // S
			Road10,      // SW
			Downs4,      // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     Road19,
		BriefMsgNo: text.HorsellLocation87_Brief,
		FullMsgNo:  text.HorsellLocation87_Full,
		Flags:      model.LfOutdoors,
		SysLoc:     text.NO_MOVEMENT_352,

		MovTab: [13]uint16{
			0,      // N
			0,      // NE
			0,      // E
			Road14, // SE
			0,      // S
			0,      // SW
			0,      // W
			0,      // NW
			0,      // Up
			0,      // Down
			0,      // In
			0,      // Out
		},
	},
	{
		Number:     GoodiesStore,
		BriefMsgNo: text.HorsellLocation88_Brief,
		FullMsgNo:  text.HorsellLocation88_Full,
		Flags:      model.LfHidden | model.LfLanding | model.LfShield,

		MovTab: [13]uint16{
			0, // N
			0, // NE
			0, // E
			0, // SE
			0, // S
			0, // SW
			0, // W
			0, // NW
			0, // Up
			0, // Down
			0, // In
			0, // Out
		},
	},
}

var Objects = [7]core.Object{
	{
		DescMessageNo: text.Artilleryman_Desc,
		Flags:         model.OfAnimate | model.OfHidden,
		MaxCounter:    -1,
		MaxLoc:        Cellar,
		MinLoc:        Cellar,
		Name:          "artilleryman",
		Number:        ObArtilleryman,
		ScanMessageNo: text.Artilleryman_Scan,
		Sex:           'm',
	},
	{
		DescMessageNo: text.Astronomer_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    -1,
		MaxLoc:        Observatory,
		MinLoc:        Observatory,
		Name:          "astronomer",
		Number:        ObAstronomer,
		ScanMessageNo: text.Astronomer_Scan,
		Sex:           'm',
	},
	{
		DescMessageNo: text.DogTag_Desc,
		Flags:         model.OfHidden,
		MaxLoc:        Trees,
		MinLoc:        Roadway,
		Name:          "dog tag",
		Number:        ObDogTag,
		ScanMessageNo: text.DogTag_Scan,
		Sex:           'n',
		Value:         1,
		Weight:        1,
	},
	{
		AttackPercent: 99,
		DescMessageNo: text.Footpad_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    2,
		MaxLoc:        EdgeOfTrees,
		MinLoc:        BendInTheRoad1,
		Name:          "footpad",
		Number:        ObFootpad,
		ScanMessageNo: text.Footpad_Scan,
		Sex:           'm',
	},
	{
		DescMessageNo: text.Marsmetal_Desc,
		Flags:         model.OfHidden,
		MaxLoc:        Cylinder1,
		MinLoc:        Cylinder1,
		Name:          "marsmetal",
		Number:        ObMarsmetal,
		ScanMessageNo: text.Marsmetal_Scan,
		Sex:           'n',
		Synonyms:      []string{"metal"},
		Value:         1,
		Weight:        1,
	},
	{
		DescMessageNo: text.Publican_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    -1,
		MaxLoc:        TheExpressPub,
		MinLoc:        TheExpressPub,
		Name:          "publican",
		Number:        ObPublican,
		ScanMessageNo: text.Publican_Scan,
		Sex:           'm',
	},
	{
		DescMessageNo: text.Sermon_Desc,
		MaxLoc:        Study,
		MinLoc:        Study,
		Name:          "sermon",
		Number:        ObSermon,
		ScanMessageNo: text.Sermon_Scan,
		Sex:           'n',
		Value:         1,
		Weight:        1,
	},
}

var Planet = core.Planet{
	Level:      uint16(model.LevelNoProduction),
	Population: 200,
	Landing:    GoodiesStore,
}
