package auth

import "crypto/rand"

func RandomKey() string {
	return rand.Text()
}
