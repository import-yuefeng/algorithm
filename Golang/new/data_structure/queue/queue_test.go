package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	errFunc := t.FailNow
	q := NewQueue(100)
	t.Log("start test...")
	for i := 0; i<100; i++ {
		err := q.Push(i)
		if err != nil {
			fmt.Println(err)
		}
	}

	for i := 0 ; i<100; i++ {
		t, err := q.Pop()
		if err != nil {
			fmt.Println(err)
		}
		val, ok := t.(int)
		if !ok {
			fmt.Println("assert int error: ", ok)
		}
		if val != i {
			errFunc()
		}
	}

}
