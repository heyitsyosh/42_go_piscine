package piscine

func Map(f func(int) bool, a []int) []bool {
	var ret []bool
	for _, n := range a {
		ret = append(ret, f(n))
	}
	return ret
}
