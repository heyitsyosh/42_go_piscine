package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20{
		return 0
	} else if nb == 0 {
		return 1
	}
	for i := nb - 1; i > 0; i-- {
		nb *= i;
	}
	return nb
}
