package piscine

import "ft"

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = -INT_MAX - 1

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

func PrintNbrBase(nbr int, base string) {
	baseLen, i := strLen(base), 1
	r_base := []rune(base)

	if !isValid(baseLen, r_base) {
		ft.PrintRune('N')
		ft.PrintRune('V')
		return
	}
	if nbr == INT_MIN {
		ft.PrintRune('-')
		PrintNbrBase((nbr / baseLen) * -1, base)
		PrintNbrBase((nbr % baseLen) * -1, base)
		return
	}
	if nbr < 0 {
		nbr *= -1
		ft.PrintRune('-')
	}
	for nbr / baseLen >= i {
		i *= baseLen
	}
	for i > 0 {
		ft.PrintRune(rune(r_base[nbr / i]))
		nbr %= i
		i /= baseLen
	}
	return
}
