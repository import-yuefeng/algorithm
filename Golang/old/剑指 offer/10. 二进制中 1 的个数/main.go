package main

import (
	"fmt"
)

func main() {
	// res := NumberOf1_1(-9)
	// fmt.Println(res)
	res := NumberOf1_2(9)
	fmt.Println(res)

}

func NumberOf1_1(n int) int {
	count := 0
	for n != 0 {
		if n&1 == 1 {
			count++
		}
		n >>= 1
	}
	return count
}

func NumberOf1_2(n int) int {
	count := 0
	for n != 0 {
		n &= (n - 1)
		count++
	}
	return count
}
