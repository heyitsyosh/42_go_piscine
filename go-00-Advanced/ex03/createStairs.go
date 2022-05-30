package main

import "fmt"
import "os"
import "strconv"

func main() {
	n := os.Args
	if len(n) != 2 {
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 0 || num > 10000 {
		return
	}
	for remainingStar, starNum := num, 1; 0 <= remainingStar-starNum; remainingStar, starNum = remainingStar-starNum, starNum+1 {
		for i := starNum; i > 0; i-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
	return
}
