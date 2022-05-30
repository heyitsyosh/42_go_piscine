package piscine

func StringHasRune(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}

func StringTrim(s string, charSet string) string {
	charMap := map[rune]bool{}
	for _, c := range charSet {
		charMap[c] = true
	}
	f := 0
	t := len(s) - 1
	if !(f <= t) {
		return ""
	}
	rs := []rune(s)
	for ; f <= t; f++ {
		_, e := charMap[rs[f]]
		if !e {
			break
		}
	}
	for ; f <= t; t-- {
		_, e := charMap[rs[t]]
		if !e {
			break
		}
	}
	return string(rs[f : t+1])
}
