package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
)

func main() {
	var total int

	filepath := "cmd/day4/input"
	it.Must(fyl.ReadEachLine(filepath, func(line []byte) error {
		fields := strings.FieldsFunc(string(line), func(f rune) bool {
			return f == '-' || f == ','
		})

		ids := it.Must2(atoiSlice(fields))
		if isWithin(ids[:2], ids[2:]) || isWithin(ids[2:], ids[:2]) {
			total++
		}

		return nil
	}))

	fmt.Println(total)
}

func atoiSlice(a []string) ([]int, error) {
	n := len(a)
	out := make([]int, n)
	var err error

	for i := 0; i < n; i++ {
		out[i], err = strconv.Atoi(a[i])
		if err != nil {
			return nil, err
		}
	}

	return out, nil
}

// isWithin determines if a is within b
func isWithin(a []int, b []int) bool {
	return a[0] >= b[0] && a[1] <= b[1]
}
