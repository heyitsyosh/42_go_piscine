package piscine

func IsNumeric(s string) bool {
	for _, r := range s {
		if !('0' <= r && r <= '9') {
			return false
		}
	}
	return true
}
