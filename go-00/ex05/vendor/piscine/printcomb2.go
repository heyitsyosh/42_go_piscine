package piscine

import "ft"

func PutNum(n int) {
	if n < 10 {
		ft.PrintRune('0')
		ft.PrintRune(rune(n) + '0')
	} else {
		ft.PrintRune((rune(n) / 10) + '0')
		ft.PrintRune((rune(n) % 10) + '0')
	}
}

func PrintComb2() {
	for a := 0; a <= 98; a++ {
		for b := a + 1; b <= 99; b++ {
			PutNum(a)
			ft.PrintRune(' ')
			PutNum(b)
			if a == 98 && b == 99 {
				ft.PrintRune('\n')
				return
			} else {
				ft.PrintRune(',')
				ft.PrintRune(' ')
			}
		}
	}
}
