package main

import (
	"fmt"

	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
)

func main() {
	var score int
	var scoremap = map[string]int{
		"A X": 1 + 3, // rock rock (draw)
		"B X": 1 + 0, // paper rock (lose)
		"C X": 1 + 6, // scissors rock (win)
		"A Y": 2 + 6, // rock paper (win)
		"B Y": 2 + 3, // paper paper (draw)
		"C Y": 2 + 0, // scissors paper (lose)
		"A Z": 3 + 0, // rock scissors (lose)
		"B Z": 3 + 6, // paper scissors (win)
		"C Z": 3 + 3, // scissors scissors (draw)
	}

	filepath := "cmd/day2/input.txt"
	it.Must(fyl.ReadEachLine(filepath, func(line string) error {
		score += scoremap[line]
		return nil
	}))

	fmt.Println(score)
}
