//go:build !linux

package auth

import "errors"

var ErrMismatchedHashAndPassword = errors.New("bigcrypt: hashedPassword is not the hash of the given password")

// bigcryptCompareHashAndPassword is not supported on this platform
func bigcryptCompareHashAndPassword(hashedPassword, password []byte) error {
	return ErrMismatchedHashAndPassword
}
