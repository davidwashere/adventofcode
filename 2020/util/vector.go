package util

import (
	"math"
)

// Vector holds direction and magnitude
type Vector struct {
	X int
	Y int
	M int
}

// NewVector creates a vector
func NewVector(x, y, magnitude int) Vector {
	return Vector{
		X: x,
		Y: y,
		M: magnitude,
	}
}

// NewNormalizedVector creates a vector with magnitude 1
func NewNormalizedVector(x, y int) Vector {
	return Vector{
		X: x,
		Y: y,
		M: 1,
	}
}

// Apply will apply this vector transform to the given coords
func (v *Vector) Apply(x, y int) (tx int, ty int) {
	tx = (v.X * v.M) + x
	ty = (v.Y * v.M) + y

	return tx, ty
}

// Rotate will rotate the vector by the given # of degrees, a positive degree
// will rotate the vector counter-clockwise, a negative degree will rotate clockwise
func (v *Vector) Rotate(deg float64) {
	rad := deg * (math.Pi / 180.0)

	ca := math.Cos(rad)
	sa := math.Sin(rad)
	tx := math.Round(ca*float64(v.X) - sa*float64(v.Y))
	ty := math.Round(sa*float64(v.X) + ca*float64(v.Y))

	v.X = int(tx)
	v.Y = int(ty)
}

// RotateInt same as Rotate but takes an int
func (v *Vector) RotateInt(deg int) {
	v.Rotate(float64(deg))
}
