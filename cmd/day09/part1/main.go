package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
	"github.com/karlhepler/advent-of-code-2022/mth"
)

func main() {
	var bridge = &Bridge{}
	var visits = make(map[[2]float64]bool)

	filepath := "cmd/day9/input"
	it.Must(fyl.ReadEachLine(filepath, func(line []byte) error {
		split := strings.Split(string(line), " ")
		numSteps := it.Must2(strconv.Atoi(split[1]))

		for step := 0; step < numSteps; step++ {
			err := bridge.Step(split[0])
			if err != nil {
				return err
			}

			visits[bridge.Tail] = true
		}

		return nil
	}))

	fmt.Println(len(visits))
}

type Bridge struct {
	Head [2]float64
	Tail [2]float64
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

	for {
		dx := b.Head[0] - b.Tail[0]
		dy := b.Head[1] - b.Tail[1]

		delta := math.Sqrt(dx*dx + dy*dy)
		unit := [2]float64{dx / delta, dy / delta}

		// fmt.Printf("delta=%v; udx=%v\n", delta, unit)
		if delta < 2 {
			break
		}

		b.Tail[0] += mth.Sign(unit[0])
		b.Tail[1] += mth.Sign(unit[1])
	}

	// fmt.Printf("%+v\n", b)

	// fmt.Println("---")

	return nil
}
