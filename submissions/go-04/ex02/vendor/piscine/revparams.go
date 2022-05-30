package piscine

import "ft"

func RevParams(args []string) {
	count := 0
	for range args {
		count++
	}

	for i := count - 1; i >= 0; i-- {
		for _, r := range args[i] {
			ft.PrintRune(r)
		}
		ft.PrintRune('\n')
	}
}

// a "b" "\\x\\\\" "\x\\"
//https://docs.microsoft.com/en-us/archive/blogs/twistylittlepassagesallalike/everyone-quotes-command-line-arguments-the-wrong-way