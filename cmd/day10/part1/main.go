package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
)

func main() {
	var vals = []int{1}

	filepath := "cmd/day10/example2"
	it.Must(fyl.ReadEachLine(filepath, func(line []byte) error {
		cmdval := strings.Split(string(line), " ")

		// noop takes one cycle and does nothing
		if cmdval[0] == "noop" {
			vals = append(vals, 0)
			return nil
		}

		// addx takes two cycles and adds the value to x
		vals = append(vals, 0, it.Must2(strconv.Atoi(cmdval[1])))

		return nil
	}))

	var x, sum int
	for i, val := range vals {
		cycle := i + 1
		x += val

		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
			sum += (cycle * x)
		}
	}

	fmt.Println(sum)
}
