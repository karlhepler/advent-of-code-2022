package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
)

const Width = 40

func main() {
	var vals []int

	filepath := "cmd/day10/input"
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

	var xreg, crty int = 1, 0
	for crtx, val := range vals {
		nextCrty := int(crtx / Width)
		if nextCrty > crty {
			crty = nextCrty
			fmt.Print("\n")
		}

		fmt.Print(pixel(crtx%Width, xreg))

		xreg += val
	}

	fmt.Print("\n")
}

func pixel(crtx, spritex int) string {
	if crtx == spritex-1 || crtx == spritex || crtx == spritex+1 {
		return "#"
	}
	return "."
}
