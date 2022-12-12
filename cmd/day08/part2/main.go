package main

import (
	"fmt"

	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
)

func main() {
	var trees = make([][]int, 0)

	filename := "cmd/day8/input"
	it.Must(fyl.ReadEachLine(filename, func(line []byte) error {
		treeline := make([]int, len(line))
		for i, n := 0, len(line); i < n; i++ {
			treeline[i] = int(line[i])
		}
		trees = append(trees, treeline)
		return nil
	}))

	var maxScore int

	for row, numrows := 0, len(trees); row < numrows; row++ {
		for col, numcols := 0, len(trees[row]); col < numcols; col++ {
			var height = trees[row][col]
			var dist [4]int

			for c := col; c >= 0; c-- {
				if c == col {
					continue
				}

				dist[0]++
				if trees[row][c] >= height {
					break
				}
			}

			for c := col; c < numcols; c++ {
				if c == col {
					continue
				}

				dist[1]++
				if trees[row][c] >= height {
					break
				}
			}

			for r := row; r >= 0; r-- {
				if r == row {
					continue
				}

				dist[2]++
				if trees[r][col] >= height {
					break
				}
			}

			for r := row; r < numrows; r++ {
				if r == row {
					continue
				}

				dist[3]++
				if trees[r][col] >= height {
					break
				}
			}

			score := dist[0] * dist[1] * dist[2] * dist[3]
			if maxScore < score {
				maxScore = score
			}
		}
	}

	fmt.Println(maxScore)
}
