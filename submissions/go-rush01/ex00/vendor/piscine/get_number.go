package piscine

func getNumber(board [][]int, row int, col int, vertical int, horizontal int) (int, int, bool) {
	for count := 0; ; count++ {
		row += vertical
		col += horizontal
		if row < 0 || len(board) <= row || col < 0 || len(board) <= col {
			return 0, 0, false
		}
		if board[row][col] > 1 {
			return row, col, true
		}
	}
}
