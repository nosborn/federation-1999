package text

func YesNoReply(input string) (bool, bool) { // returns reply, ok
	if input == "" {
		return false, false
	}
	if ToLowerString(input) == "yes"[:len(input)] {
		return true, true
	}
	if ToLowerString(input) == "no"[:len(input)] {
		return false, true
	}
	return false, false
}
