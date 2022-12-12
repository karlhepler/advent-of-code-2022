package main

import (
	"fmt"

	"github.com/karlhepler/advent-of-code-2022/com"
	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
)

func main() {
	var startIndex int

	filepath := "cmd/day6/input"
	it.Must(fyl.ReadEachLine(filepath, func(line []byte) error {
		startIndex = com.GetStartIndex(line, 4)
		return nil
	}))

	fmt.Println(startIndex)
}
