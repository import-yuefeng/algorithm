// package qSort

package main

import (
	"errors"
	"fmt"
)

func main() {
	array := []int{53, 79, 63, 30, 62, 4, 33, 53, 17, 43}
	// array := []int{53, 53}

	// array := []int{54, 79, 63, 30, 62, 4, 33, 53, 17, 43}

	fmt.Println(array)
	qSort1(array, 0, len(array)-1)
	fmt.Println(array)

}

type qSortData struct {
	left  int
	right int
}

type myQueue struct {
	front    int
	rare     int
	size     int
	capacity int
	data     []*qSortData
}

func (queue *myQueue) Init(capacity int) error {
	if capacity < 0 {
		return errors.New("capacity must be greater than 0")
	}
	queue.size, queue.front, queue.rare = 0, 0, 0
	queue.capacity = capacity
	queue.data = make([]*qSortData, queue.capacity)
	return nil
}

func (queue *myQueue) Push(value *qSortData) error {
	if queue.size == queue.capacity {
		return errors.New("queue is full")
	}
	queue.data[queue.rare] = value
	queue.rare = (queue.rare + 1) % queue.capacity
	queue.size++
	return nil
}

func (queue *myQueue) Pop() (value *qSortData, err error) {
	if queue.isEmpty() {
		return nil, errors.New("queue is empty")
	}
	value = queue.data[queue.front]
	queue.front = (queue.front + 1) % queue.capacity
	queue.size--
	return value, nil
}

func (queue *myQueue) isEmpty() bool {
	return queue.size == 0
}

type myStack struct {
	top      int
	size     int
	capacity int
	data     []*qSortData
}

func (stack *myStack) Init(capacity int) error {
	if capacity < 0 {
		return errors.New("capacity must be greater than 0")
	}
	stack.top, stack.size = -1, 0
	stack.capacity = capacity
	stack.data = make([]*qSortData, stack.capacity)
	return nil
}

func (stack *myStack) Push(value *qSortData) error {
	if stack.size == stack.capacity {
		return errors.New("stack is full")
	}
	stack.top++
	stack.data[stack.top] = value
	stack.size++
	return nil
}

func (stack *myStack) Pop() (value *qSortData, err error) {
	if stack.size == 0 {
		return nil, errors.New("stack is empty")
	}
	value = stack.data[stack.top]
	stack.top--
	stack.size--
	return value, nil
}

func (stack *myStack) isEmpty() bool {
	return stack.size == 0
}

func qSort1(array []int, left, right int) {
	if left < right {
		l, r := left, right
		key := array[(left+right)/2]

		for l < r {
			for l < right && array[l] < key {
				l++
			}
			for r > left && array[r] > key {
				r--
			}
			if l >= r {
				break
			}
			if key == array[r] && array[r] == array[l] {
				r--
			} else {
				array[l], array[r] = array[r], array[l]
			}
		}
		qSort1(array, left, l-1)
		qSort1(array, r+1, right)
	}
}

func qSort5(array []int, left, right int) {
	if left < right {
		l, r := left, right
		key := array[(left+right)/2]
		for l < r {
			for l < right && array[l] < key || (l < right && array[l] == key && l != (left+right)/2) {
				l++
			}
			for r > left && array[r] > key {
				r--
			}
			if l >= r {
				break
			}
			array[l], array[r] = array[r], array[l]
		}
		qSort5(array, left, l-1)
		qSort5(array, l+1, right)
	}
}

func qSort2(array []int, left, right int) {
	if left > right || left < 0 || right < 0 || len(array) == 0 {
		return
	}
	point := partation(array, left, right)
	qSort2(array, left, point-1)
	qSort2(array, point+1, right)
}

func qSort3(array []int, left, right int) {
	if left > right || left < 0 || right < 0 || len(array) == 0 {
		return
	}
	stack := new(myStack)
	stack.Init(10000)
	var tmp = new(qSortData)
	tmp.left, tmp.right = left, right
	stack.Push(tmp)
	for !stack.isEmpty() {
		value, err := stack.Pop()
		if err != nil {
			return
		}
		left, right := value.left, value.right
		if left > right {
			break
		}
		fmt.Println(left, right)
		point := partation(array, left, right)
		tmpLeft, tmpRight := new(qSortData), new(qSortData)
		tmpLeft.left, tmpLeft.right = left, point-1
		tmpRight.left, tmpRight.right = point+1, right
		if left < point-1 {
			stack.Push(tmpLeft)
		}
		if right > point+1 {
			stack.Push(tmpRight)
		}
	}
	return
}

func qSort6(array []int, left, right int) {
	if left > right || left < 0 || right < 0 || len(array) == 0 {
		return
	}
	queue := new(myQueue)
	queue.Init(10000)
	var tmp = new(qSortData)
	tmp.left, tmp.right = left, right
	queue.Push(tmp)
	for !queue.isEmpty() {
		value, err := queue.Pop()
		if err != nil {
			return
		}
		left, right := value.left, value.right
		if left > right {
			break
		}
		fmt.Println(left, right)
		point := partation(array, left, right)
		tmpLeft, tmpRight := new(qSortData), new(qSortData)
		tmpLeft.left, tmpLeft.right = left, point-1
		tmpRight.left, tmpRight.right = point+1, right
		if left < point-1 {
			queue.Push(tmpLeft)
		}
		if right > point+1 {
			queue.Push(tmpRight)
		}
	}
	return
}

func partation(array []int, left, right int) int {
	if left > right || left < 0 || right < 0 || len(array) == 0 {
		return -1
	}
	l, r := left, right
	point := array[(left+right)>>1]
	for l < r {
		for l < r && array[l] < point {
			l++
		}
		for l < r && array[r] > point {
			r--
		}
		if l >= r {
			break
		}
		array[l], array[r] = array[r], array[l]
	}
	return l
}

func qSort4(array []int, left, right int) {
	if left > right || left < 0 || right < 0 || len(array) == 0 {
		return
	}
	l, r := left, right
	value := array[l]
	for l < r {
		for l < r && array[r] > value {
			r--
		}
		if l < r {
			array[l] = array[r]
		}
		for l < r && array[l] <= value {
			l++
		}
		if l < r {
			array[r] = array[l]
		}
	}
	array[l] = value
	qSort4(array, left, l-1)
	qSort4(array, l+1, right)
	return
}

// int key =a[nlow];
// while(nlow < nhigh)
// {
// 	//从后向前，找第一个比key 小的
// 	while(nlow < nhigh && a[nhigh] >= key)
// 	{
// 		nhigh--;
// 	}
// 	//找到比key 小的，填到坑nlow 位置， nhigh 变成坑
// 	if(nlow < nhigh)
// 	a[nlow] = a[nhigh];
// 	//从前向后找比key 大的，放到坑nhigh中
// 	while (nlow < nhigh && a[nlow] <= key)
// 	{
// 		nlow++;
// 	}
// 		if(nlow < nhigh)
// 	a[nhigh] = a[nlow];
// }
// a[nlow] = key;
