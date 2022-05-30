package piscine

func ListClear(l *List) {
	if l == nil || l.Head == nil {
		return
	}
	l.head = nil
}