package server

import (
	"log"

	"github.com/nosborn/federation-1999/internal/server/global"
)

type Config int

const (
	CFG_ADVENTURER_TCR Config = iota // Completed jobs for Adventurer promotion
	CFG_EXPLORER_STAT
	CFG_EXPLORER_TONS
	CFG_GM_PROFIT
	CFG_JP_PROFIT
	CFG_MERCHANT_STATS
	CFG_PLANET_PROMO
	CFG_RP_JOB_CREDITS
	CFG_TRADER_TCR
)

const (
	ADVENTURER_TCR = 500 // Completed jobs for Adventurer promotion
	EXPLORER_STAT  = 120
	EXPLORER_TONS  = 5000
	GM_PROFIT      = 45000000
	JP_PROFIT      = 30000000
	MERCHANT_STATS = 184 // Total stats for Merchant promotion
	PLANET_PROMO   = 10  // Development points to promote a planet
	RP_JOB_CREDITS = 8
	TRADER_TCR     = 1000 // Completed jobs for Trader promotion
)

func getConfig(which Config) int32 {
	switch which {
	case CFG_ADVENTURER_TCR:
		return scale(ADVENTURER_TCR, 10)
	case CFG_EXPLORER_STAT:
		return EXPLORER_STAT
	case CFG_EXPLORER_TONS:
		return scale(EXPLORER_TONS, 10)
	case CFG_GM_PROFIT:
		return scale(GM_PROFIT, 10)
	case CFG_JP_PROFIT:
		return scale(JP_PROFIT, 10)
	case CFG_MERCHANT_STATS:
		return MERCHANT_STATS
	case CFG_PLANET_PROMO:
		return PLANET_PROMO
	case CFG_RP_JOB_CREDITS:
		return RP_JOB_CREDITS
	case CFG_TRADER_TCR:
		return scale(TRADER_TCR, 10)
	default:
		log.Panicf("Unknown configuration parameter %d requested", which)
		return -1 // make linter happy
	}
}

func scale(x, y int32) int32 {
	if global.TestFeaturesEnabled {
		return x / y
	} else {
		return x
	}
}
