package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"piscine"
	"time"
	"path/filepath"
)

func main() {
	arg_count := len(os.Args) - 1
	if arg_count <= 0 {
		return
	}
	switch os.Args[1] {
	case "--gen":
		rand.Seed(time.Now().UnixNano())
		if arg_count >= 5 {
			// ボード生成
			result, err := piscine.BoardGeneratorStr(os.Args[2], os.Args[3], os.Args[4], os.Args[5])

			if err != nil {
				fmt.Fprintf(os.Stderr, "%s", err.Error())
				return
			}

			// 期待される答えを、Stderrに出力
			piscine.PrintBoard(os.Stderr, result, "", true, "\n", piscine.GENERATOR_WHITE_NUM, piscine.GENERATOR_BLACK_NUM)

			// Solverに与える引数の形でStdoutに出力
			piscine.PrintBoard(os.Stdout, result, "\"", false, "\" ", piscine.GENERATOR_WHITE_NUM, piscine.GENERATOR_BLACK_NUM)
		}

	case "--serv":
		fmt.Println("HTTP Server Starting...")
		http.HandleFunc("/v1/boardgen", piscine.HttpBoardGeneratorHandler)
		http.HandleFunc("/v1/kuromasu", piscine.HttpKuromasuSolverHandler)
		err := http.ListenAndServe(":80", nil)

		if err != nil {
			_ = fmt.Errorf("%s", err.Error())
			return
		}

	case "--visual":
		result := piscine.Kuromasu(os.Args[2:], true)
		if result == nil {
			fmt.Fprintln(os.Stderr, "Could not resolve!!")
		}
	case "--help":
		fmt.Printf("Usage:\n")
		exe_name := filepath.Base(os.Args[0])
		fmt.Printf("\t./%s <board>\n", exe_name)
		fmt.Println("\t\t<board> is a string of numbers separated by space.")
		fmt.Printf("\t./%s --gen <width> <height> <white_num> <black_num>\n", exe_name)
		fmt.Println("\t\t<width> and <height> are the width and height of the board.")
		fmt.Println("\t\t<white_num> and <black_num> are the number of white and black cells.")
		fmt.Printf("\t./%s --serv\n", exe_name)
		fmt.Println("\t\tStart HTTP Server.")
		fmt.Printf("\t./%s --visual <board>\n", exe_name)
		fmt.Println("\t\tVisualize the board.")
		fmt.Printf("\t./%s --help\n", exe_name)
		fmt.Println("\t\tShow this message.")

	default:
		result := piscine.Kuromasu(os.Args[1:], false)
		if result != nil {
			piscine.PrintBoard(os.Stdout, result, "", true, "\n", piscine.SOLVER_WHITE_NUM, piscine.SOLVER_BLACK_NUM)
		}
	}
}
