package main

type Hashmap struct {
	oldBucket []*node
	newBucket []*node
	load      int
	used      int
	size      int
	rehashidx int
	inRehash  bool
}

type node struct {
	value interface{}
	key   string
	Next  *node
	Front *node
}

func NewNode(val interface{}, key string) *node {
	return &node{
		value: val,
		key:   key,
		Next:  nil,
		Front: nil,
	}
}
