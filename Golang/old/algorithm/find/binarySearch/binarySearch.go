// package binarySearch
package main

import (
	"errors"
	"fmt"
)

func main() {
	// array := []int{1, 3, 5, 8, 9, 10, 10, 10, 29, 36}
	array := []int{11, 11, 11, 11, 11, 15, 16}
	// findFirstEqual(array[:], 7, 0, len(array)-1, 0)
	// findFirstEqual(array[:], 29, 0, len(array)-1, 0)
	if res, err := findEndSmall(array[:], 10, 0, len(array)-1, 0); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

func find(array []int, value int, left, right, middle int) (site int, err error) {
	if array[left] > value {
		return -1, errors.New("not find element on array")
	}
	if left >= 0 && right >= 0 {
		for left <= right {
			middle = (left + right) / 2
			if array[middle] == value {
				return middle, nil
			} else if array[middle] > value {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
		return -1, errors.New("not find element on array")
	}
	return -1, errors.New("left or right arguments is invaild")
}

func findFirstEqual(array []int, value int, left, right, middle int) (site int, err error) {
	if array[left] > value {
		return -1, errors.New("not find element on array")
	}
	if left >= 0 && right >= 0 {
		for left <= right {
			middle = (left + right) / 2
			if array[middle] == value {
				if array[middle-1] != value {
					return middle, nil
				}
				right = middle - 1
			} else if array[middle] > value {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
		return -1, errors.New("not find element on array")
	}
	return -1, errors.New("left or right arguments is invaild")
}

func findFirstEqual2(array []int, value int, left, right, middle int) (site int, err error) {
	if array[left] > value {
		return -1, errors.New("not find element on array")
	}
	if left >= 0 && right >= 0 {
		if left > right {
			return -1, errors.New("not find element on array")
		}
		middle = (left + right) / 2
		if array[middle] >= value {
			if array[middle] == value {
				if array[middle-1] != value {
					return middle, nil
				}
			}
			return findFirstEqual2(array, value, left, middle-1, middle)
		}
		return findFirstEqual2(array, value, middle+1, right, middle)
	}
	return -1, errors.New("left or right arguments is invaild")
}

func findEndEqual(array []int, value int, left, right, middle int) (site int, err error) {
	if array[left] > value {
		return -1, errors.New("not find element on array")
	}
	if left >= 0 && right >= 0 {
		for left <= right {
			middle = (left + right) / 2
			if array[middle] == value {
				if array[middle+1] == value {
					left = middle + 1
				} else {
					return middle, nil
				}
			} else if array[middle] > value {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
		return -1, errors.New("not find element on array")
	}
	return -1, errors.New("left or right arguments is invaild")
}

func findFirstBiggerEqual(array []int, value int, left, right, middle int) (site int, err error) {
	if array[left] > value {
		return -1, errors.New("not find element on array")
	}
	if left >= 0 && right >= 0 {
		for left <= right {
			middle = (left + right) / 2
			if array[middle] < value {
				left = middle + 1
			} else {
				if array[middle-1] >= value {
					right = middle - 1
				} else {
					return middle, nil
				}
			}
		}
		return -1, errors.New("not find element on array")
	}
	return -1, errors.New("left or right arguments is invaild")
}

func findEndSmallEqual(array []int, value int, left, right, middle int) (site int, err error) {
	if array[left] > value {
		return -1, errors.New("not find element on array")
	}
	if left >= 0 && right >= 0 {
		for left <= right {
			middle = (left + right) / 2
			if array[middle] > value {
				right = middle - 1
			} else {
				if array[middle+1] <= value {
					left = middle + 1
				} else {
					return middle, nil
				}
			}
		}
		return -1, errors.New("not find element on array")
	}
	return -1, errors.New("left or right arguments is invaild")
}

func findFirstBigger(array []int, value int, left, right, middle int) (site int, err error) {
	if array[left] > value {
		return left, nil
	}
	if left >= 0 && right >= 0 {
		for left <= right {
			middle = (left + right) / 2
			if array[middle] > value {
				if array[middle-1] > value {
					right = middle - 1
				} else {
					return middle, nil
				}
			} else {
				left = middle + 1
			}
		}
		return -1, errors.New("not find element on array")
	}
	return -1, errors.New("left or right arguments is invaild")
}

func findEndSmall(array []int, value int, left, right, middle int) (site int, err error) {
	if array[left] > value {
		return -1, errors.New("not find element on array")
	}
	if left >= 0 && right >= 0 {
		for left <= right {
			middle = (left + right) / 2
			if array[middle] >= value {
				right = middle - 1
			} else {
				if array[middle+1] < value {
					left = middle + 1
				} else {
					return middle, nil
				}
			}
		}
		return -1, errors.New("not find element on array")
	}
	return -1, errors.New("left or right arguments is invaild")
}
