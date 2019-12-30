// package stringReplace
package main

import (
	"bytes"
	"fmt"
)

func main() {
	str := "We are Happy!"
	str = " "
	newStr := replace(str)
	fmt.Println(newStr)
}

func replace(str string) (newStr string) {
	originLen, blackLen, newLen := 0, 0, 0
	black := " "
	newStr = ""
	if str == "" {
		return ""
	}
	for _, v := range str {
		if string(v) == black {
			blackLen++
		}
		originLen++
	}
	newLen = originLen + 2*blackLen - 1
	originLen--

	var buffer bytes.Buffer
	buffer.WriteString(str)
	for i := 0; i < 2*blackLen; i++ {
		buffer.WriteString(black)
	}
	str = buffer.String()
	myByte := []byte(str)

	for newLen != originLen {
		if string(str[originLen]) == black {
			myByte[newLen] = '0'
			newLen--
			myByte[newLen] = '2'
			newLen--
			myByte[newLen] = '%'
			newLen--
			originLen--
		} else {
			myByte[newLen] = myByte[originLen]
			newLen--
			originLen--
		}
	}
	return string(myByte)
}
