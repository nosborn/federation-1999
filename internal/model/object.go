package model

const (
	OfLight   uint32 = 0x00000001 // Object is a light source (T)
	OfLiquid  uint32 = 0x00000040 // Object is a liquid (H)
	OfEdible  uint32 = 0x00000080 // Object is edible (E)
	OfAnimate uint32 = 0x00000200 // Object is a mobile (M)
	OfShip    uint32 = 0x00000400 // Mobile is spaceship (S)

	OfNoThe   uint32 = 0x00002000 // Don't insert a 'the' if set
	OfMusic   uint32 = 0x00004000 // Object is a music instrument
	OfCleaner uint32 = 0x00008000 // Mobile is a cleaning droid

	OfIndoors  uint32 = 0x01000000 // Must stay in INDOOR locations
	OfOutdoors uint32 = 0x02000000 // Must stay in OUTDOOR locations

	OfStoic  uint32 = 0x10000000 // No interactions
	OfDuke   uint32 = 0x20000000 // Object is set up for the Duke puzzle
	OfHidden uint32 = 0x40000000 // Object is out of the game
)

type OldSGuns struct {
	Type   int16
	Name   string
	Damage int16
	Power  int16
}
