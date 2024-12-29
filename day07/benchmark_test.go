package day07

import (
	"testing"
)

func BenchmarkD07P1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1()
	}
}

func BenchmarkD07P2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2()
	}
}
