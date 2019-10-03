package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var array [100]int
	for i := 0; i < 100; i++ {
		array[i] = rand.Intn(200)
	}
	fmt.Println(array)
	mergeSort(array[:], 0, len(array)-1)
	fmt.Println(array)
}

func merge(All []int, low int, middle int, high int) {
	leftLen := middle - low + 1
	rightLen := high - middle

	left := make([]int, leftLen+1)
	// 是否加一取决于你要将 middle 分于哪一个地方
	// 此处为分到了左边
	for i := 0; i < leftLen; i++ {
		left[i] = All[i+low]
		// 注意硬编码部分
	}
	left[leftLen] = 99999 // 哨兵
	right := make([]int, rightLen+1)

	for i := 0; i < rightLen; i++ {
		right[i] = All[middle+i+1]
		// 注意 这里的+1 是因为 middle 被分到了左边的数组中, 所以右边+1
	}
	right[rightLen] = 99999 // 哨兵
	// 哨兵的作用是 使得若有一边到达了尾部, 因为哨兵无限大, 所以使得只能走完另一边, 该请求要求哨兵足够大
	l, r := 0, 0
	for i := low; i <= high; i++ {
		if left[l] <= right[r] {
			All[i] = left[l]
			l++
		} else {
			All[i] = right[r]
			r++
		}
	}
}

func mergeSort(A []int, left int, right int) {
	if left >= right {
		return
	}
	middle := (left + right) / 2
	mergeSort(A, left, middle)
	mergeSort(A, middle+1, right)
	merge(A, left, middle, right)
}
