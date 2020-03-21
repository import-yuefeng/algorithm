package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := []string{"3", "321", "32"}
	qSort(num, 0, len(num)-1, Less2)
	fmt.Println(num)
}

func qSort(num []string, left, right int, Less func(a, b string) bool) {
	if len(num) <= 1 {
		return
	}
	if left < right && left >= 0 {
		l, r := left, right
		mid := num[left+((right-left)>>1)]
		for l < r {
			for l < r && Less(num[l], mid) {
				l++
			}
			for l < r && Less(mid, num[r]) {
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
		qSort(num, left, l-1, Less)
		qSort(num, r+1, right, Less)
	}
}

func Less(a, b string) bool {
	ab := a + b
	ba := b + a
	if ab < ba {
		return true
	}
	return false
}

func Less2(a, b string) bool {
	divisorA, divisorB := getDivisor(a), getDivisor(b)
	valA, _ := strconv.Atoi(a)
	valB, _ := strconv.Atoi(b)
	if valA%divisorA < valB%divisorB {
		return true
	}
	return false
}

func getDivisor(a string) int {
	count := len(a)
	base := 9
	for count > 1 {
		base = base*10 + 9
		count--
	}
	return base
}
