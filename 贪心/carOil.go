// package main

// import "fmt"

// func main() {
// 	n, k := 3, 7
// 	distance := []int{1, 2, 3, 4, 5, 1, 6, 6}
// 	res := schedule(n, k, distance)
// 	if res == -1 {
// 		fmt.Println("No Solution")
// 	}
// 	fmt.Println(res)
// 	return
// }

// func schedule(n int, k int, distance []int) int {
// 	if n <= 0 || k <= 0 {
// 		return -1
// 	}
// 	// sum := 0
// 	// for _, v := range distance[1:] {
// 	// 	sum += v
// 	// }
// 	// if sum > n*k {
// 	// 	return -1
// 	// }

// 	now := n
// 	addOilNum := 1
// 	for _, v := range distance[1:] {
// 		if v <= now {
// 			now -= v

// 		} else {
// 			now = now - v + n
// 			addOilNum++
// 		}
// 	}

// 	if addOilNum >= k {
// 		return -1
// 	}
// 	return addOilNum
// }
