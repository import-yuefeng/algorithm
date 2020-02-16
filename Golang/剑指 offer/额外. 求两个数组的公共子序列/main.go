package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	// a := []int{}
	b := []int{3, 5, 6, 6, 7, 9}
	// b := []int{}
	res := make([]int, 0)
	// res := LCS_3(a, b)
	LCS_4(a, b, &res)
	fmt.Println(res)
}

func LCS_1(a, b []int) (res []int) {
	p1, p2 := 0, 0
	for p1 < len(a) && p2 < len(b) {
		if a[p1] == b[p2] {
			res = append(res, a[p1])
			p1++
			p2++
		} else if a[p1] < b[p2] {
			p1++
		} else {
			p2++
		}
	}
	return
}


func LCS_2(a, b []int) (res []int) {
	for _, v := range a {
		if flag := BinarySearchByIterate(b, 0, len(b)-1, v); flag != -1 {
			res = append(res, v)
		}
	}
	return
}

func LCS_3(a, b []int) (res []int) {
	flag := 0
	for _, v := range a {
		flag = BinarySearchByRecursive(b, flag, len(b)-1, v)
		if flag != -1 {
			res = append(res, v)
		} else {
			flag = 0
		}
	}
	return
}

func LCS_4(a, b []int, res *[]int) {
	if len(a) >= 1 {
		target := a[0]
		flag := BinarySearchByRecursive(b, 0, len(b)-1, target)
		if flag != -1 {
			*res = append(*res, target)
		}
		LCS_4(a[1:], b, res)
	}
	return 
}

func BinarySearchByIterate(num []int, left, right int, target int) (res int) {
	// not found return -1
	l, r := left, right
	for l<=r {
		mid := l + ((r-l)>>1)
		if num[mid] == target {
			return mid
		} else if num[mid] > target {
			r = mid-1
		} else {
			l = mid+1
		}
	}
	return -1
}

func BinarySearchByRecursive(num []int, left, right, target int) (res int) {
	// not found return -1
	if left <= right {
		l, r := left, right
		mid := l + ((r-l)>>1)
		if num[mid] == target {
			return mid
		} else if num[mid] > target {
			return BinarySearchByRecursive(num, left, mid-1, target)
		} else {
			return BinarySearchByRecursive(num, mid+1, right, target)
		}
	} else {
		return -1
	}
}