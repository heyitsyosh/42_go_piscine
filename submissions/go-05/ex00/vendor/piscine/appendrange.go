package piscine

func AppendRange(min, max int) []int {
	var ret []int
	for i := min; i < max; i++ {
		ret = append(ret, i)
	}
	return ret
}
