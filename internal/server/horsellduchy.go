package server

import (
	"log"

	"github.com/nosborn/federation-1999/internal/collections"
)

func NewHorsellDuchy(name string) *Duchy {
	d := &Duchy{
		customsRate: 10,
		name:        name,
		systems:     collections.NewOrderedCollection[Systemer](),
	}
	log.Printf("%s duchy initialized", d.Name())
	return d
}
