package util

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

func PointsBetween(start *Point, end *Point) []*Point {
	points := make([]*Point, 0)
	if start.X == end.X {
		yRange := IntAbs(start.Y - end.Y)
		y := IntMin(start.Y, end.Y)
		for i := y; i <= y+yRange; i++ {
			points = append(points, &Point{X: start.X, Y: i})
		}
	} else if start.Y == end.Y {
		xRange := IntAbs(start.X - end.X)
		x := IntMin(start.X, end.X)
		for i := x; i <= x+xRange; i++ {
			points = append(points, &Point{X: i, Y: start.Y})
		}
	}
	return points
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
