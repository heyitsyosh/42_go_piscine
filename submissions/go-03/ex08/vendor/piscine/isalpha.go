package piscine

func check(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' &&  r <= '9')
}

func IsAlpha(s string) bool {
	for _, r := range s {
		if !check(r) {
			return false
		}
	}
	return true
}
