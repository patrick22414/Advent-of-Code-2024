package day12

import (
	"testing"
)

func BenchmarkD12P1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1()
	}
}

func BenchmarkD12P2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2()
	}
}
