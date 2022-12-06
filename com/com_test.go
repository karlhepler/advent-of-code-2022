package com_test

import (
	"testing"

	"github.com/karlhepler/advent-of-code-2022/com"
	"github.com/stretchr/testify/assert"
)

func TestGetStartIndex(t *testing.T) {
	tcs := []struct {
		data     []byte
		buffer   int
		expected int
	}{
		// day6 part1
		{[]byte("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), 4, 7},
		{[]byte("bvwbjplbgvbhsrlpgdmjqwftvncz"), 4, 5},
		{[]byte("nppdvjthqldpwncqszvftbrmjlhg"), 4, 6},
		{[]byte("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), 4, 10},
		{[]byte("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), 4, 11},
		// day6 part2
		{[]byte("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), 14, 19},
		{[]byte("bvwbjplbgvbhsrlpgdmjqwftvncz"), 14, 23},
		{[]byte("nppdvjthqldpwncqszvftbrmjlhg"), 14, 23},
		{[]byte("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), 14, 29},
		{[]byte("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), 14, 26},
	}

	for _, tc := range tcs {
		idx := com.GetStartIndex(tc.data, tc.buffer)
		assert.Equal(t, tc.expected, idx)
	}
}
