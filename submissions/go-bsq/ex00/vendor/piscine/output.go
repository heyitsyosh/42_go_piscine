package piscine

import "ft"

func Print_output(solved [][]int, b *bsq) {
	for i := 0; i < b.max_x; i++ {
		for j := 0; j < b.max_y; j++ {
			if (i > b.y-b.area) && (i <= b.y) && (j > b.x-b.area) && (j <= b.x) {
				ft.PrintRune(b.full)
			} else if solved[i][j] == 0 {
				ft.PrintRune(b.obstacle)
			} else {
				ft.PrintRune(b.empty)
			}
		}
		ft.PrintRune('\n')
	}
}
