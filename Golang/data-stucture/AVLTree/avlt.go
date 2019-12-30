package main

import (
	"errors"
)

type AVLTNode struct {
	Val   int
	Left  *AVLTNode
	Right *AVLTNode
}

func main() {
}

func Init(val int) (bt *AVLTNode, err error) {
	bt = new(AVLTNode)
	bt.Left, bt.Right = nil, nil
	bt.Val = val
	return bt, nil
}

type stack struct {
	top      int
	size     int
	capacity int
	data     []interface{}
}

func (s *stack) Init(capacity int) error {
	if capacity <= 0 {
		return errors.New("capacity is invaild!")
	}
	s.capacity = capacity
	s.top, s.size = -1, 0
	s.data = make([]interface{}, s.capacity)
	return nil
}

func (s *stack) Push(value interface{}) error {
	if s.size == s.capacity {
		return errors.New("stack is full!")
	}
	s.top++
	s.data[s.top] = value
	s.size++
	return nil
}

func (s *stack) Pop() (value interface{}, err error) {
	if s.size == 0 {
		return nil, errors.New("stack is empty!")
	}
	value = s.data[s.top]
	err = nil
	s.top--
	s.size--
	return
}

type Queue struct {
	data     []interface{}
	size     int
	rare     int
	front    int
	capacity int
}

func (q *Queue) Init(capacity int) error {
	if capacity <= 0 {
		return errors.New("capacity is invaild!")
	}
	q.capacity = capacity
	q.data = make([]interface{}, q.capacity)
	q.front, q.size, q.rare = 0, 0, 0
	return nil
}

func (q *Queue) Push(val interface{}) error {
	if q.size == q.capacity {
		return errors.New("queue is full")
	}
	q.size++
	q.rare = (q.rare + 1) % q.capacity
	q.data[q.rare] = val
	return nil
}

func (q *Queue) Pop() (val interface{}, err error) {
	if q.size == 0 {
		return nil, errors.New("queue is empty")
	}
	q.front = (q.front + 1) % q.capacity
	val = q.data[q.front]
	err = nil
	q.size--
	return
}
