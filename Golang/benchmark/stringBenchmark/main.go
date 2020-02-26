package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
}

func StringPlus() string {
	a := "hello"
	a += "world"
	a += "你好"
	a += "世界"
	return a
}

func StringFmt() string {
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
