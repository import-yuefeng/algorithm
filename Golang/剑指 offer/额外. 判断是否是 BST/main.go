package main

import (
	"errors"
	"fmt"
	"math"
)

type BST struct {
	Val         int
	Left, Right *BST
}

//         3
//       /   \
//     2     5
//   /  \
// 1    4
// return false
//         3
//       /   \
//     2     5
//   /
// 1
// return true

func main() {
	// var prev *BST
	root := new(BST)
	root.Val = 3
	root.Left = &BST{Val: 2, Left: &BST{Val: 1}, Right: &BST{Val: 4}}
	root.Right = &BST{Val: 5}
	res := isBST3(root)
	fmt.Println(res)
	// inOrderPrint(root)
}

func isBST2(root *BST) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && findMax(root.Left) > root.Val {
		return false
	}
	if root.Right != nil && findMin(root.Right) < root.Val {
		return false
	}
	return isBST2(root.Left) && isBST2(root.Right)
}

func findMax(cur *BST) int {
	if cur == nil {
		return -1 << 31
	}
	l := findMax(cur.Left)
	r := findMax(cur.Right)
	return Max(Max(l, r), cur.Val)
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func findMin(cur *BST) int {
	if cur == nil {
		return 1 << 31
	}
	l := findMin(cur.Left)
	r := findMin(cur.Right)
	return Min(Min(l, r), cur.Val)
}

func Min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

var prev int = math.MinInt64

func isBST1(root *BST) bool {
	if root == nil {
		return true
	}
	if !isBST1(root.Left) {
		return false
	}
	if root.Val < prev {
		return false
	}
	prev = root.Val
	return isBST1(root.Right)
}

func isBSTByMorris(root *BST) bool {
	if root == nil {
		return true
	}
	cur := root
	for cur != nil {

	}
}
\
func isBST3(root *BST) bool {
	stack := NewStack()
	isV := make(map[*BST]struct{})
	stack.Push(root)
	prevVal := -1 << 31
	for !stack.IsEmpty() {
		t, err := stack.Pop()
		if err != nil {
			fmt.Println(err)
			return false
		}
		if v, ok := t.(*BST); !ok {
			fmt.Println(ok)
			return false
		} else {
			if v == nil {
				continue
			}
			if _, ok := isV[v]; ok {
				if v.Val < prevVal {
					return false
				}
				prevVal = v.Val
			} else {
				stack.Push(v.Right)
				isV[v] = struct{}{}
				stack.Push(v)
				stack.Push(v.Left)
			}
		}
	}
	return true
}

func inOrderPrint(root *BST) {
	if root == nil {
		return
	}
	inOrderPrint(root.Left)
	fmt.Printf("%d ", root.Val)
	inOrderPrint(root.Right)
}

type Stack struct {
	top      int
	size     int
	capacity int
	data     []interface{}
}

func NewStack() *Stack {
	// default capacity: 2
	return &Stack{
		top:      -1,
		size:     0,
		capacity: 2,
		data:     make([]interface{}, 2),
	}
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Push(x interface{}) {
	if s.size == s.capacity {
		buf := NewStack()
		buf.data = make([]interface{}, 2*s.capacity)
		for i, v := range s.data {
			buf.data[i] = v
		}
		s.capacity *= 2
		s.data = buf.data
	}
	s.top++
	s.data[s.top] = x
	s.size++
	return
}

func (s *Stack) Pop() (x interface{}, err error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is empty")
	}
	x = s.data[s.top]
	s.size--
	s.top--
	return x, nil
}
