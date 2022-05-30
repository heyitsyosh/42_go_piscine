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

	if argCount != 4 {
		return
	}
	if args[2] != "+" && args[2] != "-" && args[2] != "/" && args[2] != "%" && args[2] != "*" {
		return
	}
	a, aValid := p.Atoi(args[1])
	b, bValid := p.Atoi(args[3])
	if !aValid || !bValid {
		return
	}
	p.Doop(args, a, b)
}
