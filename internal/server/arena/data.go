package arena

import (
	"math"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/core"
	"github.com/nosborn/federation-1999/internal/text"
)

var Locations = [78]core.Location{
	{
		Number:     InterstellarLink,
		FullMsgNo:  text.ArenaLocation9_Full,
		BriefMsgNo: text.ArenaLocation9_Brief,
		Flags:      model.LfLink | model.LfSpace,

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
			Space1, // Down
			0,      // In
			0,      // Out
		},
	},
	{
		Number:     Space1,
		FullMsgNo:  text.ArenaLocation10_Full,
		BriefMsgNo: text.ArenaLocation10_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,                // N
			0,                // NE
			0,                // E
			Space3,           // SE
			Space2,           // S
			0,                // SW
			0,                // W
			0,                // NW
			InterstellarLink, // Up
			0,                // Down
			0,                // In
			0,                // Out
		},
	},
	{
		Number:     Space2,
		FullMsgNo:  text.ArenaLocation11_Full,
		BriefMsgNo: text.ArenaLocation11_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space1, // N
			0,      // NE
			Space3, // E
			Space5, // SE
			Space4, // S
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
		Number:     Space3,
		FullMsgNo:  text.ArenaLocation12_Full,
		BriefMsgNo: text.ArenaLocation12_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			0,       // E
			Space6,  // SE
			Space5,  // S
			Space4,  // SW
			Space2,  // W
			Space1,  // NW
			0,       // Up
			Space16, // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space4,
		FullMsgNo:  text.ArenaLocation13_Full,
		BriefMsgNo: text.ArenaLocation13_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space2, // N
			Space3, // NE
			Space5, // E
			Space9, // SE
			Space8, // S
			Space7, // SW
			0,      // W
			0,      // NW
			0,      // Up
			0,      // Down
			0,      // In
			0,      // Out
		},
	},
	{
		Number:     Space5,
		FullMsgNo:  text.ArenaLocation14_Full,
		BriefMsgNo: text.ArenaLocation14_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space3,  // N
			0,       // NE
			Space6,  // E
			Space10, // SE
			Space9,  // S
			Space8,  // SW
			Space4,  // W
			Space2,  // NW
			0,       // Up
			Space18, // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space6,
		FullMsgNo:  text.ArenaLocation15_Full,
		BriefMsgNo: text.ArenaLocation15_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			0,       // E
			Space11, // SE
			Space10, // S
			Space9,  // SW
			Space5,  // W
			Space3,  // NW
			0,       // Up
			Space19, // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space7,
		FullMsgNo:  text.ArenaLocation16_Full,
		BriefMsgNo: text.ArenaLocation16_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,       // N
			Space4,  // NE
			Space8,  // E
			Space13, // SE
			Space12, // S
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
		Number:     Space8,
		FullMsgNo:  text.ArenaLocation17_Full,
		BriefMsgNo: text.ArenaLocation17_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space4,  // N
			Space5,  // NE
			Space9,  // E
			Space14, // SE
			Space13, // S
			Space12, // SW
			Space7,  // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space9,
		FullMsgNo:  text.ArenaLocation18_Full,
		BriefMsgNo: text.ArenaLocation18_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space5,  // N
			Space6,  // NE
			Space10, // E
			0,       // SE
			Space14, // S
			Space13, // SW
			Space8,  // W
			Space4,  // NW
			0,       // Up
			Space21, // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space10,
		FullMsgNo:  text.ArenaLocation19_Full,
		BriefMsgNo: text.ArenaLocation19_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space6,  // N
			0,       // NE
			Space11, // E
			0,       // SE
			0,       // S
			Space14, // SW
			Space9,  // W
			Space5,  // NW
			0,       // Up
			Space22, // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space11,
		FullMsgNo:  text.ArenaLocation20_Full,
		BriefMsgNo: text.ArenaLocation20_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			0,       // E
			0,       // SE
			0,       // S
			0,       // SW
			Space10, // W
			Space6,  // NW
			0,       // Up
			Space23, // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space12,
		FullMsgNo:  text.ArenaLocation21_Full,
		BriefMsgNo: text.ArenaLocation21_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space7,  // N
			Space8,  // NE
			Space13, // E
			Space15, // SE
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
		Number:     Space13,
		FullMsgNo:  text.ArenaLocation22_Full,
		BriefMsgNo: text.ArenaLocation22_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space8,  // N
			Space9,  // NE
			Space14, // E
			0,       // SE
			Space15, // S
			0,       // SW
			Space12, // W
			Space7,  // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space14,
		FullMsgNo:  text.ArenaLocation23_Full,
		BriefMsgNo: text.ArenaLocation23_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space9,  // N
			Space10, // NE
			0,       // E
			0,       // SE
			0,       // S
			Space15, // SW
			Space13, // W
			Space8,  // NW
			0,       // Up
			Space26, // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space15,
		FullMsgNo:  text.ArenaLocation24_Full,
		BriefMsgNo: text.ArenaLocation24_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space13, // N
			Space14, // NE
			0,       // E
			0,       // SE
			0,       // S
			0,       // SW
			0,       // W
			Space12, // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space16,
		FullMsgNo:  text.ArenaLocation25_Full,
		BriefMsgNo: text.ArenaLocation25_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			Space17, // E
			Space19, // SE
			Space18, // S
			0,       // SW
			0,       // W
			0,       // NW
			Space3,  // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space17,
		FullMsgNo:  text.ArenaLocation26_Full,
		BriefMsgNo: text.ArenaLocation26_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			0,       // E
			0,       // SE
			Space19, // S
			Space18, // SW
			Space16, // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space18,
		FullMsgNo:  text.ArenaLocation27_Full,
		BriefMsgNo: text.ArenaLocation27_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space16, // N
			Space17, // NE
			Space19, // E
			Space22, // SE
			Space21, // S
			0,       // SW
			0,       // W
			0,       // NW
			Space5,  // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space19,
		FullMsgNo:  text.ArenaLocation28_Full,
		BriefMsgNo: text.ArenaLocation28_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space17, // N
			0,       // NE
			0,       // E
			Space23, // SE
			Space22, // S
			Space21, // SW
			Space18, // W
			Space16, // NW
			Space6,  // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space20,
		FullMsgNo:  text.ArenaLocation29_Full,
		BriefMsgNo: text.ArenaLocation29_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			0,       // E
			0,       // SE
			Space25, // S
			Space24, // SW
			0,       // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space21,
		FullMsgNo:  text.ArenaLocation30_Full,
		BriefMsgNo: text.ArenaLocation30_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space18, // N
			Space19, // NE
			Space22, // E
			Space27, // SE
			Space26, // S
			0,       // SW
			0,       // W
			0,       // NW
			Space9,  // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space22,
		FullMsgNo:  text.ArenaLocation31_Full,
		BriefMsgNo: text.ArenaLocation31_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space19, // N
			0,       // NE
			Space23, // E
			Space28, // SE
			Space27, // S
			Space26, // SW
			Space21, // W
			Space18, // NW
			Space10, // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space23,
		FullMsgNo:  text.ArenaLocation32_Full,
		BriefMsgNo: text.ArenaLocation32_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			Space24, // E
			0,       // SE
			Space28, // S
			Space27, // SW
			Space22, // W
			Space19, // NW
			Space11, // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space24,
		FullMsgNo:  text.ArenaLocation33_Full,
		BriefMsgNo: text.ArenaLocation33_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,       // N
			Space20, // NE
			Space25, // E
			0,       // SE
			0,       // S
			Space28, // SW
			Space23, // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space25,
		FullMsgNo:  text.ArenaLocation34_Full,
		BriefMsgNo: text.ArenaLocation34_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space20, // N
			0,       // NE
			0,       // E
			0,       // SE
			0,       // S
			0,       // SW
			Space24, // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space26,
		FullMsgNo:  text.ArenaLocation35_Full,
		BriefMsgNo: text.ArenaLocation35_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space21, // N
			Space22, // NE
			Space27, // E
			Space30, // SE
			Space29, // S
			0,       // SW
			0,       // W
			0,       // NW
			Space14, // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space27,
		FullMsgNo:  text.ArenaLocation36_Full,
		BriefMsgNo: text.ArenaLocation36_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space22, // N
			Space23, // NE
			Space28, // E
			Space31, // SE
			Space30, // S
			Space29, // SW
			Space26, // W
			Space21, // NW
			0,       // Up
			Space34, // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space28,
		FullMsgNo:  text.ArenaLocation37_Full,
		BriefMsgNo: text.ArenaLocation37_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space23, // N
			Space24, // NE
			0,       // E
			0,       // SE
			Space31, // S
			Space30, // SW
			Space27, // W
			Space22, // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space29,
		FullMsgNo:  text.ArenaLocation38_Full,
		BriefMsgNo: text.ArenaLocation38_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space26, // N
			Space27, // NE
			Space30, // E
			Space32, // SE
			0,       // S
			0,       // SW
			0,       // W
			0,       // NW
			0,       // Up
			Space40, // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space30,
		FullMsgNo:  text.ArenaLocation39_Full,
		BriefMsgNo: text.ArenaLocation39_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space27, // N
			Space28, // NE
			Space31, // E
			Space33, // SE
			Space32, // S
			0,       // SW
			Space29, // W
			Space26, // NW
			0,       // Up
			Space41, // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space31,
		FullMsgNo:  text.ArenaLocation40_Full,
		BriefMsgNo: text.ArenaLocation40_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space28, // N
			0,       // NE
			0,       // E
			0,       // SE
			Space33, // S
			Space32, // SW
			Space30, // W
			Space27, // NW
			0,       // Up
			Space42, // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space32,
		FullMsgNo:  text.ArenaLocation41_Full,
		BriefMsgNo: text.ArenaLocation41_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space30, // N
			Space31, // NE
			Space33, // E
			0,       // SE
			0,       // S
			0,       // SW
			0,       // W
			Space29, // NW
			0,       // Up
			Space49, // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space33,
		FullMsgNo:  text.ArenaLocation42_Full,
		BriefMsgNo: text.ArenaLocation42_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space31, // N
			0,       // NE
			0,       // E
			0,       // SE
			0,       // S
			0,       // SW
			Space32, // W
			Space30, // NW
			0,       // Up
			Space50, // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space34,
		FullMsgNo:  text.ArenaLocation43_Full,
		BriefMsgNo: text.ArenaLocation43_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			0,       // E
			Space42, // SE
			Space41, // S
			Space40, // SW
			0,       // W
			0,       // NW
			Space27, // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space35,
		FullMsgNo:  text.ArenaLocation44_Full,
		BriefMsgNo: text.ArenaLocation44_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			Space36, // E
			Space44, // SE
			Space43, // S
			Space42, // SW
			0,       // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space36,
		FullMsgNo:  text.ArenaLocation45_Full,
		BriefMsgNo: text.ArenaLocation45_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			0,       // E
			0,       // SE
			Space44, // S
			Space43, // SW
			Space35, // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space37,
		FullMsgNo:  text.ArenaLocation46_Full,
		BriefMsgNo: text.ArenaLocation46_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			Space38, // E
			Space46, // SE
			Space45, // S
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
		Number:     Space38,
		FullMsgNo:  text.ArenaLocation47_Full,
		BriefMsgNo: text.ArenaLocation47_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			Space39, // E
			Space47, // SE
			Space46, // S
			Space45, // SW
			Space37, // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space39,
		FullMsgNo:  text.ArenaLocation48_Full,
		BriefMsgNo: text.ArenaLocation48_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			Space40, // E
			Space48, // SE
			Space47, // S
			Space46, // SW
			Space38, // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space40,
		FullMsgNo:  text.ArenaLocation49_Full,
		BriefMsgNo: text.ArenaLocation49_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,       // N
			Space34, // NE
			Space41, // E
			Space49, // SE
			Space48, // S
			Space47, // SW
			Space39, // W
			0,       // NW
			Space29, // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space41,
		FullMsgNo:  text.ArenaLocation50_Full,
		BriefMsgNo: text.ArenaLocation50_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space34, // N
			0,       // NE
			Space42, // E
			Space50, // SE
			Space49, // S
			Space48, // SW
			Space40, // W
			0,       // NW
			Space30, // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space42,
		FullMsgNo:  text.ArenaLocation51_Full,
		BriefMsgNo: text.ArenaLocation51_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,       // N
			Space35, // NE
			Space43, // E
			Space51, // SE
			Space50, // S
			Space49, // SW
			Space41, // W
			Space34, // NW
			Space31, // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space43,
		FullMsgNo:  text.ArenaLocation52_Full,
		BriefMsgNo: text.ArenaLocation52_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space35, // N
			Space36, // NE
			Space44, // E
			Space52, // SE
			Space51, // S
			Space50, // SW
			Space42, // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space44,
		FullMsgNo:  text.ArenaLocation53_Full,
		BriefMsgNo: text.ArenaLocation53_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space36, // N
			0,       // NE
			0,       // E
			0,       // SE
			Space52, // S
			Space51, // SW
			Space43, // W
			Space35, // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space45,
		FullMsgNo:  text.ArenaLocation54_Full,
		BriefMsgNo: text.ArenaLocation54_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space37, // N
			Space38, // NE
			Space46, // E
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
		Number:     Space46,
		FullMsgNo:  text.ArenaLocation55_Full,
		BriefMsgNo: text.ArenaLocation55_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space38, // N
			Space39, // NE
			Space47, // E
			Space53, // SE
			0,       // S
			0,       // SW
			Space45, // W
			Space37, // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space47,
		FullMsgNo:  text.ArenaLocation56_Full,
		BriefMsgNo: text.ArenaLocation56_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space39, // N
			Space40, // NE
			Space48, // E
			0,       // SE
			Space53, // S
			0,       // SW
			Space46, // W
			Space38, // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space48,
		FullMsgNo:  text.ArenaLocation57_Full,
		BriefMsgNo: text.ArenaLocation57_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space40, // N
			Space41, // NE
			Space49, // E
			0,       // SE
			0,       // S
			Space53, // SW
			Space47, // W
			Space39, // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space49,
		FullMsgNo:  text.ArenaLocation58_Full,
		BriefMsgNo: text.ArenaLocation58_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space41, // N
			Space42, // NE
			Space50, // E
			0,       // SE
			0,       // S
			0,       // SW
			Space48, // W
			Space40, // NW
			Space32, // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space50,
		FullMsgNo:  text.ArenaLocation59_Full,
		BriefMsgNo: text.ArenaLocation59_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space42, // N
			Space43, // NE
			Space51, // E
			0,       // SE
			0,       // S
			0,       // SW
			Space49, // W
			Space41, // NW
			Space33, // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space51,
		FullMsgNo:  text.ArenaLocation60_Full,
		BriefMsgNo: text.ArenaLocation60_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space43, // N
			Space44, // NE
			Space52, // E
			0,       // SE
			0,       // S
			0,       // SW
			Space50, // W
			Space42, // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space52,
		FullMsgNo:  text.ArenaLocation61_Full,
		BriefMsgNo: text.ArenaLocation61_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space44, // N
			0,       // NE
			0,       // E
			0,       // SE
			0,       // S
			0,       // SW
			Space51, // W
			Space43, // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space53,
		FullMsgNo:  text.ArenaLocation62_Full,
		BriefMsgNo: text.ArenaLocation62_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space47, // N
			Space48, // NE
			0,       // E
			0,       // SE
			Space54, // S
			0,       // SW
			0,       // W
			Space46, // NW
			0,       // Up
			Space57, // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space54,
		FullMsgNo:  text.ArenaLocation63_Full,
		BriefMsgNo: text.ArenaLocation63_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space53, // N
			0,       // NE
			0,       // E
			0,       // SE
			0,       // S
			0,       // SW
			0,       // W
			0,       // NW
			0,       // Up
			Space60, // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space55,
		FullMsgNo:  text.ArenaLocation64_Full,
		BriefMsgNo: text.ArenaLocation64_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			Space56, // E
			Space59, // SE
			Space58, // S
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
		Number:     Space56,
		FullMsgNo:  text.ArenaLocation65_Full,
		BriefMsgNo: text.ArenaLocation65_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			Space57, // E
			Space60, // SE
			Space59, // S
			Space58, // SW
			Space55, // W
			0,       // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space57,
		FullMsgNo:  text.ArenaLocation66_Full,
		BriefMsgNo: text.ArenaLocation66_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			0,       // E
			Space61, // SE
			Space60, // S
			Space59, // SW
			Space56, // W
			0,       // NW
			Space53, // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space58,
		FullMsgNo:  text.ArenaLocation67_Full,
		BriefMsgNo: text.ArenaLocation67_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space55, // N
			Space56, // NE
			Space59, // E
			Space62, // SE
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
		Number:     Space59,
		FullMsgNo:  text.ArenaLocation68_Full,
		BriefMsgNo: text.ArenaLocation68_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space56, // N
			Space57, // NE
			Space60, // E
			Space63, // SE
			Space62, // S
			0,       // SW
			Space58, // W
			Space55, // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space60,
		FullMsgNo:  text.ArenaLocation69_Full,
		BriefMsgNo: text.ArenaLocation69_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space57, // N
			0,       // NE
			Space61, // E
			Space64, // SE
			Space63, // S
			Space62, // SW
			Space59, // W
			Space56, // NW
			Space54, // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space61,
		FullMsgNo:  text.ArenaLocation70_Full,
		BriefMsgNo: text.ArenaLocation70_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			0,       // N
			0,       // NE
			0,       // E
			0,       // SE
			Space64, // S
			Space63, // SW
			Space60, // W
			Space57, // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space62,
		FullMsgNo:  text.ArenaLocation71_Full,
		BriefMsgNo: text.ArenaLocation71_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space59, // N
			Space60, // NE
			Space63, // E
			Space65, // SE
			0,       // S
			0,       // SW
			0,       // W
			Space58, // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space63,
		FullMsgNo:  text.ArenaLocation72_Full,
		BriefMsgNo: text.ArenaLocation72_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space60, // N
			Space61, // NE
			Space64, // E
			Space66, // SE
			Space65, // S
			0,       // SW
			Space62, // W
			Space59, // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space64,
		FullMsgNo:  text.ArenaLocation73_Full,
		BriefMsgNo: text.ArenaLocation73_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space61, // N
			0,       // NE
			0,       // E
			0,       // SE
			Space66, // S
			Space65, // SW
			Space63, // W
			Space60, // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     Space65,
		FullMsgNo:  text.ArenaLocation74_Full,
		BriefMsgNo: text.ArenaLocation74_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space63,      // N
			Space64,      // NE
			Space66,      // E
			DunnigansEnd, // SE
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
		Number:     Space66,
		FullMsgNo:  text.ArenaLocation75_Full,
		BriefMsgNo: text.ArenaLocation75_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space64,      // N
			0,            // NE
			0,            // E
			0,            // SE
			DunnigansEnd, // S
			0,            // SW
			Space65,      // W
			Space63,      // NW
			0,            // Up
			0,            // Down
			0,            // In
			0,            // Out
		},
	},
	{
		Number:     DunnigansEnd,
		FullMsgNo:  text.ArenaLocation76_Full,
		BriefMsgNo: text.ArenaLocation76_Brief,
		Flags:      model.LfSpace,

		MovTab: [13]uint16{
			Space66, // N
			0,       // NE
			0,       // E
			0,       // SE
			0,       // S
			0,       // SW
			0,       // W
			Space65, // NW
			0,       // Up
			0,       // Down
			0,       // In
			0,       // Out
		},
	},
	{
		Number:     StaffRoom,
		FullMsgNo:  text.ArenaLocation77_Full,
		BriefMsgNo: text.ArenaLocation77_Brief,
		Flags:      model.LfCafe | model.LfShield,

		MovTab: [13]uint16{
			0,          // N
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
			TheRamaBar, // Out
		},
	},
	{
		Number:     FoundationHospital,
		FullMsgNo:  text.ArenaLocation78_Full,
		BriefMsgNo: text.ArenaLocation78_Brief,
		Flags:      model.LfHospital | model.LfShield,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			0,           // E
			TheMonolith, // SE
			0,           // S
			0,           // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			TheMonolith, // Out
		},
	},
	{
		Number:     AmberInsurance,
		FullMsgNo:  text.ArenaLocation79_Full,
		BriefMsgNo: text.ArenaLocation79_Brief,
		Flags:      model.LfIns | model.LfShield,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			0,           // E
			0,           // SE
			TheMonolith, // S
			0,           // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			TheMonolith, // Out
		},
	},
	{
		Number:     ArrakisExchange,
		FullMsgNo:  text.ArenaLocation80_Full,
		BriefMsgNo: text.ArenaLocation80_Brief,
		Flags:      model.LfShield | model.LfTrade,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			0,           // E
			0,           // SE
			0,           // S
			TheMonolith, // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			TheMonolith, // Out
		},
	},
	{
		Number:     DiGrizWeapons,
		FullMsgNo:  text.ArenaLocation81_Full,
		BriefMsgNo: text.ArenaLocation81_Brief,
		Flags:      model.LfShield | model.LfWeap,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			TheMonolith, // E
			0,           // SE
			0,           // S
			0,           // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			TheMonolith, // Out
		},
	},
	{
		Number:     TheMonolith,
		FullMsgNo:  text.ArenaLocation82_Full,
		BriefMsgNo: text.ArenaLocation82_Brief,
		Flags:      model.LfLanding | model.LfShield,

		MovTab: [13]uint16{
			AmberInsurance,            // N
			ArrakisExchange,           // NE
			MoteShipyards,             // E
			RingworldEngineering,      // SE
			TheRamaBar,                // S
			TessierAshpoolElectronics, // SW
			DiGrizWeapons,             // W
			FoundationHospital,        // NW
			0,                         // Up
			0,                         // Down
			0,                         // In
			0,                         // Out
		},
	},
	{
		Number:     MoteShipyards,
		FullMsgNo:  text.ArenaLocation83_Full,
		BriefMsgNo: text.ArenaLocation83_Brief,
		Flags:      model.LfShield | model.LfYard,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			0,           // E
			0,           // SE
			0,           // S
			0,           // SW
			TheMonolith, // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			TheMonolith, // Out
		},
	},
	{
		Number:     TessierAshpoolElectronics,
		FullMsgNo:  text.ArenaLocation84_Full,
		BriefMsgNo: text.ArenaLocation84_Brief,
		Flags:      model.LfCom | model.LfShield,

		MovTab: [13]uint16{
			0,           // N
			TheMonolith, // NE
			0,           // E
			0,           // SE
			0,           // S
			0,           // SW
			0,           // W
			0,           // NW
			0,           // Up
			0,           // Down
			0,           // In
			TheMonolith, // Out
		},
	},
	{
		Number:     TheRamaBar,
		FullMsgNo:  text.ArenaLocation85_Full,
		BriefMsgNo: text.ArenaLocation85_Brief,
		Flags:      model.LfCafe | model.LfShield,

		MovTab: [13]uint16{
			TheMonolith, // N
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
			TheMonolith, // Out
		},
	},
	{
		Number:     RingworldEngineering,
		FullMsgNo:  text.ArenaLocation86_Full,
		BriefMsgNo: text.ArenaLocation86_Brief,
		Flags:      model.LfRep | model.LfShield,

		MovTab: [13]uint16{
			0,           // N
			0,           // NE
			0,           // E
			0,           // SE
			0,           // S
			0,           // SW
			0,           // W
			TheMonolith, // NW
			0,           // Up
			0,           // Down
			0,           // In
			TheMonolith, // Out
		},
	},
}

var Objects = [6]core.Object{
	{
		Number:        ObAngst,
		Flags:         model.OfAnimate | model.OfShip,
		Name:          "Angst",
		DescMessageNo: text.Angst_Desc,
		ScanMessageNo: text.Angst_Scan,
		Sex:           model.SexFemale,
		MinLoc:        Space2,
		MaxLoc:        Space66,
		MaxCounter:    24,

		ShipGuns: [4]model.OldSGuns{
			core.SHIP_GUN_QUAD_LASER,
			core.SHIP_GUN_QUAD_LASER,
			core.SHIP_GUN_QUAD_LASER,
			core.SHIP_GUN_QUAD_LASER,
		},

		ShipKit: core.Equipment{
			Computer: 5,
			Engine:   100,
			Fuel:     1000,
			Hold:     1000,
			Hull:     240,
			Shield:   60,
			Tonnage:  1000,
		},
	},
	{
		Number:        ObCanram,
		Flags:         model.OfAnimate | model.OfShip,
		Name:          "Canram",
		DescMessageNo: text.Canram_Desc,
		ScanMessageNo: text.Canram_Scan,
		Sex:           model.SexFemale,
		MinLoc:        Space2,
		MaxLoc:        Space66,
		MaxCounter:    2,

		ShipGuns: [4]model.OldSGuns{
			core.SHIP_GUN_QUAD_LASER,
			core.SHIP_GUN_QUAD_LASER,
			core.SHIP_GUN_QUAD_LASER,
		},

		ShipKit: core.Equipment{
			Computer: 3,
			Engine:   100,
			Fuel:     1000,
			Hold:     1000,
			Hull:     40,
			Shield:   20,
			Tonnage:  1000,
		},
	},
	{
		Number:        ObFarnell,
		Flags:         model.OfAnimate | model.OfShip,
		Name:          "Farnell",
		DescMessageNo: text.Farnell_Desc,
		ScanMessageNo: text.Farnell_Scan,
		Sex:           model.SexFemale,
		MinLoc:        Space2,
		MaxLoc:        Space66,
		MaxCounter:    16,

		ShipGuns: [4]model.OldSGuns{
			core.SHIP_GUN_QUAD_LASER,
			core.SHIP_GUN_QUAD_LASER,
			core.SHIP_GUN_QUAD_LASER,
			core.SHIP_GUN_QUAD_LASER,
		},

		ShipKit: core.Equipment{
			Computer: 6,
			Engine:   100,
			Fuel:     1000,
			Hold:     1000,
			Hull:     140,
			Shield:   35,
			Tonnage:  1000,
		},
	},
	{
		Number:        ObFleetway,
		Flags:         model.OfAnimate | model.OfShip,
		Name:          "Fleetway",
		DescMessageNo: text.Fleetway_Desc,
		ScanMessageNo: text.Fleetway_Scan,
		Sex:           model.SexFemale,
		MinLoc:        Space2,
		MaxLoc:        Space66,
		MaxCounter:    8,

		ShipGuns: [4]model.OldSGuns{
			core.SHIP_GUN_QUAD_LASER,
			core.SHIP_GUN_QUAD_LASER,
			core.SHIP_GUN_QUAD_LASER,
			core.SHIP_GUN_QUAD_LASER,
		},

		ShipKit: core.Equipment{
			Computer: 5,
			Engine:   100,
			Fuel:     1000,
			Hold:     1000,
			Hull:     150,
			Shield:   40,
			Tonnage:  1000,
		},
	},
	{
		Number:        ObStarBase1,
		Flags:         model.OfAnimate | model.OfNoThe | model.OfShip,
		Name:          "StarBase1",
		DescMessageNo: text.StarBase1_Desc,
		ScanMessageNo: text.StarBase1_Scan,
		Sex:           model.SexNeuter,
		MinLoc:        Space2,
		MaxLoc:        Space66,
		Value:         -10000,
		MaxCounter:    30,

		ShipGuns: [4]model.OldSGuns{
			core.SHIP_GUN_QUAD_LASER,
			core.SHIP_GUN_QUAD_LASER,
			core.SHIP_GUN_QUAD_LASER,
			core.SHIP_GUN_QUAD_LASER,
		},

		ShipKit: core.Equipment{
			Computer: 14,
			Engine:   1000,
			Fuel:     1000,
			Hold:     1000,
			Hull:     math.MaxUint16,
			Shield:   10000,
			Tonnage:  50000,
		},
	},
	{
		Number:        ObTolson,
		Flags:         model.OfAnimate | model.OfShip,
		Name:          "Tolson",
		DescMessageNo: text.Tolson_Desc,
		ScanMessageNo: text.Tolson_Scan,
		Sex:           model.SexFemale,
		MinLoc:        Space2,
		MaxLoc:        Space66,
		MaxCounter:    4,

		ShipGuns: [4]model.OldSGuns{
			core.SHIP_GUN_QUAD_LASER,
			core.SHIP_GUN_QUAD_LASER,
			core.SHIP_GUN_QUAD_LASER,
		},

		ShipKit: core.Equipment{
			Computer: 4,
			Engine:   100,
			Fuel:     1000,
			Hold:     1000,
			Hull:     100,
			Shield:   30,
			Tonnage:  1000,
		},
	},
}

var Planet = core.Planet{
	Name:       "StarBase1",
	Level:      uint16(model.LevelLeisure),
	Population: 1000,
	Landing:    TheMonolith,
	Hospital:   FoundationHospital,
	Exchange:   ArrakisExchange,
	Flux:       10,

	Goods: [52][6]int16{
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{100, 0, 10000, 5000, 0, -20}, // Spices
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
		{0, 1, -100, 100, 1, 50},
	},
}
