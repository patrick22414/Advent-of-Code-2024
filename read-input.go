package main

import (
	"bufio"
	"os"
)

func ReadInput(filename string) (lines chan string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	lines = make(chan string)
	go func() {
		defer f.Close()
		defer close(lines)

		for s.Scan() {
			lines <- s.Text()
		}
	}()
	return
}
