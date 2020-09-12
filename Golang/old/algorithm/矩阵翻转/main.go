package main

import "fmt"

func reverseDiagonal1(data [][]int) {
	// 1 2 3     1 4 7
	//  ↖    ==>
	// 4 5 6     2 5 8
	// 	  ↖  ==>
	// 7 8 9     3 6 9
	// 根据 左上-右下 对角线进行 翻转
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if j > i {
				data[i][j], data[j][i] = data[j][i], data[i][j]
			}
		}
	}
}

func reverseDiagonal2(data [][]int) {
	// 1 2 3     9 6 3
	//    ↗  ==>
	// 4 5 6     8 5 2
	// ↗     ==>
	// 7 8 9     7 4 1
	// 根据 右上-左下 对角线进行 翻转
	for i := 0; i < len(data); i++ {
		length := len(data)
		for j := 0; j < len(data[i]); j++ {
			if j < length-1-i {
				data[i][j], data[length-1-j][length-i-1] = data[length-1-j][length-i-1], data[i][j]
			}
		}
	}
}

func reverseStandard1(data [][]int) {
	// 1 2 3     7 8 9
	// ------ >
	// 4 5 6     4 5 6
	//------- >
	// 7 8 9     1 2 3
	// 根据 横着的中轴线 4 5 6 进行 翻转
	for i := 0; i <= len(data)/2; i++ {
		for j := 0; j < len(data[i]); j++ {
			data[i][j], data[len(data)-i-1][j] = data[len(data)-i-1][j], data[i][j]
		}
	}
}

func reverseStandard2(data [][]int) {
	// 1 |2| 3 -->  3 |2| 1
	// 4 |5| 6 -->  6 |5| 4
	// 7 |8| 9 -->  9 |8| 7
	// 根据 竖着的中轴线 2 5 8 进行 翻转
	for i := 0; i < len(data); i++ {
		for j := 0; j <= len(data[i])/2; j++ {
			data[i][j], data[i][len(data)-j-1] = data[i][len(data)-j-1], data[i][j]
		}
	}
}

func main() {
	test := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	// reverseDiagonal2(test)
	reverseStandard2(test)
	fmt.Println(test)
}
