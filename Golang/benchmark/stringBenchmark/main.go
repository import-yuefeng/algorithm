package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	res := StringFmt()
	fmt.Println(res)
}

func StringPlus() string {
	a := "hello"
	a += "world"
	a += "你好"
	a += "世界"
	return a
}

func StringFmt() string {
	var a interface{} = "234"
	switch a.(type) {
	case string:
		fmt.Println(a, "string")
	case *int:
		fmt.Println("int")
	default:
		fmt.Println("none")
	}

	return fmt.Sprint("hello", "world", "你好", "世界")
}

func StringJoin() string {
	a := []string{"hello", "world", "你好", "世界"}
	return strings.Join(a, "")
}

func StringBuffer() string {
	var a bytes.Buffer
	a.WriteString("hello")
	a.WriteString("world")
	a.WriteString("你好")
	a.WriteString("世界")
	return a.String()
}

func StringBuilder() string {
	var a strings.Builder
	a.WriteString("hello")
	a.WriteString("world")
	a.WriteString("你好")
	a.WriteString("世界")
	return a.String()
}
