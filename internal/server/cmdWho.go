package server

import (
	"fmt"
	"strings"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/text"
)

// Compiles a list of who is currently playing in the game.
func (p *Player) CmdWho() {
	playing := 0
	for _, op := range Players {
		if op.IsPlaying() {
			p.Nsoutput(op.whoEntry())
			playing++
		}
	}
	if playing == 1 {
		p.Nsoutputm(text.WHO_1_PLAYER)
	} else {
		p.Nsoutputm(text.WHO_ALL_PLAYERS, playing)
	}
}

// Compiles a list of who is currently playing in the game.
func (p *Player) CmdWhoChannel(channel int32) {
	if channel < 0 || channel > model.MAX_XT_CHANNEL {
		p.Nsoutputm(text.WHO_BAD_CHANNEL)
		return
	}

	p.Output("Not implemented. Check back in 2 weeks.\n")
}

// Compiles a list of who is currently playing in the game.
func (p *Player) CmdWhoSystem(sysName string) {
	s, ok := FindSystem(sysName)
	if !ok || (s.IsHidden() && p.Rank() < model.RankHostess) {
		p.Nsoutputm(text.WHO_BAD_SYSTEM)
		return
	}

	p.Output("Not implemented. Check back in 2 weeks.\n")
}

func (p *Player) whoEntry() string {
	var sb strings.Builder

	if p.IsCommsOff() {
		sb.WriteByte('-')
	} else {
		sb.WriteByte(' ')
	}

	if p.HasCommUnit() && p.channel > 0 && p.channel <= model.MAX_XT_CHANNEL {
		fmt.Fprintf(&sb, "%-2d ", p.channel)
	} else {
		sb.WriteString("  ")
	}
	sb.WriteByte(' ')

	if p.IsOnDutyNavigator() {
		fmt.Fprintf(&sb, "%s - DataSpace Navigator", p.name)
	} else {
		switch p.CustomRank {
		case CustomRankTheVile:
			sb.WriteString(text.Msg(text.TheVile, p.name))
		case CustomRankTheDemiGoddess:
			sb.WriteString(text.Msg(text.TheDemiGoddess, p.name))
		case CustomRankOfTheSpaceways:
			sb.WriteString(text.Msg(text.OfTheSpaceways, p.name))
		default:
			switch p.Rank() {
			case model.RankSquire, model.RankThane, model.RankIndustrialist, model.RankTechnocrat, model.RankBaron, model.RankDuke:
				fmt.Fprintf(&sb, "%s, %s of %s", p.name, p.rankName(), p.OwnSystem().Name())
			case model.RankSenator:
				sb.WriteString(text.Msg(text.TheDishonourable, p.name))
			case model.RankHostess:
				fmt.Fprintf(&sb, "%s - %s", p.name, p.rankName())
			case model.RankManager, model.RankDeity:
				fmt.Fprint(&sb, p.name)
			case model.RankEmperor:
				sb.WriteString(text.Msg(text.HisImperialMajestyThe, p.rankName(), p.name))
			default:
				fmt.Fprintf(&sb, "%s %s", p.rankName(), p.name)
			}
		}
	}

	sb.WriteByte('\n')
	return sb.String()
}
