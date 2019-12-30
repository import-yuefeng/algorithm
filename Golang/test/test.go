package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	// hashWriting := 1
	// c := 0
	// c ^= hashWriting
	// var a int = 1
	// fmt.Println(a << 4)
	// fmt.Println(c)
	// c ^= hashWriting
	// fmt.Println(c)

	// rand.Seed(time.Now().Unix())
	// var array [100]int
	// for i := 0; i < 100; i++ {
	// 	array[i] = rand.Intn(200)
	// }
	// fmt.Println(array)
	// qSort(array[:], 0, len(array)-1)
	// fmt.Println(array)
}

type quickSortData struct {
	Left, Right int
}

func qSort(array []int, left, right int) {
	if left < right && left >= 0 {
		stack := NewStack(100)
		err := stack.Push(&quickSortData{Left: left, Right: right})
		if err != nil {
			fmt.Println(err)
			return
		}
		for !stack.IsEmpty() {
			value, err := stack.Pop()
			if err != nil {
				fmt.Println(err)
				return
			}
			left, right = value.(*quickSortData).Left, value.(*quickSortData).Right
			if left < right && left >= 0 {
				l, r := left, right
				middle := array[(l+r)>>1]
				for l < r {
					for l < r && array[l] < middle {
						l++
					}
					for l < r && array[r] > middle {
						r--
					}
					if l >= r {
						break
					}
					if array[r] == array[l] && array[r] == middle {
						r--
					} else {
						array[l], array[r] = array[r], array[l]
					}
				}
				err := stack.Push(&quickSortData{Left: left, Right: l - 1})
				if err != nil {
					fmt.Println(err)
					return
				}
				err = stack.Push(&quickSortData{Left: l + 1, Right: right})
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}

func mergeSort(array []int, left, right int) {
	if left < right {
		middle := (left + right) >> 1
		mergeSort(array, left, middle)
		mergeSort(array, middle+1, right)
		merge(array, left, middle, right)
	}
}

func merge(array []int, left, middle, right int) {

}

type Stack struct {
	top      int
	size     int
	capacity int
	data     []interface{}
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) IsFull() bool {
	return s.size == s.capacity
}

func NewStack(capacity int) *Stack {
	if capacity <= 0 {
		return nil
	}
	stack := new(Stack)
	stack.top = -1
	stack.size = 0
	stack.capacity = capacity
	stack.data = make([]interface{}, stack.capacity)
	return stack
}

func (s *Stack) Push(x interface{}) error {
	if s.IsFull() {
		return errors.New("stack is full")
	}
	s.top++
	s.data[s.top] = x
	s.size++
	return nil
}

func (s *Stack) Pop() (x interface{}, err error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	x = s.data[s.top]
	err = nil
	s.top--
	s.size--
	return
}
