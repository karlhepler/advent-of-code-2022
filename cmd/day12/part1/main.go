package main

import (
	"fmt"

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

	fmt.Println(fewestStepsToEnd(heightmap, start, end))
}

func indexOf(line []byte, a byte) (i int, ok bool) {
	for i, c := range line {
		if c == a {
			return i, true
		}
	}

	return -1, false
}

func fewestStepsToEnd(heightmap [][]byte, start, end [2]int) int {
	for start != end {
		// Draw line between two points.
		// The line is the hypotenuse of a triangle.
		// Assuming right triangle, find the angle.
		// If the angle from start point is less than 45, then go down (or up depending on quadrant).
		// If the angle from start point is more than 45, then go left (or right depending on quadrant).

		// If the angle is 45, then pick randomly. This choice, however, will
		// influence everything after. So you might need to try the other way too.
		// You will need to do this for each choice.
	}

	return 0
}
