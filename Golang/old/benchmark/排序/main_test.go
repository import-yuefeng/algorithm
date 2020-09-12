package main

import (
	"math/rand"
	"testing"
	"time"
)

var MaxValueRange int = 5000

func RandList() []int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	var data []int = make([]int, MaxValueRange)
	for i := 0; i < MaxValueRange; i++ {
		data[i] = r.Intn(MaxValueRange)
	}
	return data

}

func BenchmarkQSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data := RandList()
		qSort(data, 0, len(data)-1)
	}
	// b.Log(data)
}

func BenchmarkMergeSortInplace(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data := RandList()
		mergeSortInplace(data, 0, len(data)-1)
	}
	// b.Log(data)
}

func BenchmarkShellSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data := RandList()
		ShellSort(data, 0, len(data)-1)
	}
}

func BenchmarkBubblingSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data := RandList()
		BubblingSort(data, 0, len(data)-1)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data := RandList()
		HeapSort(data, 0, len(data)-1)
	}
}

func BenchmarkCountSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data := RandList()
		CountSort(data, 0, len(data)-1)
	}
}

func BenchmarkRadixSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data := RandList()
		RadixSort(data, 0, len(data)-1)
	}
}
