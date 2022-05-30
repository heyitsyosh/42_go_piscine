package piscine

func StrLen(s string) int {
	len := 0
	for range s {
		len += 1
	}
	return len
}

func Swap(front *rune, back *rune) {
	if front == nil || back == nil {
		return
	}
	*front, *back = *back, *front
}

func StrRev(s string) string {
	len := StrLen(s)
	back := len - 1
	r := []rune(s)
	for front := 0; front < len/2; front, back = front + 1, back - 1 {
		Swap(&r[front], &r[back])
	}
	return string(r)
}
