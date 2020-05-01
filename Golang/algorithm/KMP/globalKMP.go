// package main

// import "fmt"

// var (
// 	pat  string
// 	data [][256]int
// )

// func Init(pattern string) {

// 	if len(pattern) <= 0 {
// 		return
// 	}
// 	pat = pattern
// 	data = make([][256]int, len(pat))
// 	return
// }

// func BuildKMP() {
// 	// init default state
// 	data[0][int(pat[0])] = 1
// 	// preState define
// 	X := 0
// 	for i := 1; i < len(pat); i++ {
// 		patChar := pat[i]
// 		for j := 0; j < 256; j++ {
// 			// to next state
// 			if j == int(patChar) {
// 				data[i][j] = i + 1
// 			} else { // return preState
// 				data[i][j] = data[X][j]
// 			}
// 		}
// 		// update preState
// 		X = data[i][int(patChar)]
// 	}
// }

// func SearchString(str string) int {
// 	flag := 0
// 	for i, v := range str {
// 		flag = data[flag][int(v)]
// 		if flag >= len(pat) {
// 			return i - len(pat) + 1
// 		}
// 	}
// 	return -1
// }

// func main() {
// 	Init("abaababaab")
// 	BuildKMP()
// 	res := SearchString("baababaababaababaa")
// 	fmt.Println(res)
// }
