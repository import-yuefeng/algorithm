package main

import (
	"fmt"
)

func main() {
	num := []int{1, 2, 3, 4, 5, 6, 7, 7, 6}
	res := findRepeat(num)
	fmt.Println(res)
}

func findRepeat(num []int) int {
	// 虽然是有两个 for 循环,但是每个数字只需要最多交换两次就有序,所以其时间复杂度是 O(n)
	if len(num) == 0 {
		return -1 << 31
	}
	for i := 0; i < len(num); i++ {
		for i != num[i] {
			if num[i] == num[num[i]] {
				return num[i]
			}
			num[i], num[num[i]] = num[num[i]], num[i]
		}
	}
	return -1 << 31
}
