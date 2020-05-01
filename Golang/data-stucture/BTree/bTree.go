package main

import (
	"errors"
	"fmt"
	"unsafe"
)

type BTreeNode struct {
	val      []*Node
	size     int
	capacity int
}

type Node struct {
	prev, left, right *BTreeNode
	key               int
	value             int
}

type BTree struct {
	capacity int
	root     *BTreeNode
}

func NewBTreeNode(capacity int) *BTreeNode {
	return &BTreeNode{
		capacity: capacity,
		val:      make([]*Node, capacity),
		size:     0,
	}
}

func NewNode(key, value int) *Node {
	return &Node{
		key:   key,
		value: value,
	}
}

func NewBTree(capacity int) *BTree {
	if capacity < 0 {
		return nil
	}
	return &BTree{
		capacity: capacity,
		root:     NewBTreeNode(capacity),
	}
}

func (b *BTree) Search(key int) (res *Node, err error) {
	return nil, errors.New("Not found!")
}

// func (b *BTree) Insert(key, value int) {

// }

// func (b *BTree) Delete(key int) {

// }

func main() {
	var c uint32 = 2345
	var d [4]byte
	p := unsafe.Pointer(&c)
	fmt.Println(p)
	q := (*[4]byte)(p)
	copy(d[0:], (*q)[0:])
	t := d[:]
	fmt.Println(*(*string)(unsafe.Pointer(&t)))
	fmt.Println(d)

}
