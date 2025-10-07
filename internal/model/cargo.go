package model

type Cargo struct {
	Type     Commodity // type of goods
	Quantity int32     // zero indicates an empty pallet
	Origin   string    // planet of origin
	Cost     int32     // cost of these goods
}
