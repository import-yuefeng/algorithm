// package perm
package main

import (
	"fmt"
	"runtime"
)

var array []int = []int{1, 2, 3, 3}

func main() {
	runtime.GC()
	check := make([]int, len(array))
	perm([]int{}, len(array), check)
}

func perm(now []int, end int, check []int) {
	if len(now) == end {
		fmt.Println(now)
	} else {
		for i := 0; i < end; i++ {
			if check[i] == 1 {
				continue
			}
			if i > 0 && array[i] == array[i-1] && check[i-1] == 1 {
				continue
			}
			check[i] = 1
			now = append(now, array[i])
			perm(now, end, check)
			now = now[:len(now)-1]
			check[i] = 0
		}
	}
}

func hasVal(arr []int, x int) bool {
	if len(arr) == 0 {
		return false
	}
	for _, v := range arr {
		if x == v {
			return true
		}
	}
	return false
}
