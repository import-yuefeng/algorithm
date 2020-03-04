package main

import (
	"errors"
	"fmt"
	"math/rand"
	"unsafe"
)

type tp1 struct {
	a int8
	b int32
	c int64
}

type tp2 struct {
	a int8
	c int64
	b int32
}

func main() {
	// fmt.Println("hello world")
	// skipList := NewSkipList(16)
	// skipList.Insert(10, 10)
	// skipList.Insert(4, 4)
	// skipList.Insert(2, 2)
	// if res, err := skipList.Search(2); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	if t, ok := res.Val.(int); ok {
	// 		fmt.Println(t)
	// 	}
	// }
	// A()
	fmt.Println(1 ^ 2 ^ 2 ^ 3 ^ 4 ^ 2 ^ 3 ^ 4 ^ 1)
	for i := 1; i <= 7; i++ {
		num := []int{2, 4, 78, 3, 1, 25, 345}
		res := quickSelect(num, 0, len(num)-1, i)
		fmt.Printf("%d ", res)
	}
	fmt.Println(unsafe.Sizeof(tp1{}))
	fmt.Println(unsafe.Sizeof(tp2{}))
}

const (
	p float32 = 0.25
)

type Node struct {
	Score   float64
	Val     interface{}
	forward []*Node
}

type SkipList struct {
	maxLength int
	header    *Node
	size      int
	level     int
}

func NewNode(score float64, val interface{}, level int) *Node {
	return &Node{
		Score:   score,
		Val:     val,
		forward: make([]*Node, level),
	}
}

func NewSkipList(maxLength int) *SkipList {
	return &SkipList{
		maxLength: maxLength,
		header:    NewNode(0, 0, maxLength),
	}
}

func (sk *SkipList) randomLevel() int {
	level := 1
	if rand.Float32() < p && level < sk.maxLength {
		level++
	}
	return level
}

func (sk *SkipList) Search(score float64) (res *Node, err error) {

	x := sk.header
	for i := sk.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].Score < score {
			x = x.forward[i]
		}
	}
	x = x.forward[0]
	if x != nil && x.Score == score {
		return x, nil
	}
	return nil, errors.New("Not find")
}

func (sk *SkipList) Insert(score float64, val interface{}) {
	x := sk.header
	update := make([]*Node, sk.maxLength)
	for i := sk.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].Score > score {
			x = x.forward[i]
		}
		update[i] = x
	}
	x = x.forward[0]
	if x != nil && x.Score == score {
		x.Val = val
		return
	}
	level := sk.randomLevel()
	if level > sk.level {
		level = sk.level + 1
		update[sk.level] = sk.header
		sk.level = level
	}
	newNode := NewNode(score, val, level)
	for i := 0; i < level; i++ {
		newNode.forward[i] = update[i].forward[i]
		update[i].forward[i] = newNode
	}
	sk.size++
	return
}

func (sk *SkipList) Delete(score float64) {
	if sk.size == 0 {
		return
	}
	update := make([]*Node, sk.maxLength)
	x := sk.header
	for i := sk.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].Score < score {
			x = x.forward[i]
		}
		update[i] = x
	}
	x = x.forward[0]
	if x != nil && x.Score == score {
		for i := 0; i < sk.level; i++ {
			if update[i].forward[i] != x {
				return
			}
			update[i].forward[i] = x.forward[i]
		}
	}
	return
}

func A() {
	test := make([]int, 10)
	fmt.Println(len(test))
	B(test)
	fmt.Println(len(test))
	fmt.Println(test)
}

type C struct {
	a int8
	c int64
	b int16
}

func B(num []int) {
	// num = append(num, 5)
	test := C{}
	// a := make([]int, 3)
	// b := a
	// b[0] = 2
	// res := reflect.DeepEqual(a, b)
	// fmt.Println(res)
	// num[0] = 100
	// fmt.Println(num[10])
	fmt.Println(unsafe.Sizeof(test))
}

func findMin(num []int) int {
	minVal := 1 << 31
	for _, v := range num {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}

func quickSelect(num []int, left, right, k int) int {
	if left < right {
		l, r := left, right
		mid := num[(l + ((r - l) >> 1))]
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
				num[l], num[r] = num[r], num[l]
			}
		}
		m := l - left + 1
		if m == k {
			return num[l]
		} else if k < m {
			quickSelect(num, left, l-1, k)
		} else {
			quickSelect(num, r+1, right, k-m+1)
		}
	}
	return num[k-1]
}
