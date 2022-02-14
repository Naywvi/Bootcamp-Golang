package main

import (
	"fmt"
)

type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

func pushBack(n *NodeAddL, num int) *NodeAddL {
	if n == nil {
		return nil
	}
	tmp := n
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = &NodeAddL{Num: num}
	return n
}

func pushFront(n *NodeAddL, num int) *NodeAddL {
	node := &NodeAddL{Num: num}
	node.Next = n
	return node
}

func Reverse(node *NodeAddL) *NodeAddL {
	var new *NodeAddL
	if node == nil {
		return nil
	}
	for tmp := node; tmp != nil; tmp = tmp.Next {
		new = pushFront(new, tmp.Num)
	}
	return new
}

func main() {
	num1 := &NodeAddL{Num: 1}
	num1 = pushBack(num1, 3)
	num1 = pushBack(num1, 2)
	num1 = pushBack(num1, 4)
	num1 = pushBack(num1, 5)

	result := Reverse(num1)
	for tmp := result; tmp != nil; tmp = tmp.Next {
		fmt.Print(tmp.Num)
		if tmp.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
