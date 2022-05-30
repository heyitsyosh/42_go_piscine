package piscine

func getLen[T any](t []T) (len int) {
	for range t {
		len++
	}
	return
}

func Max(a []int) int {
	if getLen(a) <= 0 {
		return 0
	}
	max := a[0]
	for _, n := range a[1:] {
		if n > max {
			max = n
		}
	}
	return max
}
