package piscine

import "ft"

func PrintAlphabet() {
	for r := 'a'; r <= 'z'; r++ {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}
