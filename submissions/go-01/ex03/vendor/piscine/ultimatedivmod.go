package piscine

func UltimateDivMod(a *int, b *int) {
	if a == nil || b == nil || *b == 0 {
		return
	}
	i := *a
	*a = *a / *b
	*b = i % *b
}
