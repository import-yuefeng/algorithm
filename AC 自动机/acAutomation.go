// package acAutomation
package main

import (
	"errors"
	"fmt"
)

type ACAutomation struct {
	fail   *ACAutomation
	child  map[rune]*ACAutomation
	key    rune
	parent *ACAutomation
}

func Init() (head *ACAutomation) {
	head = new(ACAutomation)
	head.child = make(map[rune]*ACAutomation)
	head.key = 0
	head.fail = nil
	head.parent = nil
	return head
}

func main() {
	ac := Init()
	texts := []string{"abc", "abd", "acd", "bca"}
	ac.Insert(texts)
	// ac.PrintAC()
	ac.BuildFailPoint()
	ac.PrintAC()
	res := ac.MatchText("abca")
	fmt.Println(res)
	// ac.Delete("abc")
	// ac.PrintAC()
}

func (ac *ACAutomation) MatchText(text string) int {
	num := 0
	tmp := ac
	for _, v := range text {
		for {
			if res, ok := tmp.child[v]; ok {
				tmp = res
				if len(tmp.child) == 0 {
					num++
				}
				break
			} else {
				tmp = tmp.fail
				if tmp == nil || tmp == ac {
					break
				}
			}
		}
	}
	return num
}

func (ac *ACAutomation) BuildFailPoint() error {
	tmp := ac
	for _, v := range tmp.child {
		v.fail = tmp
	}
	q := new(Queue)
	q.Init(100)
	q.Push(tmp)
	for q.size > 0 {
		t, err := q.Pop()
		if err != nil {
			fmt.Println(err)
			return err
		}
		for i, v := range t.(*ACAutomation).child {
			if err := q.Push(v); err != nil {
				fmt.Println(err)
				return err
			}
			prev := t.(*ACAutomation).fail
			for prev != nil {
				if res, ok := prev.child[i]; ok {
					v.fail = res
					break
				} else {
					prev = prev.fail
				}
			}
			if prev == nil {
				v.fail = ac
			}
		}
	}
	return nil
}

func (ac *ACAutomation) Insert(texts []string) error {
	if len(texts) == 0 {
		return errors.New("texts is empty!")
	}
	for _, v := range texts {
		tmp := ac
		for _, char := range v {
			if res, ok := tmp.child[char]; ok {
				tmp = res
			} else {
				newNode := new(ACAutomation)
				newNode.child = make(map[rune]*ACAutomation)
				newNode.key = char
				newNode.fail = nil
				tmp.child[char] = newNode
				newNode.parent = tmp
				tmp = newNode
			}
		}
	}
	return nil
}

func (ac *ACAutomation) Search(text string) bool {
	tmp := ac
	for _, v := range text {

		if res, ok := tmp.child[v]; ok {
			tmp = res
		} else {
			return false
		}
	}
	return true
}

func (ac *ACAutomation) PrintAC() {
	q := new(Queue)
	q.Init(100)
	if err := q.Push(ac); err != nil {
		fmt.Println(err)
		return
	}
	for q.size > 0 {
		if tmp, err := q.Pop(); err != nil {
			fmt.Println(err)
			return
		} else {
			for i, v := range tmp.(*ACAutomation).child {
				fmt.Printf("%v ", i)
				fmt.Printf("(%v -> %v)", i, v.fail.key)
				// fmt.Printf("(%v -> %v)", i, v.fail.child)
				if err := q.Push(v); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
	fmt.Println()
	return
}

func (ac *ACAutomation) Delete(text string) error {
	if isExist := ac.Search(text); !isExist {
		return errors.New(text + " is not exist!")
	}
	tmp := ac
	for _, v := range text {
		if res, ok := tmp.child[v]; ok {
			if len(tmp.child) == 1 {
				tmp.child = nil
				return nil
			} else if len(tmp.child) > 1 && len(res.child) <= 1 {
				delete(tmp.child, v)
				return nil
			}
			tmp = res
		} else {
			return errors.New(text + " is not exist!")
		}
	}
	return nil
}

type Queue struct {
	rare     int
	front    int
	capacity int
	size     int
	data     []interface{}
}

func (q *Queue) Init(capacity int) error {
	if capacity <= 0 {
		return errors.New("capacity is invaild!")
	}
	q.capacity = capacity
	q.size = 0
	q.rare, q.front = 0, 0
	q.data = make([]interface{}, q.capacity)
	return nil
}

func (q *Queue) Push(value interface{}) error {
	if q.size == q.capacity {
		return errors.New("queue is full!")
	}
	q.rare = (q.rare + 1) % q.capacity
	q.data[q.rare] = value
	q.size++
	return nil
}

func (q *Queue) Pop() (value interface{}, err error) {
	if q.size == 0 {
		return nil, errors.New("queue is empty!")
	}
	q.front = (q.front + 1) % q.capacity
	value = q.data[q.front]
	q.size--
	return
}
