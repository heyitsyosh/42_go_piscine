package main

import (
	"fmt"
	"piscine"
)

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(piscine.ConcatParams(test))
	// fmt.Println("-------------------")
	// test2 := []string{"こんにちは", "げんき", " ですか一二三詩後", "?"}
	// fmt.Println(piscine.ConcatParams(test2))
	// fmt.Println("-------------------")
	// test3 := []string{""}
	// fmt.Println(piscine.ConcatParams(test3))
	// fmt.Println("-------------------")
	// test3 = []string{"hello", "\n\n", "\n"}
	// fmt.Println(piscine.ConcatParams(test3))
}
