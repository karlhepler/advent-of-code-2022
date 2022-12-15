package mth

import (
	"math"
	"strconv"

	"github.com/karlhepler/advent-of-code-2022/it"
)

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func SumInts(ints []int) (sum int) {
	for _, num := range ints {
		sum += num
	}

	return sum
}

func Sign(a float64) float64 {
	if a > 0 {
		return 1
	}
	if a < 0 {
		return -1
	}
	return 0
}

func SumDigits(a int) (sum int) {
	str := strconv.Itoa(a)

	multi := 1
	for _, char := range str {
		if char == '-' {
			multi = -1
			continue
		}

		digit := it.Must2(strconv.Atoi(string(char)))
		sum += digit
	}

	return sum * multi
}

func LastDigits(a, d int) int {
	str := strconv.Itoa(a)
	if d >= len(str) {
		return a
	}

	return it.Must2(strconv.Atoi(str[len(str)-d:]))
}

func FirstDigits(a, d int) int {
	str := strconv.Itoa(a)
	if d >= len(str) {
		return a
	}

	return it.Must2(strconv.Atoi(str[:d]))
}

// FewestDigits returns the int with the fewest digits
func FewestDigits(a, b int) int {
	au := int(math.Abs(float64(a)))
	bu := int(math.Abs(float64(b)))

	if len(strconv.Itoa(au)) < len(strconv.Itoa(bu)) {
		return a
	}

	return b
}

func CongruentByModulo(a, b, mod int) bool {
	return (a-b)%mod == 0
}
