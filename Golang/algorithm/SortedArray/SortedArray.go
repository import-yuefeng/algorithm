package main

import (
	"errors"
	"fmt"
)

type SortedArray struct {
	pointer int
	n       int
	minheap *Heap
	maxheap *Heap
}

type Heap struct {
	data     []int
	capacity int
	size     int
	flag     int8
}

func NewHeap(capacity int, flag int8) *Heap {
	return &Heap{
		capacity: capacity,
		data:     make([]int, capacity),
		size:     0,
		flag:     flag,
	}
}

func (h *Heap) Insert(x int) error {
	if h.size == h.capacity {
		return errors.New("Heap is full")
	}
	h.data[h.size] = x
	h.upNode(h.size)
	h.size++
	return nil
}

func (h *Heap) Delete() (x int, err error) {
	if h.size == 0 {
		return -1, errors.New("Heap is empty")
	}
	x = h.data[0]
	err = nil
	h.size--
	h.data[0] = h.data[h.size]
	h.downNode(0)
	return
}

func (h *Heap) upNode(index int) {
	if index < 0 || h.size == 0 {
		return
	}
	parent, child := index/2, index
	for child >= 0 {
		if h.flag == 1 {
			// maxHeap
			if child+1 < h.size && h.data[child+1] > h.data[child] {
				child++
			}
			if h.data[child] <= h.data[parent] {
				break
			}
		} else {
			if child+1 < h.size && h.data[child+1] < h.data[child] {
				child++
			}
			if h.data[child] >= h.data[parent] {
				break
			}
		}
		h.data[child], h.data[parent] = h.data[parent], h.data[child]
		child = parent
		parent /= 2
	}
	return
}

func (h *Heap) downNode(index int) {
	if h.size == 0 || index < 0 {
		return
	}
	child, parent := index*2, index
	for child < h.size {
		if h.flag == 1 {
			// maxHeap
			if child+1 < h.size && h.data[child+1] > h.data[child] {
				child++
			}
			if h.data[child] <= h.data[parent] {
				break
			}
		} else {
			if child+1 < h.size && h.data[child+1] < h.data[child] {
				child++
			}
			if h.data[child] >= h.data[parent] {
				break
			}
		}
		h.data[child], h.data[parent] = h.data[parent], h.data[child]
		parent = child
		child *= 2
	}
}

func (h *Heap) Top() int {
	return h.data[0]
}

func (h *Heap) IsEmpty() bool {
	return h.size == 0
}

func NewSortedArray(n int, pointer int) *SortedArray {
	s := new(SortedArray)
	s.minheap = NewHeap(n, 0)
	s.maxheap = NewHeap(n, 1)
	for i := 0; i <= n; i++ {
		if i >= pointer {
			s.minheap.Insert(i)
		} else {
			s.maxheap.Insert(i)
		}
	}
	s.pointer = pointer
	s.n = n
	return s
}

func (s *SortedArray) moverPointerRight() {
	x, _ := s.minheap.Delete()
	s.maxheap.Insert(x)
	s.pointer = s.minheap.Top()
}

func (s *SortedArray) movePointerLeft() {
	x, _ := s.maxheap.Delete()
	s.minheap.Insert(x)
	s.pointer = s.minheap.Top()
}

func (s *SortedArray) printPointer() int {
	fmt.Println(s.pointer)
	return s.pointer
}

func main() {
	s := NewSortedArray(100, 20)
	s.moverPointerRight()
	s.printPointer()
	s.movePointerLeft()
	s.movePointerLeft()
	s.movePointerLeft()
	s.printPointer()
}
