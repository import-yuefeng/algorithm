package main

import (
	"errors"
	"fmt"
)

func main() {
	fequency, err := runSum("|.,@#${}:\">?<ancncncncncaskljfdlkshjgl478847326450009&(*&(*%*&^*(@&*#%#)));kajsl;krjfads;hgkjbvhsouaehfjkhjldsahfdjkhfkshklvbjhsgdfy")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fequency)
	root := &Huffman{
		huffmanNode: Init(fequency),
	}
	root.Build()
	// root
}

type Huffman struct {
	huffmanNode []*HuffmanNode
	root        *HuffmanNode
}

func (h *Huffman) Build() {
	heap := new(Heap)
	heap.Init(h.huffmanNode, len(h.huffmanNode))
	for !heap.isEmpty() {
		if heap.Size() > 1 {
			left, err := heap.Delete()
			if err != nil {
				fmt.Println(err)
				return
			}
			right, err := heap.Delete()
			if err != nil {
				fmt.Println(err)
				return
			}
			parent := new(HuffmanNode)
			parent.Fequency = left.Fequency + right.Fequency
			parent.Val = 0
			parent.Left, parent.Right = left, right
			err = heap.Insert(parent)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			root, err := heap.Delete()
			if err != nil {
				fmt.Println(err)
				return
			}
			h.root = root
			return
		}

	}
}

type HuffmanNode struct {
	Val         int
	Fequency    int
	Left, Right *HuffmanNode
}

func Init(fequency []int) []*HuffmanNode {
	NodeList := []*HuffmanNode{}
	for i, v := range fequency {
		node := new(HuffmanNode)
		node.Fequency = v
		node.Val = i
	}
	return NodeList
}

func runSum(text string) (fequency []int, err error) {
	if len(text) == 0 {
		return nil, errors.New("text is empty")
	}
	fequency = make([]int, 256)
	for _, v := range text {
		fequency[v]++
	}

	return fequency, nil
}

type Heap struct {
	array    []*HuffmanNode
	size     int
	capacity int
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) Init(array []*HuffmanNode, capacity int) error {
	if capacity <= 0 || len(array) > capacity {
		return errors.New("capacity is invaild or array is too biggest")
	}
	h.array = make([]*HuffmanNode, capacity)
	h.capacity = capacity
	h.size = len(array)
	for i, v := range array {
		h.array[i] = v
	}
	return nil
}

func (h *Heap) Build(option int) {
	switch option {
	case 0:
		h.BuildMinHeap()
	case 1:
		h.BuildMaxHeap()
	}
	return
}

func (h *Heap) isEmpty() bool {
	return h.size == 0
}

func (h *Heap) Insert(val *HuffmanNode) error {
	if h.size == h.capacity {
		return errors.New("heap is full")
	}
	h.array[h.size] = val
	h.size++
	h.updateMinNode()
	// h.updateMaxNode()
	return nil
}

func (h *Heap) Delete() (val *HuffmanNode, err error) {
	if h.size == 0 {
		return nil, errors.New("heap is empty")
	}
	val = h.array[0]
	err = nil
	h.array[0] = h.array[h.size-1]
	h.size--
	h.buildMinHeap(0)
	// h.buildMaxHeap(0)
	return
}

func (h *Heap) updateMinNode() {
	if h.size <= 1 {
		return
	}

	child := h.size - 1
	parent := (child - 1) / 2
	for child >= 0 {
		if h.array[parent].Fequency > h.array[child].Fequency {
			h.array[child], h.array[parent] = h.array[parent], h.array[child]
		}
		child = parent
		parent = (child - 1) / 2
	}
	return
}

func (h *Heap) updateMaxNode() {
	if h.size <= 1 {
		return
	}
	child := h.size - 1
	parent := (child - 1) / 2
	for child >= 0 {
		if h.array[child].Fequency > h.array[parent].Fequency {
			h.array[child], h.array[parent] = h.array[parent], h.array[child]
		}
		child = parent
		parent = (child - 1) / 2
	}
}

func (h *Heap) BuildMinHeap() {
	if h.size <= 1 {
		return
	}
	for i := h.size / 2; i >= 0; i-- {
		h.buildMinHeap(i)
	}
}

func (h *Heap) buildMinHeap(parent int) {
	child := parent*2 + 1
	for child < h.size {
		if child+1 < h.size && h.array[child+1].Fequency < h.array[child].Fequency {
			child++
		}
		if h.array[child].Fequency >= h.array[parent].Fequency {
			break
		}
		h.array[child], h.array[parent] = h.array[parent], h.array[child]
		parent = child
		child = parent*2 + 1
	}
}

func (h *Heap) BuildMaxHeap() {
	if h.size <= 1 {
		return
	}
	for i := h.size / 2; i >= 0; i-- {
		h.buildMaxHeap(i)
	}
}

func (h *Heap) buildMaxHeap(parent int) {
	child := parent*2 + 1
	for h.size > child {
		if child+1 < h.size && h.array[child].Fequency < h.array[child+1].Fequency {
			child++
		}
		if h.array[child].Fequency <= h.array[parent].Fequency {
			break
		}
		h.array[child], h.array[parent] = h.array[parent], h.array[child]
		parent = child
		child = parent*2 + 1
	}
}

type Queue struct {
	capacity    int
	size        int
	front, rare int
	data        []interface{}
}

func (q *Queue) Init(capacity int) error {
	if capacity <= 0 {
		return errors.New("capacity is invaild")
	}
	q.capacity = capacity
	q.size = 0
	q.front, q.rare = 0, 0
	q.data = make([]interface{}, q.capacity)
	return nil
}

func (q *Queue) Push(val interface{}) error {
	if q.size == q.capacity {
		return errors.New("queue is full")
	}
	q.rare = (q.rare + 1) % q.capacity
	q.data[q.rare] = val
	q.size++
	return nil
}

func (q *Queue) Pop() (val interface{}, err error) {
	if q.size == 0 {
		return nil, errors.New("queue is empty")
	}
	q.front = (q.front + 1) % q.capacity
	q.size--
	val = q.data[q.front]
	err = nil
	return val, err
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

type Stack struct {
	top      int
	size     int
	data     []interface{}
	capacity int
}

func (s *Stack) Init(capacity int) error {
	if capacity <= 0 {
		return errors.New("capacity is invaild")
	}
	s.top, s.size = -1, 0
	s.capacity = capacity
	s.data = make([]interface{}, s.capacity)
	return nil
}

func (s *Stack) Push(val interface{}) error {
	if s.size == s.capacity {
		return errors.New("stack is full")
	}

	s.top++
	s.data[s.top] = val
	s.size++
	return nil
}

func (s *Stack) Pop() (val interface{}, err error) {
	if s.size == 0 {
		return nil, errors.New("stack is empty")
	}
	val = s.data[s.top]
	s.size--
	s.top--
	err = nil
	return
}

func (s *Stack) IsEmpty() bool {
	return true
}
