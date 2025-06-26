package server

import (
	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/text"
)

// CmdExamine is used to get a description of a person, space_ship, object, or
// mobile.
func (p *Player) CmdExamine(name model.Name) {
	// Check for a player in the location.
	if !name.The && name.Words == 1 {
		if p.examinePlayer(name.Text) {
			return
		}
	}
	// Check for objects in the location.
	o, ok := p.curSys.FindObjectName(name)
	if ok {
		p.Output(o.Scan())
		return
	}
	// Check for objects in the player's inventory.
	o, ok = p.FindInventoryName(name)
	if ok {
		p.Output(o.Scan())
		return
	}
	// Didn't find anything.
	p.Outputm(text.MN107)
}

func (p *Player) examinePlayer(subjectName string) bool {
	debug.Precondition(subjectName != "")

	subject, ok := FindPlayer(subjectName)
	if !ok {
		return false
	}

	if subject == p {
		p.Outputm(text.MN760)
	} else {
		if !subject.IsPlaying() {
			return false
		}
		if subject.curLoc != p.curLoc {
			return false
		}
		if p.IsInsideSpaceship() {
			return false
		}
	}

	switch {
	case subject.IsDressed():
		p.Output(subject.Desc)
		p.Output("\n")
	case subject.Deaths() > 0:
		p.Outputm(text.HospitalClothes)
	default:
		p.Outputm(text.DefaultClothes)
	}

	if subject == p {
		return true
	}

	items := make([]string, len(subject.inventory))
	for i, o := range subject.inventory {
		items[i] = o.DisplayName(false)
	}
	if len(items) > 0 {
		switch subject.Sex() {
		case model.SexFemale:
			p.Output("She is carrying ")
		case model.SexMale:
			p.Output("He is carrying ")
		default:
			p.Output("It is carrying ")
		}
		p.Output(text.ListOfObjects(items))
		p.Output("\n")
	}

	subject.Outputm(text.MN759, p.Name())
	subject.FlushOutput()

	return true
}

func (p *Player) CmdExamineSpaceship(subjectName string) {
	subject, ok := FindPlayer(subjectName)
	if !ok || !subject.IsPlaying() {
		p.Outputm(text.MN107)
		return
	}
	if subject.HasSpaceship() {
		if subject.CurSys() == p.CurSys() {
			if subject.ShipLocNo() == p.CurLocNo() || subject.ShipLocNo() == p.ShipLocNo() {
				if !p.IsInHorsellSystem() {
					p.Output(subject.ShipDesc() + "\n")
					return
				}
			}
		}
	}
	if subject == p {
		p.Output("Your spaceship isn't here!\n")
	} else {
		p.outputf("%s's spaceship isn't here!\n", subject.Name())
	}
}
