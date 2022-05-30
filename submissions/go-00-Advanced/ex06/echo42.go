package main

import "flag"
import "fmt"
import "strings"

func echo42() {
	n := flag.Bool("n", false, "omit trailing newline")
	s := flag.String("s", " ", "separator")
	flag.Parse()
	args := flag.Args()

	fmt.Print(strings.Join(args, *s))
	if *n == false {
		fmt.Print("\n")
	}
}

func main() {
	echo42()
}
