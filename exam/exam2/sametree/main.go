package main

import "fmt"

type TreeNodeM struct {
	Left  *TreeNodeM
	Val   int
	Right *TreeNodeM
}

func main() {
	t1 := &TreeNodeM{Val: 7}
	t1 = insertData(t1, 5)
	t1 = insertData(t1, 10)
	t1 = insertData(t1, 3)
	t1 = insertData(t1, 6)
	t1 = insertData(t1, 9)

	t2 := &TreeNodeM{Val: 7}
	t2 = insertData(t2, 5)
	t2 = insertData(t2, 10)
	t2 = insertData(t2, 3)
	t2 = insertData(t2, 6)
	t2 = insertData(t2, 9)
	fmt.Println(IsSameTree(t1, t2))

}

func insertData(node *TreeNodeM, val int) *TreeNodeM {
	if node == nil {
		return &TreeNodeM{Val: val}
	}
	if node.Val > val {
		node.Left = insertData(node.Left, val)
	}
	if node.Val <= val {
		node.Right = insertData(node.Right, val)
	}
	return node
}

func IsSameTree(t1 *TreeNodeM, t2 *TreeNodeM) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil && t2 != nil {
		return false
	}
	if t1 != nil && t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	if IsSameTree(t1.Left, t2.Left) && IsSameTree(t1.Right, t2.Right) {
		return true
	}
	return false
}
