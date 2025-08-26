package login

import (
	"math/rand/v2"
	"strings"
)

// CheckHoneypot checks if the provided username and password are known
// honeypot credentials.
func CheckHoneypot(username, password string) bool {
	username = strings.ToLower(username)
	passwords, ok := HoneypotCredentials[username]
	if !ok {
		return false
	}
	for _, p := range passwords {
		if p == password {
			if rand.IntN(100) < 30 { // don't match all the time
				return true
			}
			return false
		}
	}
	// log.Printf("new honeypot password: %q %q", username, password)
	return false
}
