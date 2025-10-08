package goodies

import (
	"regexp"
	"sync"
)

var (
	promptRegex *regexp.Regexp
	regexOnce   sync.Once
)

// ContainsPrompt checks if a string contains a login or password prompt. It
// matches patterns like "login:" or "password:" (case-insensitive). Returns
// true if a prompt is found, false otherwise.
func ContainsPrompt(s string) bool {
	if s == "" {
		return false
	}

	regexOnce.Do(func() {
		// Compile the regex pattern: (ogin|word):
		// This matches "login:", "Login:", "password:", "Password:", etc.
		promptRegex = regexp.MustCompile(`(?i)(ogin|word):`)
	})

	return promptRegex.MatchString(s)
}
