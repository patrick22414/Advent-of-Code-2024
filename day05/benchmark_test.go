package day05

import (
	"testing"
)

func BenchmarkD05P1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1()
	}
}

func BenchmarkD05P2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2()
	}
}
