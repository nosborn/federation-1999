package server

import (
	"bytes"
	"strings"

	"github.com/nosborn/federation-1999/internal/collections"
	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/core"
	"github.com/nosborn/federation-1999/internal/text"
	"github.com/nosborn/federation-1999/pkg/ibgames"
)

type Describe int

const (
	DefaultDescription Describe = iota
	BriefDescription
	LongDescription
	FullDescription
)

type LocationMovTab [13]int

type Location struct {
	Events      [2]uint32
	MovTab      [13]uint32
	briefDesc   string
	briefMsgNo  text.MsgNum
	flags       uint32
	fullDesc    string
	fullMsgNo   text.MsgNum
	number      uint32
	players     collections.OrderedCollection[*Player]
	sysLoc      string
	sysLocMsgNo text.MsgNum
	system      Systemer
}

func NewCoreLocation(coreLoc *core.Location, system Systemer) *Location {
	l := &Location{
		Events: [2]uint32{uint32(coreLoc.Events[0]), uint32(coreLoc.Events[1])},
		MovTab: [13]uint32{
			uint32(coreLoc.MovTab[0]), uint32(coreLoc.MovTab[1]), uint32(coreLoc.MovTab[2]),
			uint32(coreLoc.MovTab[3]), uint32(coreLoc.MovTab[4]), uint32(coreLoc.MovTab[5]),
			uint32(coreLoc.MovTab[6]), uint32(coreLoc.MovTab[7]), uint32(coreLoc.MovTab[8]),
			uint32(coreLoc.MovTab[9]), uint32(coreLoc.MovTab[10]), uint32(coreLoc.MovTab[11]),
			uint32(coreLoc.MovTab[12]),
		},
		briefDesc: text.Msg(coreLoc.BriefMsgNo),
		flags:     coreLoc.Flags,
		fullDesc:  text.Msg(coreLoc.FullMsgNo),
		number:    coreLoc.Number,
		system:    system,
	}
	if coreLoc.SysLoc != 0 {
		l.sysLoc = text.Msg(coreLoc.SysLoc)
	}
	return l
}

func NewLocationFromDB(locNo uint32, dbLoc model.DBLocation, system *PlayerSystem) *Location {
	return &Location{
		Events: [2]uint32{uint32(dbLoc.Events[0]), uint32(dbLoc.Events[1])},
		MovTab: [13]uint32{
			uint32(dbLoc.MovTab[0]), uint32(dbLoc.MovTab[1]), uint32(dbLoc.MovTab[2]),
			uint32(dbLoc.MovTab[3]), uint32(dbLoc.MovTab[4]), uint32(dbLoc.MovTab[5]),
			uint32(dbLoc.MovTab[6]), uint32(dbLoc.MovTab[7]), uint32(dbLoc.MovTab[8]),
			uint32(dbLoc.MovTab[9]), uint32(dbLoc.MovTab[10]), uint32(dbLoc.MovTab[11]),
			uint32(dbLoc.MovTab[12]),
		},
		briefDesc: text.CStringToString(dbLoc.Desc[:bytes.IndexByte(dbLoc.Desc[:], '\n')+1]),
		flags:     dbLoc.MapFlag,
		fullDesc:  text.CStringToString(dbLoc.Desc[:]),
		number:    locNo,
		sysLoc:    text.CStringToString([]byte{byte(dbLoc.SysLoc)}), // Convert message number to string
		system:    system,
	}
}

func (l *Location) BriefDesc() string {
	if l.briefMsgNo > 0 {
		return text.Msg(l.briefMsgNo)
	}
	debug.Check(l.briefDesc != "")
	return l.briefDesc
}

func (l *Location) Clean(cleaner *Object) {
	var cleaned []*Object

	for _, o := range l.system.Objects() {
		if o.curLocNo != l.number {
			continue
		}
		if o.IsMobile() {
			continue
		}
		if o.IsHidden() || o.IsRecycling() {
			continue
		}
		if o.isDisplaced(l) {
			cleaned = append(cleaned, o)
			o.Recycle()
		}
	}
	if len(cleaned) == 0 {
		return
	}

	var sb strings.Builder
	sb.WriteString(text.Msg(text.CleanerPreamble, cleaner.DisplayName(true)))
	for i, o := range cleaned[:len(cleaned)-2] {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(o.DisplayName(false))
	}
	if len(cleaned) > 1 {
		if len(cleaned) > 2 {
			sb.WriteByte(',')
		}
		sb.WriteString(" and ")
		sb.WriteString(cleaned[len(cleaned)-1].DisplayName(false))
	}
	sb.WriteByte('\n')
	l.Talk(sb.String())
}

func (l *Location) Describe(caller *Player, level Describe) {
	debug.Check(l != nil)
	debug.Check(caller != nil)

	// There should always be a blank line before a location description.
	caller.Output("\n")

	if !caller.IsFlyingSpaceship() {
		// Unlit location and the player doesn't have a lamp?
		if l.IsDark() && !caller.HasLamp() {
			caller.Outputm(text.UnlitLocation)
			return
		}
	}

	var briefDesc, briefShips bool
	switch level {
	case BriefDescription:
		briefDesc = true
		briefShips = true
	case LongDescription:
		briefDesc = false
		briefShips = true
	case FullDescription:
		briefDesc = false
		briefShips = false
	default:
		briefDesc = caller.WantsBrief()
		briefShips = true
	}

	// Show the location description.
	if briefDesc && !l.IsDeath() {
		caller.Output(l.BriefDesc())
	} else {
		caller.Output(l.FullDesc())
		caller.Output("\n") // should this be elsewhere?
	}
	// log.Printf("%#v", l) // FIXME

	// There can't possibly be any players/mobiles/objects here if we're
	// inside the spaceship but not in the command centre.
	if caller.IsInsideSpaceship() && !caller.IsFlyingSpaceship() {
		// There might be a viewscreen though.
		if caller.CurLocNo() != 1 { // FIXME
			return
		}
		planet := caller.GuessCurrentPlanet()
		if planet == nil {
			caller.Outputm(text.MN844)
			return
		}
		caller.Outputm(text.MN845, planet.Name())
		return
	}

	l.ListObjects(caller)

	if caller.IsFlyingSpaceship() {
		for op := range l.players.Values() {
			if op == caller {
				continue
			}
			caller.Outputm(text.ShipIsHere, op.Name(), GetShipClass(op.ShipKit.Tonnage))
		}
		return
	}

	if caller.IsInsideSpaceship() {
		return
	}

	if l.IsLandingPad() {
		nShips := 0
		ownShip := false

		for _, op := range Players {
			if !op.IsPlaying() {
				continue
			}
			if !op.HasSpaceship() {
				continue
			}
			if op.CurSys() != caller.CurSys() || op.ShipLocNo() != caller.CurLocNo() {
				continue
			}
			if briefShips {
				nShips++
				if op == caller {
					ownShip = true
				}
				continue
			}
			if op == caller {
				caller.Outputm(text.MN1528)
				continue
			}
			caller.Outputm(text.MN743, op.Name())
		}

		if briefShips && nShips > 0 {
			if ownShip {
				switch nShips {
				case 1:
					caller.Outputm(text.MN1528)
				case 2:
					caller.Outputm(text.MN1529)
				default:
					caller.Outputm(text.MN1530, nShips-1)
				}
			} else {
				switch nShips {
				case 1:
					caller.Outputm(text.MN1531)
				default:
					caller.Outputm(text.MN1532, nShips)
				}
			}
		}
	}

	// Check for other players.
	for op := range l.players.Values() {
		if op == caller {
			continue
		}
		caller.Outputm(text.MN744, op.MoodAndName())
	}
}

func (l *Location) Destroy() {
	// TODO
}

func (l *Location) EndTimewarp(warper ibgames.AccountID) {
	for player := range l.players.Values() {
		if (player.flags2 & PL2_TIMEWARPED) != 0 {
			EndTimewarp(player, warper)
		}
	}

	// // Check for a case of the runs!
	// // FIXME: Do this for return from Snark too.
	// for (PCI iter = m_players.begin(); iter != m_players.end(); iter++) {
	// 	if (((*iter)->m_flags2 & PL2_WHOOSH) != 0) {
	// 		(*iter)->drinkWHOOSH(0);
	// 		break;
	// 	}
	// }
}

func (l *Location) FindHelper(objectID uint32) *Player {
	for player := range l.players.Values() {
		_, ok := player.FindInventoryID(objectID)
		if ok {
			return player
		}
	}
	return nil
}

func (l *Location) FindObject(num uint32) (*Object, bool) {
	for _, o := range l.system.Objects() {
		if o.CurLocNo() == l.number {
			if o.IsHidden() || o.IsRecycling() {
				continue
			}
			if o.number == num {
				return o, true
			}
		}
	}
	return nil, false
}

func (l *Location) FindObjectName(name model.Name) (*Object, bool) {
	for _, o := range l.system.Objects() {
		if o.CurLocNo() == l.number {
			if o.IsHidden() || o.IsRecycling() {
				continue
			}
			if name.The && o.noThe() {
				continue
			}
			if strings.EqualFold(o.Name(), name.Text) {
				return o, true
			}
			if name.Words != 1 {
				continue
			}
			for _, synonym := range o.Synonyms {
				if synonym == "" {
					continue
				}
				if strings.EqualFold(synonym, name.Text) {
					return o, true
				}
			}
		}
	}
	return nil, false
}

func (l *Location) FindPlayer(name string) (*Player, bool) {
	for p := range l.players.Values() {
		if strings.EqualFold(p.name, name) {
			return p, true
		}
	}
	return nil, false
}

func (l *Location) FindPlayerByID(uid ibgames.AccountID) (*Player, bool) {
	for p := range l.players.Values() {
		if p.uid == uid {
			return p, true
		}
	}
	return nil, false
}

func (l *Location) FullDesc() string {
	if l.fullMsgNo > 0 {
		return text.Msg(l.fullMsgNo)
	}
	debug.Check(l.fullDesc != "")
	return l.fullDesc
}

func (l *Location) InsertPlayer(p *Player) {
	l.players.Insert(p)
}

func (l *Location) IsCafe() bool {
	return (l.flags & model.LfCafe) != 0
}

func (l *Location) IsClothingStore() bool {
	return (l.flags & model.LfClth) != 0
}

func (l *Location) IsDark() bool {
	return (l.flags & model.LfDark) != 0
}

func (l *Location) IsDeath() bool {
	return (l.flags & model.LfDeath) != 0
}

func (l *Location) IsElectronicsStore() bool {
	return (l.flags & model.LfCom) != 0
}

func (l *Location) IsExchange() bool {
	return (l.flags & model.LfTrade) != 0
}

func (l *Location) IsGeneralStore() bool {
	return (l.flags & model.LfGen) != 0
}

func (l *Location) IsHidden() bool {
	return (l.flags & model.LfHidden) != 0
}

func (l *Location) IsHospital() bool {
	return (l.flags & model.LfHospital) != 0
}

func (l *Location) IsIndoors() bool {
	return (l.flags & model.LfIndoors) != 0
}

func (l *Location) IsInsuranceBroker() bool {
	return (l.flags & model.LfIns) != 0
}

func (l *Location) IsLandingPad() bool {
	return (l.flags & model.LfLanding) != 0
}

func (l *Location) IsLink() bool {
	return (l.flags & model.LfLink) != 0
}

func (l *Location) IsLocked() bool {
	return (l.flags & model.LfLock) != 0
}

func (l *Location) IsOutdoors() bool {
	return (l.flags & model.LfOutdoors) != 0
}

func (l *Location) IsPeaceful() bool {
	return (l.flags & model.LfPeace) != 0
}

func (l *Location) IsOrbit() bool {
	return (l.flags & model.LfOrbit) != 0
}

func (l *Location) IsRepairShop() bool {
	return (l.flags & model.LfRep) != 0
}

func (l *Location) IsShielded() bool {
	return (l.flags & model.LfShield) != 0
}

func (l *Location) IsShipyard() bool {
	return (l.flags & model.LfYard) != 0
}

func (l *Location) IsSpace() bool {
	return (l.flags & model.LfSpace) != 0
}

func (l *Location) IsWeaponsShop() bool {
	return (l.flags & model.LfWeap) != 0
}

func (l *Location) ListObjects(caller *Player) {
	for _, o := range l.system.Objects() {
		if o.CurLocNo() == l.number {
			if o.IsHidden() || o.IsRecycling() {
				continue
			}
			caller.Output(o.Desc())
		}
	}
}

func (l *Location) noExitMessage() string {
	if l.sysLocMsgNo > 0 {
		return text.Msg(l.sysLocMsgNo)
	}
	if l.sysLoc != "" {
		return l.sysLoc
	}
	if l.IsSpace() {
		return text.Msg(text.NO_MOVEMENT_2)
	}
	return text.Msg(text.NO_MOVEMENT_1)
}

func (l *Location) Number() uint32 {
	return l.number
}

func (l *Location) RemovePlayer(p *Player) {
	l.players.Remove(p)
}

func (l *Location) SetBriefMsgNo(v text.MsgNum) {
	l.briefMsgNo = v
}

func (l *Location) SetFullMsgNo(v text.MsgNum) {
	l.fullMsgNo = v
}

func (l *Location) Talk(message string, omit ...*Player) {
players:
	for p := range l.players.Values() {
		for i := range omit {
			if p == omit[i] {
				continue players
			}
		}
		p.Output(message)
		if !p.isActiveThread() {
			p.FlushOutput()
		}
	}
}
