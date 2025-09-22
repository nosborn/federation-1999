package text

// Convert null-terminated C string to Go string.
func CStringToString(cstr []byte) string {
	for i, b := range cstr {
		if b == 0 {
			return string(cstr[:i])
		}
	}
	return string(cstr)
}
