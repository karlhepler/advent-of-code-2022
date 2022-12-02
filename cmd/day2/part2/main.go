package main

import (
	"fmt"

	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
)

func main() {
	// X == lose
	// Y == draw
	// Z == win

	var score int
	var scoremap = map[string]int{
		"A X": 3 + 0, // rock (lose) = scissors
		"B X": 1 + 0, // paper (lose) = rock
		"C X": 2 + 0, // scissors (lose) = paper
		"A Y": 1 + 3, // rock (draw) = rock
		"B Y": 2 + 3, // paper (draw) = paper
		"C Y": 3 + 3, // scissors (draw) = scissors
		"A Z": 2 + 6, // rock (win) = paper
		"B Z": 3 + 6, // paper (win) = scissors
		"C Z": 1 + 6, // scissors (win) = rock
	}

	filepath := "cmd/day2/input.txt"
	it.Must(fyl.ReadEachLine(filepath, func(line string) error {
		score += scoremap[line]
		return nil
	}))

	fmt.Println(score)
}
