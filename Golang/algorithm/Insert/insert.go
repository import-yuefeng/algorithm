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
	Insert(a[:], 0, len(a)-1)
	fmt.Println(a)
}

func Insert(A []int, left int, right int) {
	left += 1
	for ; left < len(A); left++ {
		key := A[left]
		i := left - 1
		for ; i >= 0 && A[i] > key; i-- {
			A[i+1] = A[i]
		}
		A[i+1] = key
	}
}
