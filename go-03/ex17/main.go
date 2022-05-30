package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.AtoiBase("125", "0123456789"))
	fmt.Println(piscine.AtoiBase("1111101", "01"))
	fmt.Println(piscine.AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(piscine.AtoiBase("uoあ", "choumあ"))
	fmt.Println(piscine.AtoiBase("bbbbbab", "a"))
	
	fmt.Println(piscine.AtoiBase("uuu", "aiueo"))
	fmt.Println(piscine.AtoiBase("ううう", "あいうえお"))
	fmt.Println(piscine.AtoiBase("一二三", "〇一二三四五六七八九"))
}
