package main

import (
	"ft"
	"piscine"
)

func main() {
	ft.PrintRune(piscine.NRune("Hello!", 3))
	ft.PrintRune(piscine.NRune("Salut!", 2))
	ft.PrintRune(piscine.NRune("Bye!", -1))	
	ft.PrintRune(piscine.NRune("あいうえお", 3))
	ft.PrintRune(piscine.NRune("Ola!", 4))
	ft.PrintRune(piscine.NRune("", 4))
	ft.PrintRune('\n')
}
