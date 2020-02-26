package main

import (
	"fmt"
)

func main() {

	fmt.Println("hello world")
	res := B([]int{4, 7, 2, 9, 5, 2})

	fmt.Println(res)
}

func B(num []int) int {
	sum, dp := make([][]int, len(num)), make([][]int, len(num))
	for i, _ := range dp {
		dp[i] = make([]int, len(num))
		sum[i] = make([]int, len(num))
	}

	for i := 0; i < len(num); i++ {
		for j := i; j < len(num); j++ {
			for q := i; q <= j; q++ {
				sum[i][j] += num[q]
			}
		}
	}
	dp[0][0] = sum[0][0]

	for i := 0; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			dp[i][j] = sum[i][j] - min(dp[i+1][j], dp[i][j-1])
		}
	}
	a := dp[0][len(num)-1]
	b := sum[0][len(num)-1] - a
	return a - b
}

func min(a, b int) int {
	if a > b {
		return a
	}
	return b
}
