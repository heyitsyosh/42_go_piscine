package piscine

func isUpper(r rune) bool {
	return (r >= 'A' && r <= 'Z')
}

func isLower(r rune) bool {
	return (r >= 'a' && r <= 'z')
}

func isAlnum(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' &&  r <= '9')
}

func strLen(s string) (len int) {
	for range s {
		len++
	}
	return
}

func Capitalize(s string) string {
	ret := []rune(s)
	for i, len := 0, strLen(s); i < len; i++ {
		if isUpper(ret[i]) {
			ret[i] += 32
		}
		if isLower(ret[i]) && (i == 0 || !isAlnum(ret[i - 1])) {
			ret[i] -= 32
		}
	}
	return string(ret)
}
