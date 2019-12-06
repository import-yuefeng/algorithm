package main

import "fmt"

func main() {
	move(3, 'A', 'C', 'B')
}

func move(n int, from, buffer, to byte) {
	if n == 1 {
		fmt.Printf("%d %c -> %c\n", n, from, to)
		return
	}
	move(n-1, from, to, buffer)
	fmt.Printf("%d %c -> %c\n", n, from, to)
	move(n-1, buffer, from, to)
}
