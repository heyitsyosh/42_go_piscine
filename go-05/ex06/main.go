package main

import (
	"fmt"
	"piscine"
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
	s = "一二三四HA語録七HA八級"
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
	s = "HA"
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
	s = "HAHA"
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
	s = "夜HA夜HA夜"
	fmt.Printf("%#v\n", piscine.Split(s, "夜"))
	s = "HAaHA"
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
	s = "HAHAaHA"
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
	s = "HAHAaHHA"
	fmt.Printf("%#v\n", piscine.Split(s, "H"))
	s = ""
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
	s = "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", piscine.Split(s, ""))
}
