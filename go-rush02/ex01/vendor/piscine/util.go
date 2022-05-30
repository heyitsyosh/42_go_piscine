package piscine

func runesLen(s []rune) int {
	var len int = 0
	for range s {
		len++
	}
	return len
}

// assuming: |a| == |b|
func runesEqual(a, b []rune, n int) bool {
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
