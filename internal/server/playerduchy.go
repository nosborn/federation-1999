package server

import (
	"log"

	"github.com/nosborn/federation-1999/internal/collections"
)

func NewPlayerDuchy(owner *Player, name string, customsRate, favouredRate int32) *Duchy {
	d := &Duchy{
		customsRate:  customsRate,
		favouredRate: favouredRate,
		name:         name,
		owner:        owner,
		systems:      collections.NewOrderedCollection[Systemer](),
		taxRate:      10, // FIXME
	}
	if err := allDuchies.Insert(d); err != nil {
		log.Panic("PANIC: Duplicate duchy added: ", err)
	}
	duchyIndex.Insert(name, d)
	d.transportation = NewTransportation(d, name)
	return d
}
