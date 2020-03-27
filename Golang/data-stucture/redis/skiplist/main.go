package main

import (
	"fmt"
	"math/rand"
)

const (
	P float64 = 0.25
)

type skiplist struct {
	maxLevel int64
	curLevel int64
	head     *node
	size     int64
}

type node struct {
	key     string
	value   string
	forward []*node
}

func NewNode(key, value string, level int64) *node {
	return &node{
		key:     key,
		value:   value,
		forward: make([]*node, level),
	}
}

func NewSkiplist(maxLevel int64) *skiplist {
	return &skiplist{
		maxLevel: maxLevel,
		curLevel: 0,
		head:     NewNode("0", "0", maxLevel),
		size:     0,
	}
}

func (sk *skiplist) randomLevel() int64 {
	var level int64 = 1
	for rand.Float64() < P && level < sk.maxLevel {
		level++
	}
	return level
}

func (sk *skiplist) Search(key string) (res *node, isExist bool) {
	if sk.size == 0 {
		return nil, false
	}
	cur := sk.head
	for i := sk.curLevel - 1; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].key < key {
			cur = cur.forward[i]
		}
	}
	cur = cur.forward[0]
	if cur == nil || cur.key != key {
		return nil, false
	}
	return cur, true
}

func (sk *skiplist) Insert(key, value string) {
	update := make([]*node, sk.maxLevel)
	cur := sk.head
	for i := sk.curLevel - 1; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].key < key {
			cur = cur.forward[i]
		}
		update[i] = cur
	}
	cur = cur.forward[0]
	if cur != nil && cur.key == key {
		cur.value = value
		return
	}
	level := sk.randomLevel()
	if level > sk.curLevel {
		level = sk.curLevel + 1
		update[sk.curLevel] = sk.head
		sk.curLevel = level
	}
	newNode := NewNode(key, value, level)
	for i := int64(0); i < level; i++ {
		newNode.forward[i] = update[i].forward[i]
		update[i].forward[i] = newNode
	}
	sk.size++
	return
}

func (sk *skiplist) Delete(key string) {
	if sk.size == 0 {
		return
	}
	update := make([]*node, sk.maxLevel)
	cur := sk.head
	for i := sk.curLevel - 1; i >= 0; i-- {
		for cur.forward[i] != nil && cur.forward[i].key < key {
			cur = cur.forward[i]
		}
		update[i] = cur
	}
	cur = cur.forward[0]
	if cur == nil || cur.key != key {
		return
	}
	for i := int64(0); i < sk.curLevel; i++ {
		if update[i].forward[i] != cur {
			return
		}
		update[i].forward[i] = cur.forward[i]
	}
	return
}

func main() {
	skipList := NewSkiplist(16)
	skipList.Insert("10", "10")
	skipList.Insert("4", "4")
	skipList.Insert("2", "2")
	if res, ok := skipList.Search("2"); !ok {
		fmt.Println("not found")
	} else {
		fmt.Println(res)
	}
}
