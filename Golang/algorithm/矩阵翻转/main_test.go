package main

import (
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Start...")
	m.Run()
	log.Println("End...")
}

func TestReverseStandard1(m *testing.T) {
	data := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	reverseStandard1(data)
	if data[0][0] == 7 && data[1][0] == 4 && data[2][0] == 1 {
		m.Log("Success: ReverseStandard1")
		return
	}
	m.Error("Failure: ReverseStandard1")
	m.Log(data)
}

func TestReverseStandard2(m *testing.T) {
	data := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	reverseStandard2(data)
	if data[0][0] == 3 && data[1][0] == 6 && data[2][0] == 9 {
		m.Log("Success: ReverseStandard2")
		return
	}
	m.Error("Failure: ReverseStandard2")
	m.Log(data)
}

func TestReverseDiagonal1(m *testing.T) {
	data := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	reverseDiagonal1(data)
	if data[0][0] == 1 && data[1][0] == 2 && data[2][0] == 3 {
		m.Log("Success: ReverseDiagonal1")
		return
	}
	m.Error("Failure: ReverseDiagonal1")
	m.Log(data)
}

func TestReverseDiagonal2(m *testing.T) {
	data := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	reverseDiagonal2(data)
	if data[0][0] == 9 && data[1][0] == 8 && data[2][0] == 7 {
		m.Log("Success: ReverseDiagonal2")
		return
	}
	m.Error("Failure: ReverseDiagonal2")
	m.Log(data)
}

func BenchmarkReverseStantard1(b *testing.B) {
	data := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for i := 0; i < b.N; i++ {
		reverseStandard1(data)
	}
}
