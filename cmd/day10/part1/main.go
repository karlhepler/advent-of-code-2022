package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
	"github.com/karlhepler/advent-of-code-2022/mth"
)

func main() {
	var cmds []int

	filepath := "cmd/day10/example2"
	it.Must(fyl.ReadEachLine(filepath, func(line []byte) error {
		cmds = append(cmds, addxByLine(line)...)
		return nil
	}))

	x := 1
	for i, n := 1, len(cmds); i < n; i++ {
		x += cmds[i]
		if i == 20 || i == 60 || i == 100 || i == 140 || i == 180 || i == 220 {
			fmt.Println(i, x, i*x)
		}
	}
}

func addxByLine(line []byte) (cmds []int) {
	addx := strings.Split(string(line), " ")

	if addx[0] == "noop" {
		return append(cmds, 0)
	}

	num := it.Must2(strconv.ParseFloat(addx[1], 64))
	sign := mth.Sign(num)

	for i, n := 0, int(math.Abs(num)); i < n; i++ {
		cmds = append(cmds, int(sign))
	}

	return cmds
}
