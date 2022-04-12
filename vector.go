package main

import "math"

type Vector struct {
	X float64
	Y float64
}

func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{X: v1.X + v2.X, Y: v1.Y + v2.Y}
}

func (v1 Vector) Subtract(v2 Vector) Vector {
	return Vector{X: v1.X - v2.X, Y: v1.Y - v2.Y}
}

func (v1 Vector) Multiply(v2 Vector) Vector {
	return Vector{X: v1.X * v2.X, Y: v1.Y * v2.Y}
}

func (v1 Vector) AddV(d float64) Vector {
	return Vector{X: v1.X + d, Y: v1.Y + d}
}

func (v1 Vector) SubtractV(d float64) Vector {
	return Vector{X: v1.X - d, Y: v1.Y - d}
}

func (v1 Vector) MultiplyV(d float64) Vector {
	return Vector{X: v1.X * d, Y: v1.Y * d}
}

func (v1 Vector) DivisionV(d float64) Vector {
	return Vector{X: v1.X / d, Y: v1.Y / d}
}

func (v1 Vector) Limit(lower, upper float64) Vector {
	return Vector{
		X: math.Min(math.Max(v1.X, lower), upper),
		Y: math.Min(math.Max(v1.Y, lower), upper),
	}
}

func (v1 Vector) Distance(v2 Vector) float64 {
	return math.Sqrt(math.Pow(v2.X-v1.X, 2) + math.Pow(v2.Y-v1.Y, 2))
}
