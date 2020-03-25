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

func NewHashmap() *Hashmap {
	// default bucket number: 8
	// default load argument: 2
	return &Hashmap{
		load:      1,
		inRehash:  false,
		size:      2,
		used:      0,
		newBucket: make([]*node, 0),
		oldBucket: make([]*node, 2),
		rehashidx: -1,
	}
}
