package main

import (
	"container/list"
	"fmt"

	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
)

func main() {
	var startIndex int

	filepath := "cmd/day6/input"
	it.Must(fyl.ReadEachLine(filepath, func(line []byte) error {
		startIndex = GetStartIndex(line)
		return nil
	}))

	fmt.Println(startIndex)
}

func GetStartIndex(line []byte) int {
	var fifo = list.New()

	for i, char := range line {
		if fifo.Len() == 4 && isUniqueList(fifo) {
			return i
		}

		if fifo.Len() == 4 {
			fifo.Remove(fifo.Front())
		}

		fifo.PushBack(char)
	}

	return 0
}

func isUniqueList(list *list.List) bool {
	cache := make(map[any]bool, list.Len())

	for elem := list.Front(); elem != nil; elem = elem.Next() {
		if _, ok := cache[elem.Value]; !ok {
			cache[elem.Value] = true
			continue
		}

		return false
	}

	return true
}
