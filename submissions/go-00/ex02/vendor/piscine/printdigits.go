package piscine

import "ft"

func PrintDigits() {
	for r := '0'; r <= '9'; r++ {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}
