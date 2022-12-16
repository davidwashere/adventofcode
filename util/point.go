package util

import (
	"fmt"
	"math"
)

// Represents a position in 2d coordinate space
type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}

// DistOrtho calculates the number of 'steps' between two points
// where horizontal, vertical, and diagonal moves are allowed
//
// Also called: Manhattan Distance and taxicab distance
func (p Point) DistOrtho(p2 Point) int {
	return DistOrtho(p.X, p.Y, p2.X, p2.Y)
}

// DistOrthoDiag calculates the number of 'steps' between two points
// where horizontal, vertical, and diagonal moves are allowed
func (p Point) DistOrthoDiag(p2 Point) int {
	return DistOrthoDiag(p.X, p.Y, p2.X, p2.Y)
}

// Dist calculates the distance between two points 'as the bird flys'
func (p Point) Dist(p2 Point) float64 {
	return Dist(p.X, p.Y, p2.X, p2.Y)
}

func (p Point) Apply(v Vector) Point {
	var np Point
	np.X, np.Y = v.Apply(p.X, p.Y)
	return np
}

func (p Point) String() string {
	return fmt.Sprintf("[%v,%v]", p.X, p.Y)
}

// TowardVector will return a normalized vector in the semi-direction of the
// dest point.
func (p Point) TowardVector(dest Point) Vector {
	v := NewNormalizedVector(0, 0)
	if p.X != dest.X {
		if dest.X > p.X {
			v.X = 1
		} else {
			v.X = -1
		}
	}

	if p.Y != dest.Y {
		if dest.Y > p.Y {
			v.Y = 1
		} else {
			v.Y = -1
		}
	}

	return v
}

// DistOrthoDiag calculates the number of 'steps' between two points
// where horizontal, vertical, and diagonal moves are allowed
func DistOrthoDiag(x1, y1, x2, y2 int) int {
	return Max(Abs(x1-x2), Abs(y1-y2))
}

// DistOrtho calculates the number of 'steps' between two points
// where horizontal, vertical, and diagonal moves are allowed
//
// Also called: Manhattan Distance and taxicab distance
func DistOrtho(x1, y1, x2, y2 int) int {
	h := Abs(x1 - x2)
	v := Abs(y1 - y2)
	return h + v
}

// Dist calculates the distance between two points 'as the bird flys'
func Dist(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))
}
