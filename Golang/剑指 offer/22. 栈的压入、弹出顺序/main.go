package main

import (
	"errors"
	"fmt"
)

func main() {
	// res := IsPopOrder([]int{1,2,3,4,5}, []int{4,5,3,2,1})
	res := IsPopOrder([]int{1,2,3,4,5}, []int{4,5})
	// 4, 3, 5, 1, 2 false
	// 4, 5, 3, 2, 1
	// []int{}
	// []int{1}
	fmt.Println(res)

}

func IsPopOrder(pPush, pPop []int) bool {
	if len(pPush) != len(pPop) {
		return false
	}
	if len(pPush) == 0 {
		return false
	}

	p1, p2 := 0, 0
	stack := NewStack(len(pPush))
	for p1 < len(pPush) {
		if err := stack.Push(pPush[p1]); err != nil {
			fmt.Println(err)
			return false
		}
		p1++
		topVal, err := stack.Top()
		if err != nil {
			fmt.Println(err)
			return false
		}
		if !stack.IsEmpty() && topVal == pPop[p2] {
			if _, err := stack.Pop(); err != nil {
				fmt.Println(err)
				return false
			}
			p2++
		}
	}
	topVal, err := stack.Top()
	if err != nil {
		fmt.Println(err)
		return false
	}

	if !stack.IsEmpty() && pPop[p2] != topVal {
		return false
	}
	return true
}

type Stack struct {
	top int
	size int
	capacity int
	data []int
}

func NewStack(capacity int) *Stack {
	return &Stack{
		top: -1, 
		size: 0, 
		capacity: capacity,
		data: make([]int, capacity),
	}
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Push(x int) error {
	if s.size == s.capacity {
		return errors.New("Stack is full")
	}
	s.top++
	s.data[s.top] = x
	s.size++
	return nil
}

func (s *Stack) Pop() (x int, err error) {
	if s.IsEmpty() {
		return -1, errors.New("Stack is empty")
	}
	x = s.data[s.top]
	err = nil
	s.top--
	s.size--
	return 
}

func (s *Stack) Top() (x int, err error) {
	if s.IsEmpty() {
		return -1, errors.New("Stack is empty")
	}
	x = s.data[s.top]
	return x, nil
}