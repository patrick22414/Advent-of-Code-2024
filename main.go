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
	}
}
