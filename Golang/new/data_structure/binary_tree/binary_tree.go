package binary_tree

import (
	"errors"
	"fmt"
)

type BinaryTree struct {
	Left, Right *BinaryTree
	val interface{}
}

func NewBinaryTree(val interface{}) *BinaryTree {
	return &BinaryTree{
		val: val,
	}
}


func preOrder(root *BinaryTree) {
	if root == nil {
		return
	}
	fmt.Printf("%v ", root.val)
	preOrder(root.Left)
	preOrder(root.Right)
}

func inOrder(root *BinaryTree) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	fmt.Printf("%v ", root.val)
	inOrder(root.Right)
}

func postOrder(root *BinaryTree) {
	if root == nil {
		return
	}
	postOrder(root.Left)
	postOrder(root.Right)
	fmt.Printf("%v ", root.val)
}


func preOrder2(root *BinaryTree) {
	if root == nil {
		return
	}
	s := NewStack(100)
	fmt.Println(s)
	isVisit := make(map[*BinaryTree]struct{})
	s.Push(root)
	for !s.IsEmpty() {

	}
}









////



type Stack struct {
	data []interface{}
	size int
	cap  int
	top  int
}

func NewStack(cap int) *Stack {
	return &Stack{
		data: make([]interface{}, cap),
		cap:  cap,
		top:  -1,
	}
}

func (s *Stack) IsFull() bool {
	return s.size == s.cap
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Push(val interface{}) {
	if s.IsFull() {
		s.cap++
		s.data = append(s.data, val)
		s.top++
		s.size++
		return
	}
	s.top++
	s.data[s.top] = val
	s.size++
	return
}

func (s *Stack) Pop() (val interface{}, err error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is nil")
	}
	val = s.data[s.top]
	s.size--
	s.top--
	return
}
