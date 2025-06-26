package auth

import (
	"net"
	"strings"
	"time"

	"github.com/mattn/go-sqlite3"
	"github.com/nosborn/federation-1999/pkg/ibgames"
	"github.com/nosborn/federation-1999/pkg/ibgames/db"
)

func CreateCookie(addr net.Addr, uid ibgames.AccountID, sid *string) CookieResult {
	expire := time.Now().Add(30 * time.Minute).Unix()
	ipAddress := addr.String()

	// Handle IPv6 addresses by extracting just the IP part
	if tcpAddr, ok := addr.(*net.TCPAddr); ok {
		ipAddress = tcpAddr.IP.String()
	}

	// Retry loop to handle duplicate session IDs
	for {
		sessionKey := RandomKey()

		const query = `
			INSERT INTO cookies (sid, ip_address, uid, expire)
			VALUES (?, ?, ?, ?)`

		_, err := db.Exec(query, sessionKey, ipAddress, uid, expire)
		if err != nil {
			// Check if this is a duplicate key error
			if isDuplicateKeyError(err) {
				// Retry with a new random key
				continue
			}
			// Other error - return failure
			return CookieError
		}

		// Success - set the session ID and return
		*sid = sessionKey
		return CookieOK
	}
}

// isDuplicateKeyError checks if the error is a SQLite duplicate key constraint violation
func isDuplicateKeyError(err error) bool {
	if sqliteErr, ok := err.(sqlite3.Error); ok {
		return sqliteErr.ExtendedCode == sqlite3.ErrConstraintPrimaryKey ||
			sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique
	}

	// Fallback to string matching for other error types
	errStr := strings.ToLower(err.Error())
	return strings.Contains(errStr, "unique constraint") ||
		strings.Contains(errStr, "primary key") ||
		strings.Contains(errStr, "duplicate")
}
