// package binaryTree
package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	size       int
	capacity   int
	stackArray []*BinaryTree
}


func (stack *Stack) Init(capacity int) error {
	if capacity <= 0 {
		return errors.New("capacity invaild")
	}
	stack.size = 0
	stack.capacity = capacity
	stack.stackArray = make([]*BinaryTree, stack.capacity)
	return nil
}

func (stack *Stack) Push(value *BinaryTree) error {
	if stack.isFull() {
		return errors.New("stack overflow")
	}
	stack.stackArray[stack.size] = value
	stack.size++
	return nil
}

func (stack *Stack) Pop() (value *BinaryTree, err error) {
	if stack.isEmpty() {
		return nil, errors.New("stack empty")
	}
	stack.size--
	return stack.stackArray[stack.size], nil
}

func (stack *Stack) PrintStack() {
	tmpSize := stack.size
	for i := 0; i < tmpSize; i++ {
		if value, err := stack.Pop(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Printf("%v ", value)
		}
	}
	fmt.Println()
}

// Size func return stack size now
func (stack *Stack) Size() (size int, err error) {
	return stack.size, nil
}

func (stack *Stack) isFull() bool {
	return stack.size >= stack.capacity
}

func (stack *Stack) isEmpty() bool {
	return stack.size == 0
}

type Queue struct {
	size     int
	capacity int
	queue    []*BinaryTree
	front    int
	rare     int
}

func (queue *Queue) Init(capacity int) error {
	if capacity <= 0 {
		return errors.New("capacity must be greater than 0")
	}
	queue.capacity = capacity
	queue.size = 0
	queue.front, queue.rare = 0, 0
	queue.queue = make([]*BinaryTree, queue.capacity)
	return nil
}

func (queue *Queue) Size() int {
	return queue.size
}

func (queue *Queue) isFull() bool {
	return queue.size == queue.capacity
	// (queue.rare +1 ) % queue.capacity == queue.front
}

func (queue *Queue) isEmpty() bool {
	return queue.size == 0
	// queue.front == queue.rare
}

func (queue *Queue) Push(value *BinaryTree) error {
	if queue.isFull() || queue == nil {
		return errors.New("Queue is full")
	}
	queue.queue[queue.rare] = value
	queue.rare = (queue.rare + 1) % queue.capacity
	queue.size++
	return nil
}

func (queue *Queue) Pop() (value *BinaryTree, err error) {
	if queue == nil || queue.isEmpty() {
		return nil, errors.New("queue is empty")
	}
	value = queue.queue[queue.front]
	queue.front = (queue.front + 1) % queue.capacity
	queue.size--
	return value, nil

}

type BinaryTree struct {
	value   int
	size    int
	right   *BinaryTree
	left    *BinaryTree
	isVisit bool
}

func (bt *BinaryTree) Init(headValue int) {
	bt.value = headValue
	bt.size = 0
	bt.right, bt.left = nil, nil
	return
}

func (bt *BinaryTree) Exist(value int) bool {
	if bt == nil {
		return false
	}
	for bt != nil {
		if bt.value == value {
			return true
		} else if bt.value > value {
			bt = bt.left
		} else {
			bt = bt.right
		}
	}
	return false
}

func (bt *BinaryTree) Find(value int) *BinaryTree {
	if bt == nil {
		return nil
	}
	tmp := bt
	for tmp != nil {
		if tmp.value == value {
			return tmp
		} else if tmp.value > value {
			tmp = tmp.left
		} else {
			tmp = tmp.right
		}
	}
	return nil
}

func (bt *BinaryTree) Size() (size int) {
	return bt.size
}

func (bt *BinaryTree) Insert(value int) error {
	// 插入的时候不会改变树的上层结构
	newElement := new(BinaryTree)
	newElement.value = value
	newElement.right, newElement.left = nil, nil
	newElement.isVisit = false
	tmp := bt
	for tmp != nil {
		if newElement.value > tmp.value {
			if tmp.right == nil {
				tmp.right = newElement
				break
			}
			tmp = tmp.right
		} else if newElement.value <= tmp.value {
			if tmp.left == nil {
				tmp.left = newElement
				break
			}
			tmp = tmp.left
		}
	}
	bt.size++
	return nil
}

func (bt *BinaryTree) Delete(node *BinaryTree) error {
	if bt == nil {
		return errors.New("binaryTree is nil")
	}
	priorNode := bt.findPriorWithElement(node)
	newElement := new(BinaryTree)

	if node.right == nil && node.left != nil {
		newElement = node.left
	} else if node.left == nil && node.right != nil {
		newElement = node.right
	} else if node.left == nil && node.right == nil {
		newElement = nil
	} else {
		tempPoint := node
		tmp := findMaxElement(tempPoint.left)
		newElement.value = tmp.value
		bt.Delete(tmp)
		newElement.left = node.left
		newElement.right = node.right
	}
	if priorNode.right == node {
		priorNode.right = newElement
	} else if priorNode.left == node {
		priorNode.left = newElement
	}
	node = nil
	return nil
}

func findMaxElement(node *BinaryTree) (maxElement *BinaryTree) {
	if node == nil {
		return nil
	}
	if node.right != nil {
		node = findMaxElement(node.right)
	}
	maxElement = node
	return
}

func (bt *BinaryTree) findPriorWithElement(node *BinaryTree) *BinaryTree {
	if node == nil {
		return nil
	}
	point1, point2 := bt, bt
	for {
		if point1.value > node.value {
			point2 = point1
			point1 = point1.left
		} else if point1.value < node.value {
			point2 = point1
			point1 = point1.right
		} else {
			break
		}
	}
	return point2
}

func findMinElement(node *BinaryTree) (minElement *BinaryTree) {
	if node == nil {
		return nil
	}
	if node.left != nil {
		node = findMinElement(node.left)
	}
	minElement = node
	return

}

func FirstPrintRecursive(bt *BinaryTree) {
	if bt == nil {
		return
	}
	fmt.Printf("%d ", bt.value)
	FirstPrintRecursive(bt.left)
	FirstPrintRecursive(bt.right)
}

func MiddlePrintRecursive(bt *BinaryTree) {
	if bt == nil {
		return
	}
	MiddlePrintRecursive(bt.left)
	fmt.Printf("%d ", bt.value)
	MiddlePrintRecursive(bt.right)
}

func EndPrintRecursive(bt *BinaryTree) {
	if bt == nil {
		return
	}
	MiddlePrintRecursive(bt.left)
	MiddlePrintRecursive(bt.right)
	fmt.Printf("%d ", bt.value)
}

// func (bt *BinaryTree) FirstPrint() {
// 	if bt.size == 0 {
// 		return
// 	}
// 	stack := new(Stack)
// 	stack.Init(bt.size)
// 	tmpBinTree := bt
// 	for tmpBinTree != nil || !stack.isEmpty() {
// 		if tmpBinTree != nil {
// 			fmt.Printf("%d ", tmpBinTree.value)
// 			tmpBinTree = tmpBinTree.Left
// 		} else {
// 			t, err := stack.Pop()
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}
// 			tmpBinTree = t.Right
// 		}
// 	}
// 	fmt.Println()
// }

func (bt *BinaryTree) MiddlePrint() {
	if bt.size == 0 {
		return
	}
	tmp := bt
	tmpStack := new(Stack)
	tmpStack.Init(bt.size)

	for tmp != nil || !tmpStack.isEmpty() {
		if tmp != nil {
			if err := tmpStack.Push(tmp); err != nil {
				fmt.Println(err)
				break
			}
			tmp = tmp.left
		} else {
			if value, err := tmpStack.Pop(); err == nil {
				fmt.Printf("%d ", value.value)
				if value.right != nil {
					tmp = value.right
				}
			} else {
				fmt.Println(err)
				break
			}
		}
	}
	fmt.Println()
}
func (bt *BinaryTree) LevelPrint() {
	if bt == nil {
		return
	}
	tmpQueue := new(Queue)
	tmpQueue.Init(bt.size)
	tmpBT := bt
	tmpQueue.Push(tmpBT)
	for !tmpQueue.isEmpty() {
		value, err := tmpQueue.Pop()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%d ", value.value)
		if value.left != nil {
			tmpQueue.Push(value.left)
		}
		if value.right != nil {
			tmpQueue.Push(value.right)
		}
	}
	fmt.Println()
}
func (bt *BinaryTree) EndPrint() {
	if bt.size == 0 {
		return
	}
	tmp := bt
	tmpStack := new(Stack)
	tmpStack.Init(bt.size)

	for tmp != nil || !tmpStack.isEmpty() {
		if tmp != nil {
			tmpStack.Push(tmp)
			tmp = tmp.left

		} else {
			value, err := tmpStack.Pop()
			if err != nil {
				fmt.Println(err)
				break
			}
			if !value.isVisit {
				value.isVisit = true
				tmpStack.Push(value)
				tmp = value.right
			} else {
				fmt.Printf("%d ", value.value)
			}
		}
	}
	fmt.Println()
}

func main() {
	a := new(BinaryTree)
	a.Init(1)
	// for i := 1; i < 10; i++ {
	// 	a.Insert(i)
	// }

	a.Insert(3)
	a.Insert(6)
	a.Insert(6)
	a.Insert(6)
	a.Insert(4)
	a.Insert(2)
	a.Insert(5)
	// array := []int{1, 3, 6, 4, 2, 5}
	// a.FirstPrint()
	// for _, v := range array {
	// 	if node := a.Find(v); node == nil {
	// 		break
	// 	} else {
	// 		a.Delete(node)
	// 		a.FirstPrint()
	// 	}
	// }
	// a.FirstPrint()
	// a.LevelPrint()
	inOrderTraverse(a)
	// maxValue := findMaxElement(a)
	// fmt.Println(maxValue)
	// minValue := findMinElement(a)
	// fmt.Println(minValue)

	// a.FirstPrint()
	// a.MiddlePrint()
	// a.EndPrint()
	// FirstPrintRecursive(a)
	// fmt.Println()
	// MiddlePrintRecursive(a)
	// fmt.Println()
	// EndPrintRecursive(a)
	// fmt.Println()

}

func inOrderTraverse(cur *BinaryTree) {
    for cur != nil {
        if cur.left == nil {
			fmt.Println(cur.value) // 如果当前节点的左子树为空，那当然可以打印当前节点
			cur = cur.right
        } else {
            tmp := cur.left
            for tmp.right != nil && tmp.right != cur {
                // 遍历 tmp的左子树，找到左子树的最右边的叶子节点
                tmp = tmp.right
            }
            if tmp.right == nil { // 此时时第一次访问最右边的节点
                tmp.right = cur // 那么我们就记下往cur走的指针
                // 接着进一步进入 cur 的左子树中
                cur = cur.left
            } else {
				// 此时说明之前已经遍历到这，那么
				tmp.right = nil
				fmt.Printf("%d ", cur.value)
				cur = cur.right
            }
        }
    }
}
