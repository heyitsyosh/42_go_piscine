package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsAlpha("Hello! How are you?"))
	fmt.Println(piscine.IsAlpha("HelloHowareyou"))
	fmt.Println(piscine.IsAlpha("Whats this 4?"))
	fmt.Println(piscine.IsAlpha("Whatsthis4"))
	fmt.Println(piscine.IsAlpha(""))
}
