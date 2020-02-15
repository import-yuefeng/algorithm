package main

import (
	"fmt"
	// "time"
)

func main() {
	myNumber1 := []byte{'9', '0', '2', '0', '0', '9'}
	myNumber2 := []byte{'1', '2', '1', '3', '1', '4'}
	res := Sum(myNumber1, myNumber2)
	fmt.Println(string(res))
	// t1 := time.Now() // get current time
	// PrintToMaxNDigits(myNumber, len(myNumber), 0)
	// for !Increment(myNumber) {
	// 	fmt.Println(string(myNumber))
	// }
	// elapsed := time.Since(t1)
	// fmt.Println("run time: ", elapsed)

}

func Sum(a, b []byte) []byte {
	var res []byte
	if len(a) > len(b) {
		res = make([]byte, len(a)+1)
	} else {
		res = make([]byte, len(b)+1)
	}
	p1, p2 := len(a)-1, len(b)-1
	p3 := len(res) - 1
	takeOver := 0
	for p3 >= 0 {
		var x, y int
		if p1 >= 0 {
			x = int(a[p1] - '0')
			p1--
		}
		if p2 >= 0 {
			y = int(b[p2] - '0')
			p2--
		}
		val := x + y + takeOver
		takeOver = 0
		if val >= 10 {
			takeOver = val / 10
			val %= 10
		}
		res[p3] = byte(val) + '0'
		p3--
	}
	return res
}

func PrintToMaxNDigits(num []byte, length, index int) {
	if length == index {
		fmt.Println(string(num))
		return
	} else {
		for i := 0; i < 10; i++ {
			num[index] = byte(i) + '0'
			PrintToMaxNDigits(num, length, index+1)
		}
	}
}

func Increment(num []byte) bool {
	isOverflow := false
	takeOver := 0
	for i := len(num) - 1; i >= 0; i-- {
		val := int(num[i]-'0') + takeOver
		takeOver = 0
		if i == len(num)-1 {
			val += 1
		}
		if val >= 10 {
			if i == 0 {
				isOverflow = true
				return isOverflow
			}
			takeOver = val / 10
			val %= 10
			num[i] = byte(val) + '0'
		} else {
			num[i] = byte(val) + '0'
			break
		}
	}
	return isOverflow
}
