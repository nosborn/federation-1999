package parser

import (
	"testing"
)

func TestParseIdNum(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantNum  int32
		wantBool bool
	}{
		{
			name:     "valid id number",
			input:    "#123",
			wantNum:  123,
			wantBool: true,
		},
		{
			name:     "valid zero id",
			input:    "#0",
			wantNum:  0,
			wantBool: true,
		},
		{
			name:     "valid large number",
			input:    "#999999",
			wantNum:  999999,
			wantBool: true,
		},
		{
			name:     "missing hash",
			input:    "123",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "hash only",
			input:    "#",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "hash with non-numeric",
			input:    "#abc",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "negative number with hash",
			input:    "#-123",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "hash with spaces",
			input:    "# 123",
			wantNum:  0,
			wantBool: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNum, gotBool := parseIDNum(tt.input)
			if gotNum != tt.wantNum {
				t.Errorf("parseIDNum() gotNum = %v, want %v", gotNum, tt.wantNum)
			}
			if gotBool != tt.wantBool {
				t.Errorf("parseIDNum() gotBool = %v, want %v", gotBool, tt.wantBool)
			}
		})
	}
}

func TestParseIntSuccess(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantNum int32
	}{
		{
			name:    "valid positive integer",
			input:   "123",
			wantNum: 123,
		},
		{
			name:    "valid zero",
			input:   "0",
			wantNum: 0,
		},
		{
			name:    "valid negative integer",
			input:   "-456",
			wantNum: -456,
		},
		{
			name:    "valid max int32",
			input:   "2147483647",
			wantNum: 2147483647,
		},
		{
			name:    "valid min int32",
			input:   "-2147483648",
			wantNum: -2147483648,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNum, gotBool := parseInt(tt.input)
			if !gotBool {
				t.Errorf("parseInt() should succeed but gotBool = false")
			}
			if gotNum != tt.wantNum {
				t.Errorf("parseInt() gotNum = %v, want %v", gotNum, tt.wantNum)
			}
		})
	}
}

func TestParseIntFailure(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "exceeds max int32",
			input: "2147483648",
		},
		{
			name:  "below min int32",
			input: "-2147483649",
		},
		{
			name:  "way above max int32",
			input: "9223372036854775807",
		},
		{
			name:  "way below min int32",
			input: "-9223372036854775808",
		},
		{
			name:  "non-numeric string",
			input: "abc",
		},
		{
			name:  "mixed alphanumeric",
			input: "12abc",
		},
		{
			name:  "number with spaces",
			input: " 123 ",
		},
		{
			name:  "decimal number",
			input: "123.45",
		},
		{
			name:  "number with comma",
			input: "1,234",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, gotBool := parseInt(tt.input)
			if gotBool {
				t.Errorf("parseInt() should fail but gotBool = true")
			}
		})
	}
}

func TestParseFmtInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantNum  int32
		wantBool bool
	}{
		{
			name:     "number without commas",
			input:    "123",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "number with single comma",
			input:    "1,234",
			wantNum:  1234,
			wantBool: true,
		},
		{
			name:     "number with multiple commas",
			input:    "1,234,567",
			wantNum:  1234567,
			wantBool: true,
		},
		{
			name:     "zero with comma",
			input:    "0,000",
			wantNum:  0,
			wantBool: true,
		},
		{
			name:     "negative number with commas",
			input:    "-1,234,567",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "number with irregular comma placement",
			input:    "12,34,567",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "number starting with comma",
			input:    ",123",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "number ending with comma",
			input:    "123,",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "only commas",
			input:    ",,,",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "non-numeric with commas",
			input:    "1,2a3",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "large formatted number",
			input:    "999,999,999",
			wantNum:  999999999,
			wantBool: true,
		},
		{
			name:     "single digit with comma group",
			input:    "5,000",
			wantNum:  5000,
			wantBool: true,
		},
		{
			name:     "two digits with comma group",
			input:    "50,000",
			wantNum:  50000,
			wantBool: true,
		},
		{
			name:     "four digits no comma (not formatted)",
			input:    "5000",
			wantNum:  0,
			wantBool: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNum, gotBool := parseFmtInt(tt.input)
			if gotNum != tt.wantNum {
				t.Errorf("parseFmtInt() gotNum = %v, want %v", gotNum, tt.wantNum)
			}
			if gotBool != tt.wantBool {
				t.Errorf("parseFmtInt() gotBool = %v, want %v", gotBool, tt.wantBool)
			}
		})
	}
}

func TestParseSignedInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantNum  int32
		wantBool bool
	}{
		{
			name:     "positive number with plus sign",
			input:    "+123",
			wantNum:  123,
			wantBool: true,
		},
		{
			name:     "negative number with minus sign",
			input:    "-456",
			wantNum:  -456,
			wantBool: true,
		},
		{
			name:     "positive zero",
			input:    "+0",
			wantNum:  0,
			wantBool: true,
		},
		{
			name:     "negative zero",
			input:    "-0",
			wantNum:  0,
			wantBool: true,
		},
		{
			name:     "large positive number",
			input:    "+999999",
			wantNum:  999999,
			wantBool: true,
		},
		{
			name:     "large negative number",
			input:    "-999999",
			wantNum:  -999999,
			wantBool: true,
		},
		{
			name:     "number without sign",
			input:    "123",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "plus sign only",
			input:    "+",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "minus sign only",
			input:    "-",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "non-numeric with plus",
			input:    "+abc",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "non-numeric with minus",
			input:    "-xyz",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "mixed with plus",
			input:    "+12abc",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "double sign",
			input:    "++123",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "space after sign",
			input:    "+ 123",
			wantNum:  0,
			wantBool: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNum, gotBool := parseSignedInt(tt.input)
			if gotNum != tt.wantNum {
				t.Errorf("parseSignedInt() gotNum = %v, want %v", gotNum, tt.wantNum)
			}
			if gotBool != tt.wantBool {
				t.Errorf("parseSignedInt() gotBool = %v, want %v", gotBool, tt.wantBool)
			}
		})
	}
}

func TestParsePercent(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantNum  int32
		wantBool bool
	}{
		{
			name:     "valid percent",
			input:    "50%",
			wantNum:  50,
			wantBool: true,
		},
		{
			name:     "zero percent",
			input:    "0%",
			wantNum:  0,
			wantBool: true,
		},
		{
			name:     "one hundred percent",
			input:    "100%",
			wantNum:  100,
			wantBool: true,
		},
		{
			name:     "large percent",
			input:    "999%",
			wantNum:  999,
			wantBool: true,
		},
		{
			name:     "negative percent",
			input:    "-25%",
			wantNum:  -25,
			wantBool: true,
		},
		{
			name:     "decimal without percent",
			input:    "50.5",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "number without percent",
			input:    "123",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "percent sign only",
			input:    "%",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "non-numeric with percent",
			input:    "abc%",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "mixed with percent",
			input:    "12abc%",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "space before percent",
			input:    "50 %",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "decimal with percent",
			input:    "50.5%",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "multiple percent signs",
			input:    "50%%",
			wantNum:  0,
			wantBool: false,
		},
		{
			name:     "percent at beginning",
			input:    "%50",
			wantNum:  0,
			wantBool: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNum, gotBool := parsePercent(tt.input)
			if gotNum != tt.wantNum {
				t.Errorf("parsePercent() gotNum = %v, want %v", gotNum, tt.wantNum)
			}
			if gotBool != tt.wantBool {
				t.Errorf("parsePercent() gotBool = %v, want %v", gotBool, tt.wantBool)
			}
		})
	}
}

func TestIsReservedWord(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		// Test some common reserved words
		{
			name:     "look command",
			input:    "look",
			expected: true,
		},
		{
			name:     "go command",
			input:    "go",
			expected: true,
		},
		{
			name:     "buy command",
			input:    "buy",
			expected: true,
		},
		{
			name:     "sell command",
			input:    "sell",
			expected: true,
		},
		{
			name:     "quit command",
			input:    "quit",
			expected: true,
		},
		// Test case insensitivity
		{
			name:     "uppercase LOOK",
			input:    "LOOK",
			expected: true,
		},
		{
			name:     "mixed case LoOk",
			input:    "LoOk",
			expected: true,
		},
		// Test abbreviations
		{
			name:     "l abbreviation",
			input:    "l",
			expected: true,
		},
		{
			name:     "n abbreviation",
			input:    "n",
			expected: true,
		},
		{
			name:     "s abbreviation",
			input:    "s",
			expected: true,
		},
		// Test commodity names
		{
			name:     "gold commodity",
			input:    "gold",
			expected: true,
		},
		{
			name:     "alloys commodity",
			input:    "alloys",
			expected: true,
		},
		// Test stats
		{
			name:     "strength stat",
			input:    "strength",
			expected: true,
		},
		{
			name:     "str abbreviation",
			input:    "str",
			expected: true,
		},
		{
			name:     "intelligence stat",
			input:    "intelligence",
			expected: true,
		},
		{
			name:     "int abbreviation",
			input:    "int",
			expected: true,
		},
		// Test non-reserved words
		{
			name:     "non-reserved word",
			input:    "foobar",
			expected: false,
		},
		{
			name:     "number as string",
			input:    "123",
			expected: false,
		},
		{
			name:     "punctuation",
			input:    "!@#",
			expected: false,
		},
		{
			name:     "common word not in game",
			input:    "hello",
			expected: false,
		},
		{
			name:     "partial match",
			input:    "looking",
			expected: false,
		},
		// Test some edge cases
		{
			name:     "technician special case",
			input:    "technician",
			expected: true,
		},
		{
			name:     "reserved word with spaces",
			input:    "look around",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsReservedWord(tt.input)
			if result != tt.expected {
				t.Errorf("IsReservedWord(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
