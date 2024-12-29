package day06

import (
	"testing"
)

func BenchmarkD06P1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1()
	}
}

func BenchmarkD06P2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2()
	}
}
