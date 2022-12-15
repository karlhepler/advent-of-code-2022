package main

import (
	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
)

func main() {
	var start, end [2]int  // [row, col]
	var heightmap [][]byte // [row][col]

	filepath := "cmd/day12/example"
	it.Must(fyl.ReadEachLine(filepath, func(line []byte) error {
		if col, ok := indexOf(line, 'S'); ok {
			row := len(heightmap)
			start = [2]int{row, col}
			line[col] = 'a'
		}
		if col, ok := indexOf(line, 'E'); ok {
			row := len(heightmap)
			end = [2]int{row, col}
			line[col] = 'z'
		}

		heightmap = append(heightmap, line)

		return nil
	}))
}

func indexOf(line []byte, a byte) (i int, ok bool) {
	for i, c := range line {
		if c == a {
			return i, true
		}
	}

	return -1, false
}
