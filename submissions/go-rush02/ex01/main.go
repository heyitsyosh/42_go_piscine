package main

import (
	"flag"
	"fmt"
	"ft"
	"os"
	"piscine"
)

func main() {
	core := piscine.NewCore()

	if piscine.GFlag != 0 {
		piscine.Generate(core, piscine.GFlag)
		return
	}
	args := flag.Args()
	n := piscine.SliceLen(args)

	var paths []string
	if n == 0 {
		paths = []string{""}
	} else {
		paths = args
	}
	for i, path := range paths {
		if len(paths) > 1 {
			if i > 0 {
				ft.PrintRune('\n')
			}
			fmt.Printf("===> %s <===\n", path)
		}
		content, rOk := read(path)
		if !rOk {
			continue
		}
		blocks, _ := piscine.Divide(content)
		if !piscine.Validate(core, blocks) {
			continue
		}
		core.GivenMinos = piscine.Symbolize(core, blocks)
		piscine.Solve(core)
	}
}

func read(path string) (string, bool) {
	if path == "" {
		return piscine.ReadFromFile(os.Stdin)
	}
	return piscine.Read(path)
}
