package piscine

func strLen(s string) (len int) {
	for range s {
		len++
	}
	return
}

func ToLower(s string) string {
	ret := []rune(s)
	for i, len := 0, strLen(s); i < len; i++ {
		if ret[i] >= 'A' && ret[i] <= 'Z' {
			ret[i] += 32
		}
	}
	return string(ret)
}
