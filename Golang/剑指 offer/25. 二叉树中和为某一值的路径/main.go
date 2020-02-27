package main

import (
	"errors"
	"fmt"
)

func main() {
	inOrder := []int{5, 6, 7, 8, 9, 10, 11}
	postOrder := []int{5, 7, 6, 9, 11, 10, 8}
	res := ConstructBT(inOrder, postOrder)
	// inOrderPrint(res)
	stack := NewStack(len(inOrder))
	FindPath(res, stack, 0, 19)
}

type BinaryTreeNode struct {
	Val         int
	Left, Right *BinaryTreeNode
}

type Stack struct {
	top      int
	size     int
	capacity int
	data     []interface{}
}

func NewStack(capacity int) *Stack {
	return &Stack{
		top:      -1,
		size:     0,
		capacity: capacity,
		data:     make([]interface{}, capacity),
	}
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) IsFull() bool {
	return s.size == s.capacity
}

func (s *Stack) Push(x interface{}) error {
	if s.IsFull() {
		return errors.New("Stack is full")
	}
	s.top++
	s.data[s.top] = x
	s.size++
	return nil
}

func (s *Stack) Pop() (x interface{}, err error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	x = s.data[s.top]
	s.top--
	s.size--
	return x, nil
}

func FindPath(root *BinaryTreeNode, stack *Stack, curSum int, target int) {
	if root == nil {
		return
	}
	curSum += root.Val
	if err := stack.Push(root.Val); err != nil {
		fmt.Println(err)
		return
	}
	if curSum == target {
		fmt.Println(stack.data)
	}
	FindPath(root.Left, stack, curSum, target)
	FindPath(root.Right, stack, curSum, target)
	stack.Pop()
	return
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
