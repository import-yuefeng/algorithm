package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// res := StringFmt()
	// fmt.Println(res)
	// res := make(map[int]int)
	// for i := 0; i < 100000; i++ {
	tmp := RandomSelectNumber(3)
	fmt.Println(tmp)
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

func RandomSelectNumber(k int) []int {
	rawList := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = rawList[i]
	}
	for i := k; i < len(rawList); i++ {
		rand.Seed(time.Now().Unix())
		t := rand.Intn(i + 1)
		if t < k {
			res[t] = rawList[i]
		}
	}
	return res
}
