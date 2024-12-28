package day04

import (
	"testing"
)

func BenchmarkD04P1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1()
	}
}

func BenchmarkD04P2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2()
	}
}
