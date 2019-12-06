package main

import (
	"fmt"
	"math"
)

func main() {
	pageSum(11)
}

func pageSum(n int) {
	if n < 0 || n > math.MaxInt64 {
		return
	}
	pageMap := make(map[int]int, 10)
	for i := 0; i < 10; i++ {
		pageMap[i] = 0
	}
	for i := 0; i <= n; i++ {
		tmp := i
		for tmp != 0 {
			pageMap[tmp%10]++
			tmp /= 10
		}
	}
	for i := 0; i < 10; i++ {
		fmt.Println(pageMap[i])
	}
}
