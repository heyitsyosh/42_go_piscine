package main

import p "piscine"
import "os"

func main() {
	p.PrintParams(os.Args[1:])
}
