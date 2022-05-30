package main

import (
	"os"
	"piscine"
)

func main() {
	core := piscine.NewCore()
	n := piscine.SliceLen(os.Args)
	if n != 2 {
		piscine.ShowUsage()
		return
	}
	path := os.Args[1]
	content, rOk := piscine.Read(path)
	if !rOk {
		return
	}
	blocks, _ := piscine.Divide(content)
	if !piscine.Validate(core, blocks) {
		return
	}
	core.GivenMinos = piscine.Symbolize(core, blocks)
	// fmt.Println(core)

	piscine.Solve(core)
}
