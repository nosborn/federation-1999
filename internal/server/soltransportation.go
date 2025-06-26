package server

// type SolTransportation struct {
// 	Transportation
// }

func NewSolTransportation(sol *Duchy) *Transportation {
	// FIXME: needs a local timer
	return NewTransportation(sol, "Transportation Central")
}
