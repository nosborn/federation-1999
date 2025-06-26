package model

const (
	LfDark   uint32 = 0x00000001 // location is unlit [L]
	LfSpace  uint32 = 0x00000002 // location is in space [S]
	LfDeath  uint32 = 0x00000004 // location kills player [D]
	LfVacuum uint32 = 0x00000008 // vacuum location [V]
	LfTrade  uint32 = 0x00000010 // trading exchange [T]
	LfYard   uint32 = 0x00000020 // ship yard [Y]

	LfLink uint32 = 0x00000080 // interstellar link [I]

	LfGen      uint32 = 0x00000200 // general store [G]
	LfWeap     uint32 = 0x00000400 // weapon shop [W]
	LfCafe     uint32 = 0x00000800 // cafe/bar [R]
	LfRep      uint32 = 0x00001000 // ship repairer [B]
	LfCom      uint32 = 0x00002000 // comms shop [E]
	LfClth     uint32 = 0x00004000 // clothing shop [F]
	LfPeace    uint32 = 0x00008000 // no fighting [P]
	LfHospital uint32 = 0x00010000 // hospital location [H]
	LfIns      uint32 = 0x00020000 // Insurance broker
	LfLock     uint32 = 0x00040000 // Lockable - dropped objects recycled

	LfShield   uint32 = 0x00100000 // Location is teleport-shielded
	LfLanding  uint32 = 0x00200000 // Landing pad
	LfOrbit    uint32 = 0x00400000 // Planetary orbit
	LfIndoors  uint32 = 0x00800000 // Location is indoors
	LfOutdoors uint32 = 0x01000000 // Location is outdoors
	LfHidden   uint32 = 0x02000000 // Pretend location doesn't exist
)

const (
	// Allowable flags for a player planet SPACE location.
	LfSpaceAllowed uint32 = LfSpace | LfDeath | LfVacuum | LfLink | LfPeace | LfOrbit

	// Disallowed flags for a player planet non-SPACE location.
	LfGroundDenied uint32 = LfSpace | LfLink | LfOrbit
)

const ( // Symbolic names for mov_tab entries.
	MvNorth = iota
	MvNE
	MvEast
	MvSE
	MvSouth
	MvSW
	MvWest
	MvNW
	MvUp
	MvDown
	MvIn
	MvOut
	MvPlanet
)
