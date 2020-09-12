package heap

import "errors"

type Heap struct {
	size int
	data []interface{}
	flag int
	less func(a, b interface{}) bool
}

func NewHeap(data []interface{}, less func(a, b interface{})bool,flag int) *Heap {
	heap := &Heap{
		data: data,
		less: less,
		flag: flag,
	}

	heap.buildHeap()

	return heap
}

func (h *Heap) buildHeap() {
	for i := h.size>>1; i>= 0; i-- {
		h.downNode(i)
	}
}


func (h *Heap) downNode(parent int) {
	child := parent * 2
	for child < h.size {
		if h.flag == 1 { // minHeap
			if child + 1 < h.size && h.less(h.data[child+1], h.data[child]) {
				child++
			}
			if h.less(h.data[parent], h.data[child]) {
				break
			}
		} else {
			if child + 1 < h.size && h.less(h.data[child], h.data[child]) {
				child++
			}
			if h.less(h.data[child], h.data[parent]) {
				break
			}
		}
		h.data[child], h.data[parent] = h.data[parent], h.data[child]
		parent = child
		child = parent * 2
	}
}

func (h *Heap) upNode(child int) {
	parent := child >> 1
	for child > 0 {
		if h.flag == 1 {
			if child + 1 < h.size && h.less(h.data[child], h.data[child+1]) {
				child++
			}
			if h.less(h.data[parent], h.data[child]) {
				break
			}
		} else {
			if child +1 < h.size && h.less(h.data[child+1], h.data[child]) {
				child++
			}
			if h.less(h.data[child], h.data[parent]) {
				break
			}
		}
		h.data[child], h.data[parent] = h.data[parent], h.data[child]
		child = parent
		parent = child >> 1
	}
}

func (h *Heap) Insert(val interface{}) {
	h.size++
	h.data = append(h.data, val)
	h.upNode(h.size-1)
	return
}

func (h *Heap)  Delete() (val interface{}, err error) {
	if h.size == 0 {
		return nil, errors.New("heap is nil")
	}
	val = h.data[0]
	h.data[0] = h.data[h.size-1]
	h.data = h.data[:h.size-1]
	h.size--
	h.downNode(0)
	return
}