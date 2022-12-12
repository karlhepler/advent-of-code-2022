package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
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

			visits[bridge.Knots[9]] = true
		}

		return nil
	}))

	fmt.Println(len(visits))
}

type Bridge struct {
	Knots [10][2]float64
}

func (b *Bridge) Step(direction string) error {
	switch direction {
	case "U":
		b.Knots[0][1]--
	case "D":
		b.Knots[0][1]++
	case "L":
		b.Knots[0][0]--
	case "R":
		b.Knots[0][0]++
	default:
		return errors.New("invalid direction: " + direction)
	}

	for i := 0; i < 9; i++ {
		next, err := mov(b.Knots[i], b.Knots[i+1])
		if err != nil {
			break
		}

		b.Knots[i+1] = next
	}

	return nil
}

func mov(head [2]float64, tail [2]float64) (next [2]float64, err error) {
	dx := head[0] - tail[0]
	dy := head[1] - tail[1]

	delta := math.Sqrt(dx*dx + dy*dy)
	unit := [2]float64{dx / delta, dy / delta}

	if delta < 2 {
		err = errors.New("delta less than 2")
		return
	}

	next[0] = tail[0] + sign(unit[0])
	next[1] = tail[1] + sign(unit[1])

	return
}

func sign(a float64) float64 {
	if a > 0 {
		return 1
	}
	if a < 0 {
		return -1
	}
	return 0
}
