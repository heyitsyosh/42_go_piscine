package piscine

import "ft"

func PrintReverseAlphabet() {
	for r := 'z'; r >= 'a'; r-- {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}
