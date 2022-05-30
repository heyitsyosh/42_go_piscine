package main

import "os"
import "ft"
import p "piscine"

func main() {
	args := os.Args

	if len(args) == 1 {
		input, status := p.From_stdin()
		if !status {
			return
		}
		solveStatus := p.First_line(input)
		if !solveStatus {
			return
		}
	} else if len(args) > 1 {
		for i := 1; i < len(args); i++ {
			input, status := p.From_file(args[i])
			if !status {
				continue
			}
			solveStatus := p.First_line(input) //in isvalid.go
			if !solveStatus {
				continue
			}
			if i != len(args)-1 {
				ft.PrintRune('\n')
			}
		}
	}
}
