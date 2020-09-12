// package mergeArray

package main

import (
	"fmt"
)

func main() {

	array_1 := []int{1, 3, 4, 7, 9, 15}
	array_2 := []int{1, 5, 6, 8, 13, 17}
	array_3 := []int{}
	result := merge(array_1, array_2)
	result = merge(array_1, array_3)

	fmt.Println(result)
}

func merge(a, b []int) []int {
	if len(a) == 0 && len(b) == 0 {
		return []int{}
	}
	result := make([]int, len(a)+len(b))
	point, pointA, pointB := len(a)+len(b)-1, len(a)-1, len(b)-1
	tmpPoint, tmpArray := 0, []int{}

	for pointA >= 0 && pointB >= 0 {
		if a[pointA] > b[pointB] {
			result[point] = a[pointA]
			pointA--
		} else {
			result[point] = b[pointB]
			pointB--
		}
		point--
	}
	if pointA < 0 {
		tmpPoint = pointB
		tmpArray = b
	} else {
		tmpPoint = pointA
		tmpArray = a
	}
	for ; tmpPoint >= 0; tmpPoint-- {
		result[point] = tmpArray[tmpPoint]
		point--
	}
	return result
}
