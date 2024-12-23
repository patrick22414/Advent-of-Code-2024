package day02

import (
	"testing"
)

func BenchmarkD02P1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1()
	}
	b.Log(Part1())
}

func BenchmarkD02P2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2()
	}
	b.Log(Part2())
}
