package piscine

func check(r rune) bool {
	return (r >= 'a' &&  r <= 'z')
}

func IsLower(s string) bool {
	for _, r := range s {
		if !check(r) {
			return false
		}
	}
	return true
}