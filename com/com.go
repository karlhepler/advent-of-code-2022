package com

import "container/list"

func GetStartIndex(line []byte, buffer int) int {
	var fifo = list.New()

	for i, char := range line {
		if fifo.Len() == buffer && isUniqueList(fifo) {
			return i
		}

		if fifo.Len() == buffer {
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
