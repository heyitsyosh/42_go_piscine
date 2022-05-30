package piscine

func isErrorBlack(board [][]int, row int, col int) bool {
	// 黒の上下左右に黒が配置されていないことをチェック
	return isBlack(board, col-1, row) ||
		isBlack(board, col+1, row) ||
		isBlack(board, col, row-1) ||
		isBlack(board, col, row+1)
}

func isBlack(board [][]int, x int, y int) bool {
	if y < 0 || len(board) <= y || x < 0 || len(board[y]) <= x {
		return false
	}
	return board[y][x] == 0
}
