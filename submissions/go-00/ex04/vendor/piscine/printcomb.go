package piscine

import "ft"

func PutNum(a, b, c int) {
	ft.PrintRune(rune(a) + '0')
	ft.PrintRune(rune(b) + '0')
	ft.PrintRune(rune(c) + '0')
	if (a != 7) {
		ft.PrintRune(',')
		ft.PrintRune(' ')
	}
}

func PrintComb() {
	for a := 0; a <= 7; a++ {
		for b := a + 1; b <= 8; b++ {
			for c := b + 1; c <= 9; c++ {
				PutNum(a, b, c)
			}
		}
	}
	ft.PrintRune('\n')
}
