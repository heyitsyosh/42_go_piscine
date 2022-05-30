package piscine

func Unmatch(a []int) int {
	tally := map[int]int{}
	for _, n := range a {
		tally[n] += 1
	}
	
	for _, n := range a {
		if tally[n] & 1 == 1 {
			return n
		}
	}
	return -1
}
