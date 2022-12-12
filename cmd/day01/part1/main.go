package main

import (
	"fmt"
	"strconv"

	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
	"github.com/karlhepler/advent-of-code-2022/mth"
)

func main() {
	var numCalories = 0
	var mostCalories = 0

	// assume running from repo root
	filepath := "cmd/day1/input"
	it.Must(fyl.ReadEachLine(filepath, func(line []byte) error {
		if len(line) == 0 {
			mostCalories = mth.MaxInt(mostCalories, numCalories)
			numCalories = 0
			return nil
		}

		numCalories += it.Must2(strconv.Atoi(string(line)))
		return nil
	}))

	fmt.Println(mostCalories)
}
