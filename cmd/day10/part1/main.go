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

	filepath := "cmd/day10/example1"
	it.Must(fyl.ReadEachLine(filepath, func(line []byte) error {
		cmds = append(cmds, addxByLine(line)...)
		return nil
	}))

	fmt.Println(cmds)
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
