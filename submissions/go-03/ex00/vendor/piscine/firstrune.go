package piscine

func FirstRune(s string) rune {
	if s == "" {
		return 0
	}
	return []rune(s)[0]
}
