package piscine

func DivMod(a int, b int, div *int, mod *int) {
	if div == nil || mod == nil || b == 0 {
		return
	}
	*div = a / b
	*mod = a % b
}
