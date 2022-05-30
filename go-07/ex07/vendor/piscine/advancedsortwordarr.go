package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	var len int
	for range a {
		len++
	}

	for i := 0; i < len; i++ {
		for j := 0; j < len - i - 1; j++ {
			if f(a[j], a[j + 1]) ==  1{
				a[j], a[j + 1] = a[j + 1], a[j]
			}
		}
	}
}
