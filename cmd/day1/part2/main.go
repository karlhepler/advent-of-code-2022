package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/hlp"
	"github.com/karlhepler/advent-of-code-2022/mth"
)

func main() {
	var numCalories int
	var allElvesCalories []int

	// assume running from repo root
	filepath := "cmd/day1/input.txt"
	hlp.Must(fyl.ReadEachLine(filepath, func(line string) error {
		if line == "" {
			allElvesCalories = append(allElvesCalories, numCalories)
			numCalories = 0
			return nil
		}

		numCalories += hlp.Must2(strconv.Atoi(line))
		return nil
	}))

	sort.Ints(allElvesCalories)
	fmt.Println(mth.SumInts(allElvesCalories[len(allElvesCalories)-3:]))
}
