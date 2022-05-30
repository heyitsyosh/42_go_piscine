package piscine

func check(r rune) bool {
	return (r >= '0' &&  r <= '9')
}

func IsNumeric(s string) bool {
	for _, r := range s {
		if !check(r) {
			return false
		}
	}
	return true
}