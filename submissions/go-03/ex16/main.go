package main

import (
	"fmt"
	"piscine"
)

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = -INT_MAX - 1

func main() {
	piscine.PrintNbrBase(125, "0123456789")
	fmt.Println("")
	piscine.PrintNbrBase(-125, "01")
	fmt.Println("")
	piscine.PrintNbrBase(125, "0123456789ABCDEF")
	fmt.Println("")
	piscine.PrintNbrBase(-125, "choumi")
	fmt.Println("")
	piscine.PrintNbrBase(125, "aa")
	fmt.Println("")
	piscine.PrintNbrBase(125, "01")
	fmt.Println("")
	piscine.PrintNbrBase(125, "あい")
	fmt.Println("")
	piscine.PrintNbrBase(125, "0123")
	fmt.Println("")
	piscine.PrintNbrBase(125, "あいうえ")
	fmt.Println("")
	piscine.PrintNbrBase(INT_MAX, "01")
	fmt.Println("")
	piscine.PrintNbrBase(INT_MAX, "0123456789")
	fmt.Println("")
	piscine.PrintNbrBase(INT_MIN, "01")
	fmt.Println("")
	piscine.PrintNbrBase(INT_MIN, "0123456789")
	fmt.Println("")
}
