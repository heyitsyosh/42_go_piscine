package piscine

func Kuromasu(args []string) {
	if !isValidInput(args) {
		return
	}

	board := setBoard(args)

	if !isValidNumInBoard(board) {
		return
	}

	kuromasuResolve(board, 0, 0)
}

func kuromasuResolve(board [][]int, row int, col int) bool {
	if isSolved(board) {
		printBoard(board)
		return true
	}
	if row == len(board) {
		return false
	}
	if col == len(board) {
		return kuromasuResolve(board, row+1, 0)
	}
	if board[row][col] != 1 {
		return kuromasuResolve(board, row, col+1)
	}
	setKuromasu(board, row, col, 0)
	if !isError(board, row, col) && kuromasuResolve(board, row, col+1) {
		return true
	}
	setKuromasu(board, row, col, 1)
	return kuromasuResolve(board, row, col+1)
}

func setKuromasu(board [][]int, row int, col int, c int) {
	board[row][col] = c
}
