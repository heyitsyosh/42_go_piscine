package piscine

func isError(board [][]int, row int, col int) bool {
	if isErrorBlack(board, row, col) {
		return true
	}
	if containsIsland(board, 0) {
		return true
	}
	x, y, ok := getNumber(board, row, col, 1, 0)
	if ok && isErrorNumber(board, x, y) {
		return true
	}
	x, y, ok = getNumber(board, row, col, 0, 1)
	if ok && isErrorNumber(board, x, y) {
		return true
	}
	x, y, ok = getNumber(board, row, col, -1, 0)
	if ok && isErrorNumber(board, x, y) {
		return true
	}
	x, y, ok = getNumber(board, row, col, 0, -1)
	if ok && isErrorNumber(board, x, y) {
		return true
	}
	return false
}
