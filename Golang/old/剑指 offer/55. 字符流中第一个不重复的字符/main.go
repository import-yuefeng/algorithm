package main

import "fmt"

func main() {
	a := []byte{'a', 'c', 'c', 'a', 'z', '1'}
	res := FirstAppearingOnce(a)
	fmt.Println(string(res))
}

func FirstAppearingOnce(input []byte) byte {
	charMap := make([]int, 256)
	for i, _ := range charMap {
		charMap[i] = -1
	}
	for i, v := range input {
		if charMap[v] == -1 {
			charMap[v] = i
		} else if charMap[v] >= 0 {
			charMap[v] = -2
		}
	}
	minIndex := 0
	min := 1 << 31
	for i, v := range charMap {
		if v > 0 && v < min {
			min = v
			minIndex = i
		}
	}
	return byte(minIndex)
}
