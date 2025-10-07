package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueName(t *testing.T) {
	t.Run("converts uppercase to lowercase", func(t *testing.T) {
		assert.Equal(t, "hello", UniqueName("HELLO"))
		assert.Equal(t, "world", UniqueName("WORLD"))
		assert.Equal(t, "test", UniqueName("TEST"))
	})

	t.Run("preserves lowercase", func(t *testing.T) {
		assert.Equal(t, "hello", UniqueName("hello"))
		assert.Equal(t, "world", UniqueName("world"))
	})

	t.Run("handles mixed case", func(t *testing.T) {
		assert.Equal(t, "helloworld", UniqueName("HelloWorld"))
		assert.Equal(t, "testuser", UniqueName("TestUser"))
		assert.Equal(t, "mixedcase", UniqueName("MiXeDcAsE"))
	})

	t.Run("preserves digits", func(t *testing.T) {
		assert.Equal(t, "user123", UniqueName("User123"))
		assert.Equal(t, "test456", UniqueName("TEST456"))
	})

	t.Run("preserves symbols", func(t *testing.T) {
		assert.Equal(t, "user@domain.com", UniqueName("User@Domain.Com"))
		assert.Equal(t, "test_name", UniqueName("Test_Name"))
		assert.Equal(t, "user-123", UniqueName("USER-123"))
	})

	t.Run("removes spaces and control characters", func(t *testing.T) {
		assert.Equal(t, "hello", UniqueName("hello "))
		assert.Equal(t, "test", UniqueName(" test"))
		assert.Equal(t, "helloworld", UniqueName("hello world"))
		assert.Equal(t, "test", UniqueName("test\t"))
		assert.Equal(t, "test", UniqueName("test\n"))
	})

	t.Run("handles empty string", func(t *testing.T) {
		assert.Empty(t, UniqueName(""))
	})

	t.Run("handles only spaces", func(t *testing.T) {
		assert.Empty(t, UniqueName("   "))
		assert.Empty(t, UniqueName("\t\n"))
	})

	t.Run("preserves punctuation within graph range", func(t *testing.T) {
		assert.Equal(t, "!@#$%^&*()", UniqueName("!@#$%^&*()"))
		assert.Equal(t, "user.name", UniqueName("User.Name"))
		assert.Equal(t, "file.txt", UniqueName("File.Txt"))
	})
}
