package piscine

import (
	"fmt"
	"ft"
)

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
		r := []rune(s)
		for j, c := range r {
			if c == 'K' {
				return i, j
			}
		}
	}
	return -1, -1
}

func isCheckPiece(i, j int, board []string) bool {
	boardlen := boardLen(board)

	// i, jが範囲外になる場合
	if (i < 0 || i >= boardlen) || (j < 0 || j >= boardlen) {
		return false
	}
	ky, kx := getKingIndex(board)

	r := []rune(board[i])
	if r[j] == 'P' {
		if checkPawn(i, j, ky, kx) {
			return true
		}
	} else if r[j] == 'B' {
		if checkBishop(i, j, ky, kx) && checkBetweenBishop(i, j, ky, kx, board) {
			return true
		}
	} else if r[j] == 'R' {
		if checkRook(i, j, ky, kx) && checkBetweenRook(i, j, ky, kx, board) {
			return true
		}
	} else if r[j] == 'Q' {
		checker := checkQueen(i, j, ky, kx)
		if checker == ROOK && checkBetweenRook(i, j, ky, kx, board) {
			return true
		}
		if checker == BISHOP && checkBetweenBishop(i, j, ky, kx, board) {
			return true
		}
	}
	return false
}

func printStr(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	}
}

/// DEBUG
func showBoard(i, j int, board []string) {
	puts("[Show board]")
	for a, s := range board {
		for b, c := range s {
			if a == i && b == j {
				ft.PrintRune(c)
			} else if c == 'K' {
				ft.PrintRune(c)
			} else {
				ft.PrintRune('.')
			}
		}
		ft.PrintRune('\n')
	}
}

func putOutRoute(i, j int, board []string) {
	message := "Optimal route is ..\n"
	for _, c := range message {
		ft.PrintRune(c)
	}
	r := []rune(board[i])
	ft.PrintRune(r[j])
	printStr(" : (")
	ft.PrintRune(rune(i + '0'))
	printStr(", ")
	ft.PrintRune(rune(j + '0'))
	puts(")")

	// showBoard(i, j, board)
}

func printBoard(board []string) {
	for _, s := range board {
		for _, c := range s {
			fmt.Printf("%c", c)
		}
		fmt.Println("")
	}
}

func isCheckmate(board []string) bool {
	ky, kx := getKingIndex(board)

	boardlen := boardLen(board)
	for i := 1; ; i++ {
		// もうアクセスできる要素が存在しない場合
		if ky-i < 0 && ky+i >= boardlen && kx-i < 0 && kx+i >= boardlen {
			return false
		}
		// 右
		if isCheckPiece(ky, kx+i, board) {
			putOutRoute(ky, kx+i, board)
			return true
		} else if isCheckPiece(ky, kx-i, board) {
			// 左
			putOutRoute(ky, kx-i, board)
			return true
			// 上
		} else if isCheckPiece(ky-i, kx, board) {
			putOutRoute(ky-i, kx, board)
			return true
			// 左斜め上
		} else if isCheckPiece(ky-i, kx-i, board) {
			putOutRoute(ky-i, kx-i, board)
			return true
			// 右斜め上
		} else if isCheckPiece(ky-i, kx+i, board) {
			putOutRoute(ky-i, kx+i, board)
			return true
			// 下
		} else if isCheckPiece(ky+i, kx, board) {
			putOutRoute(ky+i, kx, board)
			return true
			// 左斜め下
		} else if isCheckPiece(ky+i, kx-i, board) {
			putOutRoute(ky+i, kx-i, board)
			return true
			// 右斜め下
		} else if isCheckPiece(ky+i, kx+i, board) {
			putOutRoute(ky+i, kx+i, board)
			return true
		}
	}
}

func puts(s string) {
	printStr(s)
	ft.PrintRune('\n')
}

func printAnswer(s, color string) {
	printStr(color)
	puts(s)
	printStr("\033[0m")
}

func Checkmate(board []string) {
	if !isValidBoard(board) {
		printAnswer("Fail", "\033[31m")
		return
	}

	if isCheckmate(board) {
		printAnswer("Success", "\033[32m")
	} else {
		printAnswer("Fail", "\033[31m")
	}
}
