package piscine

func UltimatePointOne(n ***int) {
	if n == nil || *n == nil || **n == nil {
		return
	}
	***n = 1
}
