package piscine

func strLen(s string) (len int) {
	for range s {
		len++
	}
	return
}

func NRune(s string, n int) rune {
	sLen := strLen(s)

	for i := 0; i < sLen; i++ {
		if i == n - 1 {
			return []rune(s)[i]
		}
	}
	return 0
}
