package main

import (
	"fmt"

	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
)

func main() {
	var grid = make([][]byte, 0)

	filepath := "cmd/day8/input"
	it.Must(fyl.ReadEachLine(filepath, func(line []byte) error {
		grid = append(grid, line)
		return nil
	}))

	var masks = make([][][]byte, 4)
	masks[0] = lrMask(grid)
	masks[1] = rlMask(grid)
	masks[2] = tbMask(grid)
	masks[3] = btMask(grid)

	fmt.Println(countVisible(orMatrices(masks)))
}

func lrMask(grid [][]byte) [][]byte {
	var mask = make([][]byte, len(grid))

	for row, maxrow := 0, len(grid); row < maxrow; row++ {

		max := byte(0)
		maxcol := len(grid[row])
		mask[row] = make([]byte, maxcol)

		for col := 0; col < maxcol; col++ {
			val := grid[row][col]
			if val > max {
				mask[row][col] = 1
				max = val
			}
		}
	}

	return mask
}

func rlMask(grid [][]byte) [][]byte {
	var mask = make([][]byte, len(grid))

	for row, maxrow := 0, len(grid); row < maxrow; row++ {

		max := byte(0)
		maxcol := len(grid[row])
		mask[row] = make([]byte, maxcol)

		for col := maxcol - 1; col >= 0; col-- {
			val := grid[row][col]
			if val > max {
				mask[row][col] = 1
				max = val
			}
		}
	}

	return mask
}

func tbMask(grid [][]byte) [][]byte {
	var mask = make([][]byte, len(grid))

	maxrow := len(grid)
	maxcol := len(grid[0])

	for col := 0; col < maxcol; col++ {

		max := byte(0)

		for row := 0; row < maxrow; row++ {
			if len(mask[row]) == 0 {
				mask[row] = make([]byte, maxcol)
			}

			val := grid[row][col]
			if val > max {
				mask[row][col] = 1
				max = val
			}
		}
	}

	return mask
}

func btMask(grid [][]byte) [][]byte {
	var mask = make([][]byte, len(grid))

	maxrow := len(grid)
	maxcol := len(grid[0])

	for col := maxcol - 1; col >= 0; col-- {

		max := byte(0)

		for row := maxrow - 1; row >= 0; row-- {
			if len(mask[row]) == 0 {
				mask[row] = make([]byte, maxcol)
			}

			val := grid[row][col]
			if val > max {
				mask[row][col] = 1
				max = val
			}
		}
	}

	return mask
}

func orMatrices(mats [][][]byte) [][]byte {
	var out = make([][]byte, len(mats[0]))

	maxrow := len(mats[0])
	maxcol := len(mats[0][0])

	for _, mat := range mats {
		for row := 0; row < maxrow; row++ {

			if len(out[row]) == 0 {
				out[row] = make([]byte, maxrow)
			}

			for col := 0; col < maxcol; col++ {
				out[row][col] = mat[row][col] | out[row][col]
			}
		}
	}

	return out
}

func countVisible(mat [][]byte) int {
	var total int

	for row, maxrow := 0, len(mat[0]); row < maxrow; row++ {
		for col, maxcol := 0, len(mat); col < maxcol; col++ {
			if mat[row][col] == 1 {
				total++
			}
		}
	}

	return total
}
