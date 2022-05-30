package piscine

import (
	"fmt"
	"io"
)

func PrintBoard(dst io.Writer, board [][]int, prePrint string, printBlack bool, postPrint string, whiteNum int, blackNum int) {
	fmt.Fprint(dst, GetBoardStr(board, prePrint, printBlack, postPrint, whiteNum, blackNum))
}

func GetBoardStr(board [][]int, prePrint string, printBlack bool, postPrint string, whiteNum int, blackNum int) (result string) {
	for _, row := range board {
		result += getRowStr(row, prePrint, printBlack, postPrint, true, whiteNum, blackNum)
	}
	return
}

func toRune(num int) rune {
	if num <= 0 {
		return '0'
	} else if num < 10 {
		return rune('0' + num)
	} else {
		return rune('a' + num - 10)
	}
}

func getRowStr(row []int, prePrint string, printBlack bool, postPrint string, isBlackColored bool, whiteNum int, blackNum int) (result string) {
	result += prePrint
	for _, col := range row {
		if col == whiteNum || (!printBlack && col == blackNum) {
			result += "."
		} else if col == blackNum {
			if isBlackColored {
				result += "\x1b[31m"
			}
			result += "B"
			if isBlackColored {
				result += "\x1b[0m"
			}
		} else {
			result += string(toRune(col))
		}
	}
	result += postPrint
	return
}
