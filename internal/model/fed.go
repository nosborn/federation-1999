package model

import "math"

const (
	GENERAL_JOB_AGING = 5
	JOB_DELIVERY_AGE  = 20
	MAX_XT_CHANNEL    = 26
	MIN_IB_RANK       = RankSenator
)

const (
	MOOD_SIZE = 36
	NAME_SIZE = 16
)

const (
	COMPANY_NAME_SIZE = 32
	EVENT_DESC_SIZE   = 160
	OBJECT_DESC_SIZE  = 82
	OBJECT_SCAN_SIZE  = 201
	SHIP_DESC_SIZE    = 160
)

const (
	MAX_GUNS   = 8  // Max number of guns a ship can carry
	MAX_LOAD   = 15 // Max cargo loads a ship can carry
	MAX_STORES = 10
)

const (
	MAX_BALANCE = math.MaxInt32 - 1
	MIN_BALANCE = math.MinInt32 + 1
)

const (
	MAX_INPUTS = 6 // Maximum inputs to make a commodity
)

const ( // Game constants
	DESC_SIZE    = 1024  // size of location text
	EVENT_SIZE   = 2     // max number of events per location
	FIRST_COMMOD = 10000 // vocab number of first commodity
	LAST_COMMOD  = 10051 // vocab number of last commodity
	MAX_FACTORY  = 12    // Number of factories in persona file record
	MAX_HOARDING = 120   // Maximum object hoarding time
	MAX_PER_DESC = 152   // Max size of player desription field
	MIN_HOARDING = 90    // Minimum object hoarding time (minutes)
	SHIP_START   = 426   // Sol start location for new space ships
)

const (
	SPYBEAM_COST   = 10000000 // 10,000,000 IG
	SPYBEAM_RESALE = 5000000  //  5,000,000 IG
	SPYSCREEN_COST = 50000000 // 50,000,000 IG
)

const (
	SPYBEAM_WEIGHT = 50 // 50 tons
)

const ( // The number of seconds in a 24 hour period.
	SECS_IN_A_DAY = 24 * 60 * 60
)
