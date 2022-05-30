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
	args := os.Args
	argCount := argLen(args)
	if argCount == 1 {
		p.FromStdin()
	} else if argCount != 0 {
		for i := 1; i < argCount; i++ {
			p.FromFile(args[i], argCount)
		}
	}
}
