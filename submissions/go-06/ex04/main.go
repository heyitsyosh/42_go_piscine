package main

import p "piscine"
import "os"

func argLen(args []string) (len int) {
	for range args {
		len++
	}
	return
}

func main() {
	var hadError bool
	args := os.Args

	if argLen(args) == 4 {
		hadError = p.Ztail(args[2], args[3], 0)
	} else if argLen(args) > 4 {
		for i := 3; i < argLen(args); i++ {
			hasError := p.Ztail(args[2], args[i], i) //return error
			if hasError == true {
				hadError = true
			}
		}
	} else {
		return
	}
	if hadError {
		os.Exit(1)
	}
}
