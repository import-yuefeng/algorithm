package main

import (
	"errors"
	"fmt"
)

func main() {
	array := []int{1, 3, 5, 8, 9, 10, 10, 12, 29, 36}
	// array := []int{11, 11, 11, 11, 11, 15, 16}
	// findFirstEqual(array[:], 7, 0, len(array)-1, 0)
	// findFirstEqual(array[:], 29, 0, len(array)-1, 0)
	if res, err := find(array[:], 5, 0, len(array)-1, 0); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func find(array []int, value, left, right, middle int) (site int, err error) {
	if left >= 0 && right >= 0 {
		if array[left] > value {
			return -1, errors.New("Not find elem on array")
		}
		for left <= right {
			middle = left + ((value - array[left]) / (array[right] - array[left]))
			fmt.Println("middle: ", middle)
			if array[middle] == value {
				return middle, nil
			} else if array[middle] > value {
				return find(array, value, left, middle-1, 0)
			}
			return find(array, value, middle+1, right, 0)
		}

	}
	return 0, errors.New("arguments error")

}
