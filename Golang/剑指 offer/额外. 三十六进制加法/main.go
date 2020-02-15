package main

import (
	"bytes"
	"fmt"
)

var num2str []byte = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
	'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y', 'z',
}

func main() {
	var buffer bytes.Buffer
	res := sum([]byte{'1', 'b'}, []byte{'2', 'x'})
	fmt.Println(res)
	buffer.WriteString("aaaa")
	// buffer.Truncate(-1)
	// fmt.Println(buffer.String())
	var a []rune = []rune{'中', '国'}
	for i:=0; i<len(a); i++ {
		fmt.Printf("%c", a[i])
	}
	fmt.Println()
    var b string = "中国"
	for i:=0; i<len(b); i++ {
	    fmt.Printf("%c", b[i])
	}
}

func sum(a, b []byte) (res []byte) {
	p1, p2 := len(a)-1, len(b)-1
	if p1 == 0 {
		return b
	}
	if p2 == 0 {
		return a
	}
	takeOver := 0
	resLength := Max(len(a), len(b)) + 1
	res = make([]byte, resLength)
	resP := resLength - 1
	for p1 >= 0 || p2 >= 0 {
		x, y := 0, 0
		if p1 >= 0 {
			x = GetInt(a[p1])
			p1--
		}
		if p2 >= 0 {
			y = GetInt(b[p2])
			p2--
		}
		sumVal := x + y + takeOver
		takeOver = 0
		if sumVal >= 36 {
			// takeover...
			takeOver = sumVal / 36
			sumVal %= 36
		}
		res[resP] = num2str[sumVal]
		resP--
	}
	if takeOver != 0 {
		res[resP] = num2str[takeOver]
	}
	return res
}

func GetInt(a uint8) int {
	if a-'0' > 0 && a <= '9' {
		return int(a - '0')
	}
	return int(a-'a') + 10
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
