// package main

// import "fmt"

// func main() {
// 	// k := 5
// 	task := [][]int{
// 		{1, 23}, {12, 28}, {25, 35}, {27, 80}, {36, 50},
// 	}
// 	qSort(task, 0, len(task)-1)
// 	res := schedule(task)
// 	fmt.Println(res)
// }

// func schedule(task [][]int) int {
// 	taskTime := task[0][1]
// 	taskSum := 1
// 	for _, v := range task[1:] {
// 		if v[0] > taskTime {
// 			taskTime = v[1]
// 			taskSum++
// 		}
// 	}
// 	return taskSum
// }

// func qSort(data interface{}, left, right int) {
// 	//reflect
// 	if len(data.([][]int)) == 0 || len(data.([][]int)[0]) == 0 {
// 		return
// 	}
// 	if left < right {
// 		l, r := left, right
// 		middle := data.([][]int)[(l+r)>>1][1]
// 		for l < r {
// 			for l < r && data.([][]int)[l][1] < middle {
// 				l++
// 			}
// 			for l < r && data.([][]int)[r][1] > middle {
// 				r--
// 			}
// 			if l >= r {
// 				break
// 			}
// 			if data.([][]int)[r][1] == data.([][]int)[l][1] && data.([][]int)[r][1] == middle {
// 				r--
// 			} else {
// 				data.([][]int)[r], data.([][]int)[l] = data.([][]int)[l], data.([][]int)[r]
// 			}
// 		}
// 		qSort(data, left, l-1)
// 		qSort(data, r+1, right)
// 	}
// 	return
// }
