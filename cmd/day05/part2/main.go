package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
)

type crate struct {
	next    *crate
	content byte
}

func main() {
	var crates []*crate
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
			crates = make([]*crate, len(line)/4)
			firstLine = false
		}

		if buildPhase {
			build(crates, line)
			return nil
		}

		return move(crates, line)
	}))

	// print answer
	for _, crate := range crates {
		fmt.Print(string(crate.content))
	}
	fmt.Print("\n")
}

func build(crates []*crate, line []byte) {
	for i, n := 0, len(line)-1; i < n; i += 4 {
		content := line[i+1]
		if content == 32 { // skip if content is a space
			continue
		}

		curr := &crate{content: content}
		si := (i / 4)
		top := crates[si]
		if top == nil {
			crates[si] = curr
			continue
		}

		bot := bottom(top)
		bot.next = curr
	}
}

func bottom(crate *crate) *crate {
	for crate != nil && crate.next != nil {
		crate = crate.next
	}
	return crate
}

func move(crates []*crate, line []byte) error {
	numSteps, fromIdx, toIdx, err := parseMove(line)
	if err != nil {
		return err
	}

	from, to := crates[fromIdx], crates[toIdx]
	lastFrom := crateAtStep(from, numSteps)

	crates[fromIdx] = lastFrom.next
	crates[toIdx] = from
	lastFrom.next = to

	return nil
}

func crateAtStep(crate *crate, numSteps int) *crate {
	for step := 1; crate != nil && crate.next != nil && step < numSteps; step++ {
		crate = crate.next
	}
	return crate
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
