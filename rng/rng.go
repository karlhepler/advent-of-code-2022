package rng

import "github.com/karlhepler/advent-of-code-2022/mth"

// IsWithin determines if a is within b.
func IsWithin(a []int, b []int) bool {
	return a[0] >= b[0] && a[1] <= b[1]
}

// Overlaps determines if two ranges overlap.
// https://stackoverflow.com/a/25369187/4987538
func Overlaps(a []int, b []int) bool {
	return mth.MaxInt(a[1], b[1])-mth.MinInt(a[0], b[0]) < (a[1]-a[0])+(b[1]-b[0])
}

// Overlaps determines if a overlaps b.
// https://noonat.github.io/intersect/#aabb-vs-aabb
// func Overlaps(a []int, b []int) bool {
// 	ax, bx := a[0], b[0]
// 	ahw, bhw := float64((a[1]-a[0]))/2.0, float64((b[1]-b[0]))/2.0

// 	dx := math.Abs(float64(ax) - float64(bx))
// 	px := ahw + bhw - dx

// 	return px >= 0
// }
