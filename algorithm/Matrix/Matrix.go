package main

import (
	"fmt"
	"math"
)

func main() {
	str := []string{"A1", "A2", "A3", "A4", "A5"}

	p := []int{30, 35, 15, 5, 10, 20}
	s := MatrixChain(p, len(str))
	PrintMatrix(s, 2, 5, str)
}

func MatrixChain(p []int, N int) [][]int {
	m := make([][]int, N+1)
	s := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		m[i] = make([]int, N+1)
	}
	for i := 0; i <= N; i++ {
		s[i] = make([]int, N+1)
	}
	for i := 0; i <= N; i++ {
		m[i][i] = 0
	}

	n := N
	for r := 2; r <= n; r++ {
		// for i := 1; i <= n-r+1; i++ {
		for i := n - r + 1; i >= 1; i-- {
			j := i + r - 1
			m[i][j] = math.MaxInt64
			for k := i; k < j; k++ {
				tmp := m[i][k] + m[k+1][j] + p[i]*p[k]*p[j]
				if tmp <= m[i][j] {
					m[i][j] = tmp
					s[i][j] = k
				}
			}
		}
	}
	for i, _ := range m {
		for _, v := range m[i] {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
	fmt.Println()
	for i, _ := range s {
		for _, v := range s[i] {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
	return s
}

func PrintMatrix(s [][]int, i, j int, str []string) {
	if i > j || i < 0 {
		fmt.Println("i or j is invaild")
		return
	}
	k := s[i][j]
	var z int
	fmt.Printf("(")
	for z = i; z <= k; z++ {
		fmt.Printf(str[z-1])
	}
	fmt.Printf(")")
	for ; z <= j; z++ {
		fmt.Printf(str[z-1])
	}
	fmt.Println()
	return
}
