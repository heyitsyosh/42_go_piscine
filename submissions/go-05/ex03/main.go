package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello how are you?"))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces(" Hello how are you?"))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("一二 三詩 語録 なな! a"))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("hello\n\t how are you  a a  "))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces(""))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces(" a"))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces(" a "))
	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces(" a  "))
}
