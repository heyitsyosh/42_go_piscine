package main

import "fmt"
import "os"
import "strconv"

const ERROR_MSG string = "Arguments is invalid."

func calculationStr(args []string) (string, bool) {
	ret := ""
	if len(args) != 3 {
		return "", false
	}
	a, a_err := strconv.Atoi(args[1])
	b, b_err := strconv.Atoi(args[2])
	if a_err != nil || b_err != nil {
		return "", false
	}
	ret += "sum: " + strconv.Itoa(a+b) + "\n"
	ret += "difference: " + strconv.Itoa(a-b) + "\n"
	ret += "product: " + strconv.Itoa(a*b) + "\n"
	if b == 0 {
		return "", false
	} else {
		ret += "quotient: " + strconv.Itoa(a/b) + "\n"
	}
	return ret, true
}

func main() {
	s, ok := calculationStr(os.Args)
	if !ok {
		fmt.Println(ERROR_MSG)
		os.Exit(1)
	}
	fmt.Print(s)
}
