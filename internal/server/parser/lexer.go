package parser

import (
	"strconv"
	"strings"

	"github.com/nosborn/federation-1999/internal/text"
)

type token struct {
	typ int
	val string
	num int32
}

type stateFn func(*lexer) stateFn

type lexer struct {
	index     int
	input     string
	nextState stateFn
	player    CommandExecutor
	pos       int
	tokens    []token
}

var keywords = map[string]int{
	"a":              vA,
	"ac":             vAccept,
	"accept":         vAccept,
	"act":            vAct,
	"add":            vAdd,
	"adventurer":     vAdventurer,
	"agri":           vAgricultural,
	"agricultural":   vAgricultural,
	"alarm":          vAlarm,
	"ale":            vAle,
	"allocate":       vAllocate,
	"ammo":           vAmmunition,
	"ammunition":     vAmmunition,
	"an":             vAn,
	"any":            vAny,
	"are":            vAre,
	"armor":          vArmour,
	"armour":         vArmour,
	"at":             vAt,
	"auto":           vAutomatic,
	"automatic":      vAutomatic,
	"bang":           vBang,
	"baroness":       vBaron,
	"bash":           vBash,
	"bay":            vBay,
	"beam":           vBeam,
	"begin":          vStarting,
	"bev":            vBev,
	"black":          vBlack,
	"blast":          vBlast,
	"bless":          vBless,
	"blue":           vBlue,
	"bribe":          vBribe,
	"brief":          vBrief,
	"broadcast":      vBroadcast,
	"brown":          vBrown,
	"build":          vBuild,
	"bump":           vBump,
	"button":         vButton,
	"buy":            vBuy,
	"c":              vCheck,
	"c3":             vC3,
	"captain":        vCaptain,
	"cargo":          vCargo,
	"change":         vChange,
	"channel":        vChannel,
	"cheat":          vCheat,
	"check":          vCheck,
	"clear":          vClear,
	"clothes":        vClothes,
	"com":            vCommunicate,
	"comm":           vCommunicate,
	"commands":       vCommands,
	"comms":          vComms,
	"communicate":    vCommunicate,
	"communicating":  vCommunicating,
	"company":        vCompany,
	"computer":       vComputer,
	"conf":           vConfiguration,
	"config":         vConfiguration,
	"configuration":  vConfiguration,
	"contract":       vContract,
	"cuddle":         vCuddle,
	"cure":           vCure,
	"d":              vDown,
	"deal":           vDeal,
	"deallocate":     vDeallocate,
	"deity":          vDeity,
	"dex":            vDexterity,
	"dexterity":      vDexterity,
	"di":             vDisplay,
	"digest":         vDigest,
	"display":        vDisplay,
	"dividend":       vDividend,
	"down":           vDown,
	"downside":       vDownside,
	"drink":          vDrink,
	"drop":           vDrop,
	"duchies":        vDuchies,
	"duchy":          vDuchy,
	"duty":           vDuty,
	"e":              vEast,
	"east":           vEast,
	"eat":            vEat,
	"education":      vEducation,
	"embargo":        vEmbargo,
	"energy":         vEnergy,
	"engines":        vEngines,
	"enter":          vEnter,
	"ex":             vExamine,
	"examine":        vExamine,
	"exchange":       vExchange,
	"expel":          vExpel,
	"explorer":       vExplorer,
	"fac":            vFactory,
	"factories":      vFactories,
	"factory":        vFactory,
	"favored":        vFavoured,
	"favoured":       vFavoured,
	"fed":            vFederation,
	"federation":     vFederation,
	"fetch":          vFetch,
	"fighting":       vFighting,
	"fire":           vFire,
	"fit":            vFit,
	"flog":           vFlog,
	"food":           vFood,
	"from":           vFrom,
	"fuck":           vFuck,
	"fuel":           vFuel,
	"full":           vFull,
	"g":              vGet,
	"gag":            vGag,
	"gamble":         vGamble,
	"get":            vGet,
	"give":           vGive,
	"gl":             vGlance,
	"glance":         vGlance,
	"go":             vGo,
	"goods":          vGoods,
	"goto":           vGoto,
	"gps":            vGps,
	"groat":          vGroats,
	"groats":         vGroats,
	"grope":          vGrope,
	"gtu":            vGtus,
	"gtus":           vGtus,
	"gun":            vGun,
	"hamsters":       vHamsters,
	"health":         vHealth,
	"help":           vHelp,
	"hgw":            vHgw,
	"host":           vHostess,
	"hostess":        vHostess,
	"hug":            vHug,
	"hull":           vHull,
	"i":              vInventory,
	"ig":             vIg,
	"imperial":       vImperial,
	"in":             vIn,
	"ind":            vIndustrial,
	"indu":           vIndustrial,
	"industrial":     vIndustrial,
	"info":           vInformation,
	"information":    vInformation,
	"infra":          vInfrastructure,
	"infrastructure": vInfrastructure,
	"install":        vInstall,
	"insure":         vInsure,
	"int":            vIntelligence,
	"intel":          vIntelligence,
	"intelligence":   vIntelligence,
	"into":           vInto,
	"inv":            vInventory,
	"inventory":      vInventory,
	"is":             vIs,
	"issue":          vIssue,
	"j":              vJump,
	"jettison":       vJettison,
	"jig":            vJig,
	"job":            vJob,
	"jobs":           vJobs,
	"join":           vJoin,
	"jump":           vJump,
	"kiss":           vKiss,
	"knit":           vKnit,
	"l":              vLook,
	"lamp":           vLamp,
	"land":           vLand,
	"laser":          vLaser,
	"launch":         vLaunch,
	"layoff":         vLayoff,
	"leis":           vLeisure,
	"leisure":        vLeisure,
	"lever":          vLever,
	"link":           vLink,
	"liquidate":      vLiquidate,
	"load":           vLoad,
	"lock":           vLock,
	"look":           vLook,
	"mag":            vMag,
	"manual":         vManual,
	"mark-up":        vMarkup,
	"markup":         vMarkup,
	"matproc":        vMatproc,
	"mattrans":       vMattrans,
	"me":             vMe,
	"milkrun":        vMilkrun,
	"mini":           vMining,
	"mining":         vMining,
	"minutes":        vMinutes,
	"missile":        vMissile,
	"missiles":       vMissiles,
	"mood":           vMood,
	"mov":            vMovement,
	"move":           vMovement,
	"movement":       vMovement,
	"moving":         vMovement,
	"my":             vMy, /* FIX ME! */
	"n":              vNorth,
	"nationalize":    vNationalize,
	"navigate":       vNavigate,
	"navigator":      vNavigator,
	"ne":             vNortheast,
	"none":           vNone,
	"north":          vNorth,
	"northeast":      vNortheast,
	"northwest":      vNorthwest,
	"notice":         vNotice,
	"nw":             vNorthwest,
	"o":              vOut,
	"of":             vOf,
	"off":            vOff,
	"offer":          vOffer,
	"offline":        vOffline,
	"on":             vOn,
	"online":         vOnline,
	"orange":         vOrange,
	"orbit":          vOrbit,
	"order":          vOrder,
	"out":            vOut,
	"output":         vOutput,
	"paint":          vPaint,
	"perivale":       vPerivale,
	"pizza":          vPizza,
	"planet":         vPlanet,
	"play":           vPlay,
	"point":          vPoints,
	"points":         vPoints,
	"post":           vPost,
	"press":          vPress,
	"pri":            vPrice,
	"price":          vPrice,
	"production":     vProduction,
	"project":        vProject,
	"promo":          vPromotional,
	"promotional":    vPromotional,
	"pull":           vPull,
	"push":           vPress,
	"put":            vPut,
	"ql":             vQL,
	"qsc":            vQuickscore,
	"qscore":         vQuickscore,
	"qst":            vQuickstatus,
	"qstatus":        vQuickstatus,
	"quad":           vQuad,
	"quark":          vQuark,
	"quit":           vQuit,
	"qw":             vQuickwho,
	"qwho":           vQuickwho,
	"rack":           vRack,
	"ranks":          vRanks,
	"rate":           vRate,
	"read":           vRead,
	"red":            vRed,
	"registry":       vRegistry,
	"reject":         vReject,
	"repair":         vRepair,
	"repay":          vRepay,
	"report":         vReport,
	"reset":          vReset,
	"reward":         vReward,
	"round":          vRound,
	"routes":         vRoutes,
	"s":              vSouth,
	"salute":         vSalute,
	"say":            vSay,
	"sc":             vScore,
	"score":          vScore,
	"screen":         vScreen,
	"se":             vSoutheast,
	"secede":         vSecede,
	"security":       vSecurity,
	"sell":           vSell,
	"senator":        vSenator,
	"set":            vSet,
	"sh":             vShow,
	"shares":         vShares,
	"shelf":          vShelving,
	"shelves":        vShelving,
	"shelving":       vShelving,
	"shields":        vShields,
	"ship":           vSpaceship,
	"show":           vShow,
	"shuffle":        vShuffle,
	"sign":           vSign,
	"single":         vSingle,
	"slap":           vSlap,
	"slide":          vSlide,
	"slot":           vSlot,
	"snog":           vSnog,
	"snuggle":        vSnuggle,
	"social":         vSocial,
	"south":          vSouth,
	"southeast":      vSoutheast,
	"southwest":      vSouthwest,
	"spaceship":      vSpaceship,
	"spy":            vSpy,
	"spybeam":        vSpybeam,
	"spynet":         vSpynet,
	"spyscreen":      vSpyscreen,
	"st":             vStatus,
	"sta":            vStamina,
	"staff":          vStaff,
	"stamina":        vStamina,
	"start":          vStarting,
	"starting":       vStarting,
	"status":         vStatus,
	"stockpile":      vStockpile,
	"store":          vStore,
	"str":            vStrength,
	"strength":       vStrength,
	"suicide":        vSuicide,
	"sulk":           vSulk,
	"sw":             vSouthwest,
	"switch":         vSwitch,
	"system":         vSystem,
	"systems":        vSystems,
	"take":           vGet,
	"target":         vTarget,
	"tax":            vTax,
	"tb":             vTell,
	"tech":           vTechnological,
	"techno":         vTechnological,
	"technological":  vTechnological,
	"teleport":       vTeleport,
	"tell":           vTell,
	"the":            vThe,
	"tickle":         vTickle,
	"time":           vTime,
	"timeout":        vTimeout,
	"timewarp":       vTimewarp,
	"tl":             vTL,
	"to":             vTo,
	"ton":            vTons,
	"tons":           vTons,
	"torch":          vTorch,
	"tour":           vTour,
	"trace":          vTrace,
	"trade":          vTrade,
	"trader":         vTrader,
	"transmit":       vTransmit,
	"travel":         vTravel,
	"tune":           vTune,
	"twin":           vTwin,
	"type":           vType,
	"u":              vUp,
	"ungag":          vUngag,
	"unload":         vUnload,
	"unlock":         vUnlock,
	"unpost":         vUnpost,
	"unravel":        vUnravel,
	"up":             vUp,
	"upside":         vUpside,
	"use":            vUse,
	"void":           vVoid,
	"w":              vWest,
	"wages":          vWages,
	"wait":           vWait,
	"wanted":         vWanted,
	"ware":           vWarehouse,
	"warehouse":      vWarehouse,
	"warehouses":     vWarehouses,
	"wares":          vWarehouses,
	"west":           vWest,
	"whereis":        vWhereis,
	"who":            vWho,
	"whois":          vWhois,
	"with":           vWith,
	"work":           vWork,
	"xmit":           vTransmit,
	"xt":             vTransmit,

	// I'm sure there's a reason for this being a special snowflake
	"technician": vTechnician,

	// Commodities
	"alloys":          vAlloys,
	"anti":            vAntiMatter,
	"anti-matter":     vAntiMatter,
	"artifacts":       vArtifacts,
	"arts":            vArtifacts,
	"bio":             vBioChips,
	"bio-chips":       vBioChips,
	"cereals":         vCereals,
	"cont":            vControllers,
	"controllers":     vControllers,
	"crys":            vCrystals,
	"crystals":        vCrystals,
	"droids":          vDroids,
	"electros":        vElectros,
	"explosives":      vExplosives,
	"exps":            vExplosives,
	"fruit":           vFruit,
	"furs":            vFurs,
	"games":           vGames,
	"gas":             vGAsChips,
	"gas-chips":       vGAsChips,
	"gen":             vGenerators,
	"generators":      vGenerators,
	"gold":            vGold,
	"hides":           vHides,
	"holos":           vHolos,
	"hypnos":          vHypnotapes,
	"hypnotapes":      vHypnotapes,
	"kats":            vKatydidics,
	"katydidics":      vKatydidics,
	"lanz":            vLanzariK,
	"lanzarik":        vLanzariK,
	"libraries":       vLibraries,
	"libs":            vLibraries,
	"livestock":       vStock,
	"lub-oils":        vLubOils,
	"lubs":            vLubOils,
	"mas":             vMasers,
	"masers":          vMasers,
	"meat":            vMeat,
	"mechparts":       vMechParts,
	"mechs":           vMechParts,
	"monopoles":       vMonopoles,
	"monos":           vMonopoles,
	"mun":             vMunitions,
	"munitions":       vMunitions,
	"musiks":          vMusiks,
	"nickel":          vNickel,
	"nitros":          vNitros,
	"petrochemicals":  vPetros,
	"petros":          vPetros,
	"pharmaceuticals": vPharms,
	"pharms":          vPharms,
	"poly":            vPolymers,
	"polymers":        vPolymers,
	"pow":             vPowerPacks,
	"powerpacks":      vPowerPacks,
	"propellants":     vPropellants,
	"props":           vPropellants,
	"radioactives":    vRads,
	"rads":            vRads,
	"rna":             vRNA,
	"sens":            vSensAmps,
	"sensamps":        vSensAmps,
	"sims":            vSims,
	"simulations":     vSims,
	"soya":            vSoya,
	"spices":          vSpices,
	"stock":           vStock,
	"studios":         vStudios,
	"synths":          vSynths,
	"tex":             vTextiles,
	"textiles":        vTextiles,
	"tools":           vTools,
	"unis":            vUnivators,
	"univators":       vUnivators,
	"vidi":            vVidi,
	"vidicasters":     vVidi,
	"wea":             vWeapons,
	"weapons":         vWeapons,
	"woods":           vWoods,
	"xmet":            vXmetals,
	"xmetals":         vXmetals,
}

func newLexer(input string, player CommandExecutor) *lexer {
	l := &lexer{
		input:     input,
		nextState: lexInitial,
		player:    player,
	}
	l.run()
	return l
}

func (l *lexer) run() {
	for l.nextState != nil {
		l.nextState = l.nextState(l)
	}
}

func lexInitial(l *lexer) stateFn {
	b := l.peek()
	if b == '"' || b == '\'' {
		l.emit(vSay, "say", 0)
		l.next()
		return lexRawText
	}
	if b == '\n' {
		l.emit(int(b), "\n", 0)
		l.next()
		return nil
	}
	return lexCommand
}

func lexCommand(l *lexer) stateFn {
	for text.IsBlank(l.peek()) {
		l.next()
	}

	if b := l.peek(); b == '\n' {
		l.emit('\n', "\n", 0)
		l.next()
		return nil // TODO: lexInitial
	}

	word := l.nextWord()
	// log.Printf("word = %q", word)
	wordLower := text.ToLowerString(word)

	if tok, ok := keywords[wordLower]; ok {
		l.emit(tok, word, 0)
		switch tok {
		case vAct:
			return lexRawText
		case vBroadcast:
			return lexRawText
		case vClothes:
			return lexRawText
		case vCommunicate:
			return lexRawText
		case vCompany:
			return lexRawText
		case vMood:
			return lexMood
		case vNationalize:
			return lexRawText
		case vPlay:
			return lexRawText
		case vPost:
			return lexRawText
		case vReport:
			return lexRawText
		case vRound:
			return lexRawText
		case vSay:
			return lexRawText
		case vSpaceship:
			return lexRawText
		case vTell:
			return lexTell
		case vTransmit:
			return lexRawText
		case vType:
			return lexRawText
		}
		return lexCommand
	}

	if n, ok := parseIDNum(word); ok {
		l.emit(ID_NUM, word, n)
		return lexCommand
	}
	if n, ok := parseInt(word); ok {
		l.emit(vInt, word, n)
		return lexCommand
	}
	if n, ok := parsePercent(word); ok {
		l.emit(vPercent, word, n)
		return lexCommand
	}
	if n, ok := parseFmtInt(word); ok {
		l.emit(PRETTY_INT, word, n)
		return lexCommand
	}
	if n, ok := parseSignedInt(word); ok {
		l.emit(vSignedInt, word, n)
		return lexCommand
	}

	if len(word) > 2 && word[len(word)-2:] == "'s" {
		l.emit(vPossessive, word[:len(word)-2], 0)
		return lexCommand
	}

	l.emit(vToken, word, 0)
	return lexCommand
}

func lexMood(l *lexer) stateFn {
	for text.IsBlank(l.peek()) {
		l.next()
	}

	start := l.pos
	for l.peek() != '\n' {
		l.next()
	}

	extracted := strings.TrimSpace(l.input[start:l.pos])

	if text.ToLowerString(extracted) == "off" {
		l.emit(vOff, extracted, 0)
	} else if extracted != "" {
		l.emit(vText, extracted, 0)
	}

	l.emit('\n', "\n", 0)
	l.next()

	return nil // TODO: lexInitial
}

func lexRawText(l *lexer) stateFn {
	for text.IsBlank(l.peek()) {
		l.next()
	}

	start := l.pos
	for l.peek() != '\n' {
		l.next()
	}

	extracted := strings.TrimSpace(l.input[start:l.pos])

	if extracted != "" {
		l.emit(vText, extracted, 0)
	}

	l.emit('\n', "\n", 0)
	l.next()

	return nil // TODO: lexInitial
}

func lexTell(l *lexer) stateFn {
	for text.IsBlank(l.peek()) {
		l.next()
	}
	b := l.peek()
	if b == '\n' {
		l.emit('\n', "\n", 0)
		l.next()
		return nil // TODO: lexInitial
	}

	word := l.nextWord()
	if word != "" {
		l.emit(vToken, word, 0)
		return lexRawText
	}
	b = l.peek()
	if b == '\n' {
		l.emit('\n', "\n", 0)
		l.next()
		return nil // TODO: lexInitial
	}

	l.emit(int(b), string(b), 0)
	l.next()
	return lexTell
}

func (l *lexer) emit(typ int, val string, num int32) {
	t := token{typ: typ, val: val, num: num}
	l.tokens = append(l.tokens, t)
}

func (l *lexer) next() byte {
	if l.pos >= len(l.input) {
		return 0
	}
	b := l.input[l.pos]
	l.pos++
	return b
}

func (l *lexer) backup() {
	if l.pos > 0 {
		l.pos--
	}
}

func (l *lexer) peek() byte {
	b := l.next()
	l.backup()
	return b
}

func (l *lexer) nextWord() string {
	start := l.pos
	for {
		b := l.next()
		// if b == 0 {
		// 	return l.input[start:l.pos]
		// }
		// if b == ' ' || b == '\t' || b == '\n' {
		// 	return l.input[start : l.pos-1]
		// }
		if !text.IsGraph(b) {
			return l.input[start : l.pos-1]
		}
	}
}

// isRawInt checks if string matches flex pattern [0-9]+
func isRawInt(s string) bool {
	for i := range s {
		if !text.IsDigit(s[i]) {
			return false
		}
	}
	return true
}

func parseFmtInt(s string) (int32, bool) {
	// Check if string matches flex pattern [0-9]{1,3}(,[0-9]{3})+
	// Must start with 1-3 digits
	i := 0
	for i < 3 && i < len(s) && text.IsDigit(s[i]) {
		i++
	}
	if i == 0 {
		return 0, false // no initial digits
	}

	// Remaining length must be divisible by 4 (,###,###,...)
	remaining := len(s) - i
	if remaining == 0 {
		return 0, false // no comma groups = not formatted
	}
	if remaining%4 != 0 {
		return 0, false // invalid length
	}

	// Check each 4-character block: ,### pattern
	for j := i; j < len(s); j += 4 {
		if s[j] != ',' ||
			!text.IsDigit(s[j+1]) ||
			!text.IsDigit(s[j+2]) ||
			!text.IsDigit(s[j+3]) {
			return 0, false
		}
	}

	// Valid format - strip commas and parse
	return parseInt(strings.ReplaceAll(s, ",", ""))
}

func parseIDNum(s string) (int32, bool) {
	if s[0] != '#' {
		return 0, false
	}
	if !isRawInt(s[1:]) {
		return 0, false
	}
	return parseInt(s[1:])
}

func parseInt(s string) (int32, bool) {
	n, err := strconv.ParseInt(s, 10, 32)
	return int32(n), err == nil
}

func parsePercent(s string) (int32, bool) {
	if s[len(s)-1] != '%' {
		return 0, false
	}
	return parseInt(s[:len(s)-1])
}

func parseSignedInt(s string) (int32, bool) {
	if s[0] != '+' && s[0] != '-' {
		return 0, false
	}
	if !isRawInt(s[1:]) {
		return 0, false
	}
	return parseInt(s)
}

func (l *lexer) Lex(lval *yySymType) int {
	if l.index >= len(l.tokens) {
		return 0
	}
	tok := l.tokens[l.index]
	lval.psz = tok.val
	lval.i = tok.num // FIXME
	lval.l = tok.num // FIXME
	l.index++
	return tok.typ
}

func (l *lexer) Error(s string) {
	//
}

// Try to determine whether the supplied name is recognised by the parser. Such
// names tend to cause trouble and/or confusion, which is obviously A Bad
// Thing.
//
// This could be somewhat more sophisticated!
func IsReservedWord(name string) bool {
	_, isReserved := keywords[text.ToLowerString(name)]
	return isReserved
}
