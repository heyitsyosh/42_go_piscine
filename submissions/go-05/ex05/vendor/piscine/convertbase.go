package piscine

func strLen(s string) (len int) {
	for range s {
		len++
	}
	return
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	r_baseTo := []rune(baseTo)
	digitsFrom, digitsTo := strLen(baseFrom), strLen(baseTo)
	n, baseIndex, div, ret := 0, 0, 1, ""

	for _, nbr_digit := range nbr {
		baseIndex = 0
		for _, base_digit := range baseFrom {
			if nbr_digit == base_digit {
				break
			}
			baseIndex++
		}
		n = n * digitsFrom + baseIndex
	}

	for n / digitsTo >= div {
		div *= digitsTo
	}

	for div > 0 {
		ret += string(r_baseTo[n / div])
		//ret += string(r_baseTo[n % digitsTo]) + ret --> from right
		n %= div
		div /= digitsTo
	}
	return ret
}
