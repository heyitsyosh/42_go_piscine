package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.AppendRange(5, 10))
	fmt.Println(piscine.AppendRange(10, 5))
	fmt.Println(piscine.AppendRange(5, 5))
	fmt.Println(piscine.AppendRange(4, 5))
}
