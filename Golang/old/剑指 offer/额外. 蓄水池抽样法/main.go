package main

import (
	"fmt"
	"math/rand"
	"time"
)

type S struct {
	r []int
}

func main() {
	for i := 0; i < 1000; i++ {
		res := RandomSelectNumber(1)
		fmt.Println(res)
	}
}

func EqualSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	point := 0
	for ; point < len(a); point++ {
		if a[point] != b[point] {
			return false
		}
	}
	return true
}

func RandomSelectNumber(k int) []int {
	rawList := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = rawList[i]
	}
	for i := k; i < len(rawList); i++ {
		rand.Seed(time.Now().UnixNano())
		t := rand.Intn(i)
		if t < k {
			res[t] = rawList[i]
		}
	}
	return res
}
