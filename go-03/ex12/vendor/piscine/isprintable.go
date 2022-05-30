package piscine

func check(r rune) bool {
	return (r >= 32 &&  r <= 126)
}

func IsPrintable(s string) bool {
	for _, r := range s {
		if !check(r) {
			return false
		}
	}
	return true
}