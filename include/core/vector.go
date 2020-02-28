package core

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
func (v Vector) AddVector(u *Vector) *Vector {
	return &Vector{
		x: v.x + u.x,
		y: v.y + u.y,
		z: v.z + u.z,
	}
}

// SubtractVector returns a pointer to a vector subtracted the given vector
func (v Vector) SubtractVector(u *Vector) *Vector {
	return &Vector{
		x: v.x - u.x,
		y: v.y - u.y,
		z: v.z - u.z,
	}
}

// AddScaledVector returns a pointer to a vector added a given vector multiplied by the given value
func (v Vector) AddScaledVector(u *Vector, value float64) *Vector {
	return &Vector{
		x: v.x + u.x*value,
		y: v.y + u.y*value,
		z: v.z + u.z*value,
	}
}

// ComponentProduct returns a vector multiplied each of the given vector's entry
func (v Vector) ComponentProduct(u *Vector) *Vector {
	return &Vector{
		x: v.x * u.x,
		y: v.y * u.y,
		z: v.z * u.z,
	}
}

// ComponentProductUpdate updates the vector by multiplying the given vector's entry
func (v *Vector) ComponentProductUpdate(u *Vector) {
	v.x *= u.x
	v.y *= u.y
	v.z *= u.z
}

// InnerProduct returns the inner product value of two vectors
func (v Vector) InnerProduct(u *Vector) float64 {
	return v.x*u.x + v.y*u.y + v.z*u.z
}

// CrossProduct returns a vector as a result of the cross product of two vectors.
// The calculating order is v.CrossProduct(u) = v X u
func (v Vector) CrossProduct(u *Vector) *Vector {
	return &Vector{
		x: v.y*u.z - v.z*u.y,
		y: v.z*u.x - v.x*u.z,
		z: v.x*u.y - v.y*u.x,
	}
}

// MakeOrthonormalBasis returns a orthonomal basisbased on the given vector
func MakeOrthonormalBasis(v, u, w *Vector) (*Vector, *Vector, *Vector, error) {
	// TODO
	return &Vector{}, &Vector{}, &Vector{}, nil
}

// Clear makes a vector 0
func (v *Vector) Clear() {
	v.x = 0
	v.y = 0
	v.z = 0
}
