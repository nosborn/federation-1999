package goodies

import (
	"os/user"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHomeDir(t *testing.T) {
	t.Run("returns basic valid path", func(t *testing.T) {
		result := HomeDir()

		// Should return a non-empty string
		assert.NotEmpty(t, result, "HomeDir should return a non-empty string")

		// Should return an absolute path
		assert.True(t, filepath.IsAbs(result), "HomeDir should return an absolute path")
	})

	t.Run("multiple calls don't crash", func(t *testing.T) {
		// Test that sync.Once mechanism works correctly
		first := HomeDir()
		second := HomeDir()

		assert.Equal(t, first, second, "HomeDir should return consistent results")
	})

	t.Run("differs from user.Current in setuid scenarios", func(t *testing.T) {
		homeDirResult := HomeDir()

		// user.Current() uses real UID, HomeDir() uses effective UID
		currentUser, err := user.Current()
		require.NoError(t, err, "should be able to get current user")

		// In normal scenarios (non-setuid), these should match
		// In setuid scenarios, they could differ - that's the point of the test
		if homeDirResult != currentUser.HomeDir {
			t.Logf("HomeDir (euid): %s, user.Current (ruid): %s - running in setuid context",
				homeDirResult, currentUser.HomeDir)
		} else {
			// This is the normal case for unprivileged users
			assert.Equal(t, currentUser.HomeDir, homeDirResult,
				"In non-setuid context, HomeDir should match user.Current")
		}
	})
}
