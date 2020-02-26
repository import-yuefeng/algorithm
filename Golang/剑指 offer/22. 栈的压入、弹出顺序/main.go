package main

import (
	"errors"
	"fmt"
)

func main() {
	res := IsInOrder([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1})

	// res := IsPopOrder([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2})
	// res := IsPopOrder([]int{1, 2, 3, 4, 5}, []int{4, 5})
	// 4, 3, 5, 1, 2 false
	// 4, 5, 3, 2, 1
	// 5, 4, 3, 2, 1
	// []int{}
	// []int{1}
	fmt.Println(res)

}

func IsPopOrder(push, pop []int) bool {
	if len(push) == 0 || len(pop) == 0 || len(push) != len(pop) {
		return false
	}
	p1, p2 := 0, 0
	stack := NewStack(len(push))
	for p2 < len(pop) {
		for stack.IsEmpty() || stack.Top() != pop[p2] {
			if p1 >= len(push) {
				break
			}
			stack.Push(push[p1])
			p1++
		}
		if stack.Top() != pop[p2] {
			break
		}
		stack.Pop()
		p2++
	}
	if p2 == len(pop) && stack.IsEmpty() {
		return true
	}
	return false
}

type Stack struct {
	top      int
	size     int
	capacity int
	data     []int
}

func NewStack(capacity int) *Stack {
	return &Stack{
		top:      -1,
		size:     0,
		capacity: capacity,
		data:     make([]int, capacity),
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

func (s *Stack) Top() (x int) {
	if s.IsEmpty() {
		return -1
	}
	x = s.data[s.top]
	return x
}

func IsInOrder(push, pop []int) bool {
	if len(push) != len(pop) || len(push) == 0 || len(pop) == 0 {
		return false
	}
	p1, p2 := 0, 0
	stack := NewStack(len(push))
	for p2 < len(pop) {
		for stack.IsEmpty() || stack.Top() != pop[p2] {
			if p1 >= len(push) {
				break
			}
			stack.Push(push[p1])
			p1++
		}
		if stack.Top() != pop[p2] && p1 >= len(push) {
			break
		}
		stack.Pop()
		p2++
	}
	if stack.IsEmpty() && p2 >= len(pop) {
		return true
	}
	return false
}
