package server

import (
	"log"
	"strings"
)

type WorkbenchDriver struct {
	player  *Player
	session *Session
}

func NewWorkbenchDriver(s *Session, p *Player) *WorkbenchDriver {
	wd := WorkbenchDriver{
		player:  p,
		session: s,
	}
	return &wd
}

func (wd *WorkbenchDriver) Destroy() {
	// TODO
}

func (wd *WorkbenchDriver) Dispatch(line string) bool {
	// input := CleanInput(line, noNL) // TODO
	input := strings.TrimSpace(line)
	_ = input // FIXME

	log.Printf("Bad state in WorkbenchDriver.Dispatch")
	return false
}
