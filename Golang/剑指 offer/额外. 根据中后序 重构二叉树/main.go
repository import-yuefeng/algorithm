package main

import "fmt"

type BinaryTreeNode struct {
	Val         int
	Left, Right *BinaryTreeNode
}

func main() {
	inOrder := []int{5, 6, 7, 8, 9, 10, 11}
	postOrder := []int{5, 7, 6, 9, 11, 10, 8}
	res := ConstructBT(inOrder, postOrder)
	inOrderPrint(res)
}

func ConstructBT(inOrder, postOrder []int) *BinaryTreeNode {
	if len(inOrder) != len(postOrder) || len(postOrder) == 0 {
		return nil
	}
	root := new(BinaryTreeNode)
	root.Val = postOrder[len(postOrder)-1]
	if len(inOrder) == 1 && inOrder[0] == postOrder[0] {
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
		root.Left = ConstructBT(inOrder[:point], postOrder[:point])
	}
	if len(postOrder) > point+1 {
		root.Right = ConstructBT(inOrder[point+1:], postOrder[point:len(postOrder)-1])
	}
	return root
}

func inOrderPrint(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	inOrderPrint(root.Left)
	fmt.Printf("%d ", root.Val)
	inOrderPrint(root.Right)
}
