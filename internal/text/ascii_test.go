package text

import "testing"

// Expected results for each classification function for all 256 byte values
// Note: Indices 128-255 are omitted as they default to false (not ASCII)
var (
	expectedIsAlnum = [256]bool{
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		true, true, true, true, true, true, true, true, true, true, false, false, false, false, false, false,
		false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, false, false, false, false, false,
		false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, false, false, false, false, false,
		// Indices 128-255: all false (non-ASCII values)
	}

	expectedIsAlpha = [256]bool{
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, false, false, false, false, false,
		false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, false, false, false, false, false,
		// Indices 128-255 omitted (all false for non-ASCII values)
	}

	expectedIsASCII = [256]bool{
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		// Indices 128-255 omitted (all false for non-ASCII values)
	}

	expectedIsBlank = [256]bool{
		false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		// Indices 128-255 omitted (all false for non-ASCII values)
	}

	expectedIsCntrl = [256]bool{
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true,
		// Indices 128-255 omitted (all false for non-ASCII values)
	}

	expectedIsDigit = [256]bool{
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		true, true, true, true, true, true, true, true, true, true, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		// Indices 128-255 omitted (all false for non-ASCII values)
	}

	expectedIsGraph = [256]bool{
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false,
		// Indices 128-255 omitted (all false for non-ASCII values)
	}

	expectedIsLower = [256]bool{
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, false, false, false, false, false,
		// Indices 128-255 omitted (all false for non-ASCII values)
	}

	expectedIsPrint = [256]bool{
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, false,
		// Indices 128-255 omitted (all false for non-ASCII values)
	}

	expectedIsSpace = [256]bool{
		false, false, false, false, false, false, false, false, false, true, true, true, true, true, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		// Indices 128-255 omitted (all false for non-ASCII values)
	}

	expectedIsUpper = [256]bool{
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true, true, true, true, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		// Indices 128-255 omitted (all false for non-ASCII values)
	}

	expectedIsXDigit = [256]bool{
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		true, true, true, true, true, true, true, true, true, true, false, false, false, false, false, false,
		false, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		false, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		// Indices 128-255 omitted (all false for non-ASCII values)
	}
)

func TestIsAlnum(t *testing.T) {
	for i := range 256 {
		b := byte(i)
		expected := expectedIsAlnum[i]
		result := IsAlnum(b)
		if result != expected {
			t.Errorf("IsAlnum(%d) = %v, expected %v", i, result, expected)
		}
	}
}

func TestIsAlpha(t *testing.T) {
	for i := range 256 {
		b := byte(i)
		expected := expectedIsAlpha[i]
		result := IsAlpha(b)
		if result != expected {
			t.Errorf("IsAlpha(%d) = %v, expected %v", i, result, expected)
		}
	}
}

func TestIsASCII(t *testing.T) {
	for i := range 256 {
		b := byte(i)
		expected := expectedIsASCII[i]
		result := IsASCII(b)
		if result != expected {
			t.Errorf("IsASCII(%d) = %v, expected %v", i, result, expected)
		}
	}
}

func TestIsBlank(t *testing.T) {
	for i := range 256 {
		b := byte(i)
		expected := expectedIsBlank[i]
		result := IsBlank(b)
		if result != expected {
			t.Errorf("IsBlank(%d) = %v, expected %v", i, result, expected)
		}
	}
}

func TestIsCntrl(t *testing.T) {
	for i := range 256 {
		b := byte(i)
		expected := expectedIsCntrl[i]
		result := IsCntrl(b)
		if result != expected {
			t.Errorf("IsCntrl(%d) = %v, expected %v", i, result, expected)
		}
	}
}

func TestIsDigit(t *testing.T) {
	for i := range 256 {
		b := byte(i)
		expected := expectedIsDigit[i]
		result := IsDigit(b)
		if result != expected {
			t.Errorf("IsDigit(%d) = %v, expected %v", i, result, expected)
		}
	}
}

func TestIsGraph(t *testing.T) {
	for i := range 256 {
		b := byte(i)
		expected := expectedIsGraph[i]
		result := IsGraph(b)
		if result != expected {
			t.Errorf("IsGraph(%d) = %v, expected %v", i, result, expected)
		}
	}
}

func TestIsLower(t *testing.T) {
	for i := range 256 {
		b := byte(i)
		expected := expectedIsLower[i]
		result := IsLower(b)
		if result != expected {
			t.Errorf("IsLower(%d) = %v, expected %v", i, result, expected)
		}
	}
}

func TestIsPrint(t *testing.T) {
	for i := range 256 {
		b := byte(i)
		expected := expectedIsPrint[i]
		result := IsPrint(b)
		if result != expected {
			t.Errorf("IsPrint(%d) = %v, expected %v", i, result, expected)
		}
	}
}

func TestIsPunct(t *testing.T) {
	for i := range 256 {
		b := byte(i)
		expected := expectedIsGraph[i] && !expectedIsAlnum[i]
		result := IsPunct(b)
		if result != expected {
			t.Errorf("IsPunct(%d) = %v, expected %v", i, result, expected)
		}
	}
}

func TestIsSpace(t *testing.T) {
	for i := range 256 {
		b := byte(i)
		expected := expectedIsSpace[i]
		result := IsSpace(b)
		if result != expected {
			t.Errorf("IsSpace(%d) = %v, expected %v", i, result, expected)
		}
	}
}

func TestIsUpper(t *testing.T) {
	for i := range 256 {
		b := byte(i)
		expected := expectedIsUpper[i]
		result := IsUpper(b)
		if result != expected {
			t.Errorf("IsUpper(%d) = %v, expected %v", i, result, expected)
		}
	}
}

func TestIsXDigit(t *testing.T) {
	for i := range 256 {
		b := byte(i)
		expected := expectedIsXDigit[i]
		result := IsXDigit(b)
		if result != expected {
			t.Errorf("IsXDigit(%d) = %v, expected %v", i, result, expected)
		}
	}
}

func TestToLower(t *testing.T) {
	for i := range 256 {
		b := byte(i)
		result := ToLower(b)

		// Expected behavior: uppercase letters get converted to lowercase,
		// everything else stays the same
		var expected byte
		if b >= 'A' && b <= 'Z' {
			expected = b | 32 // Convert to lowercase using bitwise OR
		} else {
			expected = b
		}

		if result != expected {
			t.Errorf("ToLower(%d) = %d, expected %d", i, result, expected)
		}
	}
}

func TestToUpper(t *testing.T) {
	for i := range 256 {
		b := byte(i)
		result := ToUpper(b)

		// Expected behavior: lowercase letters get converted to uppercase,
		// everything else stays the same
		var expected byte
		if b >= 'a' && b <= 'z' {
			expected = b & 0x5f // Convert to uppercase using bitwise AND
		} else {
			expected = b
		}

		if result != expected {
			t.Errorf("ToUpper(%d) = %d, expected %d", i, result, expected)
		}
	}
}

// Test that IsAlnum is consistent with IsAlpha || IsDigit
func TestIsAlnumConsistency(t *testing.T) {
	for i := range 256 {
		b := byte(i)
		expected := IsAlpha(b) || IsDigit(b)
		actual := IsAlnum(b)
		if actual != expected {
			t.Errorf("IsAlnum(%d) = %v, but IsAlpha||IsDigit = %v", i, actual, expected)
		}
	}
}

// Test that IsPunct is consistent with IsGraph && !IsAlnum
func TestIsPunctConsistency(t *testing.T) {
	for i := range 256 {
		b := byte(i)
		expected := IsGraph(b) && !IsAlnum(b)
		actual := IsPunct(b)
		if actual != expected {
			t.Errorf("IsPunct(%d) = %v, but IsGraph&&!IsAlnum = %v", i, actual, expected)
		}
	}
}

// Test case conversion round trips
func TestCaseConversionRoundTrip(t *testing.T) {
	for i := range 256 {
		b := byte(i)

		// ToLower(ToUpper(x)) should equal ToLower(x)
		if ToLower(ToUpper(b)) != ToLower(b) {
			t.Errorf("ToLower(ToUpper(%d)) != ToLower(%d)", i, i)
		}

		// ToUpper(ToLower(x)) should equal ToUpper(x)
		if ToUpper(ToLower(b)) != ToUpper(b) {
			t.Errorf("ToUpper(ToLower(%d)) != ToUpper(%d)", i, i)
		}
	}
}

func TestToLowerString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"HELLO", "hello"},
		{"hello", "hello"},
		{"Hello World!", "hello world!"},
		{"ABC123XYZ", "abc123xyz"},
		{"MiXeD CaSe StRiNg", "mixed case string"},
		{"UPPER with 123 NUMBERS", "upper with 123 numbers"},
		{"Special@#$%Characters", "special@#$%characters"},
		{"   SPACES   EVERYWHERE   ", "   spaces   everywhere   "},
		{"Tab\tAnd\nNewline\rChars", "tab\tand\nnewline\rchars"},
		{"!@#$%^&*()_+-={}[]|\\:;\"'<>?,./", "!@#$%^&*()_+-={}[]|\\:;\"'<>?,./"},
		{"AlPhAbEt SOUP 123", "alphabet soup 123"},
		{"PROGRAMMING in GO is FUN!", "programming in go is fun!"},
		{"CamelCaseVariableName", "camelcasevariablename"},
		{"ALL_CAPS_WITH_UNDERSCORES", "all_caps_with_underscores"},
	}

	for _, test := range tests {
		result := ToLowerString(test.input)
		if result != test.expected {
			t.Errorf("ToLowerString(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestToUpperString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"hello", "HELLO"},
		{"HELLO", "HELLO"},
		{"Hello World!", "HELLO WORLD!"},
		{"abc123xyz", "ABC123XYZ"},
		{"MiXeD CaSe StRiNg", "MIXED CASE STRING"},
		{"lower with 123 numbers", "LOWER WITH 123 NUMBERS"},
		{"special@#$%characters", "SPECIAL@#$%CHARACTERS"},
		{"   spaces   everywhere   ", "   SPACES   EVERYWHERE   "},
		{"tab\tand\nnewline\rchars", "TAB\tAND\nNEWLINE\rCHARS"},
		{"!@#$%^&*()_+-={}[]|\\:;\"'<>?,./", "!@#$%^&*()_+-={}[]|\\:;\"'<>?,./"},
		{"alphabet soup 123", "ALPHABET SOUP 123"},
		{"programming in go is fun!", "PROGRAMMING IN GO IS FUN!"},
		{"camelCaseVariableName", "CAMELCASEVARIABLENAME"},
		{"all_caps_with_underscores", "ALL_CAPS_WITH_UNDERSCORES"},
	}

	for _, test := range tests {
		result := ToUpperString(test.input)
		if result != test.expected {
			t.Errorf("ToUpperString(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestStringCaseConversionRoundTrip(t *testing.T) {
	testStrings := []string{
		"",
		"Hello World!",
		"MiXeD CaSe StRiNg",
		"ABC123XYZ",
		"programming in go is fun!",
		"Special@#$%Characters",
		"   spaces   everywhere   ",
		"Tab\tAnd\nNewline\rChars",
		"!@#$%^&*()_+-={}[]|\\:;\"'<>?,./",
		"AlPhAbEt SOUP 123",
		"CamelCaseVariableName",
		"ALL_CAPS_WITH_UNDERSCORES",
		"lower_case_with_underscores",
		"UPPER CASE WITH SPACES",
		"Numbers123And456Symbols!@#",
	}

	for _, s := range testStrings {
		// Test ToLower -> ToUpper -> ToLower round trip
		lowerFirst := ToLowerString(s)
		upperThenLower := ToLowerString(ToUpperString(lowerFirst))
		if lowerFirst != upperThenLower {
			t.Errorf("ToLower->ToUpper->ToLower round trip failed for %q: %q != %q", s, lowerFirst, upperThenLower)
		}

		// Test ToUpper -> ToLower -> ToUpper round trip
		upperFirst := ToUpperString(s)
		lowerThenUpper := ToUpperString(ToLowerString(upperFirst))
		if upperFirst != lowerThenUpper {
			t.Errorf("ToUpper->ToLower->ToUpper round trip failed for %q: %q != %q", s, upperFirst, lowerThenUpper)
		}

		// Test that ToLower(ToUpper(s)) == ToLower(s)
		if ToLowerString(ToUpperString(s)) != ToLowerString(s) {
			t.Errorf("ToLower(ToUpper(%q)) != ToLower(%q)", s, s)
		}

		// Test that ToUpper(ToLower(s)) == ToUpper(s)
		if ToUpperString(ToLowerString(s)) != ToUpperString(s) {
			t.Errorf("ToUpper(ToLower(%q)) != ToUpper(%q)", s, s)
		}
	}
}
