package main

import (
	"ft"
	"piscine"
)

func main() {
	ft.PrintRune(piscine.LastRune("Hello!"))
	ft.PrintRune(piscine.LastRune("Salut!"))
	ft.PrintRune(piscine.LastRune("Ola!"))
	ft.PrintRune(piscine.LastRune(""))
	ft.PrintRune('\n')
}
