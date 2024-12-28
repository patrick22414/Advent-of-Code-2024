package day14

import (
	"testing"
)

func BenchmarkD14P1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1()
	}
}

func BenchmarkD14P2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2()
	}
}
