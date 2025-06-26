package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/pkg/ibgames"
)

func (p *Player) CmdWhoIsName(subjectName string) {
	subject, _ := FindPlayer(subjectName)
	p.whoIs(subject)
}

func (p *Player) CmdWhoIsUID(subjectUID ibgames.AccountID) {
	subject, _ := FindPlayerByID(subjectUID)
	p.whoIs(subject)
}

func (p *Player) whoIs(subject *Player) {
	if p.Rank() < model.RankManager || p.Rank() > model.RankDeity {
		p.UnknownCommand()
		return
	}
	if subject == nil {
		p.Output("I haven't the faintest idea!\n")
		return
	}
	p.outputf("Persona name: %s\n fed99 UID: %d\n", subject.Name(), subject.UID())
	if p.Rank() == model.RankDeity && subject.Session() != nil {
		// There's probably no use for this except debugging.
		p.outputf(" Remote host: %s\n", subject.Session().RemoteHostname())
	}
}
