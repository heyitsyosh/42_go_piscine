package piscine

func Swap(a *int, b *int) {
	if a == nil || b == nil {
		return
	}
	*a, *b = *b, *a
}
