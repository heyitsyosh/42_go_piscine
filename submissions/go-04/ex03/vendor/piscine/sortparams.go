package piscine

import "ft"

func printParams(args []string) {
	for _, line := range args {
		for _, r := range line {
			ft.PrintRune(r)
		}
		ft.PrintRune('\n')
	}
}

func SortParams(args []string) {
	var len int
	for range args {
		len++
	}

	for i := 0; i < len; i++ {
		for j := 0; j < len - i - 1; j++ {
			if args[j] > args[j + 1] {
				args[j], args[j + 1] = args[j + 1], args[j]
			}
		}
	}
	printParams(args)
}
