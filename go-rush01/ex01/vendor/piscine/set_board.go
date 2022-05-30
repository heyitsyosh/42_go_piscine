package piscine

func setBoard(args []string) [][]int {
	height := len(args)
	width := len(args[0])

	var board [][]int = make([][]int, height)
	for row, line := range args {
		board[row] = make([]int, width)
		for col, c := range line {
			switch c {
			case '.':
				board[row][col] = 1
			default:
				tmp := toInt(c)
				board[row][col] = tmp
			}
		}
	}
	return board
}

func toInt(c rune) int {
	if '0' <= c && c <= '9' {
		return int(c - '0')
	} else if 'a' <= c && c <= 'z' {
		return int(c-'a') + 10
	} else if 'A' <= c && c <= 'Z' {
		return int(c-'A') + 10
	} else {
		return -1
	}
}
