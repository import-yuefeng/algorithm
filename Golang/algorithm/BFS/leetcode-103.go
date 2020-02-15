package main

import (
	"errors"
	"fmt"
)

func main() {
	var bt *BinaryTree = NewBinaryTree(3)
	bt.left = NewBinaryTree(9)
	bt.right = NewBinaryTree(20)
	bt.right.left = NewBinaryTree(15)
	bt.right.right = NewBinaryTree(7)
	bt.bfs()
}

type BinaryTree struct {
	left, right *BinaryTree
	val         int
}

func NewBinaryTree(x int) *BinaryTree {
	return &BinaryTree{
		left:  nil,
		right: nil,
		val:   x,
	}
}

func (bt *BinaryTree) bfs() {
	res := make([][]int, 0)

	q1 := NewQueue(100)
	q2 := NewQueue(100)
	q1.Push(bt)
	cur := make([]int, 0)
	flag := 0
	for q1.size != 0 || q2.size != 0 {
		if q1.size != 0 {
			tmp, err := q1.Pop()
			if err != nil {
				fmt.Println(err)
				return
			}
			if tmp == nil {
				continue
			}
			q2.Push(tmp.left)
			q2.Push(tmp.right)
			// fmt.Printf("%d ", tmp.val)
			cur = append(cur, tmp.val)
		} else { // stack.size != 0
			// fmt.Println()
			if flag == 1 {
				reverse(cur)
				flag = 0
			} else {
				flag = 1
			}
			res = append(res, cur)

			cur = make([]int, 0)
			for q2.size != 0 {
				tmp, _ := q2.Pop()
				q1.Push(tmp)
			}
		}
	}
	fmt.Println(res)
}

func reverse(arr []int) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

type Queue struct {
	front, rare int
	capacity    int
	data        []*BinaryTree
	size        int
}

func NewQueue(capacity int) *Queue {
	return &Queue{
		front:    0,
		rare:     0,
		data:     make([]*BinaryTree, capacity),
		capacity: capacity,
		size:     0,
	}
}

func (q *Queue) Push(x *BinaryTree) error {
	if q.size == q.capacity {
		return errors.New("queue is full")
	}
	q.data[q.rare] = x
	q.rare = (q.rare + 1) % q.capacity
	q.size++
	return nil
}

func (q *Queue) Pop() (x *BinaryTree, err error) {
	if q.size == 0 {
		return nil, errors.New("queue is empty")
	}
	x = q.data[q.front]
	err = nil
	q.front = (q.front + 1) % q.capacity
	q.size--
	return
}
