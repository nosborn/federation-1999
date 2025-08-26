package limits

import (
	"math"
)

const ( // FIXME: this doesn't belong here
	MIN_ACCOUNT_ID = 100000
	MAX_ACCOUNT_ID = math.MaxInt32
)

const ( // FIXME: this doesn't belong here
	LOCATION_LIMIT = 120 // for player planets
)
