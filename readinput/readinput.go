package readinput

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
		if err := s.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "scanner error", err)
		}
	}()
	return
}

func ReadBytes(filename string) (lines chan []byte) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	lines = make(chan []byte)
	go func() {
		defer f.Close()
		defer close(lines)

		for s.Scan() {
			lines <- slices.Clone(s.Bytes()) // it would cause a race condition without Clone
		}
		if err := s.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "scanner error", err)
		}
	}()
	return
}
