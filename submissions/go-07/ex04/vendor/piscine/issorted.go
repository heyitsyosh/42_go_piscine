package piscine

func ascending(f func(a, b int) int, a []int) bool {
	for i, _ := range a {
		if i == 0 {
			continue
		}
		if f(a[i - 1], a[i]) == 1 {
			return false
		}
	}
	return true
}

func descending(f func(a, b int) int, a []int) bool {
	for i, _ := range a {
		if i == 0 {
			continue
		}
		if f(a[i - 1], a[i]) == -1 {
			return false
		}
	}
	return true
}

func IsSorted(f func(a, b int) int, a []int) bool {
	return ascending(f, a) || descending(f, a)
}
