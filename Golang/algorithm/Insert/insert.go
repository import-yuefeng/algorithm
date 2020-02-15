package main

import (
	"fmt"
)

func main() {
	var a = [10]int{2, 5, 1, 10, 2, 4, 3, 2, 3, 7}
	// rand.Seed(time.Now().Unix())
	// for i := 0; i < 100; i++ {
	// 	a[i] = rand.Intn(100)
	// }

	fmt.Println(a)
	// CountSort(a[:], 0, len(a)-1)

	// fmt.Println(qSelect(a[:], 0, len(a)-1, 1))
	// fmt.Println(qSelect(a[:], 0, len(a)-1, 12))

	mergeSort(a[:], 0, len(a)-1)
	// ShellSort(a[:])
	// InsertSort(a[:])
	// SelectSort(a[:])
	// BubblingSort(a[:])
	// qSort(a[:], 0, len(a)-1)
	fmt.Println(a)

	// inplaceMerge([]int{1, 3, 9, 10, 13, 2, 8, 10, 11, 15}, 0, 4, 9)
	var c uint16 = 34520
	fmt.Println(c>>8 | c<<8)
}

func inplaceMerge(num []int, left, mid, right int) {
	for i := right; i > mid; i-- {
		if num[i] < num[mid] {
			num[i] ^= num[mid]
			num[mid] ^= num[i]
			num[i] ^= num[mid]
			for j := mid; j > left && num[j] < num[j-1]; j-- {
				num[j], num[j-1] = num[j-1], num[j]
			}
		}
	}
}

func InsertSort(num []int) {
	for i := 1; i < len(num); i++ {
		for j := i; j > 0 && num[j] < num[j-1]; j-- {
			num[j], num[j-1] = num[j-1], num[j]
		}
	}
}

func BubblingSort(num []int) {
	for i := 0; i < len(num)-1; i++ {
		for j := 0; j < len(num)-1; j++ {
			if num[j] > num[j+1] {
				num[j] ^= num[j+1]
				num[j+1] ^= num[j]
				num[j] ^= num[j+1]
			}
		}
	}
}

func ShellSort(num []int) {
	gap := len(num) >> 1
	for gap > 0 {
		for i := gap; i < len(num); i++ {
			for j := i; j-gap >= 0 && num[j-gap] > num[j]; j -= gap {
				num[j], num[j-gap] = num[j-gap], num[j]
			}
		}
		gap /= 2
	}
}

func SelectSort(num []int) {
	for i := 0; i < len(num); i++ {
		tmp := num[i]
		index := i
		for j := i; j < len(num); j++ {
			if num[j] < tmp {
				tmp = num[j]
				index = j
			}
		}
		num[i], num[index] = num[index], num[i]
	}
}

func qSort(num []int, left, right int) {
	if left < right && left >= 0 {
		l, r := left, right
		mid := num[l+((r-l)>>1)]
		for l < r {
			for l < r && num[l] < mid {
				l++
			}
			for l < r && num[r] > mid {
				r--
			}
			if l >= r {
				break
			}
			if num[r] == num[l] && num[r] == mid {
				r--
			} else {
				num[l], num[r] = num[r], num[l]
			}
		}
		qSort(num, left, l-1)
		qSort(num, r+1, right)
	}
}

func mergeSort(num []int, left, right int) {
	if left < right && left >= 0 {
		mid := left + ((right - left) >> 1)
		mergeSort(num, left, mid)
		mergeSort(num, mid+1, right)
		// merge(num, left, mid, right)
		inplaceMerge(num, left, mid, right)
	}
}

func merge(num []int, left, mid, right int) {
	p1, p2 := left, mid+1
	res := make([]int, right-left+1)
	tmp := 0
	for p1 <= mid && p2 <= right {
		if num[p1] < num[p2] {
			res[tmp] = num[p1]
			p1++
		} else {
			res[tmp] = num[p2]
			p2++
		}
		tmp++
	}
	for p1 <= mid {
		res[tmp] = num[p1]
		p1++
		tmp++
	}
	for p2 <= right {
		res[tmp] = num[p2]
		p2++
		tmp++
	}
	for i := left; i <= right; i++ {
		num[i] = res[i-left]
	}
}

func CountSort(num []int, left, right int) {
	min := 1 >> 31
	max := -1 << 31
	for i := left; i <= right; i++ {
		if min > num[i] {
			min = num[i]
		}
		if max < num[i] {
			max = num[i]
		}
	}
	tmp := make([]int, max-min+1)
	for i := left; i <= right; i++ {
		tmp[num[i]-min] += 1
	}
	index := 0
	for i, v := range tmp {
		for v > 0 {
			num[index] = i + min
			v--
			index++
		}
	}
}

// func RadixSort(num []int) {
// 	maxVal := qSelect(num, 0, len(num)-1, len(num))
// 	digit := 0
// 	for maxVal != 0 {
// 		maxVal /= 10
// 		digit++
// 	}
// 	for i := 1; i <= digit; i++ {
// 		tmpArray := make([]int, len(num))
// 		countArray := make([]int, 10)
// 		for j := 0; j < len(num); j++ {
// 			tmpDigit := (num[j] / math.Pow(10, i-1)) - ((num[j] / math.Pow(10, i)) * 10)
// 			countArray[tmpDigit] += 1
// 		}
// 		for m := 1; m < 10; m++ {
// 			countArray[m] += countArray[m-1]
// 		}
//         for (int n = ArrayToSort.Length - 1; n >= 0; n--)
//         {
//             int tmpSplitDigit = ArrayToSort[n] / (int)Math.Pow(10,k - 1) - (ArrayToSort[n]/(int)Math.Pow(10,k)) * 10;
//             tmpArray[tmpCountingSortArray[tmpSplitDigit]-1] = ArrayToSort[n];
//             tmpCountingSortArray[tmpSplitDigit] -= 1;
//         }

// 	}
// }

func qSelect(num []int, left, right, k int) int {
	if left < right {
		mid := num[(left+right)>>1]
		l, r := left, right
		for l < r {
			for l < r && num[l] < mid {
				l++
			}
			for l < r && num[r] > mid {
				r--
			}
			if l >= r {
				break
			}
			if num[r] == num[l] && num[r] == mid {
				r--
			} else {
				num[l], num[r] = num[r], num[l]
			}
		}
		if l-left+1 == k {
			return mid
		} else if l-left+1 > k {
			return qSelect(num, left, l-1, k)
		} else {
			return qSelect(num, r+1, right, k-l-1)
		}
	}
	return num[left]
}
