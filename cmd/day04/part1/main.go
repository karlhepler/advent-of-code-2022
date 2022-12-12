package main

import (
	"fmt"
	"strings"

	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
	"github.com/karlhepler/advent-of-code-2022/rng"
	"github.com/karlhepler/advent-of-code-2022/str"
)

func main() {
	var total int

	filepath := "cmd/day4/input"
	it.Must(fyl.ReadEachLine(filepath, func(line []byte) error {
		fields := strings.FieldsFunc(string(line), func(f rune) bool {
			return f == '-' || f == ','
		})

		ids := it.Must2(str.AtoiSlice(fields))
		if rng.IsWithin(ids[:2], ids[2:]) || rng.IsWithin(ids[2:], ids[:2]) {
			total++
		}

		return nil
	}))

	fmt.Println(total)
}
