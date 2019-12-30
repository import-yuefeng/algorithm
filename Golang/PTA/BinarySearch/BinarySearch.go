package main

import (
	"fmt"
)

func main() {
	// var array [100]int
	// for i := 0; i < 30; i++ {
	// 	array[i] = i
	// }
	// for i := 40; i < 100; i++ {
	// 	array[i] = i
	// }
	var array = [...]int{2, 4, 6, 8, 10, 12}

	fmt.Println(array)
	// fmt.Println(Search(len(array), array[:], 100))
	i, j := Search2(6, array[:], 5)

	fmt.Println(i, j)
	// fmt.Println(array[i], array[j])

}

func Search(n int, array []int, value int) (site, num int) {
	if n <= 0 || len(array) == 0 {
		return -1, 0
	}
	left, right := 0, n-1
	num = 0
	for left <= right {
		middle := (left + right) >> 1
		num++
		if array[middle] == value {
			return middle, num
		} else if array[middle] > value {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1, num
}

func Search2(n int, array []int, value int) (i, j int) {
	if n <= 0 || len(array) == 0 {
		return -1, -1
	}
	if array[n-1] < value {
		return n - 2, n - 1
	} else if array[0] > value {
		return -1, 0
	}
	i, j = 0, 0
	left, right := 0, len(array)-1
	for left <= right {
		middle := (left + right) / 2
		if array[middle] == value {
			return middle, middle
		} else if array[middle] > value {
			right = middle - 1
			j = middle
			if array[middle-1] > value {
				j = middle - 1
			}
		} else {
			left = middle + 1
			i = middle
			if array[middle+1] < value {
				i = middle + 1
			}
		}
	}
	return i, j
}
