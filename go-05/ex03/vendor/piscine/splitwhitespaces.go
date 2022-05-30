package piscine

func isWhiteSpace(r rune) bool {
	return r == '\t' || r == '\n' || r == '\v' || r == ' '
}

func SplitWhiteSpaces(s string) []string {
	r_s, lastWhiteSpace := []rune(s), -1
	var ret []string
	len := 0
	for i, r := range r_s {
		if isWhiteSpace(r) {
			if (i == 0 || isWhiteSpace(r_s[i - 1])) {
				lastWhiteSpace = i
			} else {
				ret = append(ret, string(r_s[lastWhiteSpace + 1:i]))
				lastWhiteSpace = i
			}
		}
		len++
	}
	if lastWhiteSpace != len - 1 && len != 0 {
		ret = append(ret, string(r_s[lastWhiteSpace + 1:]))
	}
	return ret
}
