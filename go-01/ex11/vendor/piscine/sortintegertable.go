package piscine

func SortIntegerTable(table []int) {
	len := 0
	for range table {
		len += 1
	}
	for i := 0; i < len; i++ {
		for j := 0; j < len - i - 1; j++ {
			if table[j] > table[j + 1] {
				table[j], table[j + 1] = table[j + 1], table[j]
			}
		}
	}
}
