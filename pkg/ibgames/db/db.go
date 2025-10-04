// Package db provides a database connection layer that emulates Informix
// ESQL/C semantics for SQLite. It maintains single global connection state and
// auto-transactions, matching the original 1999 Informix implementation.
//
// The package follows the Informix pattern where:
// - Connect establishes a database connection to a hardcoded database name
// - All operations run within auto-managed transactions
// - Commit/Rollback restart transactions automatically
// - Disconnect cleanly closes the connection
// - Exit abandons the connection (for process termination)
package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db   *sql.DB   // Global database connection
	conn *sql.Conn // Single connection from the pool
	tx   *sql.Tx   // Current auto-transaction
)

// Commit commits the current transaction and immediately starts a new one,
// matching Informix ANSI auto-transaction behaviour.
func Commit() error {
	err := tx.Commit()
	if err != nil {
		return err
	}
	tx = nil
	return startTransaction()
}

// Connect establishes a connection to the accounts database. Like the original
// Informix implementation, the database name is hardcoded and the location is
// determined by the DBPATH environment variable. Only one connection can be
// active at a time.
func Connect(readOnly bool) error {
	if db != nil {
		return fmt.Errorf("already open")
	}

	dbPath, ok := os.LookupEnv("DBPATH")
	if !ok {
		return fmt.Errorf("DBPATH not set")
	}
	dbFile := filepath.Join(dbPath, "ibgames.sqlite")

	dsnParams := []string{
		"_pragma=automatic_index(0)",
		"_pragma=busy_timeout(5000)",
		"_pragma=foreign_keys(1)",
		"_pragma=journal_mode(WAL)",
		"_time_format=sqlite",
	}
	if readOnly {
		dsnParams = append(dsnParams, "_pragma=query_only(1)") // not truly read-only
	}

	dsn := fmt.Sprintf("file:%s?%s", dbFile, strings.Join(dsnParams, "&"))
	var err error
	db, err = sql.Open("sqlite3", dsn)
	if err != nil {
		return err
	}

	conn, err = db.Conn(context.Background())
	if err != nil {
		return err
	}

	return startTransaction()
}

// Detach detaches from the database server, equivalent to Informix
// sqldetach(). Used in fork/exec scenarios where child processes need to clean
// up inherited database connections without affecting the parent process.
func Detach() int {
	// TODO: Implement if needed for process forking scenarios
	return -1
}

// Disconnect cleanly closes the database connection after committing the
// current transaction. Like Informix DISCONNECT CURRENT, this resets all
// connection state and allows a new Connect() to be called.
func Disconnect() error {
	if db == nil {
		return nil
	}
	if err := tx.Commit(); err != nil { // DISCONNECT issues a COMMIT in Informix.
		return err
	}
	tx = nil
	if err := conn.Close(); err != nil {
		return err
	}
	conn = nil
	if err := db.Close(); err != nil {
		return err
	}
	db = nil
	return nil
}

// Exec executes a SQL statement within the current auto-transaction.
func Exec(query string, args ...any) (sql.Result, error) {
	return tx.Exec(query, args...)
}

// Exit abandons the database connection, equivalent to Informix sqlexit().
// This rolls back any open transaction and closes the database without error
// handling, typically used during process termination.
func Exit() error {
	if tx != nil {
		_ = tx.Rollback() // sqlexit() rolls back open transactions
		tx = nil
	}
	if db != nil {
		_ = db.Close() // sqlexit() closes databases, ignore errors
	}
	conn, db = nil, nil
	return nil
}

// Prepare prepares a SQL statement for repeated execution. The statement is
// prepared on the connection, not the transaction.
func Prepare(query string) (*sql.Stmt, error) {
	return conn.PrepareContext(context.Background(), query)
}

// QueryRow executes a query that returns at most one row within the current
// auto-transaction.
func QueryRow(query string, args ...any) *sql.Row {
	return tx.QueryRow(query, args...)
}

// Rollback rolls back the current transaction and immediately starts a new
// one, matching Informix ANSI auto-transaction behaviour.
func Rollback() error {
	if err := tx.Rollback(); err != nil {
		return err
	}
	tx = nil
	return startTransaction()
}

// startTransaction begins a new transaction on the connection. This is called
// automatically after Connect, Commit, and Rollback to maintain the Informix
// ANSI auto-transaction semantics.
func startTransaction() error {
	var err error
	tx, err = conn.BeginTx(context.Background(), nil)
	return err
}
