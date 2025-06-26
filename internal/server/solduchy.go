package server

import "github.com/nosborn/federation-1999/internal/collections"

// func init() {
// 	if err := allDuchies.Insert(SolDuchy); err != nil {
// 		log.Panic("PANIC: Duplicate duchy added: ", err)
// 	}
// }

func NewSolDuchy() *Duchy {
	d := &Duchy{
		customsRate: 10,
		name:        "Sol",
		systems:     collections.NewOrderedCollection[Systemer](),
		taxRate:     25,
	}
	duchyIndex.Insert(d.name, d)
	d.transportation = NewSolTransportation(d)
	d.StartTransportation()
	return d
}
