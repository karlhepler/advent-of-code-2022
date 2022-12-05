package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
)

type node struct {
	next *node
	data byte
}

func main() {
	var stacks []*node
	var firstLine = true
	var buildPhase = true

	filepath := "cmd/day5/input"
	it.Must(fyl.ReadEachLine(filepath, func(line []byte) error {
		if !firstLine && len(line) == 0 {
			buildPhase = false
			return nil
		}

		line = append(line, 32) // add extra space to end for parsing

		if firstLine {
			stacks = make([]*node, len(line)/4)
			firstLine = false
		}

		if buildPhase {
			build(stacks, line)
			return nil
		}

		return move(stacks, line)
	}))

	// print answer
	for _, crate := range stacks {
		fmt.Print(string(crate.data))
	}
	fmt.Print("\n")
}

func build(stacks []*node, line []byte) {
	for i, n := 0, len(line)-1; i < n; i += 4 {
		data := line[i+1]
		if data == 32 { // skip if data is a space
			continue
		}

		curr := &node{data: data}
		si := (i / 4)
		top := stacks[si]
		if top == nil {
			stacks[si] = curr
			continue
		}

		bot := bottom(top)
		bot.next = curr
	}
}

func bottom(crate *node) *node {
	for crate != nil && crate.next != nil {
		crate = crate.next
	}
	return crate
}

func move(stacks []*node, line []byte) error {
	numToMove, fromIdx, toIdx, err := parseMove(line)
	if err != nil {
		return err
	}

	for i := 0; i < numToMove; i++ {
		from, to := stacks[fromIdx], stacks[toIdx]
		stacks[fromIdx] = from.next
		from.next = to
		stacks[toIdx] = from
	}

	return nil
}

func parseMove(line []byte) (numToMove, fromIdx, toIdx int, err error) {
	words := strings.Fields(string(line))

	numToMove, err = strconv.Atoi(words[1])
	if err != nil {
		return
	}

	fromIdx, err = strconv.Atoi(words[3])
	if err != nil {
		return
	}
	fromIdx--

	toIdx, err = strconv.Atoi(words[5])
	if err != nil {
		return
	}
	toIdx--

	return
}
