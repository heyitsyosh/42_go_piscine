package piscine

func argsCount(args []string) (len int) {
	for range args {
		len++
	}
	return
}

func ConcatParams(args []string) string {
	joined := ""
	count := argsCount(args)
	for i := 0; i < count - 1; i++ {
		joined += args[i]
		joined += "\n"
	}
	if count > 0 {
		joined += args[count - 1]
	}
	return joined
}
