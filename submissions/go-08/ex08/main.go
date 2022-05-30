package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := []int{23, 1212343, 1, 11, 55, 93}
	max := piscine.Max(a)
	fmt.Println(max)
}
