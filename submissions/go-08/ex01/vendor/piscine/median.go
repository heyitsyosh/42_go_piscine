package piscine

func Median(a, b, c, d, e int) int {
	toSort := []int{a, b, c, d, e}
	for i := 0; i < 5; i++ {
		for j := 0; j < 4 - i; j++ {
			if toSort[j] > toSort[j + 1] {
				toSort[j], toSort[j + 1] = toSort[j + 1], toSort[j]
			}
		}
	}
	return toSort[2]
}
