package main

import (
	"piscine"
	"ft"
)

func main() {
	result := piscine.Rot14("Hello! How are You?")
	
	for _, r := range result {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}