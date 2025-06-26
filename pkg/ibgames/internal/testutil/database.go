// Package testutil provides common testing utilities for the project.
package testutil

import (
	"database/sql"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nosborn/federation-1999/pkg/ibgames"
	"github.com/stretchr/testify/require"
)

// DatabaseSetup holds references to test database resources.
type DatabaseSetup struct {
	FilePath string
	TestDB   *sql.DB
}

// SetupTestDatabase creates a temporary SQLite database for testing.
// It returns a direct SQL connection for test verification.
// The caller is responsible for setting up any package-specific connections.
func SetupTestDatabase(t *testing.T) *DatabaseSetup {
	// Create temporary directory for test database
	tmpDir := t.TempDir()
	os.Setenv("DBPATH", tmpDir)

	// Create the expected database file name
	dbPath := filepath.Join(tmpDir, "ibgames.sqlite")

	// Create direct connection for test verification
	testDB, err := sql.Open("sqlite3", dbPath)
	require.NoError(t, err)

	setup := &DatabaseSetup{
		FilePath: dbPath,
		TestDB:   testDB,
	}

	t.Cleanup(func() {
		setup.Teardown()
	})

	return setup
}

// SetupTestDatabaseWithSchema creates a test database and initializes it with the accounts schema.
func SetupTestDatabaseWithSchema(t *testing.T) *DatabaseSetup {
	setup := SetupTestDatabase(t)
	setup.CreateSchema(t)
	return setup
}

// CreateSchema creates the accounts and sessions tables using the production schema file.
func (d *DatabaseSetup) CreateSchema(t *testing.T) {
	// Close the Go connection temporarily to avoid conflicts
	d.TestDB.Close()

	// Find the schema file - try different possible locations
	schemaPath := "ibgames.sql"
	if _, err := os.Stat(schemaPath); os.IsNotExist(err) {
		schemaPath = "../ibgames.sql"
		if _, err := os.Stat(schemaPath); os.IsNotExist(err) {
			schemaPath = "../../ibgames.sql"
			if _, err := os.Stat(schemaPath); os.IsNotExist(err) {
				require.NoError(t, err, "Could not find ibgames.sql schema file")
			}
		}
	}

	// Use sqlite3 CLI to load the production schema
	cmd := exec.Command("sqlite3", d.FilePath, ".read "+schemaPath) // #nosec G204
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Logf("sqlite3 output: %s", string(output))
		require.NoError(t, err, "Failed to load schema using sqlite3")
	}

	// Reopen the Go connection
	d.TestDB, err = sql.Open("sqlite3", d.FilePath)
	require.NoError(t, err)
}

// CreateTestAccount inserts a test account into the database.
func (d *DatabaseSetup) CreateTestAccount(t *testing.T, uid ibgames.AccountID, name, complimentary string, minutes int) {
	// For test accounts, name_key should be lowercase version of name
	// since UniqueName converts to lowercase and removes non-graphic chars
	nameKey := strings.ToLower(name)
	_, err := d.TestDB.Exec(`
		INSERT INTO accounts (uid, name, name_key, encrypt, complimentary, minutes)
		VALUES (?, ?, ?, 'dummy_hash', ?, ?)
	`, uid, name, nameKey, complimentary, minutes)
	require.NoError(t, err)
}

// Teardown cleans up database resources.
func (d *DatabaseSetup) Teardown() {
	if d.TestDB != nil {
		d.TestDB.Close()
	}
	if d.FilePath != "" {
		os.Remove(d.FilePath)
	}
}
