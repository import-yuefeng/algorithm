// package linkStack
package main

import (
	"errors"
	"fmt"
)

type linkStackNode struct {
	End      *linkStackNode
	Capacity int
	size     int
	Value    int
	Next     *linkStackNode
}

func (lstack *linkStackNode) Init(capacity int) {
	if capacity <= 0 {
		return
	}
	lstack.Capacity = capacity
	lstack.Next, lstack.End = lstack, lstack
	lstack.Value, lstack.size = 0, 0
	return
}

func (lstack *linkStackNode) Push(value int) error {
	if lstack.Next == nil {
		return errors.New("linkStackNode is nil")
	}
	if lstack.isFull() {
		return errors.New("linkStackNode is full")
	}
	newElement := new(linkStackNode)
	newElement.Value = value
	newElement.Next = nil
	lstack.End.Next = newElement
	lstack.End = newElement
	lstack.size++
	return nil
}

func (lstack *linkStackNode) Pop() (value int, err error) {
	if lstack.isEmpty() {
		return 0, errors.New("linkStackNode is empty")
	}
	tmp := lstack
	value = lstack.End.Value
	for tmp.Next != lstack.End {
		tmp = tmp.Next
	}
	tmp.Next = nil
	lstack.End = tmp
	return value, nil
}

func (lstack *linkStackNode) PrintStack() {
	if lstack.Next == nil {
		fmt.Println("linkStackNode is nil")
		return
	}
	tmp := lstack.Next
	for tmp.Next != nil {
		fmt.Printf("%d ", tmp.Value)
		tmp = tmp.Next
	}
	fmt.Println()
	return
}

func (lstack *linkStackNode) Size() int {
	return lstack.size
}

func (lstack *linkStackNode) isFull() bool {
	return lstack.size == lstack.Capacity
}

func (lstack *linkStackNode) isEmpty() bool {
	return lstack.size == 0
}

func main() {

	a := new(linkStackNode)
	a.Pop()

	a.Push(1)
	// a.Init(10)
	// for i := 0; i < 10; i++ {
	// 	a.Push(i)
	// }
	// a.PrintStack()
	// a.Pop()
	a.PrintStack()
}
