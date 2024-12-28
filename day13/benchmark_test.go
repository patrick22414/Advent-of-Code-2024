package day13

import (
	"testing"
)

func BenchmarkD13P1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1()
	}
}

func BenchmarkD13P2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2()
	}
}