package main

import (
	"fmt"
	"piscine"
)

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(piscine.Join(toConcat, ":"))
	checkEmpty := []string{""}
	fmt.Println(piscine.Join(checkEmpty, ":"))
}
