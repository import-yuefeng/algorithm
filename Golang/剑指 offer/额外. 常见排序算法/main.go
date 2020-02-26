package main

import (
	"errors"
	"fmt"
)

func main() {
	testVal := []int{1, 5, 5, 2, 3, 5, 3, 2, 1444, 123212312434, 12412, 4242}
	// BubblingSort(testVal)
	// SelectSort(testVal)
	// InsertSort(testVal)
	// qSort(testVal, 0, len(testVal)-1)
	// MergeSort(testVal, 0, len(testVal)-1)
	// ShellSort(testVal)
	// HeapSort(testVal)
	// CountSort(testVal)
	// RadixSort(testVal)
	qSort_2(testVal, 0, len(testVal)-1)
	fmt.Println(testVal)
}

func BubblingSort(num []int) {
	for i := 0; i < len(num)-1; i++ {
		for j := 0; j < len(num)-1-i; j++ {
			if num[j] > num[j+1] {
				num[j] ^= num[j+1]
				num[j+1] ^= num[j]
				num[j] ^= num[j+1]
			}
		}
	}
}

func SelectSort(num []int) {
	for i := 0; i < len(num)-1; i++ {
		minVal := num[i]
		minValIndex := i
		for j := i; j < len(num); j++ {
			if minVal > num[j] {
				minVal = num[j]
				minValIndex = j
			}
		}
		num[i], num[minValIndex] = num[minValIndex], num[i]
	}
}

func InsertSort(num []int) {
	for i := 0; i < len(num); i++ {
		for j := i; j > 0 && num[j] < num[j-1]; j-- {
			num[j] += num[j-1]
			num[j-1] = num[j] - num[j-1]
			num[j] -= num[j-1]
		}
	}
}

func ShellSort(num []int) {
	gap := len(num) >> 1
	for gap > 0 {
		for i := gap; i < len(num); i++ {
			for j := i; j-gap >= 0 && num[j] < num[j-gap]; j -= gap {
				num[j], num[j-gap] = num[j-gap], num[j]
			}
		}
		gap /= 2
	}
}

func qSort(num []int, left, right int) {
	if left < right && left >= 0 {
		l, r := left, right
		mid := num[left+((right-left)>>1)]
		for l < r {
			for l < r && num[l] < mid {
				l++
			}
			for l < r && num[r] > mid {
				r--
			}
			if l >= r {
				break
			}
			if num[l] == num[r] && num[r] == mid {
				r--
			} else {
				num[r] ^= num[l]
				num[l] ^= num[r]
				num[r] ^= num[l]
			}
		}
		qSort(num, left, l-1)
		qSort(num, r+1, right)
	}
}

func MergeSort(num []int, left, right int) {
	if left < right && left >= 0 {
		mid := (left + right) >> 1
		MergeSort(num, left, mid)
		MergeSort(num, mid+1, right)
		Merge(num, left, mid, right)
	}
}

func Merge(num []int, left, mid, right int) {
	p1, p2 := mid, right
	for p2 > mid {
		if num[p2] < num[p1] {
			num[p2], num[p1] = num[p1], num[p2]
		}
		p2--
		for i := mid; i > left && num[i] < num[i-1]; i-- {
			num[i-1], num[i] = num[i], num[i-1]
		}
	}
}

func HeapSort(num []int) {
	heap := NewHeap(len(num), 0)
	heap.buildHeap(num)
	for !heap.IsEmpty() {
		if x, err := heap.Pop(); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Printf("%d ", x)
		}
	}
}

type Heap struct {
	size     int
	data     []int
	capacity int
	flag     int
}

func NewHeap(capacity int, flag int) *Heap {
	return &Heap{
		size:     0,
		data:     make([]int, capacity),
		capacity: capacity,
		flag:     flag,
	}
}

func (h *Heap) Push(x int) error {
	if h.size == h.capacity {
		return errors.New("Heap is full")
	}
	h.data[h.size] = x
	h.upNode(h.size)
	h.size++
	return nil
}

func (h *Heap) Pop() (x int, err error) {
	if h.size == 0 {
		return -1, errors.New("heap is empty")
	}
	x = h.data[0]
	h.data[0] = h.data[h.size-1]
	h.size--
	h.downNode(0)
	return x, nil
}

func (h *Heap) IsEmpty() bool {
	return h.size == 0
}

func (h *Heap) downNode(i int) {
	if h.size == 0 {
		return
	}
	parent := i
	child := i*2 + 1
	for child < h.size {
		if h.flag == 0 {
			if child+1 < h.size && h.data[child+1] < h.data[child] {
				child++
			}
			if h.data[parent] <= h.data[child] {
				break
			}
		} else {
			if child+1 < h.size && h.data[child+1] > h.data[child] {
				child++
			}
			if h.data[parent] >= h.data[child] {
				break
			}
		}
		h.data[parent], h.data[child] = h.data[child], h.data[parent]
		parent = child
		child = parent*2 + 1
	}
}

func (h *Heap) upNode(i int) {
	if h.size == 0 {
		return
	}
	child := i
	parent := i / 2
	for child >= 0 {
		if h.flag == 0 {
			if child+1 < h.size && h.data[child+1] < h.data[child] {
				child++
			}
			if h.data[parent] <= h.data[child] {
				break
			}
		} else {
			if child+1 < h.size && h.data[child+1] > h.data[child] {
				child++
			}
			if h.data[parent] >= h.data[child] {
				break
			}
		}
		h.data[parent], h.data[child] = h.data[child], h.data[parent]
		child = parent
		parent = child / 2
	}
}

func (h *Heap) buildHeap(num []int) {
	h.data = num
	h.size = len(num)
	for i := h.size / 2; i >= 0; i-- {
		h.downNode(i)
	}
}

func CountSort(num []int) {
	max, min := -1<<31, 1<<31
	for _, v := range num {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	length := max - min + 1
	array := make([]int, length)
	for _, v := range num {
		array[v-min] += 1
	}
	p := 0
	for i, v := range array {
		for v > 0 {
			// v - min = i
			num[p] = i + min
			v--
			p++
		}
	}
}

func RadixSort(num []int) {
	bucket := make([]*Queue, 10)
	for i, _ := range bucket {
		bucket[i] = NewQueue(len(num))
	}
	maxVal := -1 << 31
	for _, v := range num {
		if maxVal < v {
			maxVal = v
		}
	}
	maxValLength := 0
	for maxVal != 0 {
		maxVal /= 10
		maxValLength++
	}
	tmp := make([]int, len(num))
	copy(tmp, num)
	for i := 1; i <= maxValLength; i++ {
		for _, v := range tmp {
			bucket[GetCurBit(v, i)].Push(v)
		}
		count := 0
		for _, v := range bucket {
			for !v.IsEmpty() {
				val, _ := v.Pop()
				tmp[count] = val
				count++
			}
		}
	}
	copy(num, tmp)
}

func GetCurBit(val, bit int) int {
	divisor := 1
	for bit > 1 {
		divisor *= 10
		bit--
	}
	return val / divisor % 10
}

type Queue struct {
	rare, front int
	data        []int
	size        int
	capacity    int
}

func NewQueue(capacity int) *Queue {
	return &Queue{
		front:    0,
		rare:     0,
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (q *Queue) IsEmpty() bool {
	return q.rare == q.front
}

func (q *Queue) IsFull() bool {
	return (q.rare+1)%q.capacity == q.front
}

func (q *Queue) Push(x int) error {
	if q.IsFull() {
		return errors.New("Queue is full")
	}
	q.data[q.rare] = x
	q.rare = (q.rare + 1) % q.capacity
	q.size++
	return nil
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

func qSort_2(num []int, left, right int) {
	if left < right && left >= 0 {
		mid := num[(left+right)>>1]
		l, r := left, left+1
		for r <= right {
			for l <= r && num[l] < mid {
				l++
			}
			for r <= right && num[r] > mid {
				r++
			}
			if r > right {
				break
			}
			if num[l] == num[r] && num[r] == mid {
				r++
			} else {
				num[l], num[r] = num[r], num[l]
			}
		}
		qSort(num, left, l-1)
		qSort(num, l, right)
	}
}
