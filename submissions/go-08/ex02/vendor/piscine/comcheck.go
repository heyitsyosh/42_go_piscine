package piscine

import "ft"

func printStr(s string) {
	for _, r := range s{
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func ComCheck(args []string) {
	for _, str := range args {
		if str == "42" || str == "piscine" || str == "piscine 42" {
			printStr("Alert!!!")
			return
		}
	}
}
