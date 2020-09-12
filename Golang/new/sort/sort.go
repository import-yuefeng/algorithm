package main

import (
	"errors"
	"fmt"
)

func BubblingSort(data []int, left, right int) {
	if left >= right {
		return
	}
	for i := left; i < right; i++ {
		for j := i + 1; j < right; j++ {
			if data[j] < data[i] {
				data[i] ^= data[j]
				data[j] ^= data[i]
				data[i] ^= data[j]
			}
		}
	}
}

func SelectSort(data []int, left, right int) {
	if left >= right {
		return
	}
	for i := left; i < right; i++ {
		val := data[i]
		index := i
		for j := i + 1; j <= right; j++ {
			if data[j] < val {
				val = data[j]
				index = j
			}
		}
		data[i], data[index] = data[index], data[i]
	}
}

func InsertSort(data []int, left, right int) {
	if left >= right {
		return
	}
	for i := left; i <= right; i++ {
		for j := i; j > 0 && data[j] < data[j-1]; j-- {
			data[j] += data[j-1]
			data[j-1] = data[j] - data[j-1]
			data[j] = data[j] - data[j-1]
		}
	}
}

func qSort(data []int, left, right int) {
	if left < right {
		l, r := left, right
		mid := data[l+((r-l)>>1)]
		for l < r {
			for l < r && data[l] < mid {
				l++
			}
			for l < r && data[r] > mid {
				r--
			}
			if l >= r {
				break
			}
			if data[l] == data[r] && data[r] == mid {
				r--
			} else {
				data[l], data[r] = data[r], data[l]
			}
		}
		qSort(data, left, l-1)
		qSort(data, r+1, right)
	}
	return
}

func MergeSort(data []int, left, right int) {
	if left < right {
		mid := left + ((right - left) >> 1)
		MergeSort(data, left, mid)
		MergeSort(data, mid+1, right)
		Merge(data, left, mid, right)
	}
	return
}

func MergeSortByIteration(data []int, left, right int) {
	if left >= right {
		return
	}
	step := 1
	for step < right {
		fmt.Println("子数组长度: ", step)
		l := left
		for l <= right {
			r := min(right, l+2*step)
			mid := l + step
			if mid < right {
				fmt.Println(l, mid, r)
				Merge(data, l, mid, r)
			}
			l += 2 * step
		}
		step <<= 1

	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Merge(data []int, left, mid, right int) {
	for right > mid {
		if data[right] > data[mid] {
			right--
		} else {
			data[right], data[mid] = data[mid], data[right]
			for i := mid; i > left && data[i] < data[i-1]; i-- {
				data[i] ^= data[i-1]
				data[i-1] ^= data[i]
				data[i] ^= data[i-1]
			}
		}
	}
}

type Heap struct {
	size int
	data []int
}

func NewHeap(data []int) *Heap {
	heap := &Heap{
		size: len(data),
		data: data,
	}
	heap.buildHeap()
	return heap
}

func (h *Heap) buildHeap() {
	for i := h.size / 2; i > 0; i-- {
		h.downNode(i)
	}
}

func (h *Heap) downNode(parent int) {
	child := parent * 2
	for child < h.size {
		if child+1 < h.size && h.data[child+1] < h.data[child] {
			child++
		}
		if h.data[parent] <= h.data[child] {
			break
		}
		h.data[parent], h.data[child] = h.data[child], h.data[parent]
		parent = child
		child = parent * 2
	}
}

func (h *Heap) upNode(child int) {
	parent := child / 2
	for child >= 0 {
		if child+1 < h.size && h.data[child+1] < h.data[child] {
			child++
		}
		if h.data[parent] <= h.data[child] {
			break
		}
		h.data[child], h.data[parent] = h.data[parent], h.data[child]
		child = parent
		parent = child / 2
	}
}

func (h *Heap) Insert(val int) error {
	h.data = append(h.data, val)
	h.size++
	h.upNode(h.size - 1)
	return nil
}

func (h *Heap) Delete() (val int, err error) {
	if h.size == 0 {
		return -1, errors.New("heap is nil")
	}
	val = h.data[0]
	h.data[0] = h.data[h.size-1]
	h.data = h.data[:h.size-1]
	h.size--
	h.downNode(0)

	return
}

func HeapSort(data []int, left, right int) {
	h := NewHeap(data[left : right+1])
	for h.size > 0 {
		val, err := h.Delete()
		if err != nil {
			fmt.Println("heap error: ", err)
		}
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}

func ShellSort(data []int, left, right int) {
	gap := len(data) >> 1
	for gap > 0 {
		for i := gap; i <= right; i++ {
			for j := i; j-gap >= 0 && data[j-gap] > data[j]; j -= gap {
				data[j-gap], data[j] = data[j], data[j-gap]
			}
		}
		gap >>= 1
	}
}

func CountSort(data []int, left, right int) {
	max, min := -1<<31, 1<<31
	for _, v := range data[left : right+1] {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	tmp := make([]int, max-min+1)
	for _, v := range data[left : right+1] {
		tmp[v-min] += 1
	}
	index := 0
	for i, v := range tmp {
		for v > 0 {
			v--
			data[index] = i + min
			index++
		}
	}
}

func qSortByIterteration(data []int, left, right int) {
	if left < right {
		stack := NewStack(len(data))
		_ = stack.Push(left)
		_ = stack.Push(right)
		for !stack.IsEmpty() {
			t1, err := stack.Pop()
			if err != nil {
				fmt.Println(err)
				return
			}
			t2, err := stack.Pop()
			if err != nil {
				fmt.Println(err)
				return
			}
			l, ok := t2.(int)
			if !ok {
				fmt.Println("assert left error: ", ok)
			}
			r, ok := t1.(int)
			if !ok {
				fmt.Println("assert right error: ", ok)
			}
			if l >= r {
				return
			}
			mid := data[l+((r-l)>>1)]
			for l < r {
				for l < r && data[l] < mid {
					l++
				}
				for l < r && data[r] > mid {
					r--
				}
				if l >= r {
					break
				}
				if mid == data[r] && data[l] == data[r] {
					r--
				} else {
					data[l], data[r] = data[r], data[l]
				}
			}
			_ = stack.Push(left)
			_ = stack.Push(l - 1)
			_ = stack.Push(r + 1)
			_ = stack.Push(right)
		}
	}
	return
}

type Stack struct {
	size int
	cap  int
	data []interface{}
	top  int
}

func NewStack(cap int) *Stack {
	return &Stack{
		size: 0,
		top:  -1,
		data: make([]interface{}, cap),
		cap:  cap,
	}
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) IsFull() bool {
	return s.size == s.cap
}

func (s *Stack) Push(val interface{}) error {
	if s.IsFull() {
		return errors.New("stack is full")
	}
	s.top++
	s.data[s.top] = val
	s.size++
	return nil
}

func (s *Stack) Pop() (val interface{}, err error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	val = s.data[s.top]
	s.top--
	s.size--
	return
}

type Queue struct {
	size        int
	cap         int
	data        []interface{}
	front, rare int
}

func NewQueue(cap int) *Queue {
	return &Queue{
		size: 0,
		cap:  cap,
		data: make([]interface{}, cap),
	}
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) IsFull() bool {
	return q.size == q.cap
}

func (q *Queue) Push(val interface{}) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}
	q.data[q.front] = val
	q.front = (q.front + 1) % q.cap
	q.size++
	return nil
}

func (q *Queue) Pop() (val interface{}, err error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is nil")
	}
	val = q.data[q.front]
	q.front = (q.front + 1) % q.cap
	q.size--
	return
}

func main() {
	data := []int{1, 5655, 123, 222, 32, 23, 1234, 11, 11, 11, 16, 123, 13, 3123, 123442}
	fmt.Println("data len: ", len(data))
	fmt.Println("before data: ", data)
	//MergeSortByIteration(data, 0, len(data)-1)
	//qSort(data, 0, len(data)-1)
	//qSortByIterteration(data, 0, len(data)-1)
	//CountSort(data, 0, len(data)-1)
	//MergeSort(data, 0, len(data)-1)
	//HeapSort(data, 0, len(data)-1)
	//ShellSort(data, 0, len(data)-1)
	//BubblingSort(data, 0, len(data)-1)
	//SelectSort(data, 0, len(data)-1)
	//InsertSort(data, 0, len(data)-1)
	fmt.Println("sorted data: ", data)

}
