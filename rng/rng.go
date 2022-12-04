package rng

// IsWithin determines if a is within b.
func IsWithin(a []int, b []int) bool {
	return a[0] >= b[0] && a[1] <= b[1]
}

// Overlaps determines if two ranges overlap.
// https://stackoverflow.com/a/3269471/4987538
func Overlaps(a []int, b []int) bool {
	return a[0] <= b[1] && a[1] >= b[0]
}
