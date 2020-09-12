package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	test := []int{1, 2, 3, 7, 4, 5, 6, 7}
	// fmt.Println(FindRepeat2(test, 8))
	fmt.Println(FindRepeat4(test, 8))
	// fmt.Println(FindRepeat3(test))
	// if res, err := FindRepeat(test); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(res)
	// }
}

func FindRepeat(num []int) (int, error) {
	if len(num) == 0 {
		return -1, errors.New("Not element")
	}
	count := 1
	curVal := num[0]
	for _, v := range num[1:] {
		if curVal == v {
			count++
		} else {
			count--
		}
		if count == 0 {
			curVal = v
			count = 1
		}
	}
	return curVal, nil
}

func FindRepeat2(num []int, n int) int {
	// 有 n 个数, 其范围是 1~n-1 可以用该方法
	t := 1
	for i := 1; i < n; i++ {
		t ^= i
		t ^= num[i]
	}
	return t
}

func FindRepeat3(num []int) int {
	isRepeat := make(map[int]int)
	for _, v := range num {
		if _, ok := isRepeat[v]; ok {
			isRepeat[v] += 1
			if isRepeat[v] > 1 {
				return v
			}
		} else {
			isRepeat[v] = 1
		}
	}
	return -1
}

func FindRepeat4(num []int, n int) int {
	curSum := 0
	allSum := 0
	for i, v := range num {
		allSum += i
		curSum += v
	}
	return abs(allSum - curSum)
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
