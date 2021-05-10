package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListReverse(l *List) {

	Premier := l.Head
	l.Tail = l.Head
	var Second *NodeL

	for Premier != nil {
		Premier.Next = Second
		Second = Premier
	}
	return l.Head
}
