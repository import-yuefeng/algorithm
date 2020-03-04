package main

import (
	"errors"
	"fmt"
)

type BinaryTreeNode struct {
	Val         int
	Left, Right *BinaryTreeNode
}

func main() {
	root := new(BinaryTreeNode)
	root.Val = 1
	root.Left = &BinaryTreeNode{Val: 2, Left: &BinaryTreeNode{Val: 4}, Right: &BinaryTreeNode{Val: 5}}
	root.Right = &BinaryTreeNode{Val: 3, Left: &BinaryTreeNode{Val: 6}, Right: &BinaryTreeNode{Val: 7}}
	printBT(root)
	printZBT(root)
}

func printBT(root *BinaryTreeNode) {
	queue := NewQueue()
	queue.Push(root)
	for !queue.IsEmpty() {
		size := queue.Size()
		for size > 0 {
			tmp, err := queue.Pop()
			if err != nil {
				fmt.Println(err)
				return
			}
			size--
			if val, ok := tmp.(*BinaryTreeNode); !ok {
				fmt.Println(ok)
				return
			} else {
				if val == nil {
					continue
				}
				fmt.Printf("%d ", val.Val)
				queue.Push(val.Left)
				queue.Push(val.Right)
			}
		}
		fmt.Println()
	}
}

type Queue struct {
	rare, front int
	data        []interface{}
	capacity    int
	size        int
}

func NewQueue() *Queue {
	return &Queue{
		rare:     0,
		front:    0,
		data:     make([]interface{}, 2),
		capacity: 2,
		size:     0,
	}
}

func (q *Queue) Push(x interface{}) {
	if q.size == q.capacity {
		buf := NewQueue()
		buf.data = make([]interface{}, q.capacity*2)
		for i := 0; i < q.size; i++ {
			buf.data[i] = q.data[q.front]
			q.front = (q.front + 1) % q.capacity
		}
		q.capacity *= 2
		q.front = 0
		q.rare = q.size
		q.data = buf.data
	}
	q.data[q.rare] = x
	q.rare = (q.rare + 1) % q.capacity
	q.size++
	return
}

func (q *Queue) Pop() (x interface{}, err error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}
	x = q.data[q.front]
	q.front = (q.front + 1) % q.capacity
	q.size--
	return x, nil
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Size() int {
	return q.size
}

func printZBT(root *BinaryTreeNode) {
	flag := 1
	stack1, stack2 := NewStack(), NewStack()
	stack2.Push(root)
	for !stack1.IsEmpty() || !stack2.IsEmpty() {
		var cur *Stack
		if flag == 1 {
			cur = stack2
		} else {
			cur = stack1
		}
		size := cur.Size()
		for size > 0 {
			t, err := cur.Pop()
			size--
			if err != nil {
				fmt.Println(err)
				return
			}
			if v, ok := t.(*BinaryTreeNode); !ok {
				fmt.Println(ok)
				return
			} else {
				if v == nil {
					continue
				}
				fmt.Printf("%d ", v.Val)
				if flag == 1 {
					stack1.Push(v.Left)
					stack1.Push(v.Right)
				} else {
					stack2.Push(v.Right)
					stack2.Push(v.Left)
				}
			}
		}
		fmt.Println()
		flag *= -1
	}
}

type Stack struct {
	top      int
	capacity int
	data     []interface{}
	size     int
}

func NewStack() *Stack {
	// default capacity: 2
	return &Stack{
		top:      -1,
		capacity: 2,
		data:     make([]interface{}, 2),
		size:     0,
	}
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Push(x interface{}) {
	if s.size == s.capacity {
		buf := NewStack()
		buf.data = make([]interface{}, s.capacity*2)
		for i := 0; i < s.size; i++ {
			buf.data[i] = s.data[i]
		}
		s.data = buf.data
		s.capacity *= 2
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
	s.top--
	s.size--
	return x, nil
}

func (s *Stack) Size() int {
	return s.size
}
