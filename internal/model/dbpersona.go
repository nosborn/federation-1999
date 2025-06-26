package model

// defines for access to DBPersona.Count[]
const (
	PL_G_TIMER = iota
	PL_G_JOB
	PL_G_MAINT
	PL_G_MOVES
)

type DBPersona struct { // File storage for player record
	// Personal info
	Name [NAME_SIZE]byte
	ID   uint32
	Sex  byte               // male or female
	_    [3]byte            // padding for alignment
	Rank uint32             // game level
	Desc [MAX_PER_DESC]byte // description of player
	Mood [MOOD_SIZE]byte    // xtra text to go before entry desc

	// Statistics
	MaxStr  uint32 // strength
	CurStr  uint32
	MaxSta  uint32 // stamina
	CurSta  uint32
	MaxInt  uint32 // interlligence
	CurInt  uint32
	MaxDex  uint32 // dexterity
	CurDex  uint32
	Shipped uint32    // tons of freight shipped by other players
	Games   uint32    // number of games played
	Flags   [2]uint32 // player flags

	// Money, etc
	Balance int32
	Loan    int32 // outstanding loan
	Reward  int32 // price on your head!
	Frame   [4]int32

	// Locations
	LocNo      uint32          // current location of persona
	StarSystem [NAME_SIZE]byte // name of star system player is in

	// Work
	Job DBWork // job player is current contracted to

	// Trading
	LastTrade int32 // trade goods just offered!

	// Spaceship
	ShipLoc  uint32 // current ship location
	ShipDesc [SHIP_DESC_SIZE]byte
	Registry [NAME_SIZE]byte
	ShipKit  DBEquipment // where your ship is registered
	Missiles uint32      // number of missiles carried
	Ammo     uint32      // number of shots of mag gun ammo carried
	Guns     [MAX_GUNS]SGuns
	Load     [MAX_LOAD]DBCargo

	// Misc game stuff
	LastOn int32    // FIXME: year 2038 problem
	Count  [4]int32 // player's counters
	Build  DBBuild

	PP *DBPPData // Poor People
	RP *DBRPData // Rich People
}

type DBWork struct {
	Pallet  DBCargo
	JobType int32
	From    [NAME_SIZE]byte
	To      [NAME_SIZE]byte
	Status  int32
	Value   int32
	GTU     int32
	Credits int32
	Type    DBWorkType // may contain factory job or general job
	Age     int32
}

type DBCargo struct {
	Type     int32
	Quantity int32
	Origin   [NAME_SIZE]byte
	Cost     int32
}

type DBWorkType struct {
	FactryWk DBFactoryJob
	GenWk    DBGeneralJob
}

type DBFactoryJob struct {
	Deliver DBFactoryID
	PickUp  DBFactoryID
}

type DBFactoryID struct {
	Number int32
	Owner  [COMPANY_NAME_SIZE]byte
}

type DBGeneralJob struct {
	WhereTo int32
	Owner   [NAME_SIZE]byte
}

type DBEquipment struct { // Ship's equipment record
	MaxHull     uint32 // Hull strength
	CurHull     uint32
	MaxShield   uint32 // Shield strength
	CurShield   uint32
	MaxEngine   uint32 // Engine capacity
	CurEngine   uint32
	MaxComputer uint32 // Computer size
	CurComputer uint32
	MaxFuel     uint32 // Fuel
	CurFuel     uint32
	MaxHold     uint32 // Cargo capacity
	CurHold     uint32
	Tonnage     uint32 // Overall size of ship
}

type DBBuild struct {
	IDProject uint16
	_         uint16 // padding for alignment
	Duration  int32
	Elapsed   int32
}

type DBPPData struct {
	GMLocation uint32 // loc at which player will find the GM
	Storage    [MAX_STORES]DBWarehouse
	Company    DBCompany
	Factory    [MAX_FACTORY]DBFactory
}

type DBWarehouse struct {
	Planet [NAME_SIZE]byte
	Bay    [20]DBCargo
}

type DBCompany struct {
	Name    [COMPANY_NAME_SIZE]byte // Name of company
	Balance int32                   // Company's balance
	Capital int32                   // Number of shares issued (TODO: could be int16)
	Income  int32                   // Income this account cycle
	Expend  int32                   // Expenditure this account cycle
}

type DBFactory struct { // Structure to hold factory data
	// Identification.
	Planet [NAME_SIZE]byte // Name of planet where factory is located
	_      int32           // OBSOLETE: was factory identifier number

	// Inputs.
	Product int32             // Index to goods produced
	Wages   int32             // Wage paid to workers
	Layoff  int32             // Layoff pay of workers as % of wage
	Stock   [MAX_INPUTS]int32 // Stocks of input goods.

	// Production.
	Cycle int32 // Where in the production cycle we are

	// Money.
	Income int32 // Income so far this cycle
	Expend int32 // Expenditure so far this cycle

	// Finished goods.
	Delivery   int32 // where to deliver the finished goods
	Contracted int32 // Quantity to be delivered to another factory
	Delivered  int32 // Amount delivered so far
	Carriage   int32 // Price/ton to pay hauler
	ToFactory  int32
	OpStock    int32 // Quantity of finished stock stored
}

type DBRPData struct {
	Fief       [NAME_SIZE]byte // name of player's own planet/duchy
	Planet     DBPlanet
	Storage    DBWarehouse
	Facilities [7]int32 // duke puzzle progress
	Duchy      DBDuchy
}

type DBPlanet struct { // Structure to hold planet & exchange data.
	// General info.
	Duchy      [NAME_SIZE]byte // Name of duchy to which planet belongs
	Level      int32           // planet's development level
	Population int32           // Current population index
	Tax        int32           // Percentage base tax rate for the planet
	Time       int32           // Player minutes spent on the planet
	Balance    int32           // Treasury balance
	Flags      uint32
	LastOnline int32 // Day planet was last online

	// Player set stuff.
	Jobs [4]DBPlanetJob // owner-set milk run jobs

	// Exchange info.
	ExTime     int32       // 0-100 time into exchange production cycle
	Flux       int32       // random price fluctuations (+40% to -40%)
	Markup     int32       // Percentage markup on this exchange
	Production [5][2]int32 // moveable allocation production
	Goods      [52]DBBin

	// Development stuff.
	Energy       int32 // Energy investment
	Education    int32 // Education investment
	SocialSec    int32 // Social security
	Infstr       int32 // Infrastructure investment: rail/roads/etc
	Health       int32 // Population healthcare
	IntSecurity  int32 // Internal security - police, courts etc
	Disaffection int32 // Population disaffect as a percentage
}

type DBPlanetJob struct {
	Name      [NAME_SIZE]byte
	Commodity int32
	Carriage  int32
}

type DBBin struct {
	Production  int32 // Total production per cycle
	Produce     int32 // Produce this cycle?
	Stock       int32 // Stock on hand
	Stockpile   int32 // Stockpile level
	Consumption int32 // Total consumption per cycle
	Markup      int32 // Percentage markup on this commodity
	Produced    int32 // Production so far this cycle
	Consumed    int32 // Consumption so far this cycle
}

type DBDuchy struct {
	CustomsRate  int32           // Duchy's basic customs rate
	Favoured     [NAME_SIZE]byte // Planet with favoured trading status
	FavouredRate int32           // Tax rate for favoured planet
	Embargo      [NAME_SIZE]byte // Embargoed planet/system
}
