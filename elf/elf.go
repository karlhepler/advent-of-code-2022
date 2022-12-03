package elf

func ItemPriority(item byte) int {
	// a - z
	if item >= 97 && item <= 122 {
		return int(item) - 97 + 1
	}

	// A - Z
	if item >= 65 && item <= 90 {
		return int(item) - 65 + 27
	}

	return 0
}
