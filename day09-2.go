package main

import (
	"container/list"
	"fmt"
	"strconv"
)

type ListDisk struct{ ls *list.List }

func (d ListDisk) Push(f DiskFile) {
	d.ls.PushBack(f)
}

const (
	NotAllocated int = iota - 1
	AllocatedByReplace
	AllocatedByInsert
)

// Find a free space for `f` before the n-th element and insert into it
//
// Returns:
//
// - 1: if successful and list length grew by 1
//
// - 0: if successful but list length did not grow
//
// - -1: if unsuccessful
func (d ListDisk) AllocBefore(f DiskFile, before int) int {
	e, i := d.ls.Front(), 0
	for i < before && e != nil {
		v := e.Value.(DiskFile)
		if v.id != -1 || v.size < f.size {
			i++
			e = e.Next()
			continue
		}

		d.ls.InsertBefore(f, e)
		v.size -= f.size
		if v.size == 0 {
			d.ls.Remove(e)
			return AllocatedByReplace
		}
		e.Value = v
		return AllocatedByInsert
	}
	return NotAllocated
}

func (d ListDisk) String() string {
	s := ""
	e := d.ls.Front()
	for e != nil {
		f := e.Value.(DiskFile)
		for range f.size {
			if f.id == -1 {
				s += "."
			} else {
				s += strconv.Itoa(f.id)
			}
		}
		e = e.Next()
	}
	return s
}

func Day9Part2(output bool) {
	disk := ListDisk{list.New()}
	for line := range ReadInput("input/09.txt") {
		for i := 0; i < len(line); i += 2 {
			if line[i] != '0' {
				disk.Push(DiskFile{id: i / 2, size: int(line[i] - '0')})
			}
			if i+1 < len(line) && line[i+1] != '0' {
				disk.Push(DiskFile{id: -1, size: int(line[i+1] - '0')})
			}
		}
	}

	b, j := disk.ls.Back(), disk.ls.Len()-1
	for b != nil {
		v := b.Value.(DiskFile)
		if v.id == -1 {
			j--
			b = b.Prev()
			continue
		}
		// fmt.Println(disk, j, v)

		alloc := disk.AllocBefore(v, j)
		if alloc >= 0 {
			// alloc is successful
			b.Value = DiskFile{id: -1, size: v.size}
		}
		if alloc <= 0 {
			j--
		}
		b = b.Prev()
	}

	i, total := 0, 0
	for e := disk.ls.Front(); e != nil; e = e.Next() {
		v := e.Value.(DiskFile)
		if v.id != -1 {
			for range v.size {
				total += i * v.id
				i++
			}
		} else {
			i += v.size
		}
	}
	if output {
		fmt.Println(total)
	}
}
