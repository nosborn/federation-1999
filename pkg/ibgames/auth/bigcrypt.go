//go:build linux

package auth

/*
#cgo LDFLAGS: -lcrypt
#include <stdlib.h>
#include "bigcrypt.h"
*/
import "C" //nolint:gocritic // CGO import required
import (
	"errors"
	"unsafe" //nolint:gocritic // Required for CGO C.free()
)

var ErrMismatchedHashAndPassword = errors.New("bigcrypt: hashedPassword is not the hash of the given password")

// bigcryptCompareHashAndPassword compares a bigcrypt hash with a password
// Returns nil if they match, or an error if they don't match
func bigcryptCompareHashAndPassword(hashedPassword, password []byte) error {
	keyC := C.CString(string(password))
	saltC := C.CString(string(hashedPassword))
	defer C.free(unsafe.Pointer(keyC))
	defer C.free(unsafe.Pointer(saltC))

	result := C.bigcrypt(keyC, saltC)
	resultStr := C.GoString(result)

	if resultStr == string(hashedPassword) {
		return nil
	}
	return ErrMismatchedHashAndPassword
}
