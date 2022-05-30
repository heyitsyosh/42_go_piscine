package piscine

import "ft"

// not using backtracking
var board [9]int

func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}

func printSol(n int) {
	for i := 1; i <= n; i++ {
		ft.PrintRune(rune(board[i]) + '0')
	}
	ft.PrintRune('\n')
}

func place(row, col int) bool {
	for i := 1; i < row; i++ {
		if board[i] == col || (abs(board[i] - col) == abs(i - row)) { //compare queen to place with queens from other rows (board[i])
			return false
		}
	}
	return true
}

func solve(n, row int) {
	for col := 1; col <= n; col++ {
		if place(row, col) { //check for conflicts
			board[row] = col //place queen (store col in according row)
			if row == n { //dead end
				printSol(n)
			} else { //try queen with next pos
				solve(n, row + 1)
			}
		}
	}
}

func EightQueens() {
	solve(8, 1)
}
