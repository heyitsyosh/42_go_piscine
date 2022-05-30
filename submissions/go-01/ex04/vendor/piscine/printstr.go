package piscine

import "ft"

func PrintStr(s string) {
	for _, c := range s {
		ft.PrintRune(c)
	} 
}
