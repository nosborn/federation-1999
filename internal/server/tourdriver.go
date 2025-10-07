package server

import (
	"log"
	"strings"
)

type TourDriver struct {
	player  *Player
	session *Session
}

func NewTourDriver(s *Session, p *Player) *TourDriver {
	td := TourDriver{
		player:  p,
		session: s,
	}
	return &td
}

func (td *TourDriver) Destroy() {
	// TODO
}

func (td *TourDriver) Dispatch(line string) bool {
	// input := CleanInput(line, noNL) // TODO
	input := strings.TrimSpace(line)
	_ = input // FIXME

	log.Printf("Bad state in TourDriver.Dispatch")
	return false
}
