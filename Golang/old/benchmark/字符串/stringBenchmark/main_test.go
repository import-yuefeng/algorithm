package main

import "testing"

func BenchmarkStringPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringPlus()
	}
}

func BenchmarkStringFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringFmt()
	}
}

func BenchmarkStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringJoin()
	}
}

func BenchmarkStringBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBuffer()
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringBuilder()
	}
}
