package piscine

func strLen(s string) (len int) {
	for range s {
		len++
	}
	return
}

func ToUpper(s string) string {
	ret := []rune(s)
	for i, len := 0, strLen(s); i < len; i++ {
		if ret[i] >= 'a' && ret[i] <= 'z' {
			ret[i] -= 32
		}
	}
	return string(ret)
}
