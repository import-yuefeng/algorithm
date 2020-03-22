package main

type BTreeNode struct {
	Val        []int
	childPoint []*BTreeNode
	capacity   int
	top        int
	parent     *BTreeNode
}

func NewBTree(capacity int) *BTreeNode {
	if capacity <= 0 {
		return nil
	}
	return &BTreeNode{
		Val:        make([]int, capacity),
		childPoint: make([]*BTreeNode, capacity),
		capacity:   capacity,
		top:        -1,
		parent:     nil,
	}
}

func (b *BTreeNode) NewBTreeNode() *BTreeNode {
	return &BTreeNode{
		Val:        make([]int, b.capacity),
		childPoint: make([]*BTreeNode, b.capacity),
		capacity:   b.capacity,
		top:        -1,
		parent:     nil,
	}
}

func (b *BTreeNode) Insert(val int) {
	if b.top+1 < b.capacity {
		b.Val[b.top+1] = val
		for i := b.top + 1; i > 0 && b.Val[i] < b.Val[i-1]; i-- {
			b.Val[i], b.Val[i-1] = b.Val[i-1], b.Val[i]
		}
		b.top++
		return
	}
	mid := b.top / 2
	lD, rD := b.Val[:mid], b.Val[mid:]
	if len(rD) == 1 {
		rD = []int{}
	} else {
		rD = b.Val[mid+1:]
	}

	left, right := b.NewBTreeNode(), b.NewBTreeNode()
	for i := 0; i < len(lD); i++ {
		left.Val[i] = lD[i]
	}
	for i := 0; i < len(rD); i++ {
		right.Val[i] = rD[i]
	}

}

func main() {

}
