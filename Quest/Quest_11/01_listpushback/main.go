package main

import (
	"fmt"
)

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {

}
func main() {

	link := &List{}

	ListPushBack(link, "Hello")
	ListPushBack(link, "man")
	ListPushBack(link, "how are you")

	for link.Head != nil {
		fmt.Println(link.Head.Data)
		link.Head = link.Head.Next
	}
}

/*
value := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = value
		l.Tail = value
	} else {
		l.Tail.Next = value
		l.Tail = value
	}
*/
