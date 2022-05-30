package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Index("Hello!", "l")) //2
	fmt.Println(piscine.Index("Salut!", "alu")) //1
	fmt.Println(piscine.Index("Ola!", "hOl")) //-1
	fmt.Println(piscine.Index("", "0")) //-1
	fmt.Println(piscine.Index("", "")) //0
	fmt.Println(piscine.Index("0", "")) //0
	fmt.Println(piscine.Index("あいうえお", "うえお")) //2
	fmt.Println(piscine.Index("Hello!", "lo")) //3
	fmt.Println(piscine.Index("あいうえうえお", "うえお")) //4
}
