package util

import "math"

type Point struct {
	x, y int
}

func IntAbs(a, b int) int {
	return int(math.Abs(float64(a - b)))
}
