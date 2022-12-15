package main

import (
	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
)

func main() {
	var start, end [2]int    // [row, col]
	var heightmap [][]string // [row][col]

	filepath := "cmd/day12/example"
	it.Must(fyl.ReadEachLine(filepath, func(line []byte) error {
		if line[0] == 'S' {
			line[0] = 'a'
		}

		return nil
	}))
}
