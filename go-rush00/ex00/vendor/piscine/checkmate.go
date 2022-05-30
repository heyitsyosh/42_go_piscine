package piscine

import "ft"

const (
	ROOK = iota
	BISHOP
	NOTFOUND
)

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func isPiece(r rune) bool {
	if r == 'P' || r == 'B' || r == 'R' || r == 'Q' {
		return true
	}
	return false
}

// KingがPawnの斜め上に存在するか
func checkPawn(by, bx, ky, kx int) bool {
	if by-1 == ky {
		if abs(bx-kx) == 1 {
			return true
		}
	}
	return false
}

// KingとBishopが同じ線上に存在するか
func checkBishop(by, bx, ky, kx int) bool {
	if abs(by-ky) == abs(bx-kx) {
		return true
	}
	return false
}

// KingとBishopの途中にコマがないか
func checkBetweenBishop(by, bx, ky, kx int, board []string) bool {
	loops := abs(by - ky)

	// 左上(-, -)
	if ky-by < 0 && kx-bx < 0 {
		for i := 1; i < loops; i++ {
			if isPiece(rune(board[by-i][bx-i])) {
				return false
			}
		}
	}
	// 左下(+, -)
	if ky-by > 0 && kx-bx < 0 {
		for i := 1; i < loops; i++ {
			if isPiece(rune(board[by+i][bx-i])) {
				return false
			}
		}
	}

	// 右上(-, +)
	if ky-by < 0 && kx-bx > 0 {
		for i := 1; i < loops; i++ {
			if isPiece(rune(board[by-i][bx+i])) {
				return false
			}
		}
	}

	// 右下(+, +)
	if ky-by > 0 && kx-bx > 0 {
		for i := 1; i < loops; i++ {
			if isPiece(rune(board[by+i][bx+i])) {
				return false
			}
		}
	}
	return true
}

// KingとRookの途中にコマがないか
func checkBetweenRook(by, bx, ky, kx int, board []string) bool {
	if by == ky {
		loops := abs(kx - bx)
		for i := 1; i < loops; i++ {
			if kx-bx < 0 {
				// 左に進む
				if isPiece(rune(board[by][bx-i])) {
					return false
				}
			} else {
				// 右に進む
				if isPiece(rune(board[by][bx+i])) {
					return false
				}
			}
		}
	} else if bx == kx {
		loops := abs(ky - by)
		for i := 1; i < loops; i++ {
			if ky-by < 0 {
				// 上に進む
				if isPiece(rune(board[by-i][bx])) {
					return false
				}
			} else {
				// 下に進む
				if isPiece(rune(board[by+i][bx])) {
					return false
				}
			}
		}
	}
	return true
}

// KingとRookが同じ線上に存在するか
func checkRook(by, bx, ky, kx int) bool {
	if by == ky || bx == kx {
		return true
	}
	return false
}

// KingとQueenが同じ線上に存在するか
func checkQueen(by, bx, ky, kx int) int {
	if checkRook(by, bx, ky, kx) {
		return ROOK
	} else if checkBishop(by, bx, ky, kx) {
		return BISHOP
	} else {
		return NOTFOUND
	}
}

func getKingIndex(board []string) (int, int) {
	for i, s := range board {
		for j, c := range s {
			if c == 'K' {
				return i, j
			}
		}
	}
	return -1, -1
}

func isCheckmate(board []string) bool {

	ky, kx := getKingIndex(board)

	for i, s := range board {
		for j, _ := range s {
			if board[i][j] == 'P' {
				if checkPawn(i, j, ky, kx) {
					return true
				}
			} else if board[i][j] == 'B' {
				if checkBishop(i, j, ky, kx) && checkBetweenBishop(i, j, ky, kx, board) {
					return true
				}
			} else if board[i][j] == 'R' {
				if checkRook(i, j, ky, kx) && checkBetweenRook(i, j, ky, kx, board) {
					return true
				}
			} else if board[i][j] == 'Q' {
				checker := checkQueen(i, j, ky, kx)
				if checker == ROOK && checkBetweenRook(i, j, ky, kx, board) {
					return true
				}
				if checker == BISHOP && checkBetweenBishop(i, j, ky, kx, board) {
					return true
				}
			}
		}
	}
	return false
}

func puts(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}

func Checkmate(board []string) {
	if !IsValidBoard(board) {
		puts("Fail")
		return
	}

	if isCheckmate(board) {
		puts("Success")
	} else {
		puts("Fail")
	}
}
