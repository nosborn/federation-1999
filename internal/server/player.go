package server

import (
	"fmt"
	"log"
	"math/rand/v2"
	"slices"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/nosborn/federation-1999/internal/debug"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/build"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/global"
	"github.com/nosborn/federation-1999/internal/server/goods"
	"github.com/nosborn/federation-1999/internal/server/jobs"
	"github.com/nosborn/federation-1999/internal/server/parser"
	"github.com/nosborn/federation-1999/internal/server/sol"
	"github.com/nosborn/federation-1999/internal/text"
	"github.com/nosborn/federation-1999/internal/workbench"
	"github.com/nosborn/federation-1999/pkg/ibgames"
	"github.com/nosborn/federation-1999/pkg/ibgames/rules"
)

// PlayerSaverFunc defines the function signature for saving a player.
type PlayerSaverFunc func(p *Player, when database.SaveWhen)

// defaultPlayerSaver is the production implementation of PlayerSaverFunc.
func defaultPlayerSaver(p *Player, when database.SaveWhen) {
	// We might be called for a player who isn't in the game; this can
	// happen with planet owners and Dukes in particular.
	if p.Session() != nil {
		if !p.Session().BillingTick() {
			log.Panic("PANIC: save: billingTick() failed")
		}
		if p.IsPlaying() {
			p.LastOn = int32(time.Now().Unix()) // Transaction::time();
		}
	}
	database.Modify(p, when)
}

type (
	PlayerCustomRank int8
)

const (
	CustomRankTheVile PlayerCustomRank = 1 << iota
	CustomRankTheDemiGoddess
	CustomRankOfTheSpaceways
)

type PlayerStat struct {
	Max int32
	Cur int32
}

const (
	spyPublic  = 0
	spyPrivate = (spyPublic - 1)
)

const (
	PL2_IN_PARSER       = 0x00000001 // Current player in the parser
	PL2_DECEASED        = 0x00000002 // Player is really dead
	PL2_IN_JOB_ADVERT   = 0x00000004 // Seen first in batch of new jobs
	PL2_ON_DUTY_NAV     = 0x00000008 // On-duty DataSpace Navigator
	PL2_COMMS_OFF       = 0x00000010 // Has comms turned off
	PL2_REWARD          = 0x00000020 // Has posted reward
	PL2_DNI_PASSWORD_OK = 0x00000040 // Entered correct DNI password
	PL2_TIMEWARPED      = 0x00000080 // Is being timewarped
	PL2_CORPSE          = 0x00000100 // Player is temporarily dead (!)
	PL2_WHOOSH          = 0x00000200 // Drank the WHOOSH and needs the loo!
)

const (
	tickerTimerPeriod  = 11 * time.Second
	tourismTimerPeriod = 60 * time.Second
)

type Player struct {
	name string

	uid ibgames.AccountID

	offset database.Offset // Offset into the persona database

	sex        model.Sex
	rank       model.Rank
	CustomRank PlayerCustomRank

	Desc string
	mood string

	// ObjectList

	Str     PlayerStat
	Sta     PlayerStat
	Dex     PlayerStat
	Int     PlayerStat
	Shipped uint32
	Games   uint32
	Flags0  uint32
	Flags1  uint32
	flags2  uint32

	balance      int32
	loan         int32
	reward       int32
	deaths       int32
	TradeCredits int32

	curSysName string
	curSys     Systemer
	LocNo      uint32
	curLoc     *Location

	Job model.Work

	ComType       model.CommodityGroup
	NextCommodity model.Commodity
	LastTrade     int

	company *Company
	deal    *Contract

	zoneReq *Factory

	Facilities [7]int32

	ShipLoc uint32

	shipDesc string
	Registry string

	ShipKit  Equipment
	Missiles int32
	Ammo     int32
	Guns     [model.MAX_GUNS]model.SGuns
	Load     [model.MAX_LOAD]model.Cargo

	Target ibgames.AccountID

	inventory []*Object

	channel    int32
	LastOn     int32
	lastMayday uint32

	Crypto        int
	Count         [4]int32
	gmLocation    uint32
	BuildProject  model.Project
	BuildDuration int32
	BuildElapsed  int32
	Warper        ibgames.AccountID

	storage   *Storage
	ownSystem Systemer // FIXME: *PlayerSystem
	ownDuchy  *Duchy
	session   *Session

	MsgOut         strings.Builder
	MsgOutSpyDepth int

	spied  *Player
	spyers []*Player

	alarmTimer   *time.Timer
	buildTimer   *time.Timer
	tickerTimer  *time.Timer
	tourismTimer *time.Timer

	saveFunc PlayerSaverFunc
}

var commDisplay int

func FindPlayer(name string) (*Player, bool) {
	for i := range Players {
		if strings.EqualFold(Players[i].name, name) {
			return Players[i], true
		}
	}
	return nil, false
}

func FindPlayerByID(uid ibgames.AccountID) (*Player, bool) {
	for _, p := range Players {
		if p.uid == uid {
			return p, true
		}
	}
	return nil, false
}

func (p *Player) AddSpyer(spyer *Player) {
	p.spyers = append(p.spyers, spyer)
	p.checkForBlackBox()
}

func (p *Player) AddToInventory(o *Object) {
	o.curLocNo = 0
	p.inventory = append(p.inventory, o)
}

func (p *Player) alarmTimerHandler() {
	global.Lock()
	defer global.Unlock()

	p.alarmTimerProc()
}

func (p *Player) alarmTimerProc() {
	p.alarmTimer = nil
	p.Outputm(text.AlarmNotification)
	p.FlushOutput()
}

func (p *Player) announce(msgID text.MsgNum) {
	var title string
	switch p.CustomRank {
	case CustomRankTheVile:
		title = text.Msg(text.TheVile, p.name)
	case CustomRankTheDemiGoddess:
		title = text.Msg(text.TheDemiGoddess, p.name)
	case CustomRankOfTheSpaceways:
		title = text.Msg(text.OfTheSpaceways, p.name)
	default:
		switch p.rank {
		case model.RankManager:
		case model.RankDeity:
			title = p.name
		case model.RankEmperor:
			debug.Check(p.sex == model.SexMale)
			title = text.Msg(text.HisImperialMajestyThe, p.rankName(), p.name)
		default:
			title = fmt.Sprintf("%s %s", p.rankName(), p.name)
		}
	}

	notify := text.Msg(msgID, title)
	for _, op := range Players {
		if !op.IsPlaying() {
			continue
		}
		if op == p {
			continue
		}
		if (op.Flags0 & model.PL0_INFO) != 0 {
			op.Output(notify)
			op.FlushOutput()
		}
	}
}

func (p *Player) Balance() int32 {
	return p.balance
}

func (p *Player) canCarry(weight uint32) bool {
	return (p.carriedWeight() + int32(weight)) <= (p.Str.Cur / 2)
}

func (p *Player) CanKnit() bool {
	if global.TestFeaturesEnabled {
		return p.rank == model.RankManager || p.rank == model.RankDeity
	}
	return p.rank == model.RankManager
}

func (p *Player) canSpy(subject *Player) bool {
	// Subject is a higher rank?
	if subject.Rank() > p.rank {
		return false
	}

	if p.rank < model.RankHostess {
		if subject.IsInHorsellSystem() || subject.IsInSnarkSystem() {
			return false
		}
		if subject.IsPromoCharacter() {
			return false
		}
	}

	// Spyer and subject are same rank? This is OK if the spyer is at the
	// public beam, but we have to take account of spybeam screens.
	//
	// Not allowing Barons to spy on each other is probably fairly
	// pointless. Then again, maybe it's not if we're going to have PROMO
	// characters at Baron.
	if subject.Rank() == p.rank {
		if p.IsInSolLocation(sol.SecretRoom) {
			if p.rank == model.RankBaron {
				return false
			}
			return !subject.HasSpyScreen()
		}
		return false
	}

	// The subject is spyable.
	return true
}

func (p *Player) CancelAlarm() {
	if p.alarmTimer != nil {
		p.alarmTimer.Stop()
		p.alarmTimer = nil
	}
}

func (p *Player) carriedWeight() int32 {
	var weight int32
	for i := range p.inventory {
		weight += int32(p.inventory[i].Weight())
	}
	return weight
}

func (p *Player) ChangeBalance(amount int32) {
	changeBalance(&p.balance, amount)
}

func (p *Player) ChangeReward(amount int32) {
	changeBalance(&p.reward, amount)
}

func (p *Player) Channel() int32 {
	return p.channel
}

func (p *Player) checkForBlackBox() {
	if !p.IsInSolSystem() {
		return
	}

	o, ok := p.FindInventoryID(sol.ObBlackBox)
	if !ok {
		return
	}

	var descMsgNo, scanMsgNo, msgNo text.MsgNum
	if p.isSpiedUpon() {
		descMsgNo = text.BlackBox_On_Desc
		scanMsgNo = text.BlackBox_On_Scan
		msgNo = text.BlackBoxRed
	} else {
		descMsgNo = text.BlackBox_Desc
		scanMsgNo = text.BlackBox_Scan
		msgNo = text.BlackBoxGreen
	}

	if o.Desc() == text.Msg(descMsgNo) {
		return
	}
	o.SetDesc(text.Msg(descMsgNo))
	o.SetScan(text.Msg(scanMsgNo))
	p.Nsoutputm(msgNo)
	p.FlushOutput()
}

func (p *Player) CheckForPromotion() {
	switch p.rank {
	case model.RankCommander: // -> Captain
		if p.loan > 0 {
			return
		}
		p.rank = model.RankCaptain
		p.loan = 0
	case model.RankCaptain: // -> Adventurer
		if p.TradeCredits < getConfig(CFG_ADVENTURER_TCR) {
			return
		}
		p.rank = model.RankAdventurer
		p.reward = 0
		p.Str.Max = min(p.Str.Max+5, 120)
		p.Str.Cur = min(p.Str.Cur+5, p.Str.Max)
		p.Sta.Max = min(p.Sta.Max+5, 120)
		p.Sta.Cur = min(p.Sta.Cur+5, p.Sta.Max)
		p.Int.Max = min(p.Int.Max+5, 120)
		p.Int.Cur = min(p.Int.Cur+5, p.Int.Max)
		p.Dex.Max = min(p.Dex.Max+5, 120)
		p.Dex.Cur = min(p.Dex.Cur+5, p.Dex.Max)
	case model.RankAdventurer: // -> Trader
		if p.TradeCredits < getConfig(CFG_TRADER_TCR) {
			return
		}
		p.setGMLocation()
		global.Haulers--
		if global.Haulers < 0 {
			log.Printf("Resetting global.Haulers (%d)", global.Haulers)
			global.Haulers = 0
		}
	case model.RankTrader: // -> Merchant
		if (p.Sta.Max + p.Str.Max + p.Dex.Max + p.Int.Max) < getConfig(CFG_MERCHANT_STATS) {
			return
		}
		p.rank = model.RankMerchant
		p.deaths = 0
	case model.RankMerchant: // -> JP
		if p.company == nil || p.company.Profit() < getConfig(CFG_JP_PROFIT) {
			return
		}
		p.rank = model.RankJP
		p.company.ClearBooks()
	case model.RankJP: // -> GM
		if p.company == nil || p.company.Profit() < getConfig(CFG_GM_PROFIT) {
			return
		}
		p.rank = model.RankGM
		p.company.ClearBooks()
	case model.RankGM: // -> Explorer
		if p.Str.Max < getConfig(CFG_EXPLORER_STAT) {
			return
		}
		if p.Sta.Max < getConfig(CFG_EXPLORER_STAT) {
			return
		}
		if p.Dex.Max < getConfig(CFG_EXPLORER_STAT) {
			return
		}
		if p.Int.Max < getConfig(CFG_EXPLORER_STAT) {
			return
		}
		if p.Shipped < uint32(getConfig(CFG_EXPLORER_TONS)) {
			return
		}
		p.rank = model.RankExplorer
		if p.company != nil {
			p.company.ClearBooks()
		}
	case model.RankExplorer: // -> Squire
		return // completePlanetLoad() takes care of this
	case model.RankSquire: // -> Thane
		if p.ownSystem == nil || p.ownSystem.Planets()[0].level <= model.LevelAgricultural {
			return
		}
		p.rank = model.RankThane
		p.Outputm(text.MN739, p.rankName(), p.ownSystem.Name())
	case model.RankThane: // -> Industrialist
		if p.ownSystem == nil || p.ownSystem.Planets()[0].level <= model.LevelMining {
			return
		}
		p.rank = model.RankIndustrialist
		p.Outputm(text.MN739, p.rankName(), p.ownSystem.Name())
	case model.RankIndustrialist: // -> Technocrat
		if p.ownSystem == nil || p.ownSystem.Planets()[0].level <= model.LevelIndustrial {
			return
		}
		p.rank = model.RankTechnocrat
		p.Outputm(text.MN739, p.rankName(), p.ownSystem.Name())
	case model.RankTechnocrat: // -> Baron
		if p.ownSystem == nil || p.ownSystem.Planets()[0].level <= model.LevelTechnological {
			return
		}
		p.rank = model.RankBaron
		p.Outputm(text.MN739, p.rankName(), p.ownSystem.Name())
	default:
		return
	}

	p.Save(database.SaveNow)
	p.CheckSpyers()
}

func (p *Player) CheckShields() {
	if (p.Flags1 & model.PL1_SHIELDS) == 0 {
		p.ShipKit.CurFuel -= min(p.ShipKit.CurFuel, 5)
		if p.ShipKit.CurFuel < 5 {
			p.Outputm(text.WarningBuzzer)
			p.Flags1 &^= model.PL1_SHIELDS
		}
	}
}

func (p *Player) CheckSpyers() {
	// TODO
}

func (p *Player) clearConstructionInfo() {
	p.BuildProject = build.NOTHING
	p.BuildDuration = 0
	p.BuildElapsed = 0
}

func (p *Player) ClearDNIPassword() {
	//
}

func (p *Player) ClearGMLocation() {
	p.gmLocation = 0
}

func (p *Player) clearInventory() {
	for _, o := range p.inventory {
		o.Recycle()
	}
	p.inventory = nil
}

// Involuntarily stop spying on the current subject.
func (p *Player) clearSpyer() {
	debug.Check(p.spied != nil)

	if p.spied == nil { // It shouldn't be!
		return
	}
	p.spied = nil
	p.Outputm(text.SPY_CLEARED)
	p.FlushOutput()
}

// Called by an exiting player to clear the spying and targeting fields of
// other players keyed to exiting player.
func (p *Player) clearSpyers() {
	if !p.IsPlaying() {
		log.Print("clearSpyers: Player is not playing!")
		return
	}

	for _, spyer := range p.spyers {
		spyer.clearSpyer()

		if spyer.Target == p.uid { // FIXME: this seems out of place
			spyer.Target = 0
		}
	}
	p.spyers = []*Player{}
}

func (p *Player) CurDex() int32 {
	return p.Dex.Cur
}

func (p *Player) CurLoc() *Location {
	return p.curLoc
}

func (p *Player) CurLocNo() uint32 {
	return p.LocNo
}

func (p *Player) CurSta() int32 {
	return p.Sta.Cur
}

func (p *Player) CurSys() Systemer {
	return p.curSys
}

func (p *Player) CurSysName() string {
	return p.curSysName
}

func (p *Player) Deaths() int32 {
	return p.deaths
}

func (p *Player) Deport() {
	if p.IsFlyingSpaceship() {
		p.Outputm(text.DeportInSpace)
	} else {
		p.Outputm(text.DeportOnGround)
	}
	p.Output("\n")

	if p.IsFlyingSpaceship() {
		if p.curLoc.IsLink() {
			message := text.Msg(text.ShipJumpsFrom, p.name, GetShipClass(p.ShipKit.Tonnage))
			p.curLoc.Talk(message, p)
		} else {
			message := text.Msg(text.ShipLeaves, p.name, GetShipClass(p.ShipKit.Tonnage))
			p.curLoc.Talk(message, p)
		}
	} else {
		if p.IsInsideSpaceship() {
			landingPad := p.CurSys().FindLocation(p.ShipLoc)
			debug.Check(landingPad != nil)
			message := text.Msg(text.MN561, p.name)
			landingPad.Talk(message, p)
		} else {
			message := text.Msg(text.PlayerHasLeft, p.name)
			p.curLoc.Talk(message, p)
		}
	}

	p.clearInventory()
	p.place()
	if p.IsFlyingSpaceship() {
		p.setLocation(p.ShipLoc)
	} else {
		p.setLocation(p.LocNo)
	}
	p.curLoc.Describe(p, LongDescription)
	p.FlushOutput()

	if p.IsFlyingSpaceship() {
		message := text.Msg(text.ShipJumpsTo, p.name, GetShipClass(p.ShipKit.Tonnage))
		p.curLoc.Talk(message, p)
	} else {
		if p.HasSpaceship() {
			landingPad := p.CurSys().FindLocation(p.ShipLoc)
			debug.Check(landingPad != nil)
			message := text.Msg(text.ShipLands, p.name)
			landingPad.Talk(message, p)
		}
		if !p.IsInsideSpaceship() {
			message := text.Msg(text.PlayerHasArrived, p.name)
			p.curLoc.Talk(message, p)
		}
	}
}

func (p *Player) destroySpyBeam() {
	p.StopSpying()

	p.Flags0 &^= model.PL0_SPYBEAM

	if p.HasSpaceship() {
		p.ShipKit.MaxHold += model.SPYBEAM_WEIGHT
		p.ShipKit.CurHold += model.SPYBEAM_WEIGHT
		debug.Check(p.ShipKit.CurHold <= p.ShipKit.MaxHold)
	}

	p.Save(database.SaveNow)
}

func (p *Player) Die() {
	// Log the death.
	locNo := p.LocNo
	if p.IsFlyingSpaceship() {
		locNo = p.ShipLoc
	}
	insured := 'X'
	if p.IsInsured() {
		insured = 'I'
	}
	log.Printf("%s died in %s/%d [%c]", p.name, p.curSys.Name(), locNo, insured)

	// Set the 'temporarily' dead flag; they might well be permanently
	// dead, but we'll get to that later! We only need this flag while
	// we're cleaing up prior to sending the body to the hospital (or
	// deleting the player).
	p.flags2 |= PL2_CORPSE

	// Remove any objects carried by the player.
	p.clearInventory()

	// Handle a death in Horsell.
	if p.IsInHorsellSystem() {
		p.curSys.(*HorsellSystem).Destroy(text.Msg(text.TimewarpParadox))
	}

	if p.IsInsured() {
		p.Outputm(text.InsuredDeath)

		// Tell any bystanders.
		if !p.IsInsideSpaceship() && !p.curLoc.IsSpace() {
			message := text.Msg(text.KillNonFatal, p.name)
			p.curLoc.Talk(message, p)
		}

		p.Sta.Cur = p.Sta.Max
		p.Str.Cur = p.Str.Max
		p.Dex.Cur = p.Dex.Max
		p.Int.Cur = p.Int.Max
		p.reward = 0
		p.deaths++
		p.Count[model.PL_G_MOVES] = 0

		// Twiddle player flags.
		p.Flags0 &^= model.PL0_FLYING | model.PL0_INSURED
		p.Flags1 &^= model.PL1_HILBERT | model.PL1_SHIELDS
		p.flags2 &^= PL2_CORPSE | PL2_WHOOSH

		// Shred their clothes.
		p.Desc = ""

		// Do cleanups for Sol locations. We should perhaps be doing
		// this for uninsured deaths too.
		p.curSys.CleanupHook(p)

		// Move to the nearest hospital.
		if !p.hospitalise() {
			log.Panicf("PANIC: Can't hospitalize %s", p.name)
		}

		if p.HasSpaceship() {
			for i := range p.Load {
				p.Load[i].Quantity = 0
			}
			p.Job.Status = JOB_NONE
			p.ShipKit.CurHold = p.ShipKit.MaxHold
			p.ShipKit.CurHull = p.ShipKit.MaxHull
			p.Missiles = 0
			p.Ammo = 0
		}

		// Tell any bystanders in the hospital that the clone has arrived.
		p.curLoc.Talk(text.Msg(text.KillResurrection, p.name), p)

		// If IsInsured() says they're insured, set the PL_INS flags to
		// agree with it. This isn't really needed, it's just belt and
		// braces.
		if p.IsInsured() {
			p.Flags0 |= model.PL0_INSURED
		}
		p.Save(database.SaveNow)

		if !p.IsInsured() {
			p.Outputm(text.REINSURE_REMINDER)
		}

		p.curLoc.Describe(p, LongDescription)
		p.FlushOutput()
		return
	}

	// Uninsured death! Say bye-bye...
	p.Outputm(text.UninsuredDeath, p.name)
	p.FlushOutput()

	// Stop any spying.
	p.clearSpyers()

	// Tell any bystanders about the death. The check for space is to
	// stop odd messages in the case of careless teleporting.
	if !p.IsInsideSpaceship() && !p.curLoc.IsSpace() {
		message := text.Msg(text.KillFatal, p.name)
		p.curLoc.Talk(message, p)
	}

	// Destroy the player's assets, if any. Their planet can't be deleted
	// yet if there are other players on it, so we don't want the Player
	// destructor won't find it if that's the case.
	if p.company != nil { //nolint:staticcheck // SA9003: empty branch
		// p.company.destroy()
	}
	if p.ownDuchy != nil { //nolint:staticcheck // SA9003: empty branch
		// p.ownDuchy.destroy()
	}
	if p.ownSystem != nil { //nolint:staticcheck // SA9003: empty branch
		// p.ownSystem.destroy()
		// if !p.ownSystem.isOffline() {
		// 	p.ownSystem = nil
		// }
	}

	// // Delete workbench files, if any.
	// deleteWorkbenchFiles(uid)

	// LeaveGame is going to clear out the entire in-memory persona
	// structure, so we need to snaffle the player's thread number for
	// later use.
	//
	// I think the above is a bit out-dated!
	p.flags2 |= PL2_DECEASED

	// Disconnect the session.
	debug.Check(p.session != nil)
	if p.isActiveThread() {
		p.session.Quit()
	} else { //nolint:staticcheck // SA9003: empty branch
		// p.deleteSession(uid)
		// // debug.Check(m_session = 0);
	}
}

// DoCommand parses and executes a single input line from the player.
func (p *Player) DoCommand(command string) {
	// Mark the player as being active.
	p.flags2 |= PL2_IN_PARSER
	defer func() { p.flags2 &^= PL2_IN_PARSER }()

	// Set up the escape route for players who snuff it during processing.
	if p.updateCounters() {
		_ = parser.Parse(command+"\n", p)
	}

	// Flush out any pending output.
	if p.session != nil { // A kludge, but the player may have quit.
		p.FlushOutput()
	}
}

func (p *Player) EnterGame(session *Session) {
	debug.Check(session != nil)
	debug.Check(p.session == nil)

	log.Printf("%s entering the game", p.name)

	// Sanitize the persona record.
	p.flags2 = 0
	p.Games++
	p.curLoc = nil
	p.NextCommodity = model.Commodity(rand.IntN(goods.MAX_GOODS))
	p.deal = nil
	p.zoneReq = nil
	p.Target = 0
	p.lastMayday = 0
	p.Crypto = 0
	// MsgOut
	p.session = session
	p.spied = nil

	p.curSys, _ = FindSystem(p.curSysName) // FIXME

	// Neuter Senators.
	if p.rank == model.RankSenator {
		p.balance = 0
		p.loan = 0
		p.reward = 0
	}

	// Log the player's arrival.
	var hostess string
	if p.rank == model.RankHostess {
		hostess = "{H}"
	}
	insured := "X"
	if p.IsInsured() {
		insured = "I"
	}
	log.Printf("%s%s [%d] on - Bal:%d (%s/%d) [%s]",
		hostess,
		p.name,
		p.session.UID(),
		p.balance,
		p.curSysName,
		p.LocNo,
		insured)

	p.place()

	if p.Job.Status == -1 {
		p.Job.Status = 0
	}

	if p.IsFlyingSpaceship() {
		p.setLocation(p.ShipLoc)
	} else {
		p.setLocation(p.LocNo)

		if p.curLoc.IsSpace() {
			log.Printf("enterGame: In space without spaceship")
			p.safeHaven()
		}
	}

	if p.curLoc.IsDeath() {
		log.Printf("enterGame: Logon in death location")
		p.safeHaven()
	}

	if p.HasSpaceship() {
		if p.ShipKit.CurHold > p.ShipKit.MaxHold {
			p.ShipKit.CurHold = p.ShipKit.MaxHold
		}
	}

	p.curLoc.Describe(p, DefaultDescription)

	if p.curLoc.IsCafe() {
		if p.rank > model.RankGroundHog && !p.IsInArenaLocation(77 /*StaffRoom*/) && !p.IsInHorsellSystem() {
			p.Output("\n")
			// readBarBoard(*this, true);
		}
	}

	// 1,000 games...
	if (p.Games % 1000) == 0 {
		p.deaths /= 2
		if p.Games == 1000 {
			p.Outputm(text.MN100)
		}
	}

	//
	p.CheckForPromotion()

	//
	p.Output("\n")

	if p.rank > model.RankGroundHog {
		p.quickScore()
		if p.HasSpaceship() {
			p.quickStatus()
		}
		p.Output("\n") // Make sure there's a blank line.
	}

	// Issue a warning to those without insurance.
	if p.rank > model.RankGroundHog && !p.IsInsured() {
		p.Outputm(text.UNINSURED_WARNING)
	}

	p.FlushOutput()

	// Tell other players
	p.announce(text.SPYNET_NOTICE_LOG_ON)

	if p.IsFlyingSpaceship() {
		debug.Check(p.IsInsideSpaceship())
		message := text.Msg(text.ShipAppears, p.name, GetShipClass(p.ShipKit.Tonnage))
		p.curLoc.Talk(message, p)
	} else if !p.IsInsideSpaceship() {
		if !p.IsInSolLocation(sol.MeetingPoint) {
			message := text.Msg(text.PersonAppears, p.MoodAndName())
			p.curLoc.Talk(message, p)
		}
	}

	if p.HasCommUnit() {
		switch p.rank {
		case model.RankGroundHog:
			if p.LocNo == sol.MeetingPoint {
				p.flags2 |= PL2_COMMS_OFF
			} else {
				p.channel = 1
			}
		case model.RankCaptain:
			p.channel = 2
		case model.RankAdventurer:
			p.channel = 3
		case model.RankTrader:
			p.channel = 4
		case model.RankMerchant:
			p.channel = 5
		case model.RankJP:
			p.channel = 6
		case model.RankGM:
			p.channel = 7
		case model.RankExplorer:
			p.channel = 8
		case model.RankSquire, model.RankThane, model.RankIndustrialist, model.RankTechnocrat, model.RankBaron, model.RankDuke:
			p.channel = 9
		default:
			p.channel = 1
		}
	} else {
		p.flags2 |= PL2_COMMS_OFF
		p.channel = 0
	}

	if p.rank > model.RankSenator {
		p.Str.Max = 120
		p.Str.Cur = p.Str.Max
		p.Sta.Max = 120
		p.Sta.Cur = p.Sta.Max
		p.Dex.Max = 120
		p.Dex.Cur = p.Dex.Max
		p.Int.Max = 120
		p.Int.Cur = p.Int.Max

		p.Flags0 |= model.PL0_INSURED + model.PL0_SPYBEAM + model.PL0_SPYSCREEN
		p.Flags1 |= model.PL1_MI6 + model.PL1_SHIP_PERMIT
		p.Flags1 &^= model.PL1_MI6_OFFERED

		if p.HasSpaceship() {
			p.Registry = "Arena"
		}

		switch p.rank { //nolint:exhaustive
		case model.RankHostess, model.RankManager:
			p.balance = 100000000 // 100,000,000 IG
		case model.RankDeity:
			p.balance = 0
			if p.HasSpaceship() {
				p.ShipKit.MaxComputer = 14
				p.ShipKit.CurComputer = p.ShipKit.MaxComputer
			}
		case model.RankEmperor:
			p.balance = 0
			if p.HasSpaceship() {
				p.ShipKit.MaxHull = 1000
				p.ShipKit.CurHull = p.ShipKit.MaxHull
				p.ShipKit.MaxEngine = 1000
				p.ShipKit.CurEngine = p.ShipKit.MaxHull
				p.ShipKit.MaxComputer = 14
				p.ShipKit.CurComputer = p.ShipKit.MaxComputer
				p.ShipKit.MaxFuel = 1000
				p.ShipKit.CurFuel = p.ShipKit.MaxFuel
				p.ShipKit.MaxHold = 10
				p.ShipKit.CurHold = p.ShipKit.MaxHold
				p.ShipKit.Tonnage = 100
				p.Guns[0].Type = GunQuadLaser
				p.Guns[0].Damage = 10
				p.Guns[1].Type = GunQuadLaser
				p.Guns[1].Damage = 10
			}
		}
	}

	p.Save(database.SaveNow)

	//
	// m_currentSystem->m_populated = Transaction::time();

	//
	if p.company != nil {
		p.company.Start()
	}

	//
	p.resumeConstruction()

	//
	p.tickerTimer = time.AfterFunc(tickerTimerPeriod, p.tickerTimerHandler)

	//
	if p.rank <= model.RankDuke && !p.IsPromoCharacter() {
		p.tourismTimer = time.AfterFunc(tourismTimerPeriod, p.tourismTimerHandler)
	}

	//
	if p.rank >= model.RankCommander && p.rank <= model.RankBaron && !p.IsPromoCharacter() {
		// debug.Check(g_exchangeTicks >= 0)
		global.ExchangeTicks++
		if p.rank < model.RankTrader {
			// debug.Check(g_haulers >= 0)
			global.Haulers++
		}
	}
}

// Checks to see if player is allowed to enter the Earth naval base.
func (p *Player) EnterNavalBase() {
	if !p.HasIDCard() && (p.Flags1&model.PL1_MI6_OFFERED) == 0 {
		p.Outputm(text.MN195)
		return
	}
	p.Outputm(text.MN1007)
	p.LocNo = sol.ParadeGround3
	p.setLocation(p.LocNo)
	p.curLoc.Describe(p, DefaultDescription)
	p.Save(database.SaveNow)
}

func (p *Player) EnterWorkbench() bool {
	if !p.HasPlanetPermit() {
		p.Outputm(text.WorkbenchNoAccess)
		return false
	}
	if workbench.Access(p.uid) != workbench.WB_ACCESS_OK {
		p.Outputm(text.WorkbenchNoAccess)
		return false
	}

	debug.Check(p.rank >= model.RankExplorer && p.rank <= model.RankDuke)

	if p.ownSystem != nil && !p.ownSystem.IsOffline() {
		p.Outputm(text.WorkbenchNotOffline)
		return false
	}

	// // Faking an exit message.
	// msg := text.Msg(text.PlayerHasLeft_W, p.Name())
	// p.curLoc.Talk(msg, this);

	// Switch over to the workbench driver.
	debug.Check(p.session != nil)
	p.session.SwitchToWorkbench()
	return true
}

func (p *Player) findFactory(number int32) (*Factory, bool) {
	if p.company == nil {
		p.Outputm(text.MN1050)
		return nil, false
	}
	return p.company.FindFactory(number)
}

func (p *Player) FindWarehouse(planet string) *model.Warehouse {
	if p.storage == nil {
		return nil
	}
	return p.storage.FindWarehouse(planet)
}

func (p *Player) FindInventoryID(id uint32) (*Object, bool) {
	for _, o := range p.inventory {
		if o.Number() == id {
			return o, true
		}
	}
	return nil, false
}

func (p *Player) FindInventoryName(name model.Name) (*Object, bool) {
	for _, o := range p.inventory {
		if name.The && (o.Flags&model.OfNoThe) != 0 {
			continue
		}

		if strings.EqualFold(name.Text, o.Name()) {
			return o, true
		}
		if name.Words != 1 {
			continue
		}
		for j := range o.Synonyms {
			if strings.EqualFold(name.Text, o.Synonyms[j]) {
				return o, true
			}
		}
	}
	return nil, false
}

func (p *Player) FlushOutput() {
	if p.MsgOut.Len() == 0 {
		debug.Trace("FlushOutput: MsgOut is empty")
		return
	}

	// TODO: MsgOut.spyDepth

	// This should probably be handled way earlier than this, but if we're
	// removing a player from the game after they've disconnected rather
	// than QUIT then the session won't be there any more.
	//
	// FIXME -- This is a bloody mess!
	if p.session != nil {
		// _ = p.session.Output(p.MsgOut.String(), p.MsgOutSpyDepth)
		_ = p.session.Output(p.MsgOut.String()) // FIXME

		// Repeat this message to any spyers.
		// if MsgOut.spyDepth >= spyPublic {
		// const char* text = msg_out->text + 4;          // Drop preamble
		// const size_t textSize = msg_out->textSize - 4; // Likewise
		// const int spyDepth = msg_out->spyDepth + 1;

		if p.MsgOutSpyDepth >= spyPublic {
			for _, spyer := range p.spyers {
				spyer.sendOutput(p.MsgOut.String(), p.MsgOutSpyDepth+1)
				spyer.FlushOutput()
			}
		}
		// }
	}

	p.MsgOut.Reset()
	p.MsgOutSpyDepth = 0
}

func (p *Player) getActiveBuildTemplate() *BuildTemplate {
	if p.BuildProject == build.NOTHING {
		return nil
	}
	template, ok := getBuildTemplate(p.BuildProject)
	if !ok {
		return nil
	}
	return template
}

func (p *Player) GMLocation() uint32 {
	return p.gmLocation
}

func (p *Player) GuessCurrentPlanet() *Planet {
	if p.IsInsideSpaceship() {
		return p.curSys.GuessPlanet(p.ShipLoc)
	}
	return p.curSys.GuessPlanet(p.LocNo)
}

func (p *Player) HasCommUnit() bool {
	if p.rank >= model.RankSenator || p.IsOnDutyNavigator() {
		return true
	}
	return (p.Flags0 & model.PL0_COMM_UNIT) != 0
}

func (p *Player) hasCompletedSnarkPuzzle() bool {
	if p.rank >= model.RankBaron && p.rank != model.RankDeity {
		return true
	}
	if p.IsPromoCharacter() {
		return true
	}
	return (p.Flags1 & model.PL1_DONE_SNARK) != 0
}

func (p *Player) hasHorsellAssignment() bool {
	if p.rank != model.RankBaron || p.IsPromoCharacter() {
		return false
	}
	if !p.HasIDCard() {
		return false
	}
	return (p.Flags0 & model.PL0_HORSELL_ASSIGNED) != 0
}

func (p *Player) HasIDCard() bool {
	if p.rank == model.RankSenator || p.IsPromoCharacter() {
		return false
	}
	if p.rank > model.RankSenator {
		return true
	}
	return (p.Flags1 & model.PL1_MI6) != 0
}

func (p *Player) HasLamp() bool {
	if p.rank >= model.RankSenator || p.IsOnDutyNavigator() || p.IsPromoCharacter() {
		return true
	}
	return (p.Flags0 & model.PL0_LIT) != 0
}

func (p *Player) HasPlanetPermit() bool {
	return (p.Flags1 & model.PL1_PO_PERMIT) != 0
}

// func (p *Player) hasPostedReward() bool {
// 	return (p.flags2 & PL2_REWARD) != 0
// }

func (p *Player) HasShipPermit() bool {
	return (p.Flags1 & model.PL1_SHIP_PERMIT) != 0
}

func (p *Player) hasSnarkAssignment() bool {
	if p.hasCompletedSnarkPuzzle() {
		return false
	}
	return (p.Flags0 & model.PL0_SNARK_ASSIGNED) != 0
}

func (p *Player) HasSpaceship() bool {
	if p.rank == model.RankGroundHog {
		return false
	}
	return (p.ShipKit.Tonnage > 0 && p.ShipLoc != 0)
}

func (p *Player) HasSpyBeam() bool {
	if p.rank == model.RankSenator {
		return false
	}
	if p.rank > model.RankSenator {
		return true
	}
	return (p.Flags0 & model.PL0_SPYBEAM) != 0
}

func (p *Player) HasSpyScreen() bool {
	if p.rank >= model.RankSenator {
		return true
	}
	if p.IsOnDutyNavigator() || p.IsPromoCharacter() {
		return true
	}
	return (p.Flags0 & model.PL0_SPYSCREEN) != 0
}

func (p *Player) HasTradingPermit() bool {
	if p.rank == model.RankSenator || p.IsPromoCharacter() {
		return false
	}
	return p.rank >= model.RankTrader
}

func (p *Player) HasWallet() bool {
	switch p.rank {
	case model.RankSenator, model.RankMascot, model.RankDeity, model.RankEmperor:
		return false
	default:
		return true
	}
}

func (p *Player) hospitalise() bool {
	// FIXME
	p.SetCurSys(SolSystem)
	p.LocNo = sol.HospitalWard3
	if p.HasSpaceship() {
		p.ShipLoc = sol.EarthLandingArea
	}

	p.setLocation(p.LocNo)
	return true
}

func (p *Player) isActiveThread() bool {
	return (p.flags2 & PL2_IN_PARSER) != 0
}

func (p *Player) IsAuto() bool {
	return (p.Flags1 & model.PL1_AUTO) != 0
}

func (p *Player) IsCommsOff() bool {
	if !p.HasCommUnit() {
		return true
	}
	return (p.flags2 & PL2_COMMS_OFF) != 0
}

func (p *Player) IsConstructionComplete() bool {
	if p.getActiveBuildTemplate() == nil {
		return false
	}
	return (p.BuildElapsed >= p.BuildDuration)
}

func (p *Player) IsDeceased() bool {
	return (p.flags2 & PL2_DECEASED) != 0
}

// func (p *Player) isDetectableSpyer() bool {
// 	return p.rank < model.RankManager
// }

func (p *Player) IsDressed() bool {
	return p.Desc != ""
}

func (p *Player) IsFlyingSpaceship() bool {
	if p.LocNo != 1 {
		return false
	}
	return (p.Flags0 & model.PL0_FLYING) != 0
}

func (p *Player) IsInArenaLocation(locNo uint32) bool {
	if !p.IsInArenaSystem() {
		return false
	}
	return p.LocNo == locNo
}

func (p *Player) IsInArenaSystem() bool {
	return p.curSys != nil && p.curSys.IsArena()
}

func (p *Player) IsInHorsellSystem() bool {
	return p.curSys != nil && p.curSys.IsHorsell()
}

func (p *Player) IsInSnarkSystem() bool {
	return p.curSys != nil && p.curSys.IsSnark()
}

func (p *Player) IsInSolLocation(locNo uint32) bool {
	if !p.IsInSolSystem() {
		return false
	}
	return p.LocNo == locNo
}

func (p *Player) IsInSolSystem() bool {
	return p.curSys != nil && p.curSys.IsSol()
}

func (p *Player) IsInsideSpaceship() bool {
	if !p.HasSpaceship() {
		return false
	}
	return (p.LocNo >= 1 && p.LocNo <= 8)
}

func (p *Player) IsInsured() bool {
	if p.rank >= model.RankSenator || p.IsPromoCharacter() {
		return true
	}
	return (p.Flags0 & model.PL0_INSURED) != 0
}

func (p *Player) IsLockedOut() bool {
	if p.rank >= model.RankManager {
		return false
	}
	return rules.IsLockedOut(p.uid)
}

func (p *Player) IsNavigator() bool {
	if p.rank == model.RankGroundHog || p.rank > model.RankDuke || p.IsPromoCharacter() {
		return false
	}
	return (p.Flags1 & model.PL1_NAVIGATOR) != 0
}

func (p *Player) IsOnDutyNavigator() bool {
	if !p.IsNavigator() {
		return false
	}
	return (p.flags2 & PL2_ON_DUTY_NAV) != 0
}

func (p *Player) IsPlaying() bool {
	return p.session != nil
}

func (p *Player) IsPromoCharacter() bool {
	return (p.Flags1 & model.PL1_PROMO_CHAR) != 0
}

func (p *Player) IsRoboBod() bool {
	return p.session != nil && p.session.IsRoboBod()
}

func (p *Player) isSpiedUpon() bool {
	// TODO
	return false
}

func (p *Player) IsSulking() bool {
	return (p.Flags0 & model.PL0_SULKING) != 0
}

func (p *Player) LeaveGame() {
	log.Printf("player.LeaveGame(%s)", p.name)

	// Cancel any pending timers.
	if p.tickerTimer != nil {
		p.tickerTimer.Stop()
		p.tickerTimer = nil
	}
	if p.tourismTimer != nil {
		p.tourismTimer.Stop()
		p.tourismTimer = nil
	}

	// Make sure any pending output has been sent.
	p.FlushOutput()

	// Destroy inventory.
	p.clearInventory()

	// Clean up any outstanding request for a factory purchase.
	// if p.zoneReq != 0 {
	// 	// TODO
	// }

	// p.deal = nil

	// if p.company != nil {
	// 	p.company.stop()
	// }

	// p.MsgOut = nil

	p.clearSpyers()
	p.StopSpying()

	if p.IsFlyingSpaceship() { //nolint:staticcheck // SA9003: empty branch
		// TODO
	}

	p.announce(text.SPYNET_NOTICE_LOG_OFF)

	// Remove from the current location.
	if p.curLoc != nil {
		p.curLoc.RemovePlayer(p)
		p.curLoc = nil
	}

	// A final save.
	p.Save(database.SaveNow)

	// // Clear m_flags2 but keep DECEASED -- Database::commit will need
	// // it at the end of the transaction.
	// m_flags2 &= (PL2_DECEASED + PL2_IN_PARSER);

	// Record their departure.
	var hostess string
	if p.rank == model.RankHostess {
		hostess = "{H}"
	}
	insured := "X"
	if p.IsInsured() {
		insured = "I"
	}
	log.Printf("%s%s [%d] off - Bal:%d (%s/%d) [%s]",
		hostess,
		p.name,
		p.session.UID(),
		p.balance,
		p.curSysName,
		p.LocNo,
		insured)

	// Clear the session pointer.
	p.session = nil

	// Clear the current system pointer. It's reset from
	// m_CurrentSystemName when they return.
	p.curSys = nil

	if p.rank >= model.RankCommander && p.rank <= model.RankBaron && !p.IsPromoCharacter() {
		global.ExchangeTicks--
		// debug.Check(g_exchangeTicks >= 0)
		if global.ExchangeTicks < 0 {
			log.Printf("Reseting global.ExchangeTicks (%d)", global.ExchangeTicks)
			global.ExchangeTicks = 0
		}
		if p.rank < model.RankTrader {
			global.Haulers--
			// debug.Check(g_haulers >= 0)
			if global.Haulers < 0 {
				log.Printf("Reseting global.Haulers (%d)", global.Haulers)
				global.Haulers = 0
			}
		}
	}
}

func (p *Player) Loan() int32 {
	return p.loan
}

func (p *Player) MaxShield() int32 {
	return p.ShipKit.MaxShield
}

func (p *Player) MaxSta() int32 {
	return p.Sta.Max
}

func (p *Player) Mood() string {
	return p.mood
}

func (p *Player) MoodAndName() string {
	if (p.flags2 & PL2_CORPSE) != 0 {
		return text.Msg(text.CorpseMood, p.name)
	}
	if p.IsSulking() {
		return text.Msg(text.SulkingMood, p.name)
	}
	if p.mood == "" {
		return p.name
	}
	return text.Msg(text.MoodAndName, p.mood, p.name)
}

func (p *Player) Name() string {
	return p.name
}

// Same as Output() but doesn't show up on other players' spybeams.
func (p *Player) Nsoutput(buf string) {
	if buf == "" {
		log.Printf("Nsoutput: Zero-length text for \"%s\"", p.name)
		return
	}
	if p.session == nil {
		return
	}
	p.sendOutput(buf, spyPrivate)
}

// Same as outputf() but doesn't show up on other players' spybeams.
func (p *Player) Nsoutputf(format string, args ...any) {
	if format == "" {
		log.Printf("Nsoutputf: Zero-length text for \"%s\"", p.name)
		return
	}
	if p.session == nil {
		return
	}
	p.sendOutput(fmt.Sprintf(format, args...), spyPrivate)
}

// Same as Outputm() but doesn't show up on other players' spybeams.
func (p *Player) Nsoutputm(msgID text.MsgNum, args ...any) {
	if p.session == nil {
		return
	}
	p.sendOutput(text.Msg(msgID, args...), spyPrivate)
}

func (p *Player) Offset() database.Offset {
	return p.offset
}

func (p *Player) Output(buf string) {
	if buf == "" {
		log.Printf("output: zero-length text for \"%s\"", p.name)
		return
	}
	if p.session == nil {
		return
	}
	p.sendOutput(buf, spyPublic)
}

func (p *Player) outputf(format string, args ...any) {
	if format == "" {
		log.Printf("outputf: Zero-length text for \"%s\"", p.name)
		return
	}
	if p.session == nil {
		return
	}
	p.sendOutput(fmt.Sprintf(format, args...), spyPublic)
}

func (p *Player) Outputm(msgID text.MsgNum, args ...any) {
	if p.session == nil {
		return
	}
	p.sendOutput(text.Msg(msgID, args...), spyPublic)
}

func (p *Player) OwnSystem() Systemer { // FIXME: *PlayerSystem
	return p.ownSystem
}

func (p *Player) OwnsPlanet() bool {
	if p.rank < model.RankSquire || p.rank > model.RankDuke {
		return false
	}
	return p.ownSystem != nil
}

func (p *Player) place() {
	debug.Check(p.LocNo > 0)
	// debug.Check(p.ShipLoc >= 0)

	// Players who should be in Horsell need careful handling.
	//
	// For the warper:
	//
	// If the system exists, ensure it was created for this player. If
	// that's not the case, or if the system doesn't exist, create a new
	// Horsell for them with a stopped puzzle.
	//
	// For others:
	//
	// If the system exists, ensure it was created for whoever warped this
	// player in the first place. If that's not the case, this player goes
	// back to Sol.

	if p.Warper == p.uid { //nolint:staticcheck
		if p.curSys != nil { //nolint:staticcheck // SA9003: empty branch
			// TODO
			// if (((HorsellSystem*) m_currentSystem)->warper() != uid) {
			// 	p.curSys = nil
			// }
		}
		if p.curSys == nil { //nolint:staticcheck // SA9003: empty branch
			// TODO
			// HorsellSystem* horsellSystem = createHorsell();
			// horsellSystem->stopPuzzle();
			// m_currentSystem = horsellSystem;
		}
	} else if p.Warper != 0 { //nolint:staticcheck
		if p.curSys != nil { //nolint:staticcheck // SA9003: empty branch
			// TODO
			// if (((HorsellSystem*) m_currentSystem)->warper() != m_warper) {
			// 	p.curSys = nil
			// 	p.warper = 0;
			// }
		}
	}

	// If the destination system doesn't exist then move the player to Sol.
	if p.curSys == nil {
		p.relocateToSol()
		return
	}

	// If the system is closed for any other reason then the player goes to
	// the capital planet if possible and Sol if not.
	if p.curSys.IsClosed() { //nolint:staticcheck
		// TODO
	}

	// Deal with special locations in Sol.
	if p.IsInSolSystem() { //nolint:staticcheck
		// TODO
	}

	// Sanity checks for arrivals on player planets.
	if p.IsInsideSpaceship() {
		l := p.CurSys().FindLocation(p.ShipLoc)

		if p.IsFlyingSpaceship() {
			if l == nil || !l.IsSpace() {
				p.ShipLoc = uint32(p.curSys.LinkLocNo())
			}
		} else { //nolint:staticcheck
			// p.ShipLoc = p.curSys.Planet().Landing
		}
	} else {
		l := p.CurSys().FindLocation(p.LocNo)

		if l == nil || l.IsDeath() || l.IsSpace() { //nolint:staticcheck
			// TODO: p.LocNo = p.curSys.Planet().Landing
		}

		if p.HasSpaceship() { //nolint:staticcheck
			// p.ShipLoc = p.curSys.Planet().Landing
		}
	}
}

func (p *Player) PromoteToDuke() {
	// TODO
}

func (p *Player) PromoteToSquire() {
	// Transfer the player's company balance to the planetary treasury
	// and destroy the company.
	if p.company != nil {
		log.Printf("Setting %s treasury to %d", p.ownSystem.Name(), p.company.Balance)
		p.ownSystem.Income(p.company.Balance, false)
		p.company.Destroy()
		// delete company;
		p.company = nil
	}

	// Clean up the persona record and set up the mega-warehouse.
	debug.Check(p.storage == nil)
	p.storage = NewStorage()

	// m_storage->warehouse[0] = new warehouse_t;
	// memset(m_storage->warehouse[0], '\0', sizeof(warehouse_t));
	// strcpy(m_storage->warehouse[0]->planet, m_ownSystem->name());

	// We don't need the link construction details any more.
	p.clearConstructionInfo()

	// Promote the player to Squire.
	p.rank = model.RankSquire

	if p.IsPlaying() {
		p.Outputm(text.MN99, p.ownSystem.Name())
		p.FlushOutput()
	}
	p.Save(database.SaveNow)

	NewsFlash(text.Msg(text.NewsFlash_1, p.name, p.ownSystem.Name()))
}

func (p *Player) quickScore() {
	p.Output("Stats: ")
	if p.HasWallet() {
		p.outputf("IG:%s ", humanize.Comma(int64(p.balance)))
	}
	var insured byte
	if p.IsInsured() {
		insured = 'Y'
	} else {
		insured = 'N'
	}
	p.outputf("Sta:%d/%d Str:%d Int:%d Dex:%d Ins:%c\n",
		p.Sta.Cur, p.Sta.Max, p.Str.Cur, p.Int.Cur, p.Dex.Cur,
		insured)
}

func (p *Player) quickStatus() {
	p.outputf("Stats: H:%d/%d S:%d/%d E:%d/%d C:%d/%d F:%d/%d",
		p.ShipKit.CurHull, p.ShipKit.MaxHull,
		p.ShipKit.CurShield, p.ShipKit.MaxShield,
		p.ShipKit.CurEngine, p.ShipKit.MaxEngine,
		p.ShipKit.CurComputer, p.ShipKit.MaxComputer,
		p.ShipKit.CurFuel, p.ShipKit.MaxFuel)
	if p.Ammo > 0 {
		p.outputf(" A:%d", p.Ammo)
	}
	if p.Missiles > 0 {
		p.outputf(" M:%d", p.Missiles)
	}
	p.Output("\n")
}

func (p *Player) Rank() model.Rank {
	return p.rank
}

func (p *Player) rankName() string {
	switch p.sex {
	case model.SexFemale:
		return model.FemaleRanks[p.rank]
	case model.SexMale:
		return model.MaleRanks[p.rank]
	default:
		return model.NeuterRanks[p.rank]
	}
}

func (p *Player) relocateToSol() {
	log.Print("player.relocateToSol")

	p.SetCurSys(SolSystem)
	if p.IsFlyingSpaceship() {
		p.ShipLoc = uint32(p.curSys.LinkLocNo()) // FIXME
	} else {
		p.LocNo = sol.EarthLandingArea
		if p.HasSpaceship() {
			p.ShipLoc = p.LocNo
		}
	}
	p.Warper = 0
}

func (p *Player) RemoveFromInventory(o *Object) {
	i := slices.Index(p.inventory, o)
	p.inventory = append(p.inventory[:i], p.inventory[i+1:]...)
}

func (p *Player) RemoveSpyer(spyer *Player) {
	i := slices.Index(p.spyers, spyer)
	p.spyers = append(p.spyers[:i], p.spyers[:i+1]...)
	p.checkForBlackBox()
}

func (p *Player) RepayLoan(amount int32) {
	p.loan -= amount
}

func (p *Player) RestartTourismTimer() {
	if p.tourismTimer != nil {
		p.tourismTimer.Stop()
	}
	p.tourismTimer = time.AfterFunc(tourismTimerPeriod, p.tourismTimerHandler)
}

func (p *Player) resumeConstruction() {
	// TODO
}

func (p *Player) Reward() int32 {
	return p.reward
}

func (p *Player) safeHaven() {
	// static int depth = 0

	// Abandon all attempts at recovery if we get a recursive call.
	// if (++depth > 1) {
	// 	log.Panic("safeHaven: Nested call, giving up!");
	// 	abort();
	// }
	// defer depth--;

	log.Printf("Moving %s to safe haven", p.name)

	// First make sure these flags are clear...
	p.Flags0 &^= model.PL0_FLYING
	p.Flags1 &^= model.PL1_HILBERT
	p.flags2 &^= PL2_TIMEWARPED

	// ...and then make sure these are set.
	p.Flags0 |= model.PL0_COMM_UNIT + model.PL0_INSURED

	// Move the player (snd ship) to the start location.
	p.SetCurSys(SolSystem)
	p.LocNo = sol.MeetingPoint
	if p.HasSpaceship() {
		p.ShipLoc = sol.EarthLandingArea
	}
	p.setLocation(p.LocNo)

	// Destroy the player's inventory.
	p.clearInventory()

	// Tell the player where they are not.
	p.curLoc.Describe(p, LongDescription)
}

func (p *Player) Save(when database.SaveWhen) {
	// pc, _, _, _ := runtime.Caller(1)
	// fn := runtime.FuncForPC(pc)
	// log.Printf("Player.Save called from %s", fn.Name())

	if p.saveFunc == nil {
		panic("player.saveFunc is not set!")
	}
	p.saveFunc(p, when)
}

// This is the routine that -really- puts the output message into the player's
// output buffer. It still doesn't result in any visible output unless there's
// a change in the spying depth, in which case the current output buffer
// contents are flushed out.
func (p *Player) sendOutput(buf string, spyDepth int) {
	if spyDepth < spyPrivate {
		log.Panic("sendOutput: spyDepth is less than spyPrivate")
	}

	// Deal with any change in the spying depth.
	if spyDepth != p.MsgOutSpyDepth {
		if p.MsgOut.Len() > 0 {
			p.FlushOutput()
		}
		p.MsgOutSpyDepth = spyDepth
	}

	// Queue up the new text.
	if buf == "" {
		debug.Trace("sendOutput: Got zero-length text")
		return
	}
	p.MsgOut.WriteString(buf)
}

func (p *Player) Session() *Session {
	return p.session
}

func (p *Player) SetAlarm(minutes int32) {
	if p.alarmTimer != nil {
		p.alarmTimer.Stop()
		p.alarmTimer = nil
	}
	p.alarmTimer = time.AfterFunc(time.Duration(minutes*60)*time.Second, p.alarmTimerHandler)
}

func (p *Player) SetAuto(v bool) {
	if v {
		p.Flags1 |= model.PL1_AUTO
	} else {
		p.Flags1 &^= model.PL1_AUTO
	}
}

func (p *Player) SetBalance(v int32) {
	p.balance = v
}

func (p *Player) SetBrief(v bool) {
	if v {
		p.Flags0 |= model.PL0_BRIEF
	} else {
		p.Flags0 &^= model.PL0_BRIEF
	}
}

func (p *Player) SetChannel(v int32) {
	p.channel = v
}

func (p *Player) SetCommsOff(v bool) {
	if v {
		p.flags2 |= PL2_COMMS_OFF
	} else {
		p.flags2 &^= PL2_COMMS_OFF
	}
}

func (p *Player) SetCurSta(v int32) {
	p.Sta.Cur = v
}

func (p *Player) SetCurSys(s Systemer) {
	if s != p.curSys && s != nil {
		p.curSysName = s.Name()
	}

	if p.curSysName == "" {
		log.Panic("SetCurSys: curSysName is empty")
	}

	p.curSys = s

	if s != nil { //nolint:staticcheck
		// TODO: populated
	}
}

func (p *Player) SetDeaths(v int32) {
	p.deaths = v
}

func (p *Player) setGMLocation() {
	debug.Precondition(p.rank == model.RankAdventurer)

	// if (m_gmLocation <= 0 || Transaction::time() - last_on > 60 * 60) {
	// 	output(mnPlaceGM);
	//
	// 	for (;;) {
	// 		const int loc = 10 + random() % 679;
	//
	// 		// FIX ME!!
	//
	// 		if (loc < 9 ||                         /* Ship locations */
	// 		loc == 217 ||                      /* Godfather's office */
	// 		(loc > 394 && loc < 425) ||        /* Naval base */
	// 		loc == 509 ||                      /* Deserted gates */
	// 		loc > 696)                         /* Snark/Horsell */
	// 		{
	// 			continue;
	// 		}
	//
	// 		const Location* theLocation = solSystem->findLocation(loc);
	//
	// 		if (theLocation->m_flags == 0 && theLocation->m_events[entry] == 0) {
	// 			m_gmLocation = loc;
	// 			break;
	// 		}
	// 	}
	// }

	// debug.Check(p.gmLocation > 0)
}

func (p *Player) SetIDCard(v bool) {
	if v {
		p.Flags1 |= model.PL1_MI6
	} else {
		p.Flags1 &^= model.PL1_MI6
	}
}

func (p *Player) SetInsured(v bool) {
	if v {
		p.Flags0 |= model.PL0_INSURED
	} else {
		p.Flags0 &^= model.PL0_INSURED
	}
}

func (p *Player) setLocation(locNo uint32) {
	debug.Check(locNo > 0)

	// I;m not sure this is the right place to do this!
	if p.curLoc != nil {
		p.curLoc.RemovePlayer(p)
	}

	p.curLoc = p.CurSys().FindLocation(locNo)
	if p.curLoc == nil {
		log.Printf("setLocation: Location not found (%s/%d)", p.curSys.Name(), locNo)
		p.safeHaven()
	}

	p.curLoc.InsertPlayer(p)
}

func (p *Player) SetMI6Offered(v bool) {
	if v {
		p.Flags1 |= model.PL1_MI6_OFFERED
	} else {
		p.Flags1 &^= model.PL1_MI6_OFFERED
	}
}

func (p *Player) SetMood(v string) {
	p.mood = v
}

func (p *Player) SetPlanetPermit(v bool) {
	if v {
		p.Flags1 |= model.PL1_PO_PERMIT
	} else {
		p.Flags1 &^= model.PL1_PO_PERMIT
	}
}

func (p *Player) SetRank(v model.Rank) {
	p.rank = v
}

func (p *Player) SetShipDesc(v string) {
	p.shipDesc = v
}

func (p *Player) SetShipPermit(v bool) {
	if v {
		p.Flags1 |= model.PL1_SHIP_PERMIT
	} else {
		p.Flags1 &^= model.PL1_SHIP_PERMIT
	}
}

func (p *Player) SetSulking(v bool) {
	if v {
		p.Flags0 |= model.PL0_SULKING
	} else {
		p.Flags0 &^= model.PL0_SULKING
	}
}

func (p *Player) SetTimerCount(v int32) {
	p.Count[model.PL_G_TIMER] = v
}

func (p *Player) SetTradeGroup(v model.CommodityGroup) {
	p.ComType = v
}

func (p *Player) Sex() model.Sex {
	return p.sex
}

func (p *Player) ShipDesc() string {
	if p.shipDesc == "" {
		return text.Msg(text.DefaultShipDescription)
	}
	return p.shipDesc
}

func (p *Player) ShipLocNo() uint32 {
	return p.ShipLoc
}

func (p *Player) StopNavigating() {
	p.flags2 &^= PL2_ON_DUTY_NAV
	// p.secretOutput(mnNavigateNowOff)

	log.Printf("{H} Navigator off duty: %s", p.name)

	if !p.IsInsideSpaceship() { //nolint:staticcheck // SA9003: empty branch
		// TODO
	}

	p.session.StartBilling()
}

func (p *Player) StopSpying() {
	if p.spied != nil {
		p.spied.RemoveSpyer(p)
		p.spied = nil
	}
}

func (p *Player) updateCounters() bool {
	// Make time pass.
	p.Count[model.PL_G_TIMER]++
	if p.Count[model.PL_G_TIMER] == 100 {
		p.Sta.Cur--
		if p.Sta.Cur < 6 {
			if p.Sta.Cur < 3 {
				p.Outputm(text.VeryHungry)
				if p.Sta.Cur <= 0 {
					p.Outputm(text.StarvedToDeath)
					p.Die()
					return false
				}
			} else {
				p.Outputm(text.Hungry)
			}
		}
		p.Count[model.PL_G_TIMER] = 0
		p.Save(database.SaveNow)
	}

	// Check for time up on the Snark puzzle.
	if p.Count[model.PL_G_MOVES] > 0 {
		p.Count[model.PL_G_MOVES]--
		if p.Count[model.PL_G_MOVES] == 0 {
			p.Outputm(text.GM_OUT_OF_TIME)
			p.Die()
			return false
		}
	}

	// Check for the WHOOSH effect.
	if (p.flags2 & PL2_WHOOSH) != 0 {
		if p.Count[model.PL_G_TIMER] != 0 && rand.IntN(10) == 0 {
			p.Outputm(text.WHOOSHEffect)
		}
	}

	//
	return true
}

func (p *Player) WantsBrief() bool {
	return (p.Flags0 & model.PL0_BRIEF) != 0
}

func NewPlayer(uid ibgames.AccountID, name string, sex model.Sex, strength, stamina, intelligence, dexterity int32) *Player {
	p := &Player{
		balance:    13000,
		curSysName: "Sol",
		Dex:        PlayerStat{Max: dexterity, Cur: dexterity},
		Flags0:     model.PL0_COMM_UNIT,
		Int:        PlayerStat{Max: intelligence, Cur: intelligence},
		LocNo:      sol.MeetingPoint,
		name:       name,
		offset:     database.NextPersonaOffset(),
		rank:       model.RankGroundHog,
		sex:        sex,
		Sta:        PlayerStat{Max: stamina, Cur: stamina},
		Str:        PlayerStat{Max: strength, Cur: strength},
		uid:        uid,
		saveFunc:   defaultPlayerSaver,
	}

	// Write the new record away to the persona file.
	p.Save(database.SaveNow)

	// Record the newborn's arrival.
	log.Printf("New persona for %s [%d]", name, uid)

	return p
}

// NewPlayerFromDBPersona creates a Player from DBPersona data.
func NewPlayerFromDBPersona(dbp *model.DBPersona, off database.Offset) *Player {
	name := text.CStringToString(dbp.Name[:])

	p := &Player{
		name:          name,
		uid:           ibgames.AccountID(dbp.ID),
		offset:        off,
		sex:           model.Sex(dbp.Sex),
		rank:          model.Rank(dbp.Rank),
		Desc:          extractString(dbp.Desc[:]),
		mood:          extractString(dbp.Mood[:]),
		Str:           PlayerStat{Max: int32(dbp.MaxStr), Cur: int32(dbp.CurStr)},
		Sta:           PlayerStat{Max: int32(dbp.MaxSta), Cur: int32(dbp.CurSta)},
		Dex:           PlayerStat{Max: int32(dbp.MaxDex), Cur: int32(dbp.CurDex)},
		Int:           PlayerStat{Max: int32(dbp.MaxInt), Cur: int32(dbp.CurInt)},
		Shipped:       dbp.Shipped,
		Games:         dbp.Games,
		Flags0:        dbp.Flags[0],
		Flags1:        dbp.Flags[1],
		balance:       dbp.Balance,
		loan:          dbp.Loan,
		reward:        dbp.Reward,
		deaths:        dbp.Frame[0],
		CustomRank:    PlayerCustomRank(dbp.Frame[1]),
		TradeCredits:  dbp.Frame[2],
		Warper:        ibgames.AccountID(dbp.Frame[3]),
		curSysName:    extractString(dbp.StarSystem[:]),
		LocNo:         dbp.LocNo,
		ShipLoc:       dbp.ShipLoc,
		shipDesc:      extractString(dbp.ShipDesc[:]),
		Registry:      extractString(dbp.Registry[:]),
		Missiles:      int32(dbp.Missiles),
		Ammo:          int32(dbp.Ammo),
		Guns:          dbp.Guns,
		LastTrade:     int(dbp.LastTrade),
		LastOn:        dbp.LastOn,
		Count:         dbp.Count,
		BuildProject:  model.Project(dbp.Build.IDProject),
		BuildDuration: dbp.Build.Duration,
		BuildElapsed:  dbp.Build.Elapsed,
		saveFunc:      defaultPlayerSaver,
	}

	p.ShipKit = Equipment{
		MaxHull:     int32(dbp.ShipKit.MaxHull),
		CurHull:     int32(dbp.ShipKit.CurHull),
		MaxShield:   int32(dbp.ShipKit.MaxShield),
		CurShield:   int32(dbp.ShipKit.CurShield),
		MaxEngine:   int32(dbp.ShipKit.MaxEngine),
		CurEngine:   int32(dbp.ShipKit.CurEngine),
		MaxComputer: int32(dbp.ShipKit.MaxComputer),
		CurComputer: int32(dbp.ShipKit.CurComputer),
		MaxFuel:     int32(dbp.ShipKit.MaxFuel),
		CurFuel:     int32(dbp.ShipKit.CurFuel),
		MaxHold:     int32(dbp.ShipKit.MaxHold),
		CurHold:     int32(dbp.ShipKit.CurHold),
		Tonnage:     int32(dbp.ShipKit.Tonnage),
	}

	for i := 0; i < model.MAX_LOAD && i < len(dbp.Load); i++ {
		p.Load[i] = model.Cargo{
			Type:     model.Commodity(dbp.Load[i].Type),
			Quantity: dbp.Load[i].Quantity,
			Origin:   extractString(dbp.Load[i].Origin[:]),
			Cost:     dbp.Load[i].Cost,
		}
	}

	p.Job = model.Work{
		JobType: dbp.Job.JobType,
		From:    extractString(dbp.Job.From[:]),
		To:      extractString(dbp.Job.To[:]),
		Status:  dbp.Job.Status,
		Value:   dbp.Job.Value,
		Gtu:     dbp.Job.GTU,
		Credits: dbp.Job.Credits,
		Age:     dbp.Job.Age,
	}

	p.Job.Pallet = model.Cargo{
		Type:     model.Commodity(dbp.Job.Pallet.Type),
		Quantity: dbp.Job.Pallet.Quantity,
		Origin:   extractString(dbp.Job.Pallet.Origin[:]),
		Cost:     dbp.Job.Pallet.Cost,
	}

	p.Job.FactryWk = model.FactoryJob{
		Deliver: model.FactoryID{
			Number: dbp.Job.Type.FactryWk.Deliver.Number,
			Owner:  text.CStringToString(dbp.Job.Type.FactryWk.Deliver.Owner[:]),
		},
		PickUp: model.FactoryID{
			Number: dbp.Job.Type.FactryWk.PickUp.Number,
			Owner:  text.CStringToString(dbp.Job.Type.FactryWk.PickUp.Owner[:]),
		},
	}

	p.Job.GenWk = model.GeneralJob{
		WhereTo: dbp.Job.Type.GenWk.WhereTo,
		Owner:   text.CStringToString(dbp.Job.Type.GenWk.Owner[:]),
	}

	if p.rank < model.RankSquire {
		p.gmLocation = dbp.PP.GMLocation

		for i := range model.MAX_STORES {
			if dbp.PP.Storage[i].Planet[0] == byte(0) {
				continue
			}
			planet := text.CStringToString(dbp.PP.Storage[i].Planet[:])
			_, ok := FindPlanet(planet)
			if !ok {
				log.Printf("Liquidating warehouse on %s", planet)
				continue
			}

			if p.storage == nil {
				p.storage = NewStorage()
			}

			dbWarehouse := &dbp.PP.Storage[i]
			warehouse := model.Warehouse{
				Planet: planet,
			}
			for j := range 20 {
				if dbWarehouse.Bay[j].Quantity > 0 {
					warehouse.Bay[j].Type = model.Commodity(dbWarehouse.Bay[j].Type)
					warehouse.Bay[j].Quantity = dbWarehouse.Bay[j].Quantity
					warehouse.Bay[j].Origin = text.CStringToString(dbWarehouse.Bay[j].Origin[:])
					warehouse.Bay[j].Cost = dbWarehouse.Bay[j].Cost
				}
			}
			p.storage.Warehouse[i] = &warehouse
		}

		if text.CStringToString(dbp.PP.Company.Name[:]) != "" {
			p.company = NewCompanyFromDB(p, dbp.PP.Company, dbp.PP.Factory)
		}
	} else if p.rank <= model.RankDuke {
		fief := text.CStringToString(dbp.RP.Fief[:])

		if p.rank == model.RankDuke {
			customsRate := dbp.RP.Duchy.CustomsRate
			favouredRate := dbp.RP.Duchy.FavouredRate
			p.ownDuchy = NewPlayerDuchy(p, fief, customsRate, favouredRate)
			debug.Check(p.ownDuchy != nil)
			log.Printf("Loaded %s duchy [%s]", p.ownDuchy.Name(), p.name)
			// NoteDuchyFixup(p.ownDuchy, persona.rp.duchy)
		}

		p.ownSystem = NewPlayerSystemFromDB(p, fief, dbp.RP.Planet)

		p.storage = NewStorage()
		warehouse := &model.Warehouse{Planet: p.ownSystem.Name()}
		for i := range 20 {
			if dbp.RP.Storage.Bay[i].Quantity == 0 {
				continue
			}
			warehouse.Bay[i].Type = model.Commodity(dbp.RP.Storage.Bay[i].Type)
			warehouse.Bay[i].Quantity = dbp.RP.Storage.Bay[i].Quantity
			warehouse.Bay[i].Origin = text.CStringToString(dbp.RP.Storage.Bay[i].Origin[:])
			warehouse.Bay[i].Cost = dbp.RP.Storage.Bay[i].Cost
		}
		p.storage.Warehouse[0] = warehouse

		p.Facilities = dbp.RP.Facilities

		// if p.rank == model.RankDuke {
		// 	p.ownDuchy.Start()
		// }
	}

	return p
}

// LoadPlayer creates a Player from raw DBPersona data and adds it to the global Players map
func LoadPlayer(dbp *model.DBPersona, off database.Offset) *Player {
	p := NewPlayerFromDBPersona(dbp, off)
	Players[p.name] = p
	return p
}

// extractString converts a null-terminated byte array to a Go string
func extractString(bytes []byte) string {
	for i, b := range bytes {
		if b == 0 {
			return string(bytes[:i])
		}
	}
	return string(bytes) // fallback if no null terminator
}

func (p *Player) GenderedMsg(f, m, n text.MsgNum) text.MsgNum {
	switch p.sex {
	case model.SexFemale:
		return f
	case model.SexMale:
		return m
	default:
		return n
	}
}

func (p *Player) tickerTimerHandler() {
	global.Lock()
	defer global.Unlock()

	defer database.CommitDatabase()
	p.tickerTimerProc()
}

func (p *Player) tickerTimerProc() {
	// log.Printf("Ticker timer proc for %s", p.name)

	p.tickerTimer = nil

	if p.curLoc.IsExchange() {
		planet := p.GuessCurrentPlanet()
		if planet == nil || planet.IsClosed() {
			return
		}

		if p.ComType == -1 {
			buy := planet.BuyingPrice(p.NextCommodity)
			sell := planet.SellingPrice(p.NextCommodity)

			if sell > 0 {
				p.Outputm(text.TickerTape_BuyAndSell,
					goods.GoodsArray[p.NextCommodity].Name,
					sell,
					planet.AvailableQuantity(p.NextCommodity),
					buy)
			} else {
				p.Outputm(text.TickerTape_BuyOnly, goods.GoodsArray[p.NextCommodity].Name, buy)
			}

			p.FlushOutput()

			p.NextCommodity++
			if p.NextCommodity > goods.LastCommodity {
				p.NextCommodity = goods.FirstCommodity
			}
		} else {
			commDisplay++
			if (commDisplay % 4) == 0 {
				p.Outputm(text.CommodityDisplayHeader, stardate())

				for i := range goods.MAX_GOODS {
					if goods.GoodsArray[i].Group == p.ComType {
						// const commodity_t commodity = static_cast < commodity_t > (i)

						buy := planet.BuyingPrice(model.Commodity(i))
						sell := planet.SellingPrice(model.Commodity(i))

						if sell > 0 {
							p.Outputm(text.CommodityDisplayLine_BuyAndSell,
								goods.GoodsArray[i].Name,
								humanize.Comma(int64(buy)),
								humanize.Comma(int64(sell)),
								humanize.Comma(int64(planet.AvailableQuantity(model.Commodity(i)))))
						} else {
							p.Outputm(text.CommodityDisplayLine_BuyOnly,
								goods.GoodsArray[i].Name,
								humanize.Comma(int64(buy)))
						}
					}
				}

				p.Outputm(text.CommodityDisplayFooter)
				p.FlushOutput()
			}
		}
	}

	p.tickerTimer = time.AfterFunc(tickerTimerPeriod, p.tickerTimerHandler)
}

func (p *Player) ToggleShields() {
	p.Flags1 ^= model.PL1_SHIELDS
}

func (p *Player) tourismTimerHandler() {
	global.Lock()
	defer global.Unlock()

	defer database.CommitDatabase()
	p.tourismTimerProc()
}

func (p *Player) tourismTimerProc() {
	// log.Printf("Tourism timer proc for %s", p.name)

	p.tourismTimer = nil

	if p.rank > model.RankDuke || p.IsPromoCharacter() {
		return
	}

	p.curSys.UpdateTime(p)

	p.tourismTimer = time.AfterFunc(tourismTimerPeriod, p.tourismTimerHandler)
}

func (p *Player) Serialize() *model.DBPersona {
	dbp := &model.DBPersona{
		ID:        uint32(p.uid),
		Sex:       byte(p.sex),
		Rank:      uint32(p.rank),
		MaxStr:    uint32(p.Str.Max),
		CurStr:    uint32(p.Str.Cur),
		MaxSta:    uint32(p.Sta.Max),
		CurSta:    uint32(p.Sta.Cur),
		MaxInt:    uint32(p.Int.Max),
		CurInt:    uint32(p.Int.Cur),
		MaxDex:    uint32(p.Dex.Max),
		CurDex:    uint32(p.Dex.Cur),
		Shipped:   p.Shipped,
		Games:     p.Games,
		Flags:     [2]uint32{p.Flags0, p.Flags1},
		Balance:   p.balance,
		Loan:      p.loan,
		Reward:    p.reward,
		Frame:     [4]int32{p.deaths, int32(p.CustomRank), p.TradeCredits, int32(p.Warper)},
		LocNo:     p.LocNo,
		LastTrade: int32(p.LastTrade),
		ShipLoc:   p.ShipLoc,
		Count:     [4]int32{p.Count[0], p.Count[1], p.Count[2], p.Count[3]},
		LastOn:    p.LastOn,
	}

	copy(dbp.Name[:], p.name)
	copy(dbp.Desc[:], p.Desc)
	copy(dbp.Mood[:], p.mood)
	copy(dbp.StarSystem[:], p.curSysName)

	if p.Job.Status != JOB_NONE {
		dbp.Job.Pallet.Type = int32(p.Job.Pallet.Type)
		dbp.Job.Pallet.Quantity = p.Job.Pallet.Quantity
		copy(dbp.Job.Pallet.Origin[:], p.Job.Pallet.Origin)
		dbp.Job.Pallet.Cost = p.Job.Pallet.Cost

		dbp.Job.JobType = p.Job.JobType
		copy(dbp.Job.From[:], p.Job.From)
		copy(dbp.Job.To[:], p.Job.To)
		dbp.Job.Status = p.Job.Status
		dbp.Job.Value = p.Job.Value
		dbp.Job.GTU = p.Job.Gtu
		dbp.Job.Credits = p.Job.Credits

		switch p.Job.JobType {
		case jobs.JOB_FACTORY:
			dbp.Job.Type.FactryWk.Deliver.Number = p.Job.FactryWk.Deliver.Number
			copy(dbp.Job.Type.FactryWk.Deliver.Owner[:], p.Job.FactryWk.Deliver.Owner)
			dbp.Job.Type.FactryWk.PickUp.Number = p.Job.FactryWk.PickUp.Number
			copy(dbp.Job.Type.FactryWk.PickUp.Owner[:], p.Job.FactryWk.PickUp.Owner)
		case jobs.JOB_GENERAL:
			dbp.Job.Type.GenWk.WhereTo = p.Job.GenWk.WhereTo
			copy(dbp.Job.Type.GenWk.Owner[:], p.Job.GenWk.Owner)
		}

		dbp.Job.Age = p.Job.Age
	}

	if p.HasSpaceship() {
		copy(dbp.ShipDesc[:], p.shipDesc)
		copy(dbp.Registry[:], p.Registry)

		dbp.Missiles = uint32(p.Missiles)
		dbp.Ammo = uint32(p.Ammo)
		dbp.ShipKit.MaxHull = uint32(p.ShipKit.MaxHull)
		dbp.ShipKit.CurHull = uint32(p.ShipKit.CurHull)
		dbp.ShipKit.MaxShield = uint32(p.ShipKit.MaxShield)
		dbp.ShipKit.CurShield = uint32(p.ShipKit.CurShield)
		dbp.ShipKit.MaxEngine = uint32(p.ShipKit.MaxEngine)
		dbp.ShipKit.CurEngine = uint32(p.ShipKit.CurEngine)
		dbp.ShipKit.MaxComputer = uint32(p.ShipKit.MaxComputer)
		dbp.ShipKit.CurComputer = uint32(p.ShipKit.CurComputer)
		dbp.ShipKit.MaxFuel = uint32(p.ShipKit.MaxFuel)
		dbp.ShipKit.CurFuel = uint32(p.ShipKit.CurFuel)
		dbp.ShipKit.MaxHold = uint32(p.ShipKit.MaxHold)
		dbp.ShipKit.CurHold = uint32(p.ShipKit.CurHold)
		dbp.ShipKit.Tonnage = uint32(p.ShipKit.Tonnage)

		for i := range p.Guns {
			if p.Guns[i].Type > 0 {
				dbp.Guns[i].Type = p.Guns[i].Type
				dbp.Guns[i].Damage = p.Guns[i].Damage
			}
		}

		for i := range p.Load {
			if p.Load[i].Quantity > 0 {
				dbp.Load[i].Type = int32(p.Load[i].Type)
				dbp.Load[i].Quantity = p.Load[i].Quantity
				copy(dbp.Load[i].Origin[:], p.Load[i].Origin)
				dbp.Load[i].Cost = p.Load[i].Cost
			}
		}
	}

	if p.rank >= model.RankExplorer && p.rank <= model.RankDuke {
		dbp.Build = model.DBBuild{
			IDProject: uint16(p.BuildProject),
			Duration:  p.BuildDuration,
			Elapsed:   p.BuildElapsed,
		}
	}

	if p.rank < model.RankSquire {
		dbp.PP = &model.DBPPData{}

		// GM location.
		dbp.PP.GMLocation = p.gmLocation

		// Warehouse details.
		if p.storage != nil {
			for i, warehouse := range p.storage.Warehouse {
				if warehouse == nil {
					continue
				}
				copy(dbp.PP.Storage[i].Planet[:], warehouse.Planet)
				for j, bay := range warehouse.Bay {
					dbp.PP.Storage[i].Bay[j].Type = int32(bay.Type)
					dbp.PP.Storage[i].Bay[j].Quantity = bay.Quantity
					copy(dbp.PP.Storage[i].Bay[j].Origin[:], bay.Origin)
					dbp.PP.Storage[i].Bay[j].Cost = bay.Cost
				}
			}
		}

		// Company and factory details.
		if p.company != nil {
			p.company.Serialize(&dbp.PP.Company, &dbp.PP.Factory)
		}
	} else if p.rank <= model.RankDuke {
		dbp.RP = &model.DBRPData{}

		copy(dbp.RP.Fief[:], p.ownSystem.Name())

		// Planet details.
		p.ownSystem.Serialize(&dbp.RP.Planet)

		// Warehouse details.
		for i, bay := range p.storage.Warehouse[0].Bay {
			if bay.Quantity == 0 {
				continue
			}
			dbp.RP.Storage.Bay[i].Type = int32(bay.Type)
			dbp.RP.Storage.Bay[i].Quantity = bay.Quantity
			copy(dbp.RP.Storage.Bay[i].Origin[:], bay.Origin)
			dbp.RP.Storage.Bay[i].Cost = bay.Cost
		}

		// Construction details.
		if p.rank >= model.RankBaron {
			dbp.RP.Facilities = p.Facilities
		}

		// Duchy details.
		if p.rank == model.RankDuke && p.ownDuchy != nil {
			p.ownDuchy.Serialize(&dbp.RP.Duchy)
		}
	}

	return dbp
}

func (p *Player) UID() ibgames.AccountID {
	return p.uid
}

func (p *Player) UnknownCommand() {
	p.Outputm(text.UnknownCommand)
	if p.Rank() == model.RankGroundHog {
		p.Outputm(text.UnknownCommand_Hint)
	}
}
