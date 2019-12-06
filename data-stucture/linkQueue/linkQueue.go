package main

import (
	"errors"
	"fmt"
)

type linkQueueNode struct {
	size     int
	capacity int
	value    int
	next     *linkQueueNode
	front    *linkQueueNode
	end      *linkQueueNode
}

func (lqueue *linkQueueNode) Init(capacity int) {
	if capacity <= 0 {
		return
	}
	lqueue.capacity = capacity
	lqueue.size = 0
	lqueue.front, lqueue.end, lqueue.next = lqueue, lqueue, lqueue
	return
}

func (lqueue *linkQueueNode) Push(value int) error {
	if lqueue.isFull() {
		return errors.New("linkQueueNode is full")
	}
	newElement := new(linkQueueNode)
	newElement.value = value
	newElement.next = nil
	lqueue.end.next = newElement
	lqueue.end = newElement
	lqueue.size++
	return nil
}

func (lqueue *linkQueueNode) Pop() (value int, err error) {
	if lqueue.isEmpty() {
		return 0, errors.New("linkQueueNode is empty")
	}
	value = lqueue.front.next.value
	lqueue.front.next = lqueue.front.next.next
	lqueue.size--
	return value, nil
}

func (lqueue *linkQueueNode) Size() int {
	return lqueue.size
}

func (lqueue *linkQueueNode) isEmpty() bool {
	return lqueue.size == 0
}

func (lqueue *linkQueueNode) isFull() bool {
	return lqueue.size == lqueue.capacity
}

func (lqueue *linkQueueNode) PrintQueue() {
	tmp := lqueue
	if tmp.next == nil {
		fmt.Println("linkStackNode is nil")
		return
	}
	for tmp.next != nil {
		fmt.Printf("%d ", tmp.value)
		tmp = tmp.next
	}
	fmt.Println()
}
func main() {
	a := new(linkQueueNode)
	a.Pop()

	a.Push(1)
	a.Init(10)
	for i := 0; i < 10; i++ {
		a.Push(i)
	}
	a.PrintQueue()
	a.Pop()
	a.PrintQueue()

}
