package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStartIndex(t *testing.T) {
	tcs := []struct {
		data     []byte
		expected int
	}{
		{[]byte("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), 7},
		{[]byte("bvwbjplbgvbhsrlpgdmjqwftvncz"), 5},
		{[]byte("nppdvjthqldpwncqszvftbrmjlhg"), 6},
		{[]byte("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), 10},
		{[]byte("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), 11},
	}

	for _, tc := range tcs {
		idx := GetStartIndex(tc.data)
		assert.Equal(t, tc.expected, idx)
	}
}
