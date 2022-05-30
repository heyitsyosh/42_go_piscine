package piscine

import "os"

const INT_MAX = int(^uint(0) >> 1)

func PrintError(s string) {
	errMsg := []byte("Map Error: " + s + "\n")

	_, err := os.Stderr.Write(errMsg)
	if err != nil {
		os.Exit(1)
	}
}

func Atoi(nbr string) (int, bool) {
	ret := 0
	for _, digit := range []rune(nbr) {
		if digit >= '0' && digit <= '9' {
			if ret > (INT_MAX - int(digit - '0'))/10 {
				return 0, false
			}
			ret = ret*10 + int(digit-'0')
		} else {
			return 0, false
		}
	}
	if ret == 0 {
		return 0, false
	}
	return ret, true
}
