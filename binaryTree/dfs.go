package main

import "errors"

type stack struct {
	top      int
	size     int
	capacity int
	data     []int
}

func (s *stack) Init(capacity int) error {
	if capacity <= 0 {
		return errors.New("capacity is invaild!")
	}
	s.capacity = capacity
	s.top, s.size = -1, 0
	s.data = make([]int, s.capacity)
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

type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}

func Init(val int) (bt *BinaryTree, err error) {
	bt = new(BinaryTree)
	bt.Left, bt.Right = nil, nil
	bt.Val = val
	return bt, nil
}

func Insert(value int) error {

}

func (bt *BinaryTree) DFS() {

}
