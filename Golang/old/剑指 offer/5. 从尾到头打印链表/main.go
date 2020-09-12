package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println("hello world")
	res := B([]int{4, 7, 2, 9, 5, 2})
	fmt.Println(res)
	root := new(Linkedlist)
	root.Val = 1
	root.Next = new(Linkedlist)
	root.Next.Val = 2
	root.Next.Next = new(Linkedlist)
	root.Next.Next.Val = 3
	ReversePrintLinkedlist(root)
	PrintReverseLinkedlistByIteration(root)
}

type Linkedlist struct {
	Val  int
	Next *Linkedlist
}

func ReversePrintLinkedlist(root *Linkedlist) {
	if root == nil {
		return
	}
	ReversePrintLinkedlist(root.Next)
	fmt.Println(root.Val)
	return
}

func PrintReverseLinkedlistByIteration(root *Linkedlist) {
	stack := NewStack()
	cur := root
	for cur != nil {
		stack.Push(cur)
		cur = cur.Next
	}
	for !stack.IsEmpty() {
		tmp, err := stack.Pop()
		if err != nil {
			fmt.Println(err)
			return
		}
		if val, ok := tmp.(*Linkedlist); ok {
			fmt.Printf("%d ", val)
		} else {
			fmt.Println(ok)
			return
		}
	}
}

type Stack struct {
	top      int
	size     int
	capacity int
	data     []interface{}
}

func NewStack() *Stack {
	return &Stack{
		top:      -1,
		size:     0,
		capacity: 2,
		data:     make([]interface{}, 2),
	}
}

func (s *Stack) Push(x interface{}) {
	if s.size == s.capacity {
		buf := NewStack()
		buf.data = make([]interface{}, s.capacity*2)
		for i, v := range s.data {
			buf.data[i] = v
		}
		s.data = buf.data
		s.capacity *= 2
	}
	s.top++
	s.data[s.top] = x
	s.size++
	return
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Pop() (x interface{}, err error) {
	if s.size == 0 {
		return nil, errors.New("Stack is empty")
	}
	x = s.data[s.top]
	s.top--
	s.size--
	return x, nil
}

func B(num []int) int {
	sum, dp := make([][]int, len(num)), make([][]int, len(num))
	for i, _ := range dp {
		dp[i] = make([]int, len(num))
		sum[i] = make([]int, len(num))
	}

	for i := 0; i < len(num); i++ {
		for j := i; j < len(num); j++ {
			for q := i; q <= j; q++ {
				sum[i][j] += num[q]
			}
		}
	}
	dp[0][0] = sum[0][0]

	for i := 0; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			dp[i][j] = sum[i][j] - min(dp[i+1][j], dp[i][j-1])
		}
	}
	a := dp[0][len(num)-1]
	b := sum[0][len(num)-1] - a
	return a - b
}

func min(a, b int) int {
	if a > b {
		return a
	}
	return b
}
