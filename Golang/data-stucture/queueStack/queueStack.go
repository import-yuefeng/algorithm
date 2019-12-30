// package queueStack
package main

import (
	"errors"
	"fmt"
)

type myQueue struct {
	size     int
	capacity int
	queue    []int
	front    int
	rear     int
}

type myStack struct {
	a *myQueue
	b *myQueue
}

func (queue *myQueue) Init(capacity int) error {
	if capacity <= 0 {
		return errors.New("capacity invaild!")
	}
	queue.capacity = capacity
	queue.queue = make([]int, capacity)
	queue.size = 0
	queue.front, queue.rear = 0, 0
	return nil
}

func (queue *myQueue) Push(value int) error {
	if queue.size >= queue.capacity {
		// queue.Full()
		return errors.New("queue is full")
	}
	queue.rear = (queue.rear + 1) % queue.capacity
	queue.queue[queue.rear] = value
	queue.size++
	return nil
}

func (queue *myQueue) Pop() (value int, err error) {
	if queue.size <= 0 {
		// queue.Empty()
		return 0, errors.New("queue is empty!")
	}
	queue.front = (queue.front + 1) % queue.capacity
	tmpValue := queue.queue[queue.front]
	queue.size--
	return tmpValue, nil
}

func (queue *myQueue) Empty() bool {
	if queue.front == queue.rear {
		// if queue.size == 0
		return true
	}
	return false
}

func (queue *myQueue) Full() bool {
	if queue.front == ((queue.rear + 1) % queue.capacity) {
		// queue.size == queue.capacity
		return true
	}
	return false
}

func (stack *myStack) Init(capacity int) error {
	if capacity <= 0 {
		return errors.New("capacity is invaild")
	}
	stack.a = new(myQueue)
	stack.b = new(myQueue)
	stack.a.Init(capacity)
	stack.b.Init(capacity)
	return nil
}

func (stack *myStack) Push(value int) error {
	if stack.a.Empty() || stack.b.Empty() {
		if stack.a.Empty() && !stack.b.Full() {
			stack.b.Push(value)
		} else if stack.b.Empty() && !stack.a.Full() {
			stack.a.Push(value)
		} else {
			return errors.New("stack is full!")
		}
		return nil
	}
	return errors.New("stack invaild!")
}

func (stack *myStack) Pop() (value int, err error) {
	value = 0
	if stack.a.Empty() || stack.b.Empty() {
		if stack.a.Empty() && !stack.b.Empty() {
			for {
				value, _ = stack.b.Pop()
				if stack.b.Empty() {
					return value, nil
				}
				stack.a.Push(value)
			}
		} else if stack.b.Empty() && !stack.a.Empty() {
			for {
				value, _ = stack.a.Pop()
				if stack.a.Empty() {
					return value, nil
				}
				stack.b.Push(value)
			}
		} else {
			return 0, errors.New("stack is empty")
		}
		return value, nil
	}
	return value, errors.New("stack invaild!")
}

func (stack *myStack) Full() bool {
	if stack.a.Full() || stack.b.Full() {
		return true
	}
	return false
}

func (stack *myStack) Empty() bool {
	if stack.a.Empty() && stack.b.Empty() {
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
	for i := 0; i < 20; i++ {
		if value, err := stack.Pop(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(value)
		}
	}

}
