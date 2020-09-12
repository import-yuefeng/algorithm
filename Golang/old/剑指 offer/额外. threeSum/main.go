package main

import (
	"fmt"
)

func main() {

	fmt.Println("hello world")
	// res := threeSum([]int{1, 2, 3, 4, 5, 6})
	// fmt.Println(res)
	val := []int{1, 2, 5, 6, 3, 4}
	qSort2(val, 0, len(val))
	fmt.Println(val)
}

func threeSum(num []int) [][]int {
	qSort(num, 0, len(num)-1)
	res := make([][]int, 0)
	for j := len(num) - 1; j >= 2; j-- {
		l, r := 0, j-1
		for l < r {
			if num[l]+num[r] == num[j] {
				res = append(res, []int{l, r, j})
				l++
			} else if num[r]+num[l] > num[j] {
				r--
			} else {
				l++
			}
		}
	}
	return res
}

func qSort(num []int, left, right int) {
	if left < right && left >= 0 {
		l, r := left, right
		mid := num[(left+right)>>1]
		for l < r {
			for l < r && num[l] < mid {
				l++
			}
			for r <= right && num[r] > mid {
				r--
			}
			if l >= r {
				break
			}
			if num[r] == num[l] && num[r] == mid {
				r--
			} else {
				num[l], num[r] = num[r], num[l]
			}
		}
		qSort(num, left, l-1)
		qSort(num, l+1, right)
	}
}

func qSort2(num []int, left, right int) {
	if left < right && left >= 0 {
		l, r := left, left+1
		mid := num[l+((r-l)>>1)]
		for r < right {
			for l < r && num[l] < mid {
				l++
			}
			for r < right && num[r] > mid {
				r++
			}
			if r >= right {
				break
			}
			if num[r] == num[l] && num[r] == mid {
				r++
			} else {
				num[l], num[r] = num[r], num[l]
			}
		}
		qSort2(num, left, l)
		qSort2(num, l+1, right)
	}
}
