package piscine

import "ft"

func PrintProgramName(args []string) {
	var slash int
	r_arg := []rune(args[0])
	for i, r := range r_arg {
		if r == '/' {
			slash = i
		}
	}

	for i, r := range r_arg {
		if i > slash {
			ft.PrintRune(r)
		}
	}
	ft.PrintRune('\n')
}
