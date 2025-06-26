package text

import "testing"

func TestListOfObjects(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "single item",
			input:    []string{"apple"},
			expected: "apple.",
		},
		{
			name:     "two items",
			input:    []string{"apple", "banana"},
			expected: "apple and banana.",
		},
		{
			name:     "three items with Oxford comma",
			input:    []string{"apple", "banana", "cherry"},
			expected: "apple, banana, and cherry.",
		},
		{
			name:     "four items with Oxford comma",
			input:    []string{"apple", "banana", "cherry", "date"},
			expected: "apple, banana, cherry, and date.",
		},
		{
			name:     "single item with spaces",
			input:    []string{"red apple"},
			expected: "red apple.",
		},
		{
			name:     "two items with spaces",
			input:    []string{"red apple", "green banana"},
			expected: "red apple and green banana.",
		},
		{
			name:     "three items with spaces",
			input:    []string{"red apple", "green banana", "yellow cherry"},
			expected: "red apple, green banana, and yellow cherry.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ListOfObjects(tt.input)
			if result != tt.expected {
				t.Errorf("ListOfObjects(%v) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
