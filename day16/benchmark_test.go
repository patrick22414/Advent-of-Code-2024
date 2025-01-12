package day16

import (
	"testing"
)

func BenchmarkD16P1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1()
	}
}

func BenchmarkD16P2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2()
	}
}
