package piscine

func Push[T any](s *[]T, v T) {
	*s = append(*s, v)
}

func Pop[T any](s *[]T) T {
	n := len(*s)
	v := (*s)[n-1]
	*s = (*s)[:n-1]
	return v
}

func Top[T any](s *[]T) T {
	n := len(*s)
	v := (*s)[n-1]
	return v
}

func Nth[T any](s *[]T, i int) T {
	n := len(*s)
	if 0 <= i && i < n {
		return (*s)[i]
	}
	if 0 > i && i >= -n {
		return (*s)[n+i]
	}
	panic("out of range")
}
