package piscine

import "ft"

func PrintNbr(nbr int) {
	div := 1
	if nbr < 0 {
		ft.PrintRune('-')
		PrintNbr((nbr / 10) * -1)
		PrintNbr((nbr % 10) * -1)
		return
	}
	for nbr / 10 >= div {
		div *= 10
	}
	for div > 0 {
		ft.PrintRune(rune(nbr / div) + '0')
		nbr %= div
		div /= 10
	}
}
