package piscine

import (
	"math/rand"
	"strconv"
)

func getRand(str string, min int, max int) int {
	var tmp int
	if str[0] == '?' {
		tmp, _ = strconv.Atoi(str[1:])
	} else {
		tmp, _ = strconv.Atoi(str)
	}

	if tmp <= min {
		tmp = min
	} else if tmp >= (max - min) {
		tmp = max - min
	}

	if str[0] == '?' {
		tmp = rand.Intn(tmp) + min
	}
	return tmp
}
