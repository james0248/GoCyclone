package include

import (
	"math"
)

// Vector represents a 3D vector
type Vector struct {
	x, y, z float64
	padding float64
}

// NewZeroVector returns a pointer to a default 0 vector
func NewZeroVector() *Vector {
	return NewVector(0, 0, 0)
}

// NewVector returns a pointer to a new 3D vector
func NewVector(x, y, z float64) *Vector {
	return &Vector{
		x: x,
		y: y,
		z: z,
	}
}

// InvertVector returns a vector with a same size and opposite direction
func (v Vector) InvertVector() *Vector {
	return &Vector{
		x: -v.x,
		y: -v.y,
		z: -v.z,
	}
}

//Size returns the magnitude of the vector
func (v Vector) Size() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

//SquareSize returns the square of a magnitude of the vector
func (v Vector) SquareSize() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

//NormVector returns a pointer to a normalized vector
func (v Vector) NormVector() *Vector {
	size := v.Size()
	if size == 0 {
		panic("Cannot normalize a vector with size of 0")
	}
	return &Vector{
		x: v.x / size,
		y: v.y / size,
		z: v.z / size,
	}
}

// MultiplyScalar returns a pointer to a vector multiplied by a given value
func (v Vector) MultiplyScalar(value float64) *Vector {
	return &Vector{
		x: v.x * value,
		y: v.y * value,
		z: v.z * value,
	}
}

// AddVector returns a pointer to a vector added the given vector
func (v Vector) AddVector(u Vector) *Vector {
	return &Vector{
		x: v.x + u.x,
		y: v.y + u.y,
		z: v.z + u.z,
	}
}

// SubtractVector returns a pointer to a vector subtracted the given vector
func (v Vector) SubtractVector(u Vector) *Vector {
	return &Vector{
		x: v.x - u.x,
		y: v.y - u.y,
		z: v.z - u.z,
	}
}

// AddScaledVector returns a pointer to a vector added a given vector multiplied by the given value
func (v Vector) AddScaledVector(u Vector, value float64) *Vector {
	return &Vector{
		x: v.x + u.x*value,
		y: v.y + u.y*value,
		z: v.z + u.z*value,
	}
}
