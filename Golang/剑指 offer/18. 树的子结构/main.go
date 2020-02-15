package main

import (
	"fmt"
)

type BinaryTreeNode struct {
	Left, Right *BinaryTreeNode
	Val         int
}

func main() {
	root := new(BinaryTreeNode)
	root.Val = 1
	root.Left = new(BinaryTreeNode)
	root.Left.Val = 2
	root.Right = new(BinaryTreeNode)
	root.Right.Val = 3
	root.Right.Right = new(BinaryTreeNode)
	root.Right.Right.Val = 6
	root.Left.Left = new(BinaryTreeNode)
	root.Left.Right = new(BinaryTreeNode)
	root.Left.Left.Val = 4
	root.Left.Right.Val = 5
	PreOrderPrint(root)
	subTree := new(BinaryTreeNode)
	subTree.Left, subTree.Right = new(BinaryTreeNode), new(BinaryTreeNode)
	subTree.Val, subTree.Left.Val, subTree.Right.Val = 2, 4, 5
	fmt.Println(HasSubTree(root, nil))
}

func PreOrderPrint(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	PreOrderPrint(root.Left)
	PreOrderPrint(root.Right)
}

func HasSubTree(t1, t2 *BinaryTreeNode) bool {
	result := false
	if t1 != nil && t2 != nil {
		if t1.Val == t2.Val {
			result = DoesTree1HaveTree2(t1, t2)
		}
		if !result {
			result = HasSubTree(t1.Left, t2)
		}
		if !result {
			result = HasSubTree(t1.Right, t2)
		}
	}
	return result
}

func DoesTree1HaveTree2(t1, t2 *BinaryTreeNode) bool {
	if t2 == nil {
		return true
	}
	if t1 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	return DoesTree1HaveTree2(t1.Left, t2.Left) && DoesTree1HaveTree2(t1.Right, t2.Right)
}
