package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	failed := t.Failed
	stack := NewStack(10)
	for i := 0; i < 1000; i++ {
		stack.Push(i)
	}
	for i := 999; i >= 0; i-- {
		t, err := stack.Pop()
		if err != nil {
			fmt.Println("stack pop err: ", err)
		}
		val, ok := t.(int)
		if !ok {
			fmt.Println("assert int error: ", ok)
		}
		if i != val {
			failed()
		}
	}
}
