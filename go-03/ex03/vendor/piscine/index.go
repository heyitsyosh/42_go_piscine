package piscine

func strLen(s string) (len int) {
	for range s {
		len++
	}
	return
}

func Index(s string, toFind string) int {
	r_s := []rune(s)
	r_toFind := []rune(toFind)
	if strLen(toFind) == 0 {
		return 0
	}
	if strLen(s) == 0 {
		return -1
	}
	for i, len := 0, strLen(s); i < len; i++ {
		check:
		if r_s[i] == r_toFind[0] {
			for j, r := range r_toFind {
				if r_s[i + j] != r {
					i++
					goto check
				}
			}
			return i
		}
	}
	return -1
}
