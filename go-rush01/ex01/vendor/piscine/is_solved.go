package piscine

func isSolved(board [][]int) bool {
	for row, line := range board {
		for col, c := range line {
			if c != 1 && c != 0 {
				if !isValid(board, row, col) {
					return false
				}
			}
		}
	}
	return true
}

func isValid(board [][]int, row int, col int) bool {
	var count int = 1
	count += countWhiteSquares(board, row, col, 1, 0)
	count += countWhiteSquares(board, row, col, 0, 1)
	count += countWhiteSquares(board, row, col, -1, 0)
	count += countWhiteSquares(board, row, col, 0, -1)
	return count == board[row][col]
}
