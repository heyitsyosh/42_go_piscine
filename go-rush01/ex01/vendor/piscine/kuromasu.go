package piscine

import (
	"fmt"
	"os"
)

func Kuromasu(args []string, vflg bool) [][]int {
	if !isValidInput(args) {
		return nil
	}
	board := setBoard(args)

	if !isValidNumInBoard(board) {
		return nil
	}

	if vflg {
		fmt.Print(DISABLE_CURSOR + CLEAR)
	}
	result := kuromasuResolve(board, 0, 0, vflg)
	if vflg {
		fmt.Print(ENABLE_CURSOR + RESET)
	}

	if result {
		return board
	} else {
		return nil
	}
}

func kuromasuResolve(board [][]int, row int, col int, vflg bool) bool {
	if isSolved(board) {
		if vflg {
			fmt.Print(TOP_LEFT)
			PrintBoard(os.Stdout, board, "", true, "\n", SOLVER_WHITE_NUM, SOLVER_BLACK_NUM)
		}
		return true
	}
	if row == len(board) {
		return false
	}
	if col == len(board) {
		return kuromasuResolve(board, row+1, 0, vflg)
	}
	if board[row][col] != 1 {
		return kuromasuResolve(board, row, col+1, vflg)
	}
	setKuromasu(board, row, col, 0)
	if vflg {
		fmt.Print(TOP_LEFT)
		PrintBoard(os.Stdout, board, "", true, "\n", SOLVER_WHITE_NUM, SOLVER_BLACK_NUM)
	}
	if !isError(board, row, col) && kuromasuResolve(board, row, col+1, vflg) {
		return true
	}
	setKuromasu(board, row, col, 1)
	return kuromasuResolve(board, row, col+1, vflg)
}

func setKuromasu(board [][]int, row int, col int, c int) {
	board[row][col] = c
}
