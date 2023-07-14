package main

import (
	"fmt"
	"math"
)

// A, B - sides
type Rectangle struct {
	A, B int
}

func NewRectangle(A, B int) Rectangle {
	return Rectangle{A, B}
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle. Sides: %d and %d", r.A, r.B)
}

func (r Rectangle) Square() float64 {
	return float64(r.A * r.B)
}

func (r Rectangle) Length() float64 {
	return float64(r.A*2 + r.B*2)
}

func (r Rectangle) Diagonal() float64 {
	return math.Sqrt(float64(r.A*r.A + r.B*r.B))
}
