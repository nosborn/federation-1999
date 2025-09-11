package rules

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/nosborn/federation-1999/pkg/ibgames"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRulesLockFile(t *testing.T) {
	t.Run("valid uid returns expected path format", func(t *testing.T) {
		tempDir := t.TempDir()
		err := os.MkdirAll(filepath.Join(tempDir, "lock"), 0o755)
		require.NoError(t, err)

		originalHomeDir := homeDir
		homeDir = func() string { return tempDir }
		defer func() { homeDir = originalHomeDir }()

		uid := ibgames.AccountID(666000)
		result := RulesLockFile(uid)

		expected := fmt.Sprintf("%s/lock/666000", tempDir)
		assert.Equal(t, expected, result)
	})

	t.Run("panics on uid below minimum", func(t *testing.T) {
		uid := ibgames.AccountID(ibgames.MinAccountID - 1)
		assert.Panics(t, func() { RulesLockFile(uid) })
	})

	t.Run("panics on uid above maximum", func(t *testing.T) {
		uid := ibgames.AccountID(ibgames.MaxAccountID + 1)
		assert.Panics(t, func() { RulesLockFile(uid) })
	})
}

func TestIsLockedOut(t *testing.T) {
	t.Run("returns false when lock file does not exist", func(t *testing.T) {
		tempDir := t.TempDir()
		err := os.MkdirAll(filepath.Join(tempDir, "lock"), 0o755)
		require.NoError(t, err)

		originalHomeDir := homeDir
		homeDir = func() string { return tempDir }
		defer func() { homeDir = originalHomeDir }()

		uid := ibgames.AccountID(666000)
		result := IsLockedOut(uid)
		assert.False(t, result)
	})

	t.Run("returns true when lock file exists", func(t *testing.T) {
		tempDir := t.TempDir()
		err := os.MkdirAll(filepath.Join(tempDir, "lock"), 0o755)
		require.NoError(t, err)

		originalHomeDir := homeDir
		homeDir = func() string { return tempDir }
		defer func() { homeDir = originalHomeDir }()

		uid := ibgames.AccountID(666000)
		lockFile := RulesLockFile(uid)

		// Create the lock file
		file, err := os.Create(lockFile)
		require.NoError(t, err)
		file.Close()

		// Test that it's detected as locked out
		result := IsLockedOut(uid)
		assert.True(t, result)
	})

	t.Run("panics on uid below minimum", func(t *testing.T) {
		uid := ibgames.AccountID(ibgames.MinAccountID - 1)
		assert.Panics(t, func() { IsLockedOut(uid) })
	})

	t.Run("panics on uid above maximum", func(t *testing.T) {
		uid := ibgames.AccountID(ibgames.MaxAccountID + 1)
		assert.Panics(t, func() { IsLockedOut(uid) })
	})
}
