package piscine

func containsIsland(board [][]int, blackNum int) bool {
	// 到達可能かどうかをチェックする。
	// 到達可能ならtrue
	// 黒マスもtrue
	// 到達不能な場所は、最終的にfalseのまま残る
	var reachables [][]bool = make([][]bool, len(board))

	for i := 0; i < len(reachables); i++ {
		reachables[i] = make([]bool, len(board[i]))
	}

	if board[0][0] != blackNum {
		checkIsReachable(board, reachables, 0, 0, blackNum)
	} else {
		checkIsReachable(board, reachables, 1, 0, blackNum)
	}

	for _, row := range reachables {
		for _, col := range row {
			if col == false {
				return true
			}
		}
	}
	return false
}

func checkIsReachable(board [][]int, reachables [][]bool, x int, y int, blackNum int) {
	// 対象が盤面の外に位置する場合、特に何もせずreturnする。
	if y < 0 || len(board) <= y || x < 0 || len(board[y]) <= x {
		return
	}

	if reachables[y][x] == true {
		return
	}

	reachables[y][x] = true
	if board[y][x] == blackNum {
		return
	}
	checkIsReachable(board, reachables, x-1, y, blackNum)
	checkIsReachable(board, reachables, x+1, y, blackNum)
	checkIsReachable(board, reachables, x, y-1, blackNum)
	checkIsReachable(board, reachables, x, y+1, blackNum)
}
