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

	if puzzle == "05-1" {
		Day5Part1(true)
	} else if puzzle == "05-2" {
		Day5Part2(true)
	} else if puzzle == "06-1" {
		Day6Part1(true)
	} else if puzzle == "06-2" {
		Day6Part2(true)
	} else if puzzle == "07-1" {
		Day7Part1(true)
	} else if puzzle == "07-2" {
		Day7Part2(true)
	} else if puzzle == "08-1" {
		Day8Part1(true)
	} else if puzzle == "08-2" {
		Day8Part2(true)
	} else if puzzle == "09-1" {
		Day9Part1(true)
	} else if puzzle == "09-2" {
		Day9Part2(true)
	} else {
		panic("not implemented")
	}
}
