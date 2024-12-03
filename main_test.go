package main

import "testing"

func BenchmarkDay3Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Day3Part2(false)
	}
}
