package main

import (
	"flag"
	"piscine"
)

func checkmate(board []string) {
	piscine.Checkmate(board)
}

func main() {
	l := flag.Bool("l", false, "show chess lion")
	flag.Parse()

	if *l == true {
		piscine.ShowLion()
		return
	}
	board := flag.Args()

	checkmate(board)
	piscine.SuccessCase()
	piscine.FailCase()
}

