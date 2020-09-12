// package myStack
package main

import (
	"errors"
	"fmt"
)

type myStack struct {
	stackArray []int
	topPonint  int
	size       int
	capacity   int
}

func (stack *myStack) Init(capacity int) error {
	if capacity <= 0 {
		return errors.New("capacity invaild!")
	}
	stack.capacity = capacity
	stack.size = 0
	stack.topPonint = -1
	stack.stackArray = make([]int, capacity)
	return nil
}

func (stack *myStack) Pop() (value int, err error) {
	if stack.size > 0 && stack.topPonint > -1 {
		stack.size--
		popValue := stack.stackArray[stack.topPonint]
		stack.topPonint--
		return popValue, nil
	}
	return 0, errors.New("stack is empty!")
}

func (stack *myStack) Push(value int) error {
	if stack.size >= stack.capacity {
		return errors.New("stack is full!")
	}
	stack.size++
	stack.topPonint++
	stack.stackArray[stack.topPonint] = value
	return nil
}

func (stack *myStack) Empty() bool {
	if stack.size == 0 {
		return true
	}
	return false
}

func (stack *myStack) Full() bool {
	if stack.size == stack.capacity {
		return true
	}
	return false
}

func main() {

	stack := myStack{}
	stack.Init(10)
	for i := 0; i < 10; i++ {
		stack.Push(i)
	}
	for i := 0; i < 10; i++ {
		tmp, _ := stack.Pop()
		fmt.Println(tmp)
	}

}
