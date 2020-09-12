// package Nqueen
package main

func main() {

}

func NQueen(row int, res [][]int) {
	if row == len(res) {
		return
		// res = append(res, )
	} else {
		for col := 0; col < len(res); col++ {
			if !vaild(res, row, col) {
				continue
			}
			res[row][col] = 1
			NQueen(row+1, res)
			res[row][col] = 0
		}
	}
	return
}

func vaild(res [][]int, row int, col int) bool {
	for i := row; i >= 0; i-- {
		if res[i][col] == 1 {
			return false
		}
	}
	i := row - 1
	j := col - 1
	for {
		if i < 0 || j < 0 {
			break
		}
		if res[i][j] == 1 {
			return false
		}
		i--
		j--
	}
	i = row - 1
	j = col + 1
	for {
		if i < 0 || j > len(res) {
			break
		}
		if res[i][j] == 1 {
			return false
		}
		i--
		j++
	}
	return true
}
