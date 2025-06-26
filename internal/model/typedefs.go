package model

type Commodity int32

const (
	CommodityGAsChips Commodity = iota
	CommodityBioChips
	CommodityMasers
	CommodityWeapons
	CommodityVidi
	CommodityElectros
	CommodityTools
	CommoditySynths
	CommodityDroids
	CommodityAntiMatter
	CommodityPowerPacks
	CommodityControllers
	CommodityGenerators
	CommodityPolymers
	CommodityLubOils
	CommodityPharms
	CommodityPetros
	CommodityRNA
	CommodityPropellants
	CommodityExplosives
	CommodityLanzariK
	CommodityNitros
	CommodityMunitions
	CommodityMechParts
	CommodityCereals
	CommodityWoods
	CommodityHides
	CommodityTextiles
	CommodityMeat
	CommoditySpices
	CommodityFruit
	CommoditySoya
	CommodityStock
	CommodityFurs
	CommodityRads
	CommodityNickel
	CommodityXmetals
	CommodityCrystals
	CommodityAlloys
	CommodityGold
	CommodityMonopoles
	CommodityHypnotapes
	CommodityStudios
	CommoditySensAmps
	CommodityGames
	CommodityArtifacts
	CommodityKatydidics
	CommodityMusiks
	CommodityLibraries
	CommodityHolos
	CommodityUnivators
	CommoditySims
)

type CommodityGroup int32

const (
	GroupNone CommodityGroup = -1
)

const (
	GroupAgricultural CommodityGroup = 1 + iota
	GroupMining
	GroupIndustrial
	GroupTechnological
	GroupLeisure
)

type Delivery int32

const ( // Factory delivery points
	DeliverExchange Delivery = iota
	DeliverWarehouse
	DeliverFactory
)

type Direction int32

const (
	DirectionNorth Direction = iota
	DirectionNE
	DirectionEast
	DirectionSE
	DirectionSouth
	DirectionSW
	DirectionWest
	DirectionNW
	DirectionUp
	DirectionDown
	DirectionIn
	DirectionOut
	DirectionPlanet
)

type EventResult int

const (
	EventContinue EventResult = iota
	EventStop
)

type HookResult int

const (
	HookContinue HookResult = iota
	HookStop
)

type Level int

const (
	LevelNoProduction Level = iota
	LevelAgricultural
	LevelMining
	LevelIndustrial
	LevelTechnological
	LevelLeisure
	LevelCapital
)

type Project int32

const (
	// Explorer builds
	ProjectLink Project = 1 + iota

	// Planetary investment builds
	ProjectEducation
	ProjectEnergy
	ProjectHealth
	ProjectInfra
	ProjectSecurity

	// Duke puzzle builds
	ProjectUpside
	ProjectDownside
	ProjectAccel
	ProjectMatProc
	ProjectC3
	ProjectMatTrans
	ProjectTimeMach
)

type Rank int32

const ( // FIXME: possibly not right
	RankGroundHog Rank = iota
	RankCommander
	RankCaptain
	RankAdventurer
	RankTrader
	RankMerchant
	RankJP
	RankGM
	RankExplorer
	RankSquire
	RankThane
	RankIndustrialist
	RankTechnocrat
	RankBaron
	RankDuke
	RankSenator
	RankMascot
	RankHostess
	RankManager
	RankDeity
	RankEmperor
)

type Sex int32

const (
	SexFemale Sex = 'f'
	SexMale   Sex = 'm'
	SexNeuter Sex = 'n'
)

// type Message struct {
// 	Expand bool
// 	Text   string
// }
//
// type MsgId uint32

type Trace int32

const (
	TraceNone Trace = iota
	TracePerivale
)
