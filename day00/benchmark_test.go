package day00

import (
	"testing"
)

func BenchmarkD00P1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part1()
	}
}

func BenchmarkD00P2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part2()
	}
}
