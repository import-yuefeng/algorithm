package main

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue()
	test := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(test); i++ {
		if i&1 == 1 {
			queue.PushBack(test[i])
		}
	}
	for i := 0; i < len(test); i++ {
		if i&1 == 1 {
			val, err := queue.PopFront()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("%d ", val.Val)
		}
	}
	fmt.Println()
	for i := 0; i < len(test); i++ {
		if i&1 == 1 {
			queue.PushBack(test[i])
		}
	}
	queue.Println()
}
