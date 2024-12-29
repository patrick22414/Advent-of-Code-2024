package day15

import (
	"testing"
)

func BenchmarkD15P1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1()
	}
}

func BenchmarkD15P2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2()
	}
}
