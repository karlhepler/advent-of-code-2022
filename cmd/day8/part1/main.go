package main

import (
	"fmt"

	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
)

const (
	LeftRight int = iota
	RightLeft
	TopBottom
	BottomTop
)

func main() {
	var grid = make([][]byte, 0)

	filepath := "cmd/day8/example"
	it.Must(fyl.ReadEachLine(filepath, func(line []byte) error {
		grid = append(grid, line)
		return nil
	}))

	var masks = make([][][]bool, 4)
	masks[0] = lrMask(grid)
	masks[1] = rlMask(grid)
	masks[2] = tbMask(grid)
	masks[3] = btMask(grid)

	fmt.Println(masks[1])
}

func lrMask(grid [][]byte) [][]bool {
	var mask = make([][]bool, len(grid))

	for col, maxcol := 0, len(grid); col < maxcol; col++ {

		max := byte(0)
		maxrow := len(grid[col])
		mask[col] = make([]bool, maxrow)

		for row := 0; row < maxrow; row++ {
			val := grid[col][row]
			if val > max {
				mask[col][row] = true
				max = val
			}
		}
	}

	return mask
}

func rlMask(grid [][]byte) [][]bool {
	var mask = make([][]bool, len(grid))

	for col, maxcol := 0, len(grid); col < maxcol; col++ {

		max := byte(0)
		maxrow := len(grid[col])
		mask[col] = make([]bool, maxrow)

		for row := maxrow - 1; row >= 0; row-- {
			val := grid[col][row]
			if val > max {
				mask[col][row] = true
				max = val
			}
		}
	}

	return mask
}

func tbMask(grid [][]byte) [][]bool {
	var mask = make([][]bool, len(grid))
	return mask
}

// fmt.Printf("max=%s val=%s ", string(max), string(val))
// fmt.Printf("visible=%t\n", mask[y][x])

func btMask(grid [][]byte) [][]bool {
	var mask = make([][]bool, len(grid))
	return mask
}
