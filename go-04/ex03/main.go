package main

import p "piscine"
import "os"

func main() {
	p.SortParams(os.Args[1:])
}
