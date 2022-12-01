package mth

func MaxInt(a, b int) int {
	if a > b {
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
