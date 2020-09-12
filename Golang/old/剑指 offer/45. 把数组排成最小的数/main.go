package main

import (
	"fmt"
)

func main() {
	num := []string{"3", "321", "32"}
	qSort(num, 0, len(num)-1, less)
	fmt.Println(num)
}

func qSort(data []string, left, right int, less func(a, b string) bool) {
	if len(data) <= 1 || left >= right {
		return
	}
	l, r := left, right
	mid := data[(l + ((r - l) >> 1))]
	for l < r {
		for l < r && less(data[l], mid) {
			l++
		}
		for l < r && !less(data[r], mid) {
			r--
		}
		if l >= r {
			break
		}
		if data[l] == data[r] && data[r] == mid {
			r--
		} else {
			data[l], data[r] = data[r], data[l]
		}
	}
	qSort(data, left, l-1, less)
	qSort(data, r+1, right, less)
}

// 核心就是 ab < ba 则最后结果就小
func less(a, b string) bool {
	if a+b < b+a {
		return true
	}
	return false
}
