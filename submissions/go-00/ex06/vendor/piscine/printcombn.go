package piscine

import "ft"

func PutComb(num int, arr [9]int) {
	max := (10 - num)
	for i := 0; i < num; i++ {
		ft.PrintRune(rune(arr[i]) + '0')
	}
	if arr[0] != max {
		ft.PrintRune(',')
		ft.PrintRune(' ')
	}
}

func rec(index, i, num int, arr [9]int) {
	if index < num {
		for ; i < 10; i++ {
			arr[index] = i
			rec(index + 1, i + 1, num, arr)
		}
	} else {
		PutComb(num, arr)
	}
}

func PrintCombN(n int) {
	var arr [9]int
	if n < 0 || n > 10 {
		return
	}
	rec(0, 0, n, arr)
	ft.PrintRune('\n')
}
