package main

import (
	"fmt"

	"github.com/karlhepler/advent-of-code-2022/elf"
	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
)

func main() {
	var total int

	filepath := "cmd/day3/input"
	it.Must(fyl.ReadEachLine(filepath, func(line []byte) error {
		left, right := split(line)

		rvalmap := make(map[byte]byte, len(right))
		for i, n := 0, len(right); i < n; i++ {
			rvalmap[right[i]] = right[i]
		}

		for i, n := 0, len(left); i < n; i++ {
			if item, ok := rvalmap[left[i]]; ok {
				total += elf.ItemPriority(item)
				return nil
			}
		}

		return nil
	}))

	fmt.Println(total)
}

func split(line []byte) ([]byte, []byte) {
	halfidx := len(line) / 2
	return line[:halfidx], line[halfidx:]
}
