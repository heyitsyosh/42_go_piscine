package piscine

func Atoi(s string) int {
	n := 0
	sign := 1
	for index, r := range s {
		if index == 0 {
			if r == '-' {
				sign *= -1
				continue
			}
			if r == '+' {
				continue
			}
		}
		if r >= '0' && r <= '9' {
			n = n*10 + int(r - '0')
		} else {
			return 0
		}
	}
	return n * sign
}
