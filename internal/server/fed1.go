package server

import (
	"log"
	"math/rand/v2"
)

// Returns an allegedly-random number between 'from' and 'to' inclusive.
func Random(from, to int32) int32 {
	if from > to {
		log.Panic("Random: from is greater than to")
	}
	if from == to {
		return from
	}
	return rand.Int32N((to-from)+1) + from
}
