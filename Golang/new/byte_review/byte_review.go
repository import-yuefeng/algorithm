package main

import (
    "errors"
	"fmt"
)

func qSort(data []int, left, right int) {
	if left < right {
		l, r := left, right
		mid := data[(l+r)/2]
		for l < r {
			if l < r && data[l] < mid {
				l++
			}
			if l < r && data[r] > mid {
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

type Stack struct {
	data     []interface{}
	size     int
	capacity int
	top      int
}

func NewStack(capacity int) *Stack {
	return &Stack{
		data:     make([]interface{}, capacity),
		size:     0,
		top:      -1,
		capacity: capacity,
	}
}

func (s *Stack) IsFull() bool {
	return s.size == s.capacity
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Pop() (val interface{}, err error) {
	if s.IsEmpty() {
		return -1, errors.New("stack is empty")
	}
	val = s.data[s.top]
	s.top--
	s.size--
	return val, nil
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

func qSort2(data []int, left, right int) {
	if left >= right {
		return
	}
	stack := NewStack(len(data))
	stack.Push(left)
	stack.Push(right)
	for !stack.IsEmpty() {
		t1, _ := stack.Pop()
		t2, _ := stack.Pop()
		r := t1.(int)
		l := t2.(int)
		if l >= r {
			continue
		}
		mid := data[(l+r)/2]
		for l < r {
			if data[l] < mid && l < r {
				l++
			}
			if data[r] > mid && l < r {
				r--
			}
			if l >= r {
				break
			}
			if data[l] == data[r] && data[r] == mid {
				r--
			} else {
				data[l] += data[r]
				data[r] = data[l] - data[r]
				data[l] = data[l] - data[r]
			}
		}
		stack.Push(left)
		stack.Push(l - 1)
		stack.Push(r + 1)
		stack.Push(right)
	}
	return
}

func MergeSort(data []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		MergeSort(data, left, mid)
		MergeSort(data, mid+1, right)
		Merge(data, left, mid, right)
	}
	return
}

func Merge(data []int, left, mid, right int) {
	//m := mid
	for right > mid {
		if data[right] > data[mid] {
			right--
		} else {
			data[mid], data[right] = data[right], data[mid]
			for i := mid; i > left && data[i] < data[i-1]; i-- {
				data[i], data[i-1] = data[i-1], data[i]

			}
			right--
		}
	}
	return
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

func (h *Heap) Delete() (int, error) {
	if h.size == 0 {
		return -1, errors.New("heap is empty")
	}
	val := h.data[0]
	h.data[0] = h.data[h.size-1]
	h.size--
	h.downNode(0)
	return val, nil
}

func (h *Heap) Insert(val int) error {
	h.data = append(h.data, val)
	h.size++
	h.upNode(h.size - 1)
	return nil
}

func (h *Heap) buildHeap() {
	for i := h.size / 2; i >= 0; i-- {
		h.downNode(i)
	}
	return
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
		h.data[child], h.data[parent] = h.data[parent], h.data[child]
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
		h.data[parent], h.data[child] = h.data[child], h.data[parent]
		child = parent
		parent = child / 2
	}
}

func (h *Heap) IsEmpty() bool {
	return h.size == 0
}

func HeapSort(data []int, left, right int) {
	h := NewHeap(data)
	for !h.IsEmpty() {
		val, _ := h.Delete()
		fmt.Printf("%d ", val)
	}
}

func BubblingSort(data []int, left, right int) {
	for i := left; i <= right; i++ {
		for j := i + 1; j <= right; j++ {
			if data[j] < data[i] {
				data[j], data[i] = data[i], data[j]
			}
		}
	}
}

func InsertSort(data []int, left, right int) {
	for j := left + 1; j <= right; j++ {
		for i := j; i > 0 && data[i] < data[i-1]; i-- {
			data[i], data[i-1] = data[i-1], data[i]
		}
	}
}

func SelectSort(data []int, left, right int) {
	for i := 0; i < right; i++ {
		tmp := data[i]
		index := i
		for j := i + 1; j <= right; j++ {
			if data[j] < tmp {
				index = j
				tmp = data[j]
			}
		}
		data[index], data[i] = data[i], data[index]
	}
}

func ShellSort(data []int, left, right int) {
	gap := len(data) >> 1
	for gap > 0 {
		for j := gap; j <= right; j++ {
			for i := j; i-gap > 0 && data[i] < data[i-gap]; i -= gap {
				data[i], data[i-gap] = data[i-gap], data[i]
			}
		}
		gap >>= 1
	}
}

func CountSort(data []int, left, right int) {
	max := -1 << 31
	min := 1 << 31
	for _, v := range data {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}
	tmp := make([]int, max-min+1)
	for _, v := range data {
		tmp[v-min] += 1
	}
	index := 0
	for i, v := range tmp {
		for v > 0 {
			data[index] = i + min
			index++
			v--
		}
	}
	return
}

func qSelect(data []int, left, right, k int) int {
	if k > len(data) || k <= 0 {
		return -1 << 31
	}
	if left < right {
		l, r := left, right
		mid := data[(l+r)>>1]
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
		if k == l-left+1 {
			return mid
		} else if k < l-left+1 {
			return qSelect(data, left, l-1, k)
		} else {
			return qSelect(data, r+1, right, k-l+left-1)
		}
	}
	return data[left]
}

type BinaryTree struct {
	Left, Right *BinaryTree
	Val         int
}

func NewBinaryTree(val int, left, right *BinaryTree) *BinaryTree {
	return &BinaryTree{}
}

type Queue struct {
	size        int
	front, rare int
	capacity    int
	data        []interface{}
}

func NewQueue(capacity int) *Queue {
	return &Queue{
		size:     0,
		front:    0,
		rare:     0,
		capacity: capacity,
		data:     make([]interface{}, capacity),
	}
}

func (q *Queue) IsFull() bool {
	return q.size == q.capacity
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Push(val interface{}) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}
	q.data[q.rare] = val
	q.rare = (q.rare + 1) % q.capacity
	q.size++
	return nil
}

func (q *Queue) Pop() (val interface{}, err error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	val = q.data[q.front]
	q.front = (q.front + 1) % q.capacity
	return val, nil
}

func preOrder(root *BinaryTree) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	preOrder(root.Left)
	preOrder(root.Right)
}

func inOrder(root *BinaryTree) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	fmt.Printf("%d ", root.Val)
	inOrder(root.Right)
}

func postOrder(root *BinaryTree) {
	if root == nil {
		return
	}
	postOrder(root.Left)
	postOrder(root.Right)
	fmt.Printf("%d ", root.Val)
}

func preOrder2(root *BinaryTree) {
	stack := NewStack(100)
	stack.Push(root)
	isVisit := make(map[*BinaryTree]struct{})
	for !stack.IsEmpty() {
		tmp, _ := stack.Pop()
		node := tmp.(*BinaryTree)
		if node == nil {
			continue
		}
		if _, exist := isVisit[node]; exist {
			// 节点变为黑色, 三色标记可输出
			fmt.Printf("%d ", node.Val)
		} else {
			// 节点从白色变为灰色
			isVisit[node] = struct{}{}
			stack.Push(node.Right)
			stack.Push(node.Left)
			stack.Push(node)
		}
	}
}

func inOrder2(root *BinaryTree) {
	stack := NewStack(100)
	stack.Push(root)
	isVisit := make(map[*BinaryTree]struct{})
	for !stack.IsEmpty() {
		tmp, _ := stack.Pop()
		node := tmp.(*BinaryTree)
		if node == nil {
			continue
		}
		if _, exist := isVisit[node]; exist {
			// 三色变黑, 输出
			fmt.Printf("%d ", node.Val)
		} else {
			isVisit[node] = struct{}{}
			stack.Push(node.Right)
			stack.Push(node)
			stack.Push(node.Left)
		}
	}
}

func postOrder2(root *BinaryTree) {
	stack := NewStack(100)
	stack.Push(root)
	isVisit := make(map[*BinaryTree]struct{})
	for !stack.IsEmpty() {
		tmp, _ := stack.Pop()
		node := tmp.(*BinaryTree)
		if node == nil {
			continue
		}
		if _, exist := isVisit[node]; exist {
			// 三色变黑, 输出
			fmt.Printf("%d ", node.Val)
		} else {
			isVisit[node] = struct{}{}
			stack.Push(node)
			stack.Push(node.Right)
			stack.Push(node.Left)
		}
	}
}

func inOrder3(root *BinaryTree) {
	if root == nil {
		return
	}
	for root != nil {
		if root.Left == nil {
			fmt.Printf("%d ", root.Val)
			root = root.Right
		} else {
			tmp := root.Left
			for tmp.Right != nil && tmp.Right != root {
				tmp = tmp.Right
			}
			if tmp.Right == nil {
				tmp.Right = root
				root = root.Left
			} else {
				tmp.Right = nil
				fmt.Printf("%d ", root.Val)
				root = root.Right
			}
		}
	}
}

func orderByLevel(root *BinaryTree) {
	q := NewQueue(100)
	q.Push(root)
	for !q.IsEmpty() {
		size := q.size
		for size > 0 {
			tmp, _ := q.Pop()
			size--
			node := tmp.(*BinaryTree)
			if node == nil {
				continue
			}
			fmt.Printf("%d ", node.Val)
			q.Push(node.Left)
			q.Push(node.Right)
		}
	}
}

func orderZByLevel(root *BinaryTree) {
	q := NewQueue(1000)
	q.Push(root)
	flag := 0
	for !q.IsEmpty() {
		vals := make([]int, 0)
		size := q.size
		for size > 0 {
			tmp, _ := q.Pop()
			node, _ := tmp.(*BinaryTree)
			size--
			if node == nil {
				continue
			}
			vals = append(vals, node.Val)
			q.Push(node.Left)
			q.Push(node.Right)
		}
		if flag == 0 {
			fmt.Println(vals)
			flag = -1
		} else if flag == -1 {
			reverseVal(vals, 0, len(vals)-1)
			fmt.Println(vals)
			flag = 0
		}
	}
}

func ReConstructBinaryTree(preList, inList []int) *BinaryTree {
	if len(preList) != len(inList) || len(preList) < 1 {
		return nil
	}
	root := new(BinaryTree)
	root.Val = preList[0]
	index := 0
	for i, v := range inList {
		if v == root.Val {
			index = i
			break
		}
	}
	if inList[index] != root.Val {
		return nil
	}
	if index > 0 {
		root.Left = ReConstructBinaryTree(preList[1:index+1], inList[0:index])
	}
	if index < len(inList)-1 {
		root.Right = ReConstructBinaryTree(preList[index+1:], inList[index+1:])
	}
	return root
}

func NumberOf1(n int) int {
	count := 0
	for n != 0 {
		if n&1 == 1 {
			count++
		}
		n >>= 1
	}
	return count
}

func NumberOf1_2(n int) int {
	count := 0
	for n != 0 {
		n &= (n - 1)
		count++
	}
	return count
}

func reverseVal(data []int, left, right int) {
	for left < right {
		data[left], data[right] = data[right], data[left]
		left++
		right--
	}
}

type LinkedList struct {
	Val  interface{}
	Next *LinkedList
}

func ReverseLinkedList(root *LinkedList) *LinkedList {
	if root == nil || root.Next == nil {
		return root
	}
	last := ReverseLinkedList(root.Next)
	root.Next.Next = root
	root.Next = nil
	return last
}

func ReverseLinkedList2(root *LinkedList) *LinkedList {
	if root == nil || root.Next == nil {
		return root
	}
	prev, cur, next := root, root.Next, root.Next.Next
	for cur != nil {
		cur.Next = prev
		prev = cur
		cur = next
		if next != nil {
			next = next.Next
		}
	}
	root.Next = nil
	return prev
}

func ReverseLinkedListB(root *LinkedList, b *LinkedList) *LinkedList {
	if root == nil || root.Next == nil {
		return root
	}
	prev, cur, next := root, root.Next, root.Next.Next
	for cur != b {
		cur.Next = prev
		prev = cur
		cur = next
		if next != nil {
			next = next.Next
		}
	}
	root.Next = b
	return prev
}

func ReverseLinkedListK(root *LinkedList, k int) *LinkedList {
	if root == nil || k <= 0 {
		return root
	}
	a, b := root, root

	for i := 0; i < k; i++ {
		if b == nil {
			return root
		}
		b = b.Next
	}
	newHead := ReverseLinkedListB(a, b)
	a.Next = ReverseK(b, k)
	return newHead

}

var end *LinkedList

func ReverseK(root *LinkedList, k int) *LinkedList {
	if root == nil || k <= 0 {
		return root
	}
	if k == 1 {
		end = root.Next
		return root
	}
	last := ReverseK(root.Next, k-1)
	root.Next.Next = root
	root.Next = end
	return last
}

func ReverseM2N(root *LinkedList, m, n int) *LinkedList {
	if m < 0 || n < 0 || root == nil {
		return root
	}
	if m == 1 {
		return ReverseK(root, n)
	}
	root.Next = ReverseM2N(root.Next, m-1, n-1)
	return root
}

func MergeSortedLinkedList(a, b *LinkedList) *LinkedList {
	if a == nil {
		return b
	} else if b == nil {
		return a
	}
	root := new(LinkedList)
	if a.Val.(int) < b.Val.(int) {
		root.Val = a.Val
		root.Next = MergeSortedLinkedList(a.Next, b)
	} else {
		root.Val = b.Val
		root.Next = MergeSortedLinkedList(a, b.Next)
	}
	return root
}

func MergeSortedLinkedList2(a, b *LinkedList) *LinkedList {
	if a == nil {
		return b
	} else if b == nil {
		return a
	}
	root := new(LinkedList)
	cur := root
	front := root
	for a != nil && b != nil {
		if a.Val.(int) < b.Val.(int) {
			cur.Val = a.Val
			cur.Next = new(LinkedList)
			a = a.Next
		} else {
			cur.Val = b.Val
			cur.Next = new(LinkedList)
			b = b.Next
		}
		front = cur
		cur = cur.Next
	}
	if a == nil {
		front.Next = b
	} else {
		front.Next = a
	}
	return root
}

func kNumInLinkedList(root *LinkedList, k int) interface{} {
	if isCycleLinkedList(root) {
		return -1
	}
	slow, fast := root, root
	for fast != nil && k > 0 {
		fast = fast.Next
		k--
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Val
}

func isCycleLinkedList(root *LinkedList) bool {
	slow, fast := root, root
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func CycleLinkedListEntry(root *LinkedList) *LinkedList {
	slow, fast := root, root
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		slow = slow.Next
		if slow == fast {
			fast = slow
			slow = root
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}

func findNotRepeatNumber(data []int) int {
	val := 0
	for _, v := range data {
		val ^= v
	}
	return val
}

func findRepeatNumber(data []int) int {
	val := 0
	for i, v := range data {
		val ^= i
		val ^= v
	}
	return val
}

func singleNumbers(data []int) (num1, num2 int) {
	val := 0
	for _, v := range data {
		val ^= v
	}
	tmp := val
	index := 0
	for tmp&1 == 0 {
		index++
		tmp >>= 1
	}
	num1, num2 = 0, 0
	for _, v := range data {
		c := v >> index
		if c&1 == 1 {
			num1 ^= v
		} else {
			num2 ^= v
		}
	}
	return
}

func findMax(data []int) int {
	val, count := 0, 0
	for _, v := range data {
		if count == 0 {
			val = v
			count++
		} else if val == v {
			count++
		} else if val != v {
			count--
		}
	}
	return val
}

func VerifySequenceOfBST(postList []int) bool {
	if len(postList) <= 1 {
		return true
	}
	root := postList[len(postList)-1]
	index := 0
	for i, v := range postList {
		if v > root {
			index = i
		}
	}

	for _, v := range postList[index:] {
		if v < root {
			return false
		}
	}
	left, right := true, true
	if index > 0 {
		left = VerifySequenceOfBST(postList[:index])
	}
	if index < len(postList) {
		right = VerifySequenceOfBST(postList[index : len(postList)-1])
	}
	return right && left
}

func main() {
	num := []int{1, 5, 100, 100, 12, 4, 2, 111}
	data := num
	CountSort(num, 0, 6)
	kN := qSelect(data, 0, len(data)-1, 5)
	fmt.Printf("%#v", num)
	fmt.Println()
	fmt.Println("kN: ", kN)

	val := findNotRepeatNumber([]int{1, 1, 2, 2, 5, 6, 6, 100, 100})
	fmt.Println("Not repeat val: ", val)
	val1, val2 := singleNumbers([]int{1, 1, 2, 2, 5, 6, 6, 11, 100, 100})
	fmt.Println("not repeat: ", val1, val2)
	val = findMax([]int{2, 3, 4, 5, 6, 2, 2, 2, 2, 2, 2, 2, 2})
	fmt.Println("maxVal: ", val)
	val = findRepeatNumber([]int{1, 2, 2, 3, 4, 5})
	fmt.Println("repeat val: ", val)
	flag := VerifySequenceOfBST([]int{1, 6, 3, 60, 75, 70, 50})
	fmt.Println("flag: ", flag)
}
