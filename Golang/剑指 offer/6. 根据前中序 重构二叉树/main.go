package main

import "fmt"

type BinaryTreeNode struct {
	Val         int
	Left, Right *BinaryTreeNode
}

func main() {
	preOrder := []int{1, 2, 4, 7, 3, 5, 6, 8}
	inOrder := []int{4, 7, 2, 1, 5, 3, 8, 6}
	res := ConstructBT(preOrder, inOrder)
	// fmt.Println(res.Val)
	inOrderPrintByRecursitive(res)
}

func ConstructBT(preOrder, inOrder []int) *BinaryTreeNode {
	if len(preOrder) != len(inOrder) || len(preOrder) == 0 {
		return nil
	}
	root := new(BinaryTreeNode)
	root.Val = preOrder[0]
	if preOrder[0] == inOrder[0] && len(preOrder) == 1 {
		return root
	}
	point := 0
	for i, v := range inOrder {
		if v == root.Val {
			point = i
			break
		}
	}
	if inOrder[point] != root.Val {
		return nil
	}
	if point > 0 {
		root.Left = ConstructBT(preOrder[1:point+1], inOrder[:point])
	}
	if len(inOrder) > point+1 {
		root.Right = ConstructBT(preOrder[point+1:], inOrder[point+1:])
	}
	return root
}

func inOrderPrintByRecursitive(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	inOrderPrintByRecursitive(root.Left)
	fmt.Printf("%d ", root.Val)
	inOrderPrintByRecursitive(root.Right)
}
