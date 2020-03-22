package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Val         int
	Next, Front *Node
}

type Queue struct {
	root      *Node
	head, end *Node
}

func NewNode(val int) *Node {
	return &Node{
		Val:   val,
		Next:  nil,
		Front: nil,
	}
}

func NewQueue() *Queue {
	root := NewNode(0)
	return &Queue{
		root: root,
		head: nil,
		end:  nil,
	}
}

func (q *Queue) PushFront(val int) {
	node := NewNode(val)
	if q.head == nil {
		q.root.Next = node
		q.head = node
		node.Front = q.root
		q.end = node
	} else {
		node.Next = q.root.Next
		node.Next.Front = node
		q.root.Next = node
		node.Front = q.root
		q.head = node
	}
	return
}

func (q *Queue) PushBack(val int) {
	if q.end == nil {
		q.PushFront(val)
		return
	}
	node := NewNode(val)
	q.end.Next = node
	node.Front = q.end
	q.end = node
	return
}

func (q *Queue) PopFront() (val *Node, err error) {
	if q.head == nil && q.end == nil {
		return nil, errors.New("Queue is empty")
	}
	val = q.head
	if val.Next != nil {
		q.root.Next = val.Next
		q.head = q.head.Next
	} else {
		q.root.Next = nil
		q.head, q.end = nil, nil
	}
	return val, nil
}

func (q *Queue) PopBack() (val *Node, err error) {
	if q.head == nil {
		return nil, errors.New("Queue is empty")
	}
	if q.head == q.end {
		return q.PopFront()
	}
	val = q.end
	tmp := q.end.Front
	q.end = tmp
	val.Front = nil
	tmp.Next = nil
	return val, nil
}

func (q *Queue) IsEmpty() bool {
	return q.head == nil
}

func (q *Queue) Pop(val int) (*Node, error) {
	v, err := q.PopFront()
	if err != nil {
		return nil, err
	}
	if v.Val != val {
		q.PushFront(v.Val)
		return nil, nil
	}
	return v, nil
}

func (q *Queue) Push(val int) {
	if q.IsEmpty() {
		q.PushBack(val)
		return
	}
	for !q.IsEmpty() {
		v, err := q.PopBack()
		if err != nil {
			fmt.Println(err)
		}
		if v.Val > val {
			q.PushBack(val)
			return
		}
	}
	if q.IsEmpty() {
		q.PushBack(val)
	}
	return
}

func (q *Queue) Println() {
	cur := q.head
	for cur != nil {
		fmt.Printf("%d ", cur.Val)
		cur = cur.Next
	}
	fmt.Println()
	return
}

func main() {
	queue := NewQueue()
	test := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(test); i++ {
		if i&1 == 1 {
			queue.PushBack(test[i])
		}
	}
	for i := 0; i < len(test); i++ {
		if i&1 == 1 {
			val, err := queue.PopFront()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%d ", val.Val)
		}
	}
	fmt.Println()
	for i := 0; i < len(test); i++ {
		if i&1 == 1 {
			queue.PushBack(test[i])
		}
	}
	queue.Println()
}
