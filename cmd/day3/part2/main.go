package main

import (
	"fmt"

	"github.com/karlhepler/advent-of-code-2022/elf"
	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
)

func main() {
	var total int
	var rucksacks []byte
	var numRucksacks int

	filepath := "cmd/day3/input"
	it.Must(fyl.ReadEachLine(filepath, func(line []byte) error {
		rucksacks = append(rucksacks, unique(line)...)
		numRucksacks++

		if numRucksacks == 3 {
			item := repeated(rucksacks, 3)
			total += elf.ItemPriority(item)

			rucksacks = []byte{}
			numRucksacks = 0
		}

		return nil
	}))

	fmt.Println(total)
}

func unique(input []byte) (output []byte) {
	idx := make(map[byte]byte)

	for _, val := range input {
		if _, ok := idx[val]; !ok {
			output = append(output, val)
		}

		idx[val] = val
	}

	return
}

func repeated(rucksacks []byte, max int) byte {
	idx := make(map[byte]int, len(rucksacks))
	for _, item := range rucksacks {
		idx[item]++
		if idx[item] == max {
			return item
		}
	}

	return 0
}
