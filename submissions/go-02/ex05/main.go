package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Sqrt(4))
	fmt.Println(piscine.Sqrt(3))
}

// 1 - 1 // 2 - 1 // 3 - 1 //
// 4 - 2 // 5 - 2 // 6 - 1 // 7 - 1 // 8 - 1
// 9 - 3 // 10 - 3 // 11 - 3 // 12 - 3 // 13 - 3 // 14 - 3 // 15 - 3 //
// 3 - 5 - 7 - 9 ... num of solutions between sqrts. some way to solve using that??