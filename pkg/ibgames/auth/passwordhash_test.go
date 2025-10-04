package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPasswordHash(t *testing.T) {
	t.Run("generates valid bcrypt hash", func(t *testing.T) {
		password := "testpassword123"

		hash, err := PasswordHash(password)
		require.NoError(t, err)
		require.NotEmpty(t, hash)

		// Verify the hash can be used to check the password
		err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
		assert.NoError(t, err)
	})

	t.Run("generates different hashes for same password", func(t *testing.T) {
		password := "samepassword"

		hash1, err := PasswordHash(password)
		require.NoError(t, err)

		hash2, err := PasswordHash(password)
		require.NoError(t, err)

		// Bcrypt should generate different salts
		assert.NotEqual(t, hash1, hash2)

		// But both should validate the password
		require.NoError(t, bcrypt.CompareHashAndPassword([]byte(hash1), []byte(password)))
		require.NoError(t, bcrypt.CompareHashAndPassword([]byte(hash2), []byte(password)))
	})

	t.Run("handles empty password", func(t *testing.T) {
		hash, err := PasswordHash("")
		require.NoError(t, err)
		require.NotEmpty(t, hash)

		// Should be able to verify empty password
		err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(""))
		assert.NoError(t, err)
	})

	t.Run("handles maximum length password", func(t *testing.T) {
		// bcrypt max is PasswordSize bytes
		longPassword := make([]byte, PasswordSize)
		for i := range longPassword {
			longPassword[i] = 'a'
		}

		hash, err := PasswordHash(string(longPassword))
		require.NoError(t, err)

		err = bcrypt.CompareHashAndPassword([]byte(hash), longPassword)
		assert.NoError(t, err)
	})

	t.Run("handles over-length password", func(t *testing.T) {
		// bcrypt should fail on passwords > PasswordSize bytes
		tooLongPassword := make([]byte, PasswordSize+1)
		for i := range tooLongPassword {
			tooLongPassword[i] = 'a'
		}

		_, err := PasswordHash(string(tooLongPassword))
		assert.Error(t, err)
	})
}
