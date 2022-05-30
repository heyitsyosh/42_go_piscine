package piscine

func ActiveBits(n int) int {
	nb := uint(n)
	var count uint = 0
	for nb != 0 {
		count += nb & 1
		nb >>= 1
	}
	return int(count)
}
