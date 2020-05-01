package main

import (
	"errors"
)

type Heap struct {
	data []int
	size int
}

func NewHeap(data []int) *Heap {
	return &Heap{
		data: data,
		size: len(data),
	}
}

func (h *Heap) BuildHeap() {
	for i := h.size / 2; i >= 0; i-- {
		h.downNode(i)
	}
}

func (h *Heap) Delete() (x int, err error) {
	if h.size > 0 {
		x = h.data[0]
		h.data[0] = h.data[len(h.data)-1]
		h.data = h.data[:len(h.data)-1]
		h.size--
		h.downNode(0)
		return x, nil
	}
	return -1, errors.New("Heap is empty")
}

func (h *Heap) upNode(i int) {
	child := i
	parent := child / 2
	for child >= 0 {
		if child-1 >= 0 && h.data[child] < h.data[child-1] {
			child--
		}
		if h.data[child] >= h.data[parent] {
			break
		}
		h.data[child], h.data[parent] = h.data[parent], h.data[child]
		child = parent
		parent = child / 2
	}
}

func (h *Heap) downNode(i int) {
	parent := i
	child := i*2 + 1
	for child < h.size {
		if child+1 < h.size && h.data[child+1] < h.data[child] {
			child++
		}
		if h.data[child] >= h.data[parent] {
			break
		}
		h.data[child], h.data[parent] = h.data[parent], h.data[child]
		parent = child
		child = parent*2 + 1
	}
}

type Queue struct {
	front, rare int
	data        []int
	size        int
	capacity    int
}

func NewQueue() *Queue {
	// default capacity: 2
	return &Queue{
		front:    0,
		rare:     0,
		size:     0,
		capacity: 2,
		data:     make([]int, 2),
	}
}

func (q *Queue) Push(x int) {
	if q.IsFull() {
		buf := NewQueue()
		buf.data = make([]int, q.Size()*2)
		for i := 0; i < q.size; i++ {
			buf.data[i] = q.data[q.front]
			q.front = (q.front + 1) % q.capacity
		}
		q.data = buf.data
		q.front = 0
		q.rare = q.size
		q.capacity *= 2
	}
	q.data[q.rare] = x
	q.rare = (q.rare + 1) % q.capacity
	q.size++
	return
}

func (q *Queue) Pop() (x int, err error) {
	if q.IsEmpty() {
		return -1, errors.New("Queue is empty")
	}
	x = q.data[q.front]
	q.front = (q.front + 1) % q.capacity
	q.size--
	return x, nil
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) IsFull() bool {
	return q.size == q.capacity
}

func (q *Queue) Size() int {
	return q.size
}
