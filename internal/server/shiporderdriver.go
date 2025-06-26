package server

import (
	"log"
	"strings"
)

type ShipOrderDriver struct {
	player  *Player
	session *Session
}

func NewShipOrderDriver(s *Session, p *Player) *ShipOrderDriver {
	sod := ShipOrderDriver{
		player:  p,
		session: s,
	}
	return &sod
}

func (sod *ShipOrderDriver) Destroy() {
	// TODO
}

func (sod *ShipOrderDriver) Dispatch(line string) bool {
	// input := CleanInput(line, noNL) // TODO
	input := strings.TrimSpace(line)
	_ = input // FIXME

	log.Printf("Bad state in ShipOrderDriver.Dispatch")
	return false
}
