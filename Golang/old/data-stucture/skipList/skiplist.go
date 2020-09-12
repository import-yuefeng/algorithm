package main

import (
	"errors"
	"fmt"
	"math/rand"
)

const (
	p float32 = 0.25
)

type Node struct {
	Score   float64
	Val     interface{}
	forward []*Node
}

type SkipList struct {
	maxLength int
	header    *Node
	size      int
	level     int
}

func NewNode(score float64, val interface{}, level int) *Node {
	return &Node{
		Score:   score,
		Val:     val,
		forward: make([]*Node, level),
	}
}

func NewSkipList(maxLength int) *SkipList {
	return &SkipList{
		maxLength: maxLength,
		header:    NewNode(0, 0, maxLength),
	}
}

func (sk *SkipList) randomLevel() int {
	level := 1
	if rand.Float32() < p && level < sk.maxLength {
		level++
	}
	return level
}

func (sk *SkipList) Search(score float64) (res *Node, err error) {

	x := sk.header
	for i := sk.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].Score < score {
			x = x.forward[i]
		}
	}
	x = x.forward[0]
	if x != nil && x.Score == score {
		return x, nil
	}
	return nil, errors.New("Not found")
}

func (sk *SkipList) Insert(score float64, val interface{}) {
	x := sk.header
	update := make([]*Node, sk.maxLength)
	for i := sk.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].Score < score {
			x = x.forward[i]
		}
		update[i] = x
	}
	x = x.forward[0]
	if x != nil && x.Score == score {
		x.Val = val
		return
	}
	level := sk.randomLevel()
	if level > sk.level {
		level = sk.level + 1
		update[sk.level] = sk.header
		sk.level = level
	}
	newNode := NewNode(score, val, level)
	for i := 0; i < level; i++ {
		newNode.forward[i] = update[i].forward[i]
		update[i].forward[i] = newNode
	}
	sk.size++
	return
}

func (sk *SkipList) Delete(score float64) {
	if sk.size == 0 {
		return
	}
	update := make([]*Node, sk.maxLength)
	x := sk.header
	for i := sk.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].Score < score {
			x = x.forward[i]
		}
		update[i] = x
	}
	x = x.forward[0]
	if x != nil && x.Score == score {
		for i := 0; i < sk.level; i++ {
			if update[i].forward[i] != x {
				return
			}
			update[i].forward[i] = x.forward[i]
		}
	}
	return
}

func main() {
	// fmt.Println("hello world")
	skipList := NewSkipList(16)
	skipList.Insert(10, 10)
	skipList.Insert(4, 4)
	skipList.Insert(2, 2)
	if res, err := skipList.Search(2); err != nil {
		fmt.Println(err)
	} else {
		if t, ok := res.Val.(int); ok {
			fmt.Println(t)
		}
	}
}
