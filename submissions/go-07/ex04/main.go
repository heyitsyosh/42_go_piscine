package main

import (
	"fmt"
	"piscine"
)

func main() {
	a1 := []int{0, 1, 2, 2, 3, 4, 5}
	a2 := []int{5, 4, 3, 2, 2, 1, 0}
	a3 := []int{0, 2, 1, 3}
	result1 := piscine.IsSorted(piscine.Cmp, a1)
	result2 := piscine.IsSorted(piscine.Cmp, a2)
	result3 := piscine.IsSorted(piscine.Cmp, a3)
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}
