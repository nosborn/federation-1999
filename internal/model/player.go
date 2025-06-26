package model

const (
	PL0_FLYING           = 0x00000001 // in ship command mode
	PL0_BRIEF            = 0x00000002 // brief descriptions only
	PL0_COMM_UNIT        = 0x00000008 // player has a functioning comms unit
	PL0_LIT              = 0x00000020 // player carrying a light source
	PL0_INSURED          = 0x00000100 // player is insured
	PL0_JOB              = 0x00000400 // tell player about jobs
	PL0_INFO             = 0x00002000 // inform player of logon/offs
	PL0_SULKING          = 0x00004000 // player is sulking
	PL0_SPYBEAM          = 0x00100000 // player has spybeam installed
	PL0_OFFER_TOUR       = 0x00200000 // Offer tour to new Commander
	PL0_SPYSCREEN        = 0x00800000 // player has a spybeam shield
	PL0_KNOWS_ORSONITE   = 0x20000000 // player knows about TDX (Horsell)
	PL0_SNARK_ASSIGNED   = 0x40000000 // Has assignment for Snark
	PL0_HORSELL_ASSIGNED = 0x80000000 // Has assignment for Horsell
)

const (
	PL1_SHIP_PERMIT = 0x00000001 // Has permit to operate ship
	PL1_SHIELDS     = 0x00000008 // shields are switched on
	PL1_AUTO        = 0x00000010 // battle computer on auto
	PL1_MI6_OFFERED = 0x00000080 // Commission in naval int was offered
	PL1_MI6         = 0x00000100 // Has commission in naval int
	PL1_TITAN       = 0x00000200 // For use with GOTO cmnd
	PL1_CALLISTO    = 0x00000400
	PL1_MARS        = 0x00000800
	PL1_EARTH       = 0x00001000
	PL1_MOON        = 0x00002000
	PL1_VENUS       = 0x00004000
	PL1_MERCURY     = 0x00008000
	PL1_HILBERT     = 0x00200000 // n-space converter fitted
	PL1_DONE_STA    = 0x00800000 // Done stamina puzzle
	PL1_DONE_STR    = 0x01000000 // Done strength puzzle
	PL1_DONE_DEX    = 0x02000000 // Done dexterity puzzle
	PL1_DONE_INT    = 0x04000000 // Done intelligence puzzle
	PL1_PO_PERMIT   = 0x08000000 // Has PO's permit
	PL1_DONE_SNARK  = 0x20000000 // Done Snark puzzle
	PL1_NAVIGATOR   = 0x40000000 // Is a DataSpace Navigator
	PL1_PROMO_CHAR  = 0x80000000 // Is a promotional character
)

var FemaleRanks = [...]string{
	"GroundHog",         //  0 GroundHog
	"Commander",         //  1 Commander
	"Captain",           //  2 Captain
	"Adventuress",       //  3 Adventurer
	"Trader",            //  4 Trader
	"Merchant",          //  5 Merchant
	"Journeywoman",      //  6 JP
	"Guild Mistress",    //  7 GM
	"Explorer",          //  8 Explorer
	"Squire",            //  9 Squire
	"Thane",             // 10 Thane
	"Industrialist",     // 11 Industrialist
	"Technocrat",        // 12 Technocrat
	"Baroness",          // 13 Baron
	"Duchess",           // 14 Duke
	"Senator",           // 15 Senator
	"Mascot",            // 16 Mascot
	"DataSpace Hostess", // 17 Hostess
	"Demi-Deity",        // 18 Manager
	"Deity",             // 19 Deity
	"Empress",           // 20 Emperor
}

var MaleRanks = [...]string{
	"GroundHog",      //  0 GroundHog
	"Commander",      //  1 Commander
	"Captain",        //  2 Captain
	"Adventurer",     //  3 Adventurer
	"Trader",         //  4 Trader
	"Merchant",       //  5 Merchant
	"Journeyman",     //  6 JP
	"Guild Master",   //  7 GM
	"Explorer",       //  8 Explorer
	"Squire",         //  9 Squire
	"Thane",          // 10 Thane
	"Industrialist",  // 11 Indistrialist
	"Technocrat",     // 12 Technocrat
	"Baron",          // 13 Baron
	"Duke",           // 14 Duke
	"Senator",        // 15 Senator
	"Mascot",         // 16 Mascot
	"DataSpace Host", // 17 Hostess
	"Demi-Deity",     // 18 Manager
	"Deity",          // 19 Deity
	"Emperor",        // 20 Experor
}

var NeuterRanks = [...]string{
	"GroundHog",      //  0 GroundHog
	"Commander",      //  1 Commander
	"Captain",        //  2 Captain
	"Adventuroid",    //  3 Adventurer
	"Trader",         //  4 Trader
	"Merchant",       //  5 Merchant
	"Journeything",   //  6 JP
	"Guild Maitroid", //  7 GM
	"Explorer",       //  8 Explorer
	"Squire",         //  9 Squire
	"Thane",          // 10 Thane
	"Industrialist",  // 11 Industrialist
	"Technocrat",     // 12 Technocrat
	"Baronoid",       // 13 Baron
	"Duchoid",        // 14 Duke
	"Senator",        // 15 Senator
	"Mascot",         // 16 Mascot
	"DataSpace Host", // 17 Hostess
	"Demi-Deity",     // 18 Manager
	"Deity",          // 19 Deity
	"Emperor",        // 20 Emperor
}
