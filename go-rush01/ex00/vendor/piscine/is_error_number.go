package piscine

func isErrorNumber(board [][]int, row int, col int) bool {
	var count int = 1
	count += countWhiteSquares(board, row, col, 1, 0)
	count += countWhiteSquares(board, row, col, 0, 1)
	count += countWhiteSquares(board, row, col, -1, 0)
	count += countWhiteSquares(board, row, col, 0, -1)
	return count < board[row][col]
}

func countWhiteSquares(board [][]int, row int, col int, vertical int, horizontal int) int {
	for count := 0; ; count++ {
		row += vertical
		col += horizontal
		if row < 0 || len(board) <= row || col < 0 || len(board) <= col {
			return count
		}
		if board[row][col] == 0 {
			return count
		}
	}
}
