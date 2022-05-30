package piscine

func MakeRange(min, max int) []int {
	if min < max {
		ret := make([]int, max - min)
		for i := range ret {
			ret[i] = min + i
		}
		return ret
	} else {
		ret := []int{}
		return ret //return []int{}
	}
}
