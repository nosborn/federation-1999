package text

// ASCII utility functions for 7-bit ASCII character classification. These
// mirror the C standard library ctype.h functions.
//
// The expressions used are lifted from musl: https://musl.libc.org/

func IsAlnum(b byte) bool {
	return IsAlpha(b) || IsDigit(b)
}

func IsAlpha(b byte) bool {
	return (b|32)-'a' < 26
}

func IsASCII(b byte) bool {
	return !(b&^0x7f != 0)
}

func IsBlank(b byte) bool {
	return b == ' ' || b == '\t'
}

func IsCntrl(b byte) bool {
	return b < 0x20 || b == 0x7f
}

func IsDigit(b byte) bool {
	return b-'0' < 10
}

func IsGraph(b byte) bool {
	return b-0x21 < 0x5e
}

func IsLower(b byte) bool {
	return b-'a' < 26
}

func IsPrint(b byte) bool {
	return b-0x20 < 0x5f
}

func IsPunct(b byte) bool {
	return IsGraph(b) && !IsAlnum(b)
}

func IsSpace(b byte) bool {
	return b == ' ' || b-'\t' < 5
}

func IsUpper(b byte) bool {
	return b-'A' < 26
}

func IsXDigit(b byte) bool {
	return IsDigit(b) || (b|32)-'a' < 6
}

func ToLower(b byte) byte {
	if IsUpper(b) {
		return b | 32
	}
	return b
}

func ToLowerString(s string) string {
	b := []byte(s)
	for i, ch := range b {
		b[i] = ToLower(ch)
	}
	return string(b)
}

func ToUpper(b byte) byte {
	if IsLower(b) {
		return b & 0x5f
	}
	return b
}

func ToUpperString(s string) string {
	b := []byte(s)
	for i, ch := range b {
		b[i] = ToUpper(ch)
	}
	return string(b)
}
