package main

import (
	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
)

func main() {
	filepath := "cmd/day10/example1"
	it.Must(fyl.ReadEachLine(filepath, func(line []byte) error {
		return nil
	}))
}
