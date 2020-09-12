// package cycleQueue
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

func main() {
	a := myQueue{}
	a.Init(10)
	for i := 0; i < 20; i++ {
		a.Push(i)
	}
	for i := 0; i < 3; i++ {
		if res, err := a.Pop(); err != nil {
			errors.New("pop error")
		} else {
			fmt.Println(res)
		}
	}
	for i := 0; i < 20; i++ {
		a.Push(i)
	}
	for i := 0; i < 500; i++ {
		if res, err := a.Pop(); err != nil {
			errors.New("pop error")
		} else {
			fmt.Println(res)
		}
	}

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
	queue.queue[queue.rear] = value
	queue.rear = (queue.rear + 1) % queue.capacity
	queue.size++
	return nil
}

func (queue *myQueue) Pop() (value int, err error) {
	if queue.size <= 0 {
		// queue.Empty()
		return 0, errors.New("queue is empty!")
	}
	tmpValue := queue.queue[queue.front]
	queue.front = (queue.front + 1) % queue.capacity
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
