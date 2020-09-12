package main

import (
	"errors"
	"fmt"
)

func main() {
	res, err := sqrt_2(1, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func sqrt_1(x int, t int) (res int, err error) {
	if x < 0 {
		return -1, errors.New("invalid")
	}
	left, right := 0, x
	for left <= right {
		mid := left + ((right - left) >> 1)
		tmpVal := mid * mid
		if tmpVal == x || (mid-left <= t && right-mid <= t) {
			return mid, nil
		} else if tmpVal > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left, nil
}

func sqrt_2(x float64, t float64) (res float64, err error) {
	if x < 0 {
		return -1, errors.New("invalid")
	}
	if t < 0 {
		t *= -1
	}
	last, cur := 0.0, x
	for cur-last > t || last-cur > t {
		last = cur
		cur = (cur + x/cur) / 2
	}
	return cur, nil
}
