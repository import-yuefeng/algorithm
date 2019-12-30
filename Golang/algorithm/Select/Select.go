package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var a [100]int
	rand.Seed(time.Now().Unix())
	for i := 0; i < 100; i++ {
		a[i] = rand.Intn(100)
	}
	fmt.Println(a)
	SelectSort(a[:], 0, len(a)-1)
	fmt.Println(a)
}

func SelectSort(A []int, left int, right int) {
	for ; left <= right; left++ {
		tmp := left
		min := A[left]
		for i := left; i <= right; i++ {
			if min > A[i] {
				min = A[i]
				tmp = i
			}
		}
		A[left], A[tmp] = A[tmp], A[left]
	}
}
