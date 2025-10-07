package server

import "github.com/nosborn/federation-1999/internal/server/parser"

func help(player *Player, topic, subTopic int32) bool {
	for i := range parser.HelpEntries {
		if parser.HelpEntries[i].Topic == topic && parser.HelpEntries[i].SubTopic == subTopic {
			player.Nsoutputm(parser.HelpEntries[i].MsgID)
			return true
		}
	}

	for i := range parser.HelpEntries {
		if parser.HelpEntries[i].Topic == topic {
			player.Nsoutputm(parser.HelpEntries[i].MsgID)
			return true
		}
	}

	return false
}
