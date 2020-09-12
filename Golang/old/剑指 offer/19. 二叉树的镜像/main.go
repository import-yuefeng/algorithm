package main

import (
	"errors"
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
	Mirror(root)
	fmt.Println()
	PreOrderPrint(root)
	Mirror_2(root)
	fmt.Println()
	PreOrderPrint(root)
}

func PreOrderPrint(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	PreOrderPrint(root.Left)
	PreOrderPrint(root.Right)
}

func Mirror(root *BinaryTreeNode) *BinaryTreeNode {
	if root == nil {
		return nil
	}
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	if root.Left != nil {
		root.Left = Mirror(root.Left)
	}
	if root.Right != nil {
		root.Right = Mirror(root.Right)
	}
	return root
}

func Mirror_2(root *BinaryTreeNode) *BinaryTreeNode {
	if root == nil {
		return root
	}
	stack := NewStack(100)
	if err := stack.Push(root); err != nil {
		fmt.Println(err)
		return nil
	}
	for !stack.IsEmpty() {
		cur, err := stack.Pop()
		if err != nil {
			fmt.Println(err)
			return nil
		}
		tmp := cur.Left
		cur.Left = cur.Right
		cur.Right = tmp
		if cur.Right != nil {
			if err := stack.Push(cur.Right); err != nil {
				fmt.Println(err)
				return nil
			}
		}
		if cur.Left != nil {
			if err := stack.Push(cur.Left); err != nil {
				fmt.Println(err)
				return nil
			}
		}
	}
	return root
}

type Stack struct {
	top      int
	data     []*BinaryTreeNode
	size     int
	capacity int
}

func NewStack(capacity int) *Stack {
	return &Stack{
		top:      -1,
		size:     0,
		capacity: capacity,
		data:     make([]*BinaryTreeNode, capacity),
	}
}

func (s *Stack) Push(x *BinaryTreeNode) error {
	if s.size == s.capacity {
		return errors.New("Stack is full")
	}
	s.top++
	s.data[s.top] = x
	s.size++
	return nil
}

func (s *Stack) Pop() (x *BinaryTreeNode, err error) {
	if s.size == 0 {
		return nil, errors.New("Stack is empty")
	}
	x = s.data[s.top]
	s.top--
	s.size--
	return x, nil
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}
