package main

import (
	"errors"
	"fmt"
)

type BinaryTreeNode struct {
	val         int
	Left, Right *BinaryTreeNode
}

func main() {
	root := new(BinaryTreeNode)
	root.val = 8
	root.Left = &BinaryTreeNode{val: 6}
	root.Right = &BinaryTreeNode{val: 10}
	root.Left.Left = &BinaryTreeNode{val: 5}
	root.Left.Right = &BinaryTreeNode{val: 7}
	root.Right.Left = &BinaryTreeNode{val: 9}
	root.Right.Right = &BinaryTreeNode{val: 11}
	ZBFS(root)
}

func ZBFS(root *BinaryTreeNode) {
	res := make([][]int, 1)
	if root == nil {
		return
	}
	q := NewQueue()
	q.Push(root)
	level := 0
	for !q.IsEmpty() {
		size := q.Size()
		for size > 0 {
			size--
			t, err := q.Pop()
			if err != nil {
				fmt.Println(err)
				return
			}
			if v, ok := t.(*BinaryTreeNode); ok {
				if v == nil {
					continue
				}
				if len(res) < level+1 {
					res = append(res, []int{})
				}
				res[level] = append(res[level], v.val)
				q.Push(v.Left)
				q.Push(v.Right)
			}
		}
		level++
	}
	for i := 0; i < len(res); i++ {
		if i&1 == 0 {
			fmt.Println(res[i])
		} else {
			reverse(res[i])
			fmt.Println(res[i])
		}
	}
}

func reverse(num []int) {
	for i := 0; i < len(num)/2; i++ {
		num[i], num[len(num)-i-1] = num[len(num)-i-1], num[i]
	}
}

func BFS(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	q := NewQueue()
	q.Push(root)
	for !q.IsEmpty() {
		size := q.Size()
		for size > 0 {
			size--
			t, err := q.Pop()
			if err != nil {
				fmt.Println(err)
				return
			}
			if v, ok := t.(*BinaryTreeNode); ok {
				if v == nil {
					continue
				}
				q.Push(v.Left)
				q.Push(v.Right)
				fmt.Printf("%d ", v.val)
			}
		}
		fmt.Println()
	}
}

type Queue struct {
	front, rare int
	data        []interface{}
	capacity    int
	size        int
}

func NewQueue() *Queue {
	return &Queue{
		front:    0,
		rare:     0,
		data:     make([]interface{}, 2),
		size:     0,
		capacity: 2,
	}
}

func (q *Queue) Push(x interface{}) {
	if q.IsFull() {
		buf := NewQueue()
		buf.data = make([]interface{}, q.capacity*2)
		for i, _ := range q.data {
			buf.data[i] = q.data[q.front]
			q.front = (q.front + 1) % q.capacity
		}
		q.front = 0
		q.rare = q.size
		q.data = buf.data
		q.capacity *= 2
	}
	q.size++
	q.data[q.rare] = x
	q.rare = (q.rare + 1) % q.capacity
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

func (q *Queue) IsFull() bool {
	return q.size == q.capacity
}

func (q *Queue) Size() int {
	return q.size
}
