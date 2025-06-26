package server

import "github.com/nosborn/federation-1999/internal/model"

// Stores a pallet of goods in a warehouse.
func StorePallet(warehouse *model.Warehouse, pallet model.Cargo) bool {
	for i := range 20 {
		if warehouse.Bay[i].Quantity == 0 {
			warehouse.Bay[i] = pallet
			return true
		}
	}
	return false
}
