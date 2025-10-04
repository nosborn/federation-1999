package server

import (
	"log"
	"strconv"
	"strings"

	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/text"
	"github.com/nosborn/federation-1999/pkg/ibgames"
)

type SetupState int

const (
	setupStateName SetupState = 1 + iota
	setupStateSex
	setupStateStats
	setupStateStrength
	setupStateStamina
	setupStateDexterity
)

type SetupDriver struct {
	dexterity    int32
	intelligence int32
	name         string
	replyMax     int
	replyMin     int
	session      *Session
	sex          model.Sex
	stamina      int32
	state        SetupState
	strength     int32
	uid          ibgames.AccountID
}

const (
	totalStats = 140

	defaultStr = totalStats / 4
	defaultSta = totalStats / 4
	defaultDex = totalStats / 4
	defaultInt = totalStats / 4

	statMin = 20
	statMax = totalStats / 2
)

func NewSetupDriver(session *Session, uid ibgames.AccountID) *SetupDriver {
	sd := SetupDriver{
		dexterity:    defaultDex,
		intelligence: defaultInt,
		session:      session,
		stamina:      defaultSta,
		strength:     defaultStr,
		uid:          uid,
	}
	session.Output(text.Msg(text.NewPlayerWelcome))
	sd.startName()
	return &sd
}

func (sd *SetupDriver) Destroy() {
	sd.session = nil
}

func (sd *SetupDriver) Dispatch(line string) bool {
	// input := CleanInput(line, noNL) // TODO
	input := strings.TrimSpace(line)

	switch sd.state {
	case setupStateName:
		return sd.setupName(input)
	case setupStateSex:
		return sd.setupSex(input)
	case setupStateStats:
		return sd.setupStats(input)
	case setupStateStrength:
		return sd.setupStrength(input)
	case setupStateStamina:
		return sd.setupStamina(input)
	case setupStateDexterity:
		return sd.setupDexterity(input)
	}

	log.Printf("Bad state in SetupDriver.Dispatch")
	return false
}

func (sd *SetupDriver) startName() bool {
	// May have been round once already, so make sure to re-initialise
	// properly.

	sd.state = setupStateName
	sd.name = ""

	if err := sd.session.Output(text.Msg(text.NamePreamble)); err != nil {
		log.Print(err)
		return false
	}
	if err := sd.session.Output(text.Msg(text.NamePrompt)); err != nil {
		log.Print(err)
		return false
	}
	return true
}

func (sd *SetupDriver) setupName(input string) bool {
	if strings.EqualFold(input, "quit") {
		return false
	}
	if len(input) >= 3 || len(input) < model.NAME_SIZE {
		name, ok := NormalizeName(input)
		if ok && IsNameAvailable(name) {
			sd.name = name
			return sd.startSex()
		}
	}
	if len(input) > 0 {
		if err := sd.session.Output(text.Msg(text.NameNotAvailable)); err != nil {
			log.Print(err)
			return false
		}
	}
	if err := sd.session.Output(text.Msg(text.NamePrompt)); err != nil {
		log.Print(err)
		return false
	}
	return true
}

func (sd *SetupDriver) startSex() bool {
	sd.state = setupStateSex
	if err := sd.session.Output(text.Msg(text.SexPrompt)); err != nil {
		log.Print(err)
		return false
	}
	return true
}

func (sd *SetupDriver) setupSex(input string) bool {
	if len(input) > 0 {
		if strings.ToLower(input) == "female"[:len(input)] {
			sd.sex = model.SexFemale
			return sd.startStats()
		}
		if strings.ToLower(input) == "male"[:len(input)] {
			sd.sex = model.SexMale
			return sd.startStats()
		}
	}
	if err := sd.session.Output(text.Msg(text.SexPrompt)); err != nil {
		log.Print(err)
		return false
	}
	return true
}

func (sd *SetupDriver) startStats() bool {
	if sd.state != setupStateStrength && sd.state != setupStateStamina && sd.state != setupStateDexterity {
		if err := sd.session.Output(text.Msg(text.ChangeStatsIntro, totalStats, statMin, statMax)); err != nil {
			log.Print(err)
			return false
		}
	}

	sd.state = setupStateStats

	if err := sd.session.Output(text.Msg(text.ChangeStatsPreamble, sd.strength, sd.stamina, sd.dexterity, sd.intelligence)); err != nil {
		log.Print(err)
		return false
	}
	if err := sd.session.Output(text.Msg(text.ChangeStatsPrompt)); err != nil {
		log.Print(err)
		return false
	}
	return true
}

func (sd *SetupDriver) setupStats(input string) bool {
	yes, ok := text.YesNoReply(input)
	if ok {
		if yes {
			return sd.startStrength()
		}
		player := NewPlayer(sd.uid, sd.name, sd.sex, sd.strength, sd.stamina, sd.intelligence, sd.dexterity)
		if err := sd.session.EndSetup(player); err != nil {
			return false
		}
		return true
	}
	if err := sd.session.Output(text.Msg(text.ChangeStatsPrompt)); err != nil {
		log.Print(err)
		return false
	}
	return true
}

func (sd *SetupDriver) startStrength() bool {
	reserved := statMin * 3
	available := totalStats
	maximum := min(available-reserved, statMax)

	sd.state = setupStateStrength
	sd.replyMin = statMin
	sd.replyMax = maximum

	return sd.sendStatPrompt(text.StrengthName)
}

func (sd *SetupDriver) setupStrength(input string) bool {
	if reply := sd.getStatReply(input); reply > 0 {
		sd.strength = int32(reply)
		return sd.startStamina()
	}
	return sd.sendStatPrompt(text.StrengthName)
}

func (sd *SetupDriver) startStamina() bool {
	allocated := int(sd.strength)
	reserved := statMin * 2
	available := totalStats - allocated
	maximum := min(available-reserved, statMax)

	sd.state = setupStateStamina
	sd.replyMin = statMin
	sd.replyMax = maximum

	return sd.sendStatPrompt(text.StaminaName)
}

func (sd *SetupDriver) setupStamina(input string) bool {
	if reply := sd.getStatReply(input); reply > 0 {
		sd.stamina = int32(reply)
		return sd.startDexterity()
	}
	return sd.sendStatPrompt(text.StaminaName)
}

func (sd *SetupDriver) startDexterity() bool {
	allocated := int(sd.strength + sd.stamina)
	reserved := statMin
	available := totalStats - allocated
	maximum := min(available-reserved, statMax)
	minimum := available - maximum

	if maximum == statMin {
		sd.dexterity = statMin
		sd.intelligence = statMin
		return sd.startStats()
	}

	sd.state = setupStateDexterity
	sd.replyMin = minimum
	sd.replyMax = maximum

	return sd.sendStatPrompt(text.DexterityName)
}

func (sd *SetupDriver) setupDexterity(input string) bool {
	if reply := sd.getStatReply(input); reply > 0 {
		sd.dexterity = int32(reply)
		sd.intelligence = totalStats - (sd.strength + sd.stamina + sd.dexterity)
		return sd.startStats()
	}
	return sd.sendStatPrompt(text.DexterityName)
}

func (sd *SetupDriver) sendStatPrompt(statName text.MsgNum) bool {
	if err := sd.session.Output(text.Msg(text.StatPrompt, text.Msg(statName), sd.replyMin, sd.replyMax)); err != nil {
		log.Print(err)
		return false
	}
	return true
}

func (sd *SetupDriver) getStatReply(input string) int {
	reply, err := strconv.Atoi(input)
	if err != nil || reply < sd.replyMin || reply > sd.replyMax {
		return 0
	}
	return reply
}
