package main

import (
	"errors"
	"fmt"
)

func main() {

	trie := Init()
	trie.Insert("abc")
	trie.Insert("abe")
	trie.Insert("bec")
	trie.printTrie()
	res := trie.Search("ab")
	fmt.Println(res)
}

type trieNode struct {
	key   rune
	child map[rune]*trieNode
}

func Init() (head *trieNode) {
	head = new(trieNode)
	head.key = 0
	head.child = make(map[rune]*trieNode)
	return head
}

func (t *trieNode) Insert(str string) {
	tmp := t
	for _, v := range str {
		if res, ok := tmp.child[v]; ok {
			tmp = res
		} else {
			newNode := new(trieNode)
			newNode.key = v
			newNode.child = make(map[rune]*trieNode)
			tmp.child[v] = newNode
			tmp = newNode
		}
	}
	return
}

func (t *trieNode) Search(text string) bool {
	if len(text) == 0 {
		return true
	}
	for _, v := range text {
		if res, ok := t.child[v]; ok {
			t = res
		} else {
			return false
		}
	}
	return true
}

func (t *trieNode) printTrie() {
	queue := new(Queue)
	queue.Init(100)
	if err := queue.Push(t); err != nil {
		fmt.Println(err)
		return
	}
	for queue.size != 0 {
		if tmp, err := queue.Pop(); err == nil {
			for i, v := range tmp.child {
				fmt.Printf("%v ", i)
				if err := queue.Push(v); err != nil {
					fmt.Println(err)
					return
				}
			}
		} else {
			fmt.Println(err)
			return
		}
	}
}

type Queue struct {
	front    int
	rare     int
	data     []*trieNode
	size     int
	capacity int
}

func (q *Queue) Init(capacity int) error {
	if capacity <= 0 {
		return errors.New("capacity is invaild!")
	}
	q.capacity = capacity
	q.front, q.rare = 0, 0
	q.size = 0
	q.data = make([]*trieNode, q.capacity)
	return nil
}

func (q *Queue) Push(value *trieNode) error {
	if q.size == q.capacity {
		return errors.New("queue is full!")
	}
	q.rare = (q.rare + 1) % q.capacity
	q.data[q.rare] = value
	q.size++
	return nil
}

func (q *Queue) Pop() (value *trieNode, err error) {
	if q.size == 0 {
		return nil, errors.New("queue is empty!")
	}
	q.front = (q.front + 1) % q.capacity
	value = q.data[q.front]
	q.size--
	err = nil
	return
}
