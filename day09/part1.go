package day09

import (
	"slices"
	"strconv"

	"github.com/patrick22414/Advent-of-Code-2024/readinput"
)

type DiskFile struct {
	id   int // id of the file, -1 for free space
	size int // number of blocks
}

type Disk []DiskFile

func (d Disk) String() string {
	s := ""
	for _, f := range d {
		for range f.size {
			if f.id == -1 {
				s += "."
			} else {
				s += strconv.Itoa(f.id)
			}
		}
	}
	return s
}

func (d Disk) ActualSize() (s int) {
	for _, f := range d {
		if f.id != -1 {
			s += f.size
		}
	}
	return
}

func Part1() int {
	disk := make(Disk, 0)
	for line := range readinput.ReadInput("./input.txt") {
		for i := 0; i < len(line); i += 2 {
			disk = append(disk, DiskFile{id: i / 2, size: int(line[i] - '0')})

			if i+1 < len(line) {
				disk = append(disk, DiskFile{id: -1, size: int(line[i+1] - '0')})
			}
		}
	}

	back := make(chan int)
	go func() {
		defer close(back)

		diskR := slices.Clone(disk)
		slices.Reverse(diskR)

		for _, f := range diskR {
			if f.id == -1 {
				continue
			}
			for range f.size {
				back <- f.id
			}
		}
	}()

	total := 0
	i, max := 0, disk.ActualSize()
outer:
	for _, f := range disk {
		for range f.size {
			if f.id == -1 {
				total += i * (<-back)
				// fmt.Printf("%v * %v = %v\n", i, n, i*n)
			} else {
				total += i * f.id
				// fmt.Printf("%v * %v = %v\n", i, f.id, i*f.id)
			}
			i++
			if i >= max {
				break outer
			}
		}
	}

	return total
}
