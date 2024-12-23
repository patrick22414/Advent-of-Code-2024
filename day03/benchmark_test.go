package day03

import (
	"testing"
)

func BenchmarkD03P1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1()
	}
	b.Log(Part1())
}

func BenchmarkD03P2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2()
	}
	b.Log(Part2())
}
