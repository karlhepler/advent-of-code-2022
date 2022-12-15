package main

import (
	"container/list"
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/karlhepler/advent-of-code-2022/fyl"
	"github.com/karlhepler/advent-of-code-2022/it"
)

func main() {
	var monkeyLine = regexp.MustCompile(`^Monkey\s(\d+):$`)
	var startingItemsLine = regexp.MustCompile(`^\s+Starting\sitems:\s(.+)$`)
	var operationLine = regexp.MustCompile(`^\s+Operation:\snew\s=\sold\s([+*])\s(.+)$`)
	var testLine = regexp.MustCompile(`^\s+Test:\sdivisible\sby\s(\d+)$`)
	var caseTrueLine = regexp.MustCompile(`^\s+If\strue:\sthrow\sto\smonkey\s(\d+)$`)
	var caseFalseLine = regexp.MustCompile(`^\s+If\sfalse:\sthrow\sto\smonkey\s(\d+)$`)

	var monkey *Monkey
	var monkeys = make([]*Monkey, 0)

	filepath := "cmd/day11/input"
	it.Must(fyl.ReadEachLine(filepath, func(line []byte) error {
		linestr := string(line)

		// Monkey
		if match := monkeyLine.FindStringSubmatch(linestr); len(match) > 1 {
			monkey = &Monkey{ID: len(monkeys)}
			monkeys = append(monkeys, monkey)
			return nil
		}

		// Starting Items
		if match := startingItemsLine.FindStringSubmatch(linestr); len(match) > 1 {
			monkey.Items = list.New()

			items := strings.Split(match[1], ", ")
			for _, item := range items {
				monkey.Items.PushBack(WorryLevel(it.Must2(strconv.Atoi(item))))
			}

			return nil
		}

		// Operation
		if match := operationLine.FindStringSubmatch(linestr); len(match) > 2 {
			var old bool
			num, err := strconv.Atoi(match[2])
			if err != nil {
				old = true
			}

			monkey.Operation = func(wl WorryLevel) (WorryLevel, error) {
				if match[1] == "+" && old {
					return wl + wl, nil
				}
				if match[1] == "+" && !old {
					return wl + WorryLevel(num), nil
				}
				if match[1] == "*" && old {
					return wl * wl, nil
				}
				if match[1] == "*" && !old {
					return wl * WorryLevel(num), nil
				}

				return 0, errors.New("invalid operation: " + match[0])
			}

			return nil
		}

		// Test
		if match := testLine.FindStringSubmatch(linestr); len(match) > 1 {
			num := WorryLevel(it.Must2(strconv.Atoi(match[1])))
			monkey.Test = func(wl WorryLevel) bool {
				return (wl % num) == 0
			}

			return nil
		}

		// If True
		if match := caseTrueLine.FindStringSubmatch(linestr); len(match) > 1 {
			id := MonkeyID(it.Must2(strconv.Atoi(match[1])))
			monkey.ThrowIfTrue = func(t bool) (MonkeyID, error) {
				if t {
					return id, nil
				}

				return 0, errors.New("not true")
			}
			return nil
		}

		// If False
		if match := caseFalseLine.FindStringSubmatch(linestr); len(match) > 1 {
			id := MonkeyID(it.Must2(strconv.Atoi(match[1])))
			monkey.ThrowIfFalse = func(t bool) (MonkeyID, error) {
				if !t {
					return id, nil
				}

				return 0, errors.New("not false")
			}
			return nil
		}

		return nil
	}))

	processInspections(monkeys, 10000)
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].NumInspections > monkeys[j].NumInspections
	})

	monkeyBusiness := monkeys[0].NumInspections * monkeys[1].NumInspections
	fmt.Printf("Monkey Business = %d\n", monkeyBusiness)
}

type MonkeyID int
type WorryLevel int

type Monkey struct {
	ID             int
	Items          *list.List
	Operation      func(WorryLevel) (WorryLevel, error)
	Test           func(WorryLevel) bool
	ThrowIfTrue    func(bool) (MonkeyID, error)
	ThrowIfFalse   func(bool) (MonkeyID, error)
	NumInspections int
}

func processInspections(monkeys []*Monkey, numRounds int) {
	for round := 0; round < numRounds; round++ {
		for _, monkey := range monkeys {
			processInspection(monkey, monkeys)
		}
	}
}

func processInspection(monkey *Monkey, monkeys []*Monkey) {
	for elem := monkey.Items.Front(); elem != nil; {
		next := elem.Next()
		worry := monkey.Items.Remove(elem).(WorryLevel)

		newWorry := it.Must2(monkey.Operation(worry))
		relievedWorry := relief(newWorry)
		testResult := monkey.Test(relievedWorry)

		var id MonkeyID
		var err error
		if id, err = monkey.ThrowIfTrue(testResult); err != nil {
			id = it.Must2(monkey.ThrowIfFalse(testResult))
		}

		monkey.NumInspections++
		monkeys[id].Items.PushBack(relievedWorry)

		elem = next
	}
}

func relief(worry WorryLevel) WorryLevel {
	// https://www.reddit.com/r/adventofcode/comments/zifqmh/comment/j0c0fcz/?utm_source=share&utm_medium=web2x&context=3
	return worry % 9699690
}
