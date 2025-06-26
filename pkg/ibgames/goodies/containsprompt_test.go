package goodies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsPrompt(t *testing.T) {
	t.Run("detects login prompts", func(t *testing.T) {
		testCases := []string{
			"login:",
			"Login:",
			"LOGIN:",
			"Please login:",
			"Username/login:",
			"ogin:", // matches partial "ogin"
		}

		for _, input := range testCases {
			assert.True(t, ContainsPrompt(input), "should detect prompt in: %q", input)
		}
	})

	t.Run("detects password prompts", func(t *testing.T) {
		testCases := []string{
			"password:",
			"Password:",
			"PASSWORD:",
			"Enter password:",
			"User password:",
			"word:", // matches partial "word"
		}

		for _, input := range testCases {
			assert.True(t, ContainsPrompt(input), "should detect prompt in: %q", input)
		}
	})

	t.Run("detects prompts in context", func(t *testing.T) {
		testCases := []string{
			"System login: user123",
			"Enter your password: ",
			"Failed login: try again",
			"Invalid password: access denied",
			"login: _",
			"password: ********",
		}

		for _, input := range testCases {
			assert.True(t, ContainsPrompt(input), "should detect prompt in: %q", input)
		}
	})

	t.Run("ignores non-prompts", func(t *testing.T) {
		testCases := []string{
			"",                   // empty string
			"hello world",        // no prompt
			"login",              // missing colon
			"password",           // missing colon
			"passwords:",         // "words:" not "word:"
			"logging in",         // no colon
			"password reset",     // no colon
			"user@domain.com",    // email address, no colon
			"http://example.com", // contains colon but wrong pattern
		}

		for _, input := range testCases {
			assert.False(t, ContainsPrompt(input), "should not detect prompt in: %q", input)
		}
	})

	t.Run("detects potential attack strings", func(t *testing.T) {
		// These should match as they could be used to trigger terminal auto-login
		testCases := []string{
			"blogin:",     // contains "ogin:"
			"xxogin:yy",   // "ogin:" embedded
			"myword:test", // "word:" embedded
		}

		for _, input := range testCases {
			assert.True(t, ContainsPrompt(input), "should detect potential attack in: %q", input)
		}
	})

	t.Run("handles edge cases", func(t *testing.T) {
		// Test mixed case
		assert.True(t, ContainsPrompt("LoGiN:"))
		assert.True(t, ContainsPrompt("PaSsWoRd:"))

		// Test multiple prompts in same string
		assert.True(t, ContainsPrompt("login: failed, try password: again"))

		// Test prompts at different positions
		assert.True(t, ContainsPrompt("login: at start"))
		assert.True(t, ContainsPrompt("middle login: here"))
		assert.True(t, ContainsPrompt("at end login:"))
	})

	t.Run("regex compiles once", func(t *testing.T) {
		// Multiple calls should work without issues (tests sync.Once behavior)
		for range 10 {
			assert.True(t, ContainsPrompt("login:"))
			assert.False(t, ContainsPrompt("noprompt"))
		}
	})
}
