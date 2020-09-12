package main

import (
	"fmt"
)

func main() {
	Array := []int{4, 6, 1, 2, 65, 33, 2, 1}
	// mergeSortInplace(Array, 0, len(Array)-1)
	// BubblingSort(Array, 0, len(Array)-1)
	// CountSort(Array, 0, len(Array)-1)
	RadixSort(Array, 0, len(Array)-1)
	// HeapSort(Array, 0, len(Array)-1)
	fmt.Println(Array)
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

func mergeSortInplace(num []int, left, right int) {
	if left < right && left >= 0 {
		mid := left + ((right - left) >> 1)
		mergeSortInplace(num, left, mid)
		mergeSortInplace(num, mid+1, right)
		mergeInplace(num, left, mid, right)
		return
	}
	return
}

func mergeInplace(num []int, left, mid, right int) {
	// inplace implement...
	l, r := mid, right
	for r > l {
		if num[r] < num[l] {
			num[l] ^= num[r]
			num[r] ^= num[l]
			num[l] ^= num[r]
			for i := l; i > 0 && num[i] < num[i-1]; i-- {
				num[i], num[i-1] = num[i-1], num[i]
			}
		}
		r--
	}
}

func ShellSort(num []int, left, right int) {
	gap := len(num) / 2
	for gap > 0 {
		for i := gap; i < len(num); i += gap {
			for j := i; j-gap >= 0 && num[j] < num[j-gap]; j -= gap {
				num[j], num[j-gap] = num[j-gap], num[j]
			}
		}
		gap--
	}
}

func BubblingSort(num []int, left, right int) {
	for i := left; i < right; i++ {
		for j := left; j <= right-1-i; j++ {
			if num[j] > num[j+1] {
				num[j], num[j+1] = num[j+1], num[j]
			}
		}
	}
}

func CountSort(num []int, left, right int) {
	max, min := -1<<31, 1<<31
	for i := left; i <= right; i++ {
		if max < num[i] {
			max = num[i]
		}
		if min > num[i] {
			min = num[i]
		}
	}
	length := max - min + 1
	tmp := make([]int, length)
	for _, v := range num {
		tmp[v-min]++
	}
	point := 0
	for i, v := range tmp {
		for v > 0 {
			// value - min = i
			num[point] = i + min
			v--
			point++
		}
	}
	tmp = nil
}

func HeapSort(data []int, left, right int) {
	heap := NewHeap(data)
	heap.BuildHeap()
	// res := make([]int, len(data))
	// point := 0
	for heap.size > 0 {
		_, err := heap.Delete()
		if err != nil {
			fmt.Println(err)
			return
		}
		// 	res[point] = val
		// 	point++
	}
	// fmt.Println(res)
}

func RadixSort(num []int, left, right int) {
	bucket := make([]*Queue, 10)
	for idx, _ := range bucket {
		bucket[idx] = NewQueue()
	}
	maxValue := -1 << 31
	for _, v := range num {
		if maxValue < v {
			maxValue = v
		}
	}
	length := getBit(maxValue)
	tmp := make([]int, len(num))
	copy(tmp, num)
	for i := 1; i <= length; i++ {
		for _, v := range tmp {
			bucket[GetCurBit(v, i)].Push(v)
		}
		point := 0
		for idx, _ := range bucket {
			for !bucket[idx].IsEmpty() {
				val, err := bucket[idx].Pop()
				if err != nil {
					fmt.Println(err)
					return
				}
				tmp[point] = val
				point++
			}
		}
	}
	copy(num, tmp)
}

func GetCurBit(number, i int) int {
	base := 1
	for i > 1 {
		base *= 10
		i--
	}
	return number / base % 10
}

func getBit(number int) int {
	count := 0
	for number != 0 {
		count++
		number /= 10
	}
	return count
}
