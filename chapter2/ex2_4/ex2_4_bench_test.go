package main

import (
	"tempconv/chapter2/popcount"
	"testing"
)

func BenchmarkPopCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(100) // 0.717s
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popCount(100) // 1.911s
	}
}
