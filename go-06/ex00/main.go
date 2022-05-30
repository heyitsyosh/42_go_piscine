package main

import "ft"
import "os"

type boolean int

const EvenMsg = "I have an even number of arguments"
const OddMsg = "I have an odd number of arguments"
const yes = 1
const no = 0
var lengthOfArg = argCount(os.Args)

func argCount(args []string) (count int) {
	for range args {
		count++
	}
	if count > 0 {
		return count - 1
	}
	return 0
}

func even(nbr int) (ret int) {
	if nbr % 2 == 0 {
		return 1
	}
	return 0
}

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func isEven(nbr int) boolean {
	if even(nbr) == 1 {
		return yes
	} else {
		return no
	}
}

func main() {
	if isEven(lengthOfArg) == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
