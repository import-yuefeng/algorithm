// package stackQueue
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

type myQueue struct {
	a *myStack
	b *myStack
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

	queue := myQueue{}
	queue.queueInit(10)
	for i := 0; i < 10; i++ {
		queue.queuePush(i)
	}
	for i := 0; i < 10; i++ {
		tmp := queue.queuePop()
		fmt.Println(tmp)
	}

}

func (queue *myQueue) queueInit(capacity int) error {
	if capacity <= 0 {
		return errors.New("capacity invaild!")
	}
	queue.a = new(myStack)
	queue.b = new(myStack)
	queue.a.Init(capacity)
	queue.b.Init(capacity)
	return nil
}

func (queue *myQueue) queuePush(value int) error {
	if isFull := queue.a.Full(); isFull {
		return errors.New("queue is full")
	}
	queue.a.Push(value)
	return nil
}

func (queue *myQueue) queuePop() (value int) {
	if isEmpty := queue.b.Empty(); isEmpty {
		for isEmpty := queue.a.Empty(); !isEmpty; isEmpty = queue.a.Empty() {
			value, _ := queue.a.Pop()
			queue.b.Push(value)
		}
	}
	result, _ := queue.b.Pop()
	return result
}
