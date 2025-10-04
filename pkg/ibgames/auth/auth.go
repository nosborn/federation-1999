package auth

import "github.com/nosborn/federation-1999/pkg/ibgames"

const (
	// AUTH_ENCRYPT_SIZE         = (AUTH_SALT_SIZE + ((AUTH_PASSWORD_SIZE / 8) * 11))
	NameSize     = 32
	PasswordSize = 72 // bcrypt.GenerateFromPassword does not accept passwords longer than 72 bytes
	// AUTH_RANDOM_KEY_SIZE      = 32
	// AUTH_RANDOM_PASSWORD_SIZE = (AUTH_RANDOM_KEY_SIZE + 3)
	// AUTH_SALT_SIZE            = 2
)

const (
	maxPasswordTries = 20
)

type CookieResult int

const (
	CookieOK CookieResult = iota
	CookieError
	CookieNotFound
)

type LoginResult int

const (
	LoginOK        LoginResult = iota // Valid login
	LoginError                        // Catch-all internal error
	LoginIncorrect                    // Name or password wrong
	LoginNoCredit                     // Account has no credit
	LoginSuspended                    // Account has been suspended
)

type Session struct {
	UID     ibgames.AccountID
	SLogin  string
	ULogin  string
	SucIP   string
	UnsucIP string
}
