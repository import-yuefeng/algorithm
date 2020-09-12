// The MIT License (MIT)
// Copyright (c) 2019 import-yuefeng
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"time"
)

const (
	K int = 2
)

var (
	M    = make([]float64, K)
	tmpM = make([]float64, K)
)

func main() {

	randomArray := make([]float64, 15)
	rand.Seed(time.Now().Unix())
	for i := 0; i < 15; i++ {
		randomArray[i] = float64(rand.Intn(100))
	}
	for i := 0; i < K; i++ {
		M[i] = randomArray[i]
	}
	KMeans(randomArray[:], K)
}

// KMeans func print operation result
func KMeans(dataArray []float64, k int) {

	fmt.Println(dataArray)
	fmt.Println(M)
	count := 0
	for {
		kMeans := make(map[float64][]float64, k)
		for _, value := range dataArray {
			minValue := math.MaxFloat64
			var minIndex float64
			for _, subValue := range M {
				edValue := euclideanDistance(subValue, value)
				if edValue < minValue {
					minValue = edValue
					minIndex = subValue
				}
			}
			kMeans[minIndex] = append(kMeans[minIndex], value)
		}
		copy(tmpM, M)
		// deep copy slice
		M = M[:0]
		// clear slice
		fmt.Printf("--------%d--start---------------\n", count)
		for _, v := range kMeans {
			fmt.Println(v)
			var sum float64
			for _, value := range v {
				sum += value
			}
			fmt.Printf("M: %.3f\n", sum/float64(len(v)))
			M = append(M, sum/float64(len(v)))
		}
		fmt.Printf("--------%d--end---------------\n", count)
		count++
		if reflect.DeepEqual(tmpM, M) {
			break
		}
	}
}

// euclideanDistance func return two value euclidean-distance
func euclideanDistance(a, b float64) float64 {
	return math.Abs(float64(a - b))
}
