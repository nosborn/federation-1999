package server

import (
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
)

// Allows player to give money to one another - or to the mobiles.
func (p *Player) CmdGiveGroats(groats int32, name model.Name) {
	if groats <= 0 {
		p.Outputm(text.NiceTry)
		return
	}
	if !name.The && name.Words == 1 {
		if p.giveGroatsToPlayer(groats, name.Text) {
			return
		}
	}
	if p.IsInsideSpaceship() {
		p.Outputm(text.MN537)
		return
	}
	if p.giveGroatsToMobile(groats, name) {
		return
	}
	p.Outputm(text.MN537)
}

func (p *Player) giveGroatsToPlayer(groats int32, name string) bool {
	// TODO
	return false // FIXME: true
}

func (p *Player) giveGroatsToMobile(groats int32, name model.Name) bool {
	// TODO
	return false // FIXME: true
}

// Allows player to give objects to one another - or to the mobiles.
func (p *Player) CmdGiveObject(name1 model.Name, name2 model.Name) {
	object, ok := p.FindInventoryName(name1)
	if !ok {
		p.Outputm(text.MN528)
		return
	}
	if p.IsInsideSpaceship() {
		p.Outputm(text.MN537)
		return
	}
	if !name2.The && name2.Words == 1 {
		if p.giveObjectToPlayer(object, name2.Text) {
			return
		}
	}
	if p.giveObjectToMobile(object, name2) {
		return
	}
	p.Outputm(text.MN537)
}

func (p *Player) giveObjectToPlayer(object *Object, toName string) bool {
	toPlayer, ok := FindPlayer(toName)
	if !ok || !toPlayer.IsPlaying() {
		return false
	}
	if toPlayer.curLoc != p.curLoc {
		return false
	}

	if toPlayer == p {
		p.Outputm(text.DontBeSilly)
		return true
	}

	// Can the recipient carry it?
	if !toPlayer.canCarry(object.weight) {
		p.Outputm(text.MN543, toPlayer.Name())
		return true
	}

	p.RemoveFromInventory(object)
	toPlayer.AddToInventory(object)
	toPlayer.Outputm(text.MN544, p.Name(), object.DisplayName(false))

	if object.Number() == sol.ObBlackBox {
		toPlayer.checkForBlackBox()
	}

	toPlayer.FlushOutput()

	p.Outputm(text.MN545, object.DisplayName(false))
	return true
}

func (p *Player) giveObjectToMobile(object *Object, name model.Name) bool {
	// TODO
	return false // FIXME: true
}
