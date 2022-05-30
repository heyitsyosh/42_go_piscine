package piscine

func check(r rune) bool {
	return (r >= 'A' &&  r <= 'Z')
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !check(r) {
			return false
		}
	}
	return true
}