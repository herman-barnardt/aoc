package util

import "math"

type Point struct {
	X int
	Y int
}

func IntAbs(a int) int {
	return int(math.Abs(float64(a)))
}
