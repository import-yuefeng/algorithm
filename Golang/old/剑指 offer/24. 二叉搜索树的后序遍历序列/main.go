package main

import "fmt"

func main() {
	// res := VerifySequenceOfBST([]int{5, 7, 6, 9, 11, 10, 8})
	res := VerifySequenceOfBST([]int{5, 7, 6, 9, 11, 8, 10})
	fmt.Println(res)
}

func VerifySequenceOfBST(sequence []int) bool {
	if len(sequence) == 0 {
		return false
	}
	root := sequence[len(sequence)-1]
	point := 0
	for i, v := range sequence {
		if v > root {
			point = i
		}
	}
	for j := point; j < len(sequence); j++ {
		if sequence[j] < root {
			return false
		}
	}
	left, right := true, true
	if point > 0 {
		left = VerifySequenceOfBST(sequence[:point])
	}
	if point < len(sequence)-1 {
		right = VerifySequenceOfBST(sequence[point:])
	}
	return left && right
}
