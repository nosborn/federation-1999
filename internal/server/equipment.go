package server

type Equipment struct { // Ship's equipment record
	MaxHull     int32 // Hull strength
	CurHull     int32
	MaxShield   int32 // Shield strength
	CurShield   int32
	MaxEngine   int32 // Engine capacity
	CurEngine   int32
	MaxComputer int32 // Computer size
	CurComputer int32
	MaxFuel     int32 // Fuel
	CurFuel     int32
	MaxHold     int32 // Cargo capacity
	CurHold     int32
	Tonnage     int32 // Overall size of ship
}
