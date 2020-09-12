package stack

import "errors"

type Stack struct {
	data []interface{}
	size int
	cap  int
	top  int
}

func NewStack(cap int) *Stack {
	return &Stack{
		data: make([]interface{}, cap),
		cap:  cap,
		top:  -1,
	}
}

func (s *Stack) IsFull() bool {
	return s.size == s.cap
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Push(val interface{}) {
	if s.IsFull() {
		s.cap++
		s.data = append(s.data, val)
		s.top++
		s.size++
		return
	}
	s.top++
	s.data[s.top] = val
	s.size++
	return
}

func (s *Stack) Pop() (val interface{}, err error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is nil")
	}
	val = s.data[s.top]
	s.size--
	s.top--
	return
}
