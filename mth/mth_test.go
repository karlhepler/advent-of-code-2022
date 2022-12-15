package mth_test

import (
	"testing"

	"github.com/karlhepler/advent-of-code-2022/mth"
	"github.com/stretchr/testify/assert"
)

func TestSumDigits(t *testing.T) {
	tcs := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{5, 5},
		{10, 1},
		{50, 5},
		{11, 2},
		{55, 10},
		{12345, 15},
	}

	for _, tc := range tcs {
		output := mth.SumDigits(tc.input)
		assert.Equal(t, tc.expected, output)
	}
}

func TestLastDigits(t *testing.T) {
	tcs := []struct {
		a        int
		d        int
		expected int
	}{
		{123, 2, 23},
		{567, 5, 567},
		{123456, 1, 6},
		{123456, 2, 56},
		{123456, 5, 23456},
	}

	for _, tc := range tcs {
		output := mth.LastDigits(tc.a, tc.d)
		assert.Equal(t, tc.expected, output)
	}
}

func TestFirstDigits(t *testing.T) {
	tcs := []struct {
		a        int
		d        int
		expected int
	}{
		{123, 2, 12},
		{567, 5, 567},
		{123456, 1, 1},
		{123456, 2, 12},
		{123456, 5, 12345},
	}

	for _, tc := range tcs {
		output := mth.FirstDigits(tc.a, tc.d)
		assert.Equal(t, tc.expected, output)
	}
}
