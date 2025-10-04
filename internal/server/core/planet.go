package core

type Planet struct {
	// General info.

	Name                    string // [limits.NAME_SIZE]byte // Name of the planet
	Synonym                 string // [limits.NAME_SIZE]byte // Alternate name of the planet
	Level/*level_t*/ uint16 // planet's development level
	Population              int16 // Current population index

	RouteFlag uint32 // (Player) flag for GOTO

	// Locations.

	Landing  uint16 // Landing pad location number
	Orbit    uint16 // Orbit location number
	Hospital uint16 // Hospital location number
	Exchange uint16 // Exchange location number

	// Exchange info.

	Flux   int16 /* random price fluctuations (+40% to -40%) */
	Markup int16 /* Percentage markup on this exchange */
	Goods  [52][6]int16
}
