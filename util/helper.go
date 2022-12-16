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

func IntMin(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func IntMax(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func IsType(v any, expectedType string) bool {
	return fmt.Sprintf("%T", v) == expectedType
}
