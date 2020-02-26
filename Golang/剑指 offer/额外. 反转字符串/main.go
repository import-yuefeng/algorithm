package main

import (
	"fmt"
	"strings"
)

func main() {
	res := ReverseString("   www    baidu com")
	fmt.Println(res)

}

func ReverseString(a string) string {
	if len(a) <= 1 {
		return a
	}
	array := make([]rune, len(a))
	for i, v := range a {
		array[i] = v
	}
	ReverseChar(array, 0, len(array)-1)
	var res strings.Builder
	start, end := 0, 0
	for start < len(array) {
		if array[end] == ' ' {
			if start < end {
				ReverseChar(array, start, end-1)
				start = end
			} else {
				start++
				end++
			}
		} else {
			end++
		}
	}
	for _, v := range array {
		res.WriteRune(v)
	}
	return res.String()
}

func ReverseChar(res []rune, left, right int) {
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
}
