package piscine

func getLen[T any](t []T) (len int) {
	for range t {
		len++
	}
	return
}

func Compact(ptr *[]string) int {
	emptyCount, i := 0, 0
	if ptr == nil {
		return 0
	}

	strCount := getLen(*ptr)
	for _, str := range *ptr {
		if str == "" {
			emptyCount++
		}
	}
	ret := make([]string, strCount - emptyCount)

	for _, str := range *ptr {
		if str != "" {
			ret[i] = str
			i++
		}
	}

	*ptr = ret
	return strCount - emptyCount
}
