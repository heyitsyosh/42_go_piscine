package piscine

import "fmt"

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = -INT_MAX - 1

func Doop(args []string, a, b int) {
	var result int
	switch args[2] {
	case "+":
		if (a < 0 && b < INT_MIN - a) || (a >= 0 && b > INT_MAX - a) {
			return
		} 
		result = a + b
	case "-":
		if (a < 0 && b > a - INT_MIN) || (a >= 0 && b < a - INT_MAX) {
			return
		}
		result = a - b
	case "/":
		if b == 0 {
			fmt.Println("No division by zero")
			return
		} else if a == INT_MIN && b == -1 {
			return
		}
		result = a / b
	case "%":
		if b == 0 {
			fmt.Println("No modulo by zero")
			return
		}
		result = a % b
	case "*":
		if (a != 0 && a*b/a != b) {
			return
		}
		result = a * b
	}
	fmt.Println(result)
}

//validate input, a to i
func Atoi(s string) (int, bool) {
	res := 0
	sign := 1
	len := 0

	for i, num := range s {
		if i == 0 && (num == '+' || num == '-') {
			if num == '-' {
				sign *= -1
			}
		} else if num < '0' || num > '9' {
			return 0, false
		} else {
			if (sign == 1 && res > (INT_MAX - int(num - '0'))/10) || (sign == -1 && -res < (INT_MIN + int(num - '0'))/10) {
				return 0, false
			}
			res = res*10 + int(num - '0')
		}
		len++
	}
	if len == 0 {
		return 0, false
	}
	return res * sign, true
}
