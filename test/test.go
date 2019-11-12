package main

import (
	"errors"
	"fmt"
	"math"
)

var coins = [...]int{1, 2, 5}

func main() {
	// res := coinChangeOne(11)
	// back := make([]int, 12)
	// for i, _ := range back {
	// 	back[i] = math.MaxInt64
	// }
	// fmt.Println(res)
	// back[0], back[1] = 0, 1
	// res = coinChangeTwo(back, 11)
	// fmt.Println(res)
	// res := longestCommonSubsequence("abcde", "ace")
	// res := maxSubArray([]int{-2, 11, -4, 13, -5, -2})
	// fmt.Println(res)
	// res = maxSubArray([]int{-2, 1, -3, 4, -1, 2, -1, -5, 4})
	// fmt.Println(res)
	// res = maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	// fmt.Println(res)
	// res = maxSubArray([]int{-2, -1, -3, -4, -1, -2, -1, -5, -4})
	// array := make([]int, 1000000)
	// for i := 0; i < 1000000; i++ {
	// 	array[i] = rand.Intn(10000000)
	// }
	// newArray := make([]int, 10000000)
	// go countSort(newArray, array)
	// qSort(array, 0, len(array)-1)
	// fmt.Println("quickSort")
	// qSort2(array)
	// fmt.Println(array)
	// res := bit(8, 4)
	// fmt.Println(res)
	// array := []int{1, 2, 3, 4, 5, 6, 6}
	// for i := 0; i < 32; i++ {
	// 	array1 := make([]int, 0)
	// 	array2 := make([]int, 0)
	// 	left, right := 0, len(array)-1
	// 	for _, v := range array {
	// 		if bit(v, i) { // equal 1
	// 			array1 = append(array1, v)
	// 		} else {
	// 			array2 = append(array2, v)
	// 		}
	// 		mid := len(array) / 2
	// 		if len(array1) > len(array2) {
	// 			left = mid + 1
	// 			if left < len(array) {
	// 				array = array[left:]
	// 			}
	// 		} else {
	// 			right = mid - 1
	// 			if right >= 0 {
	// 				array = array[:right]
	// 			}
	// 		}
	// 	}
	// }
	// array := []int{1, 2, 3}
	// array_1 := []int{1, 3, 2}
	// array_2 := []int{2, 1, 3}

	// stack := new(MyStack)
	// stack.Init(100)
	// stack.Push(array)
	// stack.Push(array_1)
	// stack.Push(array_2)
	// for {
	// 	if val, err := stack.Pop(); err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	} else {
	// 		fmt.Println(val)
	// 	}
	// }
	// return
	a := (-10 - 128) >> 31
	b := (2000 - 128) >> 31
	PrintBit(a)
	fmt.Println()
	PrintBit(b)
	fmt.Println()
	PrintBit(^a & 1024)
	fmt.Println()
	PrintBit(^b & 1024)

	// PrintBit(8)

}

type MyStack struct {
	top      int
	size     int
	capacity int
	data     []interface{}
}

func (s *MyStack) Init(capacity int) error {
	if capacity <= 0 {
		return errors.New("capacity is invaild!")
	}
	s.capacity = capacity
	s.top, s.size = -1, 0
	s.data = make([]interface{}, s.capacity)
	return nil
}

func (s *MyStack) Push(value interface{}) error {
	if s.size == s.capacity {
		return errors.New("stack is full!")
	}
	s.top++
	s.data[s.top] = value
	s.size++
	return nil
}

func (s *MyStack) Pop() (value interface{}, err error) {
	if s.size == 0 {
		return nil, errors.New("stack is empty!")
	}
	value = s.data[s.top]
	err = nil
	s.top--
	s.size--
	return
}

func PrintBit(num int) {
	tmp := num
	nums := 1
	for tmp > 1 {
		tmp /= 2
		nums++
	}
	for i := uint(nums) - 1; i >= 0; i-- {
		if i == 0 {
			fmt.Printf(" %d", (num&1)&1)
			break
		}
		if res := (num & (1 << i) >> i) & 1; res == 1 {
			fmt.Printf(" 1")
		} else {
			fmt.Printf(" 0")
		}
	}
}

func bit(num, site int) bool {
	// tmpNum := num
	// places := 1
	// for tmpNum > 1 {
	// 	tmpNum /= 2
	// 	places++
	// }
	site--
	a := uint(site)
	if ((num & (1 << a)) >> a) == 1 {
		return true
	}

	// for i := uint(0); i < uint(places); i++ {
	// 	// fmt.Printf("%d bit: %d\n", i+1, ((num&(1<<i))>>i)&1)
	// 	if uint(site) == i+1 && ((num&(1<<i))>>i) == 1 {
	// 		return true
	// 	}
	// }
	return false
}

func countSort(array []int, nums []int) {
	for _, v := range nums {
		array[v] = 1
	}
	fmt.Println("countSort")
	return
}

type Stack struct {
	top      int
	size     int
	capacity int
	data     []int
}

func (s *Stack) Init(capacity int) error {
	if capacity <= 0 {
		return errors.New("capacity is invaild!")
	}
	s.size, s.top = 0, -1
	s.capacity = capacity
	s.data = make([]int, s.capacity)
	return nil
}

func (s *Stack) Push(value int) error {
	if s.size == s.capacity {
		return errors.New("stack is full!")
	}
	s.top++
	s.data[s.top] = value
	s.size++
	return nil
}

func (s *Stack) Pop() (value int, err error) {
	if s.size == 0 {
		return -1, errors.New("stack is empty!")
	}
	value = s.data[s.top]
	err = nil
	s.top--
	s.size--
	return
}

func qSort2(nums []int) {
	if len(nums) <= 1 {
		return
	}
	left, right := 0, len(nums)-1
	stack := new(Stack)
	stack.Init(1000)
	stack.Push(left)
	stack.Push(right)
	for stack.size != 0 {
		r, err := stack.Pop()
		if err != nil {
			fmt.Println(err)
			return
		}
		l, err := stack.Pop()
		if err != nil {
			fmt.Println(err)
			return
		}
		if l >= r {
			break
		}
		mid := partation(nums, l, r)
		if l < mid-1 {
			stack.Push(l)
			stack.Push(mid - 1)
		}
		if r > mid+1 {
			stack.Push(mid + 1)
			stack.Push(r)
		}
	}
}

func partation(nums []int, l, r int) int {

	if l < r {
		point := nums[(l+r)>>1]
		for l < r {
			for l < r && nums[l] < point {
				l++
			}
			for l < r && nums[r] > point {
				r--
			}
			if l >= r {
				break
			}
			if nums[r] == nums[l] && nums[r] == point {
				r--
			} else {
				nums[r], nums[l] = nums[l], nums[r]
			}
		}
		return l
	}
	return -1
}

func qSort(nums []int, left, right int) {
	if left < right && left >= 0 {
		l, r := left, right
		point := nums[(l+r)>>1]
		for l < r {
			for l < r && nums[l] < point {
				l++
			}
			for l < r && nums[r] > point {
				r--
			}
			if l >= r {
				break
			}
			if nums[r] == nums[l] && nums[r] == point {
				r--
			} else {
				nums[l], nums[r] = nums[r], nums[l]
			}
		}
		qSort(nums, left, l-1)
		qSort(nums, l+1, right)
	}
	return
}

func maxSubArray(nums []int) int {
	max := nums[0]
	// left, right := 0, 0
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else {
		dp := make([]int, len(nums))
		dp[0] = nums[0]
		for i := 1; i < len(nums); i++ {
			dp[i] = Max(dp[i-1]+nums[i], nums[i])
			if dp[i] > max {
				max = dp[i]
			}
			if dp[i-1]+nums[i] < nums[i] {
				// left = i
			}

		}
	}
	// fmt.Println(left, right)
	// fmt.Println(nums[left:right])
	return max
}

func longestCommonSubsequence(text1, text2 string) int {
	a, b := len(text1), len(text2)
	dp := make([][]int, 0)
	for i := 0; i <= a; i++ {
		dp = append(dp, make([]int, b+1))
	}
	dp[0][0] = 0
	// init state
	for i := 1; i <= a; i++ {
		for j := 1; j <= b; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[a][b]
}

func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func coinChangeTwo(back []int, n int) int {
	if n == 0 {
		return 0
	}
	res := 1
	if back[n] != math.MaxInt64 {
		return back[n]
	}
	min := math.MaxInt64
	for _, v := range coins {
		if v > n {
			continue
		}
		tmp := coinChangeTwo(back, n-v)
		if tmp < min {
			min = tmp
		}
	}
	res += min
	back[n] = Min(back[n], res)
	return back[n]
}

func Min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func coinChangeOne(n int) int {
	res := 1
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		min := math.MaxInt64
		for _, v := range coins {
			if v > n {
				continue
			}
			tmp := coinChangeOne(n - v)
			if tmp < min {
				min = tmp
			}
		}
		res += min
	}
	return res
}
