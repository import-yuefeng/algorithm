package main

import (
	"errors"
	"fmt"
)

type stack struct {
	top      int
	size     int
	capacity int
	data     []interface{}
}

func (s *stack) Init(capacity int) error {
	if capacity <= 0 {
		return errors.New("capacity is invaild!")
	}
	s.capacity = capacity
	s.top, s.size = -1, 0
	s.data = make([]interface{}, s.capacity)
	return nil
}

func (s *stack) Push(value interface{}) error {
	if s.size == s.capacity {
		return errors.New("stack is full!")
	}
	s.top++
	s.data[s.top] = value
	s.size++
	return nil
}

func (s *stack) Pop() (value interface{}, err error) {
	if s.size == 0 {
		return nil, errors.New("stack is empty!")
	}
	value = s.data[s.top]
	err = nil
	s.top--
	s.size--
	return
}

type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}

func Init(val int) (bt *BinaryTree, err error) {
	bt = new(BinaryTree)
	bt.Left, bt.Right = nil, nil
	bt.Val = val
	return bt, nil
}

func (bt *BinaryTree) InsertArray(nums []int) error {
	for _, v := range nums {
		if err := bt.Insert(v); err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func (bt *BinaryTree) Insert(value int) error {
	tmp := bt
	newNode := new(BinaryTree)
	newNode.Val = value
	newNode.Left, newNode.Right = nil, nil

	for tmp != nil {
		if tmp.Val >= newNode.Val {
			if tmp.Left != nil {
				tmp = tmp.Left
			} else {
				tmp.Left = newNode
				return nil
			}
		} else {
			if tmp.Right != nil {
				tmp = tmp.Right
			} else {
				tmp.Right = newNode
				return nil
			}
		}
	}
	return errors.New("insert error")
}

func (bt *BinaryTree) Delete(value int) (val int, err error) {
	if bt == nil {
		return -1, errors.New("bt is nil!")
	}
	tmp := bt
	for bt != nil {
		if bt.Val == value {
			val, err := delMinNode(bt.Right)
			if err != nil {
				fmt.Println(err)
				if val == -1 {
					tmp.Right = nil
					return value, nil
				} else {
					bt.Right = nil
					bt.Val = val
					return val, nil
				}
			}
		} else if bt.Val > value {
			bt = bt.Left
		} else {
			bt = bt.Right
		}
	}
	return -1, errors.New("not find value!")
}

func (bt *BinaryTree) Find(value int) bool {
	if bt == nil {
		return false
	}
	for bt != nil {
		if bt.Val == value {
			return true
		} else if bt.Val > value {
			bt = bt.Left
		} else {
			bt = bt.Right
		}
	}
	return false
}

func delMinNode(sbt *BinaryTree) (val int, err error) {
	if sbt == nil {
		return -1, errors.New("sbt is nil")
	}
	if sbt.Left == nil {
		return sbt.Val, errors.New("only one node")
	}
	tmp := sbt
	for sbt.Left != nil {
		tmp = sbt
		sbt = sbt.Left
	}
	val = sbt.Val
	err = nil
	tmp.Left = nil
	return
}

func (bt *BinaryTree) DFS() {
	s := new(stack)
	s.Init(100)
	err := s.Push(bt)
	if err != nil {
		fmt.Println(err)
		return
	}
	for s.size != 0 {
		tmp, err := s.Pop()
		if err != nil {
			fmt.Println(err)
			return
		}
		val := tmp.(*BinaryTree)
		fmt.Printf(" %d", val.Val)
		// first push right child
		if val.Right != nil {
			err := s.Push(val.Right)
			if err != nil {
				fmt.Println(err)
			}
		}
		if val.Left != nil {
			err := s.Push(val.Left)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

type Queue struct {
	data     []interface{}
	size     int
	rare     int
	front    int
	capacity int
}

func (q *Queue) Init(capacity int) error {
	if capacity <= 0 {
		return errors.New("capacity is invaild!")
	}
	q.capacity = capacity
	q.data = make([]interface{}, q.capacity)
	q.front, q.size, q.rare = 0, 0, 0
	return nil
}

func (q *Queue) Push(val interface{}) error {
	if q.size == q.capacity {
		return errors.New("queue is full")
	}
	q.size++
	q.rare = (q.rare + 1) % q.capacity
	q.data[q.rare] = val
	return nil
}

func (q *Queue) Pop() (val interface{}, err error) {
	if q.size == 0 {
		return nil, errors.New("queue is empty")
	}
	q.front = (q.front + 1) % q.capacity
	val = q.data[q.front]
	err = nil
	q.size--
	return
}

func (bt *BinaryTree) BFS() {
	q := new(Queue)
	q.Init(100)
	err := q.Push(bt)
	if err != nil {
		fmt.Println(err)
		return
	}
	for q.size != 0 {
		tmp, err := q.Pop()
		if err != nil {
			fmt.Println(err)
		}
		val := tmp.(*BinaryTree)
		fmt.Printf(" %d", val.Val)
		if val.Left != nil {
			err := q.Push(val.Left)
			if err != nil {
				fmt.Println(err)
			}
		}
		if val.Right != nil {
			err := q.Push(val.Right)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func main() {
	head, err := Init(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	head.InsertArray([]int{10, 2, 0, -1, 2, 12, 1})
	head.DFS()
	fmt.Println()
	head.BFS()
	head.Delete(10)
	head.BFS()
}
