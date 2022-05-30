package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
	}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	if l == nil {
		return
	}
	n := &NodeL{data, nil}

	if l.Head == nil && l.Tail == nil {
		l.Head = n
		l.Tail = n
		return
	} else if l.Head != nil && l.Head == nil {
		
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
}
