package main

import (
	"errors"
	"fmt"
)

type MinStack struct {
	min      []int
	data     []int
	top      int
	size     int
	capacity int
	minVal   int
}

func NewMinStack(capacity int) *MinStack {
	return &MinStack{
		min:      make([]int, capacity),
		data:     make([]int, capacity),
		top:      -1,
		size:     0,
		capacity: capacity,
		minVal:   1 << 31,
	}
}

func (m *MinStack) Push(x int) error {
	if m.size == m.capacity {
		return errors.New("MinStack is empty")
	}
	m.top++

	if m.top == 0 {
		m.min[m.top] = x
		m.minVal = x
	} else if m.minVal < x {
		m.min[m.top] = m.minVal
	} else {
		m.min[m.top] = x
		m.minVal = x
	}
	m.data[m.top] = x
	m.size++
	return nil
}

func (m *MinStack) Pop() (x int, err error) {
	if m.size == 0 {
		return -1, errors.New("MinStack is empty")
	}
	x = m.data[m.top]
	if m.top-1 > 0 {
		m.minVal = m.min[m.top-1]
	}
	m.top--
	m.size--
	return x, nil
}

func (m *MinStack) Min() int {
	if m.size == 0 {
		return 1 << 31
	}
	return m.min[m.top]
}

func main() {
	stack := NewMinStack(100)
	stack.Push(3)
	stack.Push(4)
	stack.Push(2)
	stack.Push(1)
	stack.Pop()
	stack.Pop()
	stack.Push(0)
	fmt.Println(stack.data, stack.min)
}
