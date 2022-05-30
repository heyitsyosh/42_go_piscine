package piscine

import (
	"ft"
)

func printBoard(board [][]int) {
	for _, line := range board {
		for _, c := range line {
			switch c {
			case 0:
				ft.PrintRune('B')
			case 1:
				ft.PrintRune('.')
			default:
				if c < 10 {
					ft.PrintRune(rune('0' + c))
				} else {
					ft.PrintRune(rune('a' + c - 10))
				}
			}
		}
		ft.PrintRune('\n')
	}
}
