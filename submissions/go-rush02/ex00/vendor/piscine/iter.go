package piscine

func ForEach[T any](s []T, f func(a T, i int)) {
	for i, a := range s {
		f(a, i)
	}
}

func Map[T any, S any](s []T, f func(a T, i int) S) []S {
	rv := make([]S, 0, 0)
	for i, a := range s {
		rv = append(rv, f(a, i))
	}
	return rv
}

func Reduce[T any, S any](s []T, init S, f func(s S, a T, i int) S) S {
	for i, a := range s {
		init = f(init, a, i)
	}
	return init
}

// if a element of given slice `s` s.t. returns false by given predicate `f`, returns false.
// otherwise, returns true.
func Every[T any](s []T, f func(a T, i int) bool) bool {
	for i, a := range s {
		if !f(a, i) {
			return false
		}
	}
	return true
}

func Some[T any](s []T, f func(a T, i int) bool) bool {
	for i, a := range s {
		if f(a, i) {
			return true
		}
	}
	return false
}

func Filter[T any](s []T, f func(a T, i int) bool) []T {
	rv := make([]T, 0, 0)
	for i, a := range s {
		if f(a, i) {
			rv = append(rv, a)
		}
	}
	return rv
}

func Index[T any](s []T, f func(a T, i int) bool) int {
	for i, a := range s {
		if f(a, i) {
			return i
		}
	}
	return -1
}

func Seq(from, to int) []int {
	var rv []int
	if from == to {
		return rv
	}
	if from < to {
		rv = make([]int, 0, to-from)
		for ; from != to; from++ {
			rv = append(rv, from)
		}
	} else {
		rv = make([]int, 0, from-to)
		from = from - 1
		for {
			rv = append(rv, from)
			from--
			if from < to {
				break
			}
		}
	}
	return rv
}

func Reverse[T any](s []T) []T {
	n := len(s)
	rv := make([]T, n, n)
	for i, _ := range s {
		rv[n-i-1] = s[i]
	}
	return rv
}
