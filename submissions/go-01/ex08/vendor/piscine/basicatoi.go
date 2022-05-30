package piscine

func BasicAtoi(s string) int {
	n := 0
	for _, r := range s {
		if r <= '9' && r >= '0' {
			n = n*10 + (int(r) - '0')
		}
	}
	return n
}
