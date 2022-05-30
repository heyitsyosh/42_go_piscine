package main

import p "piscine"
import "ft"
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
	if argCount <= 1 {
		os.Stdout.WriteString("File name missing\n")
		return
	} else if argCount > 2 {
		os.Stdout.WriteString("Too many arguments\n")
		return
	} else {
		p.DisplayFile(args[1])
	}
}
