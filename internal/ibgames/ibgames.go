package ibgames

import "math"

type AccountID uint32

const (
	// The lowest and highest real account IDs. Values below 100000 were
	// used for Federation personas moved from AOL and don't have
	// corresponding account details.
	MIN_ACCOUNT_ID = 100000
	MAX_ACCOUNT_ID = math.MaxInt32 // NOT MaxUint32
)
