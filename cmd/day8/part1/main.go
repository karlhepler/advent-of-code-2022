package main

import (
	"fmt"

	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
)

func main() {
	var trees = make([][]byte, 0)

	filename := "cmd/day8/input"
	it.Must(fyl.ReadEachLine(filename, func(line []byte) error {
		trees = append(trees, line)
		return nil
	}))

	var numVisible int

	for row, numrows := 0, len(trees); row < numrows; row++ {
		for col, numcols := 0, len(trees[row]); col < numcols; col++ {
			var height = trees[row][col]

			// Assume the tree is visible and look left. If there are any trees of
			// equal or greater size, then this tree is not visible in that
			// direction.
			isVisible := true
			for c := col; c >= 0; c-- {
				if c != col && trees[row][c] >= height {
					isVisible = false
					break
				}
			}

			// If the tree is still visible after that test, then tally it and go to
			// the next tree.
			if isVisible {
				numVisible++
				continue
			}

			// Assume the tree is visible and look right. If there are any trees of
			// equal or greater size, then this tree is not visible in that
			// direction.
			isVisible = true
			for c := col; c < numcols; c++ {
				if c != col && trees[row][c] >= height {
					isVisible = false
					break
				}
			}

			// If the tree is still visible after that test, then tally it and go to
			// the next tree.
			if isVisible {
				numVisible++
				continue
			}

			// Assume the tree is visible and look up. If there are any trees of
			// equal or greater size, then this tree is not visible in that
			// direction.
			isVisible = true
			for r := row; r >= 0; r-- {
				if r != row && trees[r][col] >= height {
					isVisible = false
					break
				}
			}

			// If the tree is still visible after that test, then tally it and go to
			// the next tree.
			if isVisible {
				numVisible++
				continue
			}

			// Assume the tree is visible and look down. If there are any trees of
			// equal or greater size, then this tree is not visible in that
			// direction.
			isVisible = true
			for r := row; r < numrows; r++ {
				if r != row && trees[r][col] >= height {
					isVisible = false
					break
				}
			}

			// If the tree is still visible after that test, then tally it and go to
			// the next tree.
			if isVisible {
				numVisible++
				continue
			}
		}
	}

	fmt.Println(numVisible) // 1625 is WRONG
}
