package main

import (
	"errors"
)

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

func (h *Hashmap) Get(key string) (val interface{}, err error) {
	res, err := findNode(key, h.oldBucket)
	if err != nil {
		return nil, err
	}
	if h.inRehash {
		res, err := findNode(key, h.newBucket)
		if err != nil {
			return nil, err
		}
		h.rehash()
		return res, nil
	}
	h.rehash()
	return res, nil
}

func (h *Hashmap) Set(key string, value interface{}) error {
	h.rehash()
	if !h.inRehash {
		if _, err := h.insertNode(value, key, h.oldBucket); err != nil {
			return err
		}
		return nil
	}
	if _, err := h.insertNode(value, key, h.newBucket); err != nil {
		return err
	}
	return nil
}

func (h *Hashmap) Delete(key string) {
	res, err := findNode(key, h.oldBucket)
	if res != nil && err == nil {
		h.used--
		if res.Front == nil {
			h.deleteHead(key)
			return
		}
		front := res.Front
		front.value = res.value
		front.Next = res.Next
		front.Next.Front = front
		front.key = res.key
		return
	}
	if !h.inRehash {
		return
	}
	res, err = findNode(key, h.newBucket)
	if err != nil || res == nil {
		return
	}
	h.used--
	if res.Front == nil {
		h.deleteHead(key)
		return
	} else {
		front := res.Front
		front.value = res.value
		front.Next = res.Next
		front.Next.Front = front
		front.key = res.key
		return
	}
	return
}

func (h *Hashmap) deleteHead(key string) {
	k := Hash(key)
	old := h.oldBucket[k%uint32(len(h.oldBucket))]
	if old.key == key {
		h.oldBucket[k%uint32(len(h.oldBucket))] = nil
		return
	}
	if h.inRehash {
		v := h.newBucket[k%uint32(len(h.newBucket))]
		if v.key == key {
			h.newBucket[k%uint32(len(h.newBucket))] = nil
		}
	}
	return
}

func findNode(key string, bucket []*node) (res *node, err error) {
	k := Hash(key)
	curBucket := bucket[k%uint32(len(bucket))]
	if curBucket == nil {
		return nil, errors.New("Not found")
	}
	curNode := curBucket
	for curNode.Next != nil && curNode.key != key {
		curNode = curNode.Next
	}
	if curNode.key == key {
		return curNode, nil
	}
	return nil, errors.New("Not found")
}

func (h *Hashmap) insertNode(val interface{}, key string, bucket []*node) (res *node, err error) {
	k := Hash(key)
	curBucket := bucket[k%uint32(len(bucket))]
	if curBucket == nil {
		h.used++
		newNode := NewNode(val, key)
		bucket[k%uint32(len(bucket))] = newNode
		return newNode, nil
	}
	curNode := curBucket
	for curNode.Next != nil && curNode.key != key {
		curNode = curNode.Next
	}
	if curNode.key == key {
		curNode.value = val
		return curNode, nil
	}
	h.used++
	newNode := NewNode(val, key)
	newNode.Next = curBucket.Next
	if curNode.Next != nil {
		newNode.Next.Front = newNode
	}
	newNode.Front = curBucket
	curBucket.Next = newNode
	return newNode, nil
}
