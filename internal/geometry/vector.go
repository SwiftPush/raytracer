package geometry

import "math"

type Vector struct {
	X, Y, Z float64
}

func (v Vector) Add(a Vector) Vector {
	return Vector{v.X + a.X, v.Y + a.Y, v.Z + a.Z}
}

func (v Vector) Subtract(a Vector) Vector {
	return Vector{v.X - a.X, v.Y - a.Y, v.Z - a.Z}
}

func (v Vector) Multiply(a Vector) Vector {
	return Vector{v.X * a.X, v.Y * a.Y, v.Z * a.Z}
}

func (v Vector) Divide(a Vector) Vector {
	return Vector{v.X / a.X, v.Y / a.Y, v.Z / a.Z}
}

func (v Vector) AddScalar(a float64) Vector {
	return Vector{v.X + a, v.Y + a, v.Z + a}
}

func (v Vector) SubtractScalar(a float64) Vector {
	return Vector{v.X - a, v.Y - a, v.Z - a}
}

func (v Vector) MultiplyScalar(a float64) Vector {
	return Vector{v.X * a, v.Y * a, v.Z * a}
}

func (v Vector) DivideScalar(a float64) Vector {
	return Vector{v.X / a, v.Y / a, v.Z / a}
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector) SquaredLength() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vector) Normalise() Vector {
	return v.DivideScalar(v.Length())
}

func (v Vector) Dot(a Vector) float64 {
	return v.X*a.X + v.Y*a.Y + v.Z*a.Z
}

func (v Vector) Cross(a Vector) Vector {
	return Vector{
		X: v.Y*a.Z - v.Z*a.Y,
		Y: v.Z*a.X - v.X*a.Z,
		Z: v.X*a.Y - v.Y*a.X,
	}
}
