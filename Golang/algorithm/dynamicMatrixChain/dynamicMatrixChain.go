package main

import "fmt"

func main() {

	// MatrixChain()
}

func MatrixChain(p []int, n int, m [][]int, s [][]int) {
	for i := 0; i <= n; i++ {
		m[i][i] = 0
	}
	for r := 5; r >= 2; r-- {
		for i := 1; i <= n-r+1; i++ {
			j := i + r - 1
			fmt.Println(r, j)
		}
	}
}
