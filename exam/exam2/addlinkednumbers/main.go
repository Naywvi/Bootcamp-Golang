package main

import (
	"fmt"
)

type Node struct {
	Num  int
	Next *Node
}

func len(str string) int {
	var counter int
	for range str {
		counter++
	}
	return counter
}

func pushFront(node *Node, num int) *Node {
	newnode := &Node{Num: num}
	if node == nil {
		return newnode
	}
	iterator := node
	newnode.Next = iterator
	return newnode
}

func AddLinkedNumbers(num1, num2 *Node) *Node {
	iterator1 := num1
	iterator2 := num2
	var nb1 int
	var nb2 int
	var sumnb int
	//var sumnbstr string
	for iterator1.Next != nil {
		nb1 = nb1*10 + iterator1.Num
		iterator1 = iterator1.Next
	}
	nb1 = nb1*10 + iterator1.Num
	for iterator2.Next != nil {
		nb2 = nb2*10 + iterator2.Num
		iterator2 = iterator2.Next
	}
	nb2 = nb2*10 + iterator2.Num
	sumnb = nb1 + nb2
	fmt.Println("nb1", nb1)
	fmt.Println("nb2", nb2)
	fmt.Println("sumnb", sumnb)
	var resnum *Node
	for i := sumnb; i > 0; i /= 10 {
		resnum = pushFront(resnum, i%10)
	}
	//fmt.Println("len", len(sumnbstr))
	return resnum
}

func main() {
	// 3 -> 1 -> 5
	num1 := &Node{Num: 5}
	num1 = pushFront(num1, 1)
	num1 = pushFront(num1, 3)

	// 5 -> 9 -> 2
	num2 := &Node{Num: 2}
	num2 = pushFront(num2, 9)
	num2 = pushFront(num2, 5)

	// 9 -> 0 -> 7
	result := AddLinkedNumbers(num1, num2)
	for tmp := result; tmp != nil; tmp = tmp.Next {
		fmt.Print(tmp.Num)
		if tmp.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
