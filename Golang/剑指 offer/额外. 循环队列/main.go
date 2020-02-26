package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	front, rare int
	size        int
	data        []interface{}
	capacity    int
}

func NewQueue() *Queue {
	// default size: 2
	return &Queue{
		front:    0,
		rare:     0,
		size:     0,
		data:     make([]interface{}, 2),
		capacity: 2,
	}
}

func (q *Queue) Push(x interface{}) {
	if q.size == q.capacity {
		buf := NewQueue()
		buf.data = make([]interface{}, q.size*2)
		for i, _ := range q.data {
			buf.data[i] = q.data[q.front]
			q.front = (q.front + 1) % q.capacity
		}
		// 重新初始化队头接队尾
		q.rare = len(q.data)
		q.front = 0
		q.data = buf.data
		q.capacity *= 2
	}
	q.data[q.rare] = x
	q.rare = (q.rare + 1) % q.capacity
	q.size++
	return
}

func (q *Queue) Pop() (x interface{}, err error) {
	if q.size == 0 {
		return -1, errors.New("Queue is empty")
	}
	x = q.data[q.front]
	q.front = (q.front + 1) % q.capacity
	q.size--
	return x, nil
}

func (q *Queue) IsEmpty() bool {
	return q.front == q.rare
}

func (q *Queue) IsFull() bool {
	return (q.rare+1)%q.capacity == q.front
}

func main() {
	// queue := NewQueue()
	queue2 := NewQueueByLinkedlist()
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range list {
		// queue.Push(v)
		queue2.Push(v)
	}
	for !queue2.IsEmpty() {
		val, err := queue2.Pop()
		if err != nil {
			fmt.Println(err)
			return
		}
		if v, ok := val.(int); ok {
			fmt.Printf("%d ", v)
		}
	}
}

type Node struct {
	Val  interface{}
	Next *Node
}

type QueueByLinkedlist struct {
	size       int
	head, tail *Node
}

func NewQueueByLinkedlist() *QueueByLinkedlist {
	return &QueueByLinkedlist{
		size: 0,
		head: nil,
		tail: nil,
	}
}

func (q *QueueByLinkedlist) IsEmpty() bool {
	return q.size == 0
}

func (q *QueueByLinkedlist) Push(x interface{}) {
	node := new(Node)
	node.Val = x
	q.size++
	if q.size == 1 {
		q.head = node
		q.tail = node
	} else {
		q.tail.Next = node
		q.tail = node
	}
	return
}

func (q *QueueByLinkedlist) Pop() (x interface{}, err error) {
	x = q.head.Val
	if q.size == 1 {
		q.head = nil
		q.tail = nil
	} else {
		q.head.Val = q.head.Next.Val
		q.head.Next = q.head.Next.Next
	}
	q.size--
	return x, nil
}
