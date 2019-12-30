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
	Bubbling(a[:], 0, len(a)-1)
	fmt.Println(a)
}

func Bubbling(A []int, left int, right int) {
	for ; left <= right; left++ {
		for i := 0; i < left; i++ {
			if A[i] > A[left] {
				A[left], A[i] = A[i], A[left]
			}
		}
	}
}
