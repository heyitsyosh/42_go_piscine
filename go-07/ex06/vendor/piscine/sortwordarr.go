package piscine

func SortWordArr(a []string) {
	var len int
	for range a {
		len++
	}

	for i := 0; i < len; i++ {
		for j := 0; j < len - i - 1; j++ {
			if a[j] > a[j + 1] {
				a[j], a[j + 1] = a[j + 1], a[j]
			}
		}
	}
}
