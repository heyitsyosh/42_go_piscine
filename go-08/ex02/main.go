package main

import p "piscine"
import "os"

func argsCount(args []string) (count int) {
	for range args {
		count++
	}
	return
}

func main() {
	args := os.Args
	count := argsCount(args)
	if count == 0 {
		return
	}
	p.ComCheck(args[1:])
}
