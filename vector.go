package main

import "math"

type Vector struct {
	x, y, z float64
}

func (v Vector) Add(a Vector) Vector {
	return Vector{v.x + a.x, v.y + a.y, v.z + a.z}
}

func (v Vector) Subtract(a Vector) Vector {
	return Vector{v.x - a.x, v.y - a.y, v.z - a.z}
}

func (v Vector) Multiply(a Vector) Vector {
	return Vector{v.x * a.x, v.y * a.y, v.z * a.z}
}

func (v Vector) Divide(a Vector) Vector {
	return Vector{v.x / a.x, v.y / a.y, v.z / a.z}
}

func (v Vector) AddScalar(a float64) Vector {
	return Vector{v.x + a, v.y + a, v.z + a}
}

func (v Vector) SubtractScalar(a float64) Vector {
	return Vector{v.x - a, v.y - a, v.z - a}
}

func (v Vector) MultiplyScalar(a float64) Vector {
	return Vector{v.x * a, v.y * a, v.z * a}
}

func (v Vector) DivideScalar(a float64) Vector {
	return Vector{v.x / a, v.y / a, v.z / a}
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v Vector) SquaredLength() float64 {
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func (v Vector) Normalise() Vector {
	return v.DivideScalar(v.Length())
}

func (v Vector) Dot(a Vector) float64 {
	return v.x*a.x + v.y*a.y + v.z*a.z
}

func (v Vector) Cross(a Vector) Vector {
	return Vector{
		x: v.y*a.z - v.z*a.y,
		y: v.z*a.x - v.x*a.z,
		z: v.x*a.y - v.y*a.x,
	}
}
