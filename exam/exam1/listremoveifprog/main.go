package main

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = n
		l.Tail = n
	}
	l.Tail.Next = n
	l.Tail = n
}

func ListRemoveIf(l *List, data_ref interface{}) {
	listNew := &List{}
	tmp := l.Head
	for tmp != nil {
		if tmp.Data == data_ref {
			continue
		}
		ListPushBack(listNew, tmp.Data)
		l.Head = l.Head.Next
	}
	*l = *listNew
}
