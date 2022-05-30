package main

import (
	"fmt"
	"os"
	"regexp"
)

func valid(email string) bool {
	regex, err := regexp.Compile("[\\w\\-\\._]+@[\\w\\-\\._]+\\.[A-Za-z]+")
	if err != nil {
		_, ws_err := os.Stderr.WriteString(err.Error())
		if ws_err != nil {
			os.Exit(1)
		}
		return false
	}
	if len(email) > 256 {
		return false
	}
	return regex.MatchString(email)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Invalid argument\n")
	} else {
		for _, mail := range os.Args[1:] {
			if valid(mail) == true {
				fmt.Printf("%s is a valid email address.\n", mail)
			} else {
				fmt.Printf("%s is NOT a valid email address.\n", mail)
			}
		}
	}
}
