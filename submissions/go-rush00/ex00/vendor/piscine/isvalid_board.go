package piscine

func calcIntMax() int {
	var max uint = 0
	return int((^max) >> 1)
}

func boardLen(board []string) int {
	max := calcIntMax()
	l := 0
	for range board {
		if l == max {
			return -1
		}
		l += 1
	}
	return l
}

func strLen(s string) int {
	max := calcIntMax()
	l := 0
	for range s {
		if l == max {
			return -1
		}
		l += 1
	}
	return l
}

// KINGの数が1個以外 or 存在しない場合はOUT
func isOnlyOneKing(board []string) bool {
	kings := 0
	for _, s := range board {
		for _, c := range s {
			if c == 'K' {
				kings += 1
				if kings == 2 {
					return false
				}
			}
		}
	}
	return kings == 1
}

func isSquare(board []string) bool {
	height := boardLen(board)

	// 縦横の長さが同じか
	for _, s := range board {
		if strLen(s) != height {
			return false
		}
	}
	return true
}

func IsValidBoard(board []string) bool {
	if !isSquare(board) {
		return false
	}
	if !isOnlyOneKing(board) {
		return false
	}
	return true
}
