package auth

func UniqueName(name string) string {
	var unique []byte
	for _, ch := range []byte(name) {
		if isUpper(ch) {
			unique = append(unique, toLower(ch))
		} else if isGraph(ch) {
			unique = append(unique, ch)
		}
	}
	return string(unique)
}

func isUpper(b byte) bool {
	return b-'A' < 26
}

func isGraph(b byte) bool {
	return b-0x21 < 0x5e
}

func toLower(b byte) byte {
	if isUpper(b) {
		return b | 32
	}
	return b
}
