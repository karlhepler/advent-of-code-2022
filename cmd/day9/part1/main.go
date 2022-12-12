package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
)

func main() {
	var bridge = &Bridge{}

	filepath := "cmd/day9/example"
	it.Must(fyl.ReadEachLine(filepath, func(line []byte) error {
		split := strings.Split(string(line), " ")
		numSteps := it.Must2(strconv.Atoi(split[1]))

		for step := 0; step < numSteps; step++ {
			err := bridge.Step(split[0])
			if err != nil {
				return err
			}
		}

		fmt.Printf("%+v\n", bridge)

		return nil
	}))
}

type Bridge struct {
	Head Vector
	Tail Vector
}

func (b *Bridge) Step(direction string) error {
	switch direction {
	case "U":
		b.Head[1]--
	case "D":
		b.Head[1]++
	case "L":
		b.Head[0]--
	case "R":
		b.Head[0]++
	default:
		return errors.New("invalid direction: " + direction)
	}

	return nil
}

type Vector [2]float64
