package main

import (
	"fmt"
	"piscine"
)

func main() {
	s := "Hello World!あ"
	s = piscine.StrRev(s)
	fmt.Println(s)
}
