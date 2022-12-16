package util

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

func IntAbs(a int) int {
	return int(math.Abs(float64(a)))
}

func IsType(v any, expectedType string) bool {
	return fmt.Sprintf("%T", v) == expectedType
}
