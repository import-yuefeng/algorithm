// package denaryConvert
package main

import "fmt"

var (
	array = [26]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
)

func main() {
	convert(10)
	convert(27)
	convert(28)
	convert(28999)
}

func convert(n int) string {
	if n == 0 {
		return ""
	} else if n > 0 {
		n--
	}
	s := ""
	for n >= 26 {
		s += string(array[0])
		n -= 26
	}
	s += string(array[n])
	fmt.Println(s)
	return s
}
