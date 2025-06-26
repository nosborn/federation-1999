package text

import (
	"testing"
)

func TestHumanizeMinutes(t *testing.T) {
	testCases := []struct {
		name     string
		minutes  int64
		expected string
	}{
		{"zero minutes", 0, "0 minutes"},
		{"single minute", 1, "1 minute"},
		{"multiple minutes", 59, "59 minutes"},
		{"single hour", 60, "1 hour"},
		{"single hour and minute", 61, "1 hour 1 minute"},
		{"single hour and multiple minutes", 62, "1 hour 2 minutes"},
		{"multiple hours", 120, "2 hours"},
		{"multiple hours and single minute", 121, "2 hours 1 minute"},
		{"multiple hours and multiple minutes", 130, "2 hours 10 minutes"},
		{"large number of minutes", 1440, "24 hours"},
		{"negative input", -1, ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := HumanizeMinutes(tc.minutes)
			if got != tc.expected {
				t.Errorf("HumanizeMinutes(%d) = %q; want %q", tc.minutes, got, tc.expected)
			}
		})
	}
}
