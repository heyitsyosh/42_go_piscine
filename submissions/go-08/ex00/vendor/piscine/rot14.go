package piscine

func convert(r rune) rune {
	if (r >= 'm' && r <= 'z') || (r >= 'M' && r <= 'Z') {
		return r - 12
	} else if (r >= 'a' && r <= 'l') || (r >= 'A' && r <= 'L') {
		return r + 14
	} else {
		return r
	}
}

func Rot14(s string) string {
	ret := ""
	for _, r := range s {
		ret += string(convert(r))
	}
	return ret
}
