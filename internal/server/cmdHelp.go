package server

import (
	"github.com/nosborn/federation-1999/internal/text"
)

func (p *Player) CmdHelp(topic int32, subTopic int32) {
	if help(p, topic, subTopic) {
		return
	}
	if topic > 0 {
		p.Nsoutputm(text.Help_NoTopic)
	}
	p.Nsoutputm(text.Help_Default)
}
