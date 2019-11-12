package main

import (
	"errors"
	"fmt"
)

func main() {
	heap := new(MinHeap)
	array := []int{7, 1, 3, 10, 5, 2, 8, 9, 6}

	if err := heap.Init(50, array); err != nil {
		fmt.Println(err)
		return
	}

	// heap.PrintHeap()
	// heap.Insert(0)
	// heap.PrintHeap()
	// heap.Delete()
	// heap.PrintHeap()
	// heap.HeapSort(array)
}

func (h *MinHeap) HeapSort(array []int) {
	if err := h.BuildMinHeap(array); err != nil {
		fmt.Println(err)
		return
	}
	for !h.IsEmpty() {
		val, _ := h.Delete()
		fmt.Printf("%d ", val)
	}
	fmt.Println()

}

type MinHeap struct {
	array    []int
	size     int
	capacity int
}

func (h *MinHeap) PrintHeap() {
	if h.size <= 0 {
		return
	}
	for i := 0; i < h.size; i++ {
		fmt.Printf("%d ", h.array[i])
	}
	fmt.Println()
	return
}

func (h *MinHeap) Init(capacity int, array []int) error {
	if capacity <= 0 {
		return errors.New("capacity is invaild!")
	}
	h.size = 0
	h.capacity = capacity
	h.array = make([]int, h.capacity)
	if len(array) > h.capacity {
		return errors.New("array is too bigger than capacity")
	}

	for i, v := range array {
		h.array[i] = v
		h.size++
	}

	if err := h.BuildMinHeap(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (h *MinHeap) IsEmpty() bool {
	return h.size == 0
}

func (h *MinHeap) IsFull() bool {
	return !(h.size < h.capacity)
}

func (h *MinHeap) Insert(value int) error {
	// O(logn)
	if h.IsFull() {
		return errors.New("heap is full!")
	}
	h.array[h.size] = value
	h.size++
	h.upHeap()
	return nil
}

func (h *MinHeap) Delete() (value int, err error) {
	// O(logn)
	if h.IsEmpty() {
		return 0, errors.New("Heap is empty!")
	}
	value = h.array[0]
	err = nil
	parent := h.array[h.size-1]
	h.size--
	h.array[0] = parent
	h.downHeap(0)
	return
}

func (h *MinHeap) BuildMaxHeap(array []int) error {
	if len(array) > h.capacity {
		return errors.New("array is too bigger than heap capacity!")
	}
	h.size = 0
	for i, v := range array {
		h.array[i] = v
		h.size++
	}
	for i := h.size / 2; i >= 0; i-- {
		h.buildMaxHeap(i)
	}
	return nil
}

func (h *MinHeap) buildMaxHeap(parent int) {
	array := h.array[:h.size]
	child := parent*2 + 1
	for child < h.size {
		if child+1 < h.size && array[child+1] > array[child] {
			child++
		}
		if array[child] <= array[parent] {
			break
		}
		array[child], array[parent] = array[parent], array[child]
		parent = child
		child = parent*2 + 1
	}
	return
}

func (h *MinHeap) BuildMinHeap() error {
	// O(n*logn)
	for i := h.size / 2; i >= 0; i-- {
		h.downHeap(i)
	}
	return nil
}

func (h *MinHeap) upHeap() {
	// O(logn)
	if h.size <= 1 {
		return
	}
	child := h.size - 1
	parent := (child - 1) / 2
	for child > 0 {
		if h.array[parent] > h.array[child] {
			h.array[parent], h.array[child] = h.array[child], h.array[parent]
		}
		child = parent
		parent = (child - 1) / 2
	}
	return
}

func (h *MinHeap) downHeap(parent int) {
	// O(logn)
	if h.size <= 1 {
		return
	}
	array := h.array[:h.size]
	child := parent*2 + 1
	for child < len(array) {
		if child+1 < len(array) && array[child] > array[child+1] {
			child++
		}
		if array[child] > array[parent] {
			break
		}
		array[parent], array[child] = array[child], array[parent]
		parent = child
		child = child*2 + 1
	}
	return
}
