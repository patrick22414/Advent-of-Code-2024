package day10

import (
	"testing"
)

func BenchmarkD10P1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1()
	}
	b.Log(Part1())
}

func BenchmarkD10P2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2()
	}
	b.Log(Part2())
}
