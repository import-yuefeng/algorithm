package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	// hashWriting := 1
	// c := 0
	// c ^= hashWriting
	// var a int = 1
	// fmt.Println(a << 4)
	// fmt.Println(c)
	// c ^= hashWriting
	// fmt.Println(c)

	// rand.Seed(time.Now().Unix())
	// var array [100]int
	// for i := 0; i < 100; i++ {
	// 	array[i] = rand.Intn(200)
	// }
	// fmt.Println(array)
	// qSort(array[:], 0, len(array)-1)
	// fmt.Println(array)
}

type quickSortData struct {
	Left, Right int
}

func qSort(array []int, left, right int) {
	if left < right && left >= 0 {
		stack := NewStack(100)
		err := stack.Push(&quickSortData{Left: left, Right: right})
		if err != nil {
			fmt.Println(err)
			return
		}
		for !stack.IsEmpty() {
			value, err := stack.Pop()
			if err != nil {
				fmt.Println(err)
				return
			}
			left, right = value.(*quickSortData).Left, value.(*quickSortData).Right
			if left < right && left >= 0 {
				l, r := left, right
				middle := array[(l+r)>>1]
				for l < r {
					for l < r && array[l] < middle {
						l++
					}
					for l < r && array[r] > middle {
						r--
					}
					if l >= r {
						break
					}
					if array[r] == array[l] && array[r] == middle {
						r--
					} else {
						array[l], array[r] = array[r], array[l]
					}
				}
				err := stack.Push(&quickSortData{Left: left, Right: l - 1})
				if err != nil {
					fmt.Println(err)
					return
				}
				err = stack.Push(&quickSortData{Left: l + 1, Right: right})
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}

func mergeSort(array []int, left, right int) {
	if left < right {
		middle := (left + right) >> 1
		mergeSort(array, left, middle)
		mergeSort(array, middle+1, right)
		merge(array, left, middle, right)
	}
}

func merge(array []int, left, middle, right int) {

}

type Stack struct {
	top      int
	size     int
	capacity int
	data     []interface{}
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) IsFull() bool {
	return s.size == s.capacity
}

func NewStack(capacity int) *Stack {
	if capacity <= 0 {
		return nil
	}
	stack := new(Stack)
	stack.top = -1
	stack.size = 0
	stack.capacity = capacity
	stack.data = make([]interface{}, stack.capacity)
	return stack
}

func (s *Stack) Push(x interface{}) error {
	if s.IsFull() {
		return errors.New("stack is full")
	}
	s.top++
	s.data[s.top] = x
	s.size++
	return nil
}

func (s *Stack) Pop() (x interface{}, err error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	x = s.data[s.top]
	err = nil
	s.top--
	s.size--
	return
}


package main

import(
	"fmt"
	"math"
)

func main(){
	// res := A(4294967293)
	// res := hammingWeight(4294967293)
	var a uint32 = 513
	
	for i:=0; i<=31; i++ {
		if a<<uint32(i) & 0x80000000 == 1 {
			fmt.Println(32-i)
			break
		}
	}
	C([]int{1, 2, 3, 4})
	myB := []byte{'0', '0', '0', '0', '0'}
	Q(myB, len(myB), 0)
	// res := B(513)
	// fmt.Println(res)
}

func C(num []int) {
	array := make([][]int, 0)
	for start:=0; start<len(num); start++ {
		array = append(array, []int{})
		for i:=start; i<len(num); i++ {
			array[start] = append(array[start], num[i])
		}
		for end:=0; end<start; end++ {
			array[start] = append(array[start], num[end])
		}
	}
	fmt.Println(array)
}

func B(number uint32) int {
	count := 0
	var t uint32 = 1
	for i:=0; i<32; i++ {
		if (number & t) == 1 {
			count++
		}
		number >>= 1
	}
	return count
}

func hammingWeight(number uint32) int {
	if res, ok := numberPow(number); ok {
		return 1
	} else {
		tmp := math.Pow(2.0, float64(res))
		t := 1 + hammingWeight(number-uint32(tmp))
		return t
	}
}

func Q(num []byte, length int, index int) {
	if index == length {
		fmt.Println(string(num))
		return
	} else {
		for i:=0; i<10; i++ {
			num[index] = byte(i) + '0'
			Q(num, length, index+1)
		}
	}
}

func numberPow(number uint32) (count int, isExist bool) {
	var t uint32 = 1
	count = 0
	for number > t {
		t *= 2
		count++
	}
	if t > number {
		count -= 1
		t /= 2
		return count, false
	}
	return count, true
}