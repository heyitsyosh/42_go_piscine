package main

import "fmt"

func subSlice(slice []int, begin int, length int, capacity int) []int {
	if begin < 0 || length < 0 || capacity < 0 {
		var invalid []int
		return invalid
	}
	if capacity < length {
		capacity = length
	}
	sliceLen := len(slice)
	ret := make([]int, length, capacity)
	for i := range ret {
		if i > sliceLen-1-begin {
			ret[i] = 0
			continue
		}
		ret[i] = slice[begin+i]
	}
	return ret
}

func main() {
	var orig = []int{0, 1, 2, 1}
	var ret []int

	ret = subSlice(orig, 0, 4, 4)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))

	ret = subSlice(orig, 2, 7, 10)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))

	ret = subSlice(orig, 2, 7, 3)
	fmt.Printf("ret = %v, len = %d, cap = %d\n", ret, len(ret), cap(ret))
}
