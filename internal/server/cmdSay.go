package server

import (
	"math/rand/v2"
	"strings"

	"github.com/nosborn/federation-1999/internal/text"
)

var (
	ask = [2][2]text.MsgNum{
		{text.YouAsk_1, text.PlayerAsks_1},
		{text.YouAsk_2, text.PlayerAsks_2},
	}
	exclaim = [2][2]text.MsgNum{
		{text.YouExclaim_1, text.PlayerExclaims_1},
		{text.YouExclaim_2, text.PlayerExclaims_2},
	}
	frown = [2][2]text.MsgNum{
		{text.YouFrown_1, text.PlayerFrowns_1},
		{text.YouFrown_2, text.PlayerFrowns_2},
	}
	say = [2][2]text.MsgNum{
		{text.YouSay_1, text.PlayerSays_1},
		{text.YouSay_2, text.PlayerSays_2},
	}
	smile = [2][2]text.MsgNum{
		{text.YouSmile_1, text.PlayerSmiles_1},
		{text.YouSmile_2, text.PlayerSmiles_2},
	}
	shout = [2][2]text.MsgNum{
		{text.YouShout_1, text.PlayerShouts_1},
		{text.YouShout_2, text.PlayerShouts_2},
	}
	wink = [2][2]text.MsgNum{
		{text.YouWink_1, text.PlayerWinks_1},
		{text.YouWink_2, text.PlayerWinks_2},
	}
)

// Passes messages to players in the same location.
func (p *Player) CmdSay(speech string) {
	if p.IsSulking() {
		p.Outputm(text.SULKING_REMINDER)
		return
	}

	length := len(speech)
	// char* theSpeech = strcpy(static_cast<char*>(alloca(length + 1)), text);
	// speech := raw

	qualifier := rand.IntN(2)
	msgNums := say[qualifier]

	switch {
	case text.IsShouting(speech):
		msgNums = shout[qualifier]
	case speech[length-1] == '?':
		msgNums = ask[qualifier]
	case speech[length-1] == '!':
		msgNums = exclaim[qualifier]
	case len(speech) >= 2:
		switch {
		case strings.HasSuffix(speech, ":("):
			msgNums = frown[qualifier]
			speech = speech[:len(speech)-2]
		case strings.HasSuffix(speech, ":)"):
			msgNums = smile[qualifier]
			speech = speech[:len(speech)-2]
		case strings.HasSuffix(speech, ";)"):
			msgNums = wink[qualifier]
			speech = speech[:len(speech)-2]
		}
		speech = strings.TrimRight(speech, " ")
	}

	p.Outputm(msgNums[0], speech)

	//
	if !p.IsInsideSpaceship() {
		var msg string
		if (qualifier & 1) == 0 {
			msg = text.Msg(msgNums[1], p.Name(), speech)
		} else {
			msg = text.Msg(msgNums[1], speech, p.Name())
		}
		p.CurLoc().Talk(msg, p)
	}
}
