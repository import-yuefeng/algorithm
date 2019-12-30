package main

import (
	"errors"
	"fmt"
)

const (
	FIBNACCI_MAXSIZE = 30
)

func main() {
	array := []int{1, 3, 5, 8, 9, 10, 10, 12, 29, 36}

	// array := []int{11, 11, 11, 11, 11, 15, 16}
	// findFirstEqual(array[:], 7, 0, len(array)-1, 0)
	// findFirstEqual(array[:], 29, 0, len(array)-1, 0)
	// if res, err := find(array[:], 5, 0, len(array)-1, 0); err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(res)
	// }
	if res, err := produceFib(100); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func produceFib(size int) (fibArray []int, err error) {
	if size > 0 {
		fibArray = make([]int, size)
		fibArray[0] = 0
		if size > 1 {
			fibArray[1] = 1
		}
		for i := 2; i < size; i++ {
			fibArray[i] = fibArray[i-1] + fibArray[i-2]
			if fibArray[i] < fibArray[i-1] {
				// reserve overflow
				fibArray[i] = 0
				return fibArray, errors.New("fibArray overflow")
			}
		}
		return fibArray, nil
	}
	return nil, errors.New("size is invaild")
}

func find(array []int, value, left, right int) (site int, err error) {
	fibArray, err := produceFib(FIBNACCI_MAXSIZE)
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("apply for fibArray faild")
	}
	k := 0
	for len(array) > fibArray[k] {
		k++
	}
	tmpArray := make([]int, fibArray[k])
	for i := 0; i < len(tmpArray); i++ {
		if len(array)-1 < i {
			tmpArray[i] = array[len(array)-1]
		} else {
			tmpArray[i] = array[i]
		}
	}

	for left <= right {
		middle = left + fibArray[k]
		if array[middle] == value {
			if middle > right {
				return right, nil
			}
			return middle, nil
		} else if array[middle] > value {
			right = middle - 1
			k--
		} else {
			k -= 2
			left = middle + 1
		}
	}
	return 0, errors.New("Not found on array")
}
