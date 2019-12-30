package main

import "fmt"

func main() {
	task := 10
	taskList := []int{56, 12, 1, 99, 1000, 234, 33, 55, 99, 812}
	mergeSort(taskList, 0, task-1)
	fmt.Println(taskList)
}

func schedule(taskList []int, task int) float64 {
	nowTime := 0
	waitTime := 0.0
	for i, v := range taskList {
		nowTime += v
		for _, wait := range taskList[:i] {
			waitTime += float64(wait)
		}
	}
	return waitTime / float64(task)
}

func mergeSort(array []int, left, right int) {
	if len(array) <= 1 || left >= right {
		return
	}
	middle := (left + right) >> 1
	mergeSort(array, left, middle)
	mergeSort(array, middle+1, right)
	merge(array, left, middle, right)
	return
}

func merge(array []int, left, middle, right int) {
	if left >= right {
		return
	}
	i := left
	j := middle
	for i < j && j <= right {
		for i < middle && array[i] < array[j] {
			i++
		}
		tmp := j
		for j <= right && array[j] <= array[i] {
			j++
		}
		Convert(array, i, tmp, j-1)
		i = (j - tmp) + i - 1
	}
	return
}

func Convert(array []int, left, middle, right int) {
	convert(array, left, middle)
	convert(array, middle+1, right)
	convert(array, left, right)
}

func convert(array []int, left, right int) {
	for left < right {
		array[left], array[right] = array[right], array[left]
		left++
		right--
	}
	return
}
