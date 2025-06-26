package server

import (
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

// CmdWait lets the player wait for public transport to arrive.
func (p *Player) CmdWait() {
	p.Outputm(text.Wait)

	// Only move the shuttle if the player is in the shuttle or at a
	// station.
	if p.CurSys().IsSol() {
		switch p.LocNo {
		case sol.Shuttle, sol.ShuttleStation1, sol.ShuttleStation2, sol.ShuttleStation3, sol.ShuttleStation4:
			p.Output("\n")
			SolSystem.ShuttleTimerProc()
		}
	}
}
