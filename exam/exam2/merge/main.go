package main

import "fmt"

type TreeNodeM struct {
	Left  *TreeNodeM
	Val   int
	Right *TreeNodeM
}

func main() {
	mergedTree := &TreeNodeM{}
	t1 := &TreeNodeM{Val: 7}
	t1 = insertData(t1, 5)
	t1 = insertData(t1, 10)
	t1 = insertData(t1, 3)
	t1 = insertData(t1, 6)
	t1 = insertData(t1, 10)
	t1 = insertData(t1, 13)
	printTree(t1)
	t2 := &TreeNodeM{Val: 7}
	t2 = insertData(t2, 5)
	t2 = insertData(t2, 10)
	t2 = insertData(t2, 3)
	t2 = insertData(t2, 6)
	t2 = insertData(t2, 9)
	fmt.Println()
	printTree(t2)
	fmt.Println()
	mergedTree = MergeTrees(t1, t2)
	printTree(mergedTree)

}

func printTree(node *TreeNodeM) {
	if node != nil {
		printTree(node.Left)
		fmt.Println(node.Val)
		printTree(node.Right)
	}
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

func MergeTrees(t1 *TreeNodeM, t2 *TreeNodeM) *TreeNodeM {
	//var node *TreeNodeM
	node := &TreeNodeM{}
	if t1 == nil && t2 == nil {
		return nil //?
	}
	if t1 == nil && t2 != nil {
		return t2
	}
	if t1 != nil && t2 == nil {
		return t1
	}
	node.Val = t1.Val + t2.Val
	node.Left = MergeTrees(t1.Left, t2.Left)
	node.Right = MergeTrees(t1.Right, t2.Right)
	return node
}
