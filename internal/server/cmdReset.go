package server

import (
	"log"
	"syscall"

	"github.com/nosborn/federation-1999/internal/server/global"
)

func (p *Player) CmdReset() {
	if !global.TestFeaturesEnabled {
		p.UnknownCommand()
		return
	}

	log.Printf("RESET requested by %s", p.Name())

	// TODO: make unspyable

	if err := syscall.Kill(syscall.Getpid(), syscall.SIGTERM); err != nil {
		log.Printf("syscall.Kill() failed: %v", err)
		p.Output("I don't seem to be able to do that!\n")
		return
	}

	p.Output("You've killed Federation! You bastard!\n")
}
