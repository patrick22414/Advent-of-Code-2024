package main

import "testing"

func BenchmarkDay6Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Day6Part2(false)
	}
}

func BenchmarkDay7Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Day7Part1(false)
	}
}

func BenchmarkDay7Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Day7Part2(false)
	}
}

func BenchmarkDay8Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Day8Part1(false)
	}
}

func BenchmarkDay8Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Day8Part2(false)
	}
}

func BenchmarkDay9Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Day9Part1(false)
	}
}

func BenchmarkDay9Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Day9Part2(false)
	}
}
