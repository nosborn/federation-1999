package sol

import (
	"math"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/core"
	"github.com/nosborn/federation-1999/internal/text"
)

var Events = [54]core.Event{
	{
		Type:      1,
		MessageNo: text.SolEvent1,
	},
	{
		Type:      1,
		MessageNo: text.SolEvent2,
		Field2:    -15,
	},
	{
		Type:      3,
		MessageNo: text.SolEvent3,
		Field4:    55,
		Field7:    -8,
	},
	{
		Type:      1,
		MessageNo: text.SolEvent4,
		Field2:    -200,
	},
	{
		Type:      1,
		MessageNo: text.SolEvent5,
		NewLoc:    113,
	},
	{
		Type:      1,
		MessageNo: text.SolEvent6,
	},
	{
		Type:      1,
		MessageNo: text.SolEvent7,
		Field1:    -15,
		Field2:    -15,
		Field4:    -30,
	},
	{
		Type:      1,
		MessageNo: text.SolEvent8,
		Field2:    -5,
	},
	{
		Type:      1,
		MessageNo: text.SolEvent9,
		NewLoc:    485,
	},
	{
		Type:      1,
		MessageNo: text.SolEvent10,
	},
	{
		Type:      1,
		MessageNo: text.SolEvent11,
		NewLoc:    300,
	},
	{
		Type:      3,
		MessageNo: text.SolEvent12,
		Field3:    45,
	},
	{
		Type: 8,
	},
	{
		Type: 8,
	},
	{
		Type: 8,
		// Handler: increaseStrength,
	},
	{
		Type: 8,
	},
	{
		Type:      1,
		MessageNo: text.SolEvent17,
		Field2:    -199,
	},
	{
		Type: 8,
	},
	{
		Type: 8,
	},
	{
		Type: 8,
		// Handler: EnterMI6,
	},
	{
		Type: 8,
	},
	{
		Type: 8,
	},
	{
		Type: 8,
	},
	{
		Type: 8,
	},
	{
		Type:      1,
		MessageNo: text.SolEvent25,
		Field2:    -30,
	},
	{
		Type: 8,
	},
	{
		Type: 8,
	},
	{
		Type: 8,
		// Handler: snarkSapStamina,
	},
	{
		Type: 8,
		// Handler: silentSapStamina,
	},
	{
		Type: 8,
		// Handler: snarkSiloDamage,
	},
	{
		Type: 8,
	},
	{
		Type:      1,
		MessageNo: text.SolEvent32,
		Field2:    -73,
		NewLoc:    746,
	},
	{
		Type:      1,
		MessageNo: text.SolEvent33,
		Field2:    -65,
		NewLoc:    LinesRoom,
	},
	{
		Type: 8,
	},
	{
		Type: 8,
	},
	{
		Type: 8,
	},
	{
		Type:      1,
		MessageNo: text.SolEvent37,
		Field1:    -5,
		Field2:    -200,
		Field3:    -5,
		Field4:    -5,
	},
	{
		Type: 8,
	},
	{
		Type: 8,
		// Handler: restoreDogMovementTable,
	},
	{
		Type: 8,
		// Handler: restoreNisrikMovementTable,
	},
	{
		Type: 8,
	},
	{
		Type:      9,
		MessageNo: text.SolEvent42,
		Field2:    -999,
		Field8:    -ObBadge,
	},
	{
		Type: 8,
	},
	{
		Type: 8,
	},
	{
		Type: 8,
		// Handler: _DressCode,
	},
	{
		Type:      1,
		MessageNo: text.SolEvent46,
		Field2:    -199,
	},
	{
		Type: 8,
	},
	{
		Type: 8,
	},
	{
		Type: 8,
	},
	{
		Type:      1,
		MessageNo: text.SolEvent50,
		Field2:    -5,
		NewLoc:    Cell,
	},
	{
		Type: 8,
	},
	{
		Type: 8,
	},
	{
		Type: 8,
	},
	{
		Type: 8,
	},
}

var Locations = [698]core.Location{
	{
		Number:     1, // shipCommandCentre
		FullMsgNo:  text.SolLocation1_Full,
		BriefMsgNo: text.SolLocation1_Brief,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0, // N
			0, // NE
			0, // E
			0, // SE
			2, // S
			7, // SW
			0, // W
			0, // NW
			0, // Up
			0, // Down
			0, // In
			2, // Out
		},
	},
	{
		Number:     2, // shipAccessCorridor
		FullMsgNo:  text.SolLocation2_Full,
		BriefMsgNo: text.SolLocation2_Brief,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			1, // N
			4, // NE
			3, // E
			5, // SE
			8, // S
			0, // SW
			6, // W
			7, // NW
			0, // Up
			0, // Down
			0, // In
			4, // Out
		},
	},
	{
		Number:     3, // shipGalley
		BriefMsgNo: text.SolLocation3_Brief,
		FullMsgNo:  text.SolLocation3_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0, // N
			0, // NE
			0, // E
			0, // SE
			0, // S
			0, // SW
			2, // W
			0, // NW
			0, // Up
			0, // Down
			0, // In
			2, // Out
		},
	},
	{
		Number:     4, // shipAirlock
		BriefMsgNo: text.SolLocation4_Brief,
		FullMsgNo:  text.SolLocation4_Full,
		Flags:      model.LfLock | model.LfShield,

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
			2, // In
			0, // Out
		},
	},
	{
		Number:     5, // shipEngineRoom
		BriefMsgNo: text.SolLocation5_Brief,
		FullMsgNo:  text.SolLocation5_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0, // N
			0, // NE
			0, // E
			0, // SE
			0, // S
			0, // SW
			0, // W
			2, // NW
			0, // Up
			0, // Down
			0, // In
			2, // Out
		},
	},
	{
		Number:     6, // shipLivingQuarters
		BriefMsgNo: text.SolLocation6_Brief,
		FullMsgNo:  text.SolLocation6_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0, // N
			0, // NE
			2, // E
			0, // SE
			0, // S
			0, // SW
			0, // W
			0, // NW
			0, // Up
			0, // Down
			0, // In
			2, // Out
		},
	},
	{
		Number:     7, // shipSleepingQuarters
		BriefMsgNo: text.SolLocation7_Brief,
		FullMsgNo:  text.SolLocation7_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0, // N
			1, // NE
			0, // E
			2, // SE
			0, // S
			0, // SW
			0, // W
			0, // NW
			0, // Up
			0, // Down
			0, // In
			2, // Out
		},
	},
	{
		Number:     8, // shipCargoBay
		BriefMsgNo: text.SolLocation8_Brief,
		FullMsgNo:  text.SolLocation8_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			2, // N
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
			2, // Out
		},
	},
	{
		Number:     SolarSystemInterstellarLink,
		BriefMsgNo: text.SolLocation9_Brief,
		FullMsgNo:  text.SolLocation9_Full,
		Flags:      model.LfLink | model.LfPeace | model.LfSpace,

		MovTab: [13]uint16{
			0,                    // N
			0,                    // NE
			InterplanetarySpace1, // E
			InterplanetarySpace6, // SE
			InterplanetarySpace5, // S
			0,                    // SW
			0,                    // W
			0,                    // NW
			0,                    // Up
			0,                    // Down
			0,                    // In
			0,                    // Out
			InterplanetarySpace1, // Planet
		},
	},
	{
		Number:     InterplanetarySpace1,
		BriefMsgNo: text.SolLocation10_Brief,
		FullMsgNo:  text.SolLocation10_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,                           // N
			0,                           // NE
			SaturnOrbit,                 // E
			InterplanetarySpace7,        // SE
			InterplanetarySpace6,        // S
			InterplanetarySpace5,        // SW
			SolarSystemInterstellarLink, // W
			0,                           // NW
			0,                           // Up
			0,                           // Down
			0,                           // In
			0,                           // Out
			SaturnOrbit,                 // Planet
		},
	},
	{
		Number:     SaturnOrbit,
		BriefMsgNo: text.SolLocation11_Brief,
		FullMsgNo:  text.SolLocation11_Full,
		Flags:      model.LfOrbit | model.LfSpace,

		MovTab: [13]uint16{
			0,                    // N
			0,                    // NE
			TitanOrbit,           // E
			InterplanetarySpace8, // SE
			InterplanetarySpace7, // S
			InterplanetarySpace6, // SW
			InterplanetarySpace1, // W
			0,                    // NW
			0,                    // Up
			0,                    // Down
			0,                    // In
			0,                    // Out
			0,                    // Planet
		},
	},
	{
		Number:     TitanOrbit,
		BriefMsgNo: text.SolLocation12_Brief,
		FullMsgNo:  text.SolLocation12_Full,
		Flags:      model.LfOrbit | model.LfPeace | model.LfSpace,

		MovTab: [13]uint16{
			0,                    // N
			0,                    // NE
			InterplanetarySpace2, // E
			InterplanetarySpace9, // SE
			InterplanetarySpace8, // S
			InterplanetarySpace7, // SW
			SaturnOrbit,          // W
			0,                    // NW
			0,                    // Up
			0,                    // Down
			0,                    // In
			0,                    // Out
			0,                    // Planet
		},
	},
	{
		Number:     InterplanetarySpace2,
		BriefMsgNo: text.SolLocation13_Brief,
		FullMsgNo:  text.SolLocation13_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,                    // N
			0,                    // NE
			InterplanetarySpace3, // E
			CallistoOrbit,        // SE
			InterplanetarySpace9, // S
			InterplanetarySpace8, // SW
			TitanOrbit,           // W
			0,                    // NW
			0,                    // Up
			0,                    // Down
			0,                    // In
			0,                    // Out
			CallistoOrbit,        // Planet
		},
	},
	{
		Number:     InterplanetarySpace3,
		BriefMsgNo: text.SolLocation14_Brief,
		FullMsgNo:  text.SolLocation14_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,                     // N
			0,                     // NE
			InterplanetarySpace4,  // E
			InterplanetarySpace10, // SE
			CallistoOrbit,         // S
			InterplanetarySpace9,  // SW
			InterplanetarySpace2,  // W
			0,                     // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			CallistoOrbit,         // Planet
		},
	},
	{
		Number:     InterplanetarySpace4,
		BriefMsgNo: text.SolLocation15_Brief,
		FullMsgNo:  text.SolLocation15_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,                     // N
			0,                     // NE
			0,                     // E
			0,                     // SE
			InterplanetarySpace10, // S
			CallistoOrbit,         // SW
			InterplanetarySpace3,  // W
			0,                     // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			CallistoOrbit,         // Planet
		},
	},
	{
		Number:     InterplanetarySpace5,
		BriefMsgNo: text.SolLocation16_Brief,
		FullMsgNo:  text.SolLocation16_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			SolarSystemInterstellarLink, // N
			InterplanetarySpace1,        // NE
			InterplanetarySpace6,        // E
			0,                           // SE
			0,                           // S
			0,                           // SW
			0,                           // W
			0,                           // NW
			0,                           // Up
			0,                           // Down
			0,                           // In
			0,                           // Out
			InterplanetarySpace6,        // Planet
		},
	},
	{
		Number:     InterplanetarySpace6,
		BriefMsgNo: text.SolLocation17_Brief,
		FullMsgNo:  text.SolLocation17_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace1,        // N
			SaturnOrbit,                 // NE
			InterplanetarySpace7,        // E
			0,                           // SE
			0,                           // S
			0,                           // SW
			InterplanetarySpace5,        // W
			SolarSystemInterstellarLink, // NW
			0,                           // Up
			0,                           // Down
			0,                           // In
			0,                           // Out
			SaturnOrbit,                 // Planet
		},
	},
	{
		Number:     InterplanetarySpace7,
		BriefMsgNo: text.SolLocation18_Brief,
		FullMsgNo:  text.SolLocation18_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			SaturnOrbit,           // N
			TitanOrbit,            // NE
			InterplanetarySpace8,  // E
			InterplanetarySpace11, // SE
			0,                     // S
			0,                     // SW
			InterplanetarySpace6,  // W
			InterplanetarySpace1,  // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			TitanOrbit,            // Planet
		},
	},
	{
		Number:     InterplanetarySpace8,
		BriefMsgNo: text.SolLocation19_Brief,
		FullMsgNo:  text.SolLocation19_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			TitanOrbit,            // N
			InterplanetarySpace2,  // NE
			InterplanetarySpace9,  // E
			JupiterOrbit,          // SE
			InterplanetarySpace11, // S
			0,                     // SW
			InterplanetarySpace7,  // W
			SaturnOrbit,           // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			TitanOrbit,            // Planet
		},
	},
	{
		Number:     InterplanetarySpace9,
		BriefMsgNo: text.SolLocation20_Brief,
		FullMsgNo:  text.SolLocation20_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace2,  // N
			InterplanetarySpace3,  // NE
			CallistoOrbit,         // E
			InterplanetarySpace12, // SE
			JupiterOrbit,          // S
			InterplanetarySpace11, // SW
			InterplanetarySpace8,  // W
			TitanOrbit,            // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			CallistoOrbit,         // Planet
		},
	},
	{
		Number:     CallistoOrbit,
		BriefMsgNo: text.SolLocation21_Brief,
		FullMsgNo:  text.SolLocation21_Full,
		Flags:      model.LfOrbit | model.LfPeace | model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace3,  // N
			InterplanetarySpace4,  // NE
			InterplanetarySpace10, // E
			InterplanetarySpace13, // SE
			InterplanetarySpace12, // S
			JupiterOrbit,          // SW
			InterplanetarySpace9,  // W
			InterplanetarySpace2,  // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			0,                     // Planet
		},
	},
	{
		Number:     InterplanetarySpace10,
		BriefMsgNo: text.SolLocation22_Brief,
		FullMsgNo:  text.SolLocation22_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace4,  // N
			0,                     // NE
			0,                     // E
			InterplanetarySpace14, // SE
			InterplanetarySpace13, // S
			InterplanetarySpace12, // SW
			CallistoOrbit,         // W
			InterplanetarySpace3,  // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			CallistoOrbit,         // Planet
		},
	},
	{
		Number:     InterplanetarySpace11,
		BriefMsgNo: text.SolLocation23_Brief,
		FullMsgNo:  text.SolLocation23_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace8,  // N
			InterplanetarySpace9,  // NE
			JupiterOrbit,          // E
			InterplanetarySpace16, // SE
			InterplanetarySpace15, // S
			0,                     // SW
			0,                     // W
			InterplanetarySpace7,  // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			JupiterOrbit,          // Planet
		},
	},
	{
		Number:     JupiterOrbit,
		BriefMsgNo: text.SolLocation24_Brief,
		FullMsgNo:  text.SolLocation24_Full,
		Flags:      model.LfOrbit | model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace9,  // N
			CallistoOrbit,         // NE
			InterplanetarySpace12, // E
			InterplanetarySpace17, // SE
			InterplanetarySpace16, // S
			InterplanetarySpace15, // SW
			InterplanetarySpace11, // W
			InterplanetarySpace8,  // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			0,                     // Planet
		},
	},
	{
		Number:     InterplanetarySpace12,
		BriefMsgNo: text.SolLocation25_Brief,
		FullMsgNo:  text.SolLocation25_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			CallistoOrbit,         // N
			InterplanetarySpace10, // NE
			InterplanetarySpace13, // E
			MarsOrbit,             // SE
			InterplanetarySpace17, // S
			InterplanetarySpace16, // SW
			JupiterOrbit,          // W
			InterplanetarySpace9,  // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			MarsOrbit,             // Planet
		},
	},
	{
		Number:     InterplanetarySpace13,
		BriefMsgNo: text.SolLocation26_Brief,
		FullMsgNo:  text.SolLocation26_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace10, // N
			0,                     // NE
			InterplanetarySpace14, // E
			InterplanetarySpace18, // SE
			MarsOrbit,             // S
			InterplanetarySpace17, // SW
			InterplanetarySpace12, // W
			CallistoOrbit,         // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			MarsOrbit,             // Planet
		},
	},
	{
		Number:     InterplanetarySpace14,
		BriefMsgNo: text.SolLocation27_Brief,
		FullMsgNo:  text.SolLocation27_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,                     // N
			0,                     // NE
			0,                     // E
			InterplanetarySpace19, // SE
			InterplanetarySpace18, // S
			MarsOrbit,             // SW
			InterplanetarySpace13, // W
			InterplanetarySpace10, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			MarsOrbit,             // Planet
		},
	},
	{
		Number:     InterplanetarySpace15,
		BriefMsgNo: text.SolLocation28_Brief,
		FullMsgNo:  text.SolLocation28_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace11, // N
			JupiterOrbit,          // NE
			InterplanetarySpace16, // E
			InterplanetarySpace21, // SE
			0,                     // S
			0,                     // SW
			0,                     // W
			0,                     // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			JupiterOrbit,          // Planet
		},
	},
	{
		Number:     InterplanetarySpace16,
		BriefMsgNo: text.SolLocation29_Brief,
		FullMsgNo:  text.SolLocation29_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			JupiterOrbit,          // N
			InterplanetarySpace12, // NE
			InterplanetarySpace17, // E
			InterplanetarySpace22, // SE
			InterplanetarySpace21, // S
			0,                     // SW
			InterplanetarySpace15, // W
			InterplanetarySpace11, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			JupiterOrbit,          // Planet
		},
	},
	{
		Number:     InterplanetarySpace17,
		BriefMsgNo: text.SolLocation30_Brief,
		FullMsgNo:  text.SolLocation30_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace12, // N
			InterplanetarySpace13, // NE
			MarsOrbit,             // E
			InterplanetarySpace23, // SE
			InterplanetarySpace22, // S
			InterplanetarySpace21, // SW
			InterplanetarySpace16, // W
			JupiterOrbit,          // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			MarsOrbit,             // Planet
		},
	},
	{
		Number:     MarsOrbit,
		BriefMsgNo: text.SolLocation31_Brief,
		FullMsgNo:  text.SolLocation31_Full,
		Flags:      model.LfOrbit | model.LfPeace | model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace13, // N
			InterplanetarySpace14, // NE
			InterplanetarySpace18, // E
			InterplanetarySpace24, // SE
			InterplanetarySpace23, // S
			InterplanetarySpace22, // SW
			InterplanetarySpace17, // W
			InterplanetarySpace12, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			0,                     // Planet
		},
	},
	{
		Number:     InterplanetarySpace18,
		BriefMsgNo: text.SolLocation32_Brief,
		FullMsgNo:  text.SolLocation32_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace14, // N
			0,                     // NE
			InterplanetarySpace19, // E
			EarthOrbit,            // SE
			InterplanetarySpace24, // S
			InterplanetarySpace23, // SW
			MarsOrbit,             // W
			InterplanetarySpace13, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			EarthOrbit,            // Planet
		},
	},
	{
		Number:     InterplanetarySpace19,
		BriefMsgNo: text.SolLocation33_Brief,
		FullMsgNo:  text.SolLocation33_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,                     // N
			0,                     // NE
			InterplanetarySpace20, // E
			InterplanetarySpace25, // SE
			EarthOrbit,            // S
			InterplanetarySpace24, // SW
			InterplanetarySpace18, // W
			InterplanetarySpace14, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			EarthOrbit,            // Planet
		},
	},
	{
		Number:     InterplanetarySpace20,
		BriefMsgNo: text.SolLocation34_Brief,
		FullMsgNo:  text.SolLocation34_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,                     // N
			0,                     // NE
			0,                     // E
			InterplanetarySpace26, // SE
			InterplanetarySpace25, // S
			EarthOrbit,            // SW
			InterplanetarySpace19, // W
			0,                     // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			EarthOrbit,            // Planet
		},
	},
	{
		Number:     InterplanetarySpace21,
		BriefMsgNo: text.SolLocation35_Brief,
		FullMsgNo:  text.SolLocation35_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace16, // N
			InterplanetarySpace17, // NE
			InterplanetarySpace22, // E
			InterplanetarySpace28, // SE
			InterplanetarySpace27, // S
			0,                     // SW
			0,                     // W
			InterplanetarySpace15, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			InterplanetarySpace28, // Planet
		},
	},
	{
		Number:     InterplanetarySpace22,
		BriefMsgNo: text.SolLocation36_Brief,
		FullMsgNo:  text.SolLocation36_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace17, // N
			MarsOrbit,             // NE
			InterplanetarySpace23, // E
			InterplanetarySpace29, // SE
			InterplanetarySpace28, // S
			InterplanetarySpace27, // SW
			InterplanetarySpace21, // W
			InterplanetarySpace16, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			MarsOrbit,             // Planet
		},
	},
	{
		Number:     InterplanetarySpace23,
		BriefMsgNo: text.SolLocation37_Brief,
		FullMsgNo:  text.SolLocation37_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			MarsOrbit,             // N
			InterplanetarySpace18, // NE
			InterplanetarySpace24, // E
			InterplanetarySpace30, // SE
			InterplanetarySpace29, // S
			InterplanetarySpace28, // SW
			InterplanetarySpace22, // W
			InterplanetarySpace17, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			MarsOrbit,             // Planet
		},
	},
	{
		Number:     InterplanetarySpace24,
		BriefMsgNo: text.SolLocation38_Brief,
		FullMsgNo:  text.SolLocation38_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace18, // N
			InterplanetarySpace19, // NE
			EarthOrbit,            // E
			InterplanetarySpace31, // SE
			InterplanetarySpace30, // S
			InterplanetarySpace29, // SW
			InterplanetarySpace23, // W
			MarsOrbit,             // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			EarthOrbit,            // Planet
		},
	},
	{
		Number:     EarthOrbit,
		BriefMsgNo: text.SolLocation39_Brief,
		FullMsgNo:  text.SolLocation39_Full,
		Flags:      model.LfOrbit | model.LfPeace | model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace19, // N
			InterplanetarySpace20, // NE
			InterplanetarySpace25, // E
			LunarOrbit,            // SE
			InterplanetarySpace31, // S
			InterplanetarySpace30, // SW
			InterplanetarySpace24, // W
			InterplanetarySpace18, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			0,                     // Planet
		},
	},
	{
		Number:     InterplanetarySpace25,
		BriefMsgNo: text.SolLocation40_Brief,
		FullMsgNo:  text.SolLocation40_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace20, // N
			0,                     // NE
			InterplanetarySpace26, // E
			InterplanetarySpace32, // SE
			LunarOrbit,            // S
			InterplanetarySpace31, // SW
			EarthOrbit,            // W
			InterplanetarySpace19, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			EarthOrbit,            // Planet
		},
	},
	{
		Number:     InterplanetarySpace26,
		BriefMsgNo: text.SolLocation41_Brief,
		FullMsgNo:  text.SolLocation41_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,                     // N
			0,                     // NE
			0,                     // E
			0,                     // SE
			InterplanetarySpace32, // S
			LunarOrbit,            // SW
			InterplanetarySpace25, // W
			InterplanetarySpace20, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			LunarOrbit,            // Planet
		},
	},
	{
		Number:     InterplanetarySpace27,
		BriefMsgNo: text.SolLocation42_Brief,
		FullMsgNo:  text.SolLocation42_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace21, // N
			InterplanetarySpace22, // NE
			InterplanetarySpace28, // E
			VenusOrbit,            // SE
			InterplanetarySpace33, // S
			0,                     // SW
			0,                     // W
			0,                     // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			VenusOrbit,            // Planet
		},
	},
	{
		Number:     InterplanetarySpace28,
		BriefMsgNo: text.SolLocation43_Brief,
		FullMsgNo:  text.SolLocation43_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace22, // N
			InterplanetarySpace23, // NE
			InterplanetarySpace29, // E
			InterplanetarySpace34, // SE
			VenusOrbit,            // S
			InterplanetarySpace33, // SW
			InterplanetarySpace27, // W
			InterplanetarySpace21, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			VenusOrbit,            // Planet
		},
	},
	{
		Number:     InterplanetarySpace29,
		BriefMsgNo: text.SolLocation44_Brief,
		FullMsgNo:  text.SolLocation44_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace23, // N
			InterplanetarySpace24, // NE
			InterplanetarySpace30, // E
			InterplanetarySpace35, // SE
			InterplanetarySpace34, // S
			VenusOrbit,            // SW
			InterplanetarySpace28, // W
			InterplanetarySpace22, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			VenusOrbit,            // Planet
		},
	},
	{
		Number:     InterplanetarySpace30,
		BriefMsgNo: text.SolLocation45_Brief,
		FullMsgNo:  text.SolLocation45_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace24, // N
			EarthOrbit,            // NE
			InterplanetarySpace31, // E
			InterplanetarySpace36, // SE
			InterplanetarySpace35, // S
			InterplanetarySpace34, // SW
			InterplanetarySpace29, // W
			InterplanetarySpace23, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			EarthOrbit,            // Planet
		},
	},
	{
		Number:     InterplanetarySpace31,
		BriefMsgNo: text.SolLocation46_Brief,
		FullMsgNo:  text.SolLocation46_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			EarthOrbit,            // N
			InterplanetarySpace25, // NE
			LunarOrbit,            // E
			InterplanetarySpace37, // SE
			InterplanetarySpace36, // S
			InterplanetarySpace35, // SW
			InterplanetarySpace30, // W
			InterplanetarySpace24, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			LunarOrbit,            // Planet
		},
	},
	{
		Number:     LunarOrbit,
		BriefMsgNo: text.SolLocation47_Brief,
		FullMsgNo:  text.SolLocation47_Full,
		Flags:      model.LfOrbit | model.LfPeace | model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace25, // N
			InterplanetarySpace26, // NE
			InterplanetarySpace32, // E
			InterplanetarySpace38, // SE
			InterplanetarySpace37, // S
			InterplanetarySpace36, // SW
			InterplanetarySpace31, // W
			EarthOrbit,            // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			0,                     // Planet
		},
	},
	{
		Number:     InterplanetarySpace32,
		BriefMsgNo: text.SolLocation48_Brief,
		FullMsgNo:  text.SolLocation48_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace26, // N
			0,                     // NE
			0,                     // E
			0,                     // SE
			InterplanetarySpace38, // S
			InterplanetarySpace37, // SW
			LunarOrbit,            // W
			InterplanetarySpace25, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			LunarOrbit,            // Planet
		},
	},
	{
		Number:     InterplanetarySpace33,
		BriefMsgNo: text.SolLocation49_Brief,
		FullMsgNo:  text.SolLocation49_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace27, // N
			InterplanetarySpace28, // NE
			VenusOrbit,            // E
			InterplanetarySpace40, // SE
			InterplanetarySpace39, // S
			0,                     // SW
			0,                     // W
			0,                     // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			VenusOrbit,            // Planet
		},
	},
	{
		Number:     VenusOrbit,
		BriefMsgNo: text.SolLocation50_Brief,
		FullMsgNo:  text.SolLocation50_Full,
		Flags:      model.LfOrbit | model.LfPeace | model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace28, // N
			InterplanetarySpace29, // NE
			InterplanetarySpace34, // E
			InterplanetarySpace41, // SE
			InterplanetarySpace40, // S
			InterplanetarySpace39, // SW
			InterplanetarySpace33, // W
			InterplanetarySpace27, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			0,                     // Planet
		},
	},
	{
		Number:     InterplanetarySpace34,
		BriefMsgNo: text.SolLocation51_Brief,
		FullMsgNo:  text.SolLocation51_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace29, // N
			InterplanetarySpace30, // NE
			InterplanetarySpace35, // E
			InterplanetarySpace42, // SE
			InterplanetarySpace41, // S
			InterplanetarySpace40, // SW
			VenusOrbit,            // W
			InterplanetarySpace28, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			VenusOrbit,            // Planet
		},
	},
	{
		Number:     InterplanetarySpace35,
		BriefMsgNo: text.SolLocation52_Brief,
		FullMsgNo:  text.SolLocation52_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace30, // N
			InterplanetarySpace31, // NE
			InterplanetarySpace36, // E
			MercuryOrbit,          // SE
			InterplanetarySpace42, // S
			InterplanetarySpace41, // SW
			InterplanetarySpace34, // W
			InterplanetarySpace29, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			MercuryOrbit,          // Planet
		},
	},
	{
		Number:     InterplanetarySpace36,
		BriefMsgNo: text.SolLocation53_Brief,
		FullMsgNo:  text.SolLocation53_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace31, // N
			LunarOrbit,            // NE
			InterplanetarySpace37, // E
			InterplanetarySpace43, // SE
			MercuryOrbit,          // S
			InterplanetarySpace42, // SW
			InterplanetarySpace35, // W
			InterplanetarySpace30, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			MercuryOrbit,          // Planet
		},
	},
	{
		Number:     InterplanetarySpace37,
		BriefMsgNo: text.SolLocation54_Brief,
		FullMsgNo:  text.SolLocation54_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			LunarOrbit,            // N
			InterplanetarySpace32, // NE
			InterplanetarySpace38, // E
			InterplanetarySpace44, // SE
			InterplanetarySpace43, // S
			MercuryOrbit,          // SW
			InterplanetarySpace36, // W
			InterplanetarySpace31, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			MercuryOrbit,          // Planet
		},
	},
	{
		Number:     InterplanetarySpace38,
		BriefMsgNo: text.SolLocation55_Brief,
		FullMsgNo:  text.SolLocation55_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace32, // N
			0,                     // NE
			0,                     // E
			0,                     // SE
			InterplanetarySpace44, // S
			InterplanetarySpace43, // SW
			InterplanetarySpace37, // W
			LunarOrbit,            // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			LunarOrbit,            // Planet
		},
	},
	{
		Number:     InterplanetarySpace39,
		BriefMsgNo: text.SolLocation56_Brief,
		FullMsgNo:  text.SolLocation56_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace33, // N
			VenusOrbit,            // NE
			InterplanetarySpace40, // E
			0,                     // SE
			0,                     // S
			0,                     // SW
			0,                     // W
			0,                     // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			VenusOrbit,            // Planet
		},
	},
	{
		Number:     InterplanetarySpace40,
		BriefMsgNo: text.SolLocation57_Brief,
		FullMsgNo:  text.SolLocation57_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			VenusOrbit,            // N
			InterplanetarySpace34, // NE
			InterplanetarySpace41, // E
			0,                     // SE
			0,                     // S
			0,                     // SW
			InterplanetarySpace39, // W
			InterplanetarySpace33, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			VenusOrbit,            // Planet
		},
	},
	{
		Number:     InterplanetarySpace41,
		BriefMsgNo: text.SolLocation58_Brief,
		FullMsgNo:  text.SolLocation58_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace34, // N
			InterplanetarySpace35, // NE
			InterplanetarySpace42, // E
			SolarOrbit1,           // SE
			0,                     // S
			0,                     // SW
			InterplanetarySpace40, // W
			VenusOrbit,            // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			VenusOrbit,            // Planet
		},
	},
	{
		Number:     InterplanetarySpace42,
		BriefMsgNo: text.SolLocation59_Brief,
		FullMsgNo:  text.SolLocation59_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace35, // N
			InterplanetarySpace36, // NE
			MercuryOrbit,          // E
			SolarOrbit2,           // SE
			SolarOrbit1,           // S
			0,                     // SW
			InterplanetarySpace41, // W
			InterplanetarySpace34, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			MercuryOrbit,          // Planet
		},
	},
	{
		Number:     MercuryOrbit,
		BriefMsgNo: text.SolLocation60_Brief,
		FullMsgNo:  text.SolLocation60_Full,
		Flags:      model.LfOrbit | model.LfPeace | model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace36, // N
			InterplanetarySpace37, // NE
			InterplanetarySpace43, // E
			SolarOrbit3,           // SE
			SolarOrbit2,           // S
			SolarOrbit1,           // SW
			InterplanetarySpace42, // W
			InterplanetarySpace35, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			0,                     // Planet
		},
	},
	{
		Number:     InterplanetarySpace43,
		BriefMsgNo: text.SolLocation61_Brief,
		FullMsgNo:  text.SolLocation61_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace37, // N
			InterplanetarySpace38, // NE
			InterplanetarySpace44, // E
			InterplanetarySpace45, // SE
			SolarOrbit3,           // S
			SolarOrbit2,           // SW
			MercuryOrbit,          // W
			InterplanetarySpace36, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			MercuryOrbit,          // Planet
		},
	},
	{
		Number:     InterplanetarySpace44,
		BriefMsgNo: text.SolLocation62_Brief,
		FullMsgNo:  text.SolLocation62_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace38, // N
			0,                     // NE
			0,                     // E
			0,                     // SE
			InterplanetarySpace45, // S
			SolarOrbit3,           // SW
			InterplanetarySpace43, // W
			InterplanetarySpace37, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			InterplanetarySpace43, // Planet
		},
	},
	{
		Number:     SolarOrbit1,
		BriefMsgNo: text.SolLocation63_Brief,
		FullMsgNo:  text.SolLocation63_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace42, // N
			MercuryOrbit,          // NE
			SolarOrbit2,           // E
			SolarSurface,          // SE
			SolarOrbit4,           // S
			0,                     // SW
			0,                     // W
			InterplanetarySpace41, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			MercuryOrbit,          // Planet
		},
	},
	{
		Number:     SolarOrbit2,
		BriefMsgNo: text.SolLocation64_Brief,
		FullMsgNo:  text.SolLocation64_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			MercuryOrbit,          // N
			InterplanetarySpace43, // NE
			SolarOrbit3,           // E
			SolarOrbit5,           // SE
			SolarSurface,          // S
			SolarOrbit4,           // SW
			SolarOrbit1,           // W
			InterplanetarySpace42, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			MercuryOrbit,          // Planet
		},
	},
	{
		Number:     SolarOrbit3,
		BriefMsgNo: text.SolLocation65_Brief,
		FullMsgNo:  text.SolLocation65_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace43, // N
			InterplanetarySpace44, // NE
			InterplanetarySpace45, // E
			0,                     // SE
			SolarOrbit5,           // S
			SolarSurface,          // SW
			SolarOrbit2,           // W
			MercuryOrbit,          // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			MercuryOrbit,          // Planet
		},
	},
	{
		Number:     InterplanetarySpace45,
		BriefMsgNo: text.SolLocation66_Brief,
		FullMsgNo:  text.SolLocation66_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			InterplanetarySpace44, // N
			0,                     // NE
			0,                     // E
			0,                     // SE
			0,                     // S
			SolarOrbit5,           // SW
			SolarOrbit3,           // W
			InterplanetarySpace43, // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			InterplanetarySpace43, // Planet
		},
	},
	{
		Number:     SolarOrbit4,
		BriefMsgNo: text.SolLocation67_Brief,
		FullMsgNo:  text.SolLocation67_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			SolarOrbit1,  // N
			SolarOrbit2,  // NE
			SolarSurface, // E
			0,            // SE
			0,            // S
			0,            // SW
			0,            // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
			SolarOrbit1,  // Planet
		},
	},
	{
		Number:     SolarSurface,
		BriefMsgNo: text.SolLocation68_Brief,
		FullMsgNo:  text.SolLocation68_Full,
		Flags:      model.LfDeath | model.LfSpace,
	},
	{
		Number:     SolarOrbit5,
		BriefMsgNo: text.SolLocation69_Brief,
		FullMsgNo:  text.SolLocation69_Full,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			SolarOrbit3,           // N
			InterplanetarySpace45, // NE
			0,                     // E
			0,                     // SE
			0,                     // S
			0,                     // SW
			SolarSurface,          // W
			SolarOrbit2,           // NW
			0,                     // Up
			0,                     // Down
			0,                     // In
			0,                     // Out
			SolarOrbit3,           // Planet
		},
	},
	{
		Number:     Lift1,
		BriefMsgNo: text.SolLocation70_Brief,
		FullMsgNo:  text.SolLocation70_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			0,         // SE
			Reception, // S
			0,         // SW
			0,         // W
			0,         // NW
			Lift2,     // Up
			0,         // Down
			0,         // In
			Reception, // Out
		},
	},
	{
		Number:     Office1,
		BriefMsgNo: text.SolLocation71_Brief,
		FullMsgNo:  text.SolLocation71_Full,

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
		Number:     Corridor1,
		BriefMsgNo: text.SolLocation72_Brief,
		FullMsgNo:  text.SolLocation72_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			Reception, // E
			0,         // SE
			0,         // S
			0,         // SW
			Office1,   // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     Reception,
		BriefMsgNo: text.SolLocation73_Brief,
		FullMsgNo:  text.SolLocation73_Full,

		MovTab: [13]uint16{
			Lift1,            // N
			0,                // NE
			TradingExchange1, // E
			0,                // SE
			Corridor2,        // S
			0,                // SW
			Corridor1,        // W
			0,                // NW
			0,                // Up
			0,                // Down
			0,                // In
			Lift1,            // Out
		},
	},
	{
		Number:     Corridor2,
		BriefMsgNo: text.SolLocation74_Brief,
		FullMsgNo:  text.SolLocation74_Full,

		MovTab: [13]uint16{
			Reception,    // N
			0,            // NE
			0,            // E
			0,            // SE
			Corridor3,    // S
			0,            // SW
			SmallOffice1, // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     Corridor3,
		BriefMsgNo: text.SolLocation75_Brief,
		FullMsgNo:  text.SolLocation75_Full,

		MovTab: [13]uint16{
			Corridor2,       // N
			0,               // NE
			CrampedQuarters, // E
			0,               // SE
			ControlRoom1,    // S
			0,               // SW
			LivingQuarters1, // W
			0,               // NW
			0,               // Up
			0,               // Down
			0,               // In
			0,               // Out
		},
	},
	{
		Number:     ControlRoom1,
		BriefMsgNo: text.SolLocation76_Brief,
		FullMsgNo:  text.SolLocation76_Full,

		MovTab: [13]uint16{
			Corridor3, // N
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
			Corridor3, // Out
		},
	},
	{
		Number:     LivingQuarters1,
		BriefMsgNo: text.SolLocation77_Brief,
		FullMsgNo:  text.SolLocation77_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			Corridor3, // E
			0,         // SE
			0,         // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			Corridor3, // Out
		},
	},
	{
		Number:     CrampedQuarters,
		BriefMsgNo: text.SolLocation78_Brief,
		FullMsgNo:  text.SolLocation78_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			0,         // SE
			0,         // S
			0,         // SW
			Corridor3, // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			Corridor3, // Out
		},
	},
	{
		Number:     SmallOffice1,
		BriefMsgNo: text.SolLocation79_Brief,
		FullMsgNo:  text.SolLocation79_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			Corridor2, // E
			0,         // SE
			0,         // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			Corridor2, // Out
		},
	},
	{
		Number:     TradingExchange1,
		BriefMsgNo: text.SolLocation80_Brief,
		FullMsgNo:  text.SolLocation80_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Fedruckers, // E
			0,          // SE
			0,          // S
			0,          // SW
			Reception,  // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Fedruckers,
		BriefMsgNo: text.SolLocation81_Brief,
		FullMsgNo:  text.SolLocation81_Full,
		Flags:      model.LfCafe,

		MovTab: [13]uint16{
			0,                // N
			0,                // NE
			0,                // E
			0,                // SE
			0,                // S
			0,                // SW
			TradingExchange1, // W
			0,                // NW
			0,                // Up
			0,                // Down
			0,                // In
			TradingExchange1, // Out
		},
	},
	{
		Number:     Lift2,
		BriefMsgNo: text.SolLocation82_Brief,
		FullMsgNo:  text.SolLocation82_Full,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			0,        // E
			0,        // SE
			Airlock1, // S
			0,        // SW
			0,        // W
			0,        // NW
			0,        // Up
			Lift1,    // Down
			0,        // In
			Airlock1, // Out
		},
	},
	{
		Number:     Airlock1,
		BriefMsgNo: text.SolLocation83_Brief,
		FullMsgNo:  text.SolLocation83_Full,

		MovTab: [13]uint16{
			Lift2,         // N
			0,             // NE
			0,             // E
			0,             // SE
			BeforeAirlock, // S
			0,             // SW
			0,             // W
			0,             // NW
			0,             // Up
			0,             // Down
			Lift2,         // In
			BeforeAirlock, // Out
		},
	},
	{
		Number:     BeforeAirlock,
		BriefMsgNo: text.SolLocation84_Brief,
		FullMsgNo:  text.SolLocation84_Full,
		Flags:      model.LfVacuum,

		MovTab: [13]uint16{
			Airlock1,         // N
			0,                // NE
			CrumblingSurface, // E
			UnevenSurface,    // SE
			CrateredSurface,  // S
			SteepWall,        // SW
			0,                // W
			0,                // NW
			0,                // Up
			0,                // Down
			Airlock1,         // In
			0,                // Out
		},
	},
	{
		Number:     CrumblingSurface,
		BriefMsgNo: text.SolLocation85_Brief,
		FullMsgNo:  text.SolLocation85_Full,
		Flags:      model.LfVacuum,

		MovTab: [13]uint16{
			0,               // N
			0,               // NE
			0,               // E
			BottomOfCliff,   // SE
			UnevenSurface,   // S
			CrateredSurface, // SW
			BeforeAirlock,   // W
			0,               // NW
			0,               // Up
			0,               // Down
			0,               // In
			0,               // Out
		},
	},
	{
		Number:     UnevenSurface,
		BriefMsgNo: text.SolLocation86_Brief,
		FullMsgNo:  text.SolLocation86_Full,
		Flags:      model.LfVacuum,

		MovTab: [13]uint16{
			CrumblingSurface,   // N
			0,                  // NE
			BottomOfCliff,      // E
			0,                  // SE
			SheerCliff,         // S
			CallistoLandingPad, // SW
			CrateredSurface,    // W
			BeforeAirlock,      // NW
			0,                  // Up
			0,                  // Down
			0,                  // In
			0,                  // Out
		},
	},
	{
		Number:     CrateredSurface,
		BriefMsgNo: text.SolLocation87_Brief,
		FullMsgNo:  text.SolLocation87_Full,
		Flags:      model.LfVacuum,

		MovTab: [13]uint16{
			BeforeAirlock,      // N
			CrumblingSurface,   // NE
			UnevenSurface,      // E
			SheerCliff,         // SE
			CallistoLandingPad, // S
			0,                  // SW
			SteepWall,          // W
			0,                  // NW
			0,                  // Up
			0,                  // Down
			0,                  // In
			0,                  // Out
		},
	},
	{
		Number:     SteepWall,
		BriefMsgNo: text.SolLocation88_Brief,
		FullMsgNo:  text.SolLocation88_Full,
		Flags:      model.LfVacuum,

		MovTab: [13]uint16{
			0,                  // N
			BeforeAirlock,      // NE
			CrateredSurface,    // E
			CallistoLandingPad, // SE
			0,                  // S
			0,                  // SW
			0,                  // W
			0,                  // NW
			0,                  // Up
			0,                  // Down
			0,                  // In
			0,                  // Out
		},
	},
	{
		Number:     BottomOfCliff,
		BriefMsgNo: text.SolLocation89_Brief,
		FullMsgNo:  text.SolLocation89_Full,
		Flags:      model.LfVacuum,

		MovTab: [13]uint16{
			0,                // N
			0,                // NE
			0,                // E
			0,                // SE
			0,                // S
			SheerCliff,       // SW
			UnevenSurface,    // W
			CrumblingSurface, // NW
			0,                // Up
			0,                // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     SheerCliff,
		BriefMsgNo: text.SolLocation90_Brief,
		FullMsgNo:  text.SolLocation90_Full,
		Flags:      model.LfVacuum,

		MovTab: [13]uint16{
			UnevenSurface,      // N
			BottomOfCliff,      // NE
			0,                  // E
			0,                  // SE
			0,                  // S
			0,                  // SW
			CallistoLandingPad, // W
			CrateredSurface,    // NW
			0,                  // Up
			0,                  // Down
			0,                  // In
			0,                  // Out
		},
	},
	{
		Number:     CallistoLandingPad,
		BriefMsgNo: text.SolLocation91_Brief,
		FullMsgNo:  text.SolLocation91_Full,
		Flags:      model.LfLanding | model.LfVacuum,

		MovTab: [13]uint16{
			CrateredSurface, // N
			UnevenSurface,   // NE
			SheerCliff,      // E
			0,               // SE
			0,               // S
			0,               // SW
			0,               // W
			SteepWall,       // NW
			0,               // Up
			0,               // Down
			0,               // In
			0,               // Out
		},
	},
	{
		Number:     LinkBetweenDomes1,
		BriefMsgNo: text.SolLocation92_Brief,
		FullMsgNo:  text.SolLocation92_Full,

		MovTab: [13]uint16{
			MainDome2, // N
			0,         // NE
			0,         // E
			0,         // SE
			Lobby,     // S
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
		Number:     Link1,
		BriefMsgNo: text.SolLocation93_Brief,
		FullMsgNo:  text.SolLocation93_Full,

		MovTab: [13]uint16{
			LinkBetweenDomes2, // N
			0,                 // NE
			0,                 // E
			0,                 // SE
			MainDome1,         // S
			0,                 // SW
			0,                 // W
			0,                 // NW
			0,                 // Up
			0,                 // Down
			0,                 // In
			0,                 // Out
		},
	},
	{
		Number:     LinkBetweenDomes2,
		BriefMsgNo: text.SolLocation94_Brief,
		FullMsgNo:  text.SolLocation94_Full,

		MovTab: [13]uint16{
			StorageDome, // N
			0,           // NE
			0,           // E
			0,           // SE
			Link1,       // S
			0,           // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     Link2,
		BriefMsgNo: text.SolLocation95_Brief,
		FullMsgNo:  text.SolLocation95_Full,

		MovTab: [13]uint16{
			0,                 // N
			0,                 // NE
			MainDome2,         // E
			0,                 // SE
			0,                 // S
			0,                 // SW
			LinkBetweenDomes3, // W
			0,                 // NW
			0,                 // Up
			0,                 // Down
			0,                 // In
			0,                 // Out
		},
	},
	{
		Number:     LinkBetweenDomes3,
		BriefMsgNo: text.SolLocation96_Brief,
		FullMsgNo:  text.SolLocation96_Full,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			Link2,        // E
			0,            // SE
			0,            // S
			0,            // SW
			DomeEntrance, // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     Link3,
		BriefMsgNo: text.SolLocation97_Brief,
		FullMsgNo:  text.SolLocation97_Full,

		MovTab: [13]uint16{
			0,               // N
			LinkingCorridor, // NE
			0,               // E
			0,               // SE
			0,               // S
			ViewingRoom,     // SW
			0,               // W
			0,               // NW
			0,               // Up
			0,               // Down
			0,               // In
			0,               // Out
		},
	},
	{
		Number:     LinkingCorridor,
		BriefMsgNo: text.SolLocation98_Brief,
		FullMsgNo:  text.SolLocation98_Full,

		MovTab: [13]uint16{
			Entrance, // N
			0,        // NE
			0,        // E
			0,        // SE
			0,        // S
			Link3,    // SW
			0,        // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     Lobby,
		BriefMsgNo: text.SolLocation99_Brief,
		FullMsgNo:  text.SolLocation99_Full,

		MovTab: [13]uint16{
			LinkBetweenDomes1, // N
			0,                 // NE
			ThoriumStoreroom,  // E
			UraniumStoreroom,  // SE
			TransuranicsStore, // S
			0,                 // SW
			0,                 // W
			0,                 // NW
			0,                 // Up
			0,                 // Down
			0,                 // In
			LinkBetweenDomes1, // Out
		},
	},
	{
		Number:     ThoriumStoreroom,
		BriefMsgNo: text.SolLocation100_Brief,
		FullMsgNo:  text.SolLocation100_Full,
		Flags:      model.LfDeath,
	},
	{
		Number:     TransuranicsStore,
		BriefMsgNo: text.SolLocation101_Brief,
		FullMsgNo:  text.SolLocation101_Full,
		Flags:      model.LfDeath,
	},
	{
		Number:     UraniumStoreroom,
		BriefMsgNo: text.SolLocation102_Brief,
		FullMsgNo:  text.SolLocation102_Full,
		Flags:      model.LfDeath,
	},
	{
		Number:     StorageDome,
		BriefMsgNo: text.SolLocation103_Brief,
		FullMsgNo:  text.SolLocation103_Full,

		MovTab: [13]uint16{
			0,                 // N
			0,                 // NE
			Store,             // E
			0,                 // SE
			LinkBetweenDomes2, // S
			0,                 // SW
			0,                 // W
			0,                 // NW
			0,                 // Up
			0,                 // Down
			0,                 // In
			LinkBetweenDomes2, // Out
		},
	},
	{
		Number:     Store,
		BriefMsgNo: text.SolLocation104_Brief,
		FullMsgNo:  text.SolLocation104_Full,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			0,           // E
			0,           // SE
			0,           // S
			0,           // SW
			StorageDome, // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			StorageDome, // Out
		},
	},
	{
		Number:     MainDome1,
		BriefMsgNo: text.SolLocation105_Brief,
		FullMsgNo:  text.SolLocation105_Full,

		MovTab: [13]uint16{
			Link1,     // N
			0,         // NE
			Airlock2,  // E
			0,         // SE
			Market1,   // S
			MainDome2, // SW
			SmallPark, // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			Airlock2,  // Out
		},
	},
	{
		Number:     MainDome2,
		BriefMsgNo: text.SolLocation106_Brief,
		FullMsgNo:  text.SolLocation106_Full,

		MovTab: [13]uint16{
			SmallPark,         // N
			MainDome1,         // NE
			Market1,           // E
			0,                 // SE
			LinkBetweenDomes1, // S
			0,                 // SW
			Link2,             // W
			GeneralStore,      // NW
			0,                 // Up
			0,                 // Down
			0,                 // In
			0,                 // Out
		},
	},
	{
		Number:     GeneralStore,
		BriefMsgNo: text.SolLocation107_Brief,
		FullMsgNo:  text.SolLocation107_Full,
		Flags:      model.LfGen,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			MainDome2, // SE
			0,         // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			MainDome2, // Out
		},
	},
	{
		Number:     SmallPark,
		BriefMsgNo: text.SolLocation108_Brief,
		FullMsgNo:  text.SolLocation108_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			MainDome1, // E
			Market1,   // SE
			MainDome2, // S
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
		Number:     Market1,
		BriefMsgNo: text.SolLocation109_Brief,
		FullMsgNo:  text.SolLocation109_Full,
		Flags:      model.LfTrade,

		MovTab: [13]uint16{
			MainDome1, // N
			0,         // NE
			0,         // E
			0,         // SE
			0,         // S
			0,         // SW
			MainDome2, // W
			SmallPark, // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     DomeEntrance,
		BriefMsgNo: text.SolLocation110_Brief,
		FullMsgNo:  text.SolLocation110_Full,

		MovTab: [13]uint16{
			0,                 // N
			0,                 // NE
			LinkBetweenDomes3, // E
			0,                 // SE
			0,                 // S
			0,                 // SW
			Corridor6,         // W
			0,                 // NW
			0,                 // Up
			0,                 // Down
			Corridor6,         // In
			LinkBetweenDomes3, // Out
		},
	},
	{
		Number:     Corridor4,
		BriefMsgNo: text.SolLocation111_Brief,
		FullMsgNo:  text.SolLocation111_Full,

		MovTab: [13]uint16{
			Corridor5,        // N
			0,                // NE
			WestHall,         // E
			0,                // SE
			Entrance,         // S
			0,                // SW
			ProductionOffice, // W
			0,                // NW
			0,                // Up
			0,                // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     Corridor5,
		BriefMsgNo: text.SolLocation112_Brief,
		FullMsgNo:  text.SolLocation112_Full,

		MovTab: [13]uint16{
			Canteen,    // N
			0,          // NE
			Laboratory, // E
			0,          // SE
			Corridor4,  // S
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
		Number:     Corridor6,
		BriefMsgNo: text.SolLocation113_Brief,
		FullMsgNo:  text.SolLocation113_Full,

		MovTab: [13]uint16{
			ViewingRoom,     // N
			0,               // NE
			DomeEntrance,    // E
			0,               // SE
			LivingQuarters2, // S
			0,               // SW
			ControlRoom2,    // W
			0,               // NW
			0,               // Up
			0,               // Down
			0,               // In
			0,               // Out
		},
	},
	{
		Number:     LivingQuarters2,
		BriefMsgNo: text.SolLocation114_Brief,
		FullMsgNo:  text.SolLocation114_Full,

		MovTab: [13]uint16{
			Corridor6, // N
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
			Corridor6, // Out
		},
	},
	{
		Number:     ControlRoom2,
		BriefMsgNo: text.SolLocation115_Brief,
		Events:     [2]uint16{0, 5},
		FullMsgNo:  text.SolLocation115_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			Corridor6, // E
			0,         // SE
			0,         // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			Corridor6, // Out
		},
	},
	{
		Number:     ViewingRoom,
		BriefMsgNo: text.SolLocation116_Brief,
		FullMsgNo:  text.SolLocation116_Full,

		MovTab: [13]uint16{
			0,         // N
			Link3,     // NE
			0,         // E
			0,         // SE
			Corridor6, // S
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
		Number:     ProductionOffice,
		BriefMsgNo: text.SolLocation117_Brief,
		FullMsgNo:  text.SolLocation117_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			Corridor4, // E
			0,         // SE
			0,         // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			Corridor4, // Out
		},
	},
	{
		Number:     ProcessingRoom1,
		BriefMsgNo: text.SolLocation118_Brief,
		FullMsgNo:  text.SolLocation118_Full,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			Entrance, // E
			0,        // SE
			0,        // S
			0,        // SW
			0,        // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			Entrance, // Out
		},
	},
	{
		Number:     Entrance,
		BriefMsgNo: text.SolLocation119_Brief,
		FullMsgNo:  text.SolLocation119_Full,

		MovTab: [13]uint16{
			Corridor4,       // N
			0,               // NE
			0,               // E
			0,               // SE
			LinkingCorridor, // S
			0,               // SW
			ProcessingRoom1, // W
			0,               // NW
			0,               // Up
			0,               // Down
			0,               // In
			0,               // Out
		},
	},
	{
		Number:     WestHall,
		BriefMsgNo: text.SolLocation120_Brief,
		FullMsgNo:  text.SolLocation120_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			EastHall,  // E
			0,         // SE
			0,         // S
			0,         // SW
			Corridor4, // W
			0,         // NW
			0,         // Up
			0,         // Down
			EastHall,  // In
			Corridor4, // Out
		},
	},
	{
		Number:     Laboratory,
		BriefMsgNo: text.SolLocation121_Brief,
		FullMsgNo:  text.SolLocation121_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			0,         // SE
			0,         // S
			0,         // SW
			Corridor5, // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			Corridor5, // Out
		},
	},
	{
		Number:     Canteen,
		BriefMsgNo: text.SolLocation122_Brief,
		FullMsgNo:  text.SolLocation122_Full,
		Flags:      model.LfCafe,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			0,         // SE
			Corridor5, // S
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
		Number:     EastHall,
		BriefMsgNo: text.SolLocation123_Brief,
		FullMsgNo:  text.SolLocation123_Full,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			0,        // E
			0,        // SE
			0,        // S
			0,        // SW
			WestHall, // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			WestHall, // Out
		},
	},
	{
		Number:     Airlock2,
		BriefMsgNo: text.SolLocation124_Brief,
		FullMsgNo:  text.SolLocation124_Full,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			TitanSurface3, // E
			0,             // SE
			0,             // S
			0,             // SW
			MainDome1,     // W
			0,             // NW
			0,             // Up
			0,             // Down
			MainDome1,     // In
			TitanSurface3, // Out
		},
	},
	{
		Number:     TitanSurface1,
		BriefMsgNo: text.SolLocation125_Brief,
		FullMsgNo:  text.SolLocation125_Full,
		Flags:      model.LfVacuum,

		MovTab: [13]uint16{
			0,                // N
			0,                // NE
			TitanSurface2,    // E
			TitanLandingArea, // SE
			TitanSurface3,    // S
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
		Number:     TitanSurface2,
		BriefMsgNo: text.SolLocation126_Brief,
		FullMsgNo:  text.SolLocation126_Full,
		Flags:      model.LfVacuum,

		MovTab: [13]uint16{
			0,                // N
			0,                // NE
			0,                // E
			0,                // SE
			TitanLandingArea, // S
			TitanSurface3,    // SW
			TitanSurface1,    // W
			0,                // NW
			0,                // Up
			0,                // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     TitanSurface3,
		BriefMsgNo: text.SolLocation127_Brief,
		FullMsgNo:  text.SolLocation127_Full,
		Flags:      model.LfVacuum,

		MovTab: [13]uint16{
			TitanSurface1,    // N
			TitanSurface2,    // NE
			TitanLandingArea, // E
			TitanSurface5,    // SE
			TitanSurface4,    // S
			0,                // SW
			Airlock2,         // W
			0,                // NW
			0,                // Up
			0,                // Down
			Airlock2,         // In
			0,                // Out
		},
	},
	{
		Number:     TitanSurface4,
		BriefMsgNo: text.SolLocation128_Brief,
		FullMsgNo:  text.SolLocation128_Full,
		Flags:      model.LfVacuum,

		MovTab: [13]uint16{
			TitanSurface3,    // N
			TitanLandingArea, // NE
			TitanSurface5,    // E
			0,                // SE
			0,                // S
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
		Number:     TitanSurface5,
		BriefMsgNo: text.SolLocation129_Brief,
		FullMsgNo:  text.SolLocation129_Full,
		Flags:      model.LfVacuum,

		MovTab: [13]uint16{
			TitanLandingArea, // N
			0,                // NE
			0,                // E
			0,                // SE
			0,                // S
			0,                // SW
			TitanSurface4,    // W
			TitanSurface3,    // NW
			0,                // Up
			0,                // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     TitanLandingArea,
		BriefMsgNo: text.SolLocation130_Brief,
		FullMsgNo:  text.SolLocation130_Full,
		Flags:      model.LfLanding | model.LfVacuum,

		MovTab: [13]uint16{
			TitanSurface2, // N
			0,             // NE
			0,             // E
			0,             // SE
			TitanSurface5, // S
			TitanSurface4, // SW
			TitanSurface3, // W
			TitanSurface1, // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     SelenaLandingPad,
		BriefMsgNo: text.SolLocation131_Brief,
		FullMsgNo:  text.SolLocation131_Full,
		Flags:      model.LfLanding | model.LfVacuum,

		MovTab: [13]uint16{
			MoonSurface2,  // N
			0,             // NE
			MoonSurface3,  // E
			MoonRay1,      // SE
			MoonSurface4,  // S
			MoonsSurface2, // SW
			MoonsSurface1, // W
			MoonSurface1,  // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     MoonSurface1,
		BriefMsgNo: text.SolLocation132_Brief,
		FullMsgNo:  text.SolLocation132_Full,
		Flags:      model.LfVacuum,

		MovTab: [13]uint16{
			0,                // N
			0,                // NE
			MoonSurface2,     // E
			SelenaLandingPad, // SE
			MoonsSurface1,    // S
			0,                // SW
			0,                // W
			Airlock3,         // NW
			0,                // Up
			0,                // Down
			Airlock3,         // In
			0,                // Out
		},
	},
	{
		Number:     MoonSurface2,
		BriefMsgNo: text.SolLocation133_Brief,
		FullMsgNo:  text.SolLocation133_Full,
		Flags:      model.LfVacuum,

		MovTab: [13]uint16{
			0,                // N
			0,                // NE
			0,                // E
			MoonSurface3,     // SE
			SelenaLandingPad, // S
			MoonsSurface1,    // SW
			MoonSurface1,     // W
			0,                // NW
			0,                // Up
			0,                // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     MoonsSurface1,
		BriefMsgNo: text.SolLocation134_Brief,
		FullMsgNo:  text.SolLocation134_Full,
		Flags:      model.LfVacuum,

		MovTab: [13]uint16{
			MoonSurface1,     // N
			MoonSurface2,     // NE
			SelenaLandingPad, // E
			MoonSurface4,     // SE
			MoonsSurface2,    // S
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
		Number:     MoonSurface3,
		BriefMsgNo: text.SolLocation135_Brief,
		FullMsgNo:  text.SolLocation135_Full,
		Flags:      model.LfVacuum,

		MovTab: [13]uint16{
			0,                // N
			0,                // NE
			0,                // E
			0,                // SE
			MoonRay1,         // S
			MoonSurface4,     // SW
			SelenaLandingPad, // W
			MoonSurface2,     // NW
			0,                // Up
			0,                // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     MoonsSurface2,
		BriefMsgNo: text.SolLocation136_Brief,
		FullMsgNo:  text.SolLocation136_Full,
		Flags:      model.LfVacuum,

		MovTab: [13]uint16{
			MoonsSurface1,    // N
			SelenaLandingPad, // NE
			MoonSurface4,     // E
			MoonRay2,         // SE
			0,                // S
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
		Number:     MoonSurface4,
		BriefMsgNo: text.SolLocation137_Brief,
		FullMsgNo:  text.SolLocation137_Full,
		Flags:      model.LfVacuum,

		MovTab: [13]uint16{
			SelenaLandingPad, // N
			MoonSurface3,     // NE
			MoonRay1,         // E
			0,                // SE
			MoonRay2,         // S
			0,                // SW
			MoonsSurface2,    // W
			MoonsSurface1,    // NW
			0,                // Up
			0,                // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     MoonRay1,
		BriefMsgNo: text.SolLocation138_Brief,
		FullMsgNo:  text.SolLocation138_Full,
		Flags:      model.LfVacuum,

		MovTab: [13]uint16{
			MoonSurface3,     // N
			0,                // NE
			0,                // E
			0,                // SE
			0,                // S
			MoonRay2,         // SW
			MoonSurface4,     // W
			SelenaLandingPad, // NW
			0,                // Up
			0,                // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     MoonRay2,
		BriefMsgNo: text.SolLocation139_Brief,
		FullMsgNo:  text.SolLocation139_Full,
		Flags:      model.LfVacuum,

		MovTab: [13]uint16{
			MoonSurface4,  // N
			MoonRay1,      // NE
			0,             // E
			0,             // SE
			0,             // S
			MoonRay3,      // SW
			0,             // W
			MoonsSurface2, // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     MoonRay3,
		BriefMsgNo: text.SolLocation140_Brief,
		FullMsgNo:  text.SolLocation140_Full,
		Flags:      model.LfVacuum,

		MovTab: [13]uint16{
			0,        // N
			MoonRay2, // NE
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
		Number:     Airlock3,
		BriefMsgNo: text.SolLocation141_Brief,
		FullMsgNo:  text.SolLocation141_Full,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			0,            // E
			MoonSurface1, // SE
			0,            // S
			0,            // SW
			0,            // W
			MainEntrance, // NW
			0,            // Up
			0,            // Down
			MainEntrance, // In
			MoonSurface1, // Out
		},
	},
	{
		Number:     MainEntrance,
		BriefMsgNo: text.SolLocation142_Brief,
		FullMsgNo:  text.SolLocation142_Full,

		MovTab: [13]uint16{
			Park1,          // N
			Park2,          // NE
			Park3,          // E
			Airlock3,       // SE
			ViewingLounge3, // S
			ViewingLounge2, // SW
			ViewingLounge1, // W
			Lift3,          // NW
			0,              // Up
			0,              // Down
			Lift3,          // In
			Airlock3,       // Out
		},
	},
	{
		Number:     Lift3,
		BriefMsgNo: text.SolLocation143_Brief,
		FullMsgNo:  text.SolLocation143_Full,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			0,            // E
			MainEntrance, // SE
			0,            // S
			0,            // SW
			0,            // W
			0,            // NW
			0,            // Up
			Lift4,        // Down
			0,            // In
			MainEntrance, // Out
		},
	},
	{
		Number:     Lift4,
		BriefMsgNo: text.SolLocation144_Brief,
		FullMsgNo:  text.SolLocation144_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			Concourse, // SE
			0,         // S
			0,         // SW
			0,         // W
			0,         // NW
			Lift3,     // Up
			Lift7,     // Down
			0,         // In
			Concourse, // Out
		},
	},
	{
		Number:     Park1,
		BriefMsgNo: text.SolLocation145_Brief,
		FullMsgNo:  text.SolLocation145_Full,

		MovTab: [13]uint16{
			0,              // N
			0,              // NE
			Park2,          // E
			Park3,          // SE
			MainEntrance,   // S
			ViewingLounge1, // SW
			0,              // W
			0,              // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     Park2,
		BriefMsgNo: text.SolLocation146_Brief,
		FullMsgNo:  text.SolLocation146_Full,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			0,            // E
			0,            // SE
			Park3,        // S
			MainEntrance, // SW
			Park1,        // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     Park3,
		BriefMsgNo: text.SolLocation147_Brief,
		FullMsgNo:  text.SolLocation147_Full,

		MovTab: [13]uint16{
			Park2,        // N
			0,            // NE
			0,            // E
			0,            // SE
			0,            // S
			0,            // SW
			MainEntrance, // W
			Park1,        // NW
			0,            // Up
			0,            // Down
			0,            // In
			MainEntrance, // Out
		},
	},
	{
		Number:     Cafe,
		BriefMsgNo: text.SolLocation148_Brief,
		FullMsgNo:  text.SolLocation148_Full,
		Flags:      model.LfCafe,

		MovTab: [13]uint16{
			0,              // N
			0,              // NE
			ViewingLounge1, // E
			0,              // SE
			0,              // S
			0,              // SW
			0,              // W
			0,              // NW
			0,              // Up
			0,              // Down
			0,              // In
			ViewingLounge1, // Out
		},
	},
	{
		Number:     ViewingLounge1,
		BriefMsgNo: text.SolLocation149_Brief,
		FullMsgNo:  text.SolLocation149_Full,

		MovTab: [13]uint16{
			0,              // N
			Park1,          // NE
			MainEntrance,   // E
			ViewingLounge3, // SE
			ViewingLounge2, // S
			0,              // SW
			Cafe,           // W
			0,              // NW
			0,              // Up
			0,              // Down
			Cafe,           // In
			0,              // Out
		},
	},
	{
		Number:     ViewingLounge2,
		BriefMsgNo: text.SolLocation150_Brief,
		FullMsgNo:  text.SolLocation150_Full,

		MovTab: [13]uint16{
			ViewingLounge1, // N
			MainEntrance,   // NE
			ViewingLounge3, // E
			0,              // SE
			0,              // S
			0,              // SW
			0,              // W
			0,              // NW
			0,              // Up
			0,              // Down
			0,              // In
			MainEntrance,   // Out
		},
	},
	{
		Number:     ViewingLounge3,
		BriefMsgNo: text.SolLocation151_Brief,
		FullMsgNo:  text.SolLocation151_Full,

		MovTab: [13]uint16{
			MainEntrance,   // N
			0,              // NE
			0,              // E
			0,              // SE
			0,              // S
			0,              // SW
			ViewingLounge2, // W
			ViewingLounge1, // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     Corridor7,
		BriefMsgNo: text.SolLocation152_Brief,
		FullMsgNo:  text.SolLocation152_Full,

		MovTab: [13]uint16{
			Corridor8, // N
			0,         // NE
			Lift5,     // E
			0,         // SE
			0,         // S
			Concourse, // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			Lift5,     // In
			0,         // Out
		},
	},
	{
		Number:     Corridor8,
		BriefMsgNo: text.SolLocation153_Brief,
		FullMsgNo:  text.SolLocation153_Full,

		MovTab: [13]uint16{
			Corridor9,  // N
			0,          // NE
			Dormitory5, // E
			0,          // SE
			Corridor7,  // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			Dormitory5, // In
			0,          // Out
		},
	},
	{
		Number:     Corridor9,
		BriefMsgNo: text.SolLocation154_Brief,
		FullMsgNo:  text.SolLocation154_Full,

		MovTab: [13]uint16{
			Corridor10, // N
			0,          // NE
			Dormitory4, // E
			0,          // SE
			Corridor8,  // S
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
		Number:     Corridor10,
		BriefMsgNo: text.SolLocation155_Brief,
		FullMsgNo:  text.SolLocation155_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			Corridor9,  // S
			0,          // SW
			Dormitory1, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Corridor11,
		BriefMsgNo: text.SolLocation156_Brief,
		FullMsgNo:  text.SolLocation156_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			Corridor12, // S
			0,          // SW
			SmallPlaza, // W
			Concourse,  // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Corridor12,
		BriefMsgNo: text.SolLocation157_Brief,
		FullMsgNo:  text.SolLocation157_Full,

		MovTab: [13]uint16{
			Corridor11, // N
			0,          // NE
			0,          // E
			0,          // SE
			Corridor13, // S
			0,          // SW
			0,          // W
			SmallPlaza, // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Corridor13,
		BriefMsgNo: text.SolLocation158_Brief,
		FullMsgNo:  text.SolLocation158_Full,

		MovTab: [13]uint16{
			Corridor12, // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			Corridor14, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Corridor14,
		BriefMsgNo: text.SolLocation159_Brief,
		FullMsgNo:  text.SolLocation159_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Corridor13, // E
			0,          // SE
			0,          // S
			Corridor15, // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Corridor15,
		BriefMsgNo: text.SolLocation160_Brief,
		FullMsgNo:  text.SolLocation160_Full,

		MovTab: [13]uint16{
			0,          // N
			Corridor14, // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			Corridor16, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Corridor16,
		BriefMsgNo: text.SolLocation161_Brief,
		FullMsgNo:  text.SolLocation161_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Corridor15, // E
			0,          // SE
			0,          // S
			Adverts,    // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Dormitory1,
		BriefMsgNo: text.SolLocation162_Brief,
		FullMsgNo:  text.SolLocation162_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Corridor10, // E
			0,          // SE
			Dormitory3, // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			Corridor10, // Out
		},
	},
	{
		Number:     Dormitory2,
		BriefMsgNo: text.SolLocation163_Brief,
		FullMsgNo:  text.SolLocation163_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			Dormitory4, // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			Dormitory4, // Out
		},
	},
	{
		Number:     Dormitory3,
		BriefMsgNo: text.SolLocation164_Brief,
		FullMsgNo:  text.SolLocation164_Full,

		MovTab: [13]uint16{
			Dormitory1, // N
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
			Dormitory1, // Out
		},
	},
	{
		Number:     Dormitory4,
		BriefMsgNo: text.SolLocation165_Brief,
		FullMsgNo:  text.SolLocation165_Full,

		MovTab: [13]uint16{
			Dormitory2, // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			Corridor9,  // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			Corridor9,  // Out
		},
	},
	{
		Number:     Dormitory5,
		BriefMsgNo: text.SolLocation166_Brief,
		FullMsgNo:  text.SolLocation166_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			0,         // SE
			0,         // S
			0,         // SW
			Corridor8, // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			Corridor8, // Out
		},
	},
	{
		Number:     Concourse,
		BriefMsgNo: text.SolLocation167_Brief,
		FullMsgNo:  text.SolLocation167_Full,

		MovTab: [13]uint16{
			0,                         // N
			Corridor7,                 // NE
			0,                         // E
			Corridor11,                // SE
			SmallPlaza,                // S
			SportsComplexEntranceHall, // SW
			Church1,                   // W
			Lift4,                     // NW
			0,                         // Up
			0,                         // Down
			0,                         // In
			0,                         // Out
		},
	},
	{
		Number:     Church1,
		BriefMsgNo: text.SolLocation168_Brief,
		FullMsgNo:  text.SolLocation168_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			Concourse, // E
			0,         // SE
			0,         // S
			0,         // SW
			Church2,   // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			Concourse, // Out
		},
	},
	{
		Number:     Church2,
		BriefMsgNo: text.SolLocation169_Brief,
		FullMsgNo:  text.SolLocation169_Full,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			Church1, // E
			0,       // SE
			0,       // S
			0,       // SW
			0,       // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			Church1, // Out
		},
	},
	{
		Number:     SmallPlaza,
		BriefMsgNo: text.SolLocation170_Brief,
		FullMsgNo:  text.SolLocation170_Full,

		MovTab: [13]uint16{
			Concourse,  // N
			0,          // NE
			Corridor11, // E
			Corridor12, // SE
			0,          // S
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
		Number:     ZeroGravityArena,
		BriefMsgNo: text.SolLocation171_Brief,
		FullMsgNo:  text.SolLocation171_Full,

		MovTab: [13]uint16{
			0,                         // N
			0,                         // NE
			SportsComplexEntranceHall, // E
			0,                         // SE
			0,                         // S
			0,                         // SW
			0,                         // W
			0,                         // NW
			0,                         // Up
			0,                         // Down
			0,                         // In
			SportsComplexEntranceHall, // Out
		},
	},
	{
		Number:     SportsComplexEntranceHall,
		BriefMsgNo: text.SolLocation172_Brief,
		FullMsgNo:  text.SolLocation172_Full,

		MovTab: [13]uint16{
			0,                // N
			Concourse,        // NE
			0,                // E
			SwimmingPool,     // SE
			HighGravityGym,   // S
			RunningTrack,     // SW
			ZeroGravityArena, // W
			0,                // NW
			0,                // Up
			0,                // Down
			0,                // In
			Concourse,        // Out
		},
	},
	{
		Number:     RunningTrack,
		BriefMsgNo: text.SolLocation173_Brief,
		FullMsgNo:  text.SolLocation173_Full,

		MovTab: [13]uint16{
			0,                         // N
			SportsComplexEntranceHall, // NE
			0,                         // E
			0,                         // SE
			0,                         // S
			0,                         // SW
			0,                         // W
			0,                         // NW
			0,                         // Up
			0,                         // Down
			0,                         // In
			SportsComplexEntranceHall, // Out
		},
	},
	{
		Number:     HighGravityGym,
		BriefMsgNo: text.SolLocation174_Brief,
		FullMsgNo:  text.SolLocation174_Full,

		MovTab: [13]uint16{
			SportsComplexEntranceHall, // N
			0,                         // NE
			0,                         // E
			0,                         // SE
			0,                         // S
			0,                         // SW
			0,                         // W
			0,                         // NW
			0,                         // Up
			0,                         // Down
			0,                         // In
			SportsComplexEntranceHall, // Out
		},
	},
	{
		Number:     Adverts,
		BriefMsgNo: text.SolLocation175_Brief,
		FullMsgNo:  text.SolLocation175_Full,

		MovTab: [13]uint16{
			SaunaBaths,   // N
			Corridor16,   // NE
			0,            // E
			Casino,       // SE
			SmartBar,     // S
			MovieTheater, // SW
			Marios,       // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			Corridor16,   // Out
		},
	},
	{
		Number:     Casino,
		BriefMsgNo: text.SolLocation176_Brief,
		FullMsgNo:  text.SolLocation176_Full,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			0,        // E
			0,        // SE
			0,        // S
			0,        // SW
			SmartBar, // W
			Adverts,  // NW
			0,        // Up
			0,        // Down
			0,        // In
			Adverts,  // Out
		},
	},
	{
		Number:     SmartBar,
		BriefMsgNo: text.SolLocation177_Brief,
		FullMsgNo:  text.SolLocation177_Full,
		Flags:      model.LfCafe,

		MovTab: [13]uint16{
			Adverts, // N
			0,       // NE
			Casino,  // E
			0,       // SE
			0,       // S
			0,       // SW
			0,       // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			Adverts, // Out
		},
	},
	{
		Number:     Marios,
		BriefMsgNo: text.SolLocation178_Brief,
		FullMsgNo:  text.SolLocation178_Full,
		Flags:      model.LfCafe,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			Adverts, // E
			0,       // SE
			0,       // S
			0,       // SW
			0,       // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			Adverts, // Out
		},
	},
	{
		Number:     MovieTheater,
		BriefMsgNo: text.SolLocation179_Brief,
		FullMsgNo:  text.SolLocation179_Full,

		MovTab: [13]uint16{
			0,       // N
			Adverts, // NE
			0,       // E
			0,       // SE
			0,       // S
			0,       // SW
			0,       // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			Adverts, // Out
		},
	},
	{
		Number:     Lift5,
		BriefMsgNo: text.SolLocation180_Brief,
		FullMsgNo:  text.SolLocation180_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			0,         // SE
			0,         // S
			0,         // SW
			Corridor7, // W
			0,         // NW
			0,         // Up
			Lift6,     // Down
			0,         // In
			Corridor7, // Out
		},
	},
	{
		Number:     Lift6,
		BriefMsgNo: text.SolLocation181_Brief,
		FullMsgNo:  text.SolLocation181_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			Workshop17, // SE
			0,          // S
			0,          // SW
			0,          // W
			0,          // NW
			Lift5,      // Up
			0,          // Down
			0,          // In
			Workshop17, // Out
		},
	},
	{
		Number:     Lift7,
		BriefMsgNo: text.SolLocation182_Brief,
		FullMsgNo:  text.SolLocation182_Full,

		MovTab: [13]uint16{
			0,               // N
			0,               // NE
			0,               // E
			CirculationArea, // SE
			0,               // S
			0,               // SW
			0,               // W
			0,               // NW
			Lift4,           // Up
			0,               // Down
			0,               // In
			CirculationArea, // Out
		},
	},
	{
		Number:     SwimmingPool,
		BriefMsgNo: text.SolLocation183_Brief,
		FullMsgNo:  text.SolLocation183_Full,

		MovTab: [13]uint16{
			0,                         // N
			0,                         // NE
			0,                         // E
			0,                         // SE
			0,                         // S
			0,                         // SW
			0,                         // W
			SportsComplexEntranceHall, // NW
			0,                         // Up
			0,                         // Down
			0,                         // In
			SportsComplexEntranceHall, // Out
		},
	},
	{
		Number:     SaunaBaths,
		BriefMsgNo: text.SolLocation184_Brief,
		FullMsgNo:  text.SolLocation184_Full,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			0,       // E
			0,       // SE
			Adverts, // S
			0,       // SW
			0,       // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			Adverts, // Out
		},
	},
	{
		Number:     Workshop1,
		BriefMsgNo: text.SolLocation185_Brief,
		FullMsgNo:  text.SolLocation185_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			Workshop2, // E
			Workshop6, // SE
			Workshop5, // S
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
		Number:     Workshop2,
		BriefMsgNo: text.SolLocation186_Brief,
		FullMsgNo:  text.SolLocation186_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			Workshop3, // E
			Workshop7, // SE
			Workshop6, // S
			Workshop5, // SW
			Workshop1, // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     Workshop3,
		BriefMsgNo: text.SolLocation187_Brief,
		FullMsgNo:  text.SolLocation187_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			Workshop4, // E
			Workshop8, // SE
			Workshop7, // S
			Workshop6, // SW
			Workshop2, // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     Workshop4,
		BriefMsgNo: text.SolLocation188_Brief,
		FullMsgNo:  text.SolLocation188_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			0,         // SE
			Workshop8, // S
			Workshop7, // SW
			Workshop3, // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     Workshop5,
		BriefMsgNo: text.SolLocation189_Brief,
		FullMsgNo:  text.SolLocation189_Full,

		MovTab: [13]uint16{
			Workshop1,  // N
			Workshop2,  // NE
			Workshop6,  // E
			Workshop10, // SE
			Workshop9,  // S
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
		Number:     Workshop6,
		BriefMsgNo: text.SolLocation190_Brief,
		FullMsgNo:  text.SolLocation190_Full,

		MovTab: [13]uint16{
			Workshop2,  // N
			Workshop3,  // NE
			Workshop7,  // E
			Workshop11, // SE
			Workshop10, // S
			Workshop9,  // SW
			Workshop5,  // W
			Workshop1,  // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Workshop7,
		BriefMsgNo: text.SolLocation191_Brief,
		FullMsgNo:  text.SolLocation191_Full,

		MovTab: [13]uint16{
			Workshop3,  // N
			Workshop4,  // NE
			Workshop8,  // E
			Workshop12, // SE
			Workshop11, // S
			Workshop10, // SW
			Workshop6,  // W
			Workshop2,  // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Workshop8,
		BriefMsgNo: text.SolLocation192_Brief,
		FullMsgNo:  text.SolLocation192_Full,

		MovTab: [13]uint16{
			Workshop4,  // N
			0,          // NE
			0,          // E
			0,          // SE
			Workshop12, // S
			Workshop11, // SW
			Workshop7,  // W
			Workshop3,  // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Workshop9,
		BriefMsgNo: text.SolLocation193_Brief,
		FullMsgNo:  text.SolLocation193_Full,

		MovTab: [13]uint16{
			Workshop5,  // N
			Workshop6,  // NE
			Workshop10, // E
			Workshop14, // SE
			Workshop13, // S
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
		Number:     Workshop10,
		BriefMsgNo: text.SolLocation194_Brief,
		FullMsgNo:  text.SolLocation194_Full,

		MovTab: [13]uint16{
			Workshop6,  // N
			Workshop7,  // NE
			Workshop11, // E
			Workshop15, // SE
			Workshop14, // S
			Workshop13, // SW
			Workshop9,  // W
			Workshop5,  // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Workshop11,
		BriefMsgNo: text.SolLocation195_Brief,
		FullMsgNo:  text.SolLocation195_Full,

		MovTab: [13]uint16{
			Workshop7,  // N
			Workshop8,  // NE
			Workshop12, // E
			Workshop16, // SE
			Workshop15, // S
			Workshop14, // SW
			Workshop10, // W
			Workshop6,  // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Workshop12,
		BriefMsgNo: text.SolLocation196_Brief,
		FullMsgNo:  text.SolLocation196_Full,

		MovTab: [13]uint16{
			Workshop8,  // N
			0,          // NE
			0,          // E
			0,          // SE
			Workshop16, // S
			Workshop15, // SW
			Workshop11, // W
			Workshop7,  // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Workshop13,
		BriefMsgNo: text.SolLocation197_Brief,
		FullMsgNo:  text.SolLocation197_Full,

		MovTab: [13]uint16{
			Workshop9,  // N
			Workshop10, // NE
			Workshop14, // E
			Workshop18, // SE
			Workshop17, // S
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
		Number:     Workshop14,
		BriefMsgNo: text.SolLocation198_Brief,
		FullMsgNo:  text.SolLocation198_Full,

		MovTab: [13]uint16{
			Workshop10, // N
			Workshop11, // NE
			Workshop15, // E
			Workshop19, // SE
			Workshop18, // S
			Workshop17, // SW
			Workshop13, // W
			Workshop9,  // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Workshop15,
		BriefMsgNo: text.SolLocation199_Brief,
		FullMsgNo:  text.SolLocation199_Full,

		MovTab: [13]uint16{
			Workshop11, // N
			Workshop12, // NE
			Workshop16, // E
			Workshop20, // SE
			Workshop19, // S
			Workshop18, // SW
			Workshop14, // W
			Workshop10, // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Workshop16,
		BriefMsgNo: text.SolLocation200_Brief,
		FullMsgNo:  text.SolLocation200_Full,

		MovTab: [13]uint16{
			Workshop12, // N
			0,          // NE
			0,          // E
			0,          // SE
			Workshop20, // S
			Workshop19, // SW
			Workshop15, // W
			Workshop11, // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Workshop17,
		BriefMsgNo: text.SolLocation201_Brief,
		FullMsgNo:  text.SolLocation201_Full,

		MovTab: [13]uint16{
			Workshop13, // N
			Workshop14, // NE
			Workshop18, // E
			Workshop22, // SE
			Workshop21, // S
			0,          // SW
			0,          // W
			Lift6,      // NW
			0,          // Up
			0,          // Down
			Lift6,      // In
			0,          // Out
		},
	},
	{
		Number:     Workshop18,
		BriefMsgNo: text.SolLocation202_Brief,
		FullMsgNo:  text.SolLocation202_Full,

		MovTab: [13]uint16{
			Workshop14, // N
			Workshop15, // NE
			Workshop19, // E
			Workshop23, // SE
			Workshop22, // S
			Workshop21, // SW
			Workshop17, // W
			Workshop13, // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Workshop19,
		BriefMsgNo: text.SolLocation203_Brief,
		FullMsgNo:  text.SolLocation203_Full,

		MovTab: [13]uint16{
			Workshop15, // N
			Workshop16, // NE
			Workshop20, // E
			Workshop24, // SE
			Workshop23, // S
			Workshop22, // SW
			Workshop18, // W
			Workshop14, // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Workshop20,
		BriefMsgNo: text.SolLocation204_Brief,
		FullMsgNo:  text.SolLocation204_Full,

		MovTab: [13]uint16{
			Workshop16, // N
			0,          // NE
			0,          // E
			0,          // SE
			Workshop24, // S
			Workshop23, // SW
			Workshop19, // W
			Workshop15, // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Workshop21,
		BriefMsgNo: text.SolLocation205_Brief,
		FullMsgNo:  text.SolLocation205_Full,

		MovTab: [13]uint16{
			Workshop17, // N
			Workshop18, // NE
			Workshop22, // E
			0,          // SE
			0,          // S
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
		Number:     Workshop22,
		BriefMsgNo: text.SolLocation206_Brief,
		FullMsgNo:  text.SolLocation206_Full,

		MovTab: [13]uint16{
			Workshop18, // N
			Workshop19, // NE
			Workshop23, // E
			0,          // SE
			0,          // S
			0,          // SW
			Workshop21, // W
			Workshop17, // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Workshop23,
		BriefMsgNo: text.SolLocation207_Brief,
		FullMsgNo:  text.SolLocation207_Full,

		MovTab: [13]uint16{
			Workshop19, // N
			Workshop20, // NE
			Workshop24, // E
			0,          // SE
			0,          // S
			Corridor17, // SW
			Workshop22, // W
			Workshop18, // NW
			0,          // Up
			0,          // Down
			0,          // In
			Corridor17, // Out
		},
	},
	{
		Number:     Workshop24,
		BriefMsgNo: text.SolLocation208_Brief,
		FullMsgNo:  text.SolLocation208_Full,

		MovTab: [13]uint16{
			Workshop20, // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			Workshop23, // W
			Workshop19, // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Corridor17,
		BriefMsgNo: text.SolLocation209_Brief,
		FullMsgNo:  text.SolLocation209_Full,

		MovTab: [13]uint16{
			0,          // N
			Workshop23, // NE
			0,          // E
			0,          // SE
			0,          // S
			Corridor18, // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Corridor18,
		BriefMsgNo: text.SolLocation210_Brief,
		FullMsgNo:  text.SolLocation210_Full,

		MovTab: [13]uint16{
			0,              // N
			Corridor17,     // NE
			0,              // E
			0,              // SE
			ForemansOffice, // S
			0,              // SW
			Corridor19,     // W
			0,              // NW
			0,              // Up
			0,              // Down
			ForemansOffice, // In
			0,              // Out
		},
	},
	{
		Number:     Corridor19,
		BriefMsgNo: text.SolLocation211_Brief,
		FullMsgNo:  text.SolLocation211_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Corridor18, // E
			0,          // SE
			0,          // S
			0,          // SW
			Corridor20, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Corridor20,
		BriefMsgNo: text.SolLocation212_Brief,
		FullMsgNo:  text.SolLocation212_Full,

		MovTab: [13]uint16{
			Corridor21,  // N
			0,           // NE
			Corridor19,  // E
			0,           // SE
			0,           // S
			0,           // SW
			WeaponShop1, // W
			0,           // NW
			0,           // Up
			0,           // Down
			WeaponShop1, // In
			0,           // Out
		},
	},
	{
		Number:     Corridor21,
		BriefMsgNo: text.SolLocation213_Brief,
		FullMsgNo:  text.SolLocation213_Full,

		MovTab: [13]uint16{
			0,               // N
			0,               // NE
			0,               // E
			0,               // SE
			Corridor20,      // S
			0,               // SW
			0,               // W
			CirculationArea, // NW
			0,               // Up
			0,               // Down
			0,               // In
			0,               // Out
		},
	},
	{
		Number:     RockFace1,
		BriefMsgNo: text.SolLocation214_Brief,
		FullMsgNo:  text.SolLocation214_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,          // N
			Corridor22, // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			Corridor22, // Out
		},
	},
	{
		Number:     Corridor22,
		BriefMsgNo: text.SolLocation215_Brief,
		FullMsgNo:  text.SolLocation215_Full,

		MovTab: [13]uint16{
			0,               // N
			CirculationArea, // NE
			0,               // E
			0,               // SE
			0,               // S
			RockFace1,       // SW
			0,               // W
			0,               // NW
			0,               // Up
			0,               // Down
			0,               // In
			0,               // Out
		},
	},
	{
		Number:     ForemansOffice,
		BriefMsgNo: text.SolLocation216_Brief,
		FullMsgNo:  text.SolLocation216_Full,

		MovTab: [13]uint16{
			Corridor18, // N
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
			Corridor18, // Out
		},
	},
	{
		Number:     Office2,
		BriefMsgNo: text.SolLocation217_Brief,
		FullMsgNo:  text.SolLocation217_Full,
		Flags:      model.LfLock,

		MovTab: [13]uint16{
			0,   // N
			0,   // NE
			0,   // E
			Bar, // SE
			0,   // S
			0,   // SW
			0,   // W
			0,   // NW
			0,   // Up
			0,   // Down
			0,   // In
			Bar, // Out
		},
	},
	{
		Number:     Bar,
		BriefMsgNo: text.SolLocation218_Brief,
		FullMsgNo:  text.SolLocation218_Full,
		Flags:      model.LfCafe,
		SysLoc:     text.NO_MOVEMENT_32,

		MovTab: [13]uint16{
			0,               // N
			0,               // NE
			0,               // E
			0,               // SE
			CirculationArea, // S
			0,               // SW
			0,               // W
			0,               // NW
			0,               // Up
			0,               // Down
			0,               // In
			CirculationArea, // Out
		},
	},
	{
		Number:     Market2,
		BriefMsgNo: text.SolLocation219_Brief,
		FullMsgNo:  text.SolLocation219_Full,
		Flags:      model.LfTrade,

		MovTab: [13]uint16{
			0,               // N
			0,               // NE
			0,               // E
			0,               // SE
			0,               // S
			0,               // SW
			CirculationArea, // W
			0,               // NW
			0,               // Up
			0,               // Down
			0,               // In
			CirculationArea, // Out
		},
	},
	{
		Number:     CirculationArea,
		BriefMsgNo: text.SolLocation220_Brief,
		FullMsgNo:  text.SolLocation220_Full,

		MovTab: [13]uint16{
			Bar,           // N
			0,             // NE
			Market2,       // E
			Corridor21,    // SE
			0,             // S
			Corridor22,    // SW
			JewelleryShop, // W
			Lift7,         // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     JewelleryShop,
		BriefMsgNo: text.SolLocation221_Brief,
		FullMsgNo:  text.SolLocation221_Full,

		MovTab: [13]uint16{
			0,               // N
			0,               // NE
			CirculationArea, // E
			0,               // SE
			0,               // S
			0,               // SW
			0,               // W
			0,               // NW
			0,               // Up
			0,               // Down
			0,               // In
			CirculationArea, // Out
		},
	},
	{
		Number:     WeaponShop1,
		BriefMsgNo: text.SolLocation222_Brief,
		FullMsgNo:  text.SolLocation222_Full,
		Flags:      model.LfWeap,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			Corridor20,  // E
			0,           // SE
			WeaponShop2, // S
			0,           // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			Corridor20,  // Out
		},
	},
	{
		Number:     WeaponShop2,
		BriefMsgNo: text.SolLocation223_Brief,
		FullMsgNo:  text.SolLocation223_Full,
		Flags:      model.LfWeap,

		MovTab: [13]uint16{
			WeaponShop1,     // N
			0,               // NE
			WeaponShopAnnex, // E
			0,               // SE
			0,               // S
			0,               // SW
			0,               // W
			0,               // NW
			0,               // Up
			0,               // Down
			0,               // In
			WeaponShop1,     // Out
		},
	},
	{
		Number:     WeaponShopAnnex,
		BriefMsgNo: text.SolLocation224_Brief,
		FullMsgNo:  text.SolLocation224_Full,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			0,           // E
			0,           // SE
			0,           // S
			0,           // SW
			WeaponShop2, // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			WeaponShop2, // Out
		},
	},
	{
		Number:     MarsLandingArea,
		BriefMsgNo: text.SolLocation225_Brief,
		FullMsgNo:  text.SolLocation225_Full,
		Flags:      model.LfLanding,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			Marsport3, // E
			Marsport2, // SE
			Marsport1, // S
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
		Number:     Marsport1,
		BriefMsgNo: text.SolLocation226_Brief,
		FullMsgNo:  text.SolLocation226_Full,

		MovTab: [13]uint16{
			MarsLandingArea, // N
			Marsport3,       // NE
			Marsport2,       // E
			0,               // SE
			0,               // S
			0,               // SW
			0,               // W
			0,               // NW
			0,               // Up
			0,               // Down
			0,               // In
			0,               // Out
		},
	},
	{
		Number:     Marsport2,
		BriefMsgNo: text.SolLocation227_Brief,
		FullMsgNo:  text.SolLocation227_Full,

		MovTab: [13]uint16{
			Marsport3,       // N
			0,               // NE
			0,               // E
			0,               // SE
			0,               // S
			0,               // SW
			Marsport1,       // W
			MarsLandingArea, // NW
			0,               // Up
			0,               // Down
			0,               // In
			0,               // Out
		},
	},
	{
		Number:     Marsport3,
		BriefMsgNo: text.SolLocation228_Brief,
		FullMsgNo:  text.SolLocation228_Full,

		MovTab: [13]uint16{
			SpaceportTerminus, // N
			0,                 // NE
			0,                 // E
			0,                 // SE
			Marsport2,         // S
			Marsport1,         // SW
			MarsLandingArea,   // W
			0,                 // NW
			0,                 // Up
			0,                 // Down
			0,                 // In
			0,                 // Out
		},
	},
	{
		Number:     SpaceportTerminus,
		BriefMsgNo: text.SolLocation229_Brief,
		FullMsgNo:  text.SolLocation229_Full,

		MovTab: [13]uint16{
			MainRoad1,        // N
			0,                // NE
			0,                // E
			0,                // SE
			Marsport3,        // S
			0,                // SW
			SpaceportRepairs, // W
			0,                // NW
			0,                // Up
			0,                // Down
			0,                // In
			MainRoad1,        // Out
		},
	},
	{
		Number:     SpaceportRepairs,
		BriefMsgNo: text.SolLocation230_Brief,
		FullMsgNo:  text.SolLocation230_Full,
		Flags:      model.LfRep,

		MovTab: [13]uint16{
			0,                 // N
			0,                 // NE
			SpaceportTerminus, // E
			0,                 // SE
			0,                 // S
			0,                 // SW
			0,                 // W
			0,                 // NW
			0,                 // Up
			0,                 // Down
			0,                 // In
			SpaceportTerminus, // Out
		},
	},
	{
		Number:     MainRoad1,
		BriefMsgNo: text.SolLocation231_Brief,
		FullMsgNo:  text.SolLocation231_Full,

		MovTab: [13]uint16{
			MainRoad2,         // N
			0,                 // NE
			0,                 // E
			0,                 // SE
			SpaceportTerminus, // S
			0,                 // SW
			0,                 // W
			0,                 // NW
			0,                 // Up
			0,                 // Down
			0,                 // In
			0,                 // Out
		},
	},
	{
		Number:     MainRoad2,
		BriefMsgNo: text.SolLocation232_Brief,
		FullMsgNo:  text.SolLocation232_Full,

		MovTab: [13]uint16{
			CityStreet1, // N
			0,           // NE
			0,           // E
			0,           // SE
			MainRoad1,   // S
			0,           // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     CityStreet1,
		BriefMsgNo: text.SolLocation233_Brief,
		FullMsgNo:  text.SolLocation233_Full,

		MovTab: [13]uint16{
			CityStreet2,  // N
			0,            // NE
			0,            // E
			BeatenTrack1, // SE
			MainRoad2,    // S
			0,            // SW
			Workshop25,   // W
			0,            // NW
			0,            // Up
			0,            // Down
			Workshop25,   // In
			0,            // Out
		},
	},
	{
		Number:     CityStreet2,
		BriefMsgNo: text.SolLocation234_Brief,
		FullMsgNo:  text.SolLocation234_Full,

		MovTab: [13]uint16{
			Crossroads,  // N
			0,           // NE
			0,           // E
			0,           // SE
			CityStreet1, // S
			0,           // SW
			Gunsmiths,   // W
			0,           // NW
			0,           // Up
			0,           // Down
			Gunsmiths,   // In
			0,           // Out
		},
	},
	{
		Number:     CityStreet3,
		BriefMsgNo: text.SolLocation235_Brief,
		FullMsgNo:  text.SolLocation235_Full,

		MovTab: [13]uint16{
			CityStreet4, // N
			0,           // NE
			0,           // E
			0,           // SE
			Crossroads,  // S
			0,           // SW
			TuxDeluxe,   // W
			ChezDiesel,  // NW
			ChezDiesel,  // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     CityStreet4,
		BriefMsgNo: text.SolLocation236_Brief,
		FullMsgNo:  text.SolLocation236_Full,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			HotelLobby,  // E
			0,           // SE
			CityStreet3, // S
			0,           // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			HotelLobby,  // In
			0,           // Out
		},
	},
	{
		Number:     CityStreet5,
		BriefMsgNo: text.SolLocation237_Brief,
		FullMsgNo:  text.SolLocation237_Full,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			Crossroads,  // E
			0,           // SE
			0,           // S
			0,           // SW
			CityStreet6, // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     CityStreet6,
		BriefMsgNo: text.SolLocation238_Brief,
		FullMsgNo:  text.SolLocation238_Full,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			CityStreet5, // E
			0,           // SE
			Pub,         // S
			0,           // SW
			Shop1,       // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     CityStreet7,
		BriefMsgNo: text.SolLocation239_Brief,
		FullMsgNo:  text.SolLocation239_Full,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			CityStreet8, // E
			0,           // SE
			Market3,     // S
			0,           // SW
			Crossroads,  // W
			0,           // NW
			0,           // Up
			0,           // Down
			Market3,     // In
			0,           // Out
		},
	},
	{
		Number:     CityStreet8,
		BriefMsgNo: text.SolLocation240_Brief,
		FullMsgNo:  text.SolLocation240_Full,

		MovTab: [13]uint16{
			0,           // N
			Shop3,       // NE
			0,           // E
			Street1,     // SE
			0,           // S
			0,           // SW
			CityStreet7, // W
			0,           // NW
			0,           // Up
			0,           // Down
			Shop3,       // In
			0,           // Out
		},
	},
	{
		Number:     Street1,
		BriefMsgNo: text.SolLocation241_Brief,
		FullMsgNo:  text.SolLocation241_Full,

		MovTab: [13]uint16{
			0,               // N
			0,               // NE
			LargeWarehouse1, // E
			Street2,         // SE
			0,               // S
			0,               // SW
			0,               // W
			CityStreet8,     // NW
			0,               // Up
			0,               // Down
			LargeWarehouse1, // In
			0,               // Out
		},
	},
	{
		Number:     Street2,
		BriefMsgNo: text.SolLocation242_Brief,
		FullMsgNo:  text.SolLocation242_Full,

		MovTab: [13]uint16{
			0,              // N
			0,              // NE
			0,              // E
			0,              // SE
			SmallWarehouse, // S
			0,              // SW
			0,              // W
			Street1,        // NW
			0,              // Up
			0,              // Down
			SmallWarehouse, // In
			0,              // Out
		},
	},
	{
		Number:     Shop1,
		BriefMsgNo: text.SolLocation243_Brief,
		FullMsgNo:  text.SolLocation243_Full,

		MovTab: [13]uint16{
			Shop2,       // N
			0,           // NE
			CityStreet6, // E
			0,           // SE
			0,           // S
			0,           // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			CityStreet6, // Out
		},
	},
	{
		Number:     Shop2,
		BriefMsgNo: text.SolLocation244_Brief,
		FullMsgNo:  text.SolLocation244_Full,
		Flags:      model.LfGen,

		MovTab: [13]uint16{
			0,     // N
			0,     // NE
			0,     // E
			0,     // SE
			Shop1, // S
			0,     // SW
			0,     // W
			0,     // NW
			0,     // Up
			0,     // Down
			0,     // In
			Shop1, // Out
		},
	},
	{
		Number:     ChezDiesel,
		BriefMsgNo: text.SolLocation245_Brief,
		Events:     [2]uint16{45, 0},
		FullMsgNo:  text.SolLocation245_Full,
		Flags:      model.LfCafe,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			0,           // E
			CityStreet3, // SE
			0,           // S
			0,           // SW
			0,           // W
			0,           // NW
			0,           // Up
			CityStreet3, // Down
			0,           // In
			CityStreet3, // Out
		},
	},
	{
		Number:     TuxDeluxe,
		BriefMsgNo: text.SolLocation246_Brief,
		FullMsgNo:  text.SolLocation246_Full,
		Flags:      model.LfClth,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			CityStreet3, // E
			0,           // SE
			0,           // S
			0,           // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			CityStreet3, // Out
		},
	},
	{
		Number:     HotelLobby,
		BriefMsgNo: text.SolLocation247_Brief,
		FullMsgNo:  text.SolLocation247_Full,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			0,           // E
			0,           // SE
			HotelLounge, // S
			0,           // SW
			CityStreet4, // W
			0,           // NW
			0,           // Up
			0,           // Down
			HotelLounge, // In
			CityStreet4, // Out
		},
	},
	{
		Number:     HotelLounge,
		BriefMsgNo: text.SolLocation248_Brief,
		FullMsgNo:  text.SolLocation248_Full,

		MovTab: [13]uint16{
			HotelLobby, // N
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
			HotelLobby, // Out
		},
	},
	{
		Number:     Shop3,
		BriefMsgNo: text.SolLocation249_Brief,
		FullMsgNo:  text.SolLocation249_Full,
		Flags:      model.LfGen,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			0,           // E
			0,           // SE
			0,           // S
			CityStreet8, // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			CityStreet8, // Out
		},
	},
	{
		Number:     Crossroads,
		BriefMsgNo: text.SolLocation250_Brief,
		FullMsgNo:  text.SolLocation250_Full,

		MovTab: [13]uint16{
			CityStreet3, // N
			0,           // NE
			CityStreet7, // E
			0,           // SE
			CityStreet2, // S
			0,           // SW
			CityStreet5, // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     Pub,
		BriefMsgNo: text.SolLocation251_Brief,
		FullMsgNo:  text.SolLocation251_Full,
		Flags:      model.LfCafe,

		MovTab: [13]uint16{
			CityStreet6, // N
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
			CityStreet6, // Out
		},
	},
	{
		Number:     Gunsmiths,
		BriefMsgNo: text.SolLocation252_Brief,
		FullMsgNo:  text.SolLocation252_Full,
		Flags:      model.LfWeap,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			CityStreet2, // E
			0,           // SE
			0,           // S
			0,           // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			CityStreet2, // Out
		},
	},
	{
		Number:     Market3,
		BriefMsgNo: text.SolLocation253_Brief,
		FullMsgNo:  text.SolLocation253_Full,
		Flags:      model.LfTrade,

		MovTab: [13]uint16{
			CityStreet7, // N
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
			CityStreet7, // Out
		},
	},
	{
		Number:     Workshop25,
		BriefMsgNo: text.SolLocation254_Brief,
		FullMsgNo:  text.SolLocation254_Full,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			CityStreet1, // E
			0,           // SE
			0,           // S
			0,           // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			CityStreet1, // Out
		},
	},
	{
		Number:     LargeWarehouse1,
		BriefMsgNo: text.SolLocation255_Brief,
		FullMsgNo:  text.SolLocation255_Full,

		MovTab: [13]uint16{
			0,               // N
			0,               // NE
			LargeWarehouse2, // E
			0,               // SE
			0,               // S
			0,               // SW
			Street1,         // W
			0,               // NW
			0,               // Up
			0,               // Down
			LargeWarehouse2, // In
			Street1,         // Out
		},
	},
	{
		Number:     LargeWarehouse2,
		BriefMsgNo: text.SolLocation256_Brief,
		FullMsgNo:  text.SolLocation256_Full,

		MovTab: [13]uint16{
			0,               // N
			0,               // NE
			LargeWarehouse3, // E
			0,               // SE
			0,               // S
			0,               // SW
			LargeWarehouse1, // W
			0,               // NW
			0,               // Up
			0,               // Down
			LargeWarehouse3, // In
			LargeWarehouse1, // Out
		},
	},
	{
		Number:     LargeWarehouse3,
		BriefMsgNo: text.SolLocation257_Brief,
		FullMsgNo:  text.SolLocation257_Full,

		MovTab: [13]uint16{
			0,               // N
			0,               // NE
			0,               // E
			0,               // SE
			0,               // S
			0,               // SW
			LargeWarehouse2, // W
			0,               // NW
			0,               // Up
			0,               // Down
			0,               // In
			0,               // Out
		},
	},
	{
		Number:     SmallWarehouse,
		BriefMsgNo: text.SolLocation258_Brief,
		FullMsgNo:  text.SolLocation258_Full,

		MovTab: [13]uint16{
			Street2, // N
			0,       // NE
			0,       // E
			0,       // SE
			0,       // S
			0,       // SW
			0,       // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			Street2, // Out
		},
	},
	{
		Number:     BeatenTrack1,
		BriefMsgNo: text.SolLocation259_Brief,
		FullMsgNo:  text.SolLocation259_Full,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			0,            // E
			BeatenTrack2, // SE
			0,            // S
			0,            // SW
			0,            // W
			CityStreet1,  // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     BeatenTrack2,
		BriefMsgNo: text.SolLocation260_Brief,
		FullMsgNo:  text.SolLocation260_Full,

		MovTab: [13]uint16{
			0,            // N
			BeatenTrack3, // NE
			0,            // E
			0,            // SE
			0,            // S
			0,            // SW
			0,            // W
			BeatenTrack1, // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     BeatenTrack3,
		BriefMsgNo: text.SolLocation261_Brief,
		FullMsgNo:  text.SolLocation261_Full,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			0,            // E
			BeatenTrack4, // SE
			0,            // S
			BeatenTrack2, // SW
			0,            // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     BeatenTrack4,
		BriefMsgNo: text.SolLocation262_Brief,
		FullMsgNo:  text.SolLocation262_Full,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			BeatenTrack5, // E
			0,            // SE
			0,            // S
			0,            // SW
			0,            // W
			BeatenTrack3, // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     BeatenTrack5,
		BriefMsgNo: text.SolLocation263_Brief,
		FullMsgNo:  text.SolLocation263_Full,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			0,            // E
			0,            // SE
			BeatenTrack6, // S
			0,            // SW
			BeatenTrack4, // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     BeatenTrack6,
		BriefMsgNo: text.SolLocation264_Brief,
		FullMsgNo:  text.SolLocation264_Full,

		MovTab: [13]uint16{
			BeatenTrack5, // N
			0,            // NE
			0,            // E
			BeatenTrack7, // SE
			0,            // S
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
		Number:     BeatenTrack7,
		BriefMsgNo: text.SolLocation265_Brief,
		FullMsgNo:  text.SolLocation265_Full,

		MovTab: [13]uint16{
			0,              // N
			0,              // NE
			BeforeTheRuins, // E
			0,              // SE
			0,              // S
			0,              // SW
			0,              // W
			BeatenTrack6,   // NW
			0,              // Up
			0,              // Down
			BeforeTheRuins, // In
			0,              // Out
		},
	},
	{
		Number:     BeforeTheRuins,
		BriefMsgNo: text.SolLocation266_Brief,
		FullMsgNo:  text.SolLocation266_Full,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			Archway1,     // E
			0,            // SE
			0,            // S
			0,            // SW
			BeatenTrack7, // W
			0,            // NW
			0,            // Up
			0,            // Down
			Archway1,     // In
			0,            // Out
		},
	},
	{
		Number:     Archway1,
		BriefMsgNo: text.SolLocation267_Brief,
		FullMsgNo:  text.SolLocation267_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,              // N
			0,              // NE
			Alleyway1,      // E
			0,              // SE
			0,              // S
			0,              // SW
			BeforeTheRuins, // W
			0,              // NW
			0,              // Up
			0,              // Down
			Alleyway1,      // In
			BeforeTheRuins, // Out
		},
	},
	{
		Number:     Alleyway1,
		BriefMsgNo: text.SolLocation268_Brief,
		FullMsgNo:  text.SolLocation268_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			Alleyway2, // E
			0,         // SE
			RedRoom1,  // S
			0,         // SW
			Archway1,  // W
			0,         // NW
			0,         // Up
			0,         // Down
			RedRoom1,  // In
			0,         // Out
		},
	},
	{
		Number:     Alleyway2,
		BriefMsgNo: text.SolLocation269_Brief,
		FullMsgNo:  text.SolLocation269_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			OpenSpace1, // E
			LongRoom3,  // SE
			0,          // S
			0,          // SW
			Alleyway1,  // W
			0,          // NW
			0,          // Up
			0,          // Down
			LongRoom3,  // In
			0,          // Out
		},
	},
	{
		Number:     Alleyway3,
		BriefMsgNo: text.SolLocation270_Brief,
		FullMsgNo:  text.SolLocation270_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,                  // N
			MalodorousAlleyway, // NE
			0,                  // E
			0,                  // SE
			OpenSpace1,         // S
			0,                  // SW
			NarrowSlit1,        // W
			0,                  // NW
			0,                  // Up
			0,                  // Down
			NarrowSlit1,        // In
			0,                  // Out
		},
	},
	{
		Number:     MalodorousAlleyway,
		BriefMsgNo: text.SolLocation271_Brief,
		FullMsgNo:  text.SolLocation271_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			Alleyway4, // E
			0,         // SE
			0,         // S
			Alleyway3, // SW
			0,         // W
			MaresNest, // NW
			0,         // Up
			0,         // Down
			MaresNest, // In
			0,         // Out
		},
	},
	{
		Number:     Alleyway4,
		BriefMsgNo: text.SolLocation272_Brief,
		FullMsgNo:  text.SolLocation272_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			Alleyway5,          // N
			0,                  // NE
			0,                  // E
			0,                  // SE
			HazyMist1,          // S
			LongRoom1,          // SW
			MalodorousAlleyway, // W
			0,                  // NW
			0,                  // Up
			0,                  // Down
			0,                  // In
			0,                  // Out
		},
	},
	{
		Number:     Alleyway5,
		BriefMsgNo: text.SolLocation273_Brief,
		FullMsgNo:  text.SolLocation273_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			Rubble1,   // N
			0,         // NE
			Alleyway6, // E
			0,         // SE
			Alleyway4, // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			Rubble1,   // In
			0,         // Out
		},
	},
	{
		Number:     Alleyway6,
		BriefMsgNo: text.SolLocation274_Brief,
		FullMsgNo:  text.SolLocation274_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,         // N
			Alleyway7, // NE
			0,         // E
			0,         // SE
			0,         // S
			0,         // SW
			Alleyway5, // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     Alleyway7,
		BriefMsgNo: text.SolLocation275_Brief,
		FullMsgNo:  text.SolLocation275_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			Alleyway8, // SE
			0,         // S
			Alleyway6, // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     Alleyway8,
		BriefMsgNo: text.SolLocation276_Brief,
		FullMsgNo:  text.SolLocation276_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Alleyway9,  // E
			0,          // SE
			EmptyRoom2, // S
			BareRoom1,  // SW
			0,          // W
			Alleyway7,  // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     RedRoom1,
		BriefMsgNo: text.SolLocation277_Brief,
		FullMsgNo:  text.SolLocation277_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			Alleyway1, // N
			0,         // NE
			RedRoom2,  // E
			0,         // SE
			RedRoom3,  // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			Alleyway1, // Out
		},
	},
	{
		Number:     RedRoom2,
		BriefMsgNo: text.SolLocation278_Brief,
		FullMsgNo:  text.SolLocation278_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			0,        // E
			0,        // SE
			RedRoom4, // S
			0,        // SW
			RedRoom1, // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     RedRoom3,
		BriefMsgNo: text.SolLocation279_Brief,
		FullMsgNo:  text.SolLocation279_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			RedRoom1, // N
			0,        // NE
			RedRoom4, // E
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
		Number:     RedRoom4,
		BriefMsgNo: text.SolLocation280_Brief,
		FullMsgNo:  text.SolLocation280_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			RedRoom2,  // N
			0,         // NE
			HazyMist3, // E
			0,         // SE
			0,         // S
			0,         // SW
			RedRoom3,  // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     HazyMist1,
		BriefMsgNo: text.SolLocation281_Brief,
		FullMsgNo:  text.SolLocation281_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			Alleyway4, // N
			0,         // NE
			0,         // E
			0,         // SE
			HazyMist2, // S
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
		Number:     HazyMist2,
		BriefMsgNo: text.SolLocation282_Brief,
		FullMsgNo:  text.SolLocation282_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			HazyMist1, // N
			0,         // NE
			0,         // E
			0,         // SE
			0,         // S
			LinesRoom, // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     HazyMist3,
		BriefMsgNo: text.SolLocation283_Brief,
		FullMsgNo:  text.SolLocation283_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,         // N
			LinesRoom, // NE
			0,         // E
			0,         // SE
			0,         // S
			0,         // SW
			RedRoom4,  // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     LinesRoom,
		BriefMsgNo: text.SolLocation284_Brief,
		Events:     [2]uint16{0, 17},
		FullMsgNo:  text.SolLocation284_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,         // N
			HazyMist2, // NE
			0,         // E
			SlabRoom,  // SE
			0,         // S
			HazyMist3, // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     SlabRoom,
		BriefMsgNo: text.SolLocation285_Brief,
		FullMsgNo:  text.SolLocation285_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			0,         // SE
			0,         // S
			0,         // SW
			0,         // W
			LinesRoom, // NW
			0,         // Up
			0,         // Down
			0,         // In
			LinesRoom, // Out
		},
	},
	{
		Number:     OpenSpace1,
		BriefMsgNo: text.SolLocation286_Brief,
		FullMsgNo:  text.SolLocation286_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			Alleyway3, // N
			0,         // NE
			0,         // E
			0,         // SE
			0,         // S
			0,         // SW
			Alleyway2, // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     LongRoom1,
		BriefMsgNo: text.SolLocation287_Brief,
		FullMsgNo:  text.SolLocation287_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,         // N
			Alleyway4, // NE
			0,         // E
			0,         // SE
			LongRoom2, // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			Alleyway4, // Out
		},
	},
	{
		Number:     LongRoom2,
		BriefMsgNo: text.SolLocation288_Brief,
		FullMsgNo:  text.SolLocation288_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			LongRoom1, // N
			0,         // NE
			0,         // E
			0,         // SE
			0,         // S
			LongRoom3, // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     LongRoom3,
		BriefMsgNo: text.SolLocation289_Brief,
		FullMsgNo:  text.SolLocation289_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,         // N
			LongRoom2, // NE
			0,         // E
			0,         // SE
			0,         // S
			0,         // SW
			0,         // W
			Alleyway2, // NW
			0,         // Up
			0,         // Down
			0,         // In
			Alleyway2, // Out
		},
	},
	{
		Number:     BareRoom1,
		BriefMsgNo: text.SolLocation290_Brief,
		FullMsgNo:  text.SolLocation290_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,         // N
			Alleyway8, // NE
			0,         // E
			0,         // SE
			0,         // S
			BareRoom2, // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			Alleyway8, // Out
		},
	},
	{
		Number:     BareRoom2,
		BriefMsgNo: text.SolLocation291_Brief,
		FullMsgNo:  text.SolLocation291_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,          // N
			BareRoom1,  // NE
			EmptyRoom1, // E
			0,          // SE
			0,          // S
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
		Number:     EmptyRoom1,
		BriefMsgNo: text.SolLocation292_Brief,
		FullMsgNo:  text.SolLocation292_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,          // N
			EmptyRoom2, // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			BareRoom2,  // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     EmptyRoom2,
		BriefMsgNo: text.SolLocation293_Brief,
		FullMsgNo:  text.SolLocation293_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			Alleyway8,  // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			EmptyRoom1, // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			Alleyway8,  // Out
		},
	},
	{
		Number:     NarrowSlit1,
		BriefMsgNo: text.SolLocation294_Brief,
		FullMsgNo:  text.SolLocation294_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			Alleyway3,   // E
			0,           // SE
			0,           // S
			0,           // SW
			0,           // W
			NarrowSlit2, // NW
			0,           // Up
			0,           // Down
			0,           // In
			Alleyway3,   // Out
		},
	},
	{
		Number:     NarrowSlit2,
		BriefMsgNo: text.SolLocation295_Brief,
		FullMsgNo:  text.SolLocation295_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,           // N
			NarrowSlit3, // NE
			0,           // E
			NarrowSlit1, // SE
			0,           // S
			0,           // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     NarrowSlit3,
		BriefMsgNo: text.SolLocation296_Brief,
		FullMsgNo:  text.SolLocation296_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			Courtyard1,  // N
			0,           // NE
			0,           // E
			0,           // SE
			0,           // S
			NarrowSlit2, // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     NarrowSlit4,
		BriefMsgNo: text.SolLocation297_Brief,
		FullMsgNo:  text.SolLocation297_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			0,       // E
			0,       // SE
			0,       // S
			0,       // SW
			Rubble3, // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Courtyard1,
		BriefMsgNo: text.SolLocation298_Brief,
		FullMsgNo:  text.SolLocation298_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			Courtyard2,  // E
			0,           // SE
			NarrowSlit3, // S
			0,           // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     Courtyard2,
		BriefMsgNo: text.SolLocation299_Brief,
		FullMsgNo:  text.SolLocation299_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			Courtyard1, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			Courtyard1, // Out
		},
	},
	{
		Number:     TreasureRoom1,
		BriefMsgNo: text.SolLocation300_Brief,
		FullMsgNo:  text.SolLocation300_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			TreasureRoom2, // E
			0,             // SE
			0,             // S
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
		Number:     TreasureRoom2,
		BriefMsgNo: text.SolLocation301_Brief,
		FullMsgNo:  text.SolLocation301_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			0,             // E
			0,             // SE
			0,             // S
			0,             // SW
			TreasureRoom1, // W
			0,             // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     PentagonalRoom,
		BriefMsgNo: text.SolLocation302_Brief,
		Events:     [2]uint16{0, 11},
		FullMsgNo:  text.SolLocation302_Full,
		SysLoc:     text.NO_MOVEMENT_12,
	},
	{
		Number:     MaresNest,
		BriefMsgNo: text.SolLocation303_Brief,
		FullMsgNo:  text.SolLocation303_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,                  // N
			0,                  // NE
			0,                  // E
			MalodorousAlleyway, // SE
			0,                  // S
			0,                  // SW
			0,                  // W
			0,                  // NW
			0,                  // Up
			0,                  // Down
			0,                  // In
			MalodorousAlleyway, // Out
		},
	},
	{
		Number:     Rubble1,
		BriefMsgNo: text.SolLocation304_Brief,
		FullMsgNo:  text.SolLocation304_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			Rubble2,   // N
			0,         // NE
			0,         // E
			0,         // SE
			Alleyway5, // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			Alleyway5, // Out
		},
	},
	{
		Number:     Rubble2,
		BriefMsgNo: text.SolLocation305_Brief,
		FullMsgNo:  text.SolLocation305_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			0,       // E
			0,       // SE
			Rubble1, // S
			0,       // SW
			0,       // W
			CaveIn,  // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     CaveIn,
		BriefMsgNo: text.SolLocation306_Brief,
		FullMsgNo:  text.SolLocation306_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			0,           // E
			Rubble2,     // SE
			0,           // S
			0,           // SW
			NarrowSlit4, // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     Alleyway9,
		BriefMsgNo: text.SolLocation307_Brief,
		FullMsgNo:  text.SolLocation307_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			Alleyway9,        // N
			MazeOfAlleyways2, // NE
			Alleyway9,        // E
			Alleyway9,        // SE
			Alleyway9,        // S
			Alleyway9,        // SW
			Alleyway8,        // W
			Alleyway9,        // NW
			0,                // Up
			0,                // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     Passage1,
		BriefMsgNo: text.SolLocation308_Brief,
		FullMsgNo:  text.SolLocation308_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			Passage2, // E
			0,        // SE
			0,        // S
			0,        // SW
			0,        // W
			BlueRoom, // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     Passage2,
		BriefMsgNo: text.SolLocation309_Brief,
		FullMsgNo:  text.SolLocation309_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			Rubble3,  // E
			0,        // SE
			0,        // S
			0,        // SW
			Passage1, // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     Rubble3,
		BriefMsgNo: text.SolLocation310_Brief,
		FullMsgNo:  text.SolLocation310_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			0,        // E
			0,        // SE
			0,        // S
			0,        // SW
			Passage2, // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     Passage3,
		BriefMsgNo: text.SolLocation311_Brief,
		FullMsgNo:  text.SolLocation311_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,        // N
			Passage4, // NE
			0,        // E
			0,        // SE
			0,        // S
			0,        // SW
			0,        // W
			DeepPit,  // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     Passage4,
		BriefMsgNo: text.SolLocation312_Brief,
		FullMsgNo:  text.SolLocation312_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			PitRoom,  // N
			0,        // NE
			0,        // E
			0,        // SE
			0,        // S
			Passage3, // SW
			0,        // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     Passage5,
		BriefMsgNo: text.SolLocation313_Brief,
		FullMsgNo:  text.SolLocation313_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			NarrowPassage, // E
			0,             // SE
			0,             // S
			0,             // SW
			PitRoom,       // W
			0,             // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     Passage6,
		BriefMsgNo: text.SolLocation314_Brief,
		FullMsgNo:  text.SolLocation314_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			Passage7, // E
			0,        // SE
			MapRoom,  // S
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
		Number:     Passage7,
		BriefMsgNo: text.SolLocation315_Brief,
		FullMsgNo:  text.SolLocation315_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			Passage8, // E
			0,        // SE
			0,        // S
			0,        // SW
			Passage6, // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     Passage8,
		BriefMsgNo: text.SolLocation316_Brief,
		FullMsgNo:  text.SolLocation316_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			Passage9,    // E
			0,           // SE
			MachineRoom, // S
			0,           // SW
			Passage7,    // W
			0,           // NW
			0,           // Up
			0,           // Down
			MachineRoom, // In
			0,           // Out
		},
	},
	{
		Number:     Passage9,
		BriefMsgNo: text.SolLocation317_Brief,
		FullMsgNo:  text.SolLocation317_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			0,         // SE
			Passage10, // S
			0,         // SW
			Passage8,  // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     Passage10,
		BriefMsgNo: text.SolLocation318_Brief,
		FullMsgNo:  text.SolLocation318_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			Passage9,  // N
			0,         // NE
			0,         // E
			0,         // SE
			Passage12, // S
			Passage11, // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     Passage11,
		BriefMsgNo: text.SolLocation319_Brief,
		FullMsgNo:  text.SolLocation319_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,          // N
			Passage10,  // NE
			0,          // E
			0,          // SE
			0,          // S
			OpenSpace2, // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			Passage10,  // In
			OpenSpace2, // Out
		},
	},
	{
		Number:     Passage12,
		BriefMsgNo: text.SolLocation320_Brief,
		FullMsgNo:  text.SolLocation320_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			Passage10, // N
			0,         // NE
			0,         // E
			0,         // SE
			Passage13, // S
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
		Number:     Passage13,
		BriefMsgNo: text.SolLocation321_Brief,
		FullMsgNo:  text.SolLocation321_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			Passage12, // N
			0,         // NE
			0,         // E
			0,         // SE
			0,         // S
			0,         // SW
			Passage14, // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     Passage14,
		BriefMsgNo: text.SolLocation322_Brief,
		FullMsgNo:  text.SolLocation322_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Passage13,  // E
			0,          // SE
			OpenSpace3, // S
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
		Number:     BlueRoom,
		BriefMsgNo: text.SolLocation323_Brief,
		FullMsgNo:  text.SolLocation323_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,            // N
			PillaredRoom, // NE
			0,            // E
			Passage1,     // SE
			0,            // S
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
		Number:     PillaredRoom,
		BriefMsgNo: text.SolLocation324_Brief,
		FullMsgNo:  text.SolLocation324_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			DeepPit,  // N
			0,        // NE
			0,        // E
			0,        // SE
			0,        // S
			BlueRoom, // SW
			0,        // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     DeepPit,
		BriefMsgNo: text.SolLocation325_Brief,
		FullMsgNo:  text.SolLocation325_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			Pit,          // N
			Pit,          // NE
			Pit,          // E
			Passage3,     // SE
			PillaredRoom, // S
			Pit,          // SW
			Pit,          // W
			Pit,          // NW
			Pit,          // Up
			Pit,          // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     PitRoom,
		BriefMsgNo: text.SolLocation326_Brief,
		FullMsgNo:  text.SolLocation326_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,              // N
			0,              // NE
			Passage5,       // E
			0,              // SE
			Passage4,       // S
			0,              // SW
			0,              // W
			0,              // NW
			0,              // Up
			NuclearReactor, // Down
			NuclearReactor, // In
			0,              // Out
		},
	},
	{
		Number:     NuclearReactor,
		BriefMsgNo: text.SolLocation327_Brief,
		FullMsgNo:  text.SolLocation327_Full,
		Flags:      model.LfDeath,
	},
	{
		Number:     NarrowPassage,
		BriefMsgNo: text.SolLocation328_Brief,
		FullMsgNo:  text.SolLocation328_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			0,        // E
			MapRoom,  // SE
			0,        // S
			0,        // SW
			Passage5, // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     MapRoom,
		BriefMsgNo: text.SolLocation329_Brief,
		FullMsgNo:  text.SolLocation329_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			Passage6,      // N
			0,             // NE
			0,             // E
			0,             // SE
			0,             // S
			0,             // SW
			0,             // W
			NarrowPassage, // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     MachineRoom,
		BriefMsgNo: text.SolLocation330_Brief,
		FullMsgNo:  text.SolLocation330_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			Passage8,   // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			ShelfRoom1, // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     ShelfRoom1,
		BriefMsgNo: text.SolLocation331_Brief,
		FullMsgNo:  text.SolLocation331_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			ShelfRoom2,  // N
			MachineRoom, // NE
			0,           // E
			0,           // SE
			0,           // S
			0,           // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     ShelfRoom2,
		BriefMsgNo: text.SolLocation332_Brief,
		FullMsgNo:  text.SolLocation332_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			0,           // E
			0,           // SE
			ShelfRoom1,  // S
			Antechamber, // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     Antechamber,
		BriefMsgNo: text.SolLocation333_Brief,
		FullMsgNo:  text.SolLocation333_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,          // N
			ShelfRoom2, // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			0,          // W
			Dormitory7, // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Dormitory6,
		BriefMsgNo: text.SolLocation334_Brief,
		FullMsgNo:  text.SolLocation334_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Dormitory7, // E
			0,          // SE
			Dormitory8, // S
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
		Number:     Dormitory7,
		BriefMsgNo: text.SolLocation335_Brief,
		FullMsgNo:  text.SolLocation335_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			0,           // E
			Antechamber, // SE
			Dormitory9,  // S
			0,           // SW
			Dormitory6,  // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     Dormitory8,
		BriefMsgNo: text.SolLocation336_Brief,
		FullMsgNo:  text.SolLocation336_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			Dormitory6, // N
			0,          // NE
			Dormitory9, // E
			0,          // SE
			0,          // S
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
		Number:     Dormitory9,
		BriefMsgNo: text.SolLocation337_Brief,
		FullMsgNo:  text.SolLocation337_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			Dormitory7, // N
			0,          // NE
			0,          // E
			Guardroom6, // SE
			0,          // S
			0,          // SW
			Dormitory8, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Guardroom6,
		BriefMsgNo: text.SolLocation338_Brief,
		FullMsgNo:  text.SolLocation338_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			OpenSpace2, // E
			0,          // SE
			0,          // S
			0,          // SW
			0,          // W
			Dormitory9, // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     OpenSpace2,
		BriefMsgNo: text.SolLocation339_Brief,
		FullMsgNo:  text.SolLocation339_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,          // N
			Passage11,  // NE
			0,          // E
			OpenSpace3, // SE
			OpenSpace4, // S
			0,          // SW
			Guardroom6, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     OpenSpace3,
		BriefMsgNo: text.SolLocation340_Brief,
		FullMsgNo:  text.SolLocation340_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			Passage14,      // N
			0,              // NE
			PentagonalRoom, // E
			0,              // SE
			Rubble5,        // S
			0,              // SW
			OpenSpace4,     // W
			OpenSpace2,     // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     OpenSpace4,
		BriefMsgNo: text.SolLocation341_Brief,
		FullMsgNo:  text.SolLocation341_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			OpenSpace2, // N
			0,          // NE
			OpenSpace3, // E
			0,          // SE
			Rubble4,    // S
			0,          // SW
			Archway2,   // W
			0,          // NW
			0,          // Up
			0,          // Down
			Archway2,   // In
			0,          // Out
		},
	},
	{
		Number:     Rubble4,
		BriefMsgNo: text.SolLocation342_Brief,
		FullMsgNo:  text.SolLocation342_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			OpenSpace4, // N
			0,          // NE
			Rubble5,    // E
			0,          // SE
			0,          // S
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
		Number:     Rubble5,
		BriefMsgNo: text.SolLocation343_Brief,
		FullMsgNo:  text.SolLocation343_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			OpenSpace3, // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			Rubble4,    // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Archway2,
		BriefMsgNo: text.SolLocation344_Brief,
		FullMsgNo:  text.SolLocation344_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			OpenSpace4, // E
			0,          // SE
			0,          // S
			0,          // SW
			0,          // W
			Temple5,    // NW
			0,          // Up
			0,          // Down
			Temple5,    // In
			OpenSpace4, // Out
		},
	},
	{
		Number:     Temple1,
		BriefMsgNo: text.SolLocation345_Brief,
		FullMsgNo:  text.SolLocation345_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			Temple2, // E
			0,       // SE
			0,       // S
			0,       // SW
			0,       // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Temple2,
		BriefMsgNo: text.SolLocation346_Brief,
		FullMsgNo:  text.SolLocation346_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,       // N
			Altar,   // NE
			Temple3, // E
			0,       // SE
			0,       // S
			0,       // SW
			Temple1, // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Temple3,
		BriefMsgNo: text.SolLocation347_Brief,
		FullMsgNo:  text.SolLocation347_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			Altar,   // N
			0,       // NE
			Temple4, // E
			0,       // SE
			0,       // S
			0,       // SW
			Temple2, // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Temple4,
		BriefMsgNo: text.SolLocation348_Brief,
		FullMsgNo:  text.SolLocation348_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			Temple5, // E
			0,       // SE
			0,       // S
			0,       // SW
			Temple3, // W
			Altar,   // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Temple5,
		BriefMsgNo: text.SolLocation349_Brief,
		FullMsgNo:  text.SolLocation349_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			0,        // E
			Archway2, // SE
			0,        // S
			0,        // SW
			Temple4,  // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			Archway2, // Out
		},
	},
	{
		Number:     Altar,
		BriefMsgNo: text.SolLocation350_Brief,
		FullMsgNo:  text.SolLocation350_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			Temple4,   // SE
			Temple3,   // S
			Temple2,   // SW
			0,         // W
			0,         // NW
			0,         // Up
			Sacrifice, // Down
			Sacrifice, // In
			0,         // Out
		},
	},
	{
		Number:     Sacrifice,
		BriefMsgNo: text.SolLocation351_Brief,
		FullMsgNo:  text.SolLocation351_Full,
		Flags:      model.LfDeath,
	},
	{
		Number:     Pit,
		BriefMsgNo: text.SolLocation352_Brief,
		FullMsgNo:  text.SolLocation352_Full,
		Flags:      model.LfDeath,
	},
	{
		Number:     MazeOfAlleys1,
		BriefMsgNo: text.SolLocation353_Brief,
		FullMsgNo:  text.SolLocation353_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			MazeOfAlleys1, // N
			MazeOfAlleys1, // NE
			MazeOfAlleys1, // E
			MazeOfAlleys1, // SE
			MazeOfAlleys1, // S
			MazeOfAlleys2, // SW
			MazeOfAlleys1, // W
			MazeOfAlleys1, // NW
			0,             // Up
			0,             // Down
			MazeOfAlleys1, // In
			MazeOfAlleys1, // Out
		},
	},
	{
		Number:     MazeOfAlleys2,
		BriefMsgNo: text.SolLocation354_Brief,
		FullMsgNo:  text.SolLocation354_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			MazeOfAlleys2,    // N
			MazeOfAlleys1,    // NE
			MazeOfAlleyways1, // E
			MazeOfAlleys2,    // SE
			MazeOfAlleyways2, // S
			MazeOfAlleys2,    // SW
			MazeOfAlleys2,    // W
			MazeOfAlleys2,    // NW
			0,                // Up
			0,                // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     MazeOfAlleyways1,
		BriefMsgNo: text.SolLocation355_Brief,
		FullMsgNo:  text.SolLocation355_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			MazeOfAlleyways1, // N
			MazeOfAlleyways1, // NE
			MazeOfAlleyways1, // E
			MazeOfAlleyways1, // SE
			MazeOfAlleyways1, // S
			MazeOfAlleyways1, // SW
			MazeOfAlleys2,    // W
			MazeOfAlleyways1, // NW
			0,                // Up
			0,                // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     MazeOfAlleyways2,
		BriefMsgNo: text.SolLocation356_Brief,
		FullMsgNo:  text.SolLocation356_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			MazeOfAlleys2,    // N
			MazeOfAlleyways2, // NE
			MazeOfAlleyways2, // E
			MazeOfAlleyways3, // SE
			MazeOfAlleyways2, // S
			Alleyway9,        // SW
			MazeOfAlleyways2, // W
			MazeOfAlleyways2, // NW
			0,                // Up
			0,                // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     MazeOfAlleys3,
		BriefMsgNo: text.SolLocation357_Brief,
		FullMsgNo:  text.SolLocation357_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			MazeOfAlleys3,    // N
			MazeOfAlleys3,    // NE
			MazeOfAlleys3,    // E
			MazeOfAlleys3,    // SE
			MazeOfAlleyways3, // S
			MazeOfAlleys3,    // SW
			MazeOfAlleys3,    // W
			MazeOfAlleys3,    // NW
			0,                // Up
			0,                // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     MazeOfAlleyways3,
		BriefMsgNo: text.SolLocation358_Brief,
		FullMsgNo:  text.SolLocation358_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			MazeOfAlleys3,    // N
			MazeOfAlleyways3, // NE
			MazeOfAlleys4,    // E
			MazeOfAlleyways3, // SE
			MazeOfAlleyways3, // S
			MazeOfAlleyways3, // SW
			MazeOfAlleyways3, // W
			MazeOfAlleyways2, // NW
			0,                // Up
			0,                // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     MazeOfAlleys4,
		BriefMsgNo: text.SolLocation359_Brief,
		FullMsgNo:  text.SolLocation359_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			MazeOfAlleys4,    // N
			MazeOfAlleys4,    // NE
			MazeOfAlleys4,    // E
			MazeOfAlleys7,    // SE
			MazeOfAlleys4,    // S
			MazeOfAlleys4,    // SW
			MazeOfAlleyways3, // W
			MazeOfAlleys4,    // NW
			0,                // Up
			0,                // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     MazeOfAlleys5,
		BriefMsgNo: text.SolLocation360_Brief,
		FullMsgNo:  text.SolLocation360_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			MazeOfAlleys5, // N
			MazeOfAlleys5, // NE
			MazeOfAlleys6, // E
			MazeOfAlleys5, // SE
			MazeOfAlleys5, // S
			MazeOfAlleys7, // SW
			MazeOfAlleys5, // W
			MazeOfAlleys5, // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     MazeOfAlleys6,
		BriefMsgNo: text.SolLocation361_Brief,
		FullMsgNo:  text.SolLocation361_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			MazeOfAlleys6, // N
			MazeOfAlleys6, // NE
			MazeOfAlleys6, // E
			MazeOfAlleys8, // SE
			MazeOfAlleys6, // S
			MazeOfAlleys6, // SW
			MazeOfAlleys5, // W
			MazeOfAlleys6, // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     MazeOfAlleys7,
		BriefMsgNo: text.SolLocation362_Brief,
		FullMsgNo:  text.SolLocation362_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			MazeOfAlleys7,  // N
			MazeOfAlleys5,  // NE
			MazeOfAlleys7,  // E
			MazeOfAlleys10, // SE
			MazeOfAlleys7,  // S
			MazeOfAlleys9,  // SW
			MazeOfAlleys7,  // W
			MazeOfAlleys4,  // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     MazeOfAlleys8,
		BriefMsgNo: text.SolLocation363_Brief,
		FullMsgNo:  text.SolLocation363_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			MazeOfAlleys8,  // N
			MazeOfAlleys8,  // NE
			MazeOfAlleys8,  // E
			MazeOfAlleys8,  // SE
			MazeOfAlleys8,  // S
			MazeOfAlleys11, // SW
			MazeOfAlleys8,  // W
			MazeOfAlleys6,  // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     MazeOfAlleys9,
		BriefMsgNo: text.SolLocation364_Brief,
		FullMsgNo:  text.SolLocation364_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			MazeOfAlleys9,  // N
			MazeOfAlleys7,  // NE
			MazeOfAlleys9,  // E
			MazeOfAlleys9,  // SE
			MazeOfAlleys12, // S
			MazeOfAlleys9,  // SW
			MazeOfAlleys9,  // W
			MazeOfAlleys9,  // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     MazeOfAlleys10,
		BriefMsgNo: text.SolLocation365_Brief,
		FullMsgNo:  text.SolLocation365_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			MazeOfAlleys10, // N
			MazeOfAlleys10, // NE
			MazeOfAlleys10, // E
			MazeOfAlleys10, // SE
			MazeOfAlleys10, // S
			MazeOfAlleys10, // SW
			MazeOfAlleys10, // W
			MazeOfAlleys7,  // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     MazeOfAlleys11,
		BriefMsgNo: text.SolLocation366_Brief,
		FullMsgNo:  text.SolLocation366_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			MazeOfAlleys11, // N
			MazeOfAlleys8,  // NE
			MazeOfAlleys11, // E
			MazeOfAlleys11, // SE
			MazeOfAlleys11, // S
			MazeOfAlleys11, // SW
			MazeOfAlleys11, // W
			MazeOfAlleys11, // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     MazeOfAlleys12,
		BriefMsgNo: text.SolLocation367_Brief,
		FullMsgNo:  text.SolLocation367_Full,
		SysLoc:     text.NO_MOVEMENT_12,

		MovTab: [13]uint16{
			MazeOfAlleys9,  // N
			MazeOfAlleys12, // NE
			MazeOfAlleys12, // E
			MazeOfAlleys12, // SE
			MazeOfAlleys12, // S
			MazeOfAlleys12, // SW
			MazeOfAlleys12, // W
			MazeOfAlleys12, // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     Foyer,
		BriefMsgNo: text.SolLocation368_Brief,
		FullMsgNo:  text.SolLocation368_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Corridor30, // E
			0,          // SE
			Corridor23, // S
			0,          // SW
			0,          // W
			Road8,      // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Corridor23,
		BriefMsgNo: text.SolLocation369_Brief,
		FullMsgNo:  text.SolLocation369_Full,

		MovTab: [13]uint16{
			Foyer,      // N
			0,          // NE
			0,          // E
			0,          // SE
			Corridor24, // S
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
		Number:     Corridor24,
		BriefMsgNo: text.SolLocation370_Brief,
		FullMsgNo:  text.SolLocation370_Full,

		MovTab: [13]uint16{
			Corridor23, // N
			0,          // NE
			Corridor25, // E
			0,          // SE
			0,          // S
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
		Number:     Corridor25,
		BriefMsgNo: text.SolLocation371_Brief,
		FullMsgNo:  text.SolLocation371_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Corridor26, // E
			0,          // SE
			0,          // S
			0,          // SW
			Corridor24, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Corridor26,
		BriefMsgNo: text.SolLocation372_Brief,
		FullMsgNo:  text.SolLocation372_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Corridor27, // E
			0,          // SE
			Office4,    // S
			0,          // SW
			Corridor25, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Corridor27,
		BriefMsgNo: text.SolLocation373_Brief,
		FullMsgNo:  text.SolLocation373_Full,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			Corridor28,    // E
			0,             // SE
			HospitalWard4, // S
			0,             // SW
			Corridor26,    // W
			0,             // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     Corridor28,
		BriefMsgNo: text.SolLocation374_Brief,
		FullMsgNo:  text.SolLocation374_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Corridor29, // E
			0,          // SE
			0,          // S
			0,          // SW
			Corridor27, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Corridor29,
		BriefMsgNo: text.SolLocation375_Brief,
		FullMsgNo:  text.SolLocation375_Full,

		MovTab: [13]uint16{
			WaitingRoom,  // N
			0,            // NE
			CasualtyWard, // E
			0,            // SE
			0,            // S
			0,            // SW
			Corridor28,   // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     Corridor30,
		BriefMsgNo: text.SolLocation376_Brief,
		FullMsgNo:  text.SolLocation376_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Corridor31, // E
			0,          // SE
			0,          // S
			0,          // SW
			Foyer,      // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Corridor31,
		BriefMsgNo: text.SolLocation377_Brief,
		FullMsgNo:  text.SolLocation377_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Corridor32, // E
			0,          // SE
			0,          // S
			0,          // SW
			Corridor30, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Corridor32,
		BriefMsgNo: text.SolLocation378_Brief,
		FullMsgNo:  text.SolLocation378_Full,

		MovTab: [13]uint16{
			HospitalWard1,    // N
			0,                // NE
			Corridor33,       // E
			0,                // SE
			OperatingTheater, // S
			0,                // SW
			Corridor31,       // W
			0,                // NW
			0,                // Up
			0,                // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     Corridor33,
		BriefMsgNo: text.SolLocation379_Brief,
		FullMsgNo:  text.SolLocation379_Full,

		MovTab: [13]uint16{
			Office3,    // N
			0,          // NE
			Corridor34, // E
			0,          // SE
			0,          // S
			0,          // SW
			Corridor32, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Corridor34,
		BriefMsgNo: text.SolLocation380_Brief,
		FullMsgNo:  text.SolLocation380_Full,

		MovTab: [13]uint16{
			0,          // N
			Corridor35, // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			Corridor33, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Corridor35,
		BriefMsgNo: text.SolLocation381_Brief,
		FullMsgNo:  text.SolLocation381_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			Corridor36, // SE
			0,          // S
			Corridor34, // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Corridor36,
		BriefMsgNo: text.SolLocation382_Brief,
		FullMsgNo:  text.SolLocation382_Full,

		MovTab: [13]uint16{
			0,             // N
			HospitalWard2, // NE
			0,             // E
			Mortuary,      // SE
			0,             // S
			HospitalWard3, // SW
			0,             // W
			Corridor35,    // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     HospitalWard1,
		BriefMsgNo: text.SolLocation383_Brief,
		FullMsgNo:  text.SolLocation383_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			Corridor32, // S
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
		Number:     HospitalWard2,
		BriefMsgNo: text.SolLocation384_Brief,
		FullMsgNo:  text.SolLocation384_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			Corridor36, // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			Corridor36, // Out
		},
	},
	{
		Number:     HospitalWard3,
		BriefMsgNo: text.SolLocation385_Brief,
		FullMsgNo:  text.SolLocation385_Full,

		MovTab: [13]uint16{
			0,          // N
			Corridor36, // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			Corridor36, // Out
		},
	},
	{
		Number:     CasualtyWard,
		BriefMsgNo: text.SolLocation386_Brief,
		FullMsgNo:  text.SolLocation386_Full,
		Flags:      model.LfHospital,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			Corridor29, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     HospitalWard4,
		BriefMsgNo: text.SolLocation387_Brief,
		FullMsgNo:  text.SolLocation387_Full,

		MovTab: [13]uint16{
			Corridor27, // N
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
			Corridor27, // Out
		},
	},
	{
		Number:     Office3,
		BriefMsgNo: text.SolLocation388_Brief,
		FullMsgNo:  text.SolLocation388_Full,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			0,            // E
			0,            // SE
			Corridor33,   // S
			0,            // SW
			0,            // W
			PathologyLab, // NW
			0,            // Up
			0,            // Down
			PathologyLab, // In
			Corridor33,   // Out
		},
	},
	{
		Number:     Office4,
		BriefMsgNo: text.SolLocation389_Brief,
		FullMsgNo:  text.SolLocation389_Full,

		MovTab: [13]uint16{
			Corridor26, // N
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
			Corridor26, // Out
		},
	},
	{
		Number:     PathologyLab,
		BriefMsgNo: text.SolLocation390_Brief,
		FullMsgNo:  text.SolLocation390_Full,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			0,       // E
			Office3, // SE
			0,       // S
			0,       // SW
			0,       // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			Office3, // Out
		},
	},
	{
		Number:     ScannerRoom,
		BriefMsgNo: text.SolLocation391_Brief,
		FullMsgNo:  text.SolLocation391_Full,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			0,           // E
			0,           // SE
			0,           // S
			WaitingRoom, // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			WaitingRoom, // Out
		},
	},
	{
		Number:     OperatingTheater,
		BriefMsgNo: text.SolLocation392_Brief,
		FullMsgNo:  text.SolLocation392_Full,

		MovTab: [13]uint16{
			Corridor32, // N
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
			Corridor32, // Out
		},
	},
	{
		Number:     WaitingRoom,
		BriefMsgNo: text.SolLocation393_Brief,
		FullMsgNo:  text.SolLocation393_Full,

		MovTab: [13]uint16{
			0,           // N
			ScannerRoom, // NE
			0,           // E
			0,           // SE
			Corridor29,  // S
			0,           // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     Mortuary,
		BriefMsgNo: text.SolLocation394_Brief,
		FullMsgNo:  text.SolLocation394_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			0,          // W
			Corridor36, // NW
			0,          // Up
			0,          // Down
			0,          // In
			Corridor36, // Out
		},
	},
	{
		Number:     GuardRoom1,
		BriefMsgNo: text.SolLocation395_Brief,
		FullMsgNo:  text.SolLocation395_Full,
		SysLoc:     text.NO_MOVEMENT_30,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			SpacePort4, // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			SpacePort4, // Out
		},
	},
	{
		Number:     Cells,
		BriefMsgNo: text.SolLocation396_Brief,
		FullMsgNo:  text.SolLocation396_Full,
		Flags:      model.LfLock | model.LfShield,
	},
	{
		Number:     ParadeGround1,
		BriefMsgNo: text.SolLocation397_Brief,
		FullMsgNo:  text.SolLocation397_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			ParadeGround2, // E
			ParadeGround4, // SE
			ParadeGround3, // S
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
		Number:     ParadeGround2,
		BriefMsgNo: text.SolLocation398_Brief,
		FullMsgNo:  text.SolLocation398_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			Barracks2,     // N
			TarmacRoad,    // NE
			0,             // E
			0,             // SE
			ParadeGround4, // S
			ParadeGround3, // SW
			ParadeGround1, // W
			0,             // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     ParadeGround3,
		BriefMsgNo: text.SolLocation399_Brief,
		FullMsgNo:  text.SolLocation399_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			ParadeGround1, // N
			ParadeGround2, // NE
			ParadeGround4, // E
			0,             // SE
			0,             // S
			0,             // SW
			GuardRoom1,    // W
			0,             // NW
			0,             // Up
			0,             // Down
			0,             // In
			GuardRoom1,    // Out
		},
	},
	{
		Number:     ParadeGround4,
		BriefMsgNo: text.SolLocation400_Brief,
		FullMsgNo:  text.SolLocation400_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			ParadeGround2, // N
			0,             // NE
			RadarArray,    // E
			0,             // SE
			ControlTower1, // S
			0,             // SW
			ParadeGround3, // W
			ParadeGround1, // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     TarmacRoad,
		BriefMsgNo: text.SolLocation401_Brief,
		FullMsgNo:  text.SolLocation401_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			Road1,         // E
			0,             // SE
			0,             // S
			ParadeGround2, // SW
			0,             // W
			0,             // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     Road1,
		BriefMsgNo: text.SolLocation402_Brief,
		FullMsgNo:  text.SolLocation402_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0,                // N
			LaunchingCradle1, // NE
			Road2,            // E
			LaunchingCradle2, // SE
			0,                // S
			0,                // SW
			TarmacRoad,       // W
			0,                // NW
			0,                // Up
			0,                // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     Road2,
		BriefMsgNo: text.SolLocation403_Brief,
		FullMsgNo:  text.SolLocation403_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			Road3,   // E
			Office5, // SE
			0,       // S
			0,       // SW
			Road1,   // W
			0,       // NW
			0,       // Up
			0,       // Down
			Office5, // In
			0,       // Out
		},
	},
	{
		Number:     Road3,
		BriefMsgNo: text.SolLocation404_Brief,
		FullMsgNo:  text.SolLocation404_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			Hangar1, // N
			0,       // NE
			Road4,   // E
			0,       // SE
			0,       // S
			0,       // SW
			Road2,   // W
			0,       // NW
			0,       // Up
			0,       // Down
			Hangar1, // In
			0,       // Out
		},
	},
	{
		Number:     Road4,
		BriefMsgNo: text.SolLocation405_Brief,
		FullMsgNo:  text.SolLocation405_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			Hangar2, // N
			0,       // NE
			0,       // E
			0,       // SE
			Road5,   // S
			0,       // SW
			Road3,   // W
			0,       // NW
			0,       // Up
			0,       // Down
			Hangar2, // In
			0,       // Out
		},
	},
	{
		Number:     Road5,
		BriefMsgNo: text.SolLocation406_Brief,
		FullMsgNo:  text.SolLocation406_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			Road4,      // N
			0,          // NE
			0,          // E
			0,          // SE
			Road6,      // S
			Workshop26, // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			Workshop26, // In
			0,          // Out
		},
	},
	{
		Number:     Road6,
		BriefMsgNo: text.SolLocation407_Brief,
		FullMsgNo:  text.SolLocation407_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			Road5,              // N
			0,                  // NE
			LaunchCradle,       // E
			0,                  // SE
			BlockhouseEntrance, // S
			0,                  // SW
			0,                  // W
			0,                  // NW
			0,                  // Up
			0,                  // Down
			0,                  // In
			0,                  // Out
		},
	},
	{
		Number:     ControlTower1,
		BriefMsgNo: text.SolLocation408_Brief,
		FullMsgNo:  text.SolLocation408_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			ParadeGround4, // N
			0,             // NE
			0,             // E
			0,             // SE
			0,             // S
			0,             // SW
			0,             // W
			0,             // NW
			ControlTower2, // Up
			0,             // Down
			0,             // In
			ParadeGround4, // Out
		},
	},
	{
		Number:     ControlTower2,
		BriefMsgNo: text.SolLocation409_Brief,
		FullMsgNo:  text.SolLocation409_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			0,             // E
			0,             // SE
			0,             // S
			0,             // SW
			0,             // W
			0,             // NW
			0,             // Up
			ControlTower1, // Down
			0,             // In
			ControlTower1, // Out
		},
	},
	{
		Number:     RadarArray,
		BriefMsgNo: text.SolLocation410_Brief,
		FullMsgNo:  text.SolLocation410_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			0,             // E
			0,             // SE
			0,             // S
			0,             // SW
			ParadeGround4, // W
			0,             // NW
			0,             // Up
			0,             // Down
			0,             // In
			ParadeGround4, // Out
		},
	},
	{
		Number:     Barracks1,
		BriefMsgNo: text.SolLocation411_Brief,
		FullMsgNo:  text.SolLocation411_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			0,         // SE
			Barracks2, // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			Barracks2, // Out
		},
	},
	{
		Number:     Barracks2,
		BriefMsgNo: text.SolLocation412_Brief,
		FullMsgNo:  text.SolLocation412_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			Barracks1,     // N
			0,             // NE
			0,             // E
			0,             // SE
			ParadeGround2, // S
			0,             // SW
			0,             // W
			0,             // NW
			0,             // Up
			0,             // Down
			0,             // In
			ParadeGround2, // Out
		},
	},
	{
		Number:     LaunchingCradle1,
		BriefMsgNo: text.SolLocation413_Brief,
		FullMsgNo:  text.SolLocation413_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0,     // N
			0,     // NE
			0,     // E
			0,     // SE
			0,     // S
			Road1, // SW
			0,     // W
			0,     // NW
			0,     // Up
			0,     // Down
			0,     // In
			0,     // Out
		},
	},
	{
		Number:     LaunchingCradle2,
		BriefMsgNo: text.SolLocation414_Brief,
		FullMsgNo:  text.SolLocation414_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0,     // N
			0,     // NE
			0,     // E
			0,     // SE
			0,     // S
			0,     // SW
			0,     // W
			Road1, // NW
			0,     // Up
			0,     // Down
			0,     // In
			0,     // Out
		},
	},
	{
		Number:     LaunchCradle,
		BriefMsgNo: text.SolLocation415_Brief,
		FullMsgNo:  text.SolLocation415_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0,     // N
			0,     // NE
			0,     // E
			0,     // SE
			0,     // S
			0,     // SW
			Road6, // W
			0,     // NW
			0,     // Up
			0,     // Down
			0,     // In
			0,     // Out
		},
	},
	{
		Number:     Hangar1,
		BriefMsgNo: text.SolLocation416_Brief,
		FullMsgNo:  text.SolLocation416_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0,     // N
			0,     // NE
			0,     // E
			0,     // SE
			Road3, // S
			0,     // SW
			0,     // W
			0,     // NW
			0,     // Up
			0,     // Down
			0,     // In
			Road3, // Out
		},
	},
	{
		Number:     Hangar2,
		BriefMsgNo: text.SolLocation417_Brief,
		FullMsgNo:  text.SolLocation417_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0,     // N
			0,     // NE
			0,     // E
			0,     // SE
			Road4, // S
			0,     // SW
			0,     // W
			0,     // NW
			0,     // Up
			0,     // Down
			0,     // In
			Road4, // Out
		},
	},
	{
		Number:     Office5,
		BriefMsgNo: text.SolLocation418_Brief,
		FullMsgNo:  text.SolLocation418_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0,     // N
			0,     // NE
			0,     // E
			0,     // SE
			0,     // S
			0,     // SW
			0,     // W
			Road2, // NW
			0,     // Up
			0,     // Down
			0,     // In
			Road2, // Out
		},
	},
	{
		Number:     Workshop26,
		BriefMsgNo: text.SolLocation419_Brief,
		FullMsgNo:  text.SolLocation419_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0,     // N
			Road5, // NE
			0,     // E
			0,     // SE
			0,     // S
			0,     // SW
			0,     // W
			0,     // NW
			0,     // Up
			0,     // Down
			0,     // In
			Road5, // Out
		},
	},
	{
		Number:     BlockhouseEntrance,
		BriefMsgNo: text.SolLocation420_Brief,
		FullMsgNo:  text.SolLocation420_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			Road6,        // N
			0,            // NE
			0,            // E
			0,            // SE
			ControlRoom3, // S
			0,            // SW
			0,            // W
			0,            // NW
			0,            // Up
			ControlRoom3, // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     ControlRoom3,
		BriefMsgNo: text.SolLocation421_Brief,
		FullMsgNo:  text.SolLocation421_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			BlockhouseEntrance, // N
			0,                  // NE
			LaserControl,       // E
			0,                  // SE
			MissileControl,     // S
			0,                  // SW
			Telemetry,          // W
			0,                  // NW
			BlockhouseEntrance, // Up
			0,                  // Down
			0,                  // In
			0,                  // Out
		},
	},
	{
		Number:     Telemetry,
		BriefMsgNo: text.SolLocation422_Brief,
		FullMsgNo:  text.SolLocation422_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			ControlRoom3, // E
			0,            // SE
			0,            // S
			0,            // SW
			0,            // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			ControlRoom3, // Out
		},
	},
	{
		Number:     LaserControl,
		BriefMsgNo: text.SolLocation423_Brief,
		FullMsgNo:  text.SolLocation423_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			0,            // E
			0,            // SE
			0,            // S
			0,            // SW
			ControlRoom3, // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			ControlRoom3, // Out
		},
	},
	{
		Number:     MissileControl,
		BriefMsgNo: text.SolLocation424_Brief,
		FullMsgNo:  text.SolLocation424_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			ControlRoom3, // N
			0,            // NE
			0,            // E
			0,            // SE
			0,            // S
			0,            // SW
			0,            // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			ControlRoom3, // Out
		},
	},
	{
		Number:     ShipyardOffice,
		BriefMsgNo: text.SolLocation425_Brief,
		FullMsgNo:  text.SolLocation425_Full,
		Flags:      model.LfYard,

		MovTab: [13]uint16{
			0,              // N
			0,              // NE
			0,              // E
			0,              // SE
			0,              // S
			0,              // SW
			0,              // W
			PerimeterRoad1, // NW
			0,              // Up
			0,              // Down
			0,              // In
			PerimeterRoad1, // Out
		},
	},
	{
		Number:     EarthLandingArea,
		BriefMsgNo: text.SolLocation426_Brief,
		FullMsgNo:  text.SolLocation426_Full,
		Flags:      model.LfLanding,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			SpacePort4, // SE
			SpacePort3, // S
			SpacePort2, // SW
			SpacePort1, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     SpacePort1,
		BriefMsgNo: text.SolLocation427_Brief,
		FullMsgNo:  text.SolLocation427_Full,

		MovTab: [13]uint16{
			0,                // N
			0,                // NE
			EarthLandingArea, // E
			SpacePort3,       // SE
			SpacePort2,       // S
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
		Number:     SpacePort2,
		BriefMsgNo: text.SolLocation428_Brief,
		FullMsgNo:  text.SolLocation428_Full,

		MovTab: [13]uint16{
			SpacePort1,       // N
			EarthLandingArea, // NE
			SpacePort3,       // E
			0,                // SE
			0,                // S
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
		Number:     SpacePort3,
		BriefMsgNo: text.SolLocation429_Brief,
		FullMsgNo:  text.SolLocation429_Full,

		MovTab: [13]uint16{
			EarthLandingArea, // N
			0,                // NE
			SpacePort4,       // E
			0,                // SE
			Terminus1,        // S
			0,                // SW
			SpacePort2,       // W
			SpacePort1,       // NW
			0,                // Up
			0,                // Down
			Terminus1,        // In
			0,                // Out
		},
	},
	{
		Number:     SpacePort4,
		BriefMsgNo: text.SolLocation430_Brief,
		FullMsgNo:  text.SolLocation430_Full,

		MovTab: [13]uint16{
			0,                // N
			GuardRoom1,       // NE
			0,                // E
			0,                // SE
			0,                // S
			0,                // SW
			SpacePort3,       // W
			EarthLandingArea, // NW
			0,                // Up
			0,                // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     StarshipCantina,
		BriefMsgNo: text.SolLocation431_Brief,
		FullMsgNo:  text.SolLocation431_Full,
		Flags:      model.LfCafe,

		MovTab: [13]uint16{
			0,         // N
			Terminus1, // NE
			0,         // E
			0,         // SE
			0,         // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			Terminus1, // Out
		},
	},
	{
		Number:     Terminus1,
		BriefMsgNo: text.SolLocation432_Brief,
		FullMsgNo:  text.SolLocation432_Full,

		MovTab: [13]uint16{
			SpacePort3,      // N
			0,               // NE
			Terminus2,       // E
			0,               // SE
			RepairShop,      // S
			StarshipCantina, // SW
			MeetingPoint,    // W
			0,               // NW
			0,               // Up
			0,               // Down
			0,               // In
			SpacePort3,      // Out
		},
	},
	{
		Number:     Terminus2,
		BriefMsgNo: text.SolLocation433_Brief,
		FullMsgNo:  text.SolLocation433_Full,

		MovTab: [13]uint16{
			0,                // N
			0,                // NE
			TerminusExit,     // E
			0,                // SE
			TerminusEntrance, // S
			0,                // SW
			Terminus1,        // W
			0,                // NW
			0,                // Up
			0,                // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     TerminusExit,
		BriefMsgNo: text.SolLocation434_Brief,
		FullMsgNo:  text.SolLocation434_Full,

		MovTab: [13]uint16{
			0,              // N
			0,              // NE
			PerimeterRoad1, // E
			0,              // SE
			0,              // S
			0,              // SW
			0,              // W
			0,              // NW
			0,              // Up
			0,              // Down
			0,              // In
			PerimeterRoad1, // Out
		},
	},
	{
		Number:     TerminusEntrance,
		BriefMsgNo: text.SolLocation435_Brief,
		FullMsgNo:  text.SolLocation435_Full,

		MovTab: [13]uint16{
			Terminus2, // N
			0,         // NE
			0,         // E
			0,         // SE
			0,         // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			Terminus2, // In
			Terminus2, // Out
		},
	},
	{
		Number:     RepairShop,
		BriefMsgNo: text.SolLocation436_Brief,
		FullMsgNo:  text.SolLocation436_Full,
		Flags:      model.LfRep,

		MovTab: [13]uint16{
			Terminus1, // N
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
			Terminus1, // Out
		},
	},
	{
		Number:     PerimeterRoad1,
		BriefMsgNo: text.SolLocation437_Brief,
		FullMsgNo:  text.SolLocation437_Full,

		MovTab: [13]uint16{
			0,              // N
			0,              // NE
			0,              // E
			ShipyardOffice, // SE
			0,              // S
			PerimeterRoad2, // SW
			0,              // W
			0,              // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     PerimeterRoad2,
		BriefMsgNo: text.SolLocation438_Brief,
		FullMsgNo:  text.SolLocation438_Full,

		MovTab: [13]uint16{
			0,              // N
			PerimeterRoad1, // NE
			0,              // E
			0,              // SE
			0,              // S
			MainRoad3,      // SW
			0,              // W
			0,              // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     MainRoad3,
		BriefMsgNo: text.SolLocation439_Brief,
		FullMsgNo:  text.SolLocation439_Full,

		MovTab: [13]uint16{
			TerminusEntrance, // N
			PerimeterRoad2,   // NE
			0,                // E
			0,                // SE
			MainRoad4,        // S
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
		Number:     MainRoad4,
		BriefMsgNo: text.SolLocation440_Brief,
		FullMsgNo:  text.SolLocation440_Full,

		MovTab: [13]uint16{
			MainRoad3, // N
			0,         // NE
			0,         // E
			0,         // SE
			MainRoad5, // S
			0,         // SW
			Driveway2, // W
			0,         // NW
			0,         // Up
			0,         // Down
			Driveway2, // In
			0,         // Out
		},
	},
	{
		Number:     MainRoad5,
		BriefMsgNo: text.SolLocation441_Brief,
		FullMsgNo:  text.SolLocation441_Full,

		MovTab: [13]uint16{
			MainRoad4, // N
			0,         // NE
			MainRoad6, // E
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
		Number:     MainRoad6,
		BriefMsgNo: text.SolLocation442_Brief,
		FullMsgNo:  text.SolLocation442_Full,

		MovTab: [13]uint16{
			0,               // N
			0,               // NE
			MainRoad7,       // E
			0,               // SE
			ElectronicsShop, // S
			0,               // SW
			MainRoad5,       // W
			0,               // NW
			0,               // Up
			0,               // Down
			ElectronicsShop, // In
			0,               // Out
		},
	},
	{
		Number:     MainRoad7,
		BriefMsgNo: text.SolLocation443_Brief,
		FullMsgNo:  text.SolLocation443_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			MainRoad8, // E
			0,         // SE
			0,         // S
			0,         // SW
			MainRoad6, // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     MainRoad8,
		BriefMsgNo: text.SolLocation444_Brief,
		FullMsgNo:  text.SolLocation444_Full,

		MovTab: [13]uint16{
			WeaponShop3, // N
			0,           // NE
			CityRoad1,   // E
			0,           // SE
			0,           // S
			0,           // SW
			MainRoad7,   // W
			0,           // NW
			0,           // Up
			0,           // Down
			WeaponShop3, // In
			0,           // Out
		},
	},
	{
		Number:     CityRoad1,
		BriefMsgNo: text.SolLocation445_Brief,
		FullMsgNo:  text.SolLocation445_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			CityRoad6, // SE
			0,         // S
			0,         // SW
			MainRoad8, // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     Driveway1,
		BriefMsgNo: text.SolLocation446_Brief,
		FullMsgNo:  text.SolLocation446_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			Driveway2, // E
			0,         // SE
			0,         // S
			0,         // SW
			Hallway1,  // W
			0,         // NW
			0,         // Up
			0,         // Down
			Hallway1,  // In
			0,         // Out
		},
	},
	{
		Number:     Driveway2,
		BriefMsgNo: text.SolLocation447_Brief,
		FullMsgNo:  text.SolLocation447_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			MainRoad4, // E
			0,         // SE
			0,         // S
			0,         // SW
			Driveway1, // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			MainRoad4, // Out
		},
	},
	{
		Number:     Hallway1,
		BriefMsgNo: text.SolLocation448_Brief,
		FullMsgNo:  text.SolLocation448_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			Driveway1, // E
			0,         // SE
			Lounge,    // S
			0,         // SW
			Hallway2,  // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			Driveway1, // Out
		},
	},
	{
		Number:     Hallway2,
		BriefMsgNo: text.SolLocation449_Brief,
		FullMsgNo:  text.SolLocation449_Full,

		MovTab: [13]uint16{
			DrawingRoom1, // N
			0,            // NE
			Hallway1,     // E
			0,            // SE
			DiningRoom,   // S
			Kitchen,      // SW
			Backyard,     // W
			MainStairway, // NW
			MainStairway, // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     DrawingRoom1,
		BriefMsgNo: text.SolLocation450_Brief,
		FullMsgNo:  text.SolLocation450_Full,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			DrawingRoom2, // E
			0,            // SE
			Hallway2,     // S
			0,            // SW
			0,            // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			Hallway2,     // Out
		},
	},
	{
		Number:     DrawingRoom2,
		BriefMsgNo: text.SolLocation451_Brief,
		FullMsgNo:  text.SolLocation451_Full,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			0,            // E
			0,            // SE
			0,            // S
			0,            // SW
			DrawingRoom1, // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			DrawingRoom1, // Out
		},
	},
	{
		Number:     Backyard,
		BriefMsgNo: text.SolLocation452_Brief,
		FullMsgNo:  text.SolLocation452_Full,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			Hallway2, // E
			0,        // SE
			0,        // S
			0,        // SW
			0,        // W
			0,        // NW
			0,        // Up
			0,        // Down
			Hallway2, // In
			0,        // Out
		},
	},
	{
		Number:     Kitchen,
		BriefMsgNo: text.SolLocation453_Brief,
		FullMsgNo:  text.SolLocation453_Full,

		MovTab: [13]uint16{
			0,        // N
			Hallway2, // NE
			0,        // E
			0,        // SE
			0,        // S
			0,        // SW
			0,        // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			Hallway2, // Out
		},
	},
	{
		Number:     DiningRoom,
		BriefMsgNo: text.SolLocation454_Brief,
		FullMsgNo:  text.SolLocation454_Full,

		MovTab: [13]uint16{
			Hallway2, // N
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
			Hallway2, // Out
		},
	},
	{
		Number:     Lounge,
		BriefMsgNo: text.SolLocation455_Brief,
		FullMsgNo:  text.SolLocation455_Full,

		MovTab: [13]uint16{
			Hallway1, // N
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
			Hallway1, // Out
		},
	},
	{
		Number:     MainStairway,
		BriefMsgNo: text.SolLocation456_Brief,
		FullMsgNo:  text.SolLocation456_Full,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			0,             // E
			Hallway2,      // SE
			0,             // S
			0,             // SW
			0,             // W
			0,             // NW
			MainStaircase, // Up
			Hallway2,      // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     MainStaircase,
		BriefMsgNo: text.SolLocation457_Brief,
		FullMsgNo:  text.SolLocation457_Full,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			0,            // E
			Landing1,     // SE
			0,            // S
			0,            // SW
			0,            // W
			0,            // NW
			0,            // Up
			MainStairway, // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     SpiralStaircase1,
		BriefMsgNo: text.SolLocation458_Brief,
		FullMsgNo:  text.SolLocation458_Full,

		MovTab: [13]uint16{
			0,                // N
			0,                // NE
			0,                // E
			0,                // SE
			Landing1,         // S
			0,                // SW
			0,                // W
			0,                // NW
			SpiralStaircase2, // Up
			0,                // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     SpiralStaircase2,
		BriefMsgNo: text.SolLocation459_Brief,
		FullMsgNo:  text.SolLocation459_Full,

		MovTab: [13]uint16{
			0,                // N
			0,                // NE
			Library1,         // E
			0,                // SE
			0,                // S
			0,                // SW
			0,                // W
			0,                // NW
			0,                // Up
			SpiralStaircase1, // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     Study,
		BriefMsgNo: text.SolLocation460_Brief,
		FullMsgNo:  text.SolLocation460_Full,

		MovTab: [13]uint16{
			0,        // N
			Landing1, // NE
			0,        // E
			0,        // SE
			0,        // S
			0,        // SW
			0,        // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			Landing1, // Out
		},
	},
	{
		Number:     Bedroom1,
		BriefMsgNo: text.SolLocation461_Brief,
		FullMsgNo:  text.SolLocation461_Full,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			0,        // E
			0,        // SE
			0,        // S
			Landing1, // SW
			0,        // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			Landing1, // Out
		},
	},
	{
		Number:     Bedroom2,
		BriefMsgNo: text.SolLocation462_Brief,
		FullMsgNo:  text.SolLocation462_Full,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			Landing2, // E
			0,        // SE
			0,        // S
			0,        // SW
			0,        // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			Landing2, // Out
		},
	},
	{
		Number:     Library1,
		BriefMsgNo: text.SolLocation463_Brief,
		FullMsgNo:  text.SolLocation463_Full,

		MovTab: [13]uint16{
			0,                // N
			0,                // NE
			0,                // E
			0,                // SE
			0,                // S
			0,                // SW
			SpiralStaircase2, // W
			0,                // NW
			0,                // Up
			0,                // Down
			0,                // In
			SpiralStaircase2, // Out
		},
	},
	{
		Number:     Bathroom,
		BriefMsgNo: text.SolLocation464_Brief,
		FullMsgNo:  text.SolLocation464_Full,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			0,        // E
			0,        // SE
			Landing2, // S
			0,        // SW
			0,        // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			Landing2, // Out
		},
	},
	{
		Number:     Landing1,
		BriefMsgNo: text.SolLocation465_Brief,
		FullMsgNo:  text.SolLocation465_Full,

		MovTab: [13]uint16{
			SpiralStaircase1, // N
			Bedroom1,         // NE
			0,                // E
			Landing2,         // SE
			0,                // S
			Study,            // SW
			0,                // W
			MainStaircase,    // NW
			0,                // Up
			0,                // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     Landing2,
		BriefMsgNo: text.SolLocation466_Brief,
		FullMsgNo:  text.SolLocation466_Full,

		MovTab: [13]uint16{
			Bathroom, // N
			0,        // NE
			0,        // E
			0,        // SE
			0,        // S
			0,        // SW
			Bedroom2, // W
			Landing1, // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     SecretRoom,
		BriefMsgNo: text.SolLocation467_Brief,
		Events:     [2]uint16{39, 0},
		FullMsgNo:  text.SolLocation467_Full,

		MovTab: [13]uint16{
			0,        // N
			Library1, // NE
			0,        // E
			0,        // SE
			0,        // S
			0,        // SW
			0,        // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			Library1, // Out
		},
	},
	{
		Number:     WeaponShop3,
		BriefMsgNo: text.SolLocation468_Brief,
		FullMsgNo:  text.SolLocation468_Full,
		Flags:      model.LfWeap,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			0,         // SE
			MainRoad8, // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			MainRoad8, // Out
		},
	},
	{
		Number:     ElectronicsShop,
		BriefMsgNo: text.SolLocation469_Brief,
		FullMsgNo:  text.SolLocation469_Full,
		Flags:      model.LfCom,

		MovTab: [13]uint16{
			MainRoad6, // N
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
			MainRoad6, // Out
		},
	},
	{
		Number:     Terminal,
		BriefMsgNo: text.SolLocation470_Brief,
		FullMsgNo:  text.SolLocation470_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			GuardRoom5, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			GuardRoom5, // Out
		},
	},
	{
		Number:     TopOfStairs,
		BriefMsgNo: text.SolLocation471_Brief,
		FullMsgNo:  text.SolLocation471_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			Office6,    // S
			0,          // SW
			Corridor37, // W
			0,          // NW
			0,          // Up
			Stairs3,    // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Office6,
		BriefMsgNo: text.SolLocation472_Brief,
		FullMsgNo:  text.SolLocation472_Full,

		MovTab: [13]uint16{
			TopOfStairs, // N
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
			TopOfStairs, // Out
		},
	},
	{
		Number:     Corridor37,
		BriefMsgNo: text.SolLocation473_Brief,
		FullMsgNo:  text.SolLocation473_Full,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			TopOfStairs, // E
			0,           // SE
			Cloakroom,   // S
			0,           // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     Cloakroom,
		BriefMsgNo: text.SolLocation474_Brief,
		FullMsgNo:  text.SolLocation474_Full,

		MovTab: [13]uint16{
			Corridor37, // N
			0,          // NE
			0,          // E
			Corridor38, // SE
			0,          // S
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
		Number:     Corridor38,
		BriefMsgNo: text.SolLocation475_Brief,
		FullMsgNo:  text.SolLocation475_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			Stairs1,   // E
			0,         // SE
			0,         // S
			0,         // SW
			0,         // W
			Cloakroom, // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     Stairs1,
		BriefMsgNo: text.SolLocation476_Brief,
		FullMsgNo:  text.SolLocation476_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			Corridor38, // W
			0,          // NW
			Stairs2,    // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Stairs2,
		BriefMsgNo: text.SolLocation477_Brief,
		FullMsgNo:  text.SolLocation477_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Office7,    // E
			Corridor39, // SE
			0,          // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			Stairs1,    // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Office7,
		BriefMsgNo: text.SolLocation478_Brief,
		FullMsgNo:  text.SolLocation478_Full,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			0,       // E
			0,       // SE
			0,       // S
			0,       // SW
			Stairs2, // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			Stairs2, // Out
		},
	},
	{
		Number:     Corridor39,
		BriefMsgNo: text.SolLocation479_Brief,
		FullMsgNo:  text.SolLocation479_Full,

		MovTab: [13]uint16{
			0,              // N
			0,              // NE
			0,              // E
			0,              // SE
			0,              // S
			PenthouseSuite, // SW
			0,              // W
			Stairs2,        // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     PenthouseSuite,
		BriefMsgNo: text.SolLocation480_Brief,
		FullMsgNo:  text.SolLocation480_Full,

		MovTab: [13]uint16{
			0,          // N
			Corridor39, // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			Corridor39, // Out
		},
	},
	{
		Number:     CityRoad2,
		BriefMsgNo: text.SolLocation481_Brief,
		FullMsgNo:  text.SolLocation481_Full,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			CityRoad3,    // E
			0,            // SE
			SnackBar,     // S
			0,            // SW
			CityStreet12, // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     CityRoad3,
		BriefMsgNo: text.SolLocation482_Brief,
		FullMsgNo:  text.SolLocation482_Full,

		MovTab: [13]uint16{
			MarketPlace1, // N
			0,            // NE
			CityRoad4,    // E
			0,            // SE
			0,            // S
			0,            // SW
			CityRoad2,    // W
			0,            // NW
			0,            // Up
			0,            // Down
			MarketPlace1, // In
			0,            // Out
		},
	},
	{
		Number:     CityRoad4,
		BriefMsgNo: text.SolLocation483_Brief,
		FullMsgNo:  text.SolLocation483_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			0,         // SE
			CityRoad5, // S
			0,         // SW
			CityRoad3, // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     CityRoad5,
		BriefMsgNo: text.SolLocation484_Brief,
		FullMsgNo:  text.SolLocation484_Full,

		MovTab: [13]uint16{
			CityRoad4, // N
			0,         // NE
			MainRoad9, // E
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
		Number:     MainRoad9,
		BriefMsgNo: text.SolLocation485_Brief,
		FullMsgNo:  text.SolLocation485_Full,

		MovTab: [13]uint16{
			0,              // N
			0,              // NE
			MainRoad10,     // E
			PublicCounter1, // SE
			0,              // S
			0,              // SW
			CityRoad5,      // W
			0,              // NW
			0,              // Up
			0,              // Down
			PublicCounter1, // In
			0,              // Out
		},
	},
	{
		Number:     MainRoad10,
		BriefMsgNo: text.SolLocation486_Brief,
		FullMsgNo:  text.SolLocation486_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			MainRoad11, // SE
			0,          // S
			0,          // SW
			MainRoad9,  // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     SnackBar,
		BriefMsgNo: text.SolLocation487_Brief,
		FullMsgNo:  text.SolLocation487_Full,
		Flags:      model.LfCafe,

		MovTab: [13]uint16{
			CityRoad2, // N
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
			CityRoad2, // Out
		},
	},
	{
		Number:     MarketPlace1,
		BriefMsgNo: text.SolLocation488_Brief,
		FullMsgNo:  text.SolLocation488_Full,
		Flags:      model.LfCafe,

		MovTab: [13]uint16{
			ExchangeGallery, // N
			ExExchange,      // NE
			MarketPlace2,    // E
			0,               // SE
			CityRoad3,       // S
			0,               // SW
			0,               // W
			0,               // NW
			0,               // Up
			0,               // Down
			0,               // In
			CityRoad3,       // Out
		},
	},
	{
		Number:     MarketPlace2,
		BriefMsgNo: text.SolLocation489_Brief,
		FullMsgNo:  text.SolLocation489_Full,
		Flags:      model.LfClth,

		MovTab: [13]uint16{
			ExExchange,      // N
			0,               // NE
			0,               // E
			0,               // SE
			0,               // S
			0,               // SW
			MarketPlace1,    // W
			ExchangeGallery, // NW
			0,               // Up
			0,               // Down
			0,               // In
			0,               // Out
		},
	},
	{
		Number:     ExchangeGallery,
		BriefMsgNo: text.SolLocation490_Brief,
		FullMsgNo:  text.SolLocation490_Full,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			ExExchange,   // E
			MarketPlace2, // SE
			MarketPlace1, // S
			0,            // SW
			0,            // W
			0,            // NW
			0,            // Up
			ExExchange,   // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     ExExchange,
		BriefMsgNo: text.SolLocation491_Brief,
		FullMsgNo:  text.SolLocation491_Full,

		MovTab: [13]uint16{
			0,               // N
			0,               // NE
			TaxOffice,       // E
			0,               // SE
			MarketPlace2,    // S
			MarketPlace1,    // SW
			ExchangeGallery, // W
			0,               // NW
			ExchangeGallery, // Up
			0,               // Down
			TaxOffice,       // In
			0,               // Out
		},
	},
	{
		Number:     TaxOffice,
		BriefMsgNo: text.SolLocation492_Brief,
		FullMsgNo:  text.SolLocation492_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			ExExchange, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			ExExchange, // Out
		},
	},
	{
		Number:     PublicCounter1,
		BriefMsgNo: text.SolLocation493_Brief,
		FullMsgNo:  text.SolLocation493_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			GuardRoom2, // W
			MainRoad9,  // NW
			0,          // Up
			0,          // Down
			GuardRoom2, // In
			MainRoad9,  // Out
		},
	},
	{
		Number:     GuardRoom2,
		BriefMsgNo: text.SolLocation494_Brief,
		FullMsgNo:  text.SolLocation494_Full,

		MovTab: [13]uint16{
			0,              // N
			0,              // NE
			PublicCounter1, // E
			0,              // SE
			GuardRoom3,     // S
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
		Number:     GuardRoom3,
		BriefMsgNo: text.SolLocation495_Brief,
		FullMsgNo:  text.SolLocation495_Full,

		MovTab: [13]uint16{
			GuardRoom2, // N
			0,          // NE
			CourtRoom,  // E
			Cell,       // SE
			0,          // S
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
		Number:     CourtRoom,
		BriefMsgNo: text.SolLocation496_Brief,
		FullMsgNo:  text.SolLocation496_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			GuardRoom3, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			GuardRoom3, // Out
		},
	},
	{
		Number:     Cell,
		BriefMsgNo: text.SolLocation497_Brief,
		FullMsgNo:  text.SolLocation497_Full,
	},
	{
		Number:     MainRoad11,
		BriefMsgNo: text.SolLocation498_Brief,
		FullMsgNo:  text.SolLocation498_Full,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			ForkInTheRoad, // E
			0,             // SE
			0,             // S
			0,             // SW
			0,             // W
			MainRoad10,    // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     ForkInTheRoad,
		BriefMsgNo: text.SolLocation499_Brief,
		FullMsgNo:  text.SolLocation499_Full,

		MovTab: [13]uint16{
			0,           // N
			BroadRoad,   // NE
			0,           // E
			0,           // SE
			NarrowRoad1, // S
			0,           // SW
			MainRoad11,  // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     NarrowRoad1,
		BriefMsgNo: text.SolLocation500_Brief,
		FullMsgNo:  text.SolLocation500_Full,

		MovTab: [13]uint16{
			ForkInTheRoad, // N
			0,             // NE
			0,             // E
			Road7,         // SE
			0,             // S
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
		Number:     Road7,
		BriefMsgNo: text.SolLocation501_Brief,
		FullMsgNo:  text.SolLocation501_Full,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			Road8,       // E
			0,           // SE
			0,           // S
			0,           // SW
			0,           // W
			NarrowRoad1, // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     Road8,
		BriefMsgNo: text.SolLocation502_Brief,
		FullMsgNo:  text.SolLocation502_Full,

		MovTab: [13]uint16{
			Road9, // N
			0,     // NE
			0,     // E
			Foyer, // SE
			0,     // S
			0,     // SW
			Road7, // W
			0,     // NW
			0,     // Up
			0,     // Down
			Foyer, // In
			0,     // Out
		},
	},
	{
		Number:     Road9,
		BriefMsgNo: text.SolLocation503_Brief,
		FullMsgNo:  text.SolLocation503_Full,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			0,           // E
			0,           // SE
			Road8,       // S
			0,           // SW
			0,           // W
			NarrowRoad2, // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     NarrowRoad2,
		BriefMsgNo: text.SolLocation504_Brief,
		FullMsgNo:  text.SolLocation504_Full,

		MovTab: [13]uint16{
			BroadRoad, // N
			0,         // NE
			0,         // E
			Road9,     // SE
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
		Number:     BroadRoad,
		BriefMsgNo: text.SolLocation505_Brief,
		FullMsgNo:  text.SolLocation505_Full,

		MovTab: [13]uint16{
			0,                            // N
			0,                            // NE
			EndOfTheRoad,                 // E
			0,                            // SE
			NarrowRoad2,                  // S
			ForkInTheRoad,                // SW
			InstaLernFactORamaUniversity, // W
			0,                            // NW
			0,                            // Up
			0,                            // Down
			InstaLernFactORamaUniversity, // In
			0,                            // Out
		},
	},
	{
		Number:     EndOfTheRoad,
		BriefMsgNo: text.SolLocation506_Brief,
		FullMsgNo:  text.SolLocation506_Full,

		MovTab: [13]uint16{
			Loo,                                   // N
			0,                                     // NE
			BattleCreekSanitarium,                 // E
			FeeldaBurnesAerobicsAndWorkoutClasses, // SE
			PamperUHealthFarm,                     // S
			0,                                     // SW
			BroadRoad,                             // W
			0,                                     // NW
			0,                                     // Up
			0,                                     // Down
			Loo,                                   // In
			0,                                     // Out
		},
	},
	{
		Number:     InstaLernFactORamaUniversity,
		BriefMsgNo: text.SolLocation507_Brief,
		FullMsgNo:  text.SolLocation507_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			BroadRoad, // E
			0,         // SE
			0,         // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			BroadRoad, // Out
		},
	},
	{
		Number:     HagarsMusicStore,
		BriefMsgNo: text.SolLocation508_Brief,
		FullMsgNo:  text.SolLocation508_Full,
		SysLoc:     text.NO_MOVEMENT_23,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			WideRoad, // E
			0,        // SE
			0,        // S
			0,        // SW
			0,        // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			WideRoad, // Out
		},
	},
	{
		Number:     UnusedLocation509,
		BriefMsgNo: text.NullLocation,
		FullMsgNo:  text.NullLocation,
		Flags:      model.LfDeath | model.LfLock | model.LfShield,
	},
	{
		Number:     BattleCreekSanitarium,
		BriefMsgNo: text.SolLocation510_Brief,
		FullMsgNo:  text.SolLocation510_Full,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			0,            // E
			0,            // SE
			0,            // S
			0,            // SW
			EndOfTheRoad, // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			EndOfTheRoad, // Out
		},
	},
	{
		Number:     FeeldaBurnesAerobicsAndWorkoutClasses,
		BriefMsgNo: text.SolLocation511_Brief,
		FullMsgNo:  text.SolLocation511_Full,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			0,            // E
			0,            // SE
			0,            // S
			0,            // SW
			0,            // W
			EndOfTheRoad, // NW
			0,            // Up
			0,            // Down
			0,            // In
			EndOfTheRoad, // Out
		},
	},
	{
		Number:     PamperUHealthFarm,
		BriefMsgNo: text.SolLocation512_Brief,
		FullMsgNo:  text.SolLocation512_Full,

		MovTab: [13]uint16{
			EndOfTheRoad, // N
			0,            // NE
			0,            // E
			0,            // SE
			0,            // S
			0,            // SW
			0,            // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			EndOfTheRoad, // Out
		},
	},
	{
		Number:     GentsLoo,
		BriefMsgNo: text.SolLocation513_Brief,
		FullMsgNo:  text.SolLocation513_Full,

		MovTab: [13]uint16{
			0,   // N
			0,   // NE
			0,   // E
			0,   // SE
			0,   // S
			0,   // SW
			Loo, // W
			0,   // NW
			0,   // Up
			0,   // Down
			0,   // In
			Loo, // Out
		},
	},
	{
		Number:     Loo,
		BriefMsgNo: text.SolLocation514_Brief,
		FullMsgNo:  text.SolLocation514_Full,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			GentsLoo,     // E
			0,            // SE
			EndOfTheRoad, // S
			0,            // SW
			LadiesLoo,    // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     LadiesLoo,
		BriefMsgNo: text.SolLocation515_Brief,
		FullMsgNo:  text.SolLocation515_Full,

		MovTab: [13]uint16{
			0,   // N
			0,   // NE
			Loo, // E
			0,   // SE
			0,   // S
			0,   // SW
			0,   // W
			0,   // NW
			0,   // Up
			0,   // Down
			0,   // In
			Loo, // Out
		},
	},
	{
		Number:     Museum1,
		BriefMsgNo: text.SolLocation516_Brief,
		FullMsgNo:  text.SolLocation516_Full,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			Museum2, // E
			0,       // SE
			0,       // S
			0,       // SW
			0,       // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			Museum2, // Out
		},
	},
	{
		Number:     Museum2,
		BriefMsgNo: text.SolLocation517_Brief,
		FullMsgNo:  text.SolLocation517_Full,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			0,           // E
			0,           // SE
			MuseumLobby, // S
			Museum3,     // SW
			Museum1,     // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     Museum3,
		BriefMsgNo: text.SolLocation518_Brief,
		FullMsgNo:  text.SolLocation518_Full,

		MovTab: [13]uint16{
			0,           // N
			Museum2,     // NE
			MuseumLobby, // E
			0,           // SE
			0,           // S
			0,           // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			MuseumLobby, // Out
		},
	},
	{
		Number:     MuseumLobby,
		BriefMsgNo: text.SolLocation519_Brief,
		FullMsgNo:  text.SolLocation519_Full,

		MovTab: [13]uint16{
			Museum2, // N
			0,       // NE
			0,       // E
			Plaza1,  // SE
			0,       // S
			0,       // SW
			Museum3, // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			Plaza1,  // Out
		},
	},
	{
		Number:     DrFoggsMaritalArtsEmporium,
		BriefMsgNo: text.SolLocation520_Brief,
		FullMsgNo:  text.SolLocation520_Full,

		MovTab: [13]uint16{
			0,      // N
			0,      // NE
			0,      // E
			0,      // SE
			Plaza2, // S
			0,      // SW
			0,      // W
			0,      // NW
			0,      // Up
			0,      // Down
			0,      // In
			Plaza2, // Out
		},
	},
	{
		Number:     ControlRoom4,
		BriefMsgNo: text.SolLocation521_Brief,
		FullMsgNo:  text.SolLocation521_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			OpenSpace5, // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			OpenSpace5, // Out
		},
	},
	{
		Number:     StoreRoom1,
		BriefMsgNo: text.SolLocation522_Brief,
		FullMsgNo:  text.SolLocation522_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			Slideway2, // SE
			0,         // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			Slideway2, // Out
		},
	},
	{
		Number:     Plaza1,
		BriefMsgNo: text.SolLocation523_Brief,
		FullMsgNo:  text.SolLocation523_Full,

		MovTab: [13]uint16{
			0,               // N
			0,               // NE
			Plaza2,          // E
			0,               // SE
			TheSpaceportBar, // S
			Slideway2,       // SW
			0,               // W
			MuseumLobby,     // NW
			0,               // Up
			0,               // Down
			TheSpaceportBar, // In
			0,               // Out
		},
	},
	{
		Number:     Plaza2,
		BriefMsgNo: text.SolLocation524_Brief,
		FullMsgNo:  text.SolLocation524_Full,

		MovTab: [13]uint16{
			DrFoggsMaritalArtsEmporium, // N
			0,                          // NE
			Slideway1,                  // E
			0,                          // SE
			0,                          // S
			0,                          // SW
			Plaza1,                     // W
			0,                          // NW
			0,                          // Up
			0,                          // Down
			DrFoggsMaritalArtsEmporium, // In
			0,                          // Out
		},
	},
	{
		Number:     Slideway1,
		BriefMsgNo: text.SolLocation525_Brief,
		FullMsgNo:  text.SolLocation525_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			OpenSpace5, // E
			0,          // SE
			0,          // S
			0,          // SW
			Plaza2,     // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     OpenSpace5,
		BriefMsgNo: text.SolLocation526_Brief,
		FullMsgNo:  text.SolLocation526_Full,

		MovTab: [13]uint16{
			ControlRoom4, // N
			0,            // NE
			Park4,        // E
			0,            // SE
			Slideway3,    // S
			0,            // SW
			Slideway1,    // W
			0,            // NW
			0,            // Up
			0,            // Down
			ControlRoom4, // In
			0,            // Out
		},
	},
	{
		Number:     Park4,
		BriefMsgNo: text.SolLocation527_Brief,
		FullMsgNo:  text.SolLocation527_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Park5,      // E
			0,          // SE
			0,          // S
			0,          // SW
			OpenSpace5, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			OpenSpace5, // Out
		},
	},
	{
		Number:     Park5,
		BriefMsgNo: text.SolLocation528_Brief,
		FullMsgNo:  text.SolLocation528_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,                      // N
			0,                      // NE
			0,                      // E
			0,                      // SE
			InsuranceBrokersOffice, // S
			0,                      // SW
			Park4,                  // W
			0,                      // NW
			0,                      // Up
			0,                      // Down
			InsuranceBrokersOffice, // In
			Park4,                  // Out
		},
	},
	{
		Number:     LandingBay1,
		BriefMsgNo: text.SolLocation529_Brief,
		FullMsgNo:  text.SolLocation529_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			SpacePort5, // SE
			0,          // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			SpacePort5, // Out
		},
	},
	{
		Number:     LandingBay2,
		BriefMsgNo: text.SolLocation530_Brief,
		FullMsgNo:  text.SolLocation530_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			SpacePort5, // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			SpacePort5, // Out
		},
	},
	{
		Number:     SpacePort5,
		BriefMsgNo: text.SolLocation531_Brief,
		FullMsgNo:  text.SolLocation531_Full,

		MovTab: [13]uint16{
			LandingBay2, // N
			0,           // NE
			TransitArea, // E
			0,           // SE
			SpacePort6,  // S
			0,           // SW
			0,           // W
			LandingBay1, // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     TransitArea,
		BriefMsgNo: text.SolLocation532_Brief,
		FullMsgNo:  text.SolLocation532_Full,

		MovTab: [13]uint16{
			0,          // N
			Slideway2,  // NE
			0,          // E
			Slideway5,  // SE
			0,          // S
			0,          // SW
			SpacePort5, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     LandingBay3,
		BriefMsgNo: text.SolLocation533_Brief,
		FullMsgNo:  text.SolLocation533_Full,
		Flags:      model.LfLanding,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			SpacePort6, // E
			0,          // SE
			0,          // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			SpacePort6, // Out
		},
	},
	{
		Number:     SpacePort6,
		BriefMsgNo: text.SolLocation534_Brief,
		FullMsgNo:  text.SolLocation534_Full,

		MovTab: [13]uint16{
			SpacePort5,  // N
			0,           // NE
			0,           // E
			0,           // SE
			LandingBay4, // S
			0,           // SW
			LandingBay3, // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     LandingBay4,
		BriefMsgNo: text.SolLocation535_Brief,
		FullMsgNo:  text.SolLocation535_Full,

		MovTab: [13]uint16{
			SpacePort6, // N
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
			SpacePort6, // Out
		},
	},
	{
		Number:     Slideway2,
		BriefMsgNo: text.SolLocation536_Brief,
		FullMsgNo:  text.SolLocation536_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,           // N
			Plaza1,      // NE
			0,           // E
			0,           // SE
			0,           // S
			TransitArea, // SW
			0,           // W
			StoreRoom1,  // NW
			0,           // Up
			0,           // Down
			StoreRoom1,  // In
			0,           // Out
		},
	},
	{
		Number:     TheSpaceportBar,
		BriefMsgNo: text.SolLocation537_Brief,
		FullMsgNo:  text.SolLocation537_Full,
		Flags:      model.LfCafe,

		MovTab: [13]uint16{
			Plaza1, // N
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
			Plaza1, // Out
		},
	},
	{
		Number:     Slideway3,
		BriefMsgNo: text.SolLocation538_Brief,
		FullMsgNo:  text.SolLocation538_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			OpenSpace5,  // N
			0,           // NE
			0,           // E
			PowerPlant2, // SE
			Slideway4,   // S
			PartyHQ,     // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     PowerPlant1,
		BriefMsgNo: text.SolLocation539_Brief,
		FullMsgNo:  text.SolLocation539_Full,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			0,           // E
			PowerPlant3, // SE
			PowerPlant2, // S
			0,           // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     InsuranceBrokersOffice,
		BriefMsgNo: text.SolLocation540_Brief,
		FullMsgNo:  text.SolLocation540_Full,
		Flags:      model.LfIns,

		MovTab: [13]uint16{
			Park5, // N
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
			Park5, // Out
		},
	},
	{
		Number:     PartyHQ,
		BriefMsgNo: text.SolLocation541_Brief,
		FullMsgNo:  text.SolLocation541_Full,

		MovTab: [13]uint16{
			0,         // N
			Slideway3, // NE
			0,         // E
			0,         // SE
			0,         // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			Slideway3, // Out
		},
	},
	{
		Number:     Slideway4,
		BriefMsgNo: text.SolLocation542_Brief,
		FullMsgNo:  text.SolLocation542_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			Slideway3,        // N
			0,                // NE
			0,                // E
			0,                // SE
			TradingExchange2, // S
			0,                // SW
			0,                // W
			0,                // NW
			0,                // Up
			0,                // Down
			TradingExchange2, // In
			0,                // Out
		},
	},
	{
		Number:     PowerPlant2,
		BriefMsgNo: text.SolLocation543_Brief,
		FullMsgNo:  text.SolLocation543_Full,

		MovTab: [13]uint16{
			PowerPlant1, // N
			0,           // NE
			0,           // E
			0,           // SE
			0,           // S
			0,           // SW
			0,           // W
			Slideway3,   // NW
			0,           // Up
			0,           // Down
			PowerPlant1, // In
			Slideway3,   // Out
		},
	},
	{
		Number:     PowerPlant3,
		BriefMsgNo: text.SolLocation544_Brief,
		FullMsgNo:  text.SolLocation544_Full,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			0,           // E
			0,           // SE
			0,           // S
			0,           // SW
			0,           // W
			PowerPlant1, // NW
			0,           // Up
			0,           // Down
			0,           // In
			PowerPlant1, // Out
		},
	},
	{
		Number:     TradingExchange2,
		BriefMsgNo: text.SolLocation545_Brief,
		FullMsgNo:  text.SolLocation545_Full,
		Flags:      model.LfTrade,

		MovTab: [13]uint16{
			Slideway4, // N
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
			Slideway4, // Out
		},
	},
	{
		Number:     Slideway5,
		BriefMsgNo: text.SolLocation546_Brief,
		FullMsgNo:  text.SolLocation546_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			0,           // E
			Slideway6,   // SE
			0,           // S
			0,           // SW
			0,           // W
			TransitArea, // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     Slideway6,
		BriefMsgNo: text.SolLocation547_Brief,
		FullMsgNo:  text.SolLocation547_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,               // N
			0,               // NE
			0,               // E
			ShuttleStation1, // SE
			0,               // S
			0,               // SW
			0,               // W
			Slideway5,       // NW
			0,               // Up
			0,               // Down
			0,               // In
			0,               // Out
		},
	},
	{
		Number:     ShuttleStation1,
		BriefMsgNo: text.SolLocation548_Brief,
		FullMsgNo:  text.SolLocation548_Full,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			0,         // SE
			0,         // S
			0,         // SW
			0,         // W
			Slideway6, // NW
			0,         // Up
			0,         // Down
			Shuttle,   // In
			Slideway6, // Out
		},
	},
	{
		Number:     Shuttle,
		BriefMsgNo: text.SolLocation549_Brief,
		FullMsgNo:  text.SolLocation549_Full,
		Flags:      model.LfShield,

		MovTab: [13]uint16{
			0,               // N
			0,               // NE
			0,               // E
			0,               // SE
			0,               // S
			0,               // SW
			0,               // W
			0,               // NW
			0,               // Up
			0,               // Down
			0,               // In
			ShuttleStation1, // Out
		},
	},
	{
		Number:     RockFace2,
		BriefMsgNo: text.SolLocation550_Brief,
		FullMsgNo:  text.SolLocation550_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			Corridor41, // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			Corridor41, // Out
		},
	},
	{
		Number:     RockFace3,
		BriefMsgNo: text.SolLocation551_Brief,
		FullMsgNo:  text.SolLocation551_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			0,        // E
			0,        // SE
			Workings, // S
			0,        // SW
			0,        // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			Workings, // Out
		},
	},
	{
		Number:     Corridor40,
		BriefMsgNo: text.SolLocation552_Brief,
		FullMsgNo:  text.SolLocation552_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			Corridor41,    // E
			0,             // SE
			0,             // S
			MineEntrance1, // SW
			0,             // W
			0,             // NW
			0,             // Up
			0,             // Down
			0,             // In
			MineEntrance1, // Out
		},
	},
	{
		Number:     Corridor41,
		BriefMsgNo: text.SolLocation553_Brief,
		FullMsgNo:  text.SolLocation553_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			RockFace2,  // N
			0,          // NE
			Tunnel1,    // E
			0,          // SE
			RockFace4,  // S
			0,          // SW
			Corridor40, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Tunnel1,
		BriefMsgNo: text.SolLocation554_Brief,
		FullMsgNo:  text.SolLocation554_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			Tunnel2,    // SE
			0,          // S
			0,          // SW
			Corridor41, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Workings,
		BriefMsgNo: text.SolLocation555_Brief,
		FullMsgNo:  text.SolLocation555_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			RockFace3, // N
			0,         // NE
			0,         // E
			0,         // SE
			0,         // S
			Tunnel2,   // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			Tunnel2,   // Out
		},
	},
	{
		Number:     AbandonedFace1,
		BriefMsgNo: text.SolLocation556_Brief,
		FullMsgNo:  text.SolLocation556_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			MineEntrance1, // E
			0,             // SE
			0,             // S
			0,             // SW
			0,             // W
			0,             // NW
			0,             // Up
			0,             // Down
			0,             // In
			MineEntrance1, // Out
		},
	},
	{
		Number:     MineEntrance1,
		BriefMsgNo: text.SolLocation557_Brief,
		FullMsgNo:  text.SolLocation557_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,               // N
			Corridor40,      // NE
			0,               // E
			ShuttleStation2, // SE
			0,               // S
			0,               // SW
			AbandonedFace1,  // W
			0,               // NW
			0,               // Up
			0,               // Down
			0,               // In
			0,               // Out
		},
	},
	{
		Number:     RockFace4,
		BriefMsgNo: text.SolLocation558_Brief,
		FullMsgNo:  text.SolLocation558_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			Corridor41, // N
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
			Corridor41, // Out
		},
	},
	{
		Number:     Tunnel2,
		BriefMsgNo: text.SolLocation559_Brief,
		FullMsgNo:  text.SolLocation559_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,        // N
			Workings, // NE
			0,        // E
			0,        // SE
			0,        // S
			0,        // SW
			0,        // W
			Tunnel1,  // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     ShuttleStation2,
		BriefMsgNo: text.SolLocation560_Brief,
		FullMsgNo:  text.SolLocation560_Full,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			0,             // E
			0,             // SE
			0,             // S
			0,             // SW
			0,             // W
			MineEntrance1, // NW
			0,             // Up
			0,             // Down
			0,             // In
			MineEntrance1, // Out
		},
	},
	{
		Number:     ShuttleStation3,
		BriefMsgNo: text.SolLocation561_Brief,
		FullMsgNo:  text.SolLocation561_Full,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			0,            // E
			0,            // SE
			SecurityArea, // S
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
		Number:     SecurityControl,
		BriefMsgNo: text.SolLocation562_Brief,
		FullMsgNo:  text.SolLocation562_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			Corridor42, // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			Corridor42, // Out
		},
	},
	{
		Number:     Corridor42,
		BriefMsgNo: text.SolLocation563_Brief,
		FullMsgNo:  text.SolLocation563_Full,
		Flags:      model.LfLock | model.LfShield,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			SecurityControl, // N
			0,               // NE
			Corridor43,      // E
			0,               // SE
			Tunnel5,         // S
			0,               // SW
			0,               // W
			0,               // NW
			0,               // Up
			0,               // Down
			0,               // In
			0,               // Out
		},
	},
	{
		Number:     Corridor43,
		BriefMsgNo: text.SolLocation564_Brief,
		Events:     [2]uint16{40, 0},
		FullMsgNo:  text.SolLocation564_Full,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			SecurityArea, // E
			0,            // SE
			0,            // S
			0,            // SW
			Corridor42,   // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     SecurityArea,
		BriefMsgNo: text.SolLocation565_Brief,
		FullMsgNo:  text.SolLocation565_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			ShuttleStation3, // N
			0,               // NE
			Corridor44,      // E
			0,               // SE
			0,               // S
			0,               // SW
			0,               // W
			0,               // NW
			0,               // Up
			0,               // Down
			0,               // In
			0,               // Out
		},
	},
	{
		Number:     Corridor44,
		BriefMsgNo: text.SolLocation566_Brief,
		FullMsgNo:  text.SolLocation566_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			Corridor45,   // E
			0,            // SE
			0,            // S
			0,            // SW
			SecurityArea, // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     Corridor45,
		BriefMsgNo: text.SolLocation567_Brief,
		FullMsgNo:  text.SolLocation567_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			MineShaft1, // SE
			0,          // S
			0,          // SW
			Corridor44, // W
			0,          // NW
			0,          // Up
			MineShaft1, // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     MineShaft1,
		BriefMsgNo: text.SolLocation568_Brief,
		FullMsgNo:  text.SolLocation568_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,          // N
			Tunnel3,    // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			0,          // W
			Corridor45, // NW
			Corridor45, // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Tunnel3,
		BriefMsgNo: text.SolLocation569_Brief,
		FullMsgNo:  text.SolLocation569_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			Tunnel4,    // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			MineShaft1, // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Tunnel4,
		BriefMsgNo: text.SolLocation570_Brief,
		FullMsgNo:  text.SolLocation570_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			ProcessingRoom2, // N
			0,               // NE
			MineFace1,       // E
			0,               // SE
			Tunnel3,         // S
			0,               // SW
			0,               // W
			MineFace2,       // NW
			0,               // Up
			0,               // Down
			0,               // In
			0,               // Out
		},
	},
	{
		Number:     MineFace1,
		BriefMsgNo: text.SolLocation571_Brief,
		FullMsgNo:  text.SolLocation571_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			0,       // E
			0,       // SE
			0,       // S
			0,       // SW
			Tunnel4, // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			Tunnel4, // Out
		},
	},
	{
		Number:     MineFace2,
		BriefMsgNo: text.SolLocation572_Brief,
		FullMsgNo:  text.SolLocation572_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			0,       // E
			Tunnel4, // SE
			0,       // S
			0,       // SW
			0,       // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			Tunnel4, // Out
		},
	},
	{
		Number:     ProcessingRoom2,
		BriefMsgNo: text.SolLocation573_Brief,
		FullMsgNo:  text.SolLocation573_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			0,       // E
			0,       // SE
			Tunnel4, // S
			0,       // SW
			0,       // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			Tunnel4, // Out
		},
	},
	{
		Number:     Tunnel5,
		BriefMsgNo: text.SolLocation574_Brief,
		FullMsgNo:  text.SolLocation574_Full,
		Flags:      model.LfDark | model.LfShield,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			Corridor42, // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			Tunnel6,    // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Tunnel6,
		BriefMsgNo: text.SolLocation575_Brief,
		FullMsgNo:  text.SolLocation575_Full,
		Flags:      model.LfDark | model.LfShield,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,          // N
			Tunnel5,    // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			OpenSpace6, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     OpenSpace6,
		BriefMsgNo: text.SolLocation576_Brief,
		FullMsgNo:  text.SolLocation576_Full,
		Flags:      model.LfShield,

		MovTab: [13]uint16{
			MineShaft2,      // N
			0,               // NE
			Tunnel6,         // E
			0,               // SE
			Office8,         // S
			ProcessingRoom3, // SW
			NarrowCrack1,    // W
			0,               // NW
			0,               // Up
			0,               // Down
			0,               // In
			0,               // Out
		},
	},
	{
		Number:     NarrowCrack1,
		BriefMsgNo: text.SolLocation577_Brief,
		FullMsgNo:  text.SolLocation577_Full,
		Flags:      model.LfDark | model.LfShield,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			OpenSpace6,   // E
			0,            // SE
			0,            // S
			0,            // SW
			NarrowCrack2, // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     NarrowCrack2,
		BriefMsgNo: text.SolLocation578_Brief,
		FullMsgNo:  text.SolLocation578_Full,
		Flags:      model.LfDark | model.LfShield,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			NarrowCrack1, // E
			0,            // SE
			0,            // S
			0,            // SW
			0,            // W
			LargeCave,    // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     MineShaft2,
		BriefMsgNo: text.SolLocation579_Brief,
		FullMsgNo:  text.SolLocation579_Full,
		Flags:      model.LfLock | model.LfShield,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			0,             // E
			0,             // SE
			0,             // S
			0,             // SW
			0,             // W
			SlopingTunnel, // NW
			OpenSpace6,    // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     SlopingTunnel,
		BriefMsgNo: text.SolLocation580_Brief,
		FullMsgNo:  text.SolLocation580_Full,
		Flags:      model.LfLock | model.LfDark | model.LfShield,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			AbandonedFace2, // N
			0,              // NE
			0,              // E
			MineShaft2,     // SE
			0,              // S
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
		Number:     AbandonedFace2,
		BriefMsgNo: text.SolLocation581_Brief,
		FullMsgNo:  text.SolLocation581_Full,
		Flags:      model.LfDark | model.LfLock | model.LfShield,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			0,             // E
			0,             // SE
			SlopingTunnel, // S
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
		Number:     ProcessingRoom3,
		BriefMsgNo: text.SolLocation582_Brief,
		FullMsgNo:  text.SolLocation582_Full,
		Flags:      model.LfShield,

		MovTab: [13]uint16{
			0,          // N
			OpenSpace6, // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			OpenSpace6, // Out
		},
	},
	{
		Number:     Office8,
		BriefMsgNo: text.SolLocation583_Brief,
		FullMsgNo:  text.SolLocation583_Full,
		Flags:      model.LfShield,

		MovTab: [13]uint16{
			OpenSpace6, // N
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
			OpenSpace6, // Out
		},
	},
	{
		Number:     LargeCave,
		BriefMsgNo: text.SolLocation584_Brief,
		FullMsgNo:  text.SolLocation584_Full,
		Flags:      model.LfShield,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,            // N
			SmallCave,    // NE
			0,            // E
			NarrowCrack2, // SE
			0,            // S
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
		Number:     SmallCave,
		BriefMsgNo: text.SolLocation585_Brief,
		FullMsgNo:  text.SolLocation585_Full,
		Flags:      model.LfShield,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			0,           // E
			0,           // SE
			0,           // S
			LargeCave,   // SW
			GrowingArea, // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     GrowingArea,
		BriefMsgNo: text.SolLocation586_Brief,
		Events:     [2]uint16{25, 25},
		FullMsgNo:  text.SolLocation586_Full,
		Flags:      model.LfLock | model.LfShield,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			DrugProcessing, // N
			0,              // NE
			SmallCave,      // E
			0,              // SE
			0,              // S
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
		Number:     DrugProcessing,
		BriefMsgNo: text.SolLocation587_Brief,
		FullMsgNo:  text.SolLocation587_Full,
		Flags:      model.LfShield,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,           // N
			StoreRoom2,  // NE
			0,           // E
			0,           // SE
			GrowingArea, // S
			0,           // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     StoreRoom2,
		BriefMsgNo: text.SolLocation588_Brief,
		FullMsgNo:  text.SolLocation588_Full,
		Flags:      model.LfDark | model.LfShield,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,               // N
			0,               // NE
			LivingQuarters3, // E
			0,               // SE
			0,               // S
			DrugProcessing,  // SW
			0,               // W
			0,               // NW
			0,               // Up
			0,               // Down
			0,               // In
			0,               // Out
		},
	},
	{
		Number:     LivingQuarters3,
		BriefMsgNo: text.SolLocation589_Brief,
		FullMsgNo:  text.SolLocation589_Full,
		Flags:      model.LfShield,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			StoreRoom2, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			StoreRoom2, // Out
		},
	},
	{
		Number:     ShuttleStation4,
		BriefMsgNo: text.SolLocation590_Brief,
		FullMsgNo:  text.SolLocation590_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			0,             // E
			0,             // SE
			0,             // S
			MineEntrance2, // SW
			0,             // W
			0,             // NW
			0,             // Up
			0,             // Down
			0,             // In
			MineEntrance2, // Out
		},
	},
	{
		Number:     MineEntrance2,
		BriefMsgNo: text.SolLocation591_Brief,
		FullMsgNo:  text.SolLocation591_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,               // N
			ShuttleStation4, // NE
			AbandonedFace3,  // E
			0,               // SE
			0,               // S
			MineFace3,       // SW
			0,               // W
			OldMineWorkings, // NW
			0,               // Up
			0,               // Down
			0,               // In
			0,               // Out
		},
	},
	{
		Number:     AbandonedFace3,
		BriefMsgNo: text.SolLocation592_Brief,
		FullMsgNo:  text.SolLocation592_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			0,             // E
			0,             // SE
			0,             // S
			0,             // SW
			MineEntrance2, // W
			0,             // NW
			0,             // Up
			0,             // Down
			0,             // In
			MineEntrance2, // Out
		},
	},
	{
		Number:     MineFace3,
		BriefMsgNo: text.SolLocation593_Brief,
		FullMsgNo:  text.SolLocation593_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			MineFace4,     // N
			MineEntrance2, // NE
			0,             // E
			0,             // SE
			0,             // S
			0,             // SW
			0,             // W
			0,             // NW
			0,             // Up
			0,             // Down
			0,             // In
			MineEntrance2, // Out
		},
	},
	{
		Number:     MineFace4,
		BriefMsgNo: text.SolLocation594_Brief,
		FullMsgNo:  text.SolLocation594_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			0,         // E
			0,         // SE
			MineFace3, // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			MineFace3, // Out
		},
	},
	{
		Number:     OldMineWorkings,
		BriefMsgNo: text.SolLocation595_Brief,
		FullMsgNo:  text.SolLocation595_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			0,             // E
			MineEntrance2, // SE
			0,             // S
			0,             // SW
			0,             // W
			0,             // NW
			0,             // Up
			0,             // Down
			0,             // In
			MineEntrance2, // Out
		},
	},
	{
		Number:     PoliceCell,
		BriefMsgNo: text.SolLocation596_Brief,
		FullMsgNo:  text.SolLocation596_Full,
		Flags:      model.LfLock | model.LfShield,
	},
	{
		Number:     Library2,
		BriefMsgNo: text.SolLocation597_Brief,
		FullMsgNo:  text.SolLocation597_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			Corridor46, // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			Corridor46, // Out
		},
	},
	{
		Number:     CommsShop,
		BriefMsgNo: text.SolLocation598_Brief,
		FullMsgNo:  text.SolLocation598_Full,
		Flags:      model.LfCom,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			Corridor47, // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			Corridor47, // Out
		},
	},
	{
		Number:     PoliceStation,
		BriefMsgNo: text.SolLocation599_Brief,
		FullMsgNo:  text.SolLocation599_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Corridor46, // E
			0,          // SE
			0,          // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			Corridor46, // Out
		},
	},
	{
		Number:     Corridor46,
		BriefMsgNo: text.SolLocation600_Brief,
		FullMsgNo:  text.SolLocation600_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,                // N
			Library2,         // NE
			TradingExchange3, // E
			0,                // SE
			Corridor48,       // S
			0,                // SW
			PoliceStation,    // W
			0,                // NW
			0,                // Up
			0,                // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     TradingExchange3,
		BriefMsgNo: text.SolLocation601_Brief,
		FullMsgNo:  text.SolLocation601_Full,
		Flags:      model.LfTrade,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			Corridor46, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			Corridor46, // Out
		},
	},
	{
		Number:     SlartisConstructionAndDesignWorkshop,
		BriefMsgNo: text.SolLocation602_Brief,
		FullMsgNo:  text.SolLocation602_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Corridor47, // E
			0,          // SE
			0,          // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			Corridor47, // Out
		},
	},
	{
		Number:     Corridor47,
		BriefMsgNo: text.SolLocation603_Brief,
		FullMsgNo:  text.SolLocation603_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,                                    // N
			CommsShop,                            // NE
			0,                                    // E
			0,                                    // SE
			TJunction,                            // S
			0,                                    // SW
			SlartisConstructionAndDesignWorkshop, // W
			0,                                    // NW
			0,                                    // Up
			0,                                    // Down
			0,                                    // In
			0,                                    // Out
		},
	},
	{
		Number:     Corridor48,
		BriefMsgNo: text.SolLocation604_Brief,
		FullMsgNo:  text.SolLocation604_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			Corridor46, // N
			0,          // NE
			0,          // E
			0,          // SE
			TheHub,     // S
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
		Number:     LoungeBar,
		BriefMsgNo: text.SolLocation605_Brief,
		FullMsgNo:  text.SolLocation605_Full,
		Flags:      model.LfCafe,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			Corridor52, // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			Corridor52, // Out
		},
	},
	{
		Number:     Armory,
		BriefMsgNo: text.SolLocation606_Brief,
		FullMsgNo:  text.SolLocation606_Full,
		Flags:      model.LfWeap,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			Corridor54, // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			Corridor54, // Out
		},
	},
	{
		Number:     TJunction,
		BriefMsgNo: text.SolLocation607_Brief,
		FullMsgNo:  text.SolLocation607_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			Corridor47, // N
			0,          // NE
			Corridor49, // E
			0,          // SE
			Corridor56, // S
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
		Number:     Corridor49,
		BriefMsgNo: text.SolLocation608_Brief,
		FullMsgNo:  text.SolLocation608_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Corridor50, // E
			0,          // SE
			0,          // S
			0,          // SW
			TJunction,  // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Corridor50,
		BriefMsgNo: text.SolLocation609_Brief,
		FullMsgNo:  text.SolLocation609_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			TheHub,     // E
			0,          // SE
			0,          // S
			0,          // SW
			Corridor49, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     TheHub,
		BriefMsgNo: text.SolLocation610_Brief,
		FullMsgNo:  text.SolLocation610_Full,

		MovTab: [13]uint16{
			Corridor48, // N
			0,          // NE
			Corridor51, // E
			0,          // SE
			Corridor57, // S
			0,          // SW
			Corridor50, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Corridor51,
		BriefMsgNo: text.SolLocation611_Brief,
		FullMsgNo:  text.SolLocation611_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Corridor52, // E
			0,          // SE
			0,          // S
			0,          // SW
			TheHub,     // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Corridor52,
		BriefMsgNo: text.SolLocation612_Brief,
		FullMsgNo:  text.SolLocation612_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			LoungeBar,  // N
			0,          // NE
			Corridor53, // E
			0,          // SE
			0,          // S
			0,          // SW
			Corridor51, // W
			0,          // NW
			0,          // Up
			0,          // Down
			LoungeBar,  // In
			0,          // Out
		},
	},
	{
		Number:     Corridor53,
		BriefMsgNo: text.SolLocation613_Brief,
		FullMsgNo:  text.SolLocation613_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Corridor54, // E
			0,          // SE
			ShipWright, // S
			0,          // SW
			Corridor52, // W
			0,          // NW
			0,          // Up
			0,          // Down
			ShipWright, // In
			0,          // Out
		},
	},
	{
		Number:     Corridor54,
		BriefMsgNo: text.SolLocation614_Brief,
		FullMsgNo:  text.SolLocation614_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,          // N
			Armory,     // NE
			Corridor55, // E
			0,          // SE
			0,          // S
			0,          // SW
			Corridor53, // W
			0,          // NW
			0,          // Up
			0,          // Down
			Armory,     // In
			0,          // Out
		},
	},
	{
		Number:     Corridor55,
		BriefMsgNo: text.SolLocation615_Brief,
		FullMsgNo:  text.SolLocation615_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			Tunnel7,    // SE
			0,          // S
			0,          // SW
			Corridor54, // W
			0,          // NW
			0,          // Up
			0,          // Down
			Tunnel7,    // In
			0,          // Out
		},
	},
	{
		Number:     Corridor56,
		BriefMsgNo: text.SolLocation616_Brief,
		FullMsgNo:  text.SolLocation616_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			TJunction,  // N
			0,          // NE
			0,          // E
			0,          // SE
			Corridor58, // S
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
		Number:     TravelAgency,
		BriefMsgNo: text.SolLocation617_Brief,
		FullMsgNo:  text.SolLocation617_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			Corridor58, // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			Corridor58, // Out
		},
	},
	{
		Number:     Corridor57,
		BriefMsgNo: text.SolLocation618_Brief,
		FullMsgNo:  text.SolLocation618_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			TheHub,     // N
			0,          // NE
			0,          // E
			0,          // SE
			Corridor59, // S
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
		Number:     ShipWright,
		BriefMsgNo: text.SolLocation619_Brief,
		FullMsgNo:  text.SolLocation619_Full,
		Flags:      model.LfRep,

		MovTab: [13]uint16{
			Corridor53, // N
			0,          // NE
			Workshop27, // E
			0,          // SE
			0,          // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			Corridor53, // Out
		},
	},
	{
		Number:     Workshop27,
		BriefMsgNo: text.SolLocation620_Brief,
		FullMsgNo:  text.SolLocation620_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			ShipWright, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			ShipWright, // Out
		},
	},
	{
		Number:     Tunnel7,
		BriefMsgNo: text.SolLocation621_Brief,
		FullMsgNo:  text.SolLocation621_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Tunnel13,   // E
			0,          // SE
			0,          // S
			0,          // SW
			0,          // W
			Corridor55, // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     HydroponicsPlant,
		BriefMsgNo: text.SolLocation622_Brief,
		FullMsgNo:  text.SolLocation622_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			Corridor58, // E
			0,          // SE
			0,          // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			Corridor58, // Out
		},
	},
	{
		Number:     Corridor58,
		BriefMsgNo: text.SolLocation623_Brief,
		FullMsgNo:  text.SolLocation623_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			Corridor56,       // N
			TravelAgency,     // NE
			0,                // E
			Dormitory10,      // SE
			0,                // S
			0,                // SW
			HydroponicsPlant, // W
			0,                // NW
			0,                // Up
			0,                // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     Corridor59,
		BriefMsgNo: text.SolLocation624_Brief,
		FullMsgNo:  text.SolLocation624_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			Corridor57, // N
			0,          // NE
			0,          // E
			0,          // SE
			SpacePort7, // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			SpacePort7, // In
			0,          // Out
		},
	},
	{
		Number:     Dormitory10,
		BriefMsgNo: text.SolLocation625_Brief,
		FullMsgNo:  text.SolLocation625_Full,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			0,          // W
			Corridor58, // NW
			0,          // Up
			0,          // Down
			0,          // In
			Corridor58, // Out
		},
	},
	{
		Number:     LandingBay5,
		BriefMsgNo: text.SolLocation626_Brief,
		FullMsgNo:  text.SolLocation626_Full,
		Flags:      model.LfLanding,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			SpacePort7, // E
			0,          // SE
			0,          // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			SpacePort7, // Out
		},
	},
	{
		Number:     SpacePort7,
		BriefMsgNo: text.SolLocation627_Brief,
		FullMsgNo:  text.SolLocation627_Full,

		MovTab: [13]uint16{
			Corridor59,  // N
			0,           // NE
			0,           // E
			0,           // SE
			Chandler,    // S
			0,           // SW
			LandingBay5, // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     Chandler,
		BriefMsgNo: text.SolLocation628_Brief,
		FullMsgNo:  text.SolLocation628_Full,
		Flags:      model.LfGen,

		MovTab: [13]uint16{
			SpacePort7, // N
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
			SpacePort7, // Out
		},
	},
	{
		Number:     DeadEnd1,
		BriefMsgNo: text.SolLocation629_Brief,
		FullMsgNo:  text.SolLocation629_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			0,       // E
			Tunnel8, // SE
			0,       // S
			0,       // SW
			0,       // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     SharpBend,
		BriefMsgNo: text.SolLocation630_Brief,
		FullMsgNo:  text.SolLocation630_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			0,       // E
			0,       // SE
			Cave1,   // S
			Tunnel9, // SW
			0,       // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Tunnel8,
		BriefMsgNo: text.SolLocation631_Brief,
		FullMsgNo:  text.SolLocation631_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			Tunnel9,  // E
			0,        // SE
			0,        // S
			Tunnel14, // SW
			0,        // W
			DeadEnd1, // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     Tunnel9,
		BriefMsgNo: text.SolLocation632_Brief,
		FullMsgNo:  text.SolLocation632_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,         // N
			SharpBend, // NE
			0,         // E
			0,         // SE
			0,         // S
			0,         // SW
			Tunnel8,   // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     Cave1,
		BriefMsgNo: text.SolLocation633_Brief,
		FullMsgNo:  text.SolLocation633_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			SharpBend, // N
			0,         // NE
			Tunnel10,  // E
			0,         // SE
			0,         // S
			Tunnel15,  // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			0,         // Out
		},
	},
	{
		Number:     Tunnel10,
		BriefMsgNo: text.SolLocation634_Brief,
		FullMsgNo:  text.SolLocation634_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			Tunnel11, // E
			0,        // SE
			0,        // S
			0,        // SW
			Cave1,    // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     Tunnel11,
		BriefMsgNo: text.SolLocation635_Brief,
		FullMsgNo:  text.SolLocation635_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			0,        // E
			Cave2,    // SE
			0,        // S
			0,        // SW
			Tunnel10, // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     Tunnel12,
		BriefMsgNo: text.SolLocation636_Brief,
		FullMsgNo:  text.SolLocation636_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			0,             // E
			GrizzlesLair2, // SE
			Cave2,         // S
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
		Number:     Tunnel13,
		BriefMsgNo: text.SolLocation637_Brief,
		FullMsgNo:  text.SolLocation637_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			Tunnel14, // E
			0,        // SE
			0,        // S
			0,        // SW
			Tunnel7,  // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     Tunnel14,
		BriefMsgNo: text.SolLocation638_Brief,
		FullMsgNo:  text.SolLocation638_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,        // N
			Tunnel8,  // NE
			0,        // E
			Cave3,    // SE
			0,        // S
			0,        // SW
			Tunnel13, // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     GrizzlesLair1,
		BriefMsgNo: text.SolLocation639_Brief,
		FullMsgNo:  text.SolLocation639_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			Tunnel15, // E
			0,        // SE
			0,        // S
			0,        // SW
			0,        // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			Tunnel15, // Out
		},
	},
	{
		Number:     Tunnel15,
		BriefMsgNo: text.SolLocation640_Brief,
		FullMsgNo:  text.SolLocation640_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,             // N
			Cave1,         // NE
			0,             // E
			0,             // SE
			0,             // S
			0,             // SW
			GrizzlesLair1, // W
			0,             // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     Tunnel16,
		BriefMsgNo: text.SolLocation641_Brief,
		FullMsgNo:  text.SolLocation641_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			0,        // E
			0,        // SE
			Tunnel18, // S
			Tunnel17, // SW
			0,        // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     Cave2,
		BriefMsgNo: text.SolLocation642_Brief,
		FullMsgNo:  text.SolLocation642_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			Tunnel12, // N
			0,        // NE
			0,        // E
			0,        // SE
			0,        // S
			0,        // SW
			0,        // W
			Tunnel11, // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     GrizzlesLair2,
		BriefMsgNo: text.SolLocation643_Brief,
		FullMsgNo:  text.SolLocation643_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			0,        // E
			0,        // SE
			0,        // S
			0,        // SW
			0,        // W
			Tunnel12, // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     Cave3,
		BriefMsgNo: text.SolLocation644_Brief,
		FullMsgNo:  text.SolLocation644_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			Tunnel17,     // E
			0,            // SE
			NarrowCrack3, // S
			0,            // SW
			0,            // W
			Tunnel14,     // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     Tunnel17,
		BriefMsgNo: text.SolLocation645_Brief,
		FullMsgNo:  text.SolLocation645_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,        // N
			Tunnel16, // NE
			0,        // E
			0,        // SE
			0,        // S
			0,        // SW
			Cave3,    // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     Tunnel18,
		BriefMsgNo: text.SolLocation646_Brief,
		FullMsgNo:  text.SolLocation646_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			Tunnel16, // N
			0,        // NE
			Tunnel19, // E
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
		Number:     Tunnel19,
		BriefMsgNo: text.SolLocation647_Brief,
		FullMsgNo:  text.SolLocation647_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			0,        // E
			Tunnel21, // SE
			0,        // S
			Tunnel20, // SW
			Tunnel18, // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     NarrowCrack3,
		BriefMsgNo: text.SolLocation648_Brief,
		FullMsgNo:  text.SolLocation648_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			Cave3, // N
			0,     // NE
			Cave4, // E
			0,     // SE
			0,     // S
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
		Number:     Cave4,
		BriefMsgNo: text.SolLocation649_Brief,
		FullMsgNo:  text.SolLocation649_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			Tunnel20,     // E
			0,            // SE
			DeadEnd2,     // S
			0,            // SW
			NarrowCrack3, // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     Tunnel20,
		BriefMsgNo: text.SolLocation650_Brief,
		FullMsgNo:  text.SolLocation650_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,        // N
			Tunnel19, // NE
			0,        // E
			0,        // SE
			0,        // S
			0,        // SW
			Cave4,    // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     GrizzlesLair3,
		BriefMsgNo: text.SolLocation651_Brief,
		FullMsgNo:  text.SolLocation651_Full,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,     // N
			0,     // NE
			0,     // E
			Cave5, // SE
			0,     // S
			0,     // SW
			0,     // W
			0,     // NW
			0,     // Up
			0,     // Down
			0,     // In
			Cave5, // Out
		},
	},
	{
		Number:     Tunnel21,
		BriefMsgNo: text.SolLocation652_Brief,
		FullMsgNo:  text.SolLocation652_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			0,        // E
			Cave6,    // SE
			0,        // S
			0,        // SW
			0,        // W
			Tunnel19, // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     Tunnel22,
		BriefMsgNo: text.SolLocation653_Brief,
		FullMsgNo:  text.SolLocation653_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			Tunnel23, // E
			0,        // SE
			Cave6,    // S
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
		Number:     Tunnel23,
		BriefMsgNo: text.SolLocation654_Brief,
		FullMsgNo:  text.SolLocation654_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			Tunnel24, // E
			0,        // SE
			0,        // S
			0,        // SW
			Tunnel22, // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     Tunnel24,
		BriefMsgNo: text.SolLocation655_Brief,
		FullMsgNo:  text.SolLocation655_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			0,             // E
			0,             // SE
			GrizzlesLair4, // S
			0,             // SW
			Tunnel23,      // W
			0,             // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     DeadEnd2,
		BriefMsgNo: text.SolLocation656_Brief,
		FullMsgNo:  text.SolLocation656_Full,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			Cave4, // N
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
			0,     // Out
		},
	},
	{
		Number:     Cave5,
		FullMsgNo:  text.SolLocation657_Full,
		BriefMsgNo: text.SolLocation657_Brief,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			Cave6,         // E
			0,             // SE
			Tunnel26,      // S
			0,             // SW
			0,             // W
			GrizzlesLair3, // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     Cave6,
		FullMsgNo:  text.SolLocation658_Full,
		BriefMsgNo: text.SolLocation658_Brief,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			Tunnel22, // N
			0,        // NE
			Tunnel25, // E
			0,        // SE
			0,        // S
			0,        // SW
			Cave5,    // W
			Tunnel21, // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     Tunnel25,
		FullMsgNo:  text.SolLocation659_Full,
		BriefMsgNo: text.SolLocation659_Brief,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			0,             // E
			0,             // SE
			GrizzlesLair6, // S
			0,             // SW
			Cave6,         // W
			0,             // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     GrizzlesLair4,
		FullMsgNo:  text.SolLocation660_Full,
		BriefMsgNo: text.SolLocation660_Brief,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			Tunnel24, // N
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
			Tunnel24, // Out
		},
	},
	{
		Number:     GrizzlesLair5,
		FullMsgNo:  text.SolLocation661_Full,
		BriefMsgNo: text.SolLocation661_Brief,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,     // N
			0,     // NE
			0,     // E
			Cave8, // SE
			0,     // S
			0,     // SW
			0,     // W
			0,     // NW
			0,     // Up
			0,     // Down
			0,     // In
			Cave8, // Out
		},
	},
	{
		Number:     Tunnel26,
		FullMsgNo:  text.SolLocation662_Full,
		BriefMsgNo: text.SolLocation662_Brief,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			Cave5, // N
			0,     // NE
			0,     // E
			0,     // SE
			Shaft, // S
			0,     // SW
			0,     // W
			0,     // NW
			0,     // Up
			Shaft, // Down
			0,     // In
			0,     // Out
		},
	},
	{
		Number:     GrizzlesLair6,
		FullMsgNo:  text.SolLocation663_Full,
		BriefMsgNo: text.SolLocation663_Brief,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			Tunnel25, // N
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
			Tunnel25, // Out
		},
	},
	{
		Number:     Cave7,
		FullMsgNo:  text.SolLocation664_Full,
		BriefMsgNo: text.SolLocation664_Brief,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			0,        // E
			Cave9,    // SE
			0,        // S
			DeadEnd3, // SW
			0,        // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     Cave8,
		FullMsgNo:  text.SolLocation665_Full,
		BriefMsgNo: text.SolLocation665_Brief,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,             // N
			0,             // NE
			Tunnel27,      // E
			0,             // SE
			0,             // S
			Cave9,         // SW
			0,             // W
			GrizzlesLair5, // NW
			0,             // Up
			0,             // Down
			0,             // In
			0,             // Out
		},
	},
	{
		Number:     Tunnel27,
		FullMsgNo:  text.SolLocation666_Full,
		BriefMsgNo: text.SolLocation666_Brief,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			Tunnel28, // E
			0,        // SE
			0,        // S
			0,        // SW
			Cave8,    // W
			0,        // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     Tunnel28,
		FullMsgNo:  text.SolLocation667_Full,
		BriefMsgNo: text.SolLocation667_Brief,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			Shaft,    // E
			0,        // SE
			0,        // S
			0,        // SW
			Tunnel27, // W
			0,        // NW
			Shaft,    // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     Shaft,
		FullMsgNo:  text.SolLocation668_Full,
		BriefMsgNo: text.SolLocation668_Brief,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			Tunnel26, // N
			0,        // NE
			0,        // E
			0,        // SE
			0,        // S
			0,        // SW
			Tunnel28, // W
			0,        // NW
			Tunnel26, // Up
			Tunnel28, // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     DeadEnd3,
		FullMsgNo:  text.SolLocation669_Full,
		BriefMsgNo: text.SolLocation669_Brief,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,     // N
			Cave7, // NE
			0,     // E
			0,     // SE
			0,     // S
			0,     // SW
			0,     // W
			0,     // NW
			0,     // Up
			0,     // Down
			0,     // In
			Cave7, // Out
		},
	},
	{
		Number:     Cave9,
		FullMsgNo:  text.SolLocation670_Full,
		BriefMsgNo: text.SolLocation670_Brief,
		Flags:      model.LfDark,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,       // N
			Cave8,   // NE
			0,       // E
			Chamber, // SE
			0,       // S
			0,       // SW
			0,       // W
			Cave7,   // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Chamber,
		FullMsgNo:  text.SolLocation671_Full,
		BriefMsgNo: text.SolLocation671_Brief,
		SysLoc:     text.NO_MOVEMENT_26,

		MovTab: [13]uint16{
			0,     // N
			0,     // NE
			0,     // E
			0,     // SE
			0,     // S
			0,     // SW
			0,     // W
			Cave9, // NW
			0,     // Up
			0,     // Down
			0,     // In
			Cave9, // Out
		},
	},
	{
		Number:     CityRoad6,
		BriefMsgNo: text.SolLocation672_Brief,
		FullMsgNo:  text.SolLocation672_Full,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			0,           // E
			0,           // SE
			WideAvenue1, // S
			0,           // SW
			0,           // W
			CityRoad1,   // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     WideAvenue1,
		FullMsgNo:  text.SolLocation673_Full,
		BriefMsgNo: text.SolLocation673_Brief,

		MovTab: [13]uint16{
			CityRoad6,   // N
			0,           // NE
			0,           // E
			0,           // SE
			WideAvenue2, // S
			0,           // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     WideAvenue2,
		FullMsgNo:  text.SolLocation674_Full,
		BriefMsgNo: text.SolLocation674_Brief,

		MovTab: [13]uint16{
			WideAvenue1,  // N
			0,            // NE
			EntranceHall, // E
			0,            // SE
			WideRoad,     // S
			0,            // SW
			0,            // W
			0,            // NW
			0,            // Up
			0,            // Down
			EntranceHall, // In
			0,            // Out
		},
	},
	{
		Number:     WideRoad,
		FullMsgNo:  text.SolLocation675_Full,
		BriefMsgNo: text.SolLocation675_Brief,

		MovTab: [13]uint16{
			WideAvenue2,      // N
			0,                // NE
			0,                // E
			0,                // SE
			CityStreet9,      // S
			0,                // SW
			HagarsMusicStore, // W
			0,                // NW
			0,                // Up
			0,                // Down
			HagarsMusicStore, // In
			0,                // Out
		},
	},
	{
		Number:     CityStreet9,
		FullMsgNo:  text.SolLocation676_Full,
		BriefMsgNo: text.SolLocation676_Brief,

		MovTab: [13]uint16{
			WideRoad,     // N
			0,            // NE
			CityStreet10, // E
			0,            // SE
			0,            // S
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
		Number:     CityStreet10,
		FullMsgNo:  text.SolLocation677_Full,
		BriefMsgNo: text.SolLocation677_Brief,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			CityStreet11, // E
			0,            // SE
			0,            // S
			0,            // SW
			CityStreet9,  // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     CityStreet11,
		FullMsgNo:  text.SolLocation678_Full,
		BriefMsgNo: text.SolLocation678_Brief,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			0,            // E
			CityStreet12, // SE
			0,            // S
			0,            // SW
			CityStreet10, // W
			0,            // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     CityStreet12,
		FullMsgNo:  text.SolLocation679_Full,
		BriefMsgNo: text.SolLocation679_Brief,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			CityRoad2,    // E
			0,            // SE
			0,            // S
			0,            // SW
			0,            // W
			CityStreet11, // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     EntranceHall,
		FullMsgNo:  text.SolLocation680_Full,
		BriefMsgNo: text.SolLocation680_Brief,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			GuardsRoom,  // E
			0,           // SE
			0,           // S
			0,           // SW
			WideAvenue2, // W
			0,           // NW
			0,           // Up
			0,           // Down
			GuardsRoom,  // In
			WideAvenue2, // Out
		},
	},
	{
		Number:     GuardsRoom,
		FullMsgNo:  text.SolLocation681_Full,
		BriefMsgNo: text.SolLocation681_Brief,

		MovTab: [13]uint16{
			TypingPool2,    // N
			Stairs3,        // NE
			Storeroom3,     // E
			0,              // SE
			0,              // S
			PublicCounter2, // SW
			EntranceHall,   // W
			0,              // NW
			0,              // Up
			0,              // Down
			0,              // In
			EntranceHall,   // Out
		},
	},
	{
		Number:     TypingPool1,
		FullMsgNo:  text.SolLocation682_Full,
		BriefMsgNo: text.SolLocation682_Brief,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			TypingPool2, // E
			0,           // SE
			0,           // S
			0,           // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			0,           // Out
		},
	},
	{
		Number:     TypingPool2,
		FullMsgNo:  text.SolLocation683_Full,
		BriefMsgNo: text.SolLocation683_Brief,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			0,           // E
			0,           // SE
			GuardsRoom,  // S
			0,           // SW
			TypingPool1, // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			GuardsRoom,  // Out
		},
	},
	{
		Number:     Stairs3,
		FullMsgNo:  text.SolLocation684_Full,
		BriefMsgNo: text.SolLocation684_Brief,

		MovTab: [13]uint16{
			0,
			0,
			0,
			0,
			0,
			GuardsRoom,
			0,
			0,
			TopOfStairs,
			0,
			0,
			0,
			0,
		},
	},
	{
		Number:     Storeroom3,
		FullMsgNo:  text.SolLocation685_Full,
		BriefMsgNo: text.SolLocation685_Brief,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			GuardsRoom, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			GuardsRoom, // Out
		},
	},
	{
		Number:     PublicCounter2,
		FullMsgNo:  text.SolLocation686_Full,
		BriefMsgNo: text.SolLocation686_Brief,

		MovTab: [13]uint16{
			0,          // N
			GuardsRoom, // NE
			Office9,    // E
			0,          // SE
			0,          // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			GuardsRoom, // Out
		},
	},
	{
		Number:     Office9,
		FullMsgNo:  text.SolLocation687_Full,
		BriefMsgNo: text.SolLocation687_Brief,

		MovTab: [13]uint16{
			0,              // N
			0,              // NE
			0,              // E
			SmartOffice,    // SE
			0,              // S
			0,              // SW
			PublicCounter2, // W
			0,              // NW
			0,              // Up
			0,              // Down
			0,              // In
			0,              // Out
		},
	},
	{
		Number:     SmartOffice,
		FullMsgNo:  text.SolLocation688_Full,
		BriefMsgNo: text.SolLocation688_Brief,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			0,       // E
			0,       // SE
			0,       // S
			0,       // SW
			0,       // W
			Office9, // NW
			0,       // Up
			0,       // Down
			0,       // In
			Office9, // Out
		},
	},
	{
		Number:     BareRoom3,
		FullMsgNo:  text.SolLocation689_Full,
		BriefMsgNo: text.SolLocation689_Brief,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			Storeroom3, // N
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
			Storeroom3, // Out
		},
	},
	{
		Number:     BareRoom4,
		FullMsgNo:  text.SolLocation690_Full,
		BriefMsgNo: text.SolLocation690_Brief,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			GuardRoom4, // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			GuardRoom4, // Out
		},
	},
	{
		Number:     GuardRoom4,
		FullMsgNo:  text.SolLocation691_Full,
		BriefMsgNo: text.SolLocation691_Brief,
		Events:     [2]uint16{20, 0},
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			BareRoom4,    // N
			0,            // NE
			SmallOffice2, // E
			0,            // SE
			0,            // S
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
		Number:     SmallOffice2,
		FullMsgNo:  text.SolLocation692_Full,
		BriefMsgNo: text.SolLocation692_Brief,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			Lift8,      // N
			0,          // NE
			0,          // E
			0,          // SE
			0,          // S
			0,          // SW
			GuardRoom4, // W
			0,          // NW
			0,          // Up
			0,          // Down
			0,          // In
			0,          // Out
		},
	},
	{
		Number:     Lift8,
		FullMsgNo:  text.SolLocation693_Full,
		BriefMsgNo: text.SolLocation693_Brief,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0,            // N
			0,            // NE
			0,            // E
			0,            // SE
			SmallOffice2, // S
			0,            // SW
			0,            // W
			0,            // NW
			0,            // Up
			Lift9,        // Down
			0,            // In
			SmallOffice2, // Out
		},
	},
	{
		Number:     Lift9,
		FullMsgNo:  text.SolLocation694_Full,
		BriefMsgNo: text.SolLocation694_Brief,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			GuardRoom5, // SE
			0,          // S
			0,          // SW
			0,          // W
			0,          // NW
			Lift8,      // Up
			0,          // Down
			0,          // In
			GuardRoom5, // Out
		},
	},
	{
		Number:     GuardRoom5,
		FullMsgNo:  text.SolLocation695_Full,
		BriefMsgNo: text.SolLocation695_Brief,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0,        // N
			0,        // NE
			Terminal, // E
			0,        // SE
			0,        // S
			0,        // SW
			0,        // W
			Lift9,    // NW
			0,        // Up
			0,        // Down
			0,        // In
			0,        // Out
		},
	},
	{
		Number:     MeetingPoint,
		FullMsgNo:  text.SolLocation696_Full,
		BriefMsgNo: text.SolLocation696_Brief,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0,         // N
			0,         // NE
			Terminus1, // E
			0,         // SE
			0,         // S
			0,         // SW
			0,         // W
			0,         // NW
			0,         // Up
			0,         // Down
			0,         // In
			Terminus1, // Out
		},
	},
	{
		Number:     DieselsBoudoir,
		FullMsgNo:  text.SolLocation697_Full,
		BriefMsgNo: text.SolLocation697_Brief,
		Flags:      model.LfLock | model.LfShield,

		MovTab: [13]uint16{
			0,          // N
			0,          // NE
			0,          // E
			0,          // SE
			ChezDiesel, // S
			0,          // SW
			0,          // W
			0,          // NW
			0,          // Up
			ChezDiesel, // Down
			0,          // In
			ChezDiesel, // Out
		},
	},
	{
		Number:     ThroughTheLookingGlass,
		FullMsgNo:  text.SolLocation698_Full,
		BriefMsgNo: text.SolLocation698_Brief,
		Events:     [2]uint16{32, 0},
		Flags:      model.LfHidden | model.LfLock | model.LfShield,
	},
}

var Objects = [90]core.Object{
	{
		DescMessageNo: text.Barbells_Desc,
		GetEvent:      15,
		MaxLoc:        HighGravityGym,
		MinLoc:        HighGravityGym,
		Name:          "barbells",
		Number:        ObBarbells, // 860,
		ScanMessageNo: text.Barbells_Scan,
		Sex:           model.SexNeuter,
		Weight:        math.MaxUint16,
	},
	{
		DescMessageNo: text.Beaker_Desc,
		GetEvent:      3,
		MaxLoc:        Laboratory,
		MinLoc:        Laboratory,
		Name:          "beaker",
		Number:        ObBeaker, // 801,
		ScanMessageNo: text.Beaker_Scan,
		Sex:           'n',
		Synonyms:      []string{"acid"},
		Value:         10,
		Weight:        1,
	},
	{
		DescMessageNo: text.Boy_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    10,
		MaxLoc:        Landing2,
		MinLoc:        Hallway1,
		Name:          "boy",
		Number:        ObBoy, // 1024,
		ScanMessageNo: text.Boy_Scan,
		Sex:           'm',
	},
	{
		DescMessageNo: text.Butler_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    75,
		MaxLoc:        Hallway2,
		MinLoc:        Hallway1,
		Name:          "butler",
		Number:        ObButler, // 1023,
		ScanMessageNo: text.Butler_Scan,
		Sex:           'm',
	},
	{
		DescMessageNo: text.Calendar_Desc,
		MaxLoc:        Corridor46,
		MinLoc:        Library2,
		Name:          "calendar",
		Number:        ObCalendar, // 830,
		ScanMessageNo: text.Calendar_Scan,
		Sex:           'n',
		Value:         38,
		Weight:        1,
	},
	{
		DescMessageNo: text.Chart_Desc,
		MaxLoc:        EastHall,
		MinLoc:        ProductionOffice,
		Name:          "chart",
		Number:        ObChart, // 803,
		ScanMessageNo: text.Chart_Scan,
		Sex:           'n',
		Value:         120,
		Weight:        1,
	},
	{
		DescMessageNo: text.Cleaner_Desc,
		Flags:         model.OfAnimate | model.OfCleaner | model.OfStoic,
		MaxCounter:    2,
		MaxLoc:        DieselsBoudoir,
		MinLoc:        Lift1,
		Name:          "cleaner",
		Number:        ObCleaner, // 1020,
		ScanMessageNo: text.Cleaner_Scan,
		Sex:           'n',
	},
	{
		DescMessageNo: text.Coat_Desc,
		MaxLoc:        Corridor38,
		MinLoc:        Cloakroom,
		Name:          "coat",
		Number:        ObCoat, // 823,
		ScanMessageNo: text.Coat_Scan,
		Sex:           'n',
		Value:         65,
		Weight:        4,
	},
	{
		DescMessageNo: text.Cookie_Desc,
		Flags:         model.OfEdible,
		MaxLoc:        Landing2,
		MinLoc:        Kitchen,
		Name:          "cookie",
		Number:        ObCookie, // 810,
		ScanMessageNo: text.Cookie_Scan,
		Sex:           'n',
		Value:         15,
		Weight:        1,
	},
	{
		DescMessageNo: text.Curator_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    200,
		MaxLoc:        MuseumLobby,
		MinLoc:        Museum1,
		Name:          "curator",
		Number:        ObCurator, // 1035,
		PrefObject:    ObVandier, // 819,
		ScanMessageNo: text.Curator_Scan,
		Sex:           'm',
	},
	{
		DescMessageNo: text.Diamond_Desc,
		GetEvent:      12,
		MaxLoc:        ProcessingRoom2,
		MinLoc:        MineFace1,
		Name:          "diamond",
		Number:        ObDiamond, // 822,
		ScanMessageNo: text.Diamond_Scan,
		Sex:           'n',
		Value:         245,
		Weight:        2,
	},
	{
		AttackPercent: 85,
		DescMessageNo: text.Dog_Desc,
		Flags:         model.OfAnimate | model.OfStoic,
		MaxCounter:    -1,
		MaxLoc:        Backyard,
		MinLoc:        Backyard,
		Name:          "German Shepherd",
		Number:        ObDog, // 1022,
		ScanMessageNo: text.Dog_Scan,
		Sex:           'n',
		Synonyms:      []string{"dog"},
	},
	{
		DescMessageNo: text.Drugs_Desc,
		Flags:         model.OfHidden,
		MaxLoc:        StoreRoom2,
		MinLoc:        StoreRoom2,
		Name:          "drugs",
		Number:        ObDrugs, // 806,
		ScanMessageNo: text.Drugs_Scan,
		Sex:           'n',
		Value:         1500,
		Weight:        2,
	},
	{
		DescMessageNo: text.Flower_Desc,
		Flags:         model.OfEdible,
		MaxLoc:        Church2,
		MinLoc:        Church1,
		Name:          "flower",
		Number:        ObFlower, // 820,
		ScanMessageNo: text.Flower_Scan,
		Sex:           'n',
		Value:         35,
		Weight:        1,
	},
	{
		DescMessageNo: text.Foreman_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    55,
		MaxLoc:        ForemansOffice,
		MinLoc:        Corridor17,
		Name:          "foreman",
		Number:        ObForeman, // 1005,
		PrefObject:    ObChart,   // 803,
		ScanMessageNo: text.Foreman_Scan,
		Sex:           'm',
	},
	{
		AttackPercent: 20,
		DescMessageNo: text.Godfather_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    -1,
		MaxLoc:        Office2,
		MinLoc:        Office2,
		Name:          "Godfather",
		Number:        ObGodfather, // 1012,
		PrefObject:    ObDrugs,     // 806,
		ScanMessageNo: text.Godfather_Scan,
		Sex:           'm',
	},
	{
		DescMessageNo: text.Godot_Desc,
		Flags:         model.OfAnimate | model.OfNoThe,
		MaxCounter:    -1,
		MaxLoc:        SnackBar,
		MinLoc:        SnackBar,
		Name:          "Godot",
		Number:        ObGodot,    // 1007,
		PrefObject:    ObCalendar, // 830,
		ScanMessageNo: text.Godot_Scan,
		Sex:           'm',
	},
	{
		DescMessageNo: text.Hobo_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    45,
		MaxLoc:        CityRoad1,
		MinLoc:        PerimeterRoad1,
		Name:          "hobo",
		Number:        ObHobo, // 1026,
		PrefObject:    ObCoat, // 823,
		ScanMessageNo: text.Hobo_Scan,
		Sex:           'm',
	},
	{
		DescMessageNo: text.Kalindra_Desc,
		MaxLoc:        PenthouseSuite,
		MinLoc:        PenthouseSuite,
		Name:          "kalindra",
		Number:        ObKalindra, // 818,
		ScanMessageNo: text.Kalindra_Scan,
		Sex:           'n',
		Value:         200,
		Weight:        15,
	},
	{
		DescMessageNo: text.Klystron_Desc,
		MaxLoc:        TradingExchange2,
		MinLoc:        PowerPlant2,
		Name:          "klystron",
		Number:        ObKlystron, // 828,
		ScanMessageNo: text.Klystron_Scan,
		Sex:           'n',
		Value:         215,
		Weight:        12,
	},
	{
		DescMessageNo: text.Manifesto_Desc,
		MaxLoc:        Dormitory10,
		MinLoc:        TravelAgency,
		Name:          "manifesto",
		Number:        ObManifesto, // 813,
		ScanMessageNo: text.Manifesto_Scan,
		Sex:           'n',
		Value:         135,
		Weight:        1,
	},
	{
		DescMessageNo: text.Map_Desc,
		MaxLoc:        Dormitory5,
		MinLoc:        Dormitory1,
		Name:          "map",
		Number:        ObMap, // 815,
		ScanMessageNo: text.Map_Scan,
		Sex:           'n',
		Value:         115,
		Weight:        1,
	},
	{
		DescMessageNo: text.Marillion_Desc,
		Flags:         model.OfAnimate,
		GetEvent:      2,
		MaxCounter:    5,
		MaxLoc:        EastHall,
		MinLoc:        LinkBetweenDomes1,
		Name:          "marillion",
		Number:        ObMarillion, // 1001,
		ScanMessageNo: text.Marillion_Scan,
		Sex:           'n',
	},
	{
		AttackPercent: 30,
		DescMessageNo: text.Mario_Desc,
		Flags:         model.OfAnimate | model.OfNoThe,
		MaxCounter:    -1,
		MaxLoc:        Marios,
		MinLoc:        Marios,
		Name:          "Mario",
		Number:        ObMario, // 1015,
		ScanMessageNo: text.Mario_Scan,
		Sex:           'm',
	},
	{
		AttackPercent: 80,
		DescMessageNo: text.Marsrat_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    64,
		MaxLoc:        CaveIn,
		MinLoc:        Alleyway1,
		Name:          "marsrat",
		Number:        ObMarsrat, // 1019,
		ScanMessageNo: text.Marsrat_Scan,
		Sex:           'n',
		Synonyms:      []string{"rat"},
	},
	{
		DescMessageNo: text.Novel_Desc,
		MaxLoc:        WaitingRoom,
		MinLoc:        WaitingRoom,
		Name:          "novel",
		Number:        ObNovel, // 829,
		ScanMessageNo: text.Novel_Scan,
		Sex:           'n',
		Synonyms:      []string{"book"},
		Value:         65,
		Weight:        1,
	},
	{
		DescMessageNo: text.Nurse_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    50,
		MaxLoc:        Corridor26,
		MinLoc:        Foyer,
		Name:          "nurse",
		Number:        ObNurse,   // 1031,
		PrefObject:    ObSyringe, // 812,
		ScanMessageNo: text.Nurse_Scan,
		Sex:           'f',
	},
	{
		DescMessageNo: text.Official_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    -1,
		MaxLoc:        PublicCounter2,
		MinLoc:        PublicCounter2,
		Name:          "official",
		Number:        ObOfficial, // 1009,
		ScanMessageNo: text.Official_Scan,
		Sex:           'm',
	},
	{
		DescMessageNo: text.Opal_Desc,
		MaxLoc:        TreasureRoom1,
		MinLoc:        TreasureRoom1,
		Name:          "opal",
		Number:        ObOpal, // 826,
		ScanMessageNo: text.Opal_Scan,
		Sex:           'n',
		Value:         1000,
		Weight:        4,
	},
	{
		DescMessageNo: text.Oscilloscope_Desc,
		MaxLoc:        SmallOffice1,
		MinLoc:        ControlRoom1,
		Name:          "oscilloscope",
		Number:        ObOscilloscope, // 814,
		ScanMessageNo: text.Oscilloscope_Scan,
		Sex:           'n',
		Synonyms:      []string{"scope"},
		Value:         210,
		Weight:        8,
	},
	{
		DescMessageNo: text.Pearls_Desc,
		MaxLoc:        Cave9,
		MinLoc:        GrizzlesLair4,
		Name:          "pearls",
		Number:        ObPearls, // 824,
		ScanMessageNo: text.Pearls_Scan,
		Sex:           'n',
		Value:         135,
		Weight:        1,
	},
	{
		AttackPercent: 90,
		DescMessageNo: text.Pegasus_Desc,
		Flags:         model.OfAnimate | model.OfNoThe | model.OfShip,
		MaxCounter:    3,
		MaxLoc:        SolarOrbit4,
		MinLoc:        InterplanetarySpace1,
		Name:          "Pegasus",
		Number:        ObPegasus,   // 1063,
		PrefObject:    ObSugarLump, // 858,
		ScanMessageNo: text.Pegasus_Scan,
		Sex:           'm',
		Value:         2000000,

		ShipGuns: [4]model.OldSGuns{
			core.SHIP_GUN_QUAD_LASER,
			core.SHIP_GUN_QUAD_LASER,
			core.SHIP_GUN_QUAD_LASER,
			core.SHIP_GUN_TWIN_LASER,
		},

		ShipKit: core.Equipment{
			Computer: 6,
			Engine:   40,
			Fuel:     400,
			Hold:     400,
			Hull:     270,
			Shield:   10,
			Tonnage:  1000,
		},
	},
	{
		DescMessageNo: text.Photograph_Desc,
		MaxLoc:        HotelLounge,
		MinLoc:        ChezDiesel,
		Name:          "photograph",
		Number:        ObPhotograph, // 816,
		ScanMessageNo: text.Photograph_Scan,
		Sex:           'n',
		Synonyms:      []string{"photo"},
		Value:         320,
		Weight:        1,
	},
	{
		DescMessageNo: text.Piano_Desc,
		Flags:         model.OfMusic,
		MaxLoc:        ChezDiesel,
		MinLoc:        ChezDiesel,
		Name:          "piano",
		Number:        ObPiano, // 1082,
		ScanMessageNo: text.Piano_Scan,
		Sex:           'n',
		Weight:        math.MaxUint16,
	},
	{
		DescMessageNo: text.PinkFloyd_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    25,
		MaxLoc:        Corridor16,
		MinLoc:        Corridor7,
		Name:          "pink floyd",
		Number:        ObPinkFloyd, // 1011,
		PrefObject:    ObKalindra,  // 818,
		ScanMessageNo: text.PinkFloyd_Scan,
		Sex:           'n',
		Synonyms:      []string{"floyd"},
	},
	{
		DescMessageNo: text.Politician_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    -1,
		MaxLoc:        PartyHQ,
		MinLoc:        PartyHQ,
		Name:          "politician",
		Number:        ObPolitician, // 1034,
		PrefObject:    ObManifesto,  // 813,
		ScanMessageNo: text.Politician_Scan,
		Sex:           'm',
	},
	{
		DescMessageNo: text.Porter_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    40,
		MaxLoc:        Office3,
		MinLoc:        Foyer,
		Name:          "porter",
		Number:        ObPorter, // 1033,
		ScanMessageNo: text.Porter_Scan,
		Sex:           'm',
	},
	{
		DescMessageNo: text.Rabbit_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    10,
		MaxLoc:        MazeOfAlleys12,
		MinLoc:        MazeOfAlleys1,
		Name:          "white rabbit",
		Number:        ObRabbit, // 1021,
		ScanMessageNo: text.Rabbit_Scan,
		Sex:           'n',
		Synonyms:      []string{"rabbit"},
	},
	{
		DescMessageNo: text.Reality_Desc,
		Flags:         model.OfAnimate | model.OfNoThe,
		GetEvent:      10,
		MaxCounter:    5,
		MaxLoc:        MazeOfAlleys12,
		MinLoc:        Alleyway1,
		Name:          "Reality",
		Number:        ObReality, // 1049,
		ScanMessageNo: text.Reality_Scan,
		Sex:           'n',
	},
	{
		DescMessageNo: text.Receptionist_Desc,
		Flags:         model.OfAnimate,
		GetEvent:      6,
		MaxCounter:    -1,
		MaxLoc:        Reception,
		MinLoc:        Reception,
		Name:          "receptionist",
		Number:        ObReceptionist, // 1003,
		PrefObject:    ObNovel,        // 829,
		ScanMessageNo: text.Receptionist_Scan,
		Sex:           'f',
	},
	{
		DescMessageNo: text.Sandwich_Desc,
		Flags:         model.OfEdible,
		MaxLoc:        SlartisConstructionAndDesignWorkshop,
		MinLoc:        SlartisConstructionAndDesignWorkshop,
		Name:          "sandwich",
		Number:        ObSandwich, // 827,
		ScanMessageNo: text.Sandwich_Scan,
		Sex:           'n',
		Value:         55,
		Weight:        1,
	},
	{
		DescMessageNo: text.Sargeur_Desc,
		MaxLoc:        SlabRoom,
		MinLoc:        SlabRoom,
		Name:          "sargeur",
		Number:        ObSargeur, // 825,
		ScanMessageNo: text.Sargeur_Scan,
		Sex:           'n',
		Value:         255,
		Weight:        9,
	},
	{
		DescMessageNo: text.Share_Desc,
		Events:        [4]uint16{0, 53, 0, 0},
		MaxLoc:        Office1,
		MinLoc:        Office1,
		Name:          "share",
		Number:        ObShare, // 821,
		ScanMessageNo: text.Share_Scan,
		Sex:           'n',
		Value:         145,
		Weight:        2,
	},
	{
		DescMessageNo: text.Spanner_Desc,
		MaxLoc:        StoreRoom1,
		MinLoc:        StoreRoom1,
		Name:          "monkey wrench",
		Number:        ObSpanner, // 807,
		ScanMessageNo: text.Spanner_Scan,
		Sex:           'n',
		Synonyms:      []string{"spanner", "wrench"},
		Value:         75,
		Weight:        16,
	},
	{
		DescMessageNo: text.Storeman_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    100,
		MaxLoc:        LargeWarehouse3,
		MinLoc:        LargeWarehouse1,
		Name:          "storeman",
		Number:        ObStoreman, // 1017,
		PrefObject:    ObKlystron, // 828,
		ScanMessageNo: text.Storeman_Scan,
		Sex:           'm',
	},
	{
		DescMessageNo: text.Syringe_Desc,
		MaxLoc:        SmallWarehouse,
		MinLoc:        Gunsmiths,
		Name:          "syringe",
		Number:        ObSyringe, // 812,
		ScanMessageNo: text.Syringe_Scan,
		Sex:           'n',
		Value:         120,
		Weight:        1,
	},
	{
		DescMessageNo: text.TaxInspector_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    -1,
		MaxLoc:        TaxOffice,
		MinLoc:        TaxOffice,
		Name:          "tax inspector",
		Number:        ObTaxInspector, // 1028,
		ScanMessageNo: text.TaxInspector_Scan,
		Sex:           'm',
		Synonyms:      []string{"inspector"},
	},
	{
		DescMessageNo: text.TDX_Desc,
		DropEvent:     4,
		MaxLoc:        Store,
		MinLoc:        Store,
		Name:          "TDX",
		Number:        ObTDX, // 802,
		ScanMessageNo: text.TDX_Scan,
		Sex:           'n',
		Value:         80,
		Weight:        8,
	},
	{
		DescMessageNo: text.Technician_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    -1,
		MaxLoc:        ControlRoom4,
		MinLoc:        ControlRoom4,
		Name:          "technician",
		Number:        ObTechnician,   // 1036,
		PrefObject:    ObOscilloscope, // 814,
		ScanMessageNo: text.Technician_Scan,
		Sex:           'f',
	},
	{
		AttackPercent: 78,
		DescMessageNo: text.Thug_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    30,
		MaxLoc:        MoonRay3,
		MinLoc:        MoonRay1,
		Name:          "thug",
		Number:        ObThug, // 1016,
		ScanMessageNo: text.Thug_Scan,
		Sex:           'm',
	},
	{
		DescMessageNo: text.Tourist_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    82,
		MaxLoc:        RepairShop,
		MinLoc:        StarshipCantina,
		Name:          "tourist",
		Number:        ObTourist, // 1025,
		ScanMessageNo: text.Tourist_Scan,
		Sex:           'm',
	},
	{
		DescMessageNo: text.Traveler_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    350,
		MaxLoc:        Slideway2,
		MinLoc:        Plaza1,
		Name:          "traveler",
		Number:        ObTraveler, // 1037,
		PrefObject:    ObSandwich, // 827,
		ScanMessageNo: text.Traveler_Scan,
		Sex:           'f',
	},
	{
		DescMessageNo: text.Typist_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    300,
		MaxLoc:        TypingPool2,
		MinLoc:        TypingPool1,
		Name:          "typist",
		Number:        ObTypist, // 1027,
		PrefObject:    ObPearls, // 824,
		ScanMessageNo: text.Typist_Scan,
		Sex:           'f',
	},
	{
		DescMessageNo: text.UrbanSpaceman_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    3,
		MaxLoc:        Corridor56,
		MinLoc:        Corridor47,
		Name:          "urban spaceman",
		Number:        ObUrbanSpaceman, // 1042,
		PrefObject:    ObFlower,        // 820,
		ScanMessageNo: text.UrbanSpaceman_Scan,
		Sex:           'm',
		Synonyms:      []string{"spaceman"},
	},
	{
		DescMessageNo: text.Vandier_Desc,
		MaxLoc:        Temple5,
		MinLoc:        Temple1,
		Name:          "vandier",
		Number:        ObVandier, // 819,
		ScanMessageNo: text.Vandier_Scan,
		Sex:           'n',
		Value:         175,
		Weight:        17,
	},
	{
		DescMessageNo: text.Vega_Desc,
		Flags:         model.OfAnimate | model.OfShip,
		MaxCounter:    7,
		MaxLoc:        MarsOrbit,
		MinLoc:        TitanOrbit,
		Name:          "Vega",
		Number:        ObVega, // 1002,
		ScanMessageNo: text.Vega_Scan,
		Sex:           'f',
		Value:         -500,

		ShipGuns: [4]model.OldSGuns{
			core.SHIP_GUN_LASER,
			core.SHIP_GUN_MAG_GUN,
		},

		ShipKit: core.Equipment{
			Computer: 2,
			Engine:   10,
			Fuel:     150,
			Hold:     200,
			Hull:     25,
			Shield:   10,
			Tonnage:  800,
		},
	},
	{
		DescMessageNo: text.Watch_Desc,
		MaxLoc:        Corridor22,
		MinLoc:        Corridor17,
		Name:          "watch",
		Number:        ObWatch, // 809,
		ScanMessageNo: text.Watch_Scan,
		Sex:           'n',
		Value:         125,
		Weight:        1,
	},
	{
		DescMessageNo: text.Watchman_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    -1,
		MaxLoc:        SpacePort1,
		MinLoc:        SpacePort1,
		Name:          "watchman",
		Number:        ObWatchman, // 1010,
		ScanMessageNo: text.Watchman_Scan,
		Sex:           'm',
	},
	{
		DescMessageNo: text.Workman_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    40,
		MaxLoc:        Workshop24,
		MinLoc:        Workshop1,
		Name:          "workman",
		Number:        ObWorkman, // 1013,
		PrefObject:    ObSpanner, // 807,
		ScanMessageNo: text.Workman_Scan,
		Sex:           'm',
	},
	{
		DescMessageNo: text.Xerxes_Desc,
		Flags:         model.OfAnimate | model.OfShip,
		MaxCounter:    15,
		MaxLoc:        SolarOrbit4,
		MinLoc:        SolarSystemInterstellarLink,
		Name:          "Xerxes",
		Number:        ObXerxes, // 1000,
		ScanMessageNo: text.Xerxes_Scan,
		Sex:           'f',
		Value:         -1500,

		ShipGuns: [4]model.OldSGuns{
			core.SHIP_GUN_TWIN_LASER,
			core.SHIP_GUN_MISSILE,
		},

		ShipKit: core.Equipment{
			Computer: 2,
			Engine:   15,
			Fuel:     80,
			Hold:     20,
			Hull:     35,
			Shield:   15,
			Tonnage:  200,
		},
	},
	{
		DescMessageNo: text.Miner_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    45,
		MaxLoc:        Tunnel2,
		MinLoc:        RockFace2,
		Name:          "miner",
		Number:        ObMiner, // 1038,
		PrefObject:    ObTDX,   // 802,
		ScanMessageNo: text.Miner_Scan,
		Sex:           'm',
	},
	{
		AttackPercent: 99,
		DescMessageNo: text.Smuggler_Desc,
		Flags:         model.OfAnimate | model.OfCleaner,
		MaxCounter:    10,
		MaxLoc:        LivingQuarters3,
		MinLoc:        Tunnel5,
		Name:          "smuggler",
		Number:        ObSmuggler, // 1039,
		ScanMessageNo: text.Smuggler_Scan,
		Sex:           'm',
	},
	{
		DescMessageNo: text.Policeman_Desc,
		Flags:         model.OfAnimate | model.OfStoic,
		MaxCounter:    -1,
		MaxLoc:        PoliceStation,
		MinLoc:        PoliceStation,
		Name:          "police officer",
		Number:        ObPoliceman, // 1040,
		ScanMessageNo: text.Policeman_Scan,
		Sex:           'm',
		Synonyms:      []string{"officer"},
	},
	{
		DescMessageNo: text.Hunter_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    65,
		MaxLoc:        Chamber,
		MinLoc:        DeadEnd1,
		Name:          "hunter",
		Number:        ObHunter, // 1043,
		PrefObject:    ObMap,    // 815,
		ScanMessageNo: text.Hunter_Scan,
		Sex:           'm',
	},
	{
		DescMessageNo: text.Tinguey_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    15,
		MaxLoc:        Corridor5,
		MinLoc:        LinkBetweenDomes1,
		Name:          "tinguey",
		Number:        ObTinguey, // 1044,
		PrefObject:    ObSargeur, // 825,
		ScanMessageNo: text.Tinguey_Scan,
		Sex:           'n',
	},
	{
		DescMessageNo: text.Weeble_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    150,
		MaxLoc:        SmartOffice,
		MinLoc:        EntranceHall,
		Name:          "weeble",
		Number:        ObWeeble,      // 1045,
		PrefObject:    ObJamRolyPoly, // 837,
		ScanMessageNo: text.Weeble_Scan,
		Sex:           'n',
	},
	{
		DescMessageNo: text.JamRolypoly_Desc,
		Flags:         model.OfEdible,
		MaxLoc:        GrizzlesLair2,
		MinLoc:        GrizzlesLair2,
		Name:          "jam roly-poly",
		Number:        ObJamRolyPoly, // 837,
		ScanMessageNo: text.JamRolypoly_Scan,
		Sex:           'n',
		Synonyms:      []string{"roly-poly"},
		Value:         35,
		Weight:        2,
	},
	{
		DescMessageNo: text.Fogg_Desc,
		Flags:         model.OfAnimate | model.OfNoThe,
		MaxCounter:    -1,
		MaxLoc:        DrFoggsMaritalArtsEmporium,
		MinLoc:        DrFoggsMaritalArtsEmporium,
		Name:          "Dr Fogg",
		Number:        ObFogg, // 1047,
		ScanMessageNo: text.Fogg_Scan,
		Sex:           'm',
		Synonyms:      []string{"Fogg"},
	},
	{
		DescMessageNo: text.Krystal_Desc,
		Flags:         model.OfAnimate | model.OfNoThe,
		MaxCounter:    15,
		MaxLoc:        GuardRoom5,
		MinLoc:        Lift1,
		Name:          "Krystal",
		Number:        ObKrystal,    // 1048,
		PrefObject:    ObPhotograph, // 816,
		ScanMessageNo: text.Krystal_Scan,
		Sex:           'f',
	},
	{
		DescMessageNo: text.Cat_Desc,
		Flags:         model.OfAnimate,
		GetEvent:      8,
		MaxCounter:    45,
		MaxLoc:        GuardRoom5,
		MinLoc:        ThoriumStoreroom,
		Name:          "cat",
		Number:        ObCat, // 1006,
		ScanMessageNo: text.Cat_Scan,
		Sex:           'n',
	},
	{
		DescMessageNo: text.Diesel_Desc,
		Flags:         model.OfAnimate | model.OfNoThe,
		MaxCounter:    -1,
		MaxLoc:        ChezDiesel,
		MinLoc:        ChezDiesel,
		Name:          "Diesel",
		Number:        ObDiesel, // 1050,
		ScanMessageNo: text.Diesel_Scan,
		Sex:           'f',
		Value:         -500,
	},
	{
		DescMessageNo: text.Jeweller_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    -1,
		MaxLoc:        JewelleryShop,
		MinLoc:        JewelleryShop,
		Name:          "jeweler",
		Number:        ObJeweller, // 1051,
		PrefObject:    ObDiamond,  // 822,
		ScanMessageNo: text.Jeweller_Scan,
		Sex:           'f',
	},
	{
		AttackPercent: 95,
		DescMessageNo: text.Pirate_Desc,
		Flags:         model.OfAnimate | model.OfShip,
		MaxCounter:    4,
		MaxLoc:        SolarOrbit4,
		MinLoc:        InterplanetarySpace1,
		Name:          "pirate",
		Number:        ObPirate, // 1054,
		ScanMessageNo: text.Pirate_Scan,
		Sex:           'n',
		Synonyms:      []string{"Monty"},
		Value:         750000,

		ShipGuns: [4]model.OldSGuns{
			core.SHIP_GUN_TWIN_LASER,
			core.SHIP_GUN_TWIN_LASER,
			core.SHIP_GUN_MISSILE,
		},

		ShipKit: core.Equipment{
			Computer: 5,
			Engine:   30,
			Fuel:     200,
			Hold:     100,
			Hull:     50,
			Shield:   30,
			Tonnage:  600,
		},
	},
	{
		DescMessageNo: text.LV_Desc,
		MaxLoc:        Cave2,
		MinLoc:        Tunnel11,
		Name:          "Luncheon Voucher",
		Number:        ObLV, // 854,
		ScanMessageNo: text.LV_Scan,
		Sex:           'n',
		Synonyms:      []string{"LV", "voucher"},
		Value:         5,
		Weight:        1,
	},
	{
		DescMessageNo: text.Coin_Desc,
		MaxLoc:        LinkingCorridor,
		MinLoc:        Link1,
		Name:          "coin",
		Number:        ObCoin, // 856,
		ScanMessageNo: text.Coin_Scan,
		Sex:           'n',
		Value:         283,
		Weight:        1,
	},
	{
		DescMessageNo: text.SugarLump_Desc,
		Flags:         model.OfEdible,
		MaxLoc:        ShuttleStation1,
		MinLoc:        Slideway5,
		Name:          "sugar cube",
		Number:        ObSugarLump, // 858,
		ScanMessageNo: text.SugarLump_Scan,
		Sex:           'n',
		Synonyms:      []string{"sugar", "cube"},
		Value:         10,
		Weight:        1,
	},
	{
		AttackPercent: 85,
		DescMessageNo: text.Grizzle_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    18,
		MaxLoc:        Cave9,
		MinLoc:        DeadEnd1,
		Name:          "grizzle",
		Number:        ObGrizzle, // 1041,
		ScanMessageNo: text.Grizzle_Scan,
		Sex:           'n',
	},
	{
		DescMessageNo: text.BlackBox_Desc,
		Events:        [4]uint16{0, 54, 54, 0},
		MaxLoc:        SmartOffice,
		MinLoc:        Lift1,
		Name:          "black box",
		Number:        ObBlackBox, // 863,
		ScanMessageNo: text.BlackBox_Scan,
		Sex:           'n',
		Synonyms:      []string{"box"},
		Value:         600,
		Weight:        1,
	},
	{
		DescMessageNo: text.WHOOSH_Desc,
		Flags:         model.OfLiquid,
		MaxLoc:        Museum1,
		MinLoc:        Museum1,
		Name:          "WHOOSH",
		Number:        ObWHOOSH, // 865,
		ScanMessageNo: text.WHOOSH_Scan,
		Sex:           'n',
		Value:         65,
		Weight:        1,
	},
	{
		DescMessageNo: text.Katov_Desc,
		Flags:         model.OfAnimate | model.OfNoThe,
		MaxCounter:    -1,
		MaxLoc:        ControlRoom3,
		MinLoc:        ControlRoom3,
		Name:          "Admiral Katov",
		Number:        ObKatov, // 1055,
		ScanMessageNo: text.Katov_Scan,
		Sex:           'f',
		Synonyms:      []string{"Katov"},
	},
	{
		DescMessageNo: text.Squad_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    25,
		MaxLoc:        Road6,
		MinLoc:        ParadeGround1,
		Name:          "squad",
		Number:        ObSquad, // 1056,
		ScanMessageNo: text.Squad_Scan,
		Sex:           'm',
	},
	{
		DescMessageNo: text.Controller_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    -1,
		MaxLoc:        ControlTower2,
		MinLoc:        ControlTower2,
		Name:          "controller",
		Number:        ObController, // 1057,
		ScanMessageNo: text.Controller_Scan,
		Sex:           'm',
	},
	{
		DescMessageNo: text.Marine_Desc,
		Flags:         model.OfAnimate,
		MaxCounter:    -1,
		MaxLoc:        Barracks1,
		MinLoc:        Barracks1,
		Name:          "marine",
		Number:        ObMarine, // 1058,
		ScanMessageNo: text.Marine_Scan,
		Sex:           'm',
	},
	{
		DescMessageNo: text.Globe_Desc,
		Flags:         model.OfHidden,
		MaxLoc:        Courtyard2,
		MinLoc:        Courtyard2,
		Name:          "golden globe",
		Number:        ObGlobe, // 1059,
		ScanMessageNo: text.Globe_Scan,
		Sex:           'n',
		Synonyms:      []string{"globe"},
		Weight:        math.MaxUint16,
	},
	{
		DescMessageNo: text.Thimble_Desc,
		MaxLoc:        Workshop24,
		MinLoc:        Workshop1,
		Name:          "thimble",
		Number:        ObThimble, // 872,
		ScanMessageNo: text.Thimble_Scan,
		Sex:           'n',
		Value:         65,
		Weight:        1,
	},
	{
		DescMessageNo: text.Soap_Desc,
		MaxLoc:        Bathroom,
		MinLoc:        Bathroom,
		Name:          "soap",
		Number:        ObSoap, // 874,
		ScanMessageNo: text.Soap_Scan,
		Sex:           'n',
		Value:         36,
		Weight:        1,
	},
	{
		DescMessageNo: text.Badge_Desc,
		Flags:         model.OfHidden,
		Name:          "smiley badge",
		Number:        ObBadge, // 876,
		ScanMessageNo: text.Badge_Scan,
		Sex:           'n',
		Synonyms:      []string{"badge"},
		Value:         15,
		Weight:        1,
	},
	{
		DescMessageNo: text.RouletteWheel_Desc,
		MaxLoc:        Casino,
		MinLoc:        Casino,
		Name:          "roulette wheel",
		Number:        ObRouletteWheel, // 900,
		ScanMessageNo: text.RouletteWheel_Scan,
		Sex:           'n',
		Weight:        math.MaxUint16,
	},
	{
		DescMessageNo: text.Token_Desc,
		MaxLoc:        MazeOfAlleys12,
		MinLoc:        MazeOfAlleys1,
		Name:          "token",
		Number:        ObToken, // 901,
		ScanMessageNo: text.Token_Scan,
		Sex:           'n',
		Value:         200,
		Weight:        1,
	},
	{
		AttackPercent: 75,
		DescMessageNo: text.Zlitherworm_Desc,
		Flags:         model.OfAnimate,
		GetEvent:      7,
		MaxCounter:    8,
		MaxLoc:        SheerCliff,
		MinLoc:        BeforeAirlock,
		Name:          "zlitherworm",
		Number:        ObZlitherworm, // 1004,
		ScanMessageNo: text.Zlitherworm_Scan,
		Sex:           'n',
		Synonyms:      []string{"worm"},
	},
}

var Planets = [7]core.Planet{
	{
		Name:       "Titan",
		Level:      uint16(model.LevelMining),
		Population: 2000,
		RouteFlag:  model.PL1_TITAN,
		Landing:    TitanLandingArea,
		Orbit:      TitanOrbit,
		Exchange:   Market1,
		Markup:     10,

		Goods: [52][6]int16{
			{16, 1, -100, 2980, 44, 0},
			{23, 1, 3349, 2480, 16, 0},
			{19, 1, -100, 2820, 38, 0},
			{20, 1, -13, 3560, 38, 0},
			{2, 1, -100, 3140, 42, 0},
			{8, 1, -100, 3300, 48, 0},
			{12, 1, -100, 2840, 36, 0},
			{13, 0, 4520, 2260},
			{8, 1, 2761, 3660, 16, 0},
			{18, 0, 7523, 3760, 14, 0},
			{13, 1, 737, 2420, 38, 0},
			{18, 1, 1305, 2500, 34, 0},
			{1, 1, 2237, 2140, 24, 0},
			{27, 1, 5832, 2920, 8, 0},
			{42, 0, 7881, 3940, 30, 0},
			{1, 1, 1688, 2100, 34, 0},
			{31, 1, 4075, 2060},
			{0, 1, 3840, 3840},
			{0, 1, 2910, 2280},
			{0, 1, 2327, 2320},
			{0, 1, 3929, 3920},
			{0, 1, 2515, 2260},
			{0, 1, 3045, 2880},
			{0, 1, 2967, 2960},
			{0, 1, -18, 2700},
			{0, 1, 2270, 2020},
			{0, 1, 3385, 3260},
			{0, 1, -4, 2140},
			{0, 1, 163, 2160},
			{0, 1, 3230, 2980},
			{0, 1, -17, 2400},
			{0, 1, -6, 3140},
			{0, 1, 3710, 3460},
			{0, 1, 2353, 2280},
			{0, 1, 3774, 3760},
			{0, 1, 4207, 3660},
			{0, 1, 2670, 2620},
			{0, 1, 2824, 2600},
			{0, 1, 4240, 3920},
			{0, 1, 2283, 2280},
			{0, 1, 3700, 3400},
			{0, 1, 3530, 3380},
			{0, 1, 3904, 3820},
			{0, 1, 3541, 3540},
			{0, 1, 2307, 2300},
			{0, 1, 3763, 3760},
			{0, 1, -1, 2700},
			{0, 1, -5, 2380},
			{0, 1, 2640, 2640},
			{0, 1, -19, 2980},
			{0, 1, 2790, 2700},
			{0, 1, 2567, 2560},
		},
	},
	{
		Name:       "Castillo",
		Synonym:    "Callisto",
		Level:      uint16(model.LevelNoProduction),
		Population: 1000,
		RouteFlag:  model.PL1_CALLISTO,
		Landing:    CallistoLandingPad,
		Orbit:      CallistoOrbit,
	},
	{
		Name:       "Mars",
		Level:      uint16(model.LevelTechnological),
		Population: 4000,
		RouteFlag:  model.PL1_MARS,
		Landing:    MarsLandingArea,
		Orbit:      MarsOrbit,
		Exchange:   Market3,
		Markup:     10,

		Goods: [52][6]int16{
			{25, 0, 4921, 2460, 8, 0},
			{28, 1, 3324, 3080, 24, 0},
			{26, 0, 5240, 2620, 0, 0},
			{92, 0, 5923, 2960, 46, 0},
			{18, 1, -100, 3320, 50, 0},
			{49, 0, 5483, 2740, 0, 0},
			{33, 1, -100, 2120, 64, 0},
			{91, 0, 7211, 3600, 0, 0},
			{91, 0, 7327, 3660, 18, 0},
			{33, 0, 7840, 3920, 4, 0},
			{96, 1, 2738, 2160, 24, 0},
			{19, 1, 3031, 3220, 22, 0},
			{17, 1, 1138, 2280, 38, 0},
			{5, 1, -100, 2740, 18, 0},
			{33, 0, 5002, 2500, 4, 0},
			{14, 1, 2433, 2500, 18, 0},
			{19, 1, -22, 2560, 0, 0},
			{0, 1, 3726, 3720, 0, 0},
			{0, 1, 2580, 2580, 0, 0},
			{0, 1, 2014, 2000, 0, 0},
			{0, 1, 3150, 3000, 0, 0},
			{0, 1, -20, 3040, 0, 0},
			{0, 1, -3, 3340, 0, 0},
			{0, 1, 342, 2180, 0, 0},
			{0, 1, 0, 3540, 0, 0},
			{0, 1, 3712, 2260, 0, 0},
			{0, 1, 1119, 2820, 0, 0},
			{0, 1, 2262, 2000, 0, 0},
			{0, 1, 466, 3040, 0, 0},
			{0, 1, 3286, 2440, 0, 0},
			{0, 1, 705, 3920, 0, 0},
			{0, 1, -72, 3520, 0, 0},
			{0, 1, 3207, 2560, 0, 0},
			{0, 1, 2691, 2540, 0, 0},
			{0, 1, 2605, 2480, 0, 0},
			{0, 1, 3036, 2920, 0, 0},
			{0, 1, 2577, 2320, 0, 0},
			{0, 1, 3890, 3640, 0, 0},
			{0, 1, 2632, 2520, 0, 0},
			{0, 1, 475, 2000, 0, 0},
			{0, 1, 2850, 2600, 0, 0},
			{0, 1, 2830, 2680, 0, 0},
			{0, 1, -1, 3220, 0, 0},
			{0, 1, 3050, 2900, 0, 0},
			{0, 1, 3440, 3440, 0, 0},
			{0, 1, 3720, 3720, 0, 0},
			{0, 1, 3400, 3400, 0, 0},
			{0, 1, 3043, 3040, 0, 0},
			{0, 1, -15, 2480, 0, 0},
			{0, 1, 3980, 3980, 0, 0},
			{0, 1, 3190, 3040, 0, 0},
			{0, 1, 2520, 2520, 0, 0},
		},
	},
	{
		Name:       "Earth",
		Level:      uint16(model.LevelCapital),
		Population: 5000,
		RouteFlag:  model.PL1_EARTH,
		Landing:    EarthLandingArea,
		Orbit:      EarthOrbit,
		Hospital:   HospitalWard3,
	},
	{
		Name:       "Moon",
		Level:      uint16(model.LevelIndustrial),
		Population: 3000,
		RouteFlag:  model.PL1_MOON,
		Landing:    SelenaLandingPad,
		Orbit:      LunarOrbit,
		Exchange:   Market2,
		Markup:     10,

		Goods: [52][6]int16{
			{22, 1, -100, 3160, 34, 0},
			{41, 1, 3815, 2380, 26, 0},
			{6, 1, -100, 3200, 30, 0},
			{47, 0, 7280, 3640, 10, 0},
			{10, 1, -100, 3860, 26, 0},
			{7, 1, -100, 2520, 32, 0},
			{25, 0, 7280, 3640, 16, 0},
			{16, 1, 6639, 3320, 12, 0},
			{28, 1, 2955, 2900, 34, 0},
			{27, 0, 6280, 3140, 12, 0},
			{19, 1, 260, 2840, 38, 0},
			{15, 1, 1365, 3740, 20, 0},
			{57, 1, 3790, 2400, 40, 0},
			{72, 0, 6002, 3000, 14, 0},
			{45, 0, 4722, 2360, 14, 0},
			{69, 1, 4439, 2220, 18, 0},
			{29, 0, 5560, 2780, 0, 0},
			{0, 1, 2765, 2760, 0, 0},
			{0, 1, 3080, 3080, 0, 0},
			{0, 1, 2900, 2160, 0, 0},
			{0, 1, 3294, 2560, 0, 0},
			{0, 1, 3830, 3580, 0, 0},
			{0, 1, 2800, 2640, 0, 0},
			{0, 1, -50, 3200, 0, 0},
			{0, 1, 3765, 3640, 0, 0},
			{0, 1, 3060, 2660, 0, 0},
			{0, 1, -44, 2420, 0, 0},
			{0, 1, -28, 3780, 0, 0},
			{0, 1, -16, 2100, 0, 0},
			{0, 1, 4879, 3760, 0, 0},
			{0, 1, -58, 3740, 0, 0},
			{0, 1, 2550, 2300, 0, 0},
			{0, 1, -42, 2160, 0, 0},
			{0, 1, 3494, 2940, 0, 0},
			{0, 1, 2985, 2980, 0, 0},
			{0, 1, 3810, 3060, 0, 0},
			{0, 1, 4890, 3460, 0, 0},
			{0, 1, 2640, 2640, 0, 0},
			{0, 1, 2618, 2400, 0, 0},
			{0, 1, 3650, 3400, 0, 0},
			{0, 1, 4220, 2820, 0, 0},
			{0, 1, 2301, 2280, 0, 0},
			{0, 1, 2180, 2180, 0, 0},
			{0, 1, -3, 3220, 0, 0},
			{0, 1, 2354, 2340, 0, 0},
			{0, 1, 2621, 2540, 0, 0},
			{0, 1, 3462, 3460, 0, 0},
			{0, 1, 3127, 3120, 0, 0},
			{0, 1, 2374, 2360, 0, 0},
			{0, 1, -31, 3580, 0, 0},
			{0, 1, 565, 2320, 0, 0},
			{0, 1, -30, 3960, 0, 0},
		},
	},
	{
		Name:       "Venus",
		Level:      uint16(model.LevelTechnological),
		Population: 4000,
		RouteFlag:  model.PL1_VENUS,
		Landing:    LandingBay3,
		Orbit:      VenusOrbit,
		Exchange:   TradingExchange2,
		Markup:     10,

		Goods: [52][6]int16{
			{10, 1, 5948, 3160, 8, 0},
			{55, 1, -100, 2000, 78, 0},
			{62, 0, 4441, 2220, 24, 0},
			{61, 1, 102, 3860, 62, 0},
			{24, 1, 5319, 2660, 22, 0},
			{29, 1, -100, 2640, 36, 0},
			{38, 0, 7282, 3640, 8, 0},
			{36, 1, 7638, 3820, 32, 0},
			{61, 1, 2659, 2400, 60, 0},
			{50, 0, 4841, 2420, 38, 0},
			{32, 1, 464, 2120, 62, 0},
			{30, 1, 4061, 3600, 22, 0},
			{19, 1, -100, 3360, 36, 0},
			{0, 1, -100, 2160, 38, 0},
			{45, 0, 5360, 2680, 20, 0},
			{15, 1, 3940, 3880, 12, 0},
			{31, 1, -59, 2800, 0, 0},
			{0, 1, 2134, 2120, 0, 0},
			{0, 1, 3280, 3280, 0, 0},
			{0, 1, 4226, 3620, 0, 0},
			{0, 1, 2662, 2400, 0, 0},
			{0, 1, -6, 2520, 0, 0},
			{0, 1, -6, 3900, 0, 0},
			{0, 1, -8, 3780, 0, 0},
			{0, 1, 0, 2100, 0, 0},
			{0, 1, 1322, 2400, 0, 0},
			{0, 1, -29, 2900, 0, 0},
			{0, 1, 2460, 2460, 0, 0},
			{0, 1, -30, 3980, 0, 0},
			{0, 1, 1446, 3220, 0, 0},
			{0, 1, 409, 3480, 0, 0},
			{0, 1, -1, 2180, 0, 0},
			{0, 1, 2280, 2280, 0, 0},
			{0, 1, 4701, 3700, 0, 0},
			{0, 1, 2399, 2040, 0, 0},
			{0, 1, 3446, 2940, 0, 0},
			{0, 1, 2715, 2340, 0, 0},
			{0, 1, 3165, 2740, 0, 0},
			{0, 1, 3448, 2760, 0, 0},
			{0, 1, 473, 3780, 0, 0},
			{0, 1, 3827, 3160, 0, 0},
			{0, 1, 3940, 3940, 0, 0},
			{0, 1, -35, 3400, 0, 0},
			{0, 1, 3880, 3580, 0, 0},
			{0, 1, 2100, 2100, 0, 0},
			{0, 1, 3270, 3120, 0, 0},
			{0, 1, 3460, 3460, 0, 0},
			{0, 1, 3940, 3940, 0, 0},
			{0, 1, -2, 3240, 0, 0},
			{0, 1, 3780, 3780, 0, 0},
			{0, 1, 3710, 3560, 0, 0},
			{0, 1, 2680, 2680, 0, 0},
		},
	},
	{
		Name:       "Mercury",
		Level:      uint16(model.LevelIndustrial),
		Population: 3000,
		RouteFlag:  model.PL1_MERCURY,
		Landing:    LandingBay5,
		Orbit:      MercuryOrbit,
		Exchange:   TradingExchange3,
		Markup:     10,

		Goods: [52][6]int16{
			{15, 1, -100, 2300, 36, 0},
			{23, 1, 2720, 2500, 16, 0},
			{39, 1, 4738, 2480, 32, 0},
			{5, 1, -13, 2420, 10, 0},
			{37, 0, 5808, 2900, 0, 0},
			{30, 1, -100, 2420, 36, 0},
			{20, 0, 7841, 3920, 6, 0},
			{39, 1, 7838, 3920, 28, 0},
			{44, 1, 5054, 2700, 20, 0},
			{33, 1, 6799, 3400, 10, 0},
			{45, 1, 2580, 2120, 6, 0},
			{16, 1, -100, 2980, 28, 0},
			{10, 1, -100, 3100, 34, 0},
			{84, 0, 6400, 3200, 70, 0},
			{55, 1, 7839, 3920, 52, 0},
			{95, 1, 4399, 2480, 28, 0},
			{8, 0, 7880, 3940, 0, 0},
			{0, 1, 3740, 3740, 0, 0},
			{0, 1, 2280, 2280, 0, 0},
			{0, 1, 3270, 3120, 0, 0},
			{0, 1, 3170, 2720, 0, 0},
			{0, 1, 3680, 3680, 0, 0},
			{0, 1, 3444, 2700, 0, 0},
			{0, 1, -1, 2820, 0, 0},
			{0, 1, 4530, 3840, 0, 0},
			{0, 1, 3525, 3200, 0, 0},
			{0, 1, 0, 2760, 0, 0},
			{0, 1, 126, 3080, 0, 0},
			{0, 1, 0, 2980, 0, 0},
			{0, 1, 3530, 3360, 0, 0},
			{0, 1, -11, 3760, 0, 0},
			{0, 1, 2560, 2340, 0, 0},
			{0, 1, -52, 2920, 0, 0},
			{0, 1, 4696, 3280, 0, 0},
			{0, 1, 3925, 3800, 0, 0},
			{0, 1, 4385, 3820, 0, 0},
			{0, 1, 3845, 3720, 0, 0},
			{0, 1, 2475, 3720, 0, 0},
			{0, 1, 4355, 3760, 0, 0},
			{0, 1, 2420, 2420, 0, 0},
			{0, 1, 2430, 2180, 0, 0},
			{0, 1, 2920, 2920, 0, 0},
			{0, 1, 2440, 2440, 0, 0},
			{0, 1, -14, 3280, 0, 0},
			{0, 1, 2120, 2120, 0, 0},
			{0, 1, 2485, 2360, 0, 0},
			{0, 1, 3930, 3780, 0, 0},
			{0, 1, 3120, 3120, 0, 0},
			{0, 1, 3290, 3140, 0, 0},
			{0, 1, -54, 3100, 0, 0},
			{0, 1, 239, 2240, 0, 0},
			{0, 1, 0, 3860, 0, 0},
		},
	},
}
