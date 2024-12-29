package day09

import (
	"testing"
)

func BenchmarkD09P1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1()
	}
}

func BenchmarkD08P2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2()
	}
}
