package piscine

func BasicJoin(elems []string) string {
	joined := ""
	for _, str := range elems {
		joined += str
	}
	return joined
}
