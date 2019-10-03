// package doubleDimensionalArray
package main

import (
	"fmt"
)

type s1 []int
type s2 []s1

func main() {
	array := s2{
		s1{1, 2, 8, 9},
		s1{2, 4, 9, 12},
		s1{4, 7, 10, 13},
		s1{6, 8, 11, 15},
	}
	fmt.Println(find(array[:][:], 6, 4, 4))
}

func find(array s2, elem int, rows int, columns int) (result bool) {
	row, cloumn := 0, columns-1
	result = false
	if array == nil {
		return
	}
	for row < rows && cloumn >= 0 {
		if array[row][cloumn] > elem {
			cloumn--
		} else if array[row][cloumn] < elem {
			row++
		} else {
			result = true
			return
		}
	}
	return
}
