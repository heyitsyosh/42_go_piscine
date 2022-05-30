package main

import "piscine"

func main() {
	a := piscine.SplitWhiteSpaces("Hello how are you?")
	piscine.PrintWordsTables(a)
	a = piscine.SplitWhiteSpaces(" Hello how are you?")
	piscine.PrintWordsTables(a)
	a = piscine.SplitWhiteSpaces("一二 三詩 語録 なな! a")
	piscine.PrintWordsTables(a)
	a = piscine.SplitWhiteSpaces("hello\n\t how are you  a a  ")
	piscine.PrintWordsTables(a)
}
