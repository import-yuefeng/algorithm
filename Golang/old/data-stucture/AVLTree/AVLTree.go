package main

// import (
// 	"errors"
// 	"fmt"
// )

// type AVLTreeNode struct {
// 	value int
// 	right *AVLTreeNode
// 	left  *AVLTreeNode
// 	hight int
// }

// func (t *AVLTreeNode) Init() {
// 	t.value = 1
// 	t.right, t.left = nil, nil
// 	t.hight = 0
// }

// func (t *AVLTreeNode) Insert(value int) error {
// 	if t == nil {
// 		return errors.New("AVLTreeNode is nil")
// 	}
// 	tmp := t
// 	newElement := new(AVLTreeNode)
// 	newElement.right, newElement.left = nil, nil
// 	newElement.value = value
// 	newElement.hight = 0
// 	for tmp != nil {
// 		if tmp.value > value {
// 			if tmp.left == nil {
// 				tmp.left = newElement
// 				tmp.left.hight = tmp.hight + 1
// 				if tmp.left.hight-tmp.right.hight == 2 {
// 					if value >= tmp.left.value {
// 						DoubleRotateLR(tmp)
// 					} else if tmp.left.value > value {
// 						SingRotateLeft(tmp)
// 					}
// 				}
// 				break
// 			}
// 			tmp = tmp.left
// 		} else if tmp.value <= value {
// 			if tmp.right == nil {
// 				tmp.right = newElement
// 				tmp.right.hight = tmp.hight + 1
// 				if tmp.right.hight-tmp.left.hight == 2 {
// 					if tmp.right.value > value {
// 						DoubleRotateRL(tmp)
// 					} else if tmp.right.value < value {
// 						SingRotateRight(tmp)
// 					}
// 				}
// 				break
// 			}
// 			tmp = tmp.right
// 		}
// 	}
// 	return nil
// }

// func findMaxNode(t *AVLTreeNode) (node *AVLTreeNode, err error) {
// 	if t == nil {
// 		return nil, errors.New("AVLTreeNode is nil")
// 	}
// 	for t != nil {
// 		if t.right != nil {
// 			t = t.right
// 		} else {
// 			return t, nil
// 		}
// 	}
// 	return nil, nil
// }

// func (t *AVLTreeNode) findFrontNode(node *AVLTreeNode) *AVLTreeNode {
// 	tmp1, tmp2 := t, t
// 	for tmp1 != nil {
// 		if tmp1.value > node.value {
// 			tmp2 = tmp1
// 			tmp1 = tmp1.left
// 		} else if tmp1.value < node.value {
// 			tmp2 = tmp1
// 			tmp1 = tmp1.right
// 		} else {
// 			return tmp2
// 		}
// 	}
// 	return nil
// }

// func Delete2(t *AVLTreeNode, value int) {
// 	if t == nil {
// 		return
// 	}
// 	if t.value > value {
// 		Delete2(t.left, value)
// 		if t.right.hight-t.left.hight == 2 {
// 			if t.right.left != nil && t.right.left.hight-t.right.right.hight > 0 {
// 				DoubleRotateRL(t)
// 			} else {
// 				SingRotateRight(t)
// 			}
// 		}
// 	} else if t.value < value {
// 		Delete2(t.right, value)
// 		if t.left.hight-t.right.hight == 2 {
// 			if t.left.right != nil && t.left.right.hight-t.left.left.hight > 0 {
// 				DoubleRotateLR(t)
// 			} else {
// 				SingRotateLeft(t)
// 			}
// 		}
// 	} else {
// 		if t.left != nil && t.right != nil {
// 			node, err := findMaxNode(t.left)
// 			if err != nil {
// 				return
// 			}
// 			t.value = node.value
// 			Delete2(t.left, node.value)
// 			if t.right.hight-t.left.hight == 2 {
// 				if t.right.right != nil && t.right.right.hight-t.right.left.hight > 0 {
// 					SingRotateRight(t)
// 				} else {
// 					DoubleRotateRL(t)
// 				}
// 			}
// 		} else {
// 			if t.left == nil {
// 				t = t.right
// 			} else if t.right == nil {
// 				t = t.left
// 			}
// 		}
// 	}
// 	if t == nil {
// 		return
// 	}
// 	t.hight = max(t.right.hight, t.left.hight)
// 	return
// }

// func (t *AVLTreeNode) Delete(value *AVLTreeNode) error {
// 	if t.left == nil && t.right == nil {
// 		return errors.New("AVLTreeNode is empty")
// 	}
// 	frontValue := t.findFrontNode(value)
// 	if frontValue == nil {
// 		return nil
// 	}
// 	maxElement := new(AVLTreeNode)

// 	if value.left == nil {
// 		maxElement = value.right
// 	} else if value.right == nil {
// 		maxElement = value.left
// 	} else {
// 		tmp, err := findMaxNode(value.left)
// 		if err != nil {
// 			return errors.New("delete error")
// 		}
// 		maxElement.value = tmp.value
// 		t.Delete(tmp)
// 		maxElement.left, maxElement.right = value.left, value.right
// 		maxElement.hight = max(maxElement.left.hight, maxElement.right.hight) - 1
// 	}

// 	if frontValue.right == value {
// 		frontValue.right = maxElement
// 	} else if frontValue.left == value {
// 		frontValue.left = maxElement
// 	}
// 	return nil
// }

// func PriorPrintHight(tmp *AVLTreeNode) {
// 	if tmp == nil {
// 		return
// 	}
// 	fmt.Printf("%d %d\n", tmp.value, tmp.hight)
// 	PriorPrintHight(tmp.left)
// 	PriorPrintHight(tmp.right)
// }

// func PriorPrint(tmp *AVLTreeNode) {
// 	if tmp == nil {
// 		return
// 	}
// 	fmt.Printf("%d ", tmp.value)
// 	PriorPrint(tmp.left)
// 	PriorPrint(tmp.right)
// }

// func MiddlePrint(tmp *AVLTreeNode) {
// 	if tmp == nil {
// 		return
// 	}
// 	MiddlePrint(tmp.left)
// 	fmt.Printf("%d ", tmp.value)
// 	MiddlePrint(tmp.right)
// }

// func (t *AVLTreeNode) Find(value int) bool {
// 	tmp := t
// 	for tmp != nil {
// 		if tmp.value == value {
// 			return true
// 		} else if tmp.value > value {
// 			tmp = tmp.left
// 		} else {
// 			tmp = tmp.right
// 		}
// 	}
// 	return false
// }

// func (t *AVLTreeNode) FindNode(value int) *AVLTreeNode {
// 	tmp := t
// 	for tmp != nil {
// 		if tmp.value == value {
// 			return tmp
// 		} else if tmp.value > value {
// 			tmp = tmp.left
// 		} else {
// 			tmp = tmp.right
// 		}
// 	}
// 	return nil
// }

// func Find2(value int, t *AVLTreeNode) bool {
// 	if t == nil {
// 		return false
// 	} else if t.value == value {
// 		return true
// 	} else if t.value < value {
// 		return Find2(value, t.right)
// 	} else {
// 		return Find2(value, t.left)
// 	}
// }

// func (t *AVLTreeNode) LevelPrint() {
// 	if t == nil {
// 		return
// 	}
// 	tmp := t
// 	q := new(Queue)
// 	q.Init(100) // hard-code capacity
// 	q.Push(tmp)
// 	for !q.isEmpty() {
// 		value, err := q.Pop()
// 		if err != nil {
// 			break
// 		}
// 		fmt.Println(value)
// 		if value.left != nil {
// 			q.Push(value.left)
// 		}
// 		if value.right != nil {
// 			q.Push(value.right)
// 		}
// 	}
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

// func SingRotateLeft(a *AVLTreeNode) {
// 	a, b := a, a.left
// 	a.left = b.right
// 	b.right = a

// 	a.hight = max(a.left.hight, a.right.hight) + 1
// 	b.hight = max(b.left.hight, a.hight) + 1
// }

// func SingRotateRight(a *AVLTreeNode) {
// 	a, b := a, a.right
// 	a.right = b.left
// 	b.left = a

// 	a.hight = max(a.left.hight, a.right.hight) + 1
// 	b.hight = max(b.right.hight, a.hight) + 1
// }

// func DoubleRotateLR(a *AVLTreeNode) {
// 	SingRotateRight(a.left)
// 	SingRotateLeft(a)
// }

// func DoubleRotateRL(a *AVLTreeNode) {
// 	SingRotateLeft(a.right)
// 	SingRotateRight(a)
// }

// type Queue struct {
// 	front    int
// 	rare     int
// 	queue    []*AVLTreeNode
// 	size     int
// 	capacity int
// }

// func (q *Queue) Init(capacity int) {
// 	if capacity <= 0 {
// 		return
// 	}
// 	q.capacity = capacity
// 	q.size = 0
// 	q.queue = make([]*AVLTreeNode, q.capacity)
// 	q.front, q.rare = 0, 0
// }

// func (q *Queue) isEmpty() bool {
// 	return q.size == 0
// }

// func (q *Queue) Push(value *AVLTreeNode) error {
// 	if q.size == q.capacity {
// 		return errors.New("queue is full")
// 	}

// 	q.queue[q.rare] = value
// 	q.size++
// 	q.rare = (q.rare + 1) % q.capacity
// 	return nil
// }

// func (q *Queue) Pop() (value *AVLTreeNode, err error) {
// 	if q.size == 0 {
// 		return nil, errors.New("queue is empty")
// 	}
// 	value = q.queue[q.front]
// 	q.front = (q.front + 1) % q.capacity
// 	q.size--
// 	return value, nil
// }

// func main() {
// 	a := new(AVLTreeNode)
// 	a.Init()
// 	a.Insert(3)
// 	a.Insert(6)
// 	a.Insert(4)
// 	a.Insert(2)
// 	a.Insert(5)
// 	PriorPrint(a)
// 	fmt.Println()
// 	// MiddlePrint(a)
// 	// fmt.Println()
// 	// fmt.Println(a.Find(9))
// 	// fmt.Println(Find2(9, a))
// 	// PriorPrintHight(a)
// 	node := a.FindNode(3)
// 	a.Delete(node)
// 	PriorPrint(a)

// 	// a.LevelPrint()
// }
