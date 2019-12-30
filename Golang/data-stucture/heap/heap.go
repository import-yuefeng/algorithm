// package heap
package main

import (
	"errors"
	"fmt"
)

type heap struct {
	data []int
	size int
	flag int
	// flag =
	// 1 buildMaxHeap
	// 0 buildMinHeap
}

type charNode struct {
	char          rune
	charFrequency rune
}

func NewHeap(data []int, flag int) *heap {
	h := &heap{
		data: data,
		size: len(data),
		flag: flag,
	}
	h.buildHeap()
	// h.buildMinHeap
	return h
}

func (h *heap) buildHeap() {
	if h.size <= 1 {
		return
	}
	for i := h.size / 2; i >= 0; i-- {
		h.downNode(i)
	}
	return
}

func (h *heap) Delete() (x int, err error) {
	if h.size == 0 {
		return -1, errors.New("heap is nil")
	}
	x = h.data[0]
	h.data[0] = h.data[h.size-1]
	h.size--
	h.data = h.data[:h.size]
	err = nil
	h.downNode(0)
	return
}

func (h *heap) Insert(x int) error {
	// unlimited heap....
	h.data = append(h.data, x)
	h.upNode(h.size)
	h.size++
	return nil
}

func (h *heap) downNode(parent int) {
	child := parent * 2
	for child < h.size {
		if h.flag == 1 {
			if child+1 < h.size && h.data[child+1] > h.data[child] {
				child++
			}
		} else {
			if child+1 < h.size && h.data[child+1] < h.data[child] {
				child++
			}
		}
		if h.flag == 1 {
			if h.data[child] <= h.data[parent] {
				break
			}
		} else {
			if h.data[child] >= h.data[parent] {
				break
			}
		}
		h.data[child] ^= h.data[parent]
		h.data[parent] ^= h.data[child]
		h.data[child] ^= h.data[parent]
		parent = child
		child = 2 * parent
	}
}

func (h *heap) upNode(child int) {
	parent := child / 2
	for child >= 0 {
		if h.flag == 1 {
			if child+1 < h.size && h.data[child+1] > h.data[child] {
				child++
			}
		} else {
			if child+1 < h.size && h.data[child+1] < h.data[child] {
				child++
			}
		}
		if h.flag == 1 {
			if h.data[child] <= h.data[parent] {
				break
			}
		} else {
			if h.data[child] >= h.data[parent] {
				break
			}
		}
		h.data[child] ^= h.data[parent]
		h.data[parent] ^= h.data[child]
		h.data[child] ^= h.data[parent]
		child = parent
		parent = child / 2
	}
}

func (h *heap) PrintHeap() {
	for i := 0; i < h.size; i++ {
		fmt.Printf("%d ", h.data[i])
	}
	fmt.Println()
}

func main() {
	array := []int{1, 5, 4, 2, 30000, 1, 2, 4, 5, 1999}
	h := NewHeap(array, 1)
	h.PrintHeap()
	h.Insert(30001)
	h.Insert(20000)
	for h.size > 0 {
		if x, err := h.Delete(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%d ", x)
		}
	}
	fmt.Println()

	h = NewHeap(array, 0)
	h.PrintHeap()
	for h.size > 0 {
		if x, err := h.Delete(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%d ", x)
		}
	}
	fmt.Println()
}
