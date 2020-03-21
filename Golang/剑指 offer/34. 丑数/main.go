package main

func main() {
	res := getUglyNumber(1500)
	res++
	func() {
		res++
	}()
	A()
}

func A() func() int {
	x := 2
	return func() int {
		return x
	}
}

func getUglyNumber(n int) int {
	if n <= 0 {
		return 1
	}

	dp := make([]int, n+1)
	p, p2, p3, p5 := 1, 1, 1, 1
	dp[1] = 1
	var nextVal int
	for ; p <= n; p++ {
		nextVal = min(min(dp[p2]*2, dp[p3]*3), dp[p5]*5)
		dp[p] = nextVal
		for dp[p2]*2 <= nextVal {
			p2++
		}
		for dp[p3]*3 <= nextVal {
			p3++
		}
		for dp[p5]*5 <= nextVal {
			p5++
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
