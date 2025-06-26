package auth

import (
	"testing"

	"github.com/nosborn/federation-1999/pkg/ibgames"
	"github.com/nosborn/federation-1999/pkg/ibgames/db"
	"github.com/nosborn/federation-1999/pkg/ibgames/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var globalSetup *testutil.DatabaseSetup

func setupAuthTest(t *testing.T) *testutil.DatabaseSetup {
	if globalSetup != nil {
		return globalSetup
	}

	globalSetup = testutil.SetupTestDatabaseWithSchema(t)

	err := db.Connect(false)
	require.NoError(t, err)

	t.Cleanup(func() {
		db.Exit()
		globalSetup = nil
	})

	return globalSetup
}

func TestLogin(t *testing.T) {
	t.Run("successful login for active account", func(t *testing.T) {
		setup := setupAuthTest(t)
		uid := ibgames.AccountID(666000)
		password := "testpass123"
		hash, err := PasswordHash(password)
		require.NoError(t, err)

		// Create test account
		_, err = setup.TestDB.Exec(`
			INSERT INTO accounts (uid, name, name_key, encrypt, status, complimentary, minutes)
			VALUES (?, ?, ?, ?, ?, ?, ?)
		`, uid, "testuser", "testuser", hash, "A", "N", 100)
		require.NoError(t, err)

		var session Session
		result := Login("testuser", password, "192.0.2.1", &session)

		assert.Equal(t, LoginOK, result)
		assert.Equal(t, uid, session.UID)
		assert.Equal(t, "NEVER", session.SLogin) // First login
		assert.Equal(t, "NEVER", session.ULogin)
	})

	t.Run("login fails for non-existent user", func(t *testing.T) {
		setup := setupAuthTest(t)
		_ = setup // Keep linter happy

		var session Session
		result := Login("nonexistent", "password", "192.0.2.1", &session)

		assert.Equal(t, LoginIncorrect, result)
	})

	t.Run("login fails for wrong password", func(t *testing.T) {
		setup := setupAuthTest(t)
		uid := ibgames.AccountID(666001)
		hash, err := PasswordHash("rightpassword")
		require.NoError(t, err)

		setup.CreateTestAccount(t, uid, "testuser2", "N", 100)

		// Update with proper hash since CreateTestAccount uses dummy hash
		_, err = setup.TestDB.Exec("UPDATE accounts SET encrypt = ? WHERE uid = ?", hash, uid)
		require.NoError(t, err)

		var session Session
		result := Login("testuser2", "wrongpassword", "192.0.2.1", &session)

		assert.Equal(t, LoginIncorrect, result)

		// Commit to make Login's changes visible to direct DB queries
		require.NoError(t, db.Commit())

		// Should increment unsuccessful login count
		var nunsuclog int
		err = setup.TestDB.QueryRow("SELECT nunsuclog FROM accounts WHERE uid = ?", uid).Scan(&nunsuclog)
		require.NoError(t, err)
		assert.Equal(t, 1, nunsuclog)
	})

	t.Run("login fails for suspended account", func(t *testing.T) {
		setup := setupAuthTest(t)
		uid := ibgames.AccountID(666002)
		password := "testpass123"
		hash, err := PasswordHash(password)
		require.NoError(t, err)

		_, err = setup.TestDB.Exec(`
			INSERT INTO accounts (uid, name, name_key, encrypt, status, complimentary, minutes)
			VALUES (?, ?, ?, ?, ?, ?, ?)
		`, uid, "suspended", "suspended", hash, "S", "N", 100)
		require.NoError(t, err)

		var session Session
		result := Login("suspended", password, "192.0.2.1", &session)

		assert.Equal(t, LoginSuspended, result)
	})

	t.Run("login fails for canceled account", func(t *testing.T) {
		setup := setupAuthTest(t)
		uid := ibgames.AccountID(666003)
		password := "testpass123"
		hash, err := PasswordHash(password)
		require.NoError(t, err)

		_, err = setup.TestDB.Exec(`
			INSERT INTO accounts (uid, name, name_key, encrypt, status, complimentary, minutes)
			VALUES (?, ?, ?, ?, ?, ?, ?)
		`, uid, "canceled", "canceled", hash, "X", "N", 100)
		require.NoError(t, err)

		var session Session
		result := Login("canceled", password, "192.0.2.1", &session)

		assert.Equal(t, LoginIncorrect, result)
	})

	t.Run("login fails after too many password attempts", func(t *testing.T) {
		setup := setupAuthTest(t)
		uid := ibgames.AccountID(666004)
		password := "testpass123"
		hash, err := PasswordHash(password)
		require.NoError(t, err)

		_, err = setup.TestDB.Exec(`
			INSERT INTO accounts (uid, name, name_key, encrypt, status, complimentary, minutes, nunsuclog)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		`, uid, "lockedout", "lockedout", hash, "A", "N", 100, maxPasswordTries)
		require.NoError(t, err)

		var session Session
		result := Login("lockedout", password, "192.0.2.1", &session)

		assert.Equal(t, LoginIncorrect, result)
	})

	t.Run("returns session with LoginNoCredit for non-complimentary account with no minutes", func(t *testing.T) {
		setup := setupAuthTest(t)
		uid := ibgames.AccountID(666005)
		password := "testpass123"
		hash, err := PasswordHash(password)
		require.NoError(t, err)

		_, err = setup.TestDB.Exec(`
			INSERT INTO accounts (uid, name, name_key, encrypt, status, complimentary, minutes)
			VALUES (?, ?, ?, ?, ?, ?, ?)
		`, uid, "nocredit", "nocredit", hash, "A", "N", 0)
		require.NoError(t, err)

		var session Session
		result := Login("nocredit", password, "192.0.2.1", &session)

		assert.Equal(t, LoginNoCredit, result)
		assert.Equal(t, uid, session.UID)
	})

	t.Run("successful login for complimentary account with no minutes", func(t *testing.T) {
		setup := setupAuthTest(t)
		uid := ibgames.AccountID(666006)
		password := "testpass123"
		hash, err := PasswordHash(password)
		require.NoError(t, err)

		_, err = setup.TestDB.Exec(`
			INSERT INTO accounts (uid, name, name_key, encrypt, status, complimentary, minutes)
			VALUES (?, ?, ?, ?, ?, ?, ?)
		`, uid, "complimentary", "complimentary", hash, "A", "Y", 0)
		require.NoError(t, err)

		var session Session
		result := Login("complimentary", password, "192.0.2.1", &session)

		assert.Equal(t, LoginOK, result)
		assert.Equal(t, uid, session.UID)
	})
}

func TestLoginParameterValidation(t *testing.T) {
	t.Run("fails with empty name", func(t *testing.T) {
		var session Session
		result := Login("", "password", "192.0.2.1", &session)
		assert.Equal(t, LoginError, result)
	})

	t.Run("fails with empty password", func(t *testing.T) {
		var session Session
		result := Login("username", "", "192.0.2.1", &session)
		assert.Equal(t, LoginError, result)
	})

	t.Run("fails with whitespace-only name", func(t *testing.T) {
		var session Session
		result := Login("   ", "password", "192.0.2.1", &session)
		assert.Equal(t, LoginIncorrect, result)
	})

	t.Run("fails with whitespace-only password", func(t *testing.T) {
		var session Session
		result := Login("username", "   ", "192.0.2.1", &session)
		assert.Equal(t, LoginIncorrect, result)
	})

	t.Run("fails with name too long", func(t *testing.T) {
		longName := make([]byte, NameSize+1)
		for i := range longName {
			longName[i] = 'a'
		}

		var session Session
		result := Login(string(longName), "password", "192.0.2.1", &session)
		assert.Equal(t, LoginError, result)
	})

	t.Run("fails with password too long", func(t *testing.T) {
		longPassword := make([]byte, PasswordSize+1)
		for i := range longPassword {
			longPassword[i] = 'a'
		}

		var session Session
		result := Login("username", string(longPassword), "192.0.2.1", &session)
		assert.Equal(t, LoginError, result)
	})

	t.Run("trims whitespace from name and password", func(t *testing.T) {
		setup := setupAuthTest(t)
		uid := ibgames.AccountID(666007)
		password := "testpass123"
		hash, err := PasswordHash(password)
		require.NoError(t, err)

		_, err = setup.TestDB.Exec(`
			INSERT INTO accounts (uid, name, name_key, encrypt, status, complimentary, minutes)
			VALUES (?, ?, ?, ?, ?, ?, ?)
		`, uid, "trimtest", "trimtest", hash, "A", "N", 100)
		require.NoError(t, err)

		var session Session
		result := Login("  trimtest  ", "  testpass123  ", "192.0.2.1", &session)

		assert.Equal(t, LoginOK, result)
		assert.Equal(t, uid, session.UID)
	})

	t.Run("handles case-insensitive usernames", func(t *testing.T) {
		setup := setupAuthTest(t)
		uid := ibgames.AccountID(666008)
		password := "testpass123"
		hash, err := PasswordHash(password)
		require.NoError(t, err)

		// Store username in lowercase
		_, err = setup.TestDB.Exec(`
			INSERT INTO accounts (uid, name, name_key, encrypt, status, complimentary, minutes)
			VALUES (?, ?, ?, ?, ?, ?, ?)
		`, uid, "casetest", "casetest", hash, "A", "N", 100)
		require.NoError(t, err)

		// Verify the record was inserted
		var storedName string
		err = setup.TestDB.QueryRow("SELECT name FROM accounts WHERE uid = ?", uid).Scan(&storedName)
		require.NoError(t, err)
		assert.Equal(t, "casetest", storedName)

		// Test case-insensitive query directly in database
		var foundUID ibgames.AccountID
		err = setup.TestDB.QueryRow("SELECT uid FROM accounts WHERE name_key = ?", "casetest").Scan(&foundUID)
		require.NoError(t, err, "Direct database query should find case-insensitive match")
		assert.Equal(t, uid, foundUID)

		// Test various case combinations through Login function
		testCases := []string{
			"casetest", // exact match
			"CaseTest", // mixed case
			"CASETEST", // uppercase
			"cAsEtEsT", // random case
		}

		for _, username := range testCases {
			var session Session
			result := Login(username, password, "192.0.2.1", &session)
			assert.Equal(t, LoginOK, result, "should login with username: %q", username)
			assert.Equal(t, uid, session.UID, "should return correct UID for username: %q", username)
		}
	})
}
