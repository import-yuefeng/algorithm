// package qSort

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	array := make([]int, 10)
	rand.Seed(time.Now().Unix())

	for i := 0; i < 10; i++ {
		array[i] = rand.Intn(100)
	}
	fmt.Println(array)
	qSort1(array, 0, len(array)-1)

}

func qSort1(array []int, left, right int) {
	if left < 0 || right < 0 || len(array) == 0 || left >= right {
		return
	}
	value := array[(left+right)>>1]
	fmt.Println(value)
	l, r := left, right
	for l < r {
		for array[l] < value {
			l++
		}
		for array[r] > value {
			r--
		}
		if l >= r {
			break
		}
		array[l], array[r] = array[r], array[l]
	}
	qSort1(array, left, l-1)
	qSort1(array, l+1, right)
	return
}

func qSort2(array []int, left, right int) {
	if len(array) == 0 || left < 0 || right < 0 || right < left {
		return
	}
	value = array[0]
	l, r := 0, right
	for i := 1; i < right; i++ {
		if array[i] > value {
			array[i], array[r] = array[r], array[i]
			i++
			r--
		} else if array[r] < value {

		}

	}
}
