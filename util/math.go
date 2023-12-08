package util

import (
	"math"
)

func IntAbs(a int) int {
	return int(math.Abs(float64(a)))
}

func IntMin(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func IntMax(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func GreatestCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LeastCommonMultiple(a, b int, integers ...int) int {
	result := a * b / GreatestCommonDivisor(a, b)

	for i := 0; i < len(integers); i++ {
		result = LeastCommonMultiple(result, integers[i])
	}

	return result
}
