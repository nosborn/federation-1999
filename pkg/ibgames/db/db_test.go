package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) {
	tmpDir := t.TempDir()
	os.Setenv("DBPATH", tmpDir)

	t.Cleanup(func() {
		Exit()
	})
}

func TestDatabaseOperations(t *testing.T) {
	setupTestDB(t)

	err := Connect(false)
	require.NoError(t, err)

	t.Run("commit restarts transaction", func(t *testing.T) {
		// Execute something to ensure transaction is active
		_, err := Exec("SELECT 1")
		require.NoError(t, err)

		err = Commit()
		require.NoError(t, err)

		// Should be able to execute again with new transaction
		_, err = Exec("SELECT 1")
		assert.NoError(t, err)
	})

	t.Run("rollback restarts transaction", func(t *testing.T) {
		// Execute something to ensure transaction is active
		_, err := Exec("SELECT 1")
		require.NoError(t, err)

		err = Rollback()
		require.NoError(t, err)

		// Should be able to execute again with new transaction
		_, err = Exec("SELECT 1")
		assert.NoError(t, err)
	})

	t.Run("exec creates and inserts into table", func(t *testing.T) {
		// Create a simple test table
		_, err := Exec("CREATE TABLE test (id INTEGER PRIMARY KEY, name TEXT)")
		require.NoError(t, err)

		// Insert data
		result, err := Exec("INSERT INTO test (name) VALUES (?)", "test_name")
		require.NoError(t, err)

		id, err := result.LastInsertId()
		require.NoError(t, err)
		assert.Equal(t, int64(1), id)

		rows, err := result.RowsAffected()
		require.NoError(t, err)
		assert.Equal(t, int64(1), rows)
	})

	t.Run("multiple commit cycles should not crash", func(t *testing.T) {
		// First "login attempt" - do a query
		var result int
		err := QueryRow("SELECT 1").Scan(&result)
		require.NoError(t, err)
		require.Equal(t, 1, result)

		// Simulate end of first login - commit and start new transaction
		err = Commit()
		require.NoError(t, err)

		// Second "login attempt" - this should not crash with nil pointer
		err = QueryRow("SELECT 2").Scan(&result)
		require.NoError(t, err)
		require.Equal(t, 2, result)

		// Third cycle to be thorough
		err = Commit()
		require.NoError(t, err)

		err = QueryRow("SELECT 3").Scan(&result)
		require.NoError(t, err)
		require.Equal(t, 3, result)
	})

	// Test Disconnect
	err = Disconnect()
	require.NoError(t, err)
}

func TestAlreadyConnected(t *testing.T) {
	setupTestDB(t)

	err := Connect(false)
	require.NoError(t, err)

	// Second connect should fail
	err = Connect(false)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "already open")

	err = Disconnect()
	require.NoError(t, err)
}
