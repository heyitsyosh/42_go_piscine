package piscine

func CountIf(f func(string) bool, tab []string) int {
	var count int
	for _, str := range tab {
		if f(str) {
			count++
		}
	}
	return count
}
