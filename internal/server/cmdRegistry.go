package server

import (
	"github.com/nosborn/federation-1999/internal/text"
)

// Displays ship registrations for all players online.
func (p *Player) CmdRegistry() {
	p.Nsoutputm(text.RegistryHeader)

	for _, subject := range Players {
		if !subject.HasSpaceship() {
			continue
		}
		p.Nsoutputm(text.RegistryEntry, subject.Name, subject.Registry)
	}
}
