package piscine

func strCount(strs []string) (count int) {
	for range strs {
		count++
	}
	return
}

func Join(strs []string, sep string) string {
	joined := ""
	count := strCount(strs)
	for i := 0; i < count - 1; i++{
		joined += strs[i]
		joined += sep
	}
	if count > 0 {
		joined += strs[count - 1]
	}
	return joined
}
