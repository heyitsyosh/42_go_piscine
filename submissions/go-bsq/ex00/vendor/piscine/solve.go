package piscine

func bsq_min(board [][]int, x, y int) int {
	if (board[x-1][y] <= board[x][y-1]) && (board[x-1][y] <= board[x-1][y-1]) {
		return board[x-1][y]
	} else if board[x][y-1] <= board[x-1][y-1] {
		return board[x][y-1]
	} else {
		return board[x-1][y-1]
	}
}

func set_result(y, x, max_area int, b *bsq) {
	b.y = y
	b.x = x
	b.area = max_area
}

func Solve(board [][]int, b *bsq) [][]int {
	//initialize solved board
	solved := make([][]int, b.max_x)
	for k := range solved {
		solved[k] = make([]int, b.max_y)
	}
	//fill first row and column of solved with info from setUpBoard/board
	for i := 0; i < b.max_x; i++ {
		for j := 0; j < b.max_y; j++ {
			if i == 0 {
				solved[i][j] = board[i][j]
			} else if j == 0 {
				solved[i][j] = board[i][j]
			}
			if solved[i][j] > b.area {
				set_result(i, j, solved[i][j], b)
			}
		}
	}
	for i := 1; i < b.max_x; i++ {
		for j := 1; j < b.max_y; j++ {
			if board[i][j] != 0 {
				solved[i][j] = bsq_min(solved, i, j) + 1
				if solved[i][j] > b.area {
					set_result(i, j, solved[i][j], b)
				}
			}
		}
	}
	return solved
}
