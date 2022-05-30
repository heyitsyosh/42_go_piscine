package piscine

import (
	"ft"
)

func StrLen(s string) int {
	n := 0
	for range s {
		n++
	}
	return n
}

func SliceLen[T any](s []T) int {
	n := 0
	for range s {
		n++
	}
	return n
}

func PrintStr(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	}
}

func PrintLn(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}

func isRuneDigit(c rune) bool {
	return '0' <= c && c <= '9'
}

func PrintError(msg string) {
	PrintStr("Error: " + msg + "\n")
}

func Make2DSlice[T any](height, width int, init T) [][]T {
	rv := make([][]T, height, height)
	for i := 0; i < height; i++ {
		rv[i] = make([]T, width, width)
	}
	return rv
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Join(s []string, sep string) string {
	rv := ""
	for i, a := range s {
		if i > 0 {
			rv += sep
		}
		rv += a
	}
	return rv
}
