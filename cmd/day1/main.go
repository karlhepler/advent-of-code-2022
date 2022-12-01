package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/karlhepler/advent-of-code-2022/hlp"
	"github.com/karlhepler/advent-of-code-2022/mth"
)

func main() {
	// assume running from repo root
	file := hlp.Must(os.Open("cmd/day1/input.txt"))
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var numCalories = 0
	var mostCalories = 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			mostCalories = mth.MaxInt(mostCalories, numCalories)
			numCalories = 0
			continue
		}

		numCalories += hlp.Must(strconv.Atoi(line))
	}

	fmt.Println(mostCalories)
}
