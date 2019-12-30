package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(reorder(array[:]))
}

func reorder(array []int) []int {
	if len(array) == 0 {
		return nil
	} else if len(array) == 1 {
		return array
	}
	// 左边偶数, 右边奇数
	left, right := 0, len(array)-1
	for left < right {
		for array[left]&1 == 0 {
			left++
		}
		for array[right]&1 == 1 {
			right--
		}
		if left < right {
			array[left], array[right] = array[right], array[left]
		}
	}
	return array

}
