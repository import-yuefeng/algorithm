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
	// Reverse(array[:], 0, len(array)-1)
	Convert(array[:], 0, 3, 6)
	// mergeSort2(array[:], 0, len(array)-1)
	fmt.Println(array)
}

func mergeSort2(array []int, left, right int) {
	if left >= right {
		return
	}
	middle := (left + right) >> 1
	mergeSort2(array, left, middle)
	mergeSort2(array, middle+1, right)
	merge2(array, left, middle, right)
	return
}

func merge4(array []int, left, middle, right int) {
	for m := middle + 1; m < right+1; m++ {
		i := m - 1
		if array[i] > array[m] {
			key := array[m]
			for ; i > 0 && array[i] > key; i-- {
				array[i] = array[i-1]
			}
			array[i] = key
		}
	}
}

func Reverse(array []int, left, right int) {
	i, j := left, right
	for i < j {
		array[i], array[j] = array[j], array[i]
		i++
		j--
	}
}

func Convert(array []int, left, middle, right int) {
	Reverse(array, left, middle)
	Reverse(array, middle+1, right)
	Reverse(array, left, right)
}

func merge3(array []int, left, middle, right int) {

}

func merge2(array []int, left, middle, right int) {
	tmp := make([]int, len(array))
	a, b, c := left, middle+1, left
	for ; a <= middle && b <= right; c++ {
		if array[a] >= array[b] {
			tmp[c] = array[b]
			b++
		} else {
			tmp[c] = array[a]
			a++
		}

	}
	for a <= middle {
		tmp[c] = array[a]
		a++
		c++
	}
	for b <= right {
		tmp[c] = array[b]
		b++
		c++
	}
	for left <= right {
		array[left] = tmp[left]
		left++
	}
}

func mergeSort(r []int) []int {
	length := len(r)
	if length <= 1 {
		return r
	}
	num := length / 2
	left := mergeSort(r[:num])
	right := mergeSort(r[num:])

	return merge(left, right)
}

func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}
