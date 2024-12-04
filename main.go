package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Specify a puzzle")
		os.Exit(1)
	}

	puzzle := os.Args[1]

	if puzzle == "01-1" {
		Day1Part1()
	} else if puzzle == "01-2" {
		Day1Part2()
	} else if puzzle == "02-1" {
		Day2Part1()
	} else if puzzle == "02-2" {
		Day2Part2()
	} else if puzzle == "03-1" {
		Day3Part1()
	} else if puzzle == "03-2" {
		Day3Part2(true)
	} else if puzzle == "04-1" {
		Day4Part1(true)
	} else if puzzle == "04-2" {
		Day4Part2(true)
	} else {
		panic("not implemented")
	}
}
