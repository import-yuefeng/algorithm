package main

import "fmt"

func main() {
	Decompose(12)
	fmt.Println(total)
}

var total int = 1

func Decompose(n int) {
	if n <= 0 {
		total = 0
	}
	if n == 1 {
		total = 1
	} else {
		decompose(n)
	}
}

func decompose(n int) {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			total++
			decompose(n / i)
		}
	}
}
