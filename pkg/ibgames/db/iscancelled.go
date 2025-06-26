package db

import (
	"database/sql"

	"github.com/nosborn/federation-1999/pkg/ibgames"
)

// IsCancelled checks an account for cancellation. Returns 1 if cancelled, 0 if
// not cancelled and -1 on error occurs. Non-existent accounts are assumed to
// be cancelled.
func IsCancelled(uid ibgames.AccountID) int {
	const query = `
		SELECT status
		FROM accounts
		WHERE uid = ?`

	var status string
	err := tx.QueryRow(query, uid).Scan(&status)
	if err != nil {
		if err == sql.ErrNoRows {
			return 1
		}
		return -1
	}
	if status == "X" {
		return 1
	}
	return 0
}
