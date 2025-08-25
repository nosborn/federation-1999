package text

func IsShouting(s string) bool {
	if len(s) <= 3 {
		return false
	}

	shouting := false
	for _, c := range []byte(s) {
		if IsLower(c) {
			return false
		}
		if IsUpper(c) {
			shouting = true
		}
	}
	return shouting
}
