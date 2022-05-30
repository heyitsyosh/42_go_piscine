package main

import "piscine"

func checkmate(board []string) {
	piscine.Checkmate(board)
}

func main() {
	board := []string{
		"あいうえお",
		"あいうえお",
		"あいKえお",
		"あPうえお",
		"あいうえお",
	}
	checkmate(board)
	piscine.SuccessCase()
	piscine.FailCase()
}
