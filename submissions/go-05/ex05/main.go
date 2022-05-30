package main

import (
	"fmt"
	"piscine"
)

const INT_MAX = int(^uint(0) >> 1) //9223372036854775807

func main() {
	result := piscine.ConvertBase("101011", "01", "0123456789")
	fmt.Println(result)
	result = piscine.ConvertBase("9223372036854775807", "0123456789", "01")
	fmt.Println(result)
	result = piscine.ConvertBase("九二二三三七二〇三六八五四七七五八〇七", "〇一二三四五六七八九", "01")
	fmt.Println(result)
	result = piscine.ConvertBase("九二二三三七二〇三六八五四七七五八〇七", "〇一二三四五六七八九", "0123456789")
	fmt.Println(result)
}
