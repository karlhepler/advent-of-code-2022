package main

import (
	"fmt"

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
			if b, ok := rvalmap[left[i]]; ok {
				total += priority(b)
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

func priority(b byte) int {
	// a - z
	if b >= 97 && b <= 122 {
		return int(b) - 97 + 1
	}

	// A - Z
	if b >= 65 && b <= 90 {
		return int(b) - 65 + 27
	}

	return 0
}
