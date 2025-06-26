package text

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCStringToString(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "empty string with null terminator",
			input:    []byte{0},
			expected: "",
		},
		{
			name:     "single character with null terminator",
			input:    []byte{'a', 0},
			expected: "a",
		},
		{
			name:     "normal string with null terminator",
			input:    []byte{'h', 'e', 'l', 'l', 'o', 0},
			expected: "hello",
		},
		{
			name:     "string with embedded null (first null terminates)",
			input:    []byte{'h', 'i', 0, 't', 'h', 'e', 'r', 'e', 0},
			expected: "hi",
		},
		{
			name:     "string with spaces and null terminator",
			input:    []byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd', 0},
			expected: "hello world",
		},
		{
			name:     "string with special characters and null terminator",
			input:    []byte{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', 0},
			expected: "!@#$%^&*()",
		},
		{
			name:     "unicode string with null terminator",
			input:    []byte{'H', 'e', 'l', 'l', 'o', ' ', 0xF0, 0x9F, 0x8C, 0x8D, 0},
			expected: "Hello 🌍",
		},
		{
			name:     "string with numeric characters and null terminator",
			input:    []byte{'1', '2', '3', '4', '5', 0},
			expected: "12345",
		},
		{
			name:     "string with tab and newline characters and null terminator",
			input:    []byte{'h', 'e', 'l', 'l', 'o', '\t', 'w', 'o', 'r', 'l', 'd', '\n', 0},
			expected: "hello\tworld\n",
		},
		{
			name:     "null terminator at beginning",
			input:    []byte{0, 'h', 'e', 'l', 'l', 'o'},
			expected: "",
		},
		{
			name:     "longer string with null terminator",
			input:    []byte("The quick brown fox jumps over the lazy dog\000"),
			expected: "The quick brown fox jumps over the lazy dog",
		},
		{
			name:     "string with multiple consecutive nulls",
			input:    []byte{'a', 'b', 'c', 0, 0, 0, 'd', 'e', 'f'},
			expected: "abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CStringToString(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestCStringToStringEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "single null byte",
			input:    []byte{0},
			expected: "",
		},
		{
			name:     "multiple null bytes only",
			input:    []byte{0, 0, 0},
			expected: "",
		},
		{
			name:     "very long string with null terminator",
			input:    append(make([]byte, 1000, 1001), 0),
			expected: string(make([]byte, 1000)),
		},
	}

	// Initialize the very long string test case
	for i := range 1000 {
		tests[2].input[i] = byte('A' + (i % 26))
		tests[2].expected = string(tests[2].input[:1000])
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CStringToString(tt.input)
			require.Equal(t, tt.expected, result)
		})
	}
}
