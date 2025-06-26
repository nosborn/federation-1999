package model

type ButtonColour int

const (
	ColourBlue ButtonColour = iota
	ColourBrown
	ColourOrange
)

type GambleColour int

const (
	ColourBlack GambleColour = iota
	ColourRed
)

type InsureAction int

const (
	InsureGetQuote = iota
	InsureBuyPolicy
)

type BayList struct {
	Size int
	Bay  [5]int32
}

type Name struct {
	The   bool
	Words int
	Text  string
}
