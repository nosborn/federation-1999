package model

type PlanetJob struct { // used to store owner generated milkruns
	Name      string    // planet to deliver to
	Commodity Commodity // type of goods to deliver
	Carriage  int32     // IG/ton for the hauler
}
