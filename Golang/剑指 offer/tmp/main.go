package main

import (
	"fmt"
)

func main() {
	n := 0
	fmt.Scan(&n)
	strArray := make([]string, n)
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Scan(&tmp)
		strArray[i] = tmp
	}
	qSort(strArray, 0, len(strArray)-1)
	res := make([]string, 0)
	for i := 0; i < len(strArray); i++ {
		if i == 0 {
			res = append(res, strArray[i])
		} else {
			t := res[len(res)-1]
			if isSubString(t, strArray[i]) {
				continue
			} else {
				if t[len(t)-1] > strArray[i][len(strArray[i])-1] && len(strArray[i]) <= len(t) {
					continue
				} else if len(strArray[i]) > len(t) && isSubString(strArray[i], t) {
					res[len(res)-1] = strArray[i]
				} else if len(t) == 0 {
					res[len(res)-1] = strArray[i]
				} else {
					res = append(res, strArray[i])
				}
			}
		}
		fmt.Println(res)
	}
	length := 0
	for i := 0; i < len(res); i++ {
		// fmt.Println(res[i])
		length += len(res[i])
	}
	fmt.Printf("%d", length)
	return
}

func qSort(array []string, left, right int) {
	if left < right && left >= 0 {
		mid := array[left+((right-left)>>1)]
		l, r := left, right
		for l < r {
			for l < r && array[l] < mid {
				l++
			}
			for l < r && array[r] > mid {
				r--
			}
			if l >= r {
				break
			}
			if array[l] == array[r] && array[r] == mid {
				r--
			} else {
				array[l], array[r] = array[r], array[l]
			}
		}
		qSort(array, left, l-1)
		qSort(array, r+1, right)
	}
}

func isSubString(str, subStr string) bool {
	if len(subStr) == 0 {
		return true
	}
	if len(str) < len(subStr) {
		return false
	}
	p1, p2 := 0, 0
	for p1 < len(str) && p2 < len(subStr) {
		if str[p1] != str[p2] {
			return false
		}
		p1++
		p2++
	}
	if p1 == len(str) && p2 == len(subStr) {
		return true
	}
	return false
}
