package main

import (
	"fmt"
)

type NodeL struct {
	Data interface{}
	Next *NodeL
}

func ListAt(l *NodeL, pos int) *NodeL {

	compt := 0

	for l != nil {
		compt++
		if compt == pos {
			return l
		}
		l = l.Next
	}
	return l
}

func main() {
	link := &List{}

	ListPushBack(link, "hello")
	ListPushBack(link, "how are")
	ListPushBack(link, "you")
	ListPushBack(link, 1)

	fmt.Println(ListAt(link.Head, 3).Data)
	fmt.Println(ListAt(link.Head, 1).Data)
	fmt.Println(ListAt(link.Head, 7))
}
