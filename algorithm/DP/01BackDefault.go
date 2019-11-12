package main

import "fmt"

func main() {
	weight := []int{2, 2, 6, 5, 4}
	value := []int{6, 3, 5, 4, 6}
	maxWeight := 10

	res := Backage01_3(weight, value, maxWeight, len(value))
	fmt.Println(res)
}

func Backage01(weightArray []int, valueArray []int, maxWeight int, numbers int) (maxValue int) {
	dp := make([][]int, numbers+1)
	for i := 0; i <= numbers; i++ {
		dp[i] = make([]int, maxWeight+1)
	}
	for i := 0; i <= numbers; i++ {
		dp[i][0] = 0
		dp[0][i] = 0
	}

	for i := 1; i <= numbers; i++ {
		for j := 1; j <= maxWeight; j++ {
			if j < weightArray[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = Max(dp[i-1][j], dp[i-1][j-weightArray[i-1]]+valueArray[i-1])
			}
		}
	}
	maxValue = dp[numbers][maxWeight]
	return
}

func Backage01_2(weightArray []int, valueArray []int, maxWeight int, numbers int) (maxValue int) {
	maxNumbers := numbers
	numbers = 1
	dp := make([][]int, numbers+1)
	for i := 0; i <= numbers; i++ {
		dp[i] = make([]int, maxWeight+1)
	}
	for i := 0; i <= 1; i++ {
		// 滚动数组的方式
		dp[i][0] = 0
	}
	for i := 0; i <= maxWeight; i++ {
		dp[0][i] = 0
	}

	for i := 1; i <= len(valueArray); i++ {
		for j := 1; j <= maxWeight; j++ {
			if j < weightArray[i-1] {
				dp[i%2][j] = dp[(i-1)%2][j]
			} else {
				dp[i%2][j] = Max(dp[(i-1)%2][j], dp[(i-1)%2][j-weightArray[i-1]]+valueArray[i-1])
			}
		}
	}
	maxValue = dp[maxNumbers%2][maxWeight]
	return
}

func Backage01_3(weightArray []int, valueArray []int, maxWeight int, numbers int) (maxValue int) {
	dp := make([]int, maxWeight+1)
	dp[0] = 0
	for i := 1; i <= numbers; i++ {
		for j := maxWeight; j >= 1; j-- {
			// 逆序很重要, 顺序是说不通的, 在求当前 dp 值需要的是上一层状态的 DP值
			// 如果顺序的话, 上一层 DP 的值 早已被覆盖了(被本层 刚刚运算覆盖了)
			if j < weightArray[i-1] {
				dp[j] = dp[j-1]
			} else {
				dp[j] = Max(dp[j-1], dp[j-weightArray[i-1]]+valueArray[i-1])
			}
		}
	}
	maxValue = dp[maxWeight]
	return
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
