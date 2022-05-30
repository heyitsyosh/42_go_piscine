package main

import (
	"fmt"
	"piscine"
)

func main() {
	link := &piscine.List{}
	
	piscine.ListPushFront(link, "Hello")
	piscine.ListPushFront(link, "there")
	piscine.ListPushFront(link, "how are you")

	it := link.Head
	for it != nil {
		fmt.Print(it.Data, " ")
		it = it.Next
	}
	fmt.Println()
}