package server

import (
	"log"
	"strings"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/monitoring"
	"github.com/nosborn/federation-1999/internal/server/global"
)

type PlayDriver struct {
	player  *Player
	session *Session
}

func NewPlayDriver(s *Session, p *Player) *PlayDriver {
	global.Lock()
	defer global.Unlock()

	d := PlayDriver{
		player:  p,
		session: s,
	}

	p.EnterGame(s)

	// g_playing[m_player->name()] = m_player
	if p.Rank() < model.RankSenator {
		monitoring.PlayersCurrent.WithLabelValues("fedtpd").Inc()
		// if currentPlayers > peakPlayers {
		// 	peakPlayers = currentPlayers
		// }
	}

	return &d
}

func (d *PlayDriver) Destroy() {
	log.Printf("PlayDriver.Destroy(%s)", d.player.Name())

	global.Lock()
	defer global.Unlock()

	d.player.LeaveGame()

	// g_playing.erase(m_player->name());
	if d.player.Rank() < model.RankSenator {
		monitoring.PlayersCurrent.WithLabelValues("fedtpd").Dec()
	}

	d.player = nil
	d.session = nil
}

func (d *PlayDriver) Dispatch(line string) bool {
	global.Lock()
	defer global.Unlock()

	line = strings.TrimSpace(line)
	d.player.outputf(">%s\n", line)
	if line != "" {
		d.player.DoCommand(line)
	}
	d.player.FlushOutput()
	return true
}
