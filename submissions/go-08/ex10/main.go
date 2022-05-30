package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4}
	b := []int{}
	unmatch := piscine.Unmatch(a)
	unmatch2 := piscine.Unmatch(b)
	fmt.Println(unmatch)
	fmt.Println(unmatch2)
}
