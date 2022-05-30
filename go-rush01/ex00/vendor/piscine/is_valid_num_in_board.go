package piscine

const IS_MANDATORY = true
const MANDATORY_HW = 5

func isValidNumInBoard(board [][]int) bool {
	height := len(board)
	width := len(board[0])
	maxVal := height + width - 1

	if IS_MANDATORY {
		if height != MANDATORY_HW || width != MANDATORY_HW {
			return false
		}
	}
	for _, row := range board {
		for _, v := range row {
			if maxVal < v {
				return false
			}
		}
	}
	return true
}
