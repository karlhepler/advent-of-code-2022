package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
	"github.com/karlhepler/advent-of-code-2022/mth"
)

func main() {
	var numCalories int
	var allElvesCalories []int

	// assume running from repo root
	filepath := "cmd/day1/input"
	it.Must(fyl.ReadEachLine(filepath, func(line []byte) error {
		if len(line) == 0 {
			allElvesCalories = append(allElvesCalories, numCalories)
			numCalories = 0
			return nil
		}

		numCalories += it.Must2(strconv.Atoi(string(line)))
		return nil
	}))

	sort.Ints(allElvesCalories)
	fmt.Println(mth.SumInts(allElvesCalories[len(allElvesCalories)-3:]))
}
