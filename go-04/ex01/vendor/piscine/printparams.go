package piscine

import "ft"

func PrintParams(args []string) {
	for _, line := range args {
		for _, r := range line {
			ft.PrintRune(r)
		}
		ft.PrintRune('\n')
	}
}
