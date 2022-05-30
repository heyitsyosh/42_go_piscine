package piscine

func strLen(s string) (len int) {
	for range s {
		len++
	}
	return
}

func isValid(len int, base []rune) bool {
	for _, r := range base{
		if r == '+' || r == '-' {
			return false
		}
	}
	for i := 0; i < len - 1; i++ {
		for j := i + 1; j < len; j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}
	return len > 1
}

func AtoiBase(s string, base string) int {
    baseLen, ret := strLen(base), 0
    var baseIndex int

    if !isValid(baseLen, []rune(base)) {
        return 0
    }
    for _, s_digit := range s {
        baseIndex = 0
        for _, base_digit := range base {
            if s_digit == base_digit {
                ret = ret * baseLen + baseIndex
            }
            baseIndex++
        }
    }
    return ret
}
