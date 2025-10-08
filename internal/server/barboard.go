package server

import (
	"fmt"

	"github.com/nosborn/federation-1999/internal/text"
)

var board struct {
	nmessage int
	messages [20]struct {
		stcky    bool
		stardate string
		poster   string
		message  string
	}
}

func PostBarBoardMessage(p *Player) {
}

func ReadBarBoard(p *Player, login bool) string {
	var b []byte

	if board.nmessage == 0 {
		if !login {
			b = fmt.Append(b, text.Msg(text.Barboard5))
		}
		return string(b)
	}

	b = fmt.Append(b, text.Msg(text.Barboard6))

	for i := range board.nmessage {
		b = fmt.Appendf(b, "%s - %s: %s\n",
			board.messages[i].stardate,
			board.messages[i].poster,
			board.messages[i].message)
	}

	if login { // Do we even need it then?
		b = fmt.Append(b, "\n")
	}

	return string(b)
}

func UnpostBarBoardMessage(p *Player, poster string) string {
	// TODO
	return ""
}
