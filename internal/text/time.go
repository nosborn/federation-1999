package text

import (
	"fmt"
	"strings"
)

// HumanizeMinutes converts a total number of minutes into a human-readable
// string like "X hours Y minutes".
//
// It handles the following cases:
// - Omits zero parts (e.g., 120 minutes is "2 hours", not "2 hours 0 minutes").
// - Correctly pluralizes "hour/hours" and "minute/minutes".
// - Handles the zero case gracefully (e.g., 0 minutes is "0 minutes").
func HumanizeMinutes(totalMinutes int64) string {
	if totalMinutes < 0 {
		return ""
	}

	if totalMinutes == 0 {
		return "0 minutes" // TODO: use Msg()
	}

	hours := totalMinutes / 60
	minutes := totalMinutes % 60

	var parts []string

	if hours > 0 {
		if hours == 1 {
			parts = append(parts, "1 hour") // TODO: use Msg()
		} else {
			parts = append(parts, fmt.Sprintf("%d hours", hours)) // TODO: use Msg()
		}
	}

	if minutes > 0 {
		if minutes == 1 {
			parts = append(parts, "1 minute") // TODO :use Msg()
		} else {
			parts = append(parts, fmt.Sprintf("%d minutes", minutes)) // TODO: use Msg()
		}
	}

	return strings.Join(parts, " ")
}
