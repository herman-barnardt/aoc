package util

import (
	"math"
)

type Point struct {
	X int
	Y int
}

func (p Point) Add(q Point) Point {
	return Point{X: p.X + q.X, Y: p.Y + q.Y}
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

func MinMax(points []Point) (Point, Point) {
	min, max := Point{X: math.MaxInt, Y: math.MaxInt}, Point{X: math.MinInt, Y: math.MinInt}
	for _, p := range points {
		min.X, min.Y = IntMin(min.X, p.X), IntMin(min.Y, p.Y)
		max.X, max.Y = IntMax(max.X, p.X), IntMax(max.Y, p.Y)
	}
	return min, max
}

func DistanceBetween(a *Point, b *Point) int {
	return IntAbs(a.X-b.X) + IntAbs(a.Y-b.Y)
}
